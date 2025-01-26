package dto

import (
	"teacher-grading-api/internal/entity"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterUser struct {
	FullName string          `json:"full_name"`
	Password string          `json:"password"`
	Email    string          `json:"email"`
	Role     entity.TypeUser `json:"role"`
}

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v *RegisterUser) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.Email, validation.Required, is.Email),
		validation.Field(&v.FullName, validation.Required),
		validation.Field(&v.Password, validation.Required),
		validation.Field(&v.Role, validation.In(entity.StudentType, entity.TeacherType)),
	)
}
