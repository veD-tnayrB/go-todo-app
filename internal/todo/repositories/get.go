package repository

import (
	"fmt"

	"github.com/veD-tnayrB/todo-app/common/models"
)

func (r *TodoRepository) GetById(id string) (*models.Todo, error) {
	if id == "" {
		return nil, ErrIdIsRequired
	}

	fmt.Printf("HIGH TECH: %v\n", r.DB)
	fmt.Printf("HIGH TECH 2: %v %T\n", id, id)
	todo, exists := r.DB[id]
	if !exists {
		return nil, ErrRecordNotExists
	}
	return &todo, nil
}
