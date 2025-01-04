package initializers

import (
	"log"
	"os"

	"github.com/RoshanRajcmd/todo-app-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global DB connection object
var DB *gorm.DB

func ConnectDB() {
	var err error

	//Load Environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}

	//Gets the enironment vaiable named DB_URL to open db connection
	var dsn = os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.Task{})
}
