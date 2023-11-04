package userboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserBoard(c echo.Context) error {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Welcome to userboard",
	})
	return nil
}
