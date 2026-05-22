package wait

import (
	"context"
	"errors"
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

// EnableServiceWaitHandler will wait for service enablement
func EnableServiceWaitHandler(ctx context.Context, a serviceenablement.DefaultAPI, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	waitConfig := wait.WaiterHelper[serviceenablement.ServiceStatus, string]{
		FetchInstance: a.GetServiceStatusRegional(ctx, region, projectId, serviceId).Execute,
		GetState: func(s *serviceenablement.ServiceStatus) (string, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.State == nil {
				return "", errors.New("state is missing")
			}
			return *s.State, nil
		},
		ActiveState: []string{SERVICESTATUSSTATE_ENABLED},
		ErrorState:  []string{SERVICESTATUSSTATE_DISABLED, SERVICESTATUSSTATE_DISABLING},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}

// DisableServiceWaitHandler will wait for service disablement
func DisableServiceWaitHandler(ctx context.Context, a serviceenablement.DefaultAPI, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	waitConfig := wait.WaiterHelper[serviceenablement.ServiceStatus, string]{
		FetchInstance: a.GetServiceStatusRegional(ctx, region, projectId, serviceId).Execute,
		GetState: func(s *serviceenablement.ServiceStatus) (string, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.State == nil {
				return "", errors.New("status is missing")
			}
			return *s.State, nil
		},
		ActiveState: []string{SERVICESTATUSSTATE_DISABLED},
		ErrorState:  []string{SERVICESTATUSSTATE_ENABLED, SERVICESTATUSSTATE_ENABLING},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}
