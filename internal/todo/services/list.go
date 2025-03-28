package service

import (
	"errors"

	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) List() (*[]*models.Todo, error) {
	todos, err := s.TodoRepository.GetAll()
	if err != nil {
		return nil, errors.New("ERROR_WHILE_GETTING_THE_TODOS")
	}

	return todos, nil
}
