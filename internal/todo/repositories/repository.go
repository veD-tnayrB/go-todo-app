package repository

import (
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/logger"
)

type TodoRepository struct {
	DB     db.DB
	Logger logger.LoggerContract
}

func NewTodoRepository(DB db.DB, Logger logger.LoggerContract) (*TodoRepository, error) {
	if DB == nil {
		return nil, ErrDBRequired
	}

	if Logger == nil {
		return nil, ErrLoggerRequired
	}

	return &TodoRepository{
		DB,
		Logger,
	}, nil
}
