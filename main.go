package main

import (
	"os"

	"github.com/dlacreme/did-notifications/api"
	"github.com/dlacreme/did-notifications/env"
	"github.com/labstack/echo/v4"
)

func main() {
	env.LoadAndCheck()
	e := echo.New()
	api.RegisterMiddlewares(e)
	api.RegisterEndpoints(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
