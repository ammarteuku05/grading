package service

import "go.uber.org/dig"

type (
	Holder struct {
		dig.In
		UserService      UserService
		AssigmentService AssigmentService
		GradeService     GradeService
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserService); err != nil {
		return err
	}
	if err := container.Provide(NewAssignmentService); err != nil {
		return err
	}
	if err := container.Provide(NewGradeService); err != nil {
		return err
	}

	return nil
}
