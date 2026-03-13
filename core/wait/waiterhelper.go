package wait

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// WaiterHelper is a helper struct, which creates based on the configured attributes the AsyncActionCheck for wait.New
type WaiterHelper[T any, S comparable] struct {
	// FetchInstance is called periodically to get the latest resource data.
	FetchInstance func() (*T, error)

	// GetState extracts the status string from the API response.
	GetState func(*T) (S, error)

	// ActiveState represents the terminal "Happy Path" (e.g. ACTIVE, READY, CREATED).
	ActiveState []S

	// ErrorState represents the terminal "Error Path" (e.g. FAILED, ERROR).
	ErrorState []S

	// DeleteHttpErrorStatusCodes defines codes treated as a successful deletion (default: 403, 404, 410).
	DeleteHttpErrorStatusCodes []int
}

var defaultHttpErrorStatusCodes = []int{http.StatusForbidden, http.StatusNotFound, http.StatusGone}

func (w *WaiterHelper[T, S]) Wait() AsyncActionCheck[T] {
	if len(w.DeleteHttpErrorStatusCodes) == 0 {
		w.DeleteHttpErrorStatusCodes = append(w.DeleteHttpErrorStatusCodes, defaultHttpErrorStatusCodes...)
	}

	return func() (waitFinished bool, response *T, err error) {
		instance, err := w.FetchInstance()
		if err != nil {
			var oapiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiErr) {
				// If no active states are defined and one of the "Delete HTTP Status codes" is returned, finish wait without an error
				if len(w.ActiveState) == 0 && slices.Contains(w.DeleteHttpErrorStatusCodes, oapiErr.StatusCode) {
					return true, nil, nil
				}
			}
			return true, nil, err
		}

		state, err := w.GetState(instance)
		if err != nil {
			return true, nil, err
		}

		// 1. Check if the operation succeeded
		if slices.Contains(w.ActiveState, state) {
			return true, instance, nil
		}

		// 2. Check if the operation failed
		if slices.Contains(w.ErrorState, state) {
			return true, instance, fmt.Errorf("waiting failed. state is %v", state)
		}

		// 3. Default: Operation is pending
		return false, nil, nil
	}
}
