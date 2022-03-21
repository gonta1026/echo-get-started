package handler

import (
	"echo-get-started/app/application/usecase/user"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	usersUsecase user.UsersAccessor
}

func NewUser(users user.UsersAccessor) *User {
	return &User{
		usersUsecase: users,
	}
}

func (u User) Users(c echo.Context) error {
	res, _ := u.usersUsecase.Handle()
	return c.JSON(http.StatusOK, res)
}
