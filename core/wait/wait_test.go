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

	f := &AsyncActionHandler[interface{}]{
		throttle: 1 * time.Minute,
	}

	type fields struct {
		throttle time.Duration
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AsyncActionHandler[interface{}]
	}{
		{"ok", fields{throttle: 30 * time.Second}, args{d: 1 * time.Minute}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				throttle: tt.fields.throttle,
			}
			if got := w.SetThrottle(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(AsyncActionHandler[interface{}]{})) {
				t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
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

	f := &AsyncActionHandler[interface{}]{
		throttle: 1 * time.Minute,
	}

	type fields struct {
		throttle time.Duration
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AsyncActionHandler[interface{}]
	}{
		{"ok", fields{throttle: 30 * time.Second}, args{d: 1 * time.Minute}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				throttle: tt.fields.throttle,
			}
			if got := w.SetThrottle(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(AsyncActionHandler[interface{}]{})) {
				t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
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

	f := &AsyncActionHandler[interface{}]{
		throttle: 1 * time.Minute,
	}

	type fields struct {
		throttle time.Duration
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AsyncActionHandler[interface{}]
	}{
		{"ok", fields{throttle: 30 * time.Second}, args{d: 1 * time.Minute}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				throttle: tt.fields.throttle,
			}
			if got := w.SetThrottle(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(AsyncActionHandler[interface{}]{})) {
				t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetRetryLimitTempErr(t *testing.T) {
	checkFn := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }

	for _, tt := range []struct {
		desc              string
		retryLimitTempErr int
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
			want.tempErrRetryLimit = tt.retryLimitTempErr
			got := New(checkFn)
			got.SetRetryLimitTempErr(tt.retryLimitTempErr)

			diff := cmp.Diff(got, want, cmpOpts...)
			if diff != "" {
				t.Errorf("Data does not match: %s", diff)
			}
		})
	}

	f := &AsyncActionHandler[interface{}]{
		throttle: 1 * time.Minute,
	}

	type fields struct {
		throttle time.Duration
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AsyncActionHandler[interface{}]
	}{
		{"ok", fields{throttle: 30 * time.Second}, args{d: 1 * time.Minute}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				throttle: tt.fields.throttle,
			}
			if got := w.SetThrottle(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(AsyncActionHandler[interface{}]{})) {
				t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	type fields struct {
		checkFn           AsyncActionCheck[interface{}]
		throttle          time.Duration
		timeout           time.Duration
		tempErrRetryLimit int
	}
	tests := []struct {
		name     string
		fields   fields
		wantDone bool
		wantErr  bool
	}{
		{"ok", fields{throttle: 1 * time.Second, timeout: 1 * time.Hour, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			return true, nil, nil
		}}, true, false},

		{"ok 2", fields{throttle: 200 * time.Millisecond, timeout: 1 * time.Hour, tempErrRetryLimit: 5, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			if ctx.Err() == nil {
				return false, nil, nil
			}
			return true, nil, nil
		}}, true, false},

		{"err", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Hour, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			return true, nil, fmt.Errorf("something happened")
		}}, true, true},

		{"err 2", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Hour, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			return false, nil, fmt.Errorf("something happened")
		}}, true, true},

		{"timeout", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Millisecond, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			return false, nil, nil
		}}, false, true},

		{"tempErrorLimitReached", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Millisecond, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			return false, nil, nil
		}}, false, true},

		{"badThrottle", fields{throttle: 0 * time.Second, timeout: 1 * time.Hour, checkFn: func() (waitFinished bool, res *interface{}, err error) {
			return true, nil, nil
		}}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				checkFn:           tt.fields.checkFn,
				throttle:          tt.fields.throttle,
				timeout:           tt.fields.timeout,
				tempErrRetryLimit: tt.fields.tempErrRetryLimit,
			}
			_, err := w.WaitWithContext(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Wait.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	tests := []struct {
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
	}
	for _, tt := range tests {
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
