package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	alb "github.com/stackitcloud/stackit-sdk-go/services/alb/v2api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_UNSPECIFIED = alb.LOADBALANCERSTATUS_STATUS_UNSPECIFIED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_PENDING = alb.LOADBALANCERSTATUS_STATUS_PENDING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_READY = alb.LOADBALANCERSTATUS_STATUS_READY
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_ERROR = alb.LOADBALANCERSTATUS_STATUS_ERROR
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_TERMINATING = alb.LOADBALANCERSTATUS_STATUS_TERMINATING
)

func CreateOrUpdateLoadbalancerWaitHandler(ctx context.Context, client alb.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	waitConfig := wait.WaiterHelper[alb.LoadBalancer, alb.LoadBalancerStatus]{
		FetchInstance: client.GetLoadBalancer(ctx, projectId, region, name).Execute,
		GetState: func(response *alb.LoadBalancer) (alb.LoadBalancerStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Status, nil
		},
		ActiveState: []alb.LoadBalancerStatus{alb.LOADBALANCERSTATUS_STATUS_READY},
		ErrorState:  []alb.LoadBalancerStatus{alb.LOADBALANCERSTATUS_STATUS_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteLoadbalancerWaitHandler(ctx context.Context, client alb.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	waitConfig := wait.WaiterHelper[alb.LoadBalancer, alb.LoadBalancerStatus]{
		FetchInstance: client.GetLoadBalancer(ctx, projectId, region, name).Execute,
		GetState: func(response *alb.LoadBalancer) (alb.LoadBalancerStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Status, nil
		},
		ActiveState: []alb.LoadBalancerStatus{},
		ErrorState:  []alb.LoadBalancerStatus{alb.LOADBALANCERSTATUS_STATUS_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
