package wait

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	GetServiceStatusExecute(ctx context.Context, projectId string) (*loadbalancer.GetServiceStatusResponse, error)
}

// CreateLoadBalancerWaitHandler will wait for load balancer creation
func CreateLoadBalancerWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceName string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	handler := wait.New(func() (waitFinished bool, response *loadbalancer.LoadBalancer, err error) {
		s, err := a.GetLoadBalancerExecute(ctx, projectId, instanceName)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Name == nil || *s.Name != instanceName || s.Status == nil {
			return false, nil, nil
		}

		var errors []string
		if s.Errors != nil && len(*s.Errors) > 0 {
			for _, err := range *s.Errors {
				errors = append(errors, fmt.Sprintf("%s: %s", *err.Type, *err.Description))
			}
			return true, s, fmt.Errorf("create failed for instance with name %s, got status %s and errors: %s", instanceName, *s.Status, strings.Join(errors, ";"))
		}

		switch *s.Status {
		case InstanceStatusReady:
			return true, s, nil
		case InstanceStatusUnspecified:
			return false, nil, nil
		case InstanceStatusPending:
			return false, nil, nil
		case InstanceStatusTerminating:
			return true, s, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, InstanceStatusTerminating)
		case InstanceStatusError:
			return true, s, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, InstanceStatusError)
		default:
			return true, s, fmt.Errorf("instance with name %s has unexpected status %s", instanceName, *s.Status)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteLoadBalancerWaitHandler will wait for load balancer deletion
func DeleteLoadBalancerWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetLoadBalancerExecute(ctx, projectId, instanceId)
		if err == nil {
			return false, nil, nil
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
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// EnableServiceWaitHandler will wait for functionality to be enabled
func EnableServiceWaitHandler(ctx context.Context, a APIClientInterface, projectId string) *wait.AsyncActionHandler[loadbalancer.GetServiceStatusResponse] {
	handler := wait.New(func() (waitFinished bool, response *loadbalancer.GetServiceStatusResponse, err error) {
		s, err := a.GetServiceStatusExecute(ctx, projectId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Status == nil {
			return false, nil, nil
		}
		switch *s.Status {
		case FunctionalityStatusReady:
			return true, s, nil
		case FunctionalityStatusUnspecified:
			return false, nil, nil
		case FunctionalityStatusDisabled:
			return false, nil, nil
		case FunctionalityStatusUpdating:
			return false, nil, nil
		case FunctionalityStatusDeleting:
			return true, s, fmt.Errorf("enabling load balancing failed for project %s, got status %s", projectId, FunctionalityStatusDeleting)
		case FunctionalityStatusFailed:
			return true, s, fmt.Errorf("enabling load balancing failed for project %s, got status %s", projectId, FunctionalityStatusFailed)
		default:
			return true, s, fmt.Errorf("load balancing for project %s has unexpected status %s", projectId, *s.Status)
		}
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}
