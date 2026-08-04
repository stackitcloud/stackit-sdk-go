package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v1api"
)

// RunCommandWaitHandler will wait for a run command to reach a terminal state (completed or failed).
// Both completed and failed are treated as active states; the caller should inspect the returned
// CommandDetails.Status to distinguish success from failure.
func RunCommandWaitHandler(ctx context.Context, a runcommand.DefaultAPI, projectId, serverId, commandId string) *wait.AsyncActionHandler[runcommand.CommandDetails] {
	waitConfig := wait.WaiterHelper[runcommand.CommandDetails, runcommand.CommandDetailsStatus]{
		FetchInstance: a.GetCommand(ctx, projectId, serverId, commandId).Execute,
		GetState: func(d *runcommand.CommandDetails) (runcommand.CommandDetailsStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			status, ok := d.GetStatusOk()
			if !ok {
				return "", errors.New("no status in response")
			}
			return *status, nil
		},
		ActiveState: []runcommand.CommandDetailsStatus{
			runcommand.COMMANDDETAILSSTATUS_COMPLETED,
			runcommand.COMMANDDETAILSSTATUS_FAILED,
		},
		ErrorState: []runcommand.CommandDetailsStatus{},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
