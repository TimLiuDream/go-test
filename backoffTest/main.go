package main

import (
	"fmt"

	"github.com/cenkalti/backoff"
)

func main() {
	// An operation that may fail.
	operation := func() error {
		return fmt.Errorf("error hanpen")
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		panic(err)
	}
}
