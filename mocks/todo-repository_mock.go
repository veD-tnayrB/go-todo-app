package repository_mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/veD-tnayrB/todo-app/common/models"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (r *TodoRepositoryMock) GetById(id string) (*models.Todo, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Todo), nil
}

func (r *TodoRepositoryMock) List() (*[]*models.Todo, error) {
	args := r.Called()

	return args.Get(0).(*[]*models.Todo), nil
}
