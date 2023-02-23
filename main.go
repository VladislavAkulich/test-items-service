package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type testCase struct {
	ID            uuid.UUID
	Name          string
	Steps         []string
	PreConditions string
	Author        string
}

type checkList struct {
	ID            uuid.UUID
	Name          string
	Steps         []string
	PreConditions string
	Author        string
}

type testSuite struct {
	ID        uuid.UUID
	NAME      string
	TestCases []string
	Author    string
}

func main() {
	fmt.Println("test-items-service greeting...")

	router := gin.Default()

	router.GET("/test_case", getTestCase)
	router.POST("/test_case", postTestCase)
	router.PUT("/test_case", putTestCase)
	router.DELETE("/test_case", deleteTestCase)

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

func deleteTestCase(c *gin.Context) {
	fmt.Println("Test Case deleted...")

	c.IndentedJSON(http.StatusOK, "Success")
}
