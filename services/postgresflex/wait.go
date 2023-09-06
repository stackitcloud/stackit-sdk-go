package postgresflex

import (
	"context"
	"fmt"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

const (
	InstanceStateEmpty       = ""
	InstanceStateProgressing = "Progressing"
	InstanceStateSuccess     = "Ready"
	InstanceStateFailed      = "Failure"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*InstanceResponse, error)
	GetUsersExecute(ctx context.Context, projectId string, instanceId string) (*UsersResponse, error)
}

// Interface needed for tests
type APIClientUserInterface interface {
	GetUserExecute(ctx context.Context, projectId, instanceId, userId string) (*UserResponse, error)
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
	instanceCreated := false
	var instanceGetResponse *InstanceResponse

	return wait.New(func() (res interface{}, done bool, err error) {
		if !instanceCreated {
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
			case InstanceStateProgressing:
				return nil, false, nil
			case InstanceStateSuccess:
				instanceCreated = true
				instanceGetResponse = s
			case InstanceStateFailed:
				return nil, true, fmt.Errorf("create failed for instance with id %s", instanceId)
			}
		}

		// User operations aren't available right after an instance is deemed successful
		// To check if they are, perform a users request
		_, err = a.GetUsersExecute(ctx, projectId, instanceId)
		if err == nil {
			return instanceGetResponse, true, nil
		}
		oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return handleError(err)
		}
		if oapiErr.statusCode < 500 {
			return nil, true, fmt.Errorf("users request after instance creation returned %d status code", oapiErr.statusCode)
		}
		return nil, false, nil
	})
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
		case InstanceStateProgressing:
			return s, false, nil
		case InstanceStateSuccess:
			return s, true, nil
		case InstanceStateFailed:
			return s, true, fmt.Errorf("create failed for instance with id %s", instanceId)
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
			return handleError(err)
		}
		if oapiErr.statusCode != 404 {
			return handleError(err)
		}
		return nil, true, nil
	})
}

// DeleteUserWaitHandler will wait for delete
func DeleteUserWaitHandler(ctx context.Context, a APIClientUserInterface, projectId, instanceId, userId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		u, err := a.GetUserExecute(ctx, projectId, instanceId, userId)
		if err == nil {
			return u, false, nil
		}
		oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return handleError(err)
		}
		if oapiErr.statusCode != 404 {
			return handleError(err)
		}
		return nil, true, nil
	})
}
