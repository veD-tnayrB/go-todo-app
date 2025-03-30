package db

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

type DB map[string]models.Todo

func NewDB() DB {
	return DB{}
}
