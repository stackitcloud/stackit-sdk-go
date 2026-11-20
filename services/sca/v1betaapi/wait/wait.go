package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sca "github.com/stackitcloud/stackit-sdk-go/services/sca/v1betaapi"
)

// CreateApplicationWaitHandler will wait for application creation
func CreateApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, region, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	return createOrUpdateApplicationWaitHandler(ctx, a, projectID, region, environmentID, applicationID)
}

// UpdateApplicationWaitHandler will wait for application update
func UpdateApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, region, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	return createOrUpdateApplicationWaitHandler(ctx, a, projectID, region, environmentID, applicationID)
}

func createOrUpdateApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, region, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	waitConfig := wait.WaiterHelper[sca.Application, sca.ApplicationCurrentStatus]{
		FetchInstance: a.GetApplication(ctx, projectID, region, environmentID, applicationID).Execute,
		GetState:      getApplicationStatus,
		ActiveState: []sca.ApplicationCurrentStatus{
			sca.APPLICATIONCURRENTSTATUS_CURRENT_STATUS_RUNNING,
			sca.APPLICATIONCURRENTSTATUS_CURRENT_STATUS_IDLE,
			// CURRENT_STATUS_NONE implies that the application is stopped, so its active without instances
			sca.APPLICATIONCURRENTSTATUS_CURRENT_STATUS_NONE,
		},
		ErrorState: []sca.ApplicationCurrentStatus{sca.APPLICATIONCURRENTSTATUS_CURRENT_STATUS_FAILED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func getApplicationStatus(app *sca.Application) (sca.ApplicationCurrentStatus, error) {
	if app == nil {
		return "", errors.New("could not get application status: response is nil")
	}
	if app.Status == nil || app.Status.Current == nil {
		return "", errors.New("could not get application status: status is nil")
	}
	return *app.Status.Current, nil
}

// DeleteApplicationWaitHandler will wait for application deletion
func DeleteApplicationWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, region, environmentID, applicationID string) *wait.AsyncActionHandler[sca.Application] {
	waitConfig := wait.WaiterHelper[sca.Application, sca.ApplicationCurrentStatus]{
		FetchInstance:              a.GetApplication(ctx, projectID, region, environmentID, applicationID).Execute,
		GetState:                   getApplicationStatus,
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateEnvironmentWaitHandler will wait for environment creation
func CreateEnvironmentWaitHandler(ctx context.Context, a sca.DefaultAPI, projectID, region, environmentID string) *wait.AsyncActionHandler[sca.Environment] {
	waitConfig := wait.WaiterHelper[sca.Environment, sca.EnvironmentCurrentStatus]{
		FetchInstance: a.GetEnvironment(ctx, projectID, region, environmentID).Execute,
		GetState:      getEnvironmentStatus,
		ActiveState: []sca.EnvironmentCurrentStatus{
			sca.ENVIRONMENTCURRENTSTATUS_CURRENT_STATUS_CREATED,
		},
		ErrorState: []sca.EnvironmentCurrentStatus{sca.ENVIRONMENTCURRENTSTATUS_CURRENT_STATUS_FAILED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func getEnvironmentStatus(env *sca.Environment) (sca.EnvironmentCurrentStatus, error) {
	if env == nil {
		return "", errors.New("could not get environment status: response is nil")
	}
	if env.Status == nil || env.Status.Current == nil {
		return "", errors.New("could not get environment status: status is nil")
	}
	return *env.Status.Current, nil
}
