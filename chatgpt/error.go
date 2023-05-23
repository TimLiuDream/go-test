package main

import (
	"encoding/json"
	goError "errors"
	"fmt"

	"github.com/bangwork/wiki-api/app/middlewares/log"
	"github.com/bangwork/wiki-api/app/utils/errors"
)

// APIError provides error information returned by the OpenAI API.
type APIError struct {
	Code           int     `json:"code,omitempty"`
	Message        string  `json:"message"`
	Param          *string `json:"param,omitempty"`
	Type           string  `json:"type"`
	HTTPStatusCode int     `json:"-"`
}

// RequestError provides informations about generic request errors.
type RequestError struct {
	HTTPStatusCode int
	Err            error
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	if e.HTTPStatusCode > 0 {
		return fmt.Sprintf("error, status code: %d, message: %s", e.HTTPStatusCode, e.Message)
	}

	return e.Message
}

func (e *APIError) UnmarshalJSON(data []byte) (err error) {
	var rawMap map[string]json.RawMessage
	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawMap["message"], &e.Message)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawMap["type"], &e.Type)
	if err != nil {
		return
	}

	// optional fields
	if _, ok := rawMap["param"]; ok {
		err = json.Unmarshal(rawMap["param"], &e.Param)
		if err != nil {
			return
		}
	}

	if _, ok := rawMap["code"]; !ok {
		return nil
	}

	// if the api returned a number, we need to force an integer
	// since the json package defaults to float64
	var intCode int
	err = json.Unmarshal(rawMap["code"], &intCode)
	if err == nil {
		e.Code = intCode
		return nil
	}

	return json.Unmarshal(rawMap["code"], &e.Code)
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("error, status code: %d, message: %s", e.HTTPStatusCode, e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

func parseError(err error) error {
	e := &APIError{}
	if goError.As(err, &e) {
		switch e.HTTPStatusCode {
		case 401:
			// invalid auth or key (do not retry)
			log.Error("invalid auth or key, err: %+v", err)
		case 429:
			// rate limiting or engine overload (wait and retry)
			log.Error("rate limiting or engine overload, err: %+v", err)
		case 500:
			// openai server error (retry)
			log.Error("openai server error, err: %+v", err)
		default:
			// unhandled
			log.Error("unhandled error, err: %+v", err)
		}
		return errors.Trace(err)
	}
	return nil
}
