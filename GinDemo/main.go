package main

import (
	"GolangTraining/GinDemo/Model"
	"fmt"
	"net/http"

	_ "GolangTraining/GinDemo/docs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//字段名一定要大写，不然不能导出，那就不能被请求体的json反序列化
type TestJson struct {
	Para int `json:"para" binding:"required"`
}

var db *gorm.DB
var err error

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func main() {
	db, err = gorm.Open("mysql", "root:lft8306488@tcp(localhost:3306)/todo_list_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("orm is disconnect")
		return
	}

	db.AutoMigrate(&Model.Test{})

	r := gin.Default()
	r.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/change", changePost)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

// @Summary 测试post
// @Accept  json
// @Produce  json
// @Param Para body TestJson true "para"
// @Success 200 {string} json "{"message": 1991}"
// @Router /change [post]
func changePost(context *gin.Context) {
	var test TestJson
	if err := context.BindJSON(&test); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if test.Para == 1 {

		var testM Model.Test
		db.First(&testM)
		fmt.Println(testM.Year)

		context.JSON(http.StatusOK, gin.H{
			"message": testM.Year,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "false",
		})
	}
}
