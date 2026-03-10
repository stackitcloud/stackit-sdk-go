package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	observability "github.com/stackitcloud/stackit-sdk-go/services/observability/v1api"
)

const (
	GETINSTANCERESPONSESTATUS_CREATING         = "CREATING"
	GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED = "CREATE_SUCCEEDED"
	GETINSTANCERESPONSESTATUS_CREATE_FAILED    = "CREATE_FAILED"
	GETINSTANCERESPONSESTATUS_DELETING         = "DELETING"
	GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED = "DELETE_SUCCEEDED"
	GETINSTANCERESPONSESTATUS_DELETE_FAILED    = "DELETE_FAILED"
	GETINSTANCERESPONSESTATUS_UPDATING         = "UPDATING"
	GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED = "UPDATE_SUCCEEDED"
	GETINSTANCERESPONSESTATUS_UPDATE_FAILED    = "UPDATE_FAILED"
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.GetInstanceResponse, err error) {
		s, err := a.GetInstance(ctx, instanceId, projectId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil {
			return false, nil, nil
		}
		if s.Id == instanceId && s.Status == GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED {
			return true, s, nil
		}
		if s.Id == instanceId && s.Status == GETINSTANCERESPONSESTATUS_CREATE_FAILED {
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.GetInstanceResponse, err error) {
		s, err := a.GetInstance(ctx, instanceId, projectId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil {
			return false, nil, nil
		}
		// The observability instance API currently replies with create success in case the update was successful.
		if s.Id == instanceId && (s.Status == GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED || s.Status == GETINSTANCERESPONSESTATUS_CREATE_FAILED) {
			return true, s, nil
		}
		if s.Id == instanceId && (s.Status == GETINSTANCERESPONSESTATUS_UPDATE_FAILED || s.Status == GETINSTANCERESPONSESTATUS_CREATE_FAILED) {
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.GetInstanceResponse, err error) {
		s, err := a.GetInstance(ctx, instanceId, projectId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil {
			return false, nil, nil
		}
		if s.Id == instanceId && s.Status == GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED {
			return true, s, nil
		}
		if s.Id == instanceId && s.Status == GETINSTANCERESPONSESTATUS_DELETE_FAILED {
			return true, s, fmt.Errorf("delete failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// CreateScrapeConfigWaitHandler will wait for scrape config creation
func CreateScrapeConfigWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, jobName, projectId string) *wait.AsyncActionHandler[observability.ListScrapeConfigsResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.ListScrapeConfigsResponse, err error) {
		s, err := a.ListScrapeConfigs(ctx, instanceId, projectId).Execute()
		if err != nil {
			return false, nil, err
		}
		jobs := s.Data
		for i := range jobs {
			if jobs[i].JobName == jobName {
				return true, s, nil
			}
		}
		return false, nil, nil
	})
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// DeleteScrapeConfigWaitHandler will wait for scrape config deletion
func DeleteScrapeConfigWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, jobName, projectId string) *wait.AsyncActionHandler[observability.ListScrapeConfigsResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.ListScrapeConfigsResponse, err error) {
		s, err := a.ListScrapeConfigs(ctx, instanceId, projectId).Execute()
		if err != nil {
			return false, nil, err
		}
		jobs := s.Data
		for i := range jobs {
			if jobs[i].JobName == jobName {
				return false, nil, nil
			}
		}
		return true, s, nil
	})
	handler.SetTimeout(3 * time.Minute)
	return handler
}
