package service

import (
	"errors"

	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) Get(id string) (*models.Todo, error) {
	if id == "" {
		return nil, errors.New("ID_IS_REQUIRED")
	}
	todo, err := s.TodoRepository.GetById(id)
	if err != nil {
		return nil, errors.New("ERROR_WHILE_GETTING_THE_TODO")
	}

	return todo, nil
}
