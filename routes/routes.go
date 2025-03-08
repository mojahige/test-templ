package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mojahige/test-templ/handlers"
)

func SetupRoutes(e *echo.Echo) {
	setupWebRoutes(e)
	setupStaticRoutes(e)
	setupErrorRoutes(e)
}

func setupWebRoutes(e *echo.Echo) {
	homeHandler := handlers.NewHomeHandler()
	postsHandler := handlers.NewPostsHandler()

	e.GET("/", homeHandler.Get)
	e.GET("/posts", postsHandler.Get)
	e.POST("/posts", postsHandler.Post)
	e.PATCH("/posts", postsHandler.Patch)
	e.DELETE("/posts", postsHandler.Delete)
}

func setupStaticRoutes(e *echo.Echo) {
	e.Static("/assets", "assets")
}

func setupErrorRoutes(e *echo.Echo) {
	errorHandler := handlers.NewErrorHandler()
	e.HTTPErrorHandler = errorHandler.Handle
}
