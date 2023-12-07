package todo

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"math/rand"
	"todoapp/internal/entity"
	"todoapp/internal/repository"
)

// Usecase is todo usecase
type Usecase struct {
	todoRepo repository.TodoRepository
}

// New returns new todo usecase
func New(repo repository.TodoRepository) *Usecase {
	return &Usecase{
		todoRepo: repo,
	}
}

// FetchTodos returns all todos
func (u *Usecase) FetchTodos(ctx context.Context) ([]entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.todo.FetchTodos")
	defer span.Finish()

	todos, err := u.todoRepo.GetTodos(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// AddTodo adds new todo
func (u *Usecase) AddTodo(ctx context.Context, todo entity.Todo) error {
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "usecase.todo.AddTodo")
	defer span.Finish()

	err := u.todoRepo.CreateTodo(ctx, todo)
	if err != nil {
		return err
	}

	return nil
}

// GetTodoByID returns todo by id
func (u *Usecase) GetTodoByID(ctx context.Context, id int) (entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.todo.GetTodoByID")
	defer span.Finish()

	if rand.Float32() < 0.5 { // 50% chance to fail
		return entity.Todo{}, errors.New("intermittent error")
	}
	todo, err := u.todoRepo.GetTodoByID(ctx, id)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

// UpdateTodoByID updates todo by id
func (u *Usecase) UpdateTodoByID(ctx context.Context, id int, todo entity.Todo) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.todo.UpdateTodoByID")
	defer span.Finish()

	err := u.todoRepo.UpdateTodoByID(ctx, id, todo)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodoByID deletes todo by id
func (u *Usecase) DeleteTodoByID(ctx context.Context, id int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.todo.DeleteTodoByID")
	defer span.Finish()

	err := u.todoRepo.DeleteTodoByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
