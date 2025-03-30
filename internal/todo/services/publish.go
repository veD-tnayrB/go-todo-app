package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (s *TodoService) Create(param *models.Todo) error {
	s.Logger.Info("Service: Handling the CREATE method")
	param.Id = uuid.NewString()

	if param.Title == "" {
		s.Logger.Warn("Service: title is required", "error", ErrTitleIsRequired)
		return ErrTitleIsRequired
	}

	err := s.TodoRepository.Insert(param)
	if err != nil {
		if err == ErrRecordAlreadyExists {
			s.Logger.Warn("Service: Record to create already exists", "param", param)
			return err
		}

		s.Logger.Error("Service: Error creating the todo", "param", param, "error", err)
		return ErrorSaving
	}

	return nil
}

func (s *TodoService) Update(param *models.Todo) error {
	s.Logger.Info("Service: Handling the CREATE method")

	if param.Id == "" {
		s.Logger.Warn("Service: title is required", "error", ErrTitleIsRequired)
		return ErrIdIsRequired
	}

	if param.Title == "" {
		return ErrTitleIsRequired
	}

	err := s.TodoRepository.Update(param.Id, param)
	if err != nil {
		fmt.Printf("1: \"%s\" | \"%s\" | \"%v\"\n", err, ErrRecordNotExists, errors.Is(err, ErrRecordNotExists))
		if errors.Is(err, ErrRecordNotExists) {
			fmt.Printf("2: %s\n", err)

			s.Logger.Warn("Service: Record to updating dont exists", "param", param)
			return err
		}

		s.Logger.Error("Service: Error updating the todo", "param", param, "error", err)
		return ErrorUpdating
	}

	return nil
}

func (s *TodoService) Remove(id string) error {
	if id == "" {
		s.Logger.Warn("Service: id is required", "error", ErrIdIsRequired)
		return ErrIdIsRequired
	}

	err := s.TodoRepository.Remove(id)
	if err != nil {
		if err == ErrRecordNotExists {
			s.Logger.Warn("Service: Record to remove dont exists", "id", id)
			return err
		}

		s.Logger.Error("Service: Error to remove the todo", "id", id, "error", err)
		return ErrorRemoving
	}

	return nil
}
