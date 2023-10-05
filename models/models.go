package models

type UserJson struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type QuizJson struct {
	ID        uint           `json:"quizId"`
	Name      string         `json:"name"`
	Questions []QuestionJson `json:"questions"`
}

type QuestionJson struct {
	ID       uint         `json:"questionId"`
	QuizID   uint         `json:"quizId"`
	Question string       `json:"question"`
	Opts     []AnswerJson `json:"opts"`
}

type AnswerJson struct {
	ID        uint   `json:"answerId"`
	QuestinID uint   `json:"questionId"`
	Value     string `json:"value"`
	IsCorrect bool   `json:"isCorrect"`
}
