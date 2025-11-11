package postgres

import (
	"answerService/internals/models"
	"context"
)

func (st *PostgresStorage) Answer(ctx context.Context, id int) (*models.Answer, error) {
	var answer models.Answer
	if err := st.db.WithContext(ctx).First(&answer, id).Error; err != nil {
		return nil, err
	}
	return &answer, nil
}

func (st *PostgresStorage) Answers(ctx context.Context, questionID int) ([]models.Answer, error) {
	var answers []models.Answer
	if err := st.db.WithContext(ctx).
		Where("question_id = ?", questionID).
		Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (st *PostgresStorage) CreateAnswer(ctx context.Context, answer *models.Answer) (*models.Answer, error) {
	if err := st.db.WithContext(ctx).Create(answer).Error; err != nil {
		return nil, err
	}

	return answer, nil
}

func (st *PostgresStorage) DeleteAnswer(ctx context.Context, id int) error {
	return st.db.WithContext(ctx).Delete(&models.Answer{}, id).Error
}
