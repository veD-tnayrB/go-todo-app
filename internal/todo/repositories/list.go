package repository

import "github.com/veD-tnayrB/todo-app/common/models"

func (r *TodoRepository) GetAll() ([]*models.Todo, error) {
	todos := []*models.Todo{}

	for index := range r.DB {
		todo := r.DB[index]
		todos = append(todos, &todo)
	}

	return todos, nil
}
