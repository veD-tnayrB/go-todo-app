package service

import (
	"github.com/google/uuid"
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/internal/errors"
)

func (s *TodoService) Create(param *models.Todo) error {
	param.Id = uuid.NewString()

	if param.Title == "" {
		return errors.TitleIsRequired
	}

	err := s.TodoRepository.Insert(param)
	if err != nil {
		if err == errors.RecordAlreadyExists {
			return err
		}

		return errors.ErrorSaving
	}

	return nil
}

func (s *TodoService) Update(param *models.Todo) error {
	if param.Id == "" {
		return errors.IdIsRequired
	}

	if param.Title == "" {
		return errors.TitleIsRequired
	}

	err := s.TodoRepository.Update(param.Id, param)
	if err != nil {
		if err == errors.RecordNotExists {
			return err
		}

		return errors.ErrorUpdating
	}

	return nil
}

func (s *TodoService) Remove(id string) error {
	if id == "" {
		return errors.IdIsRequired
	}

	err := s.TodoRepository.Remove(id)
	if err != nil {
		if err == errors.RecordNotExists {
			return err
		}
		return errors.ErrorRemoving
	}

	return nil
}
