package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync/atomic"
)

// defaultHttpClientTransport is used as http.Client.Transport
// for next outgoing requests.
var defaultHttpClientTransport atomic.Pointer[http.RoundTripper]

// init sets the default http.Transport
func init() {
	defaultHttpClientTransport.Store(&http.DefaultTransport)
}

// ClientDo uses defaultHttpClientTransport to make outgoing http requests
func ClientDo(method, url string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(context.Background(), method, url, nil)

	client := http.Client{Transport: *defaultHttpClientTransport.Load()}

	return client.Do(req)
}

// HttpTracer is a client middleware to log requests
type HttpTracer struct {
	orig *http.RoundTripper
	log  *logrus.Entry
}

// RoundTrip implements http.RoundTripper, logs debug info.
func (tr *HttpTracer) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := (*tr.orig).RoundTrip(req)

	fields := logrus.Fields{
		"ReqMethod":  req.Method,
		"ReqURL":     req.URL,
		"ReqHeaders": req.Header,
	}
	if resp != nil {
		fields["RespStatusCode"] = resp.StatusCode
		fields["RespHeaders"] = resp.Header
	}
	tr.log.WithFields(fields).Debug("Client request")

	return resp, err
}

// RegisterHttpTracer registers HttpTracer as a client middleware.
// Returns the deregister function.
func RegisterHttpTracer(log *logrus.Entry) func() {
	httpTracer := &HttpTracer{defaultHttpClientTransport.Load(), log}
	var rt http.RoundTripper = httpTracer

	defaultHttpClientTransport.Store(&rt)

	return func() {
		defaultHttpClientTransport.Store(httpTracer.orig)
	}
}
