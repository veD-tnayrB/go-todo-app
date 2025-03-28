package repository

import (
	baseErrors "errors"
)

var (
	ErrRecordNotExists     = baseErrors.New("record not exists")
	ErrRecordAlreadyExists = baseErrors.New("record already exists")
)
