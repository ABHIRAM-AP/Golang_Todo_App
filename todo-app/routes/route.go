package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/tasks", controllers.GetTasks)
	route.POST("/tasks", controllers.CreateTask)
	route.DELETE("/tasks/:id", controllers.DeleteTask)

	return route
}
