package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex"
)

const (
	InstanceStateEmpty      = ""
	InstanceStateProcessing = "Progressing"
	InstanceStateUnknown    = "Unknown"
	InstanceStateSuccess    = "Ready"
	InstanceStateFailed     = "Failed"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId, region string) (*sqlserverflex.GetInstanceResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *sqlserverflex.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId, region)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, s, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateEmpty:
			return false, s, nil
		case InstanceStateProcessing:
			return false, s, nil
		case InstanceStateUnknown:
			return false, s, nil
		case InstanceStateSuccess:
			return true, s, nil
		case InstanceStateFailed:
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	handler.SetSleepBeforeWait(5 * time.Second)
	return handler
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *sqlserverflex.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId, region)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, s, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateEmpty:
			return false, s, nil
		case InstanceStateProcessing:
			return false, s, nil
		case InstanceStateUnknown:
			return false, s, nil
		case InstanceStateSuccess:
			return true, s, nil
		case InstanceStateFailed:
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
	})
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	return UpdateInstanceWaitHandler(ctx, a, projectId, instanceId, region)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId, region string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetInstanceExecute(ctx, projectId, instanceId, region)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}
