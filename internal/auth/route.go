package auth

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, service *Service) {
	handler := NewHandler(service)

	e.POST("/register", handler.Register)
	e.POST("/login", handler.LogIn)
	e.GET("/users", handler.GetUserByEmail)
	e.GET("/users/:id", handler.GetUserById)
}
