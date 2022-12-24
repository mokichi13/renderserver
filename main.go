package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var u struct {
	index int
	id    []int
	m     sync.Mutex
}

func create() int {
	u.m.Lock()
	defer u.m.Unlock()

	u.index++
	u.id = append(u.id, u.index)
	return u.index
}

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	e.GET("/time", func(c echo.Context) error {
		now := strconv.Itoa(int(time.Now().Unix()))
		return c.JSON(http.StatusOK, now)
	})
	e.POST("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, fmt.Sprintf("{id: %d}", create()))
	})
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, u.id)
	})

	// set middleware
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
