package wait

import (
	"context"
	"fmt"
	"net/http"

	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
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

// CreateProjectWaitHandler will wait for creation
// returned interface is nil or *resourcemanager.ProjectResponseWithParents
func CreateProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			return nil, false, err
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
// returned interface is nil or *resourcemanager.ProjectResponseWithParents
func DeleteProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			oapiErr, ok := err.(*oapiError.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return nil, false, fmt.Errorf("could not convert error to oapiError.GenericOpenAPIError")
			}
			if oapiErr.StatusCode == http.StatusNotFound || oapiErr.StatusCode == http.StatusForbidden {
				return nil, true, nil
			}
			return nil, false, err
		}
		return p, false, nil
	})
}
