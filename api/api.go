package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMiddlewares(e *echo.Echo) {

}

func RegisterEndpoints(e *echo.Echo) {
	e.GET("/", health)
	e.POST("/email", sendEmail)
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "<Notifications>: Im alive.")
}
