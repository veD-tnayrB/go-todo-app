package repository

import (
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/logger"
)

type TodoRepository struct {
	DB     db.DB
	Logger logger.LoggerContract
}
