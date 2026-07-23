package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	postgresflex "github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v3api"
)

// createOrUpdateInstanceWaitHandler will wait for instance creation
func createOrUpdateInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.GetInstanceResponse, postgresflex.State]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *postgresflex.GetInstanceResponse) (postgresflex.State, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.State == "" {
				return "", errors.New("state is missing")
			}
			return response.State, nil
		},
		ActiveState: []postgresflex.State{postgresflex.STATE_READY},
		ErrorState: []postgresflex.State{
			postgresflex.STATE_FAILURE,
			postgresflex.STATE_UNKNOWN,
			postgresflex.STATE_TERMINATING,
		},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, region, instanceId).SetSleepBeforeWait(15 * time.Second)
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, region, instanceId)
}

// CloneInstanceWaitHandler will wait for instance cloning
func CloneInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, region, instanceId)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.GetInstanceResponse, postgresflex.State]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *postgresflex.GetInstanceResponse) (postgresflex.State, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.State == "" {
				return "", errors.New("status is missing in response")
			}
			return response.State, nil
		},
		ErrorState:                 []postgresflex.State{postgresflex.STATE_FAILURE},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateUserWaitHandler will wait for user creation
func CreateUserWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string, userId int64) *wait.AsyncActionHandler[postgresflex.GetUserResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.GetUserResponse, string]{
		FetchInstance: client.GetUser(ctx, projectId, region, instanceId, userId).Execute,
		GetState: func(response *postgresflex.GetUserResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState: []string{"PROCESSED"},
		ErrorState:  []string{},
		// The API does not have a dedicated failure state for this resource,
		// so we rely on the timeout for cases where it never becomes active.
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteUserWaitHandler will wait for user deletion
func DeleteUserWaitHandler(ctx context.Context, a postgresflex.DefaultAPI, projectId, region, instanceId string, userId int64) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetUser(ctx, projectId, region, instanceId, userId).Execute()
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode != 404 {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}
