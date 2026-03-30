package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	resourcemanager "github.com/stackitcloud/stackit-sdk-go/services/resourcemanager/v0api"
)

// CreateProjectWaitHandler will wait for project creation
func CreateProjectWaitHandler(ctx context.Context, a resourcemanager.DefaultAPI, containerId string) *wait.AsyncActionHandler[resourcemanager.GetProjectResponse] {
	waitConfig := wait.WaiterHelper[resourcemanager.GetProjectResponse, resourcemanager.LifecycleState]{
		FetchInstance: a.GetProject(ctx, containerId).Execute,
		GetState: func(project *resourcemanager.GetProjectResponse) (resourcemanager.LifecycleState, error) {
			if project == nil {
				return "", errors.New("empty response")
			}
			return project.LifecycleState, nil
		},
		ActiveState: []resourcemanager.LifecycleState{resourcemanager.LIFECYCLESTATE_ACTIVE},
		ErrorState:  []resourcemanager.LifecycleState{resourcemanager.LIFECYCLESTATE_INACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(1 * time.Minute)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteProjectWaitHandler will wait for project deletion
func DeleteProjectWaitHandler(ctx context.Context, a resourcemanager.DefaultAPI, containerId string) *wait.AsyncActionHandler[resourcemanager.GetProjectResponse] {
	waitConfig := wait.WaiterHelper[resourcemanager.GetProjectResponse, resourcemanager.LifecycleState]{
		FetchInstance: a.GetProject(ctx, containerId).Execute,
		GetState: func(project *resourcemanager.GetProjectResponse) (resourcemanager.LifecycleState, error) {
			if project == nil {
				return "", errors.New("empty response")
			}
			return project.LifecycleState, nil
		},
		ErrorState:                 []resourcemanager.LifecycleState{resourcemanager.LIFECYCLESTATE_INACTIVE},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound, http.StatusForbidden},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
