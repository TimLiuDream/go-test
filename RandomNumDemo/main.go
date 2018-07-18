package main

import (
	"time"
	"fmt"
	"math/rand"
)

func randInt(min int , max int) string {
	rand.Seed(time.Now().UnixNano())
	randomNumber :=fmt.Sprintf("%d", min + rand.Intn(max-min)) 
	fmt.Println("randomNum is ",randomNumber)
	return randomNumber
}

func main()  {
	randInt(999,9999)
}