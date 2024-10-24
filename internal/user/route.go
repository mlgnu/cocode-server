package user

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, service *Service) {
	handler := NewHandler(service)

	e.GET(":id", handler.GetUser)
}
