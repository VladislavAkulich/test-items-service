package main

import (
	"context"
	"example/test-items-service/mapping"
	"fmt"
	"github.com/gin-gonic/gin"

	"example/test-items-service/properties"
)

func main() {
	fmt.Println("test-items-service greeting...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	router := gin.Default()
	mapping.ApplyTestCaseMapping(ctx, router)

	props := properties.GetProperties()
	router.Run(fmt.Sprintf("%s:%s", props.App.Host, props.App.Port))
}
