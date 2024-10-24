package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mlgnu/cocode/internal/auth"
	authrepo "github.com/mlgnu/cocode/internal/auth/repository"
	"github.com/mlgnu/cocode/internal/user"
	userrepo "github.com/mlgnu/cocode/internal/user/repository"
	authmiddleware "github.com/mlgnu/cocode/pkg"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	e := echo.New()
	ctx := context.Background()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))
	fmt.Println(connStr)
	conn, err := pgx.Connect(ctx, connStr)
	e.Validator = &CustomValidator{validator: validator.New()}

	if err != nil {
		log.Fatal("Faild to connect to database with error: ", err)
	}
	defer conn.Close(ctx)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("api")

	authRouter := api.Group("/auth")
	authRepo := authrepo.New(conn)
	authService := auth.NewService(authRepo)
	auth.RegisterRoutes(authRouter, authService)

	userRouter := api.Group("users")
	userRouter.Use(echojwt.WithConfig(authmiddleware.Config))
	userRepo := userrepo.New(conn)
	userService := user.NewService(userRepo)
	user.RegisterRoutes(userRouter, userService)

	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))
}
