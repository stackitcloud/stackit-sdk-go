package wait

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/resourcemanager"
)

const (
	ActiveState   = resourcemanager.LifecycleState("ACTIVE")
	CreatingState = resourcemanager.LifecycleState("CREATING")
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetProjectExecute(ctx context.Context, containerId string) (*resourcemanager.ProjectResponseWithParents, error)
}

// CreateProjectWaitHandler will wait for project creation
func CreateProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.AsyncActionHandler[resourcemanager.ProjectResponseWithParents] {
	return wait.New(func() (waitFinished bool, response *resourcemanager.ProjectResponseWithParents, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			return false, nil, err
		}
		if p.ContainerId == nil || p.LifecycleState == nil {
			return false, p, fmt.Errorf("creation failed: response invalid for container id %s. Container id or LifeCycleState missing", containerId)
		}
		if *p.ContainerId == containerId && *p.LifecycleState == ActiveState {
			return true, p, nil
		}
		if *p.ContainerId == containerId && *p.LifecycleState == CreatingState {
			return false, p, nil
		}
		return false, p, fmt.Errorf("creation failed: received project state '%s'", *p.LifecycleState)
	})
}

// DeleteProjectWaitHandler will wait for project deletion
func DeleteProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.AsyncActionHandler[resourcemanager.ProjectResponseWithParents] {
	return wait.New(func() (waitFinished bool, response *resourcemanager.ProjectResponseWithParents, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
			}
			if oapiErr.StatusCode == http.StatusNotFound || oapiErr.StatusCode == http.StatusForbidden {
				return true, nil, nil
			}
			return false, nil, err
		}
		return false, p, nil
	})
}
