package service

import (
	baseErrors "errors"
)

var (
	ErrTitleIsRequired     = baseErrors.New("title is required")
	ErrIdIsRequired        = baseErrors.New("id is required")
	ErrorSaving            = baseErrors.New("something went wrong while trying to save the record")
	ErrorUpdating          = baseErrors.New("something went wrong while trying to update the record")
	ErrorRemoving          = baseErrors.New("something went wrong while trying to remove the record")
	ErrRecordNotExists     = baseErrors.New("record not exists")
	ErrRecordAlreadyExists = baseErrors.New("record already exists")
	ErrGettingTodo         = baseErrors.New("error getting the todo")
)
