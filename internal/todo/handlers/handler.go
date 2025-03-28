package handler

import (
	"github.com/veD-tnayrB/todo-app/common/logger"
	"github.com/veD-tnayrB/todo-app/common/models"
)

type TodoService interface {
	List() (*[]*models.Todo, error)
	Get(id string) (*models.Todo, error)
	Create(params *models.Todo) error
	Update(params *models.Todo) error
	Remove(id string) error
}
type TodoHandler struct {
	TodoService TodoService
	Logger      logger.LoggerContract
}
