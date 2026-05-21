package wait

import (
	"context"
	"fmt"
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
		case serviceenablement.SERVICESTATUSSTATE_ENABLED:
			return true, s, nil
		case serviceenablement.SERVICESTATUSSTATE_ENABLING:
			return false, nil, nil
		case serviceenablement.SERVICESTATUSSTATE_DISABLED:
			return true, s, fmt.Errorf("enabling failed for service with id %s", serviceId)
		case serviceenablement.SERVICESTATUSSTATE_DISABLING:
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
		case serviceenablement.SERVICESTATUSSTATE_DISABLED:
			return true, s, nil
		case serviceenablement.SERVICESTATUSSTATE_DISABLING:
			return false, nil, nil
		case serviceenablement.SERVICESTATUSSTATE_ENABLED:
			return true, s, fmt.Errorf("disabling failed for service with id %s", serviceId)
		case serviceenablement.SERVICESTATUSSTATE_ENABLING:
			return true, s, fmt.Errorf("service with id %s is in state %s", serviceId, *s.State)
		}
	})
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}
