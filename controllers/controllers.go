package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"quizmaker.com/quizmakerapi/initializers"
	"quizmaker.com/quizmakerapi/models"
)

func GetQuiz(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, questions)
}

func PostQuiz(c *gin.Context) {
	var newQuiz models.QuizJson

	if err := c.BindJSON(&newQuiz); err != nil {
		log.Fatal("Unable to parse the data.")
		return
	}
	quiz := models.Quiz{Name: newQuiz.Name}
	quizR := initializers.DB.Create(&quiz)

	if err := quizR.Error; err != nil {
		log.Fatal("Unable to create quiz table.")
		return
	}

	for _, question := range newQuiz.Questions {
		qns := models.Question{QuizID: quiz.ID, Question: question.Question}
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
