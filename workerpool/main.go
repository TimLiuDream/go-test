package main

import (
	"fmt"

	"github.com/timliudream/go-test/workerpool/model"
	"github.com/timliudream/go-test/workerpool/worker"
)

func main() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 1000; i++ {
		data := model.SimpleData{ID: i}
		allData = append(allData, data)
	}
	fmt.Printf("Start processing all work \n")

	// Process
	worker.PooledWorkError(allData)
}
