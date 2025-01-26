package entity

import "time"

type TypeUser string

const (
	TeacherType TypeUser = "teacher"
	StudentType TypeUser = "student"
)

type User struct {
	Id        string   `json:"id"`
	FullName  string   `json:"full_name"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	Role      TypeUser `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
