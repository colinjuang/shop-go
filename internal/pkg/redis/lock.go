package redis

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Errors
var (
	ErrLockAcquireFailed = errors.New("failed to acquire lock")
	ErrLockReleaseFailed = errors.New("failed to release lock")
)

// Lock represents a Redis-based distributed lock
type Lock struct {
	client    *Client
	key       string
	value     string
	expiry    time.Duration
	acquired  bool
	retryTime time.Duration
	retries   int
}

// NewLock creates a new distributed lock
func NewLock(key string, expiry time.Duration) *Lock {
	return &Lock{
		client:    GetClient(),
		key:       fmt.Sprintf("lock:%s", key),
		value:     fmt.Sprintf("%d", time.Now().UnixNano()),
		expiry:    expiry,
		acquired:  false,
		retryTime: 100 * time.Millisecond,
		retries:   5,
	}
}

// TryAcquire tries to acquire the lock once without retrying
func (l *Lock) TryAcquire(ctx context.Context) (bool, error) {
	if l.acquired {
		return true, nil
	}

	acquired, err := l.client.SetNX(ctx, l.key, l.value, l.expiry)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrLockAcquireFailed, err)
	}

	l.acquired = acquired
	return acquired, nil
}

// Acquire acquires the lock with retries
func (l *Lock) Acquire(ctx context.Context) error {
	if l.acquired {
		return nil
	}

	for i := 0; i < l.retries; i++ {
		acquired, err := l.TryAcquire(ctx)
		if err != nil {
			return err
		}

		if acquired {
			return nil
		}

		// Wait before retrying
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(l.retryTime):
			// Continue with retry
		}
	}

	return ErrLockAcquireFailed
}

// Release releases the lock
// It only releases if the lock value matches the original value
// This prevents releasing a lock that has expired and been acquired by another process
func (l *Lock) Release(ctx context.Context) error {
	if !l.acquired {
		return nil
	}

	// For our simple Redis client, we'll do this non-atomically
	// In production, you'd want to use EVAL to run a Lua script for atomic operation:
	// if redis.call("GET", KEYS[1]) == ARGV[1] then return redis.call("DEL", KEYS[1]) else return 0 end
	val, err := l.client.Get(ctx, l.key)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLockReleaseFailed, err)
	}

	if val == l.value {
		err := l.client.Delete(ctx, l.key)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrLockReleaseFailed, err)
		}
		l.acquired = false
		return nil
	}

	return nil
}

// WithLock executes a function within a lock
func WithLock(ctx context.Context, key string, expiry time.Duration, fn func() error) error {
	lock := NewLock(key, expiry)

	err := lock.Acquire(ctx)
	if err != nil {
		return err
	}

	defer lock.Release(ctx)

	return fn()
}
