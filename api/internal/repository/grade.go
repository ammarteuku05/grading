package repository

import (
	"context"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/dto"
)

type (
	// GradeRepository is
	GradeRepository interface {
		Create(ctx context.Context, grade *entity.Grade) error
		FindByStudentID(ctx context.Context, studentId string, limit, offset int) ([]dto.ResponseGrade, int, error)
	}

	gradeImpl struct {
		shared.Deps
	}
)

func NewGradeRepository(deps shared.Deps) GradeRepository {
	return &gradeImpl{Deps: deps}
}

func (r *gradeImpl) Create(ctx context.Context, grade *entity.Grade) error {
	if err := r.DB.Create(&grade).Error; err != nil {
		return err
	}

	return nil
}

func (r *gradeImpl) FindByStudentID(ctx context.Context, studentId string, limit, offset int) ([]dto.ResponseGrade, int, error) {
	var (
		res   = make([]dto.ResponseGrade, 0)
		count int64
	)

	// Retrieve grades first
	err := r.DB.Table("grades").
		Select("grades.id AS id, assignment_id, teacher_id, score, feedback, assignments.subject AS assignment_subject, assignments.tittle AS assignment_tittle, assignments.content AS assignment_content, assignments.student_id AS assignment_student_id, assignments.status AS assignment_status").
		Joins("INNER JOIN assignments ON grades.assignment_id = assignments.id").
		Where("assignments.student_id = ? AND grades.deleted_at IS NULL", studentId).
		Limit(limit).Offset(offset).
		Order("grades.created_at DESC").
		Scan(&res).Error
	if err != nil {
		return res, 0, err
	}

	// Count the total records matching the filter
	err = r.DB.Table("grades").
		Joins("INNER JOIN assignments ON grades.assignment_id = assignments.id").
		Where("assignments.student_id = ? AND grades.deleted_at IS NULL", studentId).
		Count(&count).Error
	if err != nil {
		return res, 0, err
	}

	return res, int(count), nil
}
