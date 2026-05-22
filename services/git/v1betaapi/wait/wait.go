package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	git "github.com/stackitcloud/stackit-sdk-go/services/git/v1betaapi"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATE_CREATING = git.INSTANCESTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATE_WAITING_FOR_RESOURCES = git.INSTANCESTATE_WAITING_FOR_RESOURCES
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATE_UPDATING = git.INSTANCESTATE_UPDATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATE_DELETING = git.INSTANCESTATE_DELETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATE_READY = git.INSTANCESTATE_READY
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATE_ERROR = git.INSTANCESTATE_ERROR
)

func CreateGitInstanceWaitHandler(ctx context.Context, client git.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	waitConfig := wait.WaiterHelper[git.Instance, git.InstanceState]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId).Execute,
		GetState: func(instance *git.Instance) (git.InstanceState, error) {
			if instance == nil {
				return "", errors.New("empty response")
			}
			return instance.State, nil
		},
		ActiveState: []git.InstanceState{git.INSTANCESTATE_READY},
		ErrorState:  []git.InstanceState{git.INSTANCESTATE_ERROR},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteGitInstanceWaitHandler(ctx context.Context, client git.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	waitConfig := wait.WaiterHelper[git.Instance, git.InstanceState]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId).Execute,
		GetState: func(instance *git.Instance) (git.InstanceState, error) {
			if instance == nil {
				return "", errors.New("empty response")
			}
			return instance.State, nil
		},
		ActiveState:                []git.InstanceState{},
		ErrorState:                 []git.InstanceState{git.INSTANCESTATE_ERROR},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
