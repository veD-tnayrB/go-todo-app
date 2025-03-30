package repository

import (
	baseErrors "errors"
)

var (
	ErrRecordNotExists     = baseErrors.New("record not exists")
	ErrRecordAlreadyExists = baseErrors.New("record already exists")
	ErrIdIsRequired        = baseErrors.New("id is required")
	ErrDBRequired          = baseErrors.New("a db is required for todo repository")
	ErrLoggerRequired      = baseErrors.New("logger is required for todo repository")
)
