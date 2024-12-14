// internal/repository/question_repository.go
package repository

import (
	"database/sql"

	"github.com/maxzhirnov/formease/internal/models"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) GetQuestions(formID int) ([]models.Question, error) {
	// Implementation similar to the questions part in FormRepository
	return nil, nil
}
