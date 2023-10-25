package wait

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

func TestNew(t *testing.T) {
	simple := func() (waitFinished bool, res *interface{}, err error) { return true, nil, nil }
	type args struct {
		f AsyncActionCheck[interface{}]
	}
	tests := []struct {
		name string
		args args
		want *AsyncActionHandler[interface{}]
	}{
		{"ok", args{simple}, &AsyncActionHandler[interface{}]{check: simple, throttle: 5 * time.Second, tempErrRetryLimit: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.f); !cmp.Equal(got.throttle, tt.want.throttle) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetThrottle(t *testing.T) {
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
	f := &AsyncActionHandler[interface{}]{
		sleepBeforeWait: 1 * time.Minute,
	}

	type fields struct {
		sleepBeforeWait time.Duration
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
		{"ok", fields{sleepBeforeWait: 30 * time.Second}, args{d: 1 * time.Minute}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				sleepBeforeWait: tt.fields.sleepBeforeWait,
			}
			if got := w.SetSleepBeforeWait(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(AsyncActionHandler[interface{}]{})) {
				t.Errorf("Wait.SetSleepBeforeWait() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	type fields struct {
		check             AsyncActionCheck[interface{}]
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
		{"ok", fields{throttle: 1 * time.Second, timeout: 1 * time.Hour, check: func() (waitFinished bool, res *interface{}, err error) {
			return true, nil, nil
		}}, true, false},

		{"ok 2", fields{throttle: 200 * time.Millisecond, timeout: 1 * time.Hour, tempErrRetryLimit: 5, check: func() (waitFinished bool, res *interface{}, err error) {
			if ctx.Err() == nil {
				return false, nil, nil
			}
			return true, nil, nil
		}}, true, false},

		{"err", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Hour, check: func() (waitFinished bool, res *interface{}, err error) {
			return true, nil, fmt.Errorf("something happened")
		}}, true, true},

		{"err 2", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Hour, check: func() (waitFinished bool, res *interface{}, err error) {
			return false, nil, fmt.Errorf("something happened")
		}}, true, true},

		{"timeout", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Millisecond, check: func() (waitFinished bool, res *interface{}, err error) {
			return false, nil, nil
		}}, false, true},

		{"tempErrorLimitReached", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Millisecond, check: func() (waitFinished bool, res *interface{}, err error) {
			return false, nil, nil
		}}, false, true},

		{"badThrottle", fields{throttle: 0 * time.Second, timeout: 1 * time.Hour, check: func() (waitFinished bool, res *interface{}, err error) {
			return true, nil, nil
		}}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				check:             tt.fields.check,
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

func TestSetTimeout(t *testing.T) {
	f := &AsyncActionHandler[interface{}]{
		throttle: 5 * time.Second,
		timeout:  5 * time.Hour,
	}

	type fields struct {
		throttle time.Duration
		timeout  time.Duration
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
		{"ok", fields{timeout: 1 * time.Hour, throttle: 5 * time.Second}, args{d: 5 * time.Hour}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &AsyncActionHandler[interface{}]{
				throttle: tt.fields.throttle,
				timeout:  tt.fields.timeout,
			}
			if got := w.SetTimeout(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(AsyncActionHandler[interface{}]{})) {
				t.Errorf("Wait.SetTimeout() = %v, want %v", got, tt.want)
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
