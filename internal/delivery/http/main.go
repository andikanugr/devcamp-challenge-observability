package http

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"todoapp/internal/delivery/http/handler"
	"todoapp/internal/delivery/http/middleware"

	_ "todoapp/internal/delivery/http/docs"
)

type Server struct {
	echo        *echo.Echo
	handler     *handler.Handler
	middlewares middleware.Middlewares
}

// NewServer is a constructor for Server
// @title Swagger Todo App API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8080
// @BasePath /api
func NewServer(
	handler *handler.Handler,
	middlewares middleware.Middlewares,
) *Server {
	return &Server{
		handler:     handler,
		echo:        echo.New(),
		middlewares: middlewares,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(":8080")
}

func (s *Server) Stop() error {
	return s.echo.Close()
}

func (s *Server) RegisterHandler() {
	handler := s.handler

	// swagger docs
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// metrics instrumentation
	s.echo.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Create a group for /api
	api := s.echo.Group("/api")

	api.GET("/todos", handler.GetTodos)
	api.POST("/todos", handler.AddTodo)
	api.GET("/todos/:id", handler.GetTodoByID)
	api.PUT("/todos/:id", handler.UpdateTodoByID)
	api.DELETE("/todos/:id", handler.DeleteTodoByID)
}

func (s *Server) RegisterMiddleware() {
	s.echo.Use(s.middlewares...)
}
