package repository

import (
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/internal/errors"
)

func (r *TodoRepository) GetById(id string) (*models.Todo, error) {
	todo, exists := r.DB[id]
	if exists {
		return nil, errors.RecordNotExists
	}
	return &todo, nil
}
