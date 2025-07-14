package routes

import (
	"fmt"
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	route := gin.Default()

	fmt.Println(">>> ROUTES INITIALIZED") // 👈 Add this

	route.GET("/tasks", controllers.GetTasks)
	route.POST("/tasks", controllers.CreateTask)
	route.DELETE("/tasks/:id", controllers.DeleteTask)
	route.PUT("/tasks/:id", controllers.UpdateTask)
	route.PATCH("/tasks/:id/complete", controllers.MarkComplete)
	route.PATCH("/tasks/:id/incomplete", controllers.MarkIncomplete)

	return route
}
