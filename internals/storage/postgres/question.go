package postgres

import (
	"answerService/internals/models"
	"context"
)

// type QuestionStorage interface {
// 	Question(ctx context.Context, id int)
// 	Questions(ctx context.Context, offset int, limit int)
// 	CreateQuestions(ctx context.Context, question *models.Question)
// 	DeleteQuestion(ctx context.Context, id int)
// }

func (st *PostgresStorage) Question(ctx context.Context, id int) (*models.Question, error) {
	var q models.Question
	if err := st.db.WithContext(ctx).First(&q, id).Error; err != nil {
		return nil, err
	}

	return &q, nil
}

func (st *PostgresStorage) Questions(ctx context.Context, offset int, limit int) ([]models.Question, error) {
	var questions []models.Question
	if err := st.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (st *PostgresStorage) CreateQuestions(ctx context.Context, question *models.Question) (*models.Question, error) {
	if err := st.db.WithContext(ctx).Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (st *PostgresStorage) DeleteQuestion(ctx context.Context, id int) error {
	return st.db.WithContext(ctx).Delete(&models.Question{}, id).Error
}
