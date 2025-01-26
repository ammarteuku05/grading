package shared

import (
	"database/sql"
	"teacher-grading-api/logger"
	"teacher-grading-api/shared/config"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

type (
	Deps struct {
		dig.In

		DB     *gorm.DB
		Config *config.Configuration
		Logger logger.Logger

		CustomValidator *CustomValidator
	}
)

func Register(container *dig.Container) error {

	if err := container.Provide(NewCustomValidator); err != nil {
		return err
	}

	return nil
}

func (deps Deps) DBTransaction(fc func(db gorm.DB) error, opts ...*sql.TxOptions) error {
	panicked := true
	tx := deps.DB.Begin()

	var (
		err error
	)
	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err == nil {
		err = fc(*tx)
	}

	if err == nil {
		err = tx.Commit().Error
	}
	panicked = false
	return err
}
