package errors

import (
	baseErrors "errors"
)

var (
	TitleIsRequired = baseErrors.New("Title is required")
	IdIsRequired    = baseErrors.New("Id is required")
	ErrorSaving     = baseErrors.New("Something went wrong while trying to save the record")
	ErrorUpdating   = baseErrors.New("Something went wrong while trying to update the record")
	ErrorRemoving   = baseErrors.New("Something went wrong while trying to remove the record")
)
