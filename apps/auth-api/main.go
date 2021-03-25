package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	rand.Seed(time.Now().UnixNano())

	e.GET("/auth-api", func(c echo.Context) error {
		if !isOK() {
			return c.String(http.StatusInternalServerError, "Not Okay")
		}
		return c.String(http.StatusOK, "Auth API response")
	})

	e.Logger.Fatal(e.Start(":8889"))
}

func isOK() bool {
	return rand.Float32() < 0.5
}
