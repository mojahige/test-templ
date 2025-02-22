package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mojahige/test-templ/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":4000"))
}
