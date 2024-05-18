package views

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(t templ.Component, httpCode int, ctx echo.Context) error {
	buf := templ.GetBuffer()

	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(httpCode, buf.String())
}
