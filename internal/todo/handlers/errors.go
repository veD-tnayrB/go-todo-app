package handler

import (
	baseErrors "errors"
)

var (
	ErrMalformedBody   = baseErrors.New("malformed body")
	ErrIdIsRequired    = baseErrors.New("id is required")
	ErrRecordNotExists = baseErrors.New("record not exists")
	ErrErrorRemoving   = baseErrors.New("something went wrong while removing")
	ErrTitleIsRequired = baseErrors.New("title is required")
	ErrErrorSaving     = baseErrors.New("error saving")
	ErrServiceRequired = baseErrors.New("service is required for todo handlers")
	ErrLoggerRequired  = baseErrors.New("logger is required for todo repository")
)
