package handlers

import (
	"fmt"
	"log"
	"net/http"

	testItems "example/test-items-service/models"
	"example/test-items-service/services"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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
	testCaseService := *services.NewTestCaseService(ctx)
	var tc testItems.TestCase
	if err := ctx.BindJSON(&tc); err != nil {
		log.Println("Binding is ok!")
	}

	if err := testCaseService.AddOne(tc); err != nil {
		log.Println("Test Case not created...")
		ctx.IndentedJSON(http.StatusNotModified, gin.H{"data": err.Error()})
	}
	log.Println("Test Case created...")
	ctx.IndentedJSON(http.StatusOK, gin.H{"data": tc})
}

func PutTestCase(ctx *gin.Context) {
	testCaseService := *services.NewTestCaseService(ctx)

	tcId, _ := uuid.FromString(ctx.Param("id"))
	var tc testItems.TestCase
	if err := ctx.BindJSON(&tc); err != nil {
		log.Println("Binding is ok!")
	}
	tc.ID = tcId
	res, err := testCaseService.UpdateOneById(tc)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Test Case updated...")
	ctx.IndentedJSON(http.StatusOK, gin.H{"data": res})
}

func DeleteTestCase(ctx *gin.Context) {
	testCaseService := *services.NewTestCaseService(ctx)

	id := ctx.Param("id")
	log.Println("id: " + id)
	if err := testCaseService.DeleteOneById(id); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"data": "Test case deleted with id: " + id})
}
