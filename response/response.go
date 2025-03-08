package response

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func ComponentToHTML(ctx context.Context, t templ.Component) (string, error) {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx, buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func HTML(ctx echo.Context, statusCode int, t templ.Component) error {
	html, err := ComponentToHTML(ctx.Request().Context(), t)
	if err != nil {
		return err
	}

	return ctx.HTML(statusCode, html)
}
