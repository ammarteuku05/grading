package service

import (
	"context"
	"errors"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/internal/repository"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/dto"
	"teacher-grading-api/shared/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	// GradeService Service is
	GradeService interface {
		CreateGrade(ctx context.Context, req *dto.Grade) error
		GetGradeStudentId(ctx context.Context, id string, limit, offset int) ([]dto.ResponseGrade, int, error)
	}

	gradeImpl struct {
		repo repository.Holder
		deps shared.Deps
	}
)

// NewGradeService is
func NewGradeService(repo repository.Holder, deps shared.Deps) GradeService {
	return &gradeImpl{
		repo: repo,
		deps: deps,
	}
}

func (s *gradeImpl) CreateGrade(ctx context.Context, req *dto.Grade) error {
	assigment, err := s.repo.AssignmentRepository.FindByID(ctx, req.AssignmentID)
	if err != nil {
		return err
	}

	if assigment.ID == "" {
		return errors.New("assignment not found")
	}

	return s.deps.DBTransaction(func(db gorm.DB) error {
		err := s.repo.GradeRepository.Create(ctx, &entity.Grade{
			ID:           uuid.New().String(),
			TeacherID:    req.TeacherID,
			Feedback:     req.Feedback,
			AssignmentID: req.AssignmentID,
			Score:        req.Score,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		if err != nil {
			return err
		}

		err = s.repo.AssignmentRepository.Update(ctx, &entity.Assignment{
			ID:     assigment.ID,
			Status: utils.StrPtr("done"),
		})

		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (s *gradeImpl) GetGradeStudentId(ctx context.Context, id string, limit, offset int) ([]dto.ResponseGrade, int, error) {
	return s.repo.GradeRepository.FindByStudentID(ctx, id, limit, offset)
}
