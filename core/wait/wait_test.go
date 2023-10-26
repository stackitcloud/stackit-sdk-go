package wait

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// Options used for comparing AsyncActionHandler
var cmpOpts = []cmp.Option{
	cmp.AllowUnexported(AsyncActionHandler[interface{}]{}),
	cmpopts.IgnoreFields(AsyncActionHandler[interface{}]{}, "checkFn"),
}

func TestNew(t *testing.T) {
	checkFn := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }
	got := New(checkFn)
	want := &AsyncActionHandler[interface{}]{
		checkFn:           checkFn,
		sleepBeforeWait:   0 * time.Second,
		throttle:          5 * time.Second,
		timeout:           30 * time.Minute,
		tempErrRetryLimit: 5,
	}

	diff := cmp.Diff(got, want, cmpOpts...)
	if diff != "" {
		t.Errorf("Data does not match: %s", diff)
	}
}

func TestSetThrottle(t *testing.T) {
	checkFn := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }

	for _, tt := range []struct {
		desc     string
		throttle time.Duration
	}{
		{
			"base_1",
			time.Hour,
		},
		{
			"base_2",
			10 * time.Millisecond,
		},
		{
			"base_3",
			0,
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			want := New(checkFn)
			want.throttle = tt.throttle
			got := New(checkFn)
			got.SetThrottle(tt.throttle)

			diff := cmp.Diff(got, want, cmpOpts...)
			if diff != "" {
				t.Errorf("Data does not match: %s", diff)
			}
		})
	}
}

func TestSetTimeout(t *testing.T) {
	checkFn := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }

	for _, tt := range []struct {
		desc    string
		timeout time.Duration
	}{
		{
			"base_1",
			time.Hour,
		},
		{
			"base_2",
			10 * time.Millisecond,
		},
		{
			"base_3",
			0,
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			want := New(checkFn)
			want.timeout = tt.timeout
			got := New(checkFn)
			got.SetTimeout(tt.timeout)

			diff := cmp.Diff(got, want, cmpOpts...)
			if diff != "" {
				t.Errorf("Data does not match: %s", diff)
			}
		})
	}
}

func TestSetSleepBeforeWait(t *testing.T) {
	checkFn := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }

	for _, tt := range []struct {
		desc            string
		sleepBeforeWait time.Duration
	}{
		{
			"base_1",
			time.Hour,
		},
		{
			"base_2",
			10 * time.Millisecond,
		},
		{
			"base_3",
			0,
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			want := New(checkFn)
			want.sleepBeforeWait = tt.sleepBeforeWait
			got := New(checkFn)
			got.SetSleepBeforeWait(tt.sleepBeforeWait)

			diff := cmp.Diff(got, want, cmpOpts...)
			if diff != "" {
				t.Errorf("Data does not match: %s", diff)
			}
		})
	}
}

func TestSetTempErrRetryLimit(t *testing.T) {
	checkFn := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }

	for _, tt := range []struct {
		desc              string
		tempErrRetryLimit int
	}{
		{
			"base_1",
			2,
		},
		{
			"base_3",
			0,
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			want := New(checkFn)
			want.tempErrRetryLimit = tt.tempErrRetryLimit
			got := New(checkFn)
			got.SetTempErrRetryLimit(tt.tempErrRetryLimit)

			diff := cmp.Diff(got, want, cmpOpts...)
			if diff != "" {
				t.Errorf("Data does not match: %s", diff)
			}
		})
	}
}

