package wait

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/alb"
)

var (
	testProject = uuid.NewString()
	testRegion  = "eu01"
	testName    = "testlb"
)

var _ APIClientLoadbalancerInterface = (*apiClientLoadbalancerMocked)(nil)

type response struct {
	loadbalancer *alb.LoadBalancer
	err          error
}

type apiClientLoadbalancerMocked struct {
	n         int
	responses []response
}

// GetLoadBalancerExecute implements APIClientLoadbalancerInterface.
func (a *apiClientLoadbalancerMocked) GetLoadBalancerExecute(_ context.Context, _, _, _ string) (*alb.LoadBalancer, error) {
	resp := a.responses[a.n]
	a.n++
	a.n %= len(a.responses)
	return resp.loadbalancer, resp.err
}

func TestCreateOrUpdateLoadbalancerWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []response
		want      *alb.LoadBalancer
		wantErr   bool
	}{
		{
			"create succeeded immediately",
			[]response{
				{&alb.LoadBalancer{Status: alb.PtrString(StatusReady)}, nil},
			},
			&alb.LoadBalancer{Status: utils.Ptr(StatusReady)},
			false,
		},
		{
			"create succeeded delayed",
			[]response{
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: alb.PtrString(StatusReady)}, nil},
			},
			&alb.LoadBalancer{Status: utils.Ptr(StatusReady)},
			false,
		},
		{
			"create failed delayed",
			[]response{
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: alb.PtrString(StatusError)}, nil},
			},
			&alb.LoadBalancer{Status: utils.Ptr(StatusError)},
			true,
		},
		{
			"timeout",
			[]response{
				{&alb.LoadBalancer{Status: alb.PtrString(StatusPending)}, nil},
			},
			nil,
			true,
		},
		{
			"broken state",
			[]response{
				{&alb.LoadBalancer{Status: alb.PtrString("bogus")}, nil},
			},
			&alb.LoadBalancer{Status: alb.PtrString("bogus")},
			false,
		},
		// no special update tests needed as the states are the same
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiClientLoadbalancerMocked{
				responses: tt.responses,
			}
			handler := CreateOrUpdateLoadbalancerWaitHandler(ctx, client, testProject, testRegion, testName)
			got, err := handler.SetTimeout(1 * time.Second).
				SetThrottle(250 * time.Millisecond).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %qe ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing loadbalancer %s", diff)
			}
		})
	}
}

func TestDeleteLoadbalancerWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []response
		wantErr   bool
	}{
		{
			"Delete with '404' succeeded immediately",
			[]response{
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with '404' delayed",
			[]response{
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with 'gone' succeeded immediately",
			[]response{
				{nil, oapierror.NewError(http.StatusGone, "gone")},
			},
			false,
		},
		{
			"Delete with 'gone' delayed",
			[]response{
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{nil, oapierror.NewError(http.StatusGone, "not found")},
			},
			false,
		},
		{
			"Delete with error delayed",
			[]response{
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(string(StatusError))}, oapierror.NewError(http.StatusInternalServerError, "kapow")},
			},
			true,
		},
		{
			"Cannot delete",
			[]response{
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(StatusPending)}, nil},
				{&alb.LoadBalancer{Status: utils.Ptr(string(StatusError))}, oapierror.NewError(http.StatusOK, "ok")},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiClientLoadbalancerMocked{
				responses: tt.responses,
			}
			handler := DeleteLoadbalancerWaitHandler(ctx, client, testProject, testRegion, testName)
			_, err := handler.SetTimeout(1 * time.Second).
				SetThrottle(250 * time.Millisecond).
				WaitWithContext(ctx)

			if tt.wantErr != (err != nil) {
				t.Fatalf("wrong error result. want err: %v got %v", tt.wantErr, err)
			}
			if tt.wantErr {
				var apiErr *oapierror.GenericOpenAPIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("expected openapi error, got %v", err)
				}
			}
		})
	}
}
