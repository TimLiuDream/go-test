package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(rw, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(rw, r)
	})
}
