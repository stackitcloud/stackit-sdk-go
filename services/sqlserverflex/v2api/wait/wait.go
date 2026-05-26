package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sqlserverflex "github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v2api"
)

const (
	InstanceStateEmpty      = ""
	InstanceStateProcessing = "Progressing"
	InstanceStateUnknown    = "Unknown"
	InstanceStateSuccess    = "Ready"
	InstanceStateFailed     = "Failed"
)

func createOrUpdateInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[sqlserverflex.GetInstanceResponse, string]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState: func(response *sqlserverflex.GetInstanceResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Item == nil {
				return "", errors.New("empty instance")
			}
			if response.Item.Status == nil {
				return "", errors.New("status is missing")
			}
			return *response.Item.Status, nil
		},
		ActiveState: []string{InstanceStateSuccess},
		ErrorState:  []string{InstanceStateUnknown, InstanceStateFailed, InstanceStateEmpty},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(5 * time.Second)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, instanceId, region)
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, client, projectId, instanceId, region)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, client sqlserverflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[sqlserverflex.GetInstanceResponse, string]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState: func(response *sqlserverflex.GetInstanceResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Item.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Item.Status, nil
		},
		ErrorState:                 []string{InstanceStateFailed},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
