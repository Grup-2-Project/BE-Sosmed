package handler

import (
	"BE-Sosmed/features/postings"
	"context"
	"log"
	"net/http"
	"strings"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type PostingHandler struct {
	s postings.Service
}

func New(s postings.Service) postings.Handler {
	return &PostingHandler{
		s: s,
	}
}

func (pc *PostingHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(AddRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		fileHeader, err := c.FormFile("gambar")
		if input.Artikel == "" && err == nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Anda harus memasukkan artikel",
			})
		}

		var urlCloudinary = "cloudinary://533421842888945:Oish5XyXkCiiV6oTW2sEo0lEkGg@dlxvvuhph"

		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(input); err != nil {
			c.Echo().Logger.Error("Input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
		}

		var inputProcess = new(postings.Posting)
		if err != nil {
			inputProcess.Artikel = input.Artikel
		} else {
			log.Println(fileHeader.Filename)

			file, _ := fileHeader.Open()

			var ctx = context.Background()

			cldService, _ := cloudinary.NewFromURL(urlCloudinary)
			resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
			log.Println(resp.SecureURL)

			inputProcess.Artikel = input.Artikel
			inputProcess.Gambar = resp.SecureURL
		}

		result, err := pc.s.TambahPosting(c.Get("user").(*gojwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan sudah terdaftar ada sistem",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(AddResponse)
		response.ID = result.ID
		response.Artikel = result.Artikel
		response.Gambar = result.Gambar

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}