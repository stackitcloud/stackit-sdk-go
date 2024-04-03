package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/postgresflex"
)

const (
	InstanceStateEmpty       = ""
	InstanceStateProgressing = "Progressing"
	InstanceStateSuccess     = "Ready"
	InstanceStateFailed      = "Failure"
	InstanceStateDeleted     = "Deleted"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*postgresflex.InstanceResponse, error)
	ListUsersExecute(ctx context.Context, projectId string, instanceId string) (*postgresflex.ListUsersResponse, error)
}

// Interface needed for tests
type APIClientUserInterface interface {
	GetUserExecute(ctx context.Context, projectId, instanceId, userId string) (*postgresflex.GetUserResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[postgresflex.InstanceResponse] {
	instanceCreated := false
	var instanceGetResponse *postgresflex.InstanceResponse

	handler := wait.New(func() (waitFinished bool, response *postgresflex.InstanceResponse, err error) {
		if !instanceCreated {
			s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
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
				return false, nil, nil
			case InstanceStateProgressing:
				return false, nil, nil
			case InstanceStateSuccess:
				instanceCreated = true
				instanceGetResponse = s
			case InstanceStateFailed:
				return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
			}
		}

		// User operations aren't available right after an instance is deemed successful
		// To check if they are, perform a users request
		_, err = a.ListUsersExecute(ctx, projectId, instanceId)
		if err == nil {
			return true, instanceGetResponse, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode < 500 {
			return true, instanceGetResponse, fmt.Errorf("users request after instance creation returned %d status code", oapiErr.StatusCode)
		}
		return false, nil, nil
	})
	// Sleep before wait is set because sometimes API returns 404 right after creation request
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[postgresflex.InstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *postgresflex.InstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
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
			return false, nil, nil
		case InstanceStateProgressing:
			return false, nil, nil
		case InstanceStateSuccess:
			return true, s, nil
		case InstanceStateFailed:
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, nil, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateDeleted:
			return true, nil, nil
		}
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// ForceDeleteInstanceWaitHandler will wait for instance deletion
func ForceDeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetInstanceExecute(ctx, projectId, instanceId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode != 404 {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteUserWaitHandler will wait for delete
func DeleteUserWaitHandler(ctx context.Context, a APIClientUserInterface, projectId, instanceId, userId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetUserExecute(ctx, projectId, instanceId, userId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode != 404 {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}
