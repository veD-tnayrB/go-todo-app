package repository

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (r *TodoRepository) GetById(id string) (*models.Todo, error) {
	if id == "" {
		return nil, ErrIdIsRequired
	}

	todo, exists := r.DB[id]
	if !exists {
		return nil, ErrRecordNotExists
	}
	return &todo, nil
}
