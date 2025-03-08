package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mojahige/test-templ/db"
	"github.com/mojahige/test-templ/routes"
)

func main() {
	db.Init()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":4000"))
}
