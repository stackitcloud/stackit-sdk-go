package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
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
	handler := wait.New(func() (bool, *alb.LoadBalancer, error) {
		response, err := client.GetLoadBalancer(ctx, projectId, region, name).Execute()
		if err != nil {
			return false, nil, err
		}
		if response.HasStatus() {
			switch *response.Status {
			case LOADBALANCERSTATUS_PENDING:
				return false, nil, nil
			case LOADBALANCERSTATUS_UNSPECIFIED:
				return false, nil, nil
			case LOADBALANCERSTATUS_ERROR:
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

func DeleteLoadbalancerWaitHandler(ctx context.Context, client alb.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[alb.LoadBalancer] {
	handler := wait.New(func() (bool, *alb.LoadBalancer, error) {
		_, err := client.GetLoadBalancer(ctx, projectId, region, name).Execute()
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
