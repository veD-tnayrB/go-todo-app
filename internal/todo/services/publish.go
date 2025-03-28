package service

import (
	"github.com/google/uuid"
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) Create(param *models.Todo) error {
	param.Id = uuid.NewString()

	if param.Title == "" {
		return ErrTitleIsRequired
	}

	err := s.TodoRepository.Insert(param)
	if err != nil {
		if err == ErrRecordAlreadyExists {
			return err
		}

		return ErrorSaving
	}

	return nil
}

func (s *TodoService) Update(param *models.Todo) error {
	if param.Id == "" {
		return ErrIdIsRequired
	}

	if param.Title == "" {
		return ErrTitleIsRequired
	}

	err := s.TodoRepository.Update(param.Id, param)
	if err != nil {
		if err == ErrRecordNotExists {
			return err
		}

		return ErrorUpdating
	}

	return nil
}

func (s *TodoService) Remove(id string) error {
	if id == "" {
		return ErrIdIsRequired
	}

	err := s.TodoRepository.Remove(id)
	if err != nil {
		if err == ErrRecordNotExists {
			return err
		}
		return ErrorRemoving
	}

	return nil
}
