package api

import (
	"net/http"

	"github.com/dlacreme/gotifications/email"
	"github.com/labstack/echo/v4"
)

type EmailRequest struct {
	From    string   `json:"from" validate:"required"`
	To      []string `json:"to" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
	Body    string   `json:"body" validate:"required"`
}

func sendEmail(c echo.Context) error {
	data := &EmailRequest{}

	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := email.SendEmail(data.From, data.To, data.Subject, data.Body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "Email has been sent.")
}
