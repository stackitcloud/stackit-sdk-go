package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	logs "github.com/stackitcloud/stackit-sdk-go/services/logs/v1api"
)

// CreateLogsInstanceWaitHandler will wait for logs instance creation
func CreateLogsInstanceWaitHandler(ctx context.Context, client logs.DefaultAPI, projectID string, region string, instanceID string) *wait.AsyncActionHandler[logs.LogsInstance] {
	waitConfig := wait.WaiterHelper[logs.LogsInstance, logs.LogsInstanceStatus]{
		FetchInstance: client.GetLogsInstance(ctx, projectID, region, instanceID).Execute,
		GetState: func(l *logs.LogsInstance) (logs.LogsInstanceStatus, error) {
			if l == nil {
				return "", fmt.Errorf("empty response")
			}
			return l.Status, nil
		},
		ActiveState: []logs.LogsInstanceStatus{logs.LOGSINSTANCESTATUS_ACTIVE},
		ErrorState:  []logs.LogsInstanceStatus{logs.LOGSINSTANCESTATUS_DELETING},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteLogsInstanceWaitHandler will wait for logs instance deletion
func DeleteLogsInstanceWaitHandler(ctx context.Context, client logs.DefaultAPI, projectID string, region string, instanceID string) *wait.AsyncActionHandler[logs.LogsInstance] {
	waitConfig := wait.WaiterHelper[logs.LogsInstance, logs.LogsInstanceStatus]{
		FetchInstance: client.GetLogsInstance(ctx, projectID, region, instanceID).Execute,
		GetState: func(l *logs.LogsInstance) (logs.LogsInstanceStatus, error) {
			if l == nil {
				return "", errors.New("empty response")
			}
			return l.Status, nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
