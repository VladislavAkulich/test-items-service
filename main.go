package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type test_case struct {
	ID             uuid.UUID
	NAME           string
	STEPS          []string
	PRE_CONDITOINS string
	Author         string
}

type check_list struct {
	ID             uuid.UUID
	NAME           string
	STEPS          []string
	PRE_CONDITOINS string
	Author         string
}

type test_suite struct {
	ID         uuid.UUID
	NAME       string
	TEST_CASES []string
	Author     string
}

func main() {
	fmt.Println("test-items-service greeting...")

	router := gin.Default()

	router.GET("/test_case", getTestCase)
	router.POST("/test_case", postTestCase)
	router.PUT("/test_case", putTestCase)
	router.DELETE("/test_case", delteTestCase)

	router.Run("localhost:8080")
}

func getTestCase(c *gin.Context) {
	fmt.Println("Getting  Test Case...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func postTestCase(c *gin.Context) {
	fmt.Println("Test Case created...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func putTestCase(c *gin.Context) {
	fmt.Println("Test Case updated...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func delteTestCase(c *gin.Context) {
	fmt.Println("Test Case deleted...")

	c.IndentedJSON(http.StatusOK, "Success")
}
