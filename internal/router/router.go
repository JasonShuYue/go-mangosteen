package router

import (
	"go-mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", controller.Ping)

	return r
}
