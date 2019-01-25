package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Health Check Api
 */
func health(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
