package repository_test

import (
	"github.com/veD-tnayrB/todo-app/common/models"
	repository "github.com/veD-tnayrB/todo-app/internal/todo/repositories"
)

type TodoRepositoryMock struct {
	DB map[string]models.Todo
}

func (r *TodoRepositoryMock) GetById(id string) (*models.Todo, error) {
	if id == "" {
		return nil, repository.ErrIdIsRequired
	}

	todo, exists := r.DB[id]
	if !exists {
		return nil, repository.ErrRecordNotExists
	}
	return &todo, nil
}

func (r *TodoRepositoryMock) GetAll() ([]*models.Todo, error) {
	todos := []*models.Todo{}

	for index := range r.DB {
		todo := r.DB[index]
		todos = append(todos, &todo)
	}

	return todos, nil
}

func (r *TodoRepositoryMock) Insert(params *models.Todo) error {
	if _, exists := r.DB[params.Id]; exists {
		return repository.ErrRecordAlreadyExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepositoryMock) Update(id string, params *models.Todo) error {
	if _, exists := r.DB[params.Id]; !exists {
		return repository.ErrRecordNotExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepositoryMock) Remove(id string) error {
	if _, exists := r.DB[id]; !exists {
		return repository.ErrRecordNotExists
	}

	delete(r.DB, id)
	return nil
}
