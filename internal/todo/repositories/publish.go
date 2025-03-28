package repository

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (r *TodoRepository) Insert(params *models.Todo) error {
	if _, exists := r.DB[params.Id]; exists {
		return ErrRecordAlreadyExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepository) Update(id string, params *models.Todo) error {
	if _, exists := r.DB[params.Id]; !exists {
		return ErrRecordNotExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepository) Remove(id string) error {
	if _, exists := r.DB[id]; !exists {
		return ErrRecordNotExists
	}

	delete(r.DB, id)
	return nil
}
