package usecase

import (
	"context"
	"todoapp/internal/entity"
)

type TodoUsecase interface {
	FetchTodos(ctx context.Context) ([]entity.Todo, error)
	AddTodo(ctx context.Context, todo entity.Todo) error
	GetTodoByID(ctx context.Context, id int) (entity.Todo, error)
	UpdateTodoByID(ctx context.Context, id int, todo entity.Todo) error
	DeleteTodoByID(ctx context.Context, id int) error
}
