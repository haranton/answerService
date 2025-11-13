package postgres

import (
	"answerService/internals/models"
	"context"
)

func (st *PostgresStorage) Question(ctx context.Context, id int) (*models.Question, error) {
	var q models.Question
	if err := st.db.WithContext(ctx).First(&q, id).Error; err != nil {
		return nil, err
	}

	return &q, nil
}

func (st *PostgresStorage) QuestionWithAnswers(ctx context.Context, id int) (*models.Question, error) {
	var q models.Question

	if err := st.db.WithContext(ctx).
		Preload("Answers").
		First(&q, id).Error; err != nil {
		return nil, err
	}

	return &q, nil
}

func (st *PostgresStorage) Questions(ctx context.Context) ([]models.Question, error) {
	var questions []models.Question
	if err := st.db.WithContext(ctx).
		Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (st *PostgresStorage) CreateQuestion(ctx context.Context, question *models.Question) (*models.Question, error) {
	if err := st.db.WithContext(ctx).Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (st *PostgresStorage) DeleteQuestion(ctx context.Context, id int) error {
	return st.db.WithContext(ctx).Delete(&models.Question{}, id).Error
}
