package wait

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

// Load balancer instance status
const (
	InstanceStatusUnspecified = "STATUS_UNSPECIFIED"
	InstanceStatusPending     = "STATUS_PENDING"
	InstanceStatusReady       = "STATUS_READY"
	InstanceStatusError       = "STATUS_ERROR"
	InstanceStatusTerminating = "STATUS_TERMINATING"
)

// Load balancing functionality status
const (
	FunctionalityStatusUnspecified = "STATUS_UNSPECIFIED"
	FunctionalityStatusReady       = "STATUS_READY"
	FunctionalityStatusFailed      = "STATUS_FAILED"
	FunctionalityStatusUpdating    = "STATUS_UPDATING"
	FunctionalityStatusDeleting    = "STATUS_DELETING"
	FunctionalityStatusDisabled    = "STATUS_DISABLED"
)

// Interface needed for tests
type APIClientInterface interface {
	GetLoadBalancerExecute(ctx context.Context, projectId, name string) (*loadbalancer.LoadBalancer, error)
	GetStatusExecute(ctx context.Context, projectId string) (*loadbalancer.StatusResponse, error)
}

// CreateLoadBalancerWaitHandler will wait for load balancer creation
func CreateLoadBalancerWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceName string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	return wait.New(func() (waitFinished bool, response *loadbalancer.LoadBalancer, err error) {
		s, err := a.GetLoadBalancerExecute(ctx, projectId, instanceName)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Name == nil || *s.Name != instanceName || s.Status == nil {
			return false, s, nil
		}
		switch *s.Status {
		case InstanceStatusReady:
			return true, s, nil
		case InstanceStatusUnspecified:
			return false, nil, nil
		case InstanceStatusPending:
			return false, nil, nil
		case InstanceStatusTerminating:
			return true, nil, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, InstanceStatusTerminating)
		case InstanceStatusError:
			return true, nil, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, InstanceStatusError)
		default:
			return true, nil, fmt.Errorf("instance with name %s has unexpected status %s", instanceName, *s.Status)
		}
	})
}

// DeleteLoadBalancerWaitHandler will wait for load balancer deletion
func DeleteLoadBalancerWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	return wait.New(func() (waitFinished bool, response *loadbalancer.LoadBalancer, err error) {
		s, err := a.GetLoadBalancerExecute(ctx, projectId, instanceId)
		if err == nil {
			return false, s, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, nil, err
		}
		return true, nil, nil
	})
}

// EnableLoadBalancingWaitHandler will wait for functionality to be enabled
func EnableLoadBalancingWaitHandler(ctx context.Context, a APIClientInterface, projectId string) *wait.AsyncActionHandler[loadbalancer.StatusResponse] {
	return wait.New(func() (waitFinished bool, response *loadbalancer.StatusResponse, err error) {
		s, err := a.GetStatusExecute(ctx, projectId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Status == nil {
			return false, s, nil
		}
		switch *s.Status {
		case FunctionalityStatusReady:
			return true, s, nil
		case FunctionalityStatusUnspecified:
			return false, s, nil
		case FunctionalityStatusDisabled:
			return false, nil, nil
		case FunctionalityStatusUpdating:
			return false, nil, nil
		case FunctionalityStatusDeleting:
			return true, nil, fmt.Errorf("enabling load balancing failed for project %s, got status %s", projectId, FunctionalityStatusDeleting)
		case FunctionalityStatusFailed:
			return true, nil, fmt.Errorf("enabling load balancing failed for project %s, got status %s", projectId, FunctionalityStatusFailed)
		default:
			return true, nil, fmt.Errorf("load balancing for project %s has unexpected status %s", projectId, *s.Status)
		}
	})
}
