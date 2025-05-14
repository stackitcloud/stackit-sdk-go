package wait

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/redis"
)

const (
	// Deprecated: InstanceStatusActive is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCESTATUS_ACTIVE] instead.
	InstanceStatusActive = "active"
	// Deprecated: InstanceStatusFailed is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCESTATUS_FAILED] instead.
	InstanceStatusFailed = "failed"
	// Deprecated: InstanceStatusStopped is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCESTATUS_STOPPED] instead.
	InstanceStatusStopped = "stopped"
	// Deprecated: InstanceStatusCreating is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCESTATUS_CREATING] instead.
	InstanceStatusCreating = "creating"
	// Deprecated: InstanceStatusDeleting is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCESTATUS_DELETING] instead.
	InstanceStatusDeleting = "deleting"
	// Deprecated: InstanceStatusUpdating is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCESTATUS_UPDATING] instead.
	InstanceStatusUpdating = "updating"

	// Deprecated: InstanceOperationStateInProgress is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCELASTOPERATIONSTATE_IN_PROGRESS] instead.
	InstanceOperationStateInProgress = "in progress"
	// Deprecated: InstanceOperationStateSucceeded is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCELASTOPERATIONSTATE_SUCCEEDED] instead.
	InstanceOperationStateSucceeded = "succeeded"
	// Deprecated: InstanceOperationStateFailed is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCELASTOPERATIONSTATE_FAILED] instead.
	InstanceOperationStateFailed = "failed"
	// Deprecated: InstanceOperationTypeCreate is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCELASTOPERATIONTYPE_CREATE] instead.
	InstanceOperationTypeCreate = "create"
	// Deprecated: InstanceOperationTypeUpdate is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCELASTOPERATIONTYPE_UPDATE] instead.
	InstanceOperationTypeUpdate = "update"
	// Deprecated: InstanceOperationTypeDelete is deprecated and will be removed after 14th November 2025. Use [redis.INSTANCELASTOPERATIONTYPE_DELETE] instead.
	InstanceOperationTypeDelete = "delete"

	// Deprecated: InstanceStateSuccess is deprecated and will be removed after 2nd October 2025. Use [redis.INSTANCELASTOPERATIONSTATE_SUCCEEDED] instead.
	InstanceStateSuccess = "succeeded"
	// Deprecated: InstanceStateFailed is deprecated and will be removed after 2nd October 2025. Use [redis.INSTANCELASTOPERATIONSTATE_FAILED] instead.
	InstanceStateFailed = "failed"
	// Deprecated: InstanceTypeCreate is deprecated and will be removed after 2nd October 2025. Use [redis.INSTANCELASTOPERATIONTYPE_CREATE] instead.
	InstanceTypeCreate = "create"
	// Deprecated: InstanceTypeUpdate is deprecated and will be removed after 2nd October 2025. Use [redis.INSTANCELASTOPERATIONTYPE_UPDATE] instead.
	InstanceTypeUpdate = "update"
	// Deprecated: InstanceTypeDelete is deprecated and will be removed after 2nd October 2025. Use [redis.INSTANCELASTOPERATIONTYPE_DELETE] instead.
	InstanceTypeDelete = "delete"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*redis.Instance, error)
}

// Interface needed for tests
type APIClientCredentialsInterface interface {
	GetCredentialsExecute(ctx context.Context, projectId, instanceId, credentialsId string) (*redis.CredentialsResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[redis.Instance] {
	handler := wait.New(func() (waitFinished bool, response *redis.Instance, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.Status == nil {
			return false, nil, fmt.Errorf("create failed for instance with id %s. The response is not valid: the status is missing", instanceId)
		}
		switch *s.Status {
		case redis.INSTANCESTATUS_ACTIVE:
			return true, s, nil
		case redis.INSTANCESTATUS_FAILED:
			var failedDescription string
			if s.LastOperation != nil && s.LastOperation.Description != nil {
				failedDescription = *s.LastOperation.Description
			}
			return true, s, fmt.Errorf("create failed for instance with id %s: %s", instanceId, failedDescription)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[redis.Instance] {
	handler := wait.New(func() (waitFinished bool, response *redis.Instance, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.Status == nil {
			return false, nil, fmt.Errorf("update failed for instance with id %s. The response is not valid: the instance id or the status are missing", instanceId)
		}
		switch *s.Status {
		case redis.INSTANCESTATUS_ACTIVE:
			return true, s, nil
		case redis.INSTANCESTATUS_FAILED:
			var failedDescription string
			if s.LastOperation != nil && s.LastOperation.Description != nil {
				failedDescription = *s.LastOperation.Description
			}
			return true, s, fmt.Errorf("update failed for instance with id %s: %s", instanceId, failedDescription)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err == nil {
			if s.LastOperation == nil || s.LastOperation.Description == nil || s.Status == nil {
				return false, nil, fmt.Errorf("delete failed for instance with id %s. The response is not valid: The status or last operation description are missing", instanceId)
			}
			if *s.Status != redis.INSTANCESTATUS_DELETING {
				return false, nil, nil
			}
			if *s.Status == redis.INSTANCESTATUS_ACTIVE {
				if strings.Contains(*s.LastOperation.Description, "DeleteFailed") || strings.Contains(*s.LastOperation.Description, "failed") {
					return true, nil, fmt.Errorf("instance was deleted successfully but has errors: %s", *s.LastOperation.Description)
				}
				return true, nil, nil
			}
			return false, nil, nil
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
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateCredentialsWaitHandler will wait for credentials creation
func CreateCredentialsWaitHandler(ctx context.Context, a APIClientCredentialsInterface, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[redis.CredentialsResponse] {
	handler := wait.New(func() (waitFinished bool, response *redis.CredentialsResponse, err error) {
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
		return false, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteCredentialsWaitHandler will wait for credentials deletion
func DeleteCredentialsWaitHandler(ctx context.Context, a APIClientCredentialsInterface, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetCredentialsExecute(ctx, projectId, instanceId, credentialsId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusNotFound && oapiErr.StatusCode != http.StatusGone {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}
