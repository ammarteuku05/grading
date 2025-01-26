package controller

import (
	"strings"
	"teacher-grading-api/internal/entity"
	"teacher-grading-api/shared"
	"teacher-grading-api/shared/errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

func MustLoggedIn(encryptionKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			token, err := validateToken(authHeader, encryptionKey)
			pctx := shared.NewEmptyContext(c)
			if err != nil {
				return pctx.Fail(errors.ErrSession)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return pctx.Fail(errors.ErrSession)
			}
			userID := string(claims["id"].(string))
			role := string(claims["role"].(string))

			c.Set("currentUser", userID)
			c.Set("role", role)
			return next(c)
		}
	}
}

func TeacherMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pctx := shared.NewEmptyContext(c)
			role := c.Get("role").(string)
			if role != string(entity.TeacherType) {
				return pctx.Fail(errors.ErrSession)
			}

			return next(c)
		}
	}
}

func validateToken(encodedToken, key string) (*jwt.Token, error) {
	encodedToken = strings.TrimPrefix(encodedToken, "Bearer ")

	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.ErrSessionHeader
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
