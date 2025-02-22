package handlers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mojahige/test-templ/components/pages"
	"github.com/mojahige/test-templ/render"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	c.Logger().Error(err)

	if he, ok := err.(*echo.HTTPError); ok {
		message := fmt.Sprintf("%v", he.Message)
		render.Component(c, he.Code, pages.ErrorPage(pages.ErrorPageProps{
			ErrorCode: strconv.Itoa(he.Code),
			Message:   message,
		}))
	}
}
