package main

import (
	"fmt"
	"github.com/imroc/req"
	"log"
	"net/http"
)

func main() {
	resp, err := req.Get("https://httpbin.org/uuid1")
	if err != nil {
		log.Fatalln(err)
	}
	statusCode := resp.Response().StatusCode
	if statusCode != http.StatusOK {
		fmt.Printf("status: %d, resp: %s\n", statusCode, resp.String())
		return
	}
	str, err := resp.ToString()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(str)
}
