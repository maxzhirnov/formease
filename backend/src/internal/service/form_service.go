package service

import (
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/repository"
)

type FormService struct {
	formRepo     *repository.FormRepository
	questionRepo *repository.QuestionRepository
}

func NewFormService(formRepo *repository.FormRepository, questionRepo *repository.QuestionRepository) *FormService {
	return &FormService{
		formRepo:     formRepo,
		questionRepo: questionRepo,
	}
}

func (s *FormService) GetForm(id int) (*models.Form, error) {
	return s.formRepo.GetForm(id)
}
