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
	LOADBALANCERSTATUS_UNSPECIFIED = "STATUS_UNSPECIFIED"
	LOADBALANCERSTATUS_PENDING     = "STATUS_PENDING"
	LOADBALANCERSTATUS_READY       = "STATUS_READY"
	LOADBALANCERSTATUS_ERROR       = "STATUS_ERROR"
	LOADBALANCERSTATUS_TERMINATING = "STATUS_TERMINATING"
)

// CreateLoadBalancerWaitHandler will wait for load balancer creation
func CreateLoadBalancerWaitHandler(ctx context.Context, a loadbalancer.DefaultAPI, projectId, region, instanceName string) *wait.AsyncActionHandler[loadbalancer.LoadBalancer] {
	waitConfig := wait.WaiterHelper[loadbalancer.LoadBalancer, string]{
		FetchInstance: a.GetLoadBalancer(ctx, projectId, region, instanceName).Execute,
		GetState: func(r *loadbalancer.LoadBalancer) (string, error) {
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
		ActiveState: []string{LOADBALANCERSTATUS_READY},
		ErrorState:  []string{LOADBALANCERSTATUS_TERMINATING, LOADBALANCERSTATUS_ERROR},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteLoadBalancerWaitHandler will wait for load balancer deletion
func DeleteLoadBalancerWaitHandler(ctx context.Context, a loadbalancer.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[struct{}] {
	waitConfig := wait.WaiterHelper[struct{}, string]{
		FetchInstance: func() (*struct{}, error) {
			_, err := a.GetLoadBalancer(ctx, projectId, region, instanceId).Execute()
			return &struct{}{}, err
		},
		GetState: func(_ *struct{}) (string, error) {
			return "", nil
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
