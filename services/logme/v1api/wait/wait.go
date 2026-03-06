package wait

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	logme "github.com/stackitcloud/stackit-sdk-go/services/logme/v1api"
)

const (
	INSTANCESTATUS_ACTIVE   = "active"
	INSTANCESTATUS_FAILED   = "failed"
	INSTANCESTATUS_STOPPED  = "stopped"
	INSTANCESTATUS_CREATING = "creating"
	INSTANCESTATUS_DELETING = "deleting"
	INSTANCESTATUS_UPDATING = "updating"
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a logme.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[logme.Instance] {
	handler := wait.New(func() (waitFinished bool, response *logme.Instance, err error) {
		s, err := a.GetInstance(ctx, projectId, instanceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Status == nil {
			return false, nil, fmt.Errorf("create failed for instance with id %s. The response is not valid: the status is missing", instanceId)
		}
		switch *s.Status {
		case INSTANCESTATUS_ACTIVE:
			return true, s, nil
		case INSTANCESTATUS_FAILED:
			return true, s, fmt.Errorf("create failed for instance with id %s: %s", instanceId, s.LastOperation.Description)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a logme.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[logme.Instance] {
	handler := wait.New(func() (waitFinished bool, response *logme.Instance, err error) {
		s, err := a.GetInstance(ctx, projectId, instanceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Status == nil {
			return false, nil, fmt.Errorf("update failed for instance with id %s. The response is not valid: the status is missing", instanceId)
		}
		switch *s.Status {
		case INSTANCESTATUS_ACTIVE:
			return true, s, nil
		case INSTANCESTATUS_FAILED:
			return true, s, fmt.Errorf("update failed for instance with id %s: %s", instanceId, s.LastOperation.Description)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a logme.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		s, err := a.GetInstance(ctx, projectId, instanceId).Execute()
		if err == nil {
			if s.Status == nil {
				return false, nil, fmt.Errorf("delete failed for instance with id %s. The response is not valid: The status is missing", instanceId)
			}
			if *s.Status != INSTANCESTATUS_DELETING {
				return false, nil, nil
			}
			if *s.Status == INSTANCESTATUS_ACTIVE {
				if strings.Contains(s.LastOperation.Description, "DeleteFailed") || strings.Contains(s.LastOperation.Description, "failed") {
					return true, nil, fmt.Errorf("instance was deleted successfully but has errors: %s", s.LastOperation.Description)
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
func CreateCredentialsWaitHandler(ctx context.Context, a logme.DefaultAPI, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[logme.CredentialsResponse] {
	handler := wait.New(func() (waitFinished bool, response *logme.CredentialsResponse, err error) {
		s, err := a.GetCredentials(ctx, projectId, instanceId, credentialsId).Execute()
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
		if s.Id == credentialsId {
			return true, s, nil
		}
		return false, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteCredentialsWaitHandler will wait for credentials deletion
func DeleteCredentialsWaitHandler(ctx context.Context, a logme.DefaultAPI, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetCredentials(ctx, projectId, instanceId, credentialsId).Execute()
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
