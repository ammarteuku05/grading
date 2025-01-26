package repository

import (
	"context"
	"fmt"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/shared"

	"gorm.io/gorm"
)

type (
	// AssignmentRepository is
	AssignmentRepository interface {
		Create(ctx context.Context, assigment *entity.Assignment) error
		FindByID(ctx context.Context, ID string) (entity.Assignment, error)
		FindAll(ctx context.Context, subject string, limit, offset int) ([]entity.Assignment, int, error)
		Update(ctx context.Context, assignment *entity.Assignment) error
	}

	assignmentImpl struct {
		shared.Deps
	}
)

func NewAssignmentRepository(deps shared.Deps) AssignmentRepository {
	return &assignmentImpl{Deps: deps}
}

func (r *assignmentImpl) Create(ctx context.Context, assigment *entity.Assignment) error {
	if err := r.DB.Create(&assigment).Error; err != nil {
		return err
	}

	return nil
}

func (r *assignmentImpl) FindByID(ctx context.Context, ID string) (entity.Assignment, error) {
	var (
		err error
		res entity.Assignment
	)

	if err = r.DB.Where("id = ?", ID).Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, err
		}
		return res, err
	}

	return res, nil
}

func (r *assignmentImpl) FindAll(ctx context.Context, subject string, limit, offset int) ([]entity.Assignment, int, error) {
	var (
		assigment []entity.Assignment
		count     int64
	)

	stmt := r.DB.Model(&assigment)
	stmt = stmt.Debug()
	stmt = stmt.Where("status IS NULL")
	if subject != "" {
		stmt = stmt.Where("subject LIKE ?", fmt.Sprintf("%%%s%%", subject))
	}

	err := stmt.Limit(limit).Offset(offset).Order("created_at DESC").Find(&assigment).Error
	if err != nil {
		return nil, 0, err
	}

	err = stmt.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return assigment, int(count), nil
}

func (r *assignmentImpl) Update(ctx context.Context, assigment *entity.Assignment) error {
	if err := r.DB.UpdateColumns(&assigment).Error; err != nil {
		return err
	}

	return nil
}
