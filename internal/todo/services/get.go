package service

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) Get(id string) (*models.Todo, error) {
	s.Logger.Info("Service: Handling the GET method")

	if id == "" {
		s.Logger.Warn("Service: Missing id in request", "id", id, "error", ErrIdIsRequired)
		return nil, ErrIdIsRequired
	}

	todo, err := s.TodoRepository.GetById(id)
	if err != nil {
		s.Logger.Error("Service: Error getting the todo", "id", id, "error", err)
		return nil, ErrGettingTodo
	}

	return todo, nil
}
