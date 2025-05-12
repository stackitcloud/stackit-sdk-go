package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/git"
)

const (
	// Deprecated: InstanceStateReady is deprecated and will be removed after 12th November 2025. Use [git.INSTANCESTATE_READY] instead.
	InstanceStateReady = "Ready"
	// Deprecated: InstanceStateCreating is deprecated and will be removed after 12th November 2025. Use [git.INSTANCESTATE_CREATING] instead.
	InstanceStateCreating = "Creating"
	// Deprecated: InstanceStateError is deprecated and will be removed after 12th November 2025. Use [git.INSTANCESTATE_ERROR] instead.
	InstanceStateError = "Error"
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
		if *instance.Id == instanceId && *instance.State == git.INSTANCESTATE_READY {
			return true, instance, nil
		}
		if *instance.Id == instanceId && *instance.State == git.INSTANCESTATE_ERROR {
			return true, instance, fmt.Errorf("create failed for Instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteGitInstanceWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[git.Instance] {
	handler := wait.New(func() (waitFinished bool, response *git.Instance, err error) {
		_, err = a.GetInstanceExecute(ctx, projectId, instanceId)
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
