package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300

	azureAPIPrefix         = "openai"
	azureDeploymentsPrefix = "deployments"
)

type APIType string

const (
	APITypeOpenAI  APIType = "OPEN_AI"
	APITypeAzure   APIType = "AZURE"
	APITypeAzureAD APIType = "AZURE_AD"
)

const AzureAPIKeyHeader = "api-key"

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL    string
	OrgID      string
	APIType    APIType
	APIVersion string // required when APIType is APITypeAzure or APITypeAzureAD
	Engine     string // required when APIType is APITypeAzure or APITypeAzureAD

	HTTPClient *http.Client

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken: authToken,
		BaseURL:   openaiAPIURLv1,
		APIType:   APITypeOpenAI,
		OrgID:     "",

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

// Client is OpenAI GPT-3 API client.
type Client struct {
	config ClientConfig

	requestBuilder    requestBuilder
	createFormBuilder func(io.Writer) formBuilder
}

// NewClient creates new OpenAI API client.
func NewClient(authToken string) *Client {
	config := DefaultConfig(authToken)
	return NewClientWithConfig(config)
}

// NewClientWithConfig creates new OpenAI API client for specified config.
func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: newRequestBuilder(),
		createFormBuilder: func(body io.Writer) formBuilder {
			return newFormBuilder(body)
		},
	}
}

func (c *Client) fullURL(suffix string) string {
	// /openai/deployments/{engine}/chat/completions?api-version={api_version}
	if c.config.APIType == APITypeAzure || c.config.APIType == APITypeAzureAD {
		baseURL := c.config.BaseURL
		baseURL = strings.TrimRight(baseURL, "/")
		// if suffix is /models change to {endpoint}/openai/models?api-version=2022-12-01
		// https://learn.microsoft.com/en-us/rest/api/cognitiveservices/azureopenaistable/models/list?tabs=HTTP
		if strings.Contains(suffix, "/models") {
			return fmt.Sprintf("%s/%s%s?api-version=%s", baseURL, azureAPIPrefix, suffix, c.config.APIVersion)
		}
		return fmt.Sprintf("%s/%s/%s/%s%s?api-version=%s",
			baseURL, azureAPIPrefix, azureDeploymentsPrefix, c.config.Engine, suffix, c.config.APIVersion)
	}

	// c.config.APIType == APITypeOpenAI || c.config.APIType == ""
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	// Azure API Key authentication
	if c.config.APIType == APITypeAzure {
		req.Header.Set(AzureAPIKeyHeader, c.config.authToken)
	} else {
		// OpenAI or Azure AD authentication
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.authToken))
	}

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if len(c.config.OrgID) > 0 {
		req.Header.Set("OpenAI-Organization", c.config.OrgID)
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return c.handleErrorResp(res)
	}

	return decodeResponse(res.Body, v)
}

func decodeResponse(body io.Reader, v interface{}) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return decodeString(body, result)
	}
	return json.NewDecoder(body).Decode(v)
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := &RequestError{
			HTTPStatusCode: resp.StatusCode,
			Err:            err,
		}
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	return errRes.Error
}
