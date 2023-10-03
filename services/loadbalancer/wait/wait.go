package wait

import (
	"context"
	"fmt"
	"net/http"

	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
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

// CreateInstanceWaitHandler will wait for creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceName string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetLoadBalancerExecute(ctx, projectId, instanceName)
		if err != nil {
			return nil, false, err
		}
		if s == nil || s.Name == nil || *s.Name != instanceName || s.Status == nil {
			return s, false, nil
		}
		switch *s.Status {
		case InstanceStatusReady:
			return s, true, nil
		case InstanceStatusUnspecified:
			return nil, false, nil
		case InstanceStatusPending:
			return nil, false, nil
		case InstanceStatusTerminating:
			return nil, true, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, InstanceStatusTerminating)
		case InstanceStatusError:
			return nil, true, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, InstanceStatusError)
		default:
			return nil, true, fmt.Errorf("instance with name %s has unexpected status %s", instanceName, *s.Status)
		}
	})
}

// DeleteInstanceWaitHandler will wait for delete
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetLoadBalancerExecute(ctx, projectId, instanceId)
		if err == nil {
			return s, false, nil
		}
		oapiErr, ok := err.(*oapiError.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return nil, false, fmt.Errorf("could not convert error to oapiError.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return nil, false, err
		}
		return nil, true, nil
	})
}

// EnableLoadBalancingWaitHandler will wait for functionality to be enabled
func EnableLoadBalancingWaitHandler(ctx context.Context, a APIClientInterface, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetStatusExecute(ctx, projectId)
		if err != nil {
			return nil, false, err
		}
		if s == nil || s.Status == nil {
			return s, false, nil
		}
		switch *s.Status {
		case FunctionalityStatusReady:
			return s, true, nil
		case FunctionalityStatusUnspecified:
			return s, false, nil
		case FunctionalityStatusDisabled:
			return nil, false, nil
		case FunctionalityStatusUpdating:
			return nil, false, nil
		case FunctionalityStatusDeleting:
			return nil, true, fmt.Errorf("enabling load balancing failed for project %s, got status %s", projectId, FunctionalityStatusDeleting)
		case FunctionalityStatusFailed:
			return nil, true, fmt.Errorf("enabling load balancing failed for project %s, got status %s", projectId, FunctionalityStatusFailed)
		default:
			return nil, true, fmt.Errorf("load balancing for project %s has unexpected status %s", projectId, *s.Status)
		}
	})
}
