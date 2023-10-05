package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"not null"`
	Email        string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`
}

type Quiz struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Questions []Question
}

type Question struct {
	gorm.Model
	QuizID   uint   `gorm:"index;not null"`
	Question string `gorm:"not null"`
	Opts     []Answer
}

type Answer struct {
	gorm.Model
	QuestionID uint   `gorm:"index;not null"`
	Value      string `gorm:"not null"`
	IsCorrect  bool   `gorm:"not null"`
}
