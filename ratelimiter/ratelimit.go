package ratelimiter

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"web4app/locale" // import your locale package
)

// --- Custom rate limits ---
type customRateLimit struct {
	suffix   string
	requests int
	reset    time.Duration
}

// --- Bucket ---
type Bucket struct {
	sync.Mutex
	Key             string
	Locale          locale.Locale // <-- locale-aware
	Remaining       int
	reset           time.Time
	global          *int64
	lastReset       time.Time
	customRateLimit *customRateLimit
	Userdata        interface{}
	backoffCount    int
}

func (b *Bucket) getRemaining() int {
	b.Lock()
	defer b.Unlock()
	return b.Remaining
}

// --- RateLimiter ---
type RateLimiter struct {
	sync.Mutex
	global           *int64
	buckets          map[string]map[locale.Locale]*Bucket // <-- map of buckets per locale
	customRateLimits []*customRateLimit
}

// NewRatelimiter initializes RateLimiter
func NewRatelimiter() *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]map[locale.Locale]*Bucket),
		global:  new(int64),
		customRateLimits: []*customRateLimit{
			{
				suffix:   "//reactions//",
				requests: 1,
				reset:    200 * time.Millisecond,
			},
		},
	}
}

// GetBucket retrieves or creates a bucket for a specific locale
func (r *RateLimiter) GetBucket(key string, loc locale.Locale) *Bucket {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.buckets[key]; !ok {
		r.buckets[key] = make(map[locale.Locale]*Bucket)
	}

	if b, ok := r.buckets[key][loc]; ok {
		return b
	}

	b := &Bucket{
		Remaining: 1,
		Key:       key,
		Locale:    loc,
		global:    r.global,
	}

	for _, rl := range r.customRateLimits {
		if strings.HasSuffix(b.Key, rl.suffix) {
			b.customRateLimit = rl
			break
		}
	}

	r.buckets[key][loc] = b
	return b
}

// GetWaitTime returns non-blocking remaining wait
func (r *RateLimiter) GetWaitTime(b *Bucket, minRemaining int) time.Duration {
	b.Lock()
	defer b.Unlock()

	if b.Remaining < minRemaining && b.reset.After(time.Now()) {
		return b.reset.Sub(time.Now())
	}

	sleepTo := time.Unix(0, atomic.LoadInt64(r.global))
	if now := time.Now(); now.Before(sleepTo) {
		return sleepTo.Sub(now)
	}

	return 0
}

// LockBucketObject locks with context support
func (r *RateLimiter) LockBucketObject(ctx context.Context, b *Bucket) error {
	for {
		wait := r.GetWaitTime(b, 1)
		if wait > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(wait):
			}
		}

		b.Lock()
		if b.Remaining > 0 {
			b.Remaining--
			b.Unlock()
			return nil
		}
		b.Unlock()
	}
}

// Release updates bucket based on headers
func (b *Bucket) Release(headers http.Header) error {
	b.Lock()
	defer b.Unlock()

	if rl := b.customRateLimit; rl != nil {
		if time.Since(b.lastReset) >= rl.reset {
			b.Remaining = rl.requests
			b.lastReset = time.Now()
			b.backoffCount = 0
		} else if b.Remaining < 1 {
			b.reset = time.Now().Add(rl.reset)
		}
		return nil
	}

	if headers == nil {
		return nil
	}

	remaining := headers.Get("X-RateLimit-Remaining")
	resetAfter := headers.Get("X-RateLimit-Reset-After")
	global := headers.Get("X-RateLimit-Global")

	if remaining != "" {
		if r, err := strconv.Atoi(remaining); err == nil {
			b.Remaining = r
		}
	}

	if resetAfter != "" {
		if parsed, err := strconv.ParseFloat(resetAfter, 64); err == nil {
			whole, frac := math.Modf(parsed)
			resetAt := time.Now().Add(time.Duration(whole) * time.Second).Add(time.Duration(frac*1000) * time.Millisecond)
			if global != "" {
				atomic.StoreInt64(b.global, resetAt.UnixNano())
			} else {
				b.reset = resetAt
			}
		}
	}

	// Exponential backoff on repeated 429s
	if b.Remaining == 0 {
		b.backoffCount++
		backoff := time.Duration(math.Pow(2, float64(b.backoffCount))) * time.Second
		if backoff > 60*time.Second {
			backoff = 60 * time.Second
		}
		b.reset = time.Now().Add(backoff)
	}

	return nil
}
