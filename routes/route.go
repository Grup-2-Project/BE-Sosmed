package routes

import (
	"BE-Sosmed/features/users"
	"BE-Sosmed/features/postings"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uh users.Handler, ph postings.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uh)
	routePosting(e, ph)
}

func routeUser(e *echo.Echo, uh users.Handler) {
	e.POST("/users", uh.Register())
	e.POST("/login", uh.Login())
}

func routePosting(e *echo.Echo, ph postings.Handler) {
	e.POST("/posting", ph.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}