package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mojahige/test-templ/components/pages"
	"github.com/mojahige/test-templ/response"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Get(c echo.Context) error {
	return response.HTML(c, http.StatusOK, pages.HomePage())
}
