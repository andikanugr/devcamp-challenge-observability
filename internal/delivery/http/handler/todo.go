package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"todoapp/internal/entity"
	"todoapp/pkg/logger"
)

// GetTodos is a handler for getting all todos
// @Summary Get todos
// @Description Get todos
// @ID get-todos
// @Accept  json
// @Produce  json
// @Router /todos [get]
func (h *Handler) GetTodos(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "delivery.http.handler.GetTodos")
	defer span.Finish()

	todos, err := h.todoUC.FetchTodos(ctx)
	if err != nil {
		h.log.WithField(logger.Fields{
			"error": err.Error(),
		}).Error("failed to get todos")

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed to get todos",
		})
	}

	return c.JSON(200, todos)
}

// AddTodo is a handler for adding new todo
// @Summary Add todo
// @Description Add todo
// @ID add-todo
// @Accept  json
// @Produce  json
// @Param todo body entity.Todo true "Todo"
// @Router /todos [post]
func (h *Handler) AddTodo(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "delivery.http.handler.AddTodo")
	defer span.Finish()

	var todo entity.Todo
	if err := c.Bind(&todo); err != nil {
		h.log.WithField(logger.Fields{
			"error": err.Error(),
		}).Error("invalid request body on adding new todo")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request body",
		})
	}

	if err := h.todoUC.AddTodo(ctx, todo); err != nil {
		h.log.WithField(logger.Fields{
			"todo":  todo,
			"error": err.Error(),
		}).Error("failed to add todo on adding new todo")

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed to add todo",
		})
	}

	return c.JSON(200, todo)
}

// GetTodoByID is a handler for getting todo by id
// @Summary Get todo by id
// @Description Get todo by id
// @ID get-todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Router /todos/{id} [get]
func (h *Handler) GetTodoByID(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "delivery.http.handler.GetTodoByID")
	defer span.Finish()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.WithField(logger.Fields{
			"id":    idStr,
			"error": err.Error(),
		}).Error("id is not a valid integer on getting todo by id")

		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "id is not a valid integer",
		})
	}

	todo, err := h.todoUC.GetTodoByID(ctx, id)
	if err != nil {
		h.log.WithField(logger.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("failed to get todo on getting todo by id")

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed to get todo",
		})
	}

	return c.JSON(200, todo)
}

// UpdateTodoByID is a handler for updating todo
// @Summary Update todo
// @Description Update todo
// @ID update-todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body entity.Todo true "Todo"
// @Router /todos/{id} [put]
func (h *Handler) UpdateTodoByID(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "delivery.http.handler.UpdateTodoByID")
	defer span.Finish()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.WithField(logger.Fields{
			"id":    idStr,
			"error": err.Error(),
		}).Error("id is not a valid integer on updating todo")

		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "id is not a valid integer",
		})
	}

	var todo entity.Todo
	if err := c.Bind(&todo); err != nil {
		h.log.WithField(logger.Fields{
			"id":    id,
			"todo":  todo,
			"error": err.Error(),
		}).Error("invalid request body on updating todo")

		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request body",
		})
	}

	if err := h.todoUC.UpdateTodoByID(ctx, id, todo); err != nil {
		h.log.WithField(logger.Fields{
			"id":    id,
			"todo":  todo,
			"error": err.Error(),
		}).Error("failed to update todo on updating todo")
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed to update todo",
		})
	}

	return c.JSON(200, todo)
}

// DeleteTodoByID is a handler for deleting todo
// @Summary Delete todo
// @Description Delete todo
// @ID delete-todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Router /todos/{id} [delete]
func (h *Handler) DeleteTodoByID(c echo.Context) error {
	span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "delivery.http.handler.DeleteTodoByID")
	defer span.Finish()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.WithField(logger.Fields{
			"id":    idStr,
			"error": err.Error(),
		}).Error("id is not a valid integer on deleting todo")

		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "id is not a valid integer",
		})
	}

	if err := h.todoUC.DeleteTodoByID(ctx, id); err != nil {
		h.log.WithField(logger.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("failed to delete todo on deleting todo")

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed to delete todo",
		})
	}

	return c.JSON(200, nil)
}
