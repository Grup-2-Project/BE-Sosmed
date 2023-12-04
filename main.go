package main

import (
	"BE-Sosmed/config"
	ph "BE-Sosmed/features/postings/handler"
	pr "BE-Sosmed/features/postings/repository"
	ps "BE-Sosmed/features/postings/service"
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

	db.AutoMigrate(&ur.UserModel{}, &pr.PostingModel{})

	userRepo := ur.New(db)
	hash := enkrip.New()
	userService := us.New(userRepo, hash)
	userHandler := uh.New(userService)

	postingRepo := pr.New(db)
	postingService := ps.New(postingRepo)
	postingHandler := ph.New(postingService)

	routes.InitRoute(e, userHandler, postingHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
