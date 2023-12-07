package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

func jaegerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "http_request")
			defer span.Finish()

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
