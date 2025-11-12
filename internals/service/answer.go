package service

import (
	"answerService/internals/models"
	"answerService/internals/storage"
	"context"
	"errors"
)

var (
	ErrQuestionNotFound = errors.New("question not found")
	ErrAnswerNotFound   = errors.New("answer not found")
)

type AnswerService struct {
	storage storage.Storage
}

func NewAnswerService(storage storage.Storage) *AnswerService {
	return &AnswerService{storage: storage}
}

func (a *AnswerService) Answer(ctx context.Context, id int) (*models.Answer, error) {

	answer, err := a.storage.Answer(ctx, id)
	if err != nil {
		return nil, err
	}
	if answer == nil {
		return nil, ErrAnswerNotFound
	}
	return answer, nil
}

func (a *AnswerService) CreateAnswer(ctx context.Context, answer *models.Answer) (*models.Answer, error) {
	question, err := a.storage.Question(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, ErrQuestionNotFound
	}
	return a.storage.CreateAnswer(ctx, answer)
}

func (a *AnswerService) DeleteAnswer(ctx context.Context, id int) error {
	question, err := a.storage.Answer(ctx, id)
	if err != nil {
		return err
	}
	if question == nil {
		return ErrAnswerNotFound
	}
	return a.storage.DeleteAnswer(ctx, id)
}
