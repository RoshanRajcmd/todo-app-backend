package routes

import (
	"github.com/RoshanRajcmd/todo-app-backend/controllers"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	var userGroup = r.Group("/todos")
	{
		// Task Routes
		//Create a Task
		userGroup.POST("/createTask", controllers.CreateTask)
		// Read all
		userGroup.GET("/getAllTasks", controllers.GetAllTasks)
		// Read One
		userGroup.GET("/getTaskById/:id", controllers.GetTaskById)
		// Update
		userGroup.PUT("/updateTask/:id", controllers.UpdateTask)
		// Delete
		userGroup.DELETE("/deleteTask/:id", controllers.DeleteTask)
	}
}
