package wait

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	simple := func() (res interface{}, tempErrFound, done bool, err error) { return nil, false, true, nil }
	type args struct {
		f WaitFn
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{"ok", args{simple}, &Handler{fn: simple, throttle: 5 * time.Second, retryLimitTempErr: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.f); !cmp.Equal(got.throttle, tt.want.throttle) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWait_SetThrottle(t *testing.T) {
	simple := func() (res interface{}, tempErrFound, done bool, err error) { return nil, false, true, nil }
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"ok", args{10 * time.Second}, nil},
		{"err", args{0 * time.Second}, fmt.Errorf("throttle can't be 0")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(simple)
			got := w.SetThrottle(tt.args.d)
			if got == nil && tt.want != nil {
				t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
			}
			if got != nil && tt.want != nil {
				if got.Error() != tt.want.Error() {
					t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestWait_SetSleepBeforeWait(t *testing.T) {
	f := &Handler{
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
		want   *Handler
	}{
		{"ok", fields{sleepBeforeWait: 30 * time.Second}, args{d: 1 * time.Minute}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Handler{
				sleepBeforeWait: tt.fields.sleepBeforeWait,
			}
			if got := w.SetSleepBeforeWait(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(Handler{})) {
				t.Errorf("Wait.SetSleepBeforeWait() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitWithCtx_Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	type fields struct {
		fn                WaitFn
		throttle          time.Duration
		timeout           time.Duration
		retryLimitTempErr int
	}
	tests := []struct {
		name     string
		fields   fields
		wantDone bool
		wantErr  bool
	}{
		{"ok", fields{throttle: 1 * time.Second, timeout: 1 * time.Hour, fn: func() (res interface{}, tempErrFound, done bool, err error) {
			return nil, false, true, nil
		}}, true, false},

		{"ok 2", fields{throttle: 200 * time.Millisecond, timeout: 1 * time.Hour, retryLimitTempErr: 5, fn: func() (res interface{}, tempErrFound, done bool, err error) {
			if ctx.Err() == nil {
				return nil, false, false, nil
			}
			return nil, false, true, nil
		}}, true, false},

		{"err", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Hour, fn: func() (res interface{}, tempErrFound, done bool, err error) {
			return nil, false, true, fmt.Errorf("something happened")
		}}, true, true},

		{"err 2", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Hour, fn: func() (res interface{}, tempErrFound, done bool, err error) {
			return nil, false, false, fmt.Errorf("something happened")
		}}, true, true},

		{"timeout", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Millisecond, fn: func() (res interface{}, tempErrFound, done bool, err error) {
			return nil, false, false, nil
		}}, false, true},

		{"tempErrorLimitReached", fields{throttle: 1 * time.Millisecond, timeout: 1 * time.Millisecond, fn: func() (res interface{}, tempErrFound, done bool, err error) {
			return nil, true, false, nil
		}}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Handler{
				fn:                tt.fields.fn,
				throttle:          tt.fields.throttle,
				timeout:           tt.fields.timeout,
				retryLimitTempErr: tt.fields.retryLimitTempErr,
			}
			_, err := w.WaitWithContext(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Wait.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWait_SetTimeout(t *testing.T) {
	f := &Handler{
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
		want   *Handler
	}{
		{"ok", fields{timeout: 1 * time.Hour, throttle: 5 * time.Second}, args{d: 5 * time.Hour}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Handler{
				throttle: tt.fields.throttle,
				timeout:  tt.fields.timeout,
			}
			if got := w.SetTimeout(tt.args.d); !cmp.Equal(got, tt.want, cmp.AllowUnexported(Handler{})) {
				t.Errorf("Wait.SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
