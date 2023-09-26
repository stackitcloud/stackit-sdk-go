package mongodbflex

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

const (
	InstanceStateEmpty      = ""
	InstanceStateProcessing = "PROCESSING"
	InstanceStateUnknown    = "UNKNOWN"
	InstanceStateSuccess    = "READY"
	InstanceStateFailed     = "FAILED"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*GetInstanceResponse, error)
}

func handleError(reqErr error) (res interface{}, done bool, err error) {
	oapiErr, ok := reqErr.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
	if !ok {
		return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
	}
	// Some APIs may return temporary errors and the request should be retried
	if utils.Contains(wait.RetryHttpErrorStatusCodes, oapiErr.statusCode) {
		return nil, false, nil
	}
	return nil, false, reqErr
}

// CreateInstanceWaitHandler will wait for creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.Handler {
	waitHandler := wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return s, false, nil
		}
		switch *s.Item.Status {
		default:
			return nil, true, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateEmpty:
			return nil, false, nil
		case InstanceStateProcessing:
			return nil, false, nil
		case InstanceStateUnknown:
			return nil, false, nil
		case InstanceStateSuccess:
			return s, true, nil
		case InstanceStateFailed:
			return nil, true, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
	})
	return waitHandler.SetSleepBeforeWait(5 * time.Second)
}

// UpdateInstanceWaitHandler will wait for update
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return s, false, nil
		}
		switch *s.Item.Status {
		default:
			return s, true, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateEmpty:
			return s, false, nil
		case InstanceStateProcessing:
			return s, false, nil
		case InstanceStateUnknown:
			return s, false, nil
		case InstanceStateSuccess:
			return s, true, nil
		case InstanceStateFailed:
			return s, true, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
	})
}

// DeleteInstanceWaitHandler will wait for delete
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err == nil {
			return s, false, nil
		}
		oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
		}
		if oapiErr.statusCode != http.StatusNotFound {
			return handleError(err)
		}
		return nil, true, nil
	})
}
