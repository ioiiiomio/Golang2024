package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ioiiiomio/Golang2024/handlers"
)

func SetupRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.GET("", handlers.GetTasks)
		taskRoutes.POST("", handlers.CreateTask)
		taskRoutes.DELETE(":id", handlers.DeleteTask)
		taskRoutes.PUT(":id", handlers.UpdateTask)
	}
}