func TestWaitWithContext(t *testing.T) {
	for _, tt := range []struct {
		desc                           string
		checkFnNumberCallsToFinishWait int
		checkFnWaitSucceeds            bool
		checkFnNumberCallsUntilErr     int
		checkFnReturnsTempErr          bool
		handlerSleepBeforeWait         time.Duration
		handlerThrottle                time.Duration
		handlerTimeout                 time.Duration
		handlerTempErrRetryLimit       int
		contextTimeout                 time.Duration
		wantCheckFnNumberCalls         int
		wantErr                        bool
	}{
		{
			desc:                           "base",
			checkFnNumberCallsToFinishWait: 1,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 100 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 100 * time.Millisecond,
			wantCheckFnNumberCalls:         1,
			wantErr:                        false,
		},
		{
			desc:                           "bad_trottle",
			checkFnNumberCallsToFinishWait: 1,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                0 * time.Millisecond,
			handlerTimeout:                 100 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 100 * time.Millisecond,
			wantCheckFnNumberCalls:         0,
			wantErr:                        true,
		},
		{
			desc:                           "throttle",
			checkFnNumberCallsToFinishWait: 3,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 100 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 100 * time.Millisecond,
			wantCheckFnNumberCalls:         3,
			wantErr:                        false,
		},
		{
			desc:                           "throttle_timeout_1",
			checkFnNumberCallsToFinishWait: 999999,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 100 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 1000 * time.Millisecond,
			wantCheckFnNumberCalls:         4,
			wantErr:                        true,
		},
		{
			desc:                           "throttle_timeout_2",
			checkFnNumberCallsToFinishWait: 999999,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 1000 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 100 * time.Millisecond,
			wantCheckFnNumberCalls:         4,
			wantErr:                        true,
		},
		{
			desc:                           "set_sleep_before_wait_and_throttle",
			checkFnNumberCallsToFinishWait: 999999,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         60 * time.Millisecond,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 100 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 100 * time.Millisecond,
			wantCheckFnNumberCalls:         2,
			wantErr:                        true,
		},
		{
			desc:                           "set_sleep_before_wait_timeout_1",
			checkFnNumberCallsToFinishWait: 2,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         200 * time.Millisecond,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 100 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 1000 * time.Millisecond,
			wantCheckFnNumberCalls:         1,
			wantErr:                        true,
		},
		{
			desc:                           "set_sleep_before_wait_timeout_2",
			checkFnNumberCallsToFinishWait: 2,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     999999,
			handlerSleepBeforeWait:         200 * time.Millisecond,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 1000 * time.Millisecond,
			handlerTempErrRetryLimit:       0,
			contextTimeout:                 100 * time.Millisecond,
			wantCheckFnNumberCalls:         1,
			wantErr:                        true,
		},
		{
			desc:                           "retry_limit_temp_err_1",
			checkFnNumberCallsToFinishWait: 999999,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     0,
			checkFnReturnsTempErr:          false,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 1000 * time.Millisecond,
			handlerTempErrRetryLimit:       5,
			contextTimeout:                 1000 * time.Millisecond,
			wantCheckFnNumberCalls:         1,
			wantErr:                        true,
		},
		{
			desc:                           "retry_limit_temp_err_2",
			checkFnNumberCallsToFinishWait: 999999,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     1,
			checkFnReturnsTempErr:          true,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 1000 * time.Millisecond,
			handlerTempErrRetryLimit:       5,
			contextTimeout:                 1000 * time.Millisecond,
			wantCheckFnNumberCalls:         5,
			wantErr:                        true,
		},
		{
			desc:                           "retry_limit_temp_err_3",
			checkFnNumberCallsToFinishWait: 3,
			checkFnWaitSucceeds:            true,
			checkFnNumberCallsUntilErr:     1,
			checkFnReturnsTempErr:          true,
			handlerSleepBeforeWait:         0,
			handlerThrottle:                30 * time.Millisecond,
			handlerTimeout:                 1000 * time.Millisecond,
			handlerTempErrRetryLimit:       5,
			contextTimeout:                 1000 * time.Millisecond,
			wantCheckFnNumberCalls:         3,
			wantErr:                        false,
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			type respType struct{}

			numberCheckFnCalls := 0
			checkFn := func() (waitFinished bool, response *respType, err error) {
				numberCheckFnCalls++
				if numberCheckFnCalls == tt.checkFnNumberCallsToFinishWait {
					if tt.checkFnWaitSucceeds {
						return true, &respType{}, nil
					}
					return true, &respType{}, fmt.Errorf("the async action couldn't be done")
				}

				if numberCheckFnCalls < tt.checkFnNumberCallsUntilErr {
					return false, nil, nil
				}

				if tt.checkFnReturnsTempErr {
					return false, nil, &oapiError.GenericOpenAPIError{
						StatusCode:   RetryHttpErrorStatusCodes[0],
						ErrorMessage: "something bad happened when checking if the async action was finished",
					}
				}
				return false, nil, fmt.Errorf("something bad happened when checking if the async action was finished")
			}
			handler := AsyncActionHandler[respType]{
				checkFn:           checkFn,
				sleepBeforeWait:   tt.handlerSleepBeforeWait,
				throttle:          tt.handlerThrottle,
				timeout:           tt.handlerTimeout,
				tempErrRetryLimit: tt.handlerTempErrRetryLimit,
			}
			ctx, cancel := context.WithTimeout(context.Background(), tt.contextTimeout)
			defer cancel()

			resp, err := handler.WaitWithContext(ctx)

			if tt.wantErr && (err == nil) {
				t.Errorf("expected error but got none")
			}
			if !tt.wantErr && (err != nil) {
				t.Errorf("expected no error but got \"%v\"", err)
			}
			if (err == nil) && (resp == nil) {
				t.Errorf("got nil err but nil resp")
			}
			if numberCheckFnCalls != tt.wantCheckFnNumberCalls {
				t.Errorf("expected %d calls to checkFn but got %d instead", tt.wantCheckFnNumberCalls, numberCheckFnCalls)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	for _, tt := range []struct {
		desc              string
		reqErr            error
		tempErrRetryLimit int
		wantErr           bool
	}{
		{
			desc: "handle_oapi_error",
			reqErr: &oapiError.GenericOpenAPIError{
				StatusCode: http.StatusInternalServerError,
			},
			tempErrRetryLimit: 5,
			wantErr:           true,
		},
		{
			desc:              "not_generic_oapi_error",
			reqErr:            fmt.Errorf("some error"),
			tempErrRetryLimit: 5,
			wantErr:           true,
		},
		{
			desc: "bad_gateway_error",
			reqErr: &oapiError.GenericOpenAPIError{
				StatusCode: http.StatusBadGateway,
			},
			tempErrRetryLimit: 5,
			wantErr:           false,
		},
		{
			desc: "gateway_timeout_error",
			reqErr: &oapiError.GenericOpenAPIError{
				StatusCode: http.StatusBadGateway,
			},
			tempErrRetryLimit: 5,
			wantErr:           false,
		},
		{
			desc: "temp_error_retry_limit_reached",
			reqErr: &oapiError.GenericOpenAPIError{
				StatusCode: http.StatusBadGateway,
			},
			tempErrRetryLimit: 1,
			wantErr:           true,
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				tempErrRetryLimit: tt.tempErrRetryLimit,
			}
			_, err := w.handleError(0, tt.reqErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
