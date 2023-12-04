package users

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	FirstName     string
	LastName string
	Gender string
	Hp string
	Email string
	Password string
	Image string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Service interface {
	Register(newUser User) (User, error)
	Login(email string, password string) (User, error)
}

type Repository interface {
	InsertUser(newUser User) (User, error)
	Login(email string) (User, error)
}
