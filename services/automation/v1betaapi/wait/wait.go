package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	automation "github.com/stackitcloud/stackit-sdk-go/services/automation/v1betaapi"
)

// CreateVolumeExecutionWaitHandler will wait for a volume automation execution to finish.
func CreateVolumeExecutionWaitHandler(ctx context.Context, client automation.DefaultAPI, projectId, region, automationId, executionId string) *wait.AsyncActionHandler[automation.VolumeExecutionResponse] {
	waitConfig := wait.WaiterHelper[automation.VolumeExecutionResponse, automation.VolumeExecutionResponseStatus]{
		FetchInstance: client.GetVolumeExecution(ctx, projectId, region, automationId, executionId).Execute,
		GetState: func(response *automation.VolumeExecutionResponse) (automation.VolumeExecutionResponseStatus, error) {
			if response == nil {
				return "", errors.New("execution is nil")
			}
			return response.Status, nil
		},
		ActiveState: []automation.VolumeExecutionResponseStatus{automation.VOLUMEEXECUTIONRESPONSESTATUS_COMPLETED},
		ErrorState: []automation.VolumeExecutionResponseStatus{
			automation.VOLUMEEXECUTIONRESPONSESTATUS_FAILED,
			automation.VOLUMEEXECUTIONRESPONSESTATUS_TERMINATED,
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}
