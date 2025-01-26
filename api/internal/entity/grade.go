package entity

import (
	"time"
)

type Grade struct {
	ID           string  `json:"id"`
	AssignmentID string  `json:"assignment_id"`
	Score        float64 `json:"score"`
	Feedback     string  `json:"feedback"`
	TeacherID    string  `json:"teacher_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
