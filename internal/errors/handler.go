package errors

import (
	baseErrors "errors"
)

var (
	MalformedBody = baseErrors.New("Malformed body")
)
