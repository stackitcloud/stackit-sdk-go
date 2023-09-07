package argus

import (
	"context"
	"fmt"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

const (
	CreateSuccess = "CREATE_SUCCEEDED"
	CreateFail    = "CREATE_FAILED"
	UpdateSuccess = "UPDATE_SUCCEEDED"
	UpdateFail    = "UPDATE_FAILED"
	DeleteSuccess = "DELETE_SUCCEEDED"
	DeleteFail    = "DELETE_FAILED"
)

// APIClientInterface Interfaces needed for tests
type APIClientInterface interface {
	GetInstanceExecute(ctx context.Context, instanceId, projectId string) (*InstanceResponse, error)
	GetScrapeConfigsExecute(ctx context.Context, instanceId, projectId string) (*ScrapeConfigsResponse, error)
}

func handleError(reqErr error) (res interface{}, done bool, err error) {
	oapiErr, ok := reqErr.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
	if !ok {
		return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
	}
	// Some APIs may return temporary errors and the request should be retried
	if utils.Contains(wait.RetryHttpErrorStatusCodes, oapiErr.statusCode) {
		return nil, false, nil
	}
	return nil, false, reqErr
}

//	will wait for creation
//
// returned interface is nil or *InstanceResponse
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return handleError(err)
		}
		if s.Id == nil || s.Status == nil {
			return s, false, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		if *s.Id == instanceId && *s.Status == CreateSuccess {
			return s, true, nil
		}
		if *s.Id == instanceId && *s.Status == CreateFail {
			return s, true, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return s, false, nil
	})
}

// UpdateInstanceWaitHandler will wait for update
// returned interface is nil or *InstanceResponse
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return handleError(err)
		}
		if s.Id == nil || s.Status == nil {
			return s, false, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		// The argus instance API currently replies with create success in case the update was successful.
		if *s.Id == instanceId && (*s.Status == UpdateSuccess || *s.Status == CreateSuccess) {
			return s, true, nil
		}
		if *s.Id == instanceId && (*s.Status == UpdateFail || *s.Status == CreateFail) {
			return s, true, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
		return s, false, nil
	})
}

// DeleteInstanceWaitHandler will wait for delete
// returned interface is nil or *InstanceResponse
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return handleError(err)
		}
		if s.Id == nil || s.Status == nil {
			return s, false, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		if *s.Id == instanceId && *s.Status == DeleteSuccess {
			return s, true, nil
		}
		if *s.Id == instanceId && *s.Status == DeleteFail {
			return s, true, fmt.Errorf("delete failed for instance with id %s", instanceId)
		}
		return s, false, nil
	})
}

func CreateScrapeConfigWaitHandler(ctx context.Context, a APIClientInterface, instanceId, jobName, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetScrapeConfigsExecute(ctx, instanceId, projectId)
		if err != nil {
			return handleError(err)
		}
		jobs := *s.Data
		for i := range jobs {
			if *jobs[i].JobName == jobName {
				return s, true, nil
			}
		}
		return s, false, nil
	})
}

func DeleteScrapeConfigWaitHandler(ctx context.Context, a APIClientInterface, instanceId, jobName, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetScrapeConfigsExecute(ctx, instanceId, projectId)
		if err != nil {
			return handleError(err)
		}
		jobs := *s.Data
		for i := range jobs {
			if *jobs[i].JobName == jobName {
				return s, false, nil
			}
		}
		return s, true, nil
	})
}
