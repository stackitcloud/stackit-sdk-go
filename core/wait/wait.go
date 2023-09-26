package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

var RetryHttpErrorStatusCodes = []int{http.StatusBadGateway, http.StatusGatewayTimeout}

type WaitFn func() (res interface{}, done bool, err error)

type Handler struct {
	fn                WaitFn
	sleepBeforeWait   time.Duration
	throttle          time.Duration
	timeout           time.Duration
	retryLimitTempErr int
}

// New creates a new Wait instance
func New(f WaitFn) *Handler {
	return &Handler{
		fn:                f,
		sleepBeforeWait:   0 * time.Second,
		throttle:          5 * time.Second,
		timeout:           30 * time.Minute,
		retryLimitTempErr: 10,
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

// SetRetryLimitTempErr sets the retry limit if a temporary error is found. The list of temporary errors is defined in the RetryHttpErrorStatusCodes variable
func (w *Handler) SetRetryLimitTempErr(l int) *Handler {
	w.retryLimitTempErr = l
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

	var retryTempErrorCounter = 0
	for {
		res, done, err = w.fn()
		retryTempErrorCounter, err = handleError(retryTempErrorCounter, w.retryLimitTempErr, err)
		if err != nil {
			return res, err
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

func handleError(retryTempErrorCounter, retryLimitTempErr int, err error) (int, error) {
	if err != nil {
		oapiErr, ok := err.(*oapiError.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return retryTempErrorCounter, fmt.Errorf("could not convert error to GenericOpenApiError, %w", err)
		}
		// Some APIs may return temporary errors and the request should be retried
		if utils.Contains(RetryHttpErrorStatusCodes, oapiErr.StatusCode) {
			retryTempErrorCounter++
			if retryTempErrorCounter == retryLimitTempErr {
				return retryTempErrorCounter, fmt.Errorf("temporary error was found and the retry limit was reached: %w", err)
			}
			return retryTempErrorCounter, nil
		} else {
			return retryTempErrorCounter, fmt.Errorf("executing wait function: %w", err)
		}
	}
	return retryTempErrorCounter, nil
}
