package shared

import (
	"github.com/labstack/echo"
)

type User struct {
	ID       string
	Fullname string
	Email    string
	Role     string
}

func GetLoggedInUser(c echo.Context) User {
	// Initialize an empty User
	var user User

	// Safely retrieve and cast values from context
	if role, ok := c.Get("role").(string); ok {
		user.Role = role
	}

	if id, ok := c.Get("currentUser").(string); ok {
		user.ID = id
	}

	if fullName, ok := c.Get("fullname").(string); ok {
		user.Fullname = fullName
	}

	if email, ok := c.Get("email").(string); ok {
		user.Email = email
	}

	return user

}
