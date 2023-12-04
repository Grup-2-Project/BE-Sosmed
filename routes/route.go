package routes

import (
	"BE-Sosmed/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uh users.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uh)
}

func routeUser(e *echo.Echo, uh users.Handler) {
	e.POST("/users", uh.Register())
	e.POST("/login", uh.Login())
	e.GET("/users/:id", uh.ReadById(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/users", uh.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/users", uh.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
