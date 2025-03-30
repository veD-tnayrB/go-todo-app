package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/logger"
	"github.com/veD-tnayrB/todo-app/common/models"
	repository "github.com/veD-tnayrB/todo-app/internal/todo/repositories"
	service "github.com/veD-tnayrB/todo-app/internal/todo/services"
)

func setupService(t *testing.T) (*service.TodoService, db.DB) {
	logger, err := logger.NewLogger("tests")
	if err != nil {
		t.Fatalf("logger wasnt created, cant proceed with the case: %s\n", err)
	}

	DB := make(db.DB)
	repo, err := repository.NewTodoRepositoryMock(DB)
	if err != nil {
		t.Fatalf("repo wasnt created, cant proceed with the case: %s\n", err)
	}

	instance, err := service.NewTodoService(repo, logger)
	if err != nil {
		t.Fatalf("service wasnt created, cant proceed with the case: %s\n", err)
	}

	return instance, DB
}

func TestInitializaton(t *testing.T) {
	t.Run("Check instanciation API", func(t *testing.T) {
		t.Run("Check with Logger in nil", func(t *testing.T) {
			DB := make(db.DB)
			repo, err := repository.NewTodoRepositoryMock(DB)
			if err != nil {
				t.Errorf("unexpected error occured: %s\n", err)
			}

			_, err = service.NewTodoService(repo, nil)
			if !errors.Is(err, service.ErrLoggerIsRequired) {
				t.Errorf("error message is not ErrLoggerRequired, instead is: %s\n", err.Error())
			}
		})

		t.Run("Check with repo in nil", func(t *testing.T) {
			logger, err := logger.NewLogger("tests")
			if err != nil {
				t.Errorf("logger wasnt created, cant proceed with the case: %s\n", err)
			}

			_, err = service.NewTodoService(nil, logger)
			if !errors.Is(err, service.ErrRepositoryIsRequired) {
				t.Errorf("error message is not ErrRepositoryIsRequired, instead is: %s\n", err.Error())
			}
		})
	})

}

func TestGet(t *testing.T) {
	instance, db := setupService(t)
	registry := models.Todo{
		Id:        "1",
		Title:     "Study",
		Completed: false,
	}
	db["1"] = registry

	tests := []struct {
		name          string
		id            string
		expectedError error
		expectedTodo  *models.Todo
	}{
		{
			name:          "Existing already in DB case",
			id:            "1",
			expectedError: nil,
			expectedTodo:  &registry,
		},
		{
			name:          "Not existing in DB case",
			id:            "2",
			expectedError: service.ErrRecordNotExists,
			expectedTodo:  nil,
		},
		{
			name:          "Not sending ID",
			id:            "",
			expectedError: service.ErrIdIsRequired,
			expectedTodo:  nil,
		},
	}

	for _, testingCase := range tests {
		t.Run(testingCase.name, func(t *testing.T) {
			register, err := instance.Get(testingCase.id)
			if !errors.Is(err, testingCase.expectedError) {
				t.Fatalf("unexpected error: %s\n", err)
			}

			if !assert.Equal(t, register, testingCase.expectedTodo) {
				t.Fatalf("invalid output, expecting for: %v | obtained: %v \n", testingCase.expectedTodo, register)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	instance, _ := setupService(t)

	tests := []struct {
		name          string
		expectedError error
		param         *models.Todo
	}{
		{
			name:          "Normal insert",
			expectedError: nil,
			param:         &models.Todo{Title: "New todo", Completed: false},
		},
		{
			name:          "Empty title",
			expectedError: service.ErrTitleIsRequired,
			param:         &models.Todo{Title: "", Completed: false},
		},
	}

	for _, ttc := range tests {
		t.Run(ttc.name, func(t *testing.T) {
			err := instance.Create(ttc.param)

			if !errors.Is(err, ttc.expectedError) {
				t.Errorf("invalid error, expected: '%s' | received: '%s'\n", ttc.expectedError, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	instance, db := setupService(t)
	db["1"] = models.Todo{Id: "1", Title: "Do homework", Completed: false}

	tests := []struct {
		name          string
		expectedError error
		param         *models.Todo
	}{
		{
			name:          "Normal update",
			expectedError: nil,
			param:         &models.Todo{Id: "1", Title: "Do homework", Completed: true},
		},
		{
			name:          "Updating not existing record",
			expectedError: service.ErrRecordNotExists,
			param:         &models.Todo{Id: "5", Title: "Do homework", Completed: true},
		},
		{
			name:          "Empty title",
			expectedError: service.ErrTitleIsRequired,
			param:         &models.Todo{Id: "1", Title: "", Completed: false},
		},
		{
			name:          "Empty id",
			expectedError: service.ErrIdIsRequired,
			param:         &models.Todo{Title: "", Completed: false},
		},
	}

	for _, ttc := range tests {
		t.Run(ttc.name, func(t *testing.T) {
			err := instance.Update(ttc.param)

			if !errors.Is(err, ttc.expectedError) {
				t.Errorf("invalid error, expected: '%s' | received: '%s'\n", ttc.expectedError, err)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	instance, db := setupService(t)
	db["1"] = models.Todo{Id: "1", Title: "Do homework", Completed: false}

	tests := []struct {
		name          string
		expectedError error
		id            string
	}{
		{
			name:          "Normal remove",
			expectedError: nil,
			id:            "1",
		},
		{
			name:          "Empty id",
			expectedError: service.ErrIdIsRequired,
			id:            "",
		},
		{
			name:          "Removing not existing record",
			expectedError: service.ErrRecordNotExists,
			id:            "2",
		},
	}

	for _, ttc := range tests {
		t.Run(ttc.name, func(t *testing.T) {
			err := instance.Remove(ttc.id)

			if !errors.Is(err, ttc.expectedError) {
				t.Errorf("invalid error, expected: '%s' | received: '%s'\n", ttc.expectedError, err)
			}
		})
	}
}
