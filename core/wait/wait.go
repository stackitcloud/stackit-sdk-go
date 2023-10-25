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

// AsyncActionCheck reports whether a specific async action has finished.
//   - waitFinished == true if the async action is finished, false otherwise.
//   - response contains data regarding the current state of the resource targeted by the async action (if applicable). resource != nil if waitFinished == true.
//   - err != nil if there was an error checking if the aync action finished, or if it finished unsuccessfully.
type AsyncActionCheck[T any] func() (waitFinished bool, response *T, err error)

// AsyncActionHandler handles waiting for a specific async action to be finished.
type AsyncActionHandler[T any] struct {
	check             AsyncActionCheck[T]
	sleepBeforeWait   time.Duration
	throttle          time.Duration
	timeout           time.Duration
	tempErrRetryLimit int
}

// New creates a new AsyncHandler instance
func New[T any](f AsyncActionCheck[T]) *AsyncActionHandler[T] {
	return &AsyncActionHandler[T]{
		check:             f,
		sleepBeforeWait:   0 * time.Second,
		throttle:          5 * time.Second,
		timeout:           30 * time.Minute,
		tempErrRetryLimit: 5,
	}
}

// SetThrottle sets the duration between func triggering
func (w *AsyncActionHandler[T]) SetThrottle(d time.Duration) error {
	if d == 0 {
		return fmt.Errorf("throttle can't be 0")
	}
	w.throttle = d
	return nil
}

// SetTimeout sets the duration for wait timeout
func (w *AsyncActionHandler[T]) SetTimeout(d time.Duration) *AsyncActionHandler[T] {
	w.timeout = d
	return w
}

// SetSleepBeforeWait sets the duration for sleep before wait
func (w *AsyncActionHandler[T]) SetSleepBeforeWait(d time.Duration) *AsyncActionHandler[T] {
	w.sleepBeforeWait = d
	return w
}

// SetRetryLimitTempErr sets the retry limit if a temporary error is found. The list of temporary errors is defined in the RetryHttpErrorStatusCodes variable
func (w *AsyncActionHandler[T]) SetRetryLimitTempErr(l int) *AsyncActionHandler[T] {
	w.tempErrRetryLimit = l
	return w
}

// WaitWithContext starts the wait until there's an error or wait is done
func (w *AsyncActionHandler[T]) WaitWithContext(ctx context.Context) (res *T, err error) {
	ctx, cancel := context.WithTimeout(ctx, w.timeout)
	defer cancel()

	// Wait some seconds for the API to process the request
	time.Sleep(w.sleepBeforeWait)

	ticker := time.NewTicker(w.throttle)
	defer ticker.Stop()

	var retryTempErrorCounter = 0
	for {
		done, res, err := w.check()
		if err != nil {
			retryTempErrorCounter, err = w.handleError(retryTempErrorCounter, err)
			if err != nil {
				return res, err
			}
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

func (w *AsyncActionHandler[T]) handleError(retryTempErrorCounter int, err error) (int, error) {
	oapiErr, ok := err.(*oapiError.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
	if !ok {
		return retryTempErrorCounter, fmt.Errorf("could not convert error to GenericOpenApiError, %w", err)
	}
	// Some APIs may return temporary errors and the request should be retried
	if utils.Contains(RetryHttpErrorStatusCodes, oapiErr.StatusCode) {
		retryTempErrorCounter++
		if retryTempErrorCounter == w.tempErrRetryLimit {
			return retryTempErrorCounter, fmt.Errorf("temporary error was found and the retry limit was reached: %w", err)
		}
		return retryTempErrorCounter, nil
	}
	return retryTempErrorCounter, fmt.Errorf("executing wait function: %w", err)
}
