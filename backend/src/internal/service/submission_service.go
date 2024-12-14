package service

import (
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/repository"
)

type SubmissionService struct {
	formRepo     *repository.FormRepository
	questionRepo *repository.QuestionRepository
}

func NewSubmissionService(formRepo *repository.FormRepository, questionRepo *repository.QuestionRepository) *SubmissionService {
	return &SubmissionService{
		formRepo:     formRepo,
		questionRepo: questionRepo,
	}
}

func (s *SubmissionService) SubmitForm(submission *models.Submission) error {
	// Implementation
	return nil
}
