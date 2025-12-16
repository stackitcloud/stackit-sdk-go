package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/logs"
)

type APIClientInterface interface {
	GetLogsInstanceExecute(ctx context.Context, projectId string, regionId string, instanceId string) (*logs.LogsInstance, error)
}

func CreateLogsInstanceWaitHandler(ctx context.Context, client APIClientInterface, projectID, region, instanceID string) *wait.AsyncActionHandler[logs.LogsInstance] {
	handler := wait.New(func() (waitFinished bool, response *logs.LogsInstance, err error) {
		instance, err := client.GetLogsInstanceExecute(ctx, projectID, region, instanceID)
		if err != nil {
			return false, nil, err
		}
		if instance.Id == nil || instance.Status == nil {
			return false, nil, fmt.Errorf("get instance, project: %q, region: %q, instanceID: %q: missing id or status", projectID, region, instanceID)
		}
		if *instance.Id == instanceID && *instance.Status == logs.LOGSINSTANCESTATUS_ACTIVE {
			return true, instance, nil
		}
		if *instance.Status == logs.LOGSINSTANCESTATUS_DELETING {
			return true, nil, fmt.Errorf("creating log instance failed, instance is being deleted")
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteLogsInstanceWaitHandler(ctx context.Context, client APIClientInterface, projectID, region, instanceID string) *wait.AsyncActionHandler[logs.LogsInstance] {
	handler := wait.New(func() (waitFinished bool, response *logs.LogsInstance, err error) {
		_, err = client.GetLogsInstanceExecute(ctx, projectID, region, instanceID)
		// the instances is still gettable, e.g. not deleted, when the errors is null
		if err == nil {
			return false, nil, nil
		}
		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
