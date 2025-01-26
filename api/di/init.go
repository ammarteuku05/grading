package di

import (
	"teacher-grading-api/controller"
	"teacher-grading-api/internal/repository"
	"teacher-grading-api/internal/service"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/config"

	"go.uber.org/dig"
)

var (
	Container *dig.Container = dig.New()
)

func init() {
	if err := Container.Provide(config.New); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewOrm); err != nil {
		panic(err)
	}

	if err := shared.Register(Container); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewLogger); err != nil {
		panic(err)
	}

	if err := service.Register(Container); err != nil {
		panic(err)
	}

	if err := repository.Register(Container); err != nil {
		panic(err)
	}

	if err := controller.Register(Container); err != nil {
		panic(err)
	}
}
