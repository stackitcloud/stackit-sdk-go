package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceenablement"
)

const (
	// Deprecated: ServiceStateEnabled is deprecated and will be removed after 14th November 2025. Use [serviceenablement.SERVICESTATUSSTATE_ENABLED] instead.
	ServiceStateEnabled = "ENABLED"
	// Deprecated: ServiceStateEnabling is deprecated and will be removed after 14th November 2025. Use [serviceenablement.SERVICESTATUSSTATE_ENABLING] instead.
	ServiceStateEnabling = "ENABLING"
	// Deprecated: ServiceStateDisabled is deprecated and will be removed after 14th November 2025. Use [serviceenablement.SERVICESTATUSSTATE_DISABLED] instead.
	ServiceStateDisabled = "DISABLED"
	// Deprecated: ServiceStateDisabling is deprecated and will be removed after 14th November 2025. Use [serviceenablement.SERVICESTATUSSTATE_DISABLING] instead.
	ServiceStateDisabling = "DISABLING"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetServiceStatusRegionalExecute(ctx context.Context, region, projectId, serviceId string) (*serviceenablement.ServiceStatus, error)
}

func EnableServiceWaitHandler(ctx context.Context, a APIClientInstanceInterface, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	handler := wait.New(func() (waitFinished bool, response *serviceenablement.ServiceStatus, err error) {
		s, err := a.GetServiceStatusRegionalExecute(ctx, region, projectId, serviceId)
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

func DisableServiceWaitHandler(ctx context.Context, a APIClientInstanceInterface, region, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	handler := wait.New(func() (waitFinished bool, response *serviceenablement.ServiceStatus, err error) {
		s, err := a.GetServiceStatusRegionalExecute(ctx, region, projectId, serviceId)
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
