package main

import (
	"github.com/cxr-inc/charting_apis/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * main
 */
func main() {
	router := gin.Default()

	router.GET("/", controller.health)
	router.GET("/health", controller.health)

	// api v1
	apiv1 := router.Group("/v1")
	{
		apiv1.GET("/example", controller.getExample)
		apiv1.POST("/study_templates", controller.postExample)
	}

	router.Run(":3003")
}
