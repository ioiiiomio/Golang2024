package routes

import (
	handlers "github.com/ioiiiomio/Golang2024/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
}
