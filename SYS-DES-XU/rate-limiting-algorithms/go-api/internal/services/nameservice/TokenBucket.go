package nameservice

import (
	"sync"
	"time"
)

// TokenBucket struct defines the token bucket
type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   time.Duration
	lastRefilled time.Time
	mu           sync.Mutex
}

// NewTokenBucket creates a new Token Bucket.
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity, // Start with full capacity
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

// Allow checks if a token is available and consumes it if so.
func (tb *TokenBucket) Allow() bool {
	// Ensure thread safety
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Refill tokens based on time elapsed since last refill
	now := time.Now()
	elapsed := now.Sub(tb.lastRefilled)
	// update lastRefilled to now
	tb.lastRefilled = now

	refillTokens := int(elapsed / tb.refillRate)
	if refillTokens > 0 {
		tb.tokens += refillTokens
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
	}

	// Check and consume token
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucket) SecondsUntilNextRequest() int {
	// Ensure thread safety
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Calculate the time elapsed since the last refill
	now := time.Now()
	elapsed := now.Sub(tb.lastRefilled)

	// Calculate how many tokens should have been refilled during that time
	refillTokens := int(elapsed / tb.refillRate)

	// If there are no tokens available, calculate the time needed to refill one token
	if tb.tokens == 0 {
		// Calculate the time required to refill one token
		secondsLeft := int(tb.refillRate.Seconds()) * (1 - refillTokens) // Corrected calculation
		return secondsLeft
	}

	return 0 // There is at least one token available, so no need to wait
}
