package ratelimit

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
)

// ------------------------
// Custom rate limit definition
// ------------------------
type customRateLimit struct {
	suffix   string
	requests int
	reset    time.Duration
}

// ------------------------
// Bucket: handles queued requests
// ------------------------
type Bucket struct {
	Key             string
	Remaining       int
	lastReset       time.Time
	reset           time.Time
	customRateLimit *customRateLimit
	backoff         int
	global          *int64

	mu    sync.Mutex
	queue chan requestTask
}

type requestTask struct {
	ctx context.Context
	fn  func() (*http.Response, error)
	ch  chan result
}

type result struct {
	resp *http.Response
	err  error
}

// ------------------------
// RateLimiter: manages all buckets
// ------------------------
type RateLimiter struct {
	buckets          map[string]*Bucket
	global           *int64
	customRateLimits []*customRateLimit
	mu               sync.Mutex
}

// ------------------------
// NewRateLimiter
// ------------------------
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]*Bucket),
		global:  new(int64),
		customRateLimits: []*customRateLimit{
			{suffix: "//reactions//", requests: 1, reset: 200 * time.Millisecond},
		},
	}
}

// ------------------------
// GetBucket
// ------------------------
func (r *RateLimiter) GetBucket(key string) *Bucket {
	r.mu.Lock()
	defer r.mu.Unlock()

	if b, ok := r.buckets[key]; ok {
		return b
	}

	b := &Bucket{
		Key:       key,
		Remaining: 1,
		global:    r.global,
		queue:     make(chan requestTask, 10000), // allow thousands of queued requests
	}

	for _, rl := range r.customRateLimits {
		if strings.HasSuffix(key, rl.suffix) {
			b.customRateLimit = rl
			break
		}
	}

	r.buckets[key] = b
	go b.worker() // spawn worker for async execution

	return b
}

// ------------------------
// worker: handles queued requests asynchronously
// ------------------------
func (b *Bucket) worker() {
	for task := range b.queue {
		select {
		case <-task.ctx.Done():
			task.ch <- result{nil, task.ctx.Err()}
			continue
		default:
		}

		// wait if rate-limited
		b.mu.Lock()
		wait := b.getWaitTime(1)
		b.mu.Unlock()

		if wait > 0 {
			select {
			case <-time.After(wait):
			case <-task.ctx.Done():
				task.ch <- result{nil, task.ctx.Err()}
				continue
			}
		}

		// Execute the request
		resp, err := task.fn()
		if err != nil {
			task.ch <- result{nil, err}
			continue
		}

		// Release and update rate limit info
		_ = b.Release(resp.Header)

		task.ch <- result{resp, nil}
	}
}

// ------------------------
// Submit: enqueue request
// ------------------------
func (b *Bucket) Submit(ctx context.Context, fn func() (*http.Response, error)) (*http.Response, error) {
	ch := make(chan result, 1)
	task := requestTask{
		ctx: ctx,
		fn:  fn,
		ch:  ch,
	}

	select {
	case b.queue <- task:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	select {
	case res := <-ch:
		return res.resp, res.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ------------------------
// GetRemaining: get quota
// ------------------------
func (b *Bucket) GetRemaining() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Remaining
}

// ------------------------
// Release: updates bucket from headers
// ------------------------
func (b *Bucket) Release(headers http.Header) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()

	// Custom rate limit
	if rl := b.customRateLimit; rl != nil {
		if now.Sub(b.lastReset) >= rl.reset {
			b.Remaining = rl.requests
			b.lastReset = now
			b.backoff = 0
		}
		return nil
	}

	if headers == nil {
		return nil
	}

	remaining := headers.Get("X-RateLimit-Remaining")
	resetAfter := headers.Get("X-RateLimit-Reset-After")
	global := headers.Get("X-RateLimit-Global")

	if resetAfter != "" {
		parsed, err := strconv.ParseFloat(resetAfter, 64)
		if err != nil {
			return err
		}
		whole, frac := math.Modf(parsed)
		resetTime := now.Add(time.Duration(whole)*time.Second + time.Duration(frac*1e9))
		if global != "" {
			atomic.StoreInt64(b.global, resetTime.UnixNano())
		} else {
			b.reset = resetTime
		}
	}

	if remaining != "" {
		val, err := strconv.Atoi(remaining)
		if err != nil {
			return err
		}
		b.Remaining = val
		b.backoff = 0
	} else if b.Remaining < 1 {
		// exponential backoff if repeated 429
		b.backoff++
		backoffDuration := time.Duration(math.Pow(2, float64(b.backoff))) * time.Second
		b.reset = now.Add(backoffDuration)
	}

	return nil
}

// ------------------------
// getWaitTime: compute non-blocking wait
// ------------------------
func (b *Bucket) getWaitTime(minRemaining int) time.Duration {
	now := time.Now()
	if b.Remaining < minRemaining && b.reset.After(now) {
		return b.reset.Sub(now)
	}

	globalTime := time.Unix(0, atomic.LoadInt64(b.global))
	if now.Before(globalTime) {
		return globalTime.Sub(now)
	}

	return 0
}
