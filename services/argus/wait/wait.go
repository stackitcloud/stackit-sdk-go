package wait

import (
	"context"
	"fmt"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/argus"
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
	GetInstanceExecute(ctx context.Context, instanceId, projectId string) (*argus.InstanceResponse, error)
	GetScrapeConfigsExecute(ctx context.Context, instanceId, projectId string) (*argus.ScrapeConfigsResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.AsyncActionHandler[argus.InstanceResponse] {
	return wait.New(func() (waitFinished bool, response *argus.InstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		if *s.Id == instanceId && *s.Status == CreateSuccess {
			return true, s, nil
		}
		if *s.Id == instanceId && *s.Status == CreateFail {
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.AsyncActionHandler[argus.InstanceResponse] {
	return wait.New(func() (waitFinished bool, response *argus.InstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		// The argus instance API currently replies with create success in case the update was successful.
		if *s.Id == instanceId && (*s.Status == UpdateSuccess || *s.Status == CreateSuccess) {
			return true, s, nil
		}
		if *s.Id == instanceId && (*s.Status == UpdateFail || *s.Status == CreateFail) {
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.AsyncActionHandler[argus.InstanceResponse] {
	return wait.New(func() (waitFinished bool, response *argus.InstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		if *s.Id == instanceId && *s.Status == DeleteSuccess {
			return true, s, nil
		}
		if *s.Id == instanceId && *s.Status == DeleteFail {
			return true, s, fmt.Errorf("delete failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
}

// CreateScrapeConfigWaitHandler will wait for scrape config creation
func CreateScrapeConfigWaitHandler(ctx context.Context, a APIClientInterface, instanceId, jobName, projectId string) *wait.AsyncActionHandler[argus.ScrapeConfigsResponse] {
	return wait.New(func() (waitFinished bool, response *argus.ScrapeConfigsResponse, err error) {
		s, err := a.GetScrapeConfigsExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		jobs := *s.Data
		for i := range jobs {
			if *jobs[i].JobName == jobName {
				return true, s, nil
			}
		}
		return false, nil, nil
	})
}

// DeleteScrapeConfigWaitHandler will wait for scrape config deletion
func DeleteScrapeConfigWaitHandler(ctx context.Context, a APIClientInterface, instanceId, jobName, projectId string) *wait.AsyncActionHandler[argus.ScrapeConfigsResponse] {
	return wait.New(func() (waitFinished bool, response *argus.ScrapeConfigsResponse, err error) {
		s, err := a.GetScrapeConfigsExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		jobs := *s.Data
		for i := range jobs {
			if *jobs[i].JobName == jobName {
				return false, nil, nil
			}
		}
		return true, s, nil
	})
}
