package authmiddleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

var Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return &Claims{}
	},
	SigningKey: []byte("adfadsfasdfasdfasdaf"),
}
