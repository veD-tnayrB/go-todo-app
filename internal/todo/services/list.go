package service

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) List() (*[]*models.Todo, error) {
	s.Logger.Info("Service: Handling the LIST method")
	todos, err := s.TodoRepository.GetAll()

	if err != nil {
		s.Logger.Error("Service: Error while trying to get the todo list", "error", err)
		return nil, ErrGettingTodo
	}

	return todos, nil
}
