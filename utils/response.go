package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// JSONResponse - упрощенный ответ в JSON
func JSONResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
