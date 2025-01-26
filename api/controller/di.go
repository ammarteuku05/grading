package controller

import (
	"net/http"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/errors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		Deps shared.Deps

		UserController       *UserController
		AssignmentController *AssignmentController
		GradeController      *GradeController
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserController); err != nil {
		return err
	}
	if err := container.Provide(NewAssignmentController); err != nil {
		return err
	}
	if err := container.Provide(NewGradeController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {

	app.Validator = h.Deps.CustomValidator
	app.HTTPErrorHandler = h.ErrorHandler

	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	v1 := app.Group("/v1")
	userRoute := v1.Group("/user")
	userRoute.POST("/register", h.UserController.Register)
	userRoute.POST("/login", h.UserController.Login)

	assignmentRoute := v1.Group("/assignment")
	assignmentRoute.Use(MustLoggedIn(h.Deps.Config.JwtSecret))
	assignmentRoute.POST("/create-assignment", h.AssignmentController.CreateAssignment)
	assignmentRoute.GET("/get-assignment", h.AssignmentController.GetAllAssignment, TeacherMiddleware())

	gradeRoute := v1.Group("/grade")
	gradeRoute.Use(MustLoggedIn(h.Deps.Config.JwtSecret))
	gradeRoute.POST("/create-grade", h.GradeController.CreateGrade, TeacherMiddleware())
	gradeRoute.GET("/get-grade", h.GradeController.GetAllGrade)
}

func (h *Holder) ErrorHandler(err error, ctx echo.Context) {
	var (
		sctx, ok = ctx.(*shared.Context)
	)

	if !ok {
		sctx = shared.NewEmptyContext(ctx)
	}

	e, ok := err.(*echo.HTTPError)
	if ok {
		msg, ok := e.Message.(string)
		if !ok {
			msg = err.Error()
		}
		err = errors.ErrBase.New(msg).WithProperty(errors.ErrCodeProperty, e.Code).WithProperty(errors.ErrHttpCodeProperty, e.Code)
	}

	h.Deps.Logger.Errorf(
		"path=%s error=%s",
		sctx.Path(),
		err,
	)

	_ = sctx.Fail(err)
}
