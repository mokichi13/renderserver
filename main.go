package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	e.GET("/time", func(c echo.Context) error {
		now := strconv.Itoa(int(time.Now().Unix()))
		return c.JSON(http.StatusOK, now)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
