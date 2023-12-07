package app

import (
	"github.com/google/wire"
	"todoapp/internal/delivery/http"
	"todoapp/internal/delivery/http/handler"
	"todoapp/internal/delivery/http/middleware"
	"todoapp/internal/repository"
	"todoapp/internal/repository/psql"
	"todoapp/internal/usecase"
	"todoapp/internal/usecase/todo"
	"todoapp/pkg/logger"
)

var (
	middlewareSet = wire.NewSet(
		middleware.ProvideMiddleware,
	)

	repositorySet = wire.NewSet(
		psql.NewDatabase,
		psql.NewTodoRepository,
		wire.Bind(new(repository.TodoRepository), new(*psql.TodoRepo)),
	)

	usecaseSet = wire.NewSet(
		todo.New,
		wire.Bind(new(usecase.TodoUsecase), new(*todo.Usecase)),
	)

	httpSet = wire.NewSet(
		logger.NewLogger,
		middlewareSet,
		repositorySet,
		usecaseSet,
		handler.NewHandler,
		http.NewServer,
	)
)
