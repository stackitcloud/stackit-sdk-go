package rabbitmq

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
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
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*Instance, error)
}

// Interface needed for tests
type APIClientCredentialsInterface interface {
	GetCredentialsExecute(ctx context.Context, projectId, instanceId, credentialsId string) (*CredentialsResponse, error)
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
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s.InstanceId == nil || s.LastOperation == nil || s.LastOperation.Type == nil || s.LastOperation.State == nil {
			return s, false, fmt.Errorf("create failed for instance with id %s. The response is not valid: the instance id, the last operation type or the state are missing", instanceId)
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeCreate && *s.LastOperation.State == InstanceStateSuccess {
			return s, true, nil
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeCreate && *s.LastOperation.State == InstanceStateFailed {
			return s, true, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return s, false, nil
	})
}

// UpdateInstanceWaitHandler will wait for update
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s.InstanceId == nil || s.LastOperation == nil || s.LastOperation.Type == nil || s.LastOperation.State == nil {
			return s, false, fmt.Errorf("update failed for instance with id %s. The response is not valid: the instance id, the last operation type or the state are missing", instanceId)
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeUpdate && *s.LastOperation.State == InstanceStateSuccess {
			return s, true, nil
		}
		if *s.InstanceId == instanceId && *s.LastOperation.Type == InstanceTypeUpdate && *s.LastOperation.State == InstanceStateFailed {
			return s, true, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return s, false, nil
	})
}

// DeleteInstanceWaitHandler will wait for delete
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err == nil {
			if s.LastOperation == nil || s.LastOperation.Type == nil || s.LastOperation.State == nil || s.LastOperation.Description == nil {
				return s, false, fmt.Errorf("delete failed for instance with id %s. The response is not valid: The last operation type, description or the state are missing", instanceId)
			}
			if *s.LastOperation.Type != InstanceTypeDelete {
				return nil, false, nil
			}
			if *s.LastOperation.State == InstanceStateSuccess {
				if strings.Contains(*s.LastOperation.Description, "DeleteFailed") || strings.Contains(*s.LastOperation.Description, "failed") {
					return s, true, fmt.Errorf("instance was deleted successfully but has errors: %s", *s.LastOperation.Description)
				}
				return s, true, nil
			}
			return s, false, nil
		}
		oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
		}
		if oapiErr.statusCode != http.StatusGone {
			return handleError(err)
		}
		return nil, true, nil
	})
}

// CreateCredentialsWaitHandler will wait for creation
func CreateCredentialsWaitHandler(ctx context.Context, a APIClientCredentialsInterface, projectId, instanceId, credentialsId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetCredentialsExecute(ctx, projectId, instanceId, credentialsId)
		if err != nil {
			oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
			}
			// Some APIs may return temporary errors and the request should be retried
			if utils.Contains(wait.RetryHttpErrorStatusCodes, oapiErr.statusCode) {
				return nil, false, nil
			}
			// If the request returns 404, the credentials have not been created yet
			if oapiErr.statusCode == http.StatusNotFound {
				return nil, false, nil
			}
			return nil, true, err
		}
		if *s.Id == credentialsId {
			return s, true, nil
		}
		return s, false, nil
	})
}

// DeleteCredentialsWaitHandler will wait for deletion
func DeleteCredentialsWaitHandler(ctx context.Context, a APIClientCredentialsInterface, projectId, instanceId, credentialsId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetCredentialsExecute(ctx, projectId, instanceId, credentialsId)
		if err != nil {
			oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
			}
			if oapiErr.statusCode != http.StatusNotFound && oapiErr.statusCode != http.StatusGone {
				return handleError(err)
			}
			return nil, true, nil
		}
		return s, false, nil
	})
}
