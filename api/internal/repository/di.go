package repository

import "go.uber.org/dig"

type (
	//Holder is
	Holder struct {
		dig.In

		UserRepository
		AssignmentRepository
		GradeRepository
	}
)

// Register is
func Register(container *dig.Container) error {
	if err := container.Provide(NewUserRepository); err != nil {
		return err
	}
	if err := container.Provide(NewAssignmentRepository); err != nil {
		return err
	}
	if err := container.Provide(NewGradeRepository); err != nil {
		return err
	}

	return nil
}
