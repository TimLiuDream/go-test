package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")

	//docs.SwaggerInfo.Title = "maintenance"

	r := gin.Default()

	//routers.SetupRouter(r)

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":" + "10007"); err != nil {
		fmt.Printf("startup server failed, err: %v", err)
	}
}

func init() {
	fmt.Println("main init start")

	fmt.Println("init global settings")
	//err := setupSetting()
	//
	//if err != nil {
	//	log.Fatalf("init.setupSetting err: %v", err)
	//}
}
