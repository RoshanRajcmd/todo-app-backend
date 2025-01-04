package main

import (
	"github.com/RoshanRajcmd/todo-app-backend/initializers"
	"github.com/RoshanRajcmd/todo-app-backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDB()
}

func main() {

	var app = gin.Default()

	// Todo Routes
	routes.TodoRoutes(app)

	app.Run()
}
