package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	serviceenablement "github.com/stackitcloud/stackit-sdk-go/services/serviceenablement/v2api"
)

const (
	SERVICESTATUSSTATE_ENABLED   = "ENABLED"
	SERVICESTATUSSTATE_ENABLING  = "ENABLING"
	SERVICESTATUSSTATE_DISABLED  = "DISABLED"
	SERVICESTATUSSTATE_DISABLING = "DISABLING"
)

func EnableServiceWaitHandler(ctx context.Context, a serviceenablement.DefaultAPI, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	handler := wait.New(func() (waitFinished bool, response *serviceenablement.ServiceStatus, err error) {
		s, err := a.GetServiceStatusRegional(ctx, region, projectId, serviceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.State == nil {
			return false, nil, nil
		}
		switch *s.State {
		default:
			return true, s, fmt.Errorf("service with id %s has unexpected state %s", serviceId, *s.State)
		case SERVICESTATUSSTATE_ENABLED:
			return true, s, nil
		case SERVICESTATUSSTATE_ENABLING:
			return false, nil, nil
		case SERVICESTATUSSTATE_DISABLED:
			return true, s, fmt.Errorf("enabling failed for service with id %s", serviceId)
		case SERVICESTATUSSTATE_DISABLING:
			return true, s, fmt.Errorf("service with id %s is in state %s", serviceId, *s.State)
		}
	})

	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}

func DisableServiceWaitHandler(ctx context.Context, a serviceenablement.DefaultAPI, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	handler := wait.New(func() (waitFinished bool, response *serviceenablement.ServiceStatus, err error) {
		s, err := a.GetServiceStatusRegional(ctx, region, projectId, serviceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.State == nil {
			return false, nil, nil
		}
		switch *s.State {
		default:
			return true, s, fmt.Errorf("service with id %s has unexpected state %s", serviceId, *s.State)
		case SERVICESTATUSSTATE_DISABLED:
			return true, s, nil
		case SERVICESTATUSSTATE_DISABLING:
			return false, nil, nil
		case SERVICESTATUSSTATE_ENABLED:
			return true, s, fmt.Errorf("disabling failed for service with id %s", serviceId)
		case SERVICESTATUSSTATE_ENABLING:
			return true, s, fmt.Errorf("service with id %s is in state %s", serviceId, *s.State)
		}
	})
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}
