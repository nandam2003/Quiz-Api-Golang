package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"quizmaker.com/quizmakerapi/controllers"
	"quizmaker.com/quizmakerapi/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173"}         // Specify the allowed origin(s)
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}  // Specify the allowed HTTP methods
	config.AllowHeaders = []string{"Content-Type", "Authorization"} // Specify the allowed headers

	// Use the CORS middleware with the configured options
	r.Use(cors.New(config))

	r.GET("/", controllers.GetQuiz)
	r.POST("/", controllers.PostQuiz)
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignUp)

	r.Run()
}
