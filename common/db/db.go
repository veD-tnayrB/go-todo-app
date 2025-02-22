package db

import (
	"github.com/veD-tnayrB/todo-app/common/models"
)

type DB map[string]models.Todo

func NewDB() DB {
	// this is returning the direction of where db is stored, thats correct
	return DB{}
}
