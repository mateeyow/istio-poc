package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/auth-api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Auth API response")
	})

	e.Logger.Fatal(e.Start(":8889"))
}
