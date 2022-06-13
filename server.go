package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FRONTEND_LINK")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/api/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "API!")
	})

	e.GET("/api/v1/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello from backend :) It is nice to be here!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
