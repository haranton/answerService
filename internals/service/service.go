package service

import (
	"answerService/internals/storage"
	"log/slog"
)

type Service struct {
	SrvAnswer   *AnswerService
	SrvQuestion *QuestionService
	storage     storage.Storage
}

func NewService(st storage.Storage, logger *slog.Logger) *Service {
	srvAnswer := NewAnswerService(st)
	srvQuestion := NewQuestionService(st)

	return &Service{
		SrvAnswer:   srvAnswer,
		SrvQuestion: srvQuestion,
		storage:     st,
	}
}
