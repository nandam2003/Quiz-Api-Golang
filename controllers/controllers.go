package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"quizmaker.com/quizmakerapi/initializers"
	"quizmaker.com/quizmakerapi/models"
)

func Login(c *gin.Context) {
	db := initializers.DB
	var userJSON models.UserJson
	if err := c.BindJSON(&userJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	var user models.User
	if err := db.Where("email = ?", userJSON.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User not found."})
			log.Println("User not found")
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			log.Fatal(err)
			return
		}
	}
	if hashErr := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userJSON.Password)); hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Password"})
		log.Fatal(hashErr)
		return
	}
	c.JSON(http.StatusOK, &user)
}

func SignUp(c *gin.Context) {
	var newUser models.UserJson
	if c.BindJSON(&newUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to create User."})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to Hash password."})
		return
	}
	user := models.User{Username: newUser.Username, Email: newUser.Email, PasswordHash: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to Create user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User create succesfully"})

}

func GetQuiz(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, questions)
}

func PostQuiz(c *gin.Context) {
	var newQuiz models.Quiz

	if err := c.BindJSON(&newQuiz); err != nil {
		log.Fatal("Unable to parse the data.")
		return
	}
	quiz := initializers.DB.Create(&newQuiz)

	if err := quiz.Error; err != nil {
		log.Fatal("Unable to create quiz table.")
		return
	}

	for _, question := range newQuiz.Questions {
		qns := models.Question{QuizID: newQuiz.ID, Question: question.Question}
		qnsR := initializers.DB.Create(&qns)

		if err := qnsR.Error; err != nil {
			log.Fatal("Unable to create question table.")
			return
		}
		for _, option := range question.Opts {
			opt := models.Answer{QuestionID: qns.ID, Value: option.Value, IsCorrect: option.IsCorrect}
			optR := initializers.DB.Create(&opt)

			if err := optR.Error; err != nil {
				log.Fatal("Unable to create option table.")
				return
			}
		}

	}
}
