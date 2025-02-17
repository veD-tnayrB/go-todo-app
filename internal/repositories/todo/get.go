package repository

import "github.com/veD-tnayrB/todo-app/common/models"

func (r *TodoRepository) GetById(id string) (*models.Todo, error) {
	todo := r.DB[id]
	return &todo, nil
}
