package main

import (
	"fmt"

	"example/test-items-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	props := GetProperties()
	fmt.Println("test-items-service greeting...")

	router := gin.Default()

	router.GET("/test_case", handlers.GetTestCase)
	router.POST("/test_case", handlers.PostTestCase)
	router.PUT("/test_case", handlers.PutTestCase)
	router.DELETE("/test_case", handlers.DeleteTestCase)

	router.Run(fmt.Sprintf("%s:%s", props.App.Host, props.App.Port))
}
