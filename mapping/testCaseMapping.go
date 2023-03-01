package mapping

import (
	"context"

	"example/test-items-service/handlers"

	"github.com/gin-gonic/gin"
)

const (
	RootTestCasePath = "/test_case"
)

func ApplyTestCaseMapping(ctx context.Context, router *gin.Engine) {

	api := router.Group(RootTestCasePath)
	{
		api.GET("", handlers.GetTestCases)
		api.GET(":id", handlers.GetTestCase)
		api.POST("", handlers.PostTestCase)
		api.PUT(":id", handlers.PutTestCase)
		api.DELETE(":id", handlers.DeleteTestCase)
	}
}
