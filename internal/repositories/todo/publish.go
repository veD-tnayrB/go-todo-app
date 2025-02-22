package repository

import (
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/internal/errors"
)

func (r *TodoRepository) Insert(params *models.Todo) error {
	if _, exists := r.DB[params.Id]; exists {
		return errors.RecordAlreadyExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepository) Update(id string, params *models.Todo) error {
	if _, exists := r.DB[params.Id]; !exists {
		return errors.RecordNotExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepository) Remove(id string) error {
	if _, exists := r.DB[id]; !exists {
		return errors.RecordNotExists
	}

	delete(r.DB, id)
	return nil
}
