// client.go

package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	endpoint := "http://localhost:15001/api/v1/get_something?name=something"
	totalRequests := 5

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			makeAPICall(endpoint)
		}(i)
	}
	wg.Wait()

}

func makeAPICall(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	result := string(body)
	log.Printf(result)
}
