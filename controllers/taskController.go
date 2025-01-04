package controllers

import (
	"fmt"

	"github.com/RoshanRajcmd/todo-app-backend/initializers"
	"github.com/RoshanRajcmd/todo-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateTask(c *gin.Context) {
	// Get data from req body
	var body struct {
		Content string
		IsRead  bool
	}
	c.Bind(&body)

	// Create a todo
	//var task = models.Task{Content: body.Content, IsRead: body.IsRead}
	var task = models.NewTask(body.Content, body.IsRead)
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
	var response = initializers.DB.First(&task, taskId)
	// fmt.Print(response)
	// fmt.Println(response.Error)
	if response.Error == nil {
		// Return todo in response
		c.JSON(200, gin.H{
			"data":    task,
			"message": "Fetched Successfully",
		})
	} else {
		pgError := response.Error.(*pgconn.PgError)
		fmt.Print("Error Code: ", pgError.Code)

		c.JSON(500, gin.H{
			"code":    pgError.Code,
			"message": pgError.Message,
		})
	}
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
		if pgError, ok := response.Error.(*pgconn.PgError); ok {
			fmt.Print(pgError.Code)
			// if pgError.Code == UniqueViolation && pgError.ConstraintName == "users_email_key" {
			// 	// Handle specifically the email constraint broken
			// }
			// // Handle unknown error
		} else {
			// Handle non-MySQL errors or unknown errors
		}
	} else {
		// Return response
		c.JSON(200, gin.H{
			"message": "Task removed Successfully",
		})
	}
}
