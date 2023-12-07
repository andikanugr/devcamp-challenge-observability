package middleware

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func prometheusMiddleware() echo.MiddlewareFunc {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests, labeled by success or error",
		},
		[]string{"method", "endpoint", "status", "is_error"},
	)
	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Register the defined metrics with Prometheus
	prometheus.MustRegister(requestCounter, requestDuration)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Start timer
			timer := prometheus.NewTimer(requestDuration.WithLabelValues(c.Request().Method, c.Path()))

			// Process request
			err := next(c)

			// Stop timer
			timer.ObserveDuration()

			// Increment request counter
			// Determine the value of is_error label
			isErrorLabel := "false"
			if err != nil {
				isErrorLabel = "true"
			}

			status := c.Response().Status
			requestCounter.WithLabelValues(c.Request().Method, c.Path(), strconv.Itoa(status), isErrorLabel).Inc()

			return err
		}
	}
}
