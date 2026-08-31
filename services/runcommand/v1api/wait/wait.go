package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v1api"
)

// AgentReadyWaitHandler retries CreateCommand until the server agent registers.
// The API returns 404 while the agent is booting; any other error is terminal.
// On success, it returns the NewCommandResponse with the submitted command ID.
func AgentReadyWaitHandler(ctx context.Context, a runcommand.DefaultAPI, projectId, serverId string, payload runcommand.CreateCommandPayload) *wait.AsyncActionHandler[runcommand.NewCommandResponse] {
	handler := wait.New(func() (bool, *runcommand.NewCommandResponse, error) {
		resp, err := a.CreateCommand(ctx, projectId, serverId).CreateCommandPayload(payload).Execute()
		if err != nil {
			var oapiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiErr) && oapiErr.StatusCode == http.StatusNotFound {
				return false, nil, nil
			}
			return false, nil, err
		}
		return true, resp, nil
	})
	handler.SetThrottle(10 * time.Second)
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// RunCommandWaitHandler will wait for a run command to reach a terminal state (completed or failed).
// Both completed and failed are treated as active states; the caller should inspect the returned
// CommandDetails.Status to distinguish success from failure.
func RunCommandWaitHandler(ctx context.Context, a runcommand.DefaultAPI, projectId, serverId, commandId string) *wait.AsyncActionHandler[runcommand.CommandDetails] {
	waitConfig := wait.WaiterHelper[runcommand.CommandDetails, runcommand.CommandDetailsStatus]{
		FetchInstance: a.GetCommand(ctx, projectId, serverId, commandId).Execute,
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
			runcommand.COMMANDDETAILSSTATUS_FAILED,
		},
		ErrorState: []runcommand.CommandDetailsStatus{},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}
