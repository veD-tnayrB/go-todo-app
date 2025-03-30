package rate_limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu              sync.Mutex
	Tokens          int
	MaxTokens       int
	RefreshTime     time.Duration
	LastTimeChecked time.Time
}

func NewRateLimiter(tokens int, refreshTime time.Duration) (*RateLimiter, error) {
	if tokens <= 0 {
		return nil, ErrInvalidTokensNumber
	}

	if refreshTime <= 0 {
		return nil, ErrInvalidRefreshTime
	}

	return &RateLimiter{MaxTokens: tokens, Tokens: tokens, RefreshTime: refreshTime, LastTimeChecked: time.Now()}, nil
}

func (r *RateLimiter) CheckSpace() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	passedTime := time.Since(r.LastTimeChecked)
	if passedTime >= r.RefreshTime {
		r.handleNewToken()
		r.LastTimeChecked = time.Now()

	}

	if r.Tokens > 0 {
		r.Tokens = r.Tokens - 1
		return true
	}
	return false
}

func (r *RateLimiter) handleNewToken() {
	r.Tokens = r.MaxTokens
}
