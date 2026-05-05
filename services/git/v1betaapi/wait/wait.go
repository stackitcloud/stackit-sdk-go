package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	git "github.com/stackitcloud/stackit-sdk-go/services/git/v1betaapi"
)

const (
	INSTANCESTATE_CREATING              = "Creating"
	INSTANCESTATE_WAITING_FOR_RESOURCES = "WaitingForResources"
	INSTANCESTATE_UPDATING              = "Updating"
	INSTANCESTATE_DELETING              = "Deleting"
	INSTANCESTATE_READY                 = "Ready"
	INSTANCESTATE_ERROR                 = "Error"
)

func CreateGitInstanceWaitHandler(ctx context.Context, client git.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	waitConfig := wait.WaiterHelper[git.Instance, string]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId).Execute,
		GetState: func(instance *git.Instance) (string, error) {
			if instance == nil {
				return "", errors.New("empty response")
			}
			return instance.State, nil
		},
		ActiveState: []string{INSTANCESTATE_READY},
		ErrorState:  []string{INSTANCESTATE_ERROR},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteGitInstanceWaitHandler(ctx context.Context, client git.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	waitConfig := wait.WaiterHelper[git.Instance, string]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId).Execute,
		GetState: func(instance *git.Instance) (string, error) {
			if instance == nil {
				return "", errors.New("empty response")
			}
			return instance.State, nil
		},
		ActiveState: []string{},
		ErrorState:  []string{INSTANCESTATE_ERROR},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
