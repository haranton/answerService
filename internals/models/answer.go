package models

import "time"

type Answer struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	QuestionID int       `json:"question_id" gorm:"not null;index"`
	UserID     string    `json:"user_id" gorm:"type:varchar(100);not null"`
	Text       string    `json:"text" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`

	Question Question `json:"-" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}
