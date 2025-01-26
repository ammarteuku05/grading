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
	//GradeController is
	GradeController struct {
		services service.Holder
		deps     shared.Deps
	}
)

// NewGradeController is
func NewGradeController(services service.Holder, deps shared.Deps) (*GradeController, error) {
	return &GradeController{
		services: services,
		deps:     deps,
	}, nil
}

// CreateAssignment is
func (ctrl *GradeController) CreateGrade(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = ctx.Request().Context()
		request = new(dto.Grade)
		auth    = shared.GetLoggedInUser(ctx)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}
	if err := ctx.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	request.TeacherID = auth.ID
	err := ctrl.services.GradeService.CreateGrade(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}

func (ctrl *GradeController) GetAllGrade(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		page    = pagination.NewFromRequest(pctx.Request())
		auth    = shared.GetLoggedInUser(ctx)
	)

	res, total, err := ctrl.services.GradeService.GetGradeStudentId(context, auth.ID, page.Limit(), page.Offset())
	if err != nil {
		return pctx.Fail(err)
	}

	page.SetData(res, total)
	return pctx.Success(page)
}
