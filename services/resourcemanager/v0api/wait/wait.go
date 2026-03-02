package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	resourcemanager "github.com/stackitcloud/stackit-sdk-go/services/resourcemanager/v0api"
)

// CreateProjectWaitHandler will wait for project creation
func CreateProjectWaitHandler(ctx context.Context, a resourcemanager.DefaultAPI, containerId string) *wait.AsyncActionHandler[resourcemanager.GetProjectResponse] {
	handler := wait.New(func() (waitFinished bool, response *resourcemanager.GetProjectResponse, err error) {
		p, err := a.GetProject(ctx, containerId).Execute()
		if err != nil {
			return false, nil, err
		}
		if p.ContainerId == containerId && p.LifecycleState == resourcemanager.LIFECYCLESTATE_ACTIVE {
			return true, p, nil
		}
		if p.ContainerId == containerId && p.LifecycleState == resourcemanager.LIFECYCLESTATE_CREATING {
			return false, nil, nil
		}
		return true, p, fmt.Errorf("creation failed: received project state '%s'", p.LifecycleState)
	})
	handler.SetSleepBeforeWait(1 * time.Minute)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteProjectWaitHandler will wait for project deletion
func DeleteProjectWaitHandler(ctx context.Context, a resourcemanager.DefaultAPI, containerId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetProject(ctx, containerId).Execute()
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
