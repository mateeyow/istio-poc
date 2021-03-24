package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	authAPI := os.Getenv("AUTH_URL")
	if authAPI == "" {
		authAPI = "http://auth:8887"
	}

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/admin-api", func(c echo.Context) error {
		_, err := http.Get(authAPI + "/auth-api")

		if err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusInternalServerError, "Error")
		}

		return c.String(http.StatusOK, "Got it")
	})

	e.Logger.Fatal(e.Start(":8888"))
}
