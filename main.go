package main

import (
	"BE-Sosmed/config"
	uh "BE-Sosmed/features/users/handler"
	ur "BE-Sosmed/features/users/repository"
	us "BE-Sosmed/features/users/services"
	"BE-Sosmed/helper/enkrip"
	"BE-Sosmed/routes"
	"BE-Sosmed/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
e := echo.New()

	cfg := config.InitConfig()

	if cfg == nil {
		e.Logger.Fatal("tidak bisa start karena ENV error")
		return
	}

	db, err := database.InitMySQL(*cfg)

	if err != nil {
		e.Logger.Fatal("tidak bisa start karena DB error:", err.Error())
		return
	}

	db.AutoMigrate(&ur.UserModel{})

	userRepo := ur.New(db)
	hash := enkrip.New()
	userService := us.New(userRepo, hash)
	userHandler := uh.New(userService)

	routes.InitRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
