package handlers

import (
	"fmt"
	"net/http"

	"example/test-items-service/services"

	"github.com/gin-gonic/gin"
)

func GetTestCases(ctx *gin.Context) {
	testCaseService := *services.NewTestCaseService(ctx)
	fmt.Println("Getting Test Cases...")
	cases, err := testCaseService.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	ctx.IndentedJSON(http.StatusOK, cases)
}

func GetTestCase(ctx *gin.Context) {
	testCaseService := *services.NewTestCaseService(ctx)

	id := ctx.Param("id")
	fmt.Println("Getting Test Case with id = " + id)
	testCase, err := testCaseService.FindOneById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	ctx.IndentedJSON(http.StatusOK, testCase)
}

func PostTestCase(ctx *gin.Context) {
	fmt.Println("Test Case created...")

	ctx.IndentedJSON(http.StatusOK, "Success")
}

func PutTestCase(ctx *gin.Context) {
	fmt.Println("Test Case updated...")

	ctx.IndentedJSON(http.StatusOK, "Success")
}

func DeleteTestCase(ctx *gin.Context) {
	fmt.Println("Test Case deleted...")

	ctx.IndentedJSON(http.StatusOK, "Success")
}
