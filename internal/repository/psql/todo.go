package psql

import (
	"context"
	"math/rand"
	"time"

	"github.com/opentracing/opentracing-go"
	"todoapp/internal/entity"
)

type TodoRepo struct {
	db *Database
}

func NewTodoRepository(db *Database) *TodoRepo {
	return &TodoRepo{
		db: db,
	}
}

func (r *TodoRepo) CreateTodo(ctx context.Context, todo entity.Todo) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.todo.CreateTodo")
	defer span.Finish()

	query := `INSERT INTO todo (title, description, priority, is_done) 
	          VALUES (:title, :description, :priority, :is_done)`

	_, err := r.db.NamedExecContext(ctx, query, todo)
	if err != nil {
		// Handle the error properly here
		return err
	}

	return nil
}

func (r *TodoRepo) GetTodos(ctx context.Context) ([]entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.todo.GetTodos")
	defer span.Finish()

	query := `SELECT id, title, description, priority, is_done  FROM todo`

	r.slowFunc()

	var todos []entity.Todo
	err := r.db.SelectContext(ctx, &todos, query)
	if err != nil {
		// Handle the error properly here
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepo) GetTodoByID(ctx context.Context, id int) (entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.todo.GetTodoByID")
	defer span.Finish()

	query := `SELECT id, title, description, priority, is_done  FROM todo WHERE id = $1`

	var todo entity.Todo
	err := r.db.GetContext(ctx, &todo, query, id)
	if err != nil {
		// Handle the error properly here
		return entity.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepo) UpdateTodoByID(ctx context.Context, id int, todo entity.Todo) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.todo.UpdateTodoByID")
	defer span.Finish()

	query := `UPDATE todo SET title = :title, description = :description, priority = :priority, is_done = :is_done WHERE id = :id`

	_, err := r.db.NamedExecContext(ctx, query, todo)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepo) DeleteTodoByID(ctx context.Context, id int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.todo.DeleteTodoByID")
	defer span.Finish()

	query := `DELETE FROM todo WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepo) slowFunc() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
