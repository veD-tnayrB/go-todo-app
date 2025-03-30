package rate_limiter

import (
	baseError "errors"
)

var (
	ErrInvalidTokensNumber = baseError.New("invalid tokens number for rate limiter, must be > 0")
	ErrInvalidRefreshTime  = baseError.New("invalid refresh time for rate limiter, must be bigger than 0 seconds")
)
