package controllers

import (
	"net/http"

	"github.com/RoshanRajcmd/todo-app-backend/initializers"
	"github.com/RoshanRajcmd/todo-app-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	// Get data from req body
	var body struct {
		Content string
		IsRead  bool
	}
	c.Bind(&body)

	// Create a Task
	//var task = models.Task{Content: body.Content, IsRead: body.IsRead}
	var task = models.NewTask(body.Content, body.IsRead)
	var response = initializers.DB.Create(&task)

	if response.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": response.Error.Error(),
		})
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"data":    task,
		"message": "Task Created Successfully",
	})
}

func GetAllTasks(c *gin.Context) {
	// Get all the tasks
	var tasks []models.Task
	var response = initializers.DB.Find(&tasks)

	if response.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": response.Error.Error(),
		})
		return
	}

	// Return Tasks in response
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"data":    tasks,
		"message": "Fetched Successfully",
	})
}

func GetTaskById(c *gin.Context) {
	// Get id from URL param
	var taskId = c.Param("id")

	// Get the task model defined for DB
	var task models.Task
	var response = initializers.DB.First(&task, taskId)
	if response.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": response.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"data":    task,
		"message": "Fetched Successfuly",
	})

}

func UpdateTask(c *gin.Context) {
	// Get id from URL param
	var id = c.Param("id")

	// get the data of req body
	var body struct {
		Content string
		IsRead  bool
	}
	c.Bind(&body)

	// Get a single Task that we want to update
	// Check if the Task exists
	var task models.Task
	var findResponse = initializers.DB.First(&task, id)
	if findResponse.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Task not found",
		})
		return
	}

	// Update it
	var updateResponse = initializers.DB.Model(&task).Select("Content", "IsRead").Updates(models.Task{
		Content: body.Content,
		IsRead:  body.IsRead})

	if updateResponse.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": updateResponse.Error.Error(),
		})
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"data":    task,
		"message": "Updated Task Successfully",
	})
}

func DeleteTask(c *gin.Context) {
	// Get id from URL param
	var id = c.Param("id")

	// Check if the task exists
	var task models.Task
	var findResponse = initializers.DB.First(&task, id)
	if findResponse.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Task not found",
		})
		return
	}

	// Delete the Task
	var deleteResponse = initializers.DB.Delete(&task)
	if deleteResponse.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": deleteResponse.Error.Error(),
		})
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "Task removed Successfully",
	})
}
