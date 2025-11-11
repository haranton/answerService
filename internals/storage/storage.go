package storage

import (
	"answerService/internals/models"
	"context"
)

type Storage interface {
	QuestionStorage
	AnswerStorage
}

type QuestionStorage interface {
	Question(ctx context.Context, id int) (*models.Question, error)
	Questions(ctx context.Context) ([]models.Question, error)
	CreateQuestion(ctx context.Context, question *models.Question) (*models.Question, error)
	DeleteQuestion(ctx context.Context, id int) error
}

type AnswerStorage interface {
	Answer(ctx context.Context, id int) (*models.Answer, error)
	Answers(ctx context.Context, questionID int) ([]models.Answer, error)
	CreateAnswer(ctx context.Context, answer *models.Answer) (*models.Answer, error)
	DeleteAnswer(ctx context.Context, id int) error
}
