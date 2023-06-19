package main

import (
	"net/http"
	"testing"
)

// https://blog.devgenius.io/atomic-pointer-in-golang-1-19-and-dependency-injection-ad090fbb597d
func TestClientDo(t *testing.T) {
	// given
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log := logger.WithFields(logrus.Fields{"!TestCase": t.Name()})
	cleanupHttpDebug := RegisterHttpTracer(log)
	defer cleanupHttpDebug()
	testURL := "http://google.com"
	// when
	resp, err := ClientDo(http.MethodGet, testURL)
	// then
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
