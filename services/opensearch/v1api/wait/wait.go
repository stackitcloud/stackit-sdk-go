package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	opensearch "github.com/stackitcloud/stackit-sdk-go/services/opensearch/v1api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_ACTIVE = opensearch.INSTANCESTATUS_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_FAILED = opensearch.INSTANCESTATUS_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_STOPPED = opensearch.INSTANCESTATUS_STOPPED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_CREATING = opensearch.INSTANCESTATUS_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_DELETING = opensearch.INSTANCESTATUS_DELETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_UPDATING = opensearch.INSTANCESTATUS_UPDATING

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCELASTOPERATIONTYPE_CREATE = opensearch.INSTANCELASTOPERATIONTYPE_CREATE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCELASTOPERATIONTYPE_UPDATE = opensearch.INSTANCELASTOPERATIONTYPE_UPDATE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCELASTOPERATIONTYPE_DELETE = opensearch.INSTANCELASTOPERATIONTYPE_DELETE
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a opensearch.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[opensearch.Instance] {
	return createOrUpdateInstanceWaitHandler(ctx, a, projectId, instanceId)
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a opensearch.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[opensearch.Instance] {
	return createOrUpdateInstanceWaitHandler(ctx, a, projectId, instanceId)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a opensearch.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		s, err := a.GetInstance(ctx, projectId, instanceId).Execute()
		if err == nil {
			if s.Status == nil {
				return false, nil, fmt.Errorf("delete failed for instance with id %s. The response is not valid: The status is missing", instanceId)
			}
			if *s.Status != opensearch.INSTANCESTATUS_DELETING {
				return false, nil, nil
			}
			if *s.Status == opensearch.INSTANCESTATUS_ACTIVE {
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
func CreateCredentialsWaitHandler(ctx context.Context, a opensearch.DefaultAPI, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[opensearch.CredentialsResponse] {
	handler := wait.New(func() (waitFinished bool, response *opensearch.CredentialsResponse, err error) {
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
func DeleteCredentialsWaitHandler(ctx context.Context, a opensearch.DefaultAPI, projectId, instanceId, credentialsId string) *wait.AsyncActionHandler[struct{}] {
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

func createOrUpdateInstanceWaitHandler(ctx context.Context, a opensearch.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[opensearch.Instance] {
	waitConfig := wait.WaiterHelper[opensearch.Instance, opensearch.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, instanceId).Execute,
		GetState: func(s *opensearch.Instance) (opensearch.InstanceStatus, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.Status == nil {
				return "", errors.New("status is missing")
			}
			return *s.Status, nil
		},
		ActiveState: []opensearch.InstanceStatus{opensearch.INSTANCESTATUS_ACTIVE},
		ErrorState:  []opensearch.InstanceStatus{opensearch.INSTANCESTATUS_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}
