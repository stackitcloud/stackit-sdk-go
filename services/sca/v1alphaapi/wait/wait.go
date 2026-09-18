package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sca "github.com/stackitcloud/stackit-sdk-go/services/sca/v1alphaapi"
)

// CreateApplicationWaitHandler will wait for application creation
func CreateApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	return createOrUpdateApplicationWaitHandler(ctx, a, projectID, environmentID, applicationID)
}

// UpdateApplicationWaitHandler will wait for application update
func UpdateApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	return createOrUpdateApplicationWaitHandler(ctx, a, projectID, environmentID, applicationID)
}

func createOrUpdateApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	waitConfig := wait.WaiterHelper[sca.Application, sca.CurrentStatus]{
		FetchInstance: a.GetApplication(ctx, projectID, environmentID, applicationID).Execute,
		GetState:      getApplicationState,
		ActiveState:   []sca.CurrentStatus{sca.CURRENTSTATUS_CURRENT_STATUS_RUNNING, sca.CURRENTSTATUS_CURRENT_STATUS_IDLE},
		// There is an issue where the API reports a transient FAILED status during the startup of applications.
		// It will be solved before the beta.
		// ErrorState:    []sca.CurrentStatus{sca.CURRENTSTATUS_CURRENT_STATUS_FAILED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func getApplicationState(app *sca.Application) (sca.CurrentStatus, error) {
	if app == nil {
		return "", errors.New("could not get application status: response is nil")
	}
	if app.RuntimeStatus == nil || app.RuntimeStatus.CurrentStatus == nil {
		return "", errors.New("could not get application status: status is nil")
	}
	return *app.RuntimeStatus.CurrentStatus, nil
}

// DeleteApplicationWaitHandler will wait for application deletion
func DeleteApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	waitConfig := wait.WaiterHelper[sca.Application, sca.CurrentStatus]{
		FetchInstance:              a.GetApplication(ctx, projectID, environmentID, applicationID).Execute,
		GetState:                   getApplicationState,
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}
