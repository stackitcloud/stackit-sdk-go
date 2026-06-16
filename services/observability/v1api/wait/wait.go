package wait

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	observability "github.com/stackitcloud/stackit-sdk-go/services/observability/v1api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_CREATING = observability.STATUS_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED = observability.STATUS_CREATE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_CREATE_FAILED = observability.STATUS_CREATE_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_DELETING = observability.STATUS_DELETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED = observability.STATUS_DELETE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_DELETE_FAILED = observability.STATUS_DELETE_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_UPDATING = observability.STATUS_UPDATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED = observability.STATUS_UPDATE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	GETINSTANCERESPONSESTATUS_UPDATE_FAILED = observability.STATUS_UPDATE_FAILED
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[observability.GetInstanceResponse, observability.Status]{
		FetchInstance: a.GetInstance(ctx, instanceId, projectId).Execute,
		GetState: func(s *observability.GetInstanceResponse) (observability.Status, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.Id != instanceId {
				return "", fmt.Errorf("instance id mismatch: expected %s, got %s", instanceId, s.Id)
			}
			return s.Status, nil
		},
		ActiveState: []observability.Status{observability.STATUS_CREATE_SUCCEEDED},
		ErrorState:  []observability.Status{observability.STATUS_CREATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
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
		if s.Id == instanceId && (s.Status == observability.STATUS_UPDATE_SUCCEEDED || s.Status == observability.STATUS_CREATE_SUCCEEDED) {
			return true, s, nil
		}
		if s.Id == instanceId && (s.Status == observability.STATUS_UPDATE_FAILED || s.Status == observability.STATUS_CREATE_FAILED) {
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a observability.DefaultAPI, instanceId, projectId string) *wait.AsyncActionHandler[observability.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[observability.GetInstanceResponse, observability.Status]{
		FetchInstance: a.GetInstance(ctx, instanceId, projectId).Execute,
		GetState: func(s *observability.GetInstanceResponse) (observability.Status, error) {
			if s == nil {
				return "", errors.New("empty response")
			}
			if s.Id != instanceId {
				return "", fmt.Errorf("instance id mismatch: expected %s, got %s", instanceId, s.Id)
			}
			return s.Status, nil
		},
		ActiveState: []observability.Status{observability.STATUS_DELETE_SUCCEEDED},
		ErrorState:  []observability.Status{observability.STATUS_DELETE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
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
