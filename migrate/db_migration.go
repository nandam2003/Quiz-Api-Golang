package main

import (
	"quizmaker.com/quizmakerapi/initializers"
	"quizmaker.com/quizmakerapi/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Create tables based on the defined structs and auto-migrate the schema.
	initializers.DB.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Question{}, &models.Answer{})
}
