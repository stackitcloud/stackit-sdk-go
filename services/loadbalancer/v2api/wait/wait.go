package wait

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	loadbalancer "github.com/stackitcloud/stackit-sdk-go/services/loadbalancer/v2api"
)

// Load balancer instance status
const (
	LOADBALANCERSTATUS_UNSPECIFIED = "STATUS_UNSPECIFIED"
	LOADBALANCERSTATUS_PENDING     = "STATUS_PENDING"
	LOADBALANCERSTATUS_READY       = "STATUS_READY"
	LOADBALANCERSTATUS_ERROR       = "STATUS_ERROR"
	LOADBALANCERSTATUS_TERMINATING = "STATUS_TERMINATING"
)

// CreateLoadBalancerWaitHandler will wait for load balancer creation
func CreateLoadBalancerWaitHandler(ctx context.Context, a loadbalancer.DefaultAPI, projectId, region, instanceName string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	handler := wait.New(func() (waitFinished bool, response *loadbalancer.LoadBalancer, err error) {
		s, err := a.GetLoadBalancer(ctx, projectId, region, instanceName).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Name == nil || *s.Name != instanceName || s.Status == nil {
			return false, nil, nil
		}

		var errors []string
		if len(s.Errors) > 0 {
			for _, err := range s.Errors {
				errors = append(errors, fmt.Sprintf("%s: %s", *err.Type, *err.Description))
			}
			return true, s, fmt.Errorf("create failed for instance with name %s, got status %s and errors: %s", instanceName, *s.Status, strings.Join(errors, ";"))
		}

		switch *s.Status {
		case LOADBALANCERSTATUS_READY:
			return true, s, nil
		case LOADBALANCERSTATUS_UNSPECIFIED:
			return false, nil, nil
		case LOADBALANCERSTATUS_PENDING:
			return false, nil, nil
		case LOADBALANCERSTATUS_TERMINATING:
			return true, s, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, LOADBALANCERSTATUS_TERMINATING)
		case LOADBALANCERSTATUS_ERROR:
			return true, s, fmt.Errorf("create failed for instance with name %s, got status %s", instanceName, LOADBALANCERSTATUS_ERROR)
		default:
			return true, s, fmt.Errorf("instance with name %s has unexpected status %s", instanceName, *s.Status)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteLoadBalancerWaitHandler will wait for load balancer deletion
func DeleteLoadBalancerWaitHandler(ctx context.Context, a loadbalancer.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetLoadBalancer(ctx, projectId, region, instanceId).Execute()
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
