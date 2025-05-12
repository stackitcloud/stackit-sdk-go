package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/observability"
)

const (
	// Deprecated: CreateSuccess is deprecated and will be removed after 12th November 2025. Use [observability.GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED] instead.
	CreateSuccess = "CREATE_SUCCEEDED"
	// Deprecated: CreateFail is deprecated and will be removed after 12th November 2025. Use [observability.GETINSTANCERESPONSESTATUS_CREATE_FAILED] instead.
	CreateFail = "CREATE_FAILED"
	// Deprecated: UpdateSuccess is deprecated and will be removed after 12th November 2025. Use [observability.GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED] instead.
	UpdateSuccess = "UPDATE_SUCCEEDED"
	// Deprecated: UpdateFail is deprecated and will be removed after 12th November 2025. Use [observability.GETINSTANCERESPONSESTATUS_UPDATE_FAILED] instead.
	UpdateFail = "UPDATE_FAILED"
	// Deprecated: DeleteSuccess is deprecated and will be removed after 12th November 2025. Use [observability.GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED] instead.
	DeleteSuccess = "DELETE_SUCCEEDED"
	// Deprecated: DeleteFail is deprecated and will be removed after 12th November 2025. Use [observability.GETINSTANCERESPONSESTATUS_DELETE_FAILED] instead.
	DeleteFail = "DELETE_FAILED"
)

// APIClientInterface Interfaces needed for tests
type APIClientInterface interface {
	GetInstanceExecute(ctx context.Context, instanceId, projectId string) (*observability.GetInstanceResponse, error)
	ListScrapeConfigsExecute(ctx context.Context, instanceId, projectId string) (*observability.ListScrapeConfigsResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		if *s.Id == instanceId && *s.Status == observability.GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Id == instanceId && *s.Status == observability.GETINSTANCERESPONSESTATUS_CREATE_FAILED {
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		// The observability instance API currently replies with create success in case the update was successful.
		if *s.Id == instanceId && (*s.Status == observability.GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED || *s.Status == observability.GETINSTANCERESPONSESTATUS_CREATE_FAILED) {
			return true, s, nil
		}
		if *s.Id == instanceId && (*s.Status == observability.GETINSTANCERESPONSESTATUS_UPDATE_FAILED || *s.Status == observability.GETINSTANCERESPONSESTATUS_CREATE_FAILED) {
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInterface, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, instanceId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get instance id or status from response for project %s and instance %s", projectId, instanceId)
		}
		if *s.Id == instanceId && *s.Status == observability.GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Id == instanceId && *s.Status == observability.GETINSTANCERESPONSESTATUS_DELETE_FAILED {
			return true, s, fmt.Errorf("delete failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// CreateScrapeConfigWaitHandler will wait for scrape config creation
func CreateScrapeConfigWaitHandler(ctx context.Context, a APIClientInterface, instanceId, jobName, projectId string) *wait.AsyncActionHandler[observability.ListScrapeConfigsResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.ListScrapeConfigsResponse, err error) {
		s, err := a.ListScrapeConfigsExecute(ctx, instanceId, projectId)
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
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// DeleteScrapeConfigWaitHandler will wait for scrape config deletion
func DeleteScrapeConfigWaitHandler(ctx context.Context, a APIClientInterface, instanceId, jobName, projectId string) *wait.AsyncActionHandler[observability.ListScrapeConfigsResponse] {
	handler := wait.New(func() (waitFinished bool, response *observability.ListScrapeConfigsResponse, err error) {
		s, err := a.ListScrapeConfigsExecute(ctx, instanceId, projectId)
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
	handler.SetTimeout(3 * time.Minute)
	return handler
}
