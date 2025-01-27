package entity

import (
	"time"
)

type Grade struct {
	ID           string     `json:"id"`
	AssignmentID string     `json:"assignment_id"`
	Score        float64    `json:"score"`
	Feedback     string     `json:"feedback"`
	TeacherID    string     `json:"teacher_id"`
	CreatedAt    time.Time  `json:"create_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
