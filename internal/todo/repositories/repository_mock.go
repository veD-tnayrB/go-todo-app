package repository

import (
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/models"
)

type TodoRepositoryMock struct {
	DB db.DB
}

func NewTodoRepositoryMock(DB db.DB) (*TodoRepositoryMock, error) {
	if DB == nil {
		return nil, ErrDBRequired
	}
	return &TodoRepositoryMock{DB: DB}, nil
}

func (r *TodoRepositoryMock) GetById(id string) (*models.Todo, error) {
	if id == "" {
		return nil, ErrIdIsRequired
	}

	todo, exists := r.DB[id]
	if !exists {
		return nil, ErrRecordNotExists
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
		return ErrRecordAlreadyExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepositoryMock) Update(id string, params *models.Todo) error {
	if _, exists := r.DB[params.Id]; !exists {
		return ErrRecordNotExists
	}

	r.DB[params.Id] = *params
	return nil
}

func (r *TodoRepositoryMock) Remove(id string) error {
	if _, exists := r.DB[id]; !exists {
		return ErrRecordNotExists
	}

	delete(r.DB, id)
	return nil
}
