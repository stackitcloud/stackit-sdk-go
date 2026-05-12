package wait

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

// Load balancer instance status
const (
	// Deprecated: InstanceStatusUnspecified is deprecated and will be removed after 14th November 2025. Use [loadbalancer.LOADBALANCERSTATUS_UNSPECIFIED] instead.
	InstanceStatusUnspecified = "STATUS_UNSPECIFIED"
	// Deprecated: InstanceStatusPending is deprecated and will be removed after 14th November 2025. Use [loadbalancer.LOADBALANCERSTATUS_PENDING] instead.
	InstanceStatusPending = "STATUS_PENDING"
	// Deprecated: InstanceStatusReady is deprecated and will be removed after 14th November 2025. Use [loadbalancer.LOADBALANCERSTATUS_READY] instead.
	InstanceStatusReady = "STATUS_READY"
	// Deprecated: InstanceStatusError is deprecated and will be removed after 14th November 2025. Use [loadbalancer.LOADBALANCERSTATUS_ERROR] instead.
	InstanceStatusError = "STATUS_ERROR"
	// Deprecated: InstanceStatusTerminating is deprecated and will be removed after 12th November 2025. Use [loadbalancer.LOADBALANCERSTATUS_TERMINATING] instead.
	InstanceStatusTerminating = "STATUS_TERMINATING"
)

var _ APIClientInterface = &loadbalancer.APIClient{}

// Interface needed for tests
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
type APIClientInterface interface {
	GetLoadBalancerExecute(ctx context.Context, projectId, region, name string) (*loadbalancer.LoadBalancer, error)
}

// CreateLoadBalancerWaitHandler will wait for load balancer creation
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateLoadBalancerWaitHandler(ctx context.Context, a APIClientInterface, projectId, region, instanceName string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	waitConfig := wait.WaiterHelper[loadbalancer.LoadBalancer, loadbalancer.LoadBalancerStatus]{
		FetchInstance: func() (*loadbalancer.LoadBalancer, error) {
			return a.GetLoadBalancerExecute(ctx, projectId, region, instanceName)
		},
		GetState: func(r *loadbalancer.LoadBalancer) (loadbalancer.LoadBalancerStatus, error) {
			if r == nil || r.Status == nil {
				return "", errors.New("response or status is nil")
			}
			var sb strings.Builder
			if r.Errors != nil && len(*r.Errors) > 0 {
				for _, err := range *r.Errors {
					sb.WriteString(fmt.Sprintf("%s: %s; ", *err.Type, *err.Description))
				}
				return "", fmt.Errorf("create failed for instance with name %s, got status %s and errors: %s", instanceName, *r.Status, sb.String())
			}
			return *r.Status, nil
		},
		ActiveState: []loadbalancer.LoadBalancerStatus{loadbalancer.LOADBALANCERSTATUS_READY},
		ErrorState:  []loadbalancer.LoadBalancerStatus{loadbalancer.LOADBALANCERSTATUS_ERROR, loadbalancer.LOADBALANCERSTATUS_TERMINATING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteLoadBalancerWaitHandler will wait for load balancer deletion
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func DeleteLoadBalancerWaitHandler(ctx context.Context, a APIClientInterface, projectId, region, instanceId string) *wait.AsyncActionHandler[struct{}] {
	waitConfig := wait.WaiterHelper[struct{}, loadbalancer.LoadBalancerStatus]{
		FetchInstance: func() (*struct{}, error) {
			_, err := a.GetLoadBalancerExecute(ctx, projectId, region, instanceId)
			return &struct{}{}, err
		},
		GetState: func(r *struct{}) (loadbalancer.LoadBalancerStatus, error) {
			return "", nil
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
