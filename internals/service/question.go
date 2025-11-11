package service

import (
	"answerService/internals/models"
	"answerService/internals/storage"
	"context"
)

type QuestionService struct {
	storage storage.Storage
}

func NewQuestionService(storage storage.Storage) *QuestionService {
	return &QuestionService{storage: storage}
}

func (q *QuestionService) Question(ctx context.Context, id int) (*models.Question, error) {
	return q.storage.Question(ctx, id)
}

func (q *QuestionService) CreateQuestion(ctx context.Context, question *models.Question) (*models.Question, error) {
	return q.storage.CreateQuestion(ctx, question)
}

func (q *QuestionService) DeleteQuestion(ctx context.Context, id int) error {
	return q.storage.DeleteQuestion(ctx, id)
}
