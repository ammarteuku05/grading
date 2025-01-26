package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Assignment struct {
	Subject   string `json:"subject"`
	Tittle    string `json:"tittle"`
	StudentID string `json:"student_id"`
	Content   string `json:"content"`
}

func (v *Assignment) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.Subject, validation.Required),
		validation.Field(&v.Tittle, validation.Required),
		validation.Field(&v.Content, validation.Required),
	)
}
