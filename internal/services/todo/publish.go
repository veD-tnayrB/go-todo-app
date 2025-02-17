package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) Create(param *models.Todo) error {
	param.Id = uuid.NewString()

	if param.Title == "" {
		return errors.New("TITLE_IS_REQUIRED")
	}

	err := s.TodoRepository.Insert(param)
	if err != nil {
		return errors.New("ERROR_WHILE_SAVING_THE_TODO")
	}

	return nil
}

func (s *TodoService) Update(param *models.Todo) error {
	if param.Id == "" {
		return errors.New("ID_IS_REQUIRED")
	}

	if param.Title == "" {
		return errors.New("TITLE_IS_REQUIRED")
	}

	err := s.TodoRepository.Update(param.Id, param)
	if err != nil {
		return errors.New("ERROR_WHILE_SAVING_THE_CHANGES")
	}

	return nil
}

func (s *TodoService) Remove(id string) error {
	if id == "" {
		return errors.New("ID_IS_REQUIRED")
	}

	err := s.TodoRepository.Remove(id)
	if err != nil {
		return errors.New("ERROR_WHILE_REMOVING_THE_TODO")
	}

	return nil
}
