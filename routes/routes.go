package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mojahige/test-templ/handlers"
)

func SetupRoutes(e *echo.Echo) {
	setupWebRoutes(e)
	setupErrorRoutes(e)
}

func setupWebRoutes(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
}

func setupErrorRoutes(e *echo.Echo) {
	e.HTTPErrorHandler = handlers.ErrorHandler
}
