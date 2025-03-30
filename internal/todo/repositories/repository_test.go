package repository_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/logger"
	"github.com/veD-tnayrB/todo-app/common/models"
	repository "github.com/veD-tnayrB/todo-app/internal/todo/repositories"
)

func setupRepo(t *testing.T) *repository.TodoRepository {
	DB := make(db.DB)
	logger, err := logger.NewLogger("tests")
	if err != nil {
		t.Fatalf("logger wasnt created, cant proceed with the case: %s\n", err)
	}

	instance, err := repository.NewTodoRepository(DB, logger)
	if err != nil {
		t.Fatalf("repository instance wasnt created, cant proceed with the case: %s\n", err)
	}

	return instance
}

func TestInitializaton(t *testing.T) {
	t.Run("Check instanciation API", func(t *testing.T) {
		t.Run("Check with Logger in nil", func(t *testing.T) {
			_, err := repository.NewTodoRepository(make(db.DB), nil)
			if err == nil {
				t.Error("instance can be created without a logger")
			}

			if err != repository.ErrLoggerRequired {
				t.Errorf("error message is not ErrLoggerRequired, instead is: %s\n", err.Error())
			}
		})

		t.Run("Check with DB in nil", func(t *testing.T) {
			logger, err := logger.NewLogger("tests")
			if err != nil {
				t.Errorf("logger wasnt created, cant proceed with the case: %s\n", err)
			}

			_, err = repository.NewTodoRepository(nil, logger)
			if err == nil {
				t.Errorf("instance can be created without a logger")
			}

			if err != repository.ErrDBRequired {
				t.Errorf("error message is not ErrDBRequired, instead is: %s\n", err.Error())
			}
		})
	})

}

func TestGet(t *testing.T) {
	instance := setupRepo(t)
	registry := models.Todo{
		Id:        "1",
		Title:     "Study",
		Completed: false,
	}
	instance.DB["1"] = registry

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
			expectedError: repository.ErrRecordNotExists,
			expectedTodo:  nil,
		},
	}

	for _, testingCase := range tests {
		t.Run(testingCase.name, func(t *testing.T) {
			register, err := instance.GetById(testingCase.id)
			if !errors.Is(err, testingCase.expectedError) {
				t.Log(1)
				t.Fatalf("unexpected error: %s\n", err)
			}

			if !assert.Equal(t, register, testingCase.expectedTodo) {
				t.Log(3)
				t.Fatalf("invalid output, expecting for: %v | obtained: %v \n", testingCase.expectedTodo, register)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	repo := setupRepo(t)

	tests := []struct {
		name          string
		todo          *models.Todo
		expectedError error
	}{
		{
			name:          "Insert a new Todo successfully",
			todo:          &models.Todo{Id: "1", Title: "Learn Go", Completed: false},
			expectedError: nil,
		},
		{
			name:          "Insert fails when ID already exists",
			todo:          &models.Todo{Id: "1", Title: "Duplicate Todo", Completed: true},
			expectedError: repository.ErrRecordAlreadyExists, // ID "1" was inserted in the previous case
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Insert(tc.todo)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			// Check if Todo was actually inserted
			if err == nil {
				inserted, exists := repo.DB[tc.todo.Id]
				assert.True(t, exists, "Todo should exist in the DB")
				assert.Equal(t, tc.todo.Title, inserted.Title, "Inserted title mismatch")
				assert.Equal(t, tc.todo.Completed, inserted.Completed, "Inserted completion status mismatch")
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	repo := setupRepo(t)

	// Pre-populate with one entry
	existingTodo := &models.Todo{Id: "1", Title: "Learn Go", Completed: false}
	repo.DB[existingTodo.Id] = *existingTodo

	tests := []struct {
		name          string
		id            string
		updatedTodo   *models.Todo
		expectedError error
	}{
		{
			name:          "Update an existing Todo",
			id:            "1",
			updatedTodo:   &models.Todo{Id: "1", Title: "Master Go", Completed: true},
			expectedError: nil,
		},
		{
			name:          "Fail to update a non-existent Todo",
			id:            "2",
			updatedTodo:   &models.Todo{Id: "2", Title: "New Task", Completed: false},
			expectedError: repository.ErrRecordNotExists,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Update(tc.id, tc.updatedTodo)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			// Verify update only if it succeeded
			if err == nil {
				updated, exists := repo.DB[tc.id]
				assert.True(t, exists, "Updated Todo should exist")
				assert.Equal(t, tc.updatedTodo.Title, updated.Title, "Updated title mismatch")
				assert.Equal(t, tc.updatedTodo.Completed, updated.Completed, "Updated completion status mismatch")
			}
		})
	}
}

func TestRemove(t *testing.T) {
	repo := setupRepo(t)

	// Pre-populate DB
	existingTodo := models.Todo{Id: "1", Title: "Learn Go", Completed: false}
	repo.DB[existingTodo.Id] = existingTodo

	tests := []struct {
		name          string
		id            string
		expectedError error
	}{
		{
			name:          "Remove an existing Todo",
			id:            "1",
			expectedError: nil,
		},
		{
			name:          "Fail to remove a non-existent Todo",
			id:            "2",
			expectedError: repository.ErrRecordNotExists,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Remove(tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			// Verify deletion only if remove succeeded
			if err == nil {
				_, exists := repo.DB[tc.id]
				assert.False(t, exists, "Todo should no longer exist in the DB")
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	repo := setupRepo(t) // Fresh repository

	t.Run("Returns empty list when no Todos exist", func(t *testing.T) {
		todos, err := repo.GetAll()

		assert.Nil(t, err, "Error should be nil")
		assert.NotNil(t, todos, "Todos should not be nil")
		assert.Equal(t, 0, len(todos), "List should be empty")
	})

	t.Run("Returns all existing Todos", func(t *testing.T) {
		repo.DB["1"] = models.Todo{Id: "1", Title: "Learn Go", Completed: false}
		repo.DB["2"] = models.Todo{Id: "2", Title: "Write Tests", Completed: true}

		todos, err := repo.GetAll()

		assert.Nil(t, err, "Error should be nil")
		assert.NotNil(t, todos, "Todos should not be nil")
		assert.Equal(t, 2, len(todos), "Should return 2 todos")

		// Check each Todo
		todoMap := map[string]*models.Todo{
			todos[0].Id: todos[0],
			todos[1].Id: todos[1],
		}

		assert.Contains(t, todoMap, "1", "Todo ID 1 should be present")
		assert.Contains(t, todoMap, "2", "Todo ID 2 should be present")
		assert.Equal(t, "Learn Go", todoMap["1"].Title, "Todo 1 title mismatch")
		assert.Equal(t, true, todoMap["2"].Completed, "Todo 2 completion status mismatch")
	})
}
