package handlers

import (
	"context"
	"fmt"
	"net/http"

	"example/test-items-service/services"
	"github.com/gin-gonic/gin"
)

type TestCaseHandler struct {
	testCaseService services.TestCaseService
}

func NewTestCaseHandler(ctx context.Context) *TestCaseHandler {
	return &TestCaseHandler{*services.NewTestCaseService(ctx)}
}

func (h *TestCaseHandler) GetTestCases(c *gin.Context) {
	fmt.Println("Getting Test Cases...")
	cases, err := h.testCaseService.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	c.IndentedJSON(http.StatusOK, cases)
}

func (h *TestCaseHandler) GetTestCaseById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Getting Test Case with id = " + id)
	testCase, err := h.testCaseService.FindOneById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	c.IndentedJSON(http.StatusOK, testCase)
}

func (h *TestCaseHandler) PostTestCase(c *gin.Context) {
	fmt.Println("Test Case created...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func (h *TestCaseHandler) PutTestCase(c *gin.Context) {
	fmt.Println("Test Case updated...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func (h *TestCaseHandler) DeleteTestCase(c *gin.Context) {
	fmt.Println("Test Case deleted...")

	c.IndentedJSON(http.StatusOK, "Success")
}
