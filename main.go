package main

import (
	"context"
	"fmt"

	"example/test-items-service/handlers"
	"example/test-items-service/properties"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test-items-service greeting...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	props := properties.GetProperties()

	router := gin.Default()

	testCaseHandler := handlers.NewTestCaseHandler(ctx)

	router.GET("/test_case", testCaseHandler.GetTestCases)
	router.GET("/test_case/:id", testCaseHandler.GetTestCaseById)
	router.POST("/test_case", testCaseHandler.PostTestCase)
	router.PUT("/test_case", testCaseHandler.PutTestCase)
	router.DELETE("/test_case", testCaseHandler.DeleteTestCase)

	router.Run(fmt.Sprintf("%s:%s", props.App.Host, props.App.Port))
}
