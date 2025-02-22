package errors

import (
	baseErrors "errors"
)

var (
	RecordNotExists     = baseErrors.New("Record not exists")
	RecordAlreadyExists = baseErrors.New("Record already exists")
)
