package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var chat = make(chan string)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if _, err := h.Service.GetUserByEmail(c.Request().Context(), req.Email); err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already exists"})
	}

	err := h.Service.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to register a user: %s", err)})
	}

	c.NoContent(http.StatusCreated)
	return nil
}

func (h *Handler) LogIn(e echo.Context) error {
	var req LoginRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	token, err := h.Service.Login(e.Request().Context(), req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to login: %s", err)})
	}

	e.JSON(http.StatusOK, map[string]string{"token": token})
	return nil
}

func (h *Handler) GetUserByEmail(e echo.Context) error {
	var req GetUserByEmailRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.Service.GetUserByEmail(e.Request().Context(), req.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to get user by email: %s", err)})
	}

	e.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) GetUserById(e echo.Context) error {
	var req GetUserByIdRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.Service.GetUserById(e.Request().Context(), req.Id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to get user by id: %s", err)})
	}

	e.JSON(http.StatusOK, user)
	return nil
}
