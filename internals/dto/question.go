package dto

import "time"

type QuestionCreateRequest struct {
	Text string `json:"text" binding:"required"`
}

type QuestionResponse struct {
	ID        int              `json:"id"`
	Text      string           `json:"text"`
	CreatedAt time.Time        `json:"created_at"`
	Answers   []AnswerResponse `json:"answers,omitempty"`
}
