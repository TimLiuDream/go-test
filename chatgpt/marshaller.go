package main

import "encoding/json"

type marshaller interface {
	marshal(value interface{}) ([]byte, error)
}

type jsonMarshaller struct{}

func (jm *jsonMarshaller) marshal(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}
