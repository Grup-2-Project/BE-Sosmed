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
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	TambahPosting(token *jwt.Token, newPosting Posting) (Posting, error)
	SemuaPosting() ([]Posting, error)
	UpdatePosting(token *jwt.Token, updatePosting Posting) (Posting, error)
	DeletePosting(token *jwt.Token, postID uint) error
}

type Repository interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
	GetAllPost() ([]Posting, error)
	UpdatePost(userID uint, updatePosting Posting) (Posting, error)
	DeletePost(userID uint, postID uint) error
}
