package db

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

// i dont remember how to initialize a map in go...
type DB map[string]models.Todo

func NewDB() *DB {
	return &DB{}
}
