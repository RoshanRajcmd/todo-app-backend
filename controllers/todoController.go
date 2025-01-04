package controllers

import (
	"github.com/RoshanRajcmd/todo-app-backend/initializers"
	"github.com/RoshanRajcmd/todo-app-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	// Get data from req body
	var body struct {
		content string
		isRead  bool
	}
	c.Bind(&body)

	// Create a todo
	//var task = models.Task{Content: body.content, IsRead: body.isRead}
	var task = models.NewTask(body.content, body.isRead)
	var result = initializers.DB.Create(&task)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"data":    task,
		"message": "Task Created Successfully",
	})
}

func GetAllTasks(c *gin.Context) {
	// Get all the tasks
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	// Return todos in response
	c.JSON(200, gin.H{
		"data":    tasks,
		"message": "Fetched Successfully",
	})
}

func GetTaskById(c *gin.Context) {
	// Get id from URL param
	var taskId = c.Param("id")

	// Get a get the task todo
	var task models.Task
	initializers.DB.First(&task, taskId)

	// Return todo in response
	c.JSON(200, gin.H{
		"data":    task,
		"message": "Fetched Successfully",
	})
}

func UpdateTask(c *gin.Context) {
	// Get id from URL param
	var id = c.Param("id")

	// get the data of req body
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	// Get a single todo that we want to update
	var todo models.Task
	initializers.DB.First(&todo, id)

	// Update it
	initializers.DB.Model(&todo).Updates(models.Task{
		Content: body.Content,
		IsRead:  body.Status,
	})

	// Return response
	c.JSON(200, gin.H{
		"data":    todo,
		"message": "Updated Task Successfully",
	})
}

func DeleteTask(c *gin.Context) {
	// Get id from URL param
	var id = c.Param("id")

	// Delete the Todo
	var response = initializers.DB.Delete(&models.Task{}, id)

	if response.Error != nil {
		// if dbErr, ok := response.Error.(); ok {
		// 	switch dbErr.Number {
		// 	case 1062: // MySQL code for duplicate entry
		// 		// Handle duplicate entry
		// 	// Add cases for other specific error codes
		// 	default:
		// 		// Handle other errors
		// 	}
		// } else {
		// 	// Handle non-MySQL errors or unknown errors
		// }
	} else {
		// Return response
		c.JSON(200, gin.H{
			"message": "Task removed Successfully",
		})
	}
}
