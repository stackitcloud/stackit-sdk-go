package wait

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/logme"
)

const (
	InstanceStateSuccess = "succeeded"
	InstanceStateFailed  = "failed"
	InstanceTypeCreate   = "create"
	InstanceTypeUpdate   = "update"
	InstanceTypeDelete   = "delete"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*logme.Instance, error)
}

// Interface needed for tests
type APIClientCredentialsInterface interface {
	GetCredentialsExecute(ctx context.Context, projectId, instanceId, credentialsId string) (*logme.CredentialsResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[logme.Instance] {
	return wait.New(func() (waitFinished bool, response *logme.Instance, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.InstanceId == nil || s.LastOperation == nil || s.LastOperation.Type == nil || s.LastOperation.State == nil {
			return false, s, fmt.Errorf("create failed for instance with id %s. The response is not valid: the instance id, the last operation type or the state are missing", instanceId)
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeCreate && *s.LastOperation.State == InstanceStateSuccess {
			return true, s, nil
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeCreate && *s.LastOperation.State == InstanceStateFailed {
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return false, s, nil
	})
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[logme.Instance] {
	return wait.New(func() (waitFinished bool, response *logme.Instance, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.InstanceId == nil || s.LastOperation == nil || s.LastOperation.Type == nil || s.LastOperation.State == nil {
			return false, s, fmt.Errorf("update failed for instance with id %s. The response is not valid: the instance id, the last operation type or the state are missing", instanceId)
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeUpdate && *s.LastOperation.State == InstanceStateSuccess {
			return true, s, nil
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeUpdate && *s.LastOperation.State == InstanceStateFailed {
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return false, s, nil
	})
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[logme.Instance] {
	return wait.New(func() (waitFinished bool, response *logme.Instance, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err == nil {
			if s.LastOperation == nil || s.LastOperation.Type == nil || s.LastOperation.State == nil || s.LastOperation.Description == nil {
				return false, s, fmt.Errorf("delete failed for instance with id %s. The response is not valid: The last operation type, description or the state are missing", instanceId)
			}
			if *s.LastOperation.Type != InstanceTypeDelete {
				return false, nil, nil
			}
			if *s.LastOperation.State == InstanceStateSuccess {
				if strings.Contains(*s.LastOperation.Description, "DeleteFailed") || strings.Contains(*s.LastOperation.Description, "failed") {
					return true, s, fmt.Errorf("instance was deleted successfully but has errors: %s", *s.LastOperation.Description)
				}
				return true, s, nil
			}
			return false, s, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusGone {
			return false, nil, err
		}
		return true, nil, nil
	})
}

// CreateCredentialsWaitHandler will wait for credentials creation
func CreateCredentialsWaitHandler(ctx context.Context, a APIClientCredentialsInterface, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[logme.CredentialsResponse] {
	return wait.New(func() (waitFinished bool, response *logme.CredentialsResponse, err error) {
		s, err := a.GetCredentialsExecute(ctx, projectId, instanceId, credentialsId)
		if err != nil {
			oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
			}
			// If the request returns 404, the credentials have not been created yet
			if oapiErr.StatusCode == http.StatusNotFound {
				return false, nil, nil
			}
			return false, nil, err
		}
		if *s.Id == credentialsId {
			return true, s, nil
		}
		return false, s, nil
	})
}

// DeleteCredentialsWaitHandler will wait for credentials deletion
func DeleteCredentialsWaitHandler(ctx context.Context, a APIClientCredentialsInterface, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[logme.CredentialsResponse] {
	return wait.New(func() (waitFinished bool, response *logme.CredentialsResponse, err error) {
		s, err := a.GetCredentialsExecute(ctx, projectId, instanceId, credentialsId)
		if err != nil {
			oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
			}
			if oapiErr.StatusCode != http.StatusNotFound && oapiErr.StatusCode != http.StatusGone {
				return false, nil, err
			}
			return true, nil, nil
		}
		return false, s, nil
	})
}
