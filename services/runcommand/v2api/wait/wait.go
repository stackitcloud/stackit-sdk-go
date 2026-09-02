package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v2api"
)

// RunCommandWaitHandler will wait for a run command to reach a terminal state.
// COMPLETED is treated as success and returns the CommandDetails with no error.
// FAILED is treated as an error: the handler returns a non-nil error and the
// CommandDetails (containing exit code and output) so callers can surface
// diagnostic information without an additional API call.
func RunCommandWaitHandler(ctx context.Context, a runcommand.DefaultAPI, projectId, serverId, region, commandId string) *wait.AsyncActionHandler[runcommand.CommandDetails] {
	waitConfig := wait.WaiterHelper[runcommand.CommandDetails, runcommand.CommandDetailsStatus]{
		FetchInstance: a.GetCommand(ctx, projectId, region, serverId, commandId).Execute,
		GetState: func(d *runcommand.CommandDetails) (runcommand.CommandDetailsStatus, error) {
			if d == nil {
				return "", fmt.Errorf("failed to get command %s: empty response", commandId)
			}
			status, ok := d.GetStatusOk()
			if !ok {
				return "", fmt.Errorf("command %s: status missing in response", commandId)
			}
			return *status, nil
		},
		ActiveState: []runcommand.CommandDetailsStatus{
			runcommand.COMMANDDETAILSSTATUS_COMPLETED,
		},
		ErrorState: []runcommand.CommandDetailsStatus{
			runcommand.COMMANDDETAILSSTATUS_FAILED,
		},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}
