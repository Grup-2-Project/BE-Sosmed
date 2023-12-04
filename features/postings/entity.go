package postings

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type Posting struct {
	ID          uint
	Artikel string
	Gambar      string
	UserID      uint
}

type Handler interface {
	Add() echo.HandlerFunc
}

type Service interface {
	TambahPosting(token *jwt.Token, newPosting Posting) (Posting, error)
}

type Repository interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
}
