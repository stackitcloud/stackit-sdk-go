package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/git"
)

const (
	CreateSuccess = "Ready"
	Creating      = "Creating"
	CreateFail    = "Error"
)

// APIClientInterface Interfaces needed for tests
type APIClientInterface interface {
	GetGitExecute(ctx context.Context, projectId string, instanceId string) (*git.Instance, error)
}

func CreateGitInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		instance, err := a.GetGitExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if *instance.Id == "" || *instance.State == "" {
			return false, nil, fmt.Errorf("could not get Instance id or State from response for project %s and instanceId %s", projectId, instanceId)
		}
		if *instance.Id == instanceId && *instance.State == CreateSuccess {
			return true, instance, nil
		}
		if *instance.Id == instanceId && *instance.State == CreateFail {
			return true, instance, fmt.Errorf("create failed for Instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteGitInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		instance, err := a.GetGitExecute(ctx, projectId, instanceId)
		if err != nil {
			return true, nil, err
		}
		if instance != nil {
			return true, instance, nil
		}
		return true, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
