package service

import (
	"github.com/veD-tnayrB/todo-app/common/logger"
	"github.com/veD-tnayrB/todo-app/common/models"
)

type TodoRepository interface {
	GetAll() (*[]*models.Todo, error)
	GetById(id string) (*models.Todo, error)
	Insert(params *models.Todo) error
	Update(id string, params *models.Todo) error
	Remove(id string) error
}

type TodoService struct {
	TodoRepository TodoRepository
	Logger         logger.LoggerContract
}
