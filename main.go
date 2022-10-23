package main

import (
	"os"

	"github.com/dlacreme/gotifications/api"
	"github.com/dlacreme/gotifications/env"
	"github.com/labstack/echo/v4"
)

func main() {
	env.LoadAndCheck()
	e := echo.New()
	api.RegisterMiddlewares(e)
	api.RegisterEndpoints(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
