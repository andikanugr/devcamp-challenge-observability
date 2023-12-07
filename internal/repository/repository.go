package repository

import (
	"context"
	"todoapp/internal/entity"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo entity.Todo) error
	GetTodos(ctx context.Context) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, id int) (entity.Todo, error)
	UpdateTodoByID(ctx context.Context, id int, todo entity.Todo) error
	DeleteTodoByID(ctx context.Context, id int) error
}
