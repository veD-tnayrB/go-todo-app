package rate_limiter

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	type Test struct {
		Tokens      int
		RefreshTime time.Duration
		Requests    int
		ExpectPass  int // Expected successful requests
	}

	tests := []Test{
		{Tokens: 10, RefreshTime: 10 * time.Second, Requests: 15, ExpectPass: 10}, // Exceeds available tokens
		{Tokens: 5, RefreshTime: 5 * time.Second, Requests: 5, ExpectPass: 5},     // Uses exactly available tokens
		{Tokens: 0, RefreshTime: 10 * time.Second, Requests: 5, ExpectPass: 0},    // No tokens available
	}

	for _, tt := range tests {
		rateLimiter, err := NewRateLimiter(tt.Tokens, tt.RefreshTime)
		if tt.RefreshTime <= 0 && !errors.Is(err, ErrInvalidRefreshTime) {
			t.Log("missing minimun refresh time validations")
			return
		}

		if tt.RefreshTime <= 0 && errors.Is(err, ErrInvalidRefreshTime) {
			t.Log("missing minimun refresh time validations PASSES")
			return
		}

		if tt.Tokens <= 0 && !errors.Is(err, ErrInvalidTokensNumber) {
			t.Log("missing minimun number of token validations")
			return
		}

		if tt.Tokens <= 0 && errors.Is(err, ErrInvalidTokensNumber) {
			t.Log("missing minimun number of token validations PASSES")
			return
		}

		if err != nil {
			t.Logf("Unexpected error: %v\n", err)
			return
		}

		wg := sync.WaitGroup{}
		passCount := 0
		mu := sync.Mutex{} // Protects passCount from concurrent writes

		for i := 0; i < tt.Requests; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if rateLimiter.CheckSpace() {
					mu.Lock()
					passCount += 1
					mu.Unlock()
				}
			}()
		}

		wg.Wait()

		if passCount != tt.ExpectPass {
			t.Errorf("Expected %d successful requests, but got %d", tt.ExpectPass, passCount)
		}

		t.Logf("Test with %d tokens and %d requests passed. (%d successful)", tt.Tokens, tt.Requests, passCount)
	}
}
