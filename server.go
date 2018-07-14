package main

import (
	"github.com/labstack/echo"
	"github.com/ntrv/foods-database/handler"
)

func main() {
	e := echo.New()
	e.GET("/", handler.AllFoods())
	e.Logger.Fatal(e.Start(":1323"))
}
