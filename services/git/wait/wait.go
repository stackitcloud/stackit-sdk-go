package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/git"
)

const (
	InstanceStateReady    = "Ready"
	InstanceStateCreating = "Creating"
	InstanceStateError    = "Error"
)

// APIClientInterface Interfaces needed for tests
type APIClientInterface interface {
	GetInstanceExecute(ctx context.Context, projectId string, instanceId string) (*git.Instance, error)
}

func CreateGitInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		instance, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if instance.Id == nil || instance.State == nil {
			return false, nil, fmt.Errorf("could not get Instance id or State from response for project %s and instanceId %s", projectId, instanceId)
		}
		if *instance.Id == instanceId && *instance.State == InstanceStateReady {
			return true, instance, nil
		}
		if *instance.Id == instanceId && *instance.State == InstanceStateError {
			return true, instance, fmt.Errorf("create failed for Instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteGitInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		instance, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if instance != nil {
			return false, instance, nil
		}
		return true, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
