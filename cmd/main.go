package main

import (
	"github.com/RoshanRajcmd/todo-app-backend/initializers"
	"github.com/RoshanRajcmd/todo-app-backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	var r = gin.Default()

	// Todo Routes
	routes.TodoRoutes(r)

	r.Run()
}
