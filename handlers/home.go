package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mojahige/test-templ/components/pages"
	"github.com/mojahige/test-templ/render"
)

func HomeHandler(c echo.Context) error {
	return render.Component(c, http.StatusOK, pages.HomePage())
}
