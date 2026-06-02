package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	serviceenablement "github.com/stackitcloud/stackit-sdk-go/services/serviceenablement/v2api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	SERVICESTATUSSTATE_ENABLED = serviceenablement.SERVICESTATUSSTATE_ENABLED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	SERVICESTATUSSTATE_ENABLING = serviceenablement.SERVICESTATUSSTATE_ENABLING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	SERVICESTATUSSTATE_DISABLED = serviceenablement.SERVICESTATUSSTATE_DISABLED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	SERVICESTATUSSTATE_DISABLING = serviceenablement.SERVICESTATUSSTATE_DISABLING
)

// EnableServiceWaitHandler will wait for service enablement
func EnableServiceWaitHandler(ctx context.Context, a serviceenablement.DefaultAPI, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	waitConfig := wait.WaiterHelper[serviceenablement.ServiceStatus, serviceenablement.ServiceStatusState]{
		FetchInstance: a.GetServiceStatusRegional(ctx, region, projectId, serviceId).Execute,
		GetState: func(s *serviceenablement.ServiceStatus) (serviceenablement.ServiceStatusState, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.State == nil {
				return "", errors.New("state is missing")
			}
			return *s.State, nil
		},
		ActiveState: []serviceenablement.ServiceStatusState{serviceenablement.SERVICESTATUSSTATE_ENABLED},
		ErrorState:  []serviceenablement.ServiceStatusState{serviceenablement.SERVICESTATUSSTATE_DISABLED, serviceenablement.SERVICESTATUSSTATE_DISABLING},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}

// DisableServiceWaitHandler will wait for service disablement
func DisableServiceWaitHandler(ctx context.Context, a serviceenablement.DefaultAPI, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	waitConfig := wait.WaiterHelper[serviceenablement.ServiceStatus, serviceenablement.ServiceStatusState]{
		FetchInstance: a.GetServiceStatusRegional(ctx, region, projectId, serviceId).Execute,
		GetState: func(s *serviceenablement.ServiceStatus) (serviceenablement.ServiceStatusState, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.State == nil {
				return "", errors.New("status is missing")
			}
			return *s.State, nil
		},
		ActiveState: []serviceenablement.ServiceStatusState{serviceenablement.SERVICESTATUSSTATE_DISABLED},
		ErrorState:  []serviceenablement.ServiceStatusState{serviceenablement.SERVICESTATUSSTATE_ENABLED, serviceenablement.SERVICESTATUSSTATE_ENABLING},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}
