package service

import (
	"fmt"

	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormService struct {
	formRepo *repository.FormRepository
}

func NewFormService(formRepo *repository.FormRepository) *FormService {
	return &FormService{
		formRepo: formRepo,
	}
}

func (s *FormService) CreateForm(form *models.Form) error {
	if err := s.validateForm(form); err != nil {
		return fmt.Errorf("form validation failed: %w", err)
	}

	return s.formRepo.CreateForm(form)
}

func (s *FormService) GetForm(id string) (*models.Form, error) {
	form, err := s.formRepo.GetForm(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get form: %w", err)
	}
	return form, nil
}

func (s *FormService) ListForms(userID string) ([]models.Form, error) {
	forms, err := s.formRepo.ListForms(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list forms: %w", err)
	}
	return forms, nil
}

func (s *FormService) UpdateForm(form *models.Form) error {
	if err := s.validateForm(form); err != nil {
		return fmt.Errorf("form validation failed: %w", err)
	}

	return s.formRepo.UpdateForm(form)
}

func (s *FormService) DeleteForm(id string) error {
	return s.formRepo.DeleteForm(id)
}

func (s *FormService) validateForm(form *models.Form) error {
	if form.Name == "" {
		return fmt.Errorf("form name is required")
	}

	// if len(form.Questions) == 0 {
	// 	return fmt.Errorf("form must have at least one question")
	// }

	// Validate questions
	// for i, question := range form.Questions {
	// 	if err := s.validateQuestion(question, i); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (s *FormService) validateQuestion(question models.Question, index int) error {
	if question.Question == "" {
		return fmt.Errorf("question text is required for question %d", index+1)
	}

	if question.Type == "" {
		return fmt.Errorf("question type is required for question %d", index+1)
	}

	switch question.Type {
	case "single-choice", "multiple-choice":
		if len(question.Options) == 0 {
			return fmt.Errorf("options are required for question %d", index+1)
		}
	case "input":
		if question.InputType == "" {
			return fmt.Errorf("input type is required for input question %d", index+1)
		}
	default:
		return fmt.Errorf("invalid question type for question %d", index+1)
	}

	return nil
}

func (s *FormService) ToggleDraftStatus(formID string, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(formID)
	if err != nil {
		return fmt.Errorf("invalid form ID: %w", err)
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	return s.formRepo.ToggleDraftStatus(objectID, userObjectID)
}
