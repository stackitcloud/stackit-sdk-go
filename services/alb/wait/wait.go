package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/alb"
)

const (
	StatusUnspecified = "STATUS_UNSPECIFIED"
	StatusPending     = "STATUS_PENDING"
	StatusReady       = "STATUS_READY"
	StatusError       = "STATUS_ERROR"
	StatusTerminating = "STATUS_TERMINATING"
)

type APIClientLoadbalancerInterface interface {
	GetLoadBalancerExecute(ctx context.Context, projectId string, region string, name string) (*alb.LoadBalancer, error)
}

func CreateOrUpdateLoadbalancerWaitHandler(ctx context.Context, client APIClientLoadbalancerInterface, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	handler := wait.New(func() (bool, *alb.LoadBalancer, error) {
		response, err := client.GetLoadBalancerExecute(ctx, projectId, region, name)
		if err != nil {
			return false, nil, err
		}
		if response.HasStatus() {
			switch *response.Status {
			case StatusPending:
				return false, nil, nil
			case StatusUnspecified:
				return false, nil, nil
			case StatusError:
				return true, response, fmt.Errorf("loadbalancer in error: %s", *response.Status)
			default:
				return true, response, nil
			}
		}

		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteLoadbalancerWaitHandler(ctx context.Context, client APIClientLoadbalancerInterface, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	handler := wait.New(func() (bool, *alb.LoadBalancer, error) {
		_, err := client.GetLoadBalancerExecute(ctx, projectId, region, name)
		if err != nil {
			var apiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &apiErr) {
				if statusCode := apiErr.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, nil, nil
				}
			}
			return true, nil, err
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
