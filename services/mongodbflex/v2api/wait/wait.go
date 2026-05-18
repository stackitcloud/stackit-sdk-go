package wait

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	mongodbflex "github.com/stackitcloud/stackit-sdk-go/services/mongodbflex/v2api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_READY = mongodbflex.INSTANCESTATUS_READY
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_PENDING = mongodbflex.INSTANCESTATUS_PENDING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_PROCESSING = mongodbflex.INSTANCESTATUS_PROCESSING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_FAILED = mongodbflex.INSTANCESTATUS_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_UNKNOWN = mongodbflex.INSTANCESTATUS_UNKNOWN

	RestoreJobProcessing = "IN_PROGRESS"
	RestoreJobFinished   = "FINISHED"
	RestoreJobBroken     = "BROKEN"
	RestoreJobKilled     = "KILLED"
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	waitConfig := wait.WaiterHelper[mongodbflex.InstanceResponse, string]{
		FetchInstance: a.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState:      getStateInstance,
		ActiveState:   []string{INSTANCESTATUS_READY},
		ErrorState:    []string{INSTANCESTATUS_FAILED},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(45 * time.Minute)
	handler.SetSleepBeforeWait(5 * time.Second)
	return handler
}

// CloneInstanceWaitHandler will wait for instance clone to be created
func CloneInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	return CreateInstanceWaitHandler(ctx, a, projectId, instanceId, region)
}

func RestoreInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, backupId, region string) *wait.AsyncActionHandler[mongodbflex.ListRestoreJobsResponse] {
	handler := wait.New(func() (waitFinished bool, response *mongodbflex.ListRestoreJobsResponse, err error) {
		s, err := a.ListRestoreJobs(ctx, projectId, instanceId, region).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Items == nil {
			return false, nil, nil
		}

		restoreJobsSlice := s.Items

		// sort array by descending date
		sort.Slice(restoreJobsSlice, func(i, j int) bool {
			// swap elements to sort by descending order
			return *restoreJobsSlice[i].Date > *restoreJobsSlice[j].Date
		})

		var status string
		for _, restoreJob := range restoreJobsSlice {
			if *restoreJob.BackupID == backupId {
				status = *restoreJob.Status
				break
			}
		}

		switch status {
		default:
			return true, s, fmt.Errorf("restore job for backup with id %s has unexpected status %s", backupId, status)
		case RestoreJobProcessing:
			return false, nil, nil
		case RestoreJobFinished:
			return true, s, nil
		case RestoreJobBroken:
			return true, s, fmt.Errorf("restore job for backup with id %s is broken", backupId)
		case RestoreJobKilled:
			return true, s, fmt.Errorf("restore job for backup with id %s was killed", backupId)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	handler.SetSleepBeforeWait(5 * time.Second)
	return handler
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	waitConfig := wait.WaiterHelper[mongodbflex.InstanceResponse, string]{
		FetchInstance: a.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState:      getStateInstance,
		ActiveState:   []string{INSTANCESTATUS_READY},
		ErrorState:    []string{INSTANCESTATUS_FAILED},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	return UpdateInstanceWaitHandler(ctx, a, projectId, instanceId, region)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[struct{}] {
	w := wait.WaiterHelper[mongodbflex.InstanceResponse, string]{
		FetchInstance: a.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState:      getStateInstance,
		ActiveState:   []string{},
		ErrorState:    []string{},
	}

	// adapter for adhering to the wait helper type schema
	genericCheck := w.Wait()
	adaptedCheck := func() (waitFinished bool, response *struct{}, err error) {
		finished, _, err := genericCheck()
		if err != nil {
			return finished, nil, err
		}
		if finished {
			return true, nil, nil
		}
		return false, nil, nil
	}

	handler := wait.New(adaptedCheck)
	handler.SetTimeout(15 * time.Minute)
	return handler
}

func getStateInstance(response *mongodbflex.InstanceResponse) (string, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	if response.Item == nil {
		return "", errors.New("emtpy items")
	}
	if response.Item.Id == nil {
		return "", errors.New("emtpy item id")
	}
	if response.Item.Status == nil {
		return "", errors.New("emtpy item status")
	}
	return *response.Item.Status, nil
}
