package dto

import "time"

type AnswerCreateRequest struct {
	QuestionID int    `json:"question_id" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
}

type AnswerResponse struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	UserID     string    `json:"user_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
