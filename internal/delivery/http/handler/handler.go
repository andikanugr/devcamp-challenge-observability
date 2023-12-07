package handler

import (
	"todoapp/internal/usecase"
	"todoapp/pkg/logger"
)

type Handler struct {
	todoUC usecase.TodoUsecase
	log    logger.Logger
}

func NewHandler(
	todoUC usecase.TodoUsecase,
	log logger.Logger,
) *Handler {
	return &Handler{
		todoUC: todoUC,
		log:    log,
	}
}
