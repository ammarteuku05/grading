package controller

import (
	"teacher-grading-api/internal/service"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/dto"
	"teacher-grading-api/shared/errors"

	"github.com/labstack/echo"
)

type (
	//UserController is
	UserController struct {
		services service.Holder
		deps     shared.Deps
	}
)

// NewUserController is
func NewUserController(services service.Holder, deps shared.Deps) (*UserController, error) {
	return &UserController{
		services: services,
		deps:     deps,
	}, nil
}

// Transfer is
func (ctrl *UserController) Login(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.LoginUser)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}
	if err := ctx.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	user, err := ctrl.services.UserService.LoginUser(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(user)
}

func (ctrl *UserController) Register(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.RegisterUser)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}
	if err := ctx.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	if err := request.Validate(); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	if err := request.Validate(); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	err := ctrl.services.UserService.RegisterUser(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}
