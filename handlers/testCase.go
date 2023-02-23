package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTestCase(c *gin.Context) {
	fmt.Println("Getting  Test Case...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func PostTestCase(c *gin.Context) {
	fmt.Println("Test Case created...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func PutTestCase(c *gin.Context) {
	fmt.Println("Test Case updated...")

	c.IndentedJSON(http.StatusOK, "Success")
}

func DeleteTestCase(c *gin.Context) {
	fmt.Println("Test Case deleted...")

	c.IndentedJSON(http.StatusOK, "Success")
}
