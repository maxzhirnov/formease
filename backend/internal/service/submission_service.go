package service

import (
	"fmt"

	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/repository"
)

type SubmissionService struct {
	subRepo *repository.SubmissionRepository
}

func NewSubmissionService(subRepo *repository.SubmissionRepository) *SubmissionService {
	return &SubmissionService{
		subRepo: subRepo,
	}
}

func (s *SubmissionService) CreateSubmission(sub *models.Submission) error {
	if err := s.validateSubmission(sub); err != nil {
		return fmt.Errorf("submission validation failed: %w", err)
	}

	return s.subRepo.CreateSubmission(sub)
}

func (s *SubmissionService) validateSubmission(sub *models.Submission) error {
	//TODO: Add validation
	return nil
}
