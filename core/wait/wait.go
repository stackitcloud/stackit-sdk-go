package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var RetryHttpErrorStatusCodes = []int{http.StatusBadGateway, http.StatusGatewayTimeout}

type WaitFn func() (res interface{}, done bool, err error)

type Handler struct {
	fn              WaitFn
	sleepBeforeWait time.Duration
	throttle        time.Duration
	timeout         time.Duration
}

// New creates a new Wait instance
func New(f WaitFn) *Handler {
	return &Handler{
		fn:              f,
		sleepBeforeWait: 0 * time.Second,
		throttle:        5 * time.Second,
		timeout:         30 * time.Minute,
	}
}

// SetThrottle sets the duration between func triggering
func (w *Handler) SetThrottle(d time.Duration) error {
	if d == 0 {
		return fmt.Errorf("throttle can't be 0")
	}
	w.throttle = d
	return nil
}

// SetTimeout sets the duration for wait timeout
func (w *Handler) SetTimeout(d time.Duration) *Handler {
	w.timeout = d
	return w
}

// SetSleepBeforeWait sets the duration for sleep before wait
func (w *Handler) SetSleepBeforeWait(d time.Duration) *Handler {
	w.sleepBeforeWait = d
	return w
}

// WaitWithContext starts the wait until there's an error or wait is done
func (w *Handler) WaitWithContext(ctx context.Context) (res interface{}, err error) {
	var done bool

	ctx, cancel := context.WithTimeout(ctx, w.timeout)
	defer cancel()

	// Wait some seconds for the API to process the request
	time.Sleep(w.sleepBeforeWait)

	ticker := time.NewTicker(w.throttle)
	defer ticker.Stop()

	for {
		res, done, err = w.fn()
		if err != nil {
			return res, fmt.Errorf("executing wait function: %w", err)
		}
		if done {
			return res, nil
		}

		select {
		case <-ticker.C:
			// continue
		case <-ctx.Done():
			return res, fmt.Errorf("WaitWithContext() has timed out")
		}
	}
}
