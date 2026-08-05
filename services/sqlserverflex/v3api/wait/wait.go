package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sqlserverflex "github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3api"
)

const userActiveState = "PROCESSED"

func createOrUpdateInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[sqlserverflex.GetInstanceResponse, sqlserverflex.State]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *sqlserverflex.GetInstanceResponse) (sqlserverflex.State, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.State == "" {
				return "", errors.New("state is missing")
			}
			return response.State, nil
		},
		ActiveState: []sqlserverflex.State{sqlserverflex.STATE_READY},
		ErrorState: []sqlserverflex.State{
			sqlserverflex.STATE_FAILURE,
			sqlserverflex.STATE_UNKNOWN,
			sqlserverflex.STATE_TERMINATING,
		},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(5 * time.Second)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, region, instanceId)
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, region, instanceId)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[sqlserverflex.GetInstanceResponse, sqlserverflex.State]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *sqlserverflex.GetInstanceResponse) (sqlserverflex.State, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.State == "" {
				return "", errors.New("state is missing in response")
			}
			return response.State, nil
		},
		ErrorState:                 []sqlserverflex.State{sqlserverflex.STATE_FAILURE},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateUserWaitHandler will wait for user creation
func CreateUserWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, region, instanceId string, userId int64) *wait.AsyncActionHandler[sqlserverflex.GetUserResponse] {
	waitConfig := wait.WaiterHelper[sqlserverflex.GetUserResponse, string]{
		FetchInstance: client.GetUser(ctx, projectId, region, instanceId, userId).Execute,
		GetState: func(resp *sqlserverflex.GetUserResponse) (string, error) {
			if resp == nil {
				return "", errors.New("empty response")
			}
			if resp.Status == "" {
				return "", errors.New("state is missing in response")
			}
			return resp.Status, nil
		},
		ActiveState: []string{userActiveState},
		ErrorState:  []string{},
		// The API does not have a dedicated failure state for this resource,
		// so we rely on the timeout for cases where it never becomes active.
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(5 * time.Second)
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteUserWaitHandler will wait for user deletion
func DeleteUserWaitHandler(ctx context.Context, a sqlserverflex.DefaultAPI, projectId, region, instanceId string, userId int64) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetUser(ctx, projectId, region, instanceId, userId).Execute()
		if err == nil {
			return false, nil, nil
		}
		var oapiErr *oapierror.GenericOpenAPIError
		if !errors.As(err, &oapiErr) {
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
