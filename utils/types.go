package utils

import "github.com/labstack/echo/v4"

type (
	Any             = interface{}
	AnyStruct       = map[string]Any
	MessageResponse struct {
		Message string `json:"message"`
	}
	TokenResponse struct {
		MessageResponse
		Token string `json:"token"`
	}
	Handler func(echo.Context) error
)
