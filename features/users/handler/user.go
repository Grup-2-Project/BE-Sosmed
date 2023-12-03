package handler

import (
	"BE-Sosmed/features/users"
	"BE-Sosmed/helper/jwt"
	"BE-Sosmed/helper/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	s users.Service
}

func New(s users.Service) users.Handler {
	return &userHandler{
		s: s,
	}
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(users.User)
		inputProcess.FirstName = input.FirstName
		inputProcess.LastName = input.LastName
		inputProcess.Gender = input.Gender
		inputProcess.Hp = input.Hp
		inputProcess.Email = input.Email
		inputProcess.Password = input.Password


		result, err := uh.s.Register(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar pada sistem"
			}

			return responses.PrintResponse(c, statusCode, message, nil)
		}

		var response = new(RegisterResponse)
		response.ID = result.ID
		response.FirstName = result.FirstName
		response.LastName = result.LastName
		response.Gender = result.Gender
		response.Hp = result.Hp
		response.Email = result.Email


		return responses.PrintResponse(c, http.StatusCreated, "success create data", response)
		// return c.JSON(http.StatusCreated, map[string]any{
		// 	"message": "success create data",
		// 	"data":    response,
		// })
	}
}


func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		strToken, err := jwt.GenerateJWT(result.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.ID = result.ID
		response.FirstName = result.FirstName
		response.LastName = result.LastName
		response.Gender = result.Gender
		response.Hp = result.Hp
		response.Email = result.Email
		response.Image = result.Image
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]any{
			"message": "login success",
			"data":    response,
		})
	}
}
