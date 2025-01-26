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
	var grades []entity.Grade
	err := r.DB.Table("grades").
		Select("id, assignment_id, teacher_id, score, feedback").
		Where("assignment_id IN (SELECT id FROM assignments WHERE student_id = ?) AND deleted_at IS NULL", studentId).
		Limit(limit).Offset(offset).
		Order("grades.created_at DESC").
		Scan(&grades).Error
	if err != nil {
		return res, 0, err
	}

	// Retrieve the corresponding assignments for these grades
	var assignments []entity.Assignment
	err = r.DB.Table("assignments").
		Select("id, subject, tittle, student_id, content").
		Where("student_id = ?", studentId).
		Scan(&assignments).Error
	if err != nil {
		return res, 0, err
	}

	// Create a map of assignment IDs to assignments
	assignmentMap := make(map[string]entity.Assignment)
	for _, assignment := range assignments {
		assignmentMap[assignment.ID] = assignment
	}

	// Combine grades with their corresponding assignments
	for _, grade := range grades {
		assignment, exists := assignmentMap[grade.AssignmentID]
		if exists {
			res = append(res, dto.ResponseGrade{
				ID:           grade.ID,
				AssignmentID: grade.AssignmentID,
				TeacherID:    grade.TeacherID,
				Score:        grade.Score,
				Feedback:     grade.Feedback,
				Assignment: struct {
					ID        string "json:\"id\""
					Subject   string "json:\"subject\""
					Tittle    string "json:\"tittle\""
					StudentID string "json:\"student_id\""
					Content   string "json:\"content\""
				}{
					ID:        assignment.ID,
					Subject:   assignment.Subject,
					Tittle:    assignment.Tittle,
					StudentID: assignment.StudentID,
					Content:   assignment.Content,
				},
			})
		}
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
