package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	alb "github.com/stackitcloud/stackit-sdk-go/services/alb/v2api"
)

const (
	LOADBALANCERSTATUS_UNSPECIFIED = "STATUS_UNSPECIFIED"
	LOADBALANCERSTATUS_PENDING     = "STATUS_PENDING"
	LOADBALANCERSTATUS_READY       = "STATUS_READY"
	LOADBALANCERSTATUS_ERROR       = "STATUS_ERROR"
	LOADBALANCERSTATUS_TERMINATING = "STATUS_TERMINATING"
)

func CreateOrUpdateLoadbalancerWaitHandler(ctx context.Context, client alb.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	waitConfig := wait.WaiterHelper[alb.LoadBalancer, string]{
		FetchInstance: client.GetLoadBalancer(ctx, projectId, region, name).Execute,
		GetState: func(response *alb.LoadBalancer) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Status, nil
		},
		ActiveState: []string{LOADBALANCERSTATUS_READY},
		ErrorState:  []string{LOADBALANCERSTATUS_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteLoadbalancerWaitHandler(ctx context.Context, client alb.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	waitConfig := wait.WaiterHelper[alb.LoadBalancer, string]{
		FetchInstance: client.GetLoadBalancer(ctx, projectId, region, name).Execute,
		GetState: func(response *alb.LoadBalancer) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Status, nil
		},
		ActiveState: []string{},
		ErrorState:  []string{LOADBALANCERSTATUS_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
