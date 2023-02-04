package routes

import (
	"git-example/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	group := r.Group("v1")
	group.GET("/health", handlers.HealthHandler)
	group.GET("/sum", handlers.SumHandler)
}
