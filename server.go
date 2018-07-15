package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ntrv/foods-database/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/", handler.AllFoods())
	e.POST("/", handler.RegisterFood())

	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(s))
}
