package routes

import (
	"github.com/RoshanRajcmd/todo-app-backend/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	var userGroup = r.Group("/todos")
	{
		// Todo Routes
		userGroup.POST("", controllers.TodosCreate)       // Create
		userGroup.GET("", controllers.TodosIndex)         // Read all
		userGroup.GET("/:id", controllers.TodosShow)      // Read One
		userGroup.PUT("/:id", controllers.TodosUpdate)    // Update
		userGroup.DELETE("/:id", controllers.TodosDelete) // Delete
	}
}
