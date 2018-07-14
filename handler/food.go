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
