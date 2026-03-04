package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	logs "github.com/stackitcloud/stackit-sdk-go/services/logs/v1api"
)

const (
	instanceStatusActive   = "active"
	instanceStatusDeleting = "deleting"
)

func CreateLogsInstanceWaitHandler(ctx context.Context, client logs.DefaultAPI, projectID, region, instanceID string) *wait.AsyncActionHandler[logs.LogsInstance] {
	handler := wait.New(func() (waitFinished bool, response *logs.LogsInstance, err error) {
		instance, err := client.GetLogsInstance(ctx, projectID, region, instanceID).Execute()
		if err != nil {
			return false, nil, err
		}
		if instance.Id == instanceID && instance.Status == instanceStatusActive {
			return true, instance, nil
		}
		if instance.Status == instanceStatusDeleting {
			return true, nil, fmt.Errorf("creating log instance failed, instance is being deleted")
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteLogsInstanceWaitHandler(ctx context.Context, client logs.DefaultAPI, projectID, region, instanceID string) *wait.AsyncActionHandler[logs.LogsInstance] {
	handler := wait.New(func() (waitFinished bool, response *logs.LogsInstance, err error) {
		_, err = client.GetLogsInstance(ctx, projectID, region, instanceID).Execute()
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
