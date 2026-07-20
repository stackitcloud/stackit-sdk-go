package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sqlserverflex "github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3api"
)

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
