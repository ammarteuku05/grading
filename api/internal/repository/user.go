package repository

import (
	"context"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/shared"

	"gorm.io/gorm"
)

type (
	// UserRepository is
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) error
		FindByID(ctx context.Context, ID string) (*entity.User, bool, error)
		UpdateByID(ctx context.Context, ID string, dataUpdate map[string]interface{}) (*entity.User, error)
		FindByEmail(ctx context.Context, email string) (*entity.User, error)
	}

	userImpl struct {
		shared.Deps
	}
)

func NewUserRepository(deps shared.Deps) UserRepository {
	return &userImpl{Deps: deps}
}

func (r *userImpl) Create(ctx context.Context, user *entity.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userImpl) FindByID(ctx context.Context, id string) (*entity.User, bool, error) {
	var (
		err  error
		user *entity.User
	)

	if err = r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}

	return user, true, nil
}

func (r *userImpl) UpdateByID(ctx context.Context, id string, dataUpdate map[string]interface{}) (*entity.User, error) {

	var user *entity.User

	if err := r.DB.Model(&user).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User

	if err := r.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
