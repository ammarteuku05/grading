package controller

import (
	"teacher-grading-api/internal/service"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/dto"
	"teacher-grading-api/shared/errors"
	"teacher-grading-api/shared/pagination"

	"github.com/labstack/echo"
)

type (
	//AssignmentController is
	AssignmentController struct {
		services service.Holder
		deps     shared.Deps
	}
)

// NewAssignmentController is
func NewAssignmentController(services service.Holder, deps shared.Deps) (*AssignmentController, error) {
	return &AssignmentController{
		services: services,
		deps:     deps,
	}, nil
}

// CreateAssignment is
func (ctrl *AssignmentController) CreateAssignment(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = ctx.Request().Context()
		request = new(dto.Assignment)
		auth    = shared.GetLoggedInUser(ctx)
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

	request.StudentID = auth.ID
	err := ctrl.services.AssigmentService.CreateAssignment(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}

func (ctrl *AssignmentController) GetAllAssignment(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		page    = pagination.NewFromRequest(pctx.Request())
	)

	res, total, err := ctrl.services.AssigmentService.GetAllAssignment(context, ctx.QueryParam("subject"), page.Limit(), page.Offset())
	if err != nil {
		return pctx.Fail(err)
	}

	page.SetData(res, total)
	return pctx.Success(page)
}
