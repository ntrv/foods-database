package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ntrv/foods-database/models"
)

func AllFoods() echo.HandlerFunc {
	return func(c echo.Context) error {
		food := models.Food{}
		foods, err := food.AllFoods()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, foods)
	}
}

func RegisterFood() echo.HandlerFunc {
	return func(c echo.Context) error {
		food := models.Food{}
		food.Name = c.FormValue("name")
		if err := food.RegisterFood(); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, food)
	}
}
