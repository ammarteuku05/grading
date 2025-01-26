package service

import (
	"context"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/internal/repository"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/dto"
	"time"

	"github.com/google/uuid"
)

type (
	// Assigment Service is
	AssigmentService interface {
		CreateAssignment(ctx context.Context, req *dto.Assignment) error
		GetAllAssignment(ctx context.Context, subject string, limit, offset int) ([]entity.Assignment, int, error)
		GetAssignmentById(ctx context.Context, id string) (*entity.Assignment, error)
	}

	assigmentImpl struct {
		repo repository.Holder
		deps shared.Deps
	}
)

// NewAssignmentService is
func NewAssignmentService(repo repository.Holder, deps shared.Deps) AssigmentService {
	return &assigmentImpl{
		repo: repo,
		deps: deps,
	}
}

func (s *assigmentImpl) CreateAssignment(ctx context.Context, req *dto.Assignment) error {
	err := s.repo.AssignmentRepository.Create(ctx, &entity.Assignment{
		ID:        uuid.New().String(),
		Subject:   req.Subject,
		StudentID: req.StudentID,
		Tittle:    req.Tittle,
		Content:   req.Content,
		Status:    nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *assigmentImpl) GetAllAssignment(ctx context.Context, subject string, limit, offset int) ([]entity.Assignment, int, error) {
	return s.repo.AssignmentRepository.FindAll(ctx, subject, limit, offset)
}

func (s *assigmentImpl) GetAssignmentById(ctx context.Context, id string) (*entity.Assignment, error) {
	res, err := s.repo.AssignmentRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
