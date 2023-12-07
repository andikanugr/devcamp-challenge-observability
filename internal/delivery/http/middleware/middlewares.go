package middleware

import "github.com/labstack/echo/v4"

type Middlewares []echo.MiddlewareFunc

func ProvideMiddleware() Middlewares {
	return Middlewares{
		prometheusMiddleware(),
		jaegerMiddleware(),
	}
}
