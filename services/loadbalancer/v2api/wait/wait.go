package wait

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	loadbalancer "github.com/stackitcloud/stackit-sdk-go/services/loadbalancer/v2api"
)

// Load balancer instance status
const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_UNSPECIFIED = loadbalancer.LOADBALANCERSTATUS_STATUS_UNSPECIFIED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_PENDING = loadbalancer.LOADBALANCERSTATUS_STATUS_PENDING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_READY = loadbalancer.LOADBALANCERSTATUS_STATUS_READY
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_ERROR = loadbalancer.LOADBALANCERSTATUS_STATUS_ERROR
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	LOADBALANCERSTATUS_TERMINATING = loadbalancer.LOADBALANCERSTATUS_STATUS_TERMINATING
)

// CreateLoadBalancerWaitHandler will wait for load balancer creation
func CreateLoadBalancerWaitHandler(ctx context.Context, a loadbalancer.DefaultAPI, projectId, region, instanceName string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	waitConfig := wait.WaiterHelper[loadbalancer.LoadBalancer, loadbalancer.LoadBalancerStatus]{
		FetchInstance: a.GetLoadBalancer(ctx, projectId, region, instanceName).Execute,
		GetState: func(r *loadbalancer.LoadBalancer) (loadbalancer.LoadBalancerStatus, error) {
			if r == nil || r.Status == nil {
				return "", errors.New("response or status is nil")
			}
			var sb strings.Builder
			if r.Errors != nil {
				for _, err := range r.Errors {
					sb.WriteString(fmt.Sprintf("%s: %s; ", *err.Type, *err.Description))
				}
				return "", fmt.Errorf("create failed for instance with name %s, got status %s and errors: %s", instanceName, *r.Status, sb.String())
			}
			return *r.Status, nil
		},
		ActiveState: []loadbalancer.LoadBalancerStatus{loadbalancer.LOADBALANCERSTATUS_STATUS_READY},
		ErrorState:  []loadbalancer.LoadBalancerStatus{loadbalancer.LOADBALANCERSTATUS_STATUS_TERMINATING, loadbalancer.LOADBALANCERSTATUS_STATUS_ERROR},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteLoadBalancerWaitHandler will wait for load balancer deletion
func DeleteLoadBalancerWaitHandler(ctx context.Context, a loadbalancer.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	waitConfig := wait.WaiterHelper[loadbalancer.LoadBalancer, loadbalancer.LoadBalancerStatus]{
		FetchInstance: a.GetLoadBalancer(ctx, projectId, region, instanceId).Execute,
		GetState: func(l *loadbalancer.LoadBalancer) (loadbalancer.LoadBalancerStatus, error) {
			if l == nil || l.Status == nil {
				return "", errors.New("response or status is nil")
			}
			return *l.Status, nil
		},
		ErrorState: []loadbalancer.LoadBalancerStatus{loadbalancer.LOADBALANCERSTATUS_STATUS_ERROR},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
