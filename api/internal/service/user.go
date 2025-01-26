package service

import (
	"context"
	"errors"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/internal/repository"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UserService is
	UserService interface {
		RegisterUser(ctx context.Context, req *dto.RegisterUser) error
		LoginUser(ctx context.Context, req *dto.LoginUser) (string, error)
	}

	userImpl struct {
		repo repository.Holder
		deps shared.Deps
	}
)

// NewUserService is
func NewUserService(repo repository.Holder, deps shared.Deps) UserService {
	return &userImpl{
		repo: repo,
		deps: deps,
	}
}

func (s *userImpl) RegisterUser(ctx context.Context, req *dto.RegisterUser) error {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	newUser := entity.User{
		Id:        uuid.NewString(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  string(genPassword),
		Role:      entity.TypeUser(req.Role),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.repo.UserRepository.Create(ctx, &newUser)

	if err != nil {
		return err
	}

	return nil
}

func (s *userImpl) LoginUser(ctx context.Context, req *dto.LoginUser) (string, error) {
	user, err := s.repo.UserRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return "", err
	}

	if user.Id == "" {
		newError := "user is not found"
		return "", errors.New(newError)
	}

	// pengecekan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("password invalid")
	}

	auth, err := s.generateToken(user.Id, user.FullName, user.Email, string(user.Role))
	if err != nil {
		return "", err
	}

	return auth, nil
}

func (s *userImpl) generateToken(userID, fullName, email, role string) (string, error) {
	claim := jwt.MapClaims{
		"id":        userID,
		"full_name": fullName,
		"email":     email,
		"role":      role,
	}

	// generate token using HS256 with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(s.deps.Config.JwtSecret))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
