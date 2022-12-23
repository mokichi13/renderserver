package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// set middleware
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
