package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
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

func CreateGitInstanceWaitHandler(ctx context.Context, a git.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		instance, err := a.GetInstance(ctx, projectId, instanceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if instance.Id == instanceId && instance.State == INSTANCESTATE_READY {
			return true, instance, nil
		}
		if instance.Id == instanceId && instance.State == INSTANCESTATE_ERROR {
			return true, instance, fmt.Errorf("create failed for Instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteGitInstanceWaitHandler(ctx context.Context, a git.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		_, err = a.GetInstance(ctx, projectId, instanceId).Execute()
		// the instances is still gettable, e.g. not deleted, when the errors is null
		if err == nil {
			return false, nil, nil
		}
		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
