package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceenablement"
)

const (
	ServiceStateEnabled   = "ENABLED"
	ServiceStateEnabling  = "ENABLING"
	ServiceStateDisabled  = "DISABLED"
	ServiceStateDisabling = "DISABLING"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetServiceStatusExecute(ctx context.Context, projectId, serviceId string) (*serviceenablement.ServiceStatus, error)
}

func EnableServiceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	handler := wait.New(func() (waitFinished bool, response *serviceenablement.ServiceStatus, err error) {
		s, err := a.GetServiceStatusExecute(ctx, projectId, serviceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.State == nil {
			return false, nil, nil
		}
		switch *s.State {
		default:
			return true, s, fmt.Errorf("service with id %s has unexpected state %s", serviceId, *s.State)
		case ServiceStateEnabled:
			return true, s, nil
		case ServiceStateEnabling:
			return false, nil, nil
		case ServiceStateDisabled:
			return true, s, fmt.Errorf("enabling failed for service with id %s", serviceId)
		case ServiceStateDisabling:
			return true, s, fmt.Errorf("service with id %s is in state %s", serviceId, *s.State)
		}
	})

	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(45 * time.Second)
	return handler
}

func DisableServiceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, serviceId string) *wait.AsyncActionHandler[serviceenablement.ServiceStatus] {
	handler := wait.New(func() (waitFinished bool, response *serviceenablement.ServiceStatus, err error) {
		s, err := a.GetServiceStatusExecute(ctx, projectId, serviceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.State == nil {
			return false, nil, nil
		}
		switch *s.State {
		default:
			return true, s, fmt.Errorf("service with id %s has unexpected state %s", serviceId, *s.State)
		case ServiceStateDisabled:
			return true, s, nil
		case ServiceStateDisabling:
			return false, nil, nil
		case ServiceStateEnabled:
			return true, s, fmt.Errorf("disabling failed for service with id %s", serviceId)
		case ServiceStateEnabling:
			return true, s, fmt.Errorf("service with id %s is in state %s", serviceId, *s.State)
		}
	})
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(30 * time.Second)
	return handler
}
