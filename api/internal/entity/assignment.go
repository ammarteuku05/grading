package entity

import (
	"time"
)

type Assignment struct {
	ID        string  `json:"id"`
	Subject   string  `json:"subject"`
	Tittle    string  `json:"tittle"`
	StudentID string  `json:"student_id"`
	Content   string  `json:"content"`
	Status    *string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
