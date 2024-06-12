package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	GetProjectExecute(ctx context.Context, containerId string) (*resourcemanager.GetProjectResponse, error)
}

// CreateProjectWaitHandler will wait for project creation
func CreateProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.AsyncActionHandler[resourcemanager.GetProjectResponse] {
	handler := wait.New(func() (waitFinished bool, response *resourcemanager.GetProjectResponse, err error) {
		p, err := a.GetProjectExecute(ctx, containerId)
		if err != nil {
			return false, nil, err
		}
		if p.ContainerId == nil || p.LifecycleState == nil {
			return false, nil, fmt.Errorf("creation failed: response invalid for container id %s. Container id or LifeCycleState missing", containerId)
		}
		if *p.ContainerId == containerId && *p.LifecycleState == ActiveState {
			return true, p, nil
		}
		if *p.ContainerId == containerId && *p.LifecycleState == CreatingState {
			return false, nil, nil
		}
		return true, p, fmt.Errorf("creation failed: received project state '%s'", *p.LifecycleState)
	})
	handler.SetSleepBeforeWait(1 * time.Minute)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteProjectWaitHandler will wait for project deletion
func DeleteProjectWaitHandler(ctx context.Context, a APIClientInterface, containerId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetProjectExecute(ctx, containerId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode == http.StatusNotFound || oapiErr.StatusCode == http.StatusForbidden {
			return true, nil, nil
		}
		return false, nil, err
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}
