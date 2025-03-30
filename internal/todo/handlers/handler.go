package handler

import (
	"github.com/veD-tnayrB/todo-app/common/logger"
	"github.com/veD-tnayrB/todo-app/common/models"
)

type TodoService interface {
	List() ([]*models.Todo, error)
	Get(id string) (*models.Todo, error)
	Create(params *models.Todo) error
	Update(params *models.Todo) error
	Remove(id string) error
}
type TodoHandler struct {
	TodoService TodoService
	Logger      logger.LoggerContract
}

func NewTodoHandler(service TodoService, logger *logger.Logger) (*TodoHandler, error) {
	if service == nil {
		return nil, ErrServiceRequired
	}

	if logger == nil {
		return nil, ErrLoggerRequired
	}

	return &TodoHandler{TodoService: service, Logger: logger}, nil
}
