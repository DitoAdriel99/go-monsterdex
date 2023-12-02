package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Check Health
// @Description Check server is running
// @ID check-health
// @Router /check-health [get]
func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
