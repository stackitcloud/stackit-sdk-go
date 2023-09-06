package resourcemanager

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

const (
	ActiveState   = LifecycleState("ACTIVE")
	CreatingState = LifecycleState("CREATING")
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetProjectExecute(ctx context.Context, containerId string) (*ProjectResponseWithParents, error)
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

// CreateProjectWaitHandler will wait for creation
// returned interface is nil or *ProjectResponseWithParents
func CreateProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			return handleError(err)
		}
		if p.ContainerId == nil || p.LifecycleState == nil {
			return p, false, fmt.Errorf("creation failed: response invalid for container id %s. Container id or LifeCycleState missing", containerId)
		}
		if *p.ContainerId == containerId && *p.LifecycleState == ActiveState {
			return p, true, nil
		}
		if *p.ContainerId == containerId && *p.LifecycleState == CreatingState {
			return p, false, nil
		}
		return p, false, fmt.Errorf("creation failed: received project state '%s'", *p.LifecycleState)
	})
}

// DeleteProjectWaitHandler will wait for delete
// returned interface is nil or *ProjectResponseWithParents
func DeleteProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
			}
			// Some APIs may return temporary errors and the request should be retried
			if utils.Contains(wait.RetryHttpErrorStatusCodes, oapiErr.statusCode) {
				return nil, false, nil
			}

			if oapiErr.statusCode == http.StatusNotFound || oapiErr.statusCode == http.StatusForbidden {
				return nil, true, nil
			}

			return nil, false, err
		}
		return p, false, nil
	})
}
