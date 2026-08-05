package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	workflows "github.com/stackitcloud/stackit-sdk-go/services/workflows/v1alphaapi"
)

// CreateInstanceWaitHandler will wait for STACKIT Workflows instance creation.
// Returns the instance once it reaches the active state.
func CreateInstanceWaitHandler(ctx context.Context, a workflows.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[workflows.Instance] {
	waitConfig := wait.WaiterHelper[workflows.Instance, workflows.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *workflows.Instance) (workflows.InstanceStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []workflows.InstanceStatus{workflows.INSTANCESTATUS_ACTIVE},
		ErrorState:  []workflows.InstanceStatus{workflows.INSTANCESTATUS_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// UpdateInstanceWaitHandler will wait for STACKIT Workflows instance updates.
func UpdateInstanceWaitHandler(ctx context.Context, a workflows.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[workflows.Instance] {
	waitConfig := wait.WaiterHelper[workflows.Instance, workflows.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *workflows.Instance) (workflows.InstanceStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []workflows.InstanceStatus{workflows.INSTANCESTATUS_ACTIVE},
		ErrorState:  []workflows.InstanceStatus{workflows.INSTANCESTATUS_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for STACKIT Workflows instance deletion.
func DeleteInstanceWaitHandler(ctx context.Context, a workflows.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[workflows.Instance] {
	waitConfig := wait.WaiterHelper[workflows.Instance, workflows.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *workflows.Instance) (workflows.InstanceStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ErrorState:                 []workflows.InstanceStatus{workflows.INSTANCESTATUS_FAILED},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
