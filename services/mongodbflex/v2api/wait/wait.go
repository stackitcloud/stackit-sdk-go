package wait

import (
	"context"
	"errors"
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
	waitConfig := wait.WaiterHelper[mongodbflex.InstanceResponse, mongodbflex.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState:      getStateInstance,
		ActiveState:   []mongodbflex.InstanceStatus{INSTANCESTATUS_READY},
		ErrorState:    []mongodbflex.InstanceStatus{INSTANCESTATUS_FAILED},
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
	waitConfig := wait.WaiterHelper[mongodbflex.ListRestoreJobsResponse, string]{
		FetchInstance: a.ListRestoreJobs(ctx, projectId, instanceId, region).Execute,
		GetState: func(response *mongodbflex.ListRestoreJobsResponse) (string, error) {
			if response == nil {
				return "", errors.New("response is nil")
			}
			if len(response.Items) == 0 {
				return "", errors.New("response items is empty")
			}
			restoreJobsSlice := response.Items
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
			return status, nil
		},
		ActiveState: []string{RestoreJobFinished},
		ErrorState:  []string{RestoreJobKilled, RestoreJobBroken},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(45 * time.Minute)
	handler.SetSleepBeforeWait(5 * time.Second)
	return handler
}

// UpdateInstanceWaitHandler will wait for instance update
func UpdateInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	waitConfig := wait.WaiterHelper[mongodbflex.InstanceResponse, mongodbflex.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState:      getStateInstance,
		ActiveState:   []mongodbflex.InstanceStatus{INSTANCESTATUS_READY},
		ErrorState:    []mongodbflex.InstanceStatus{INSTANCESTATUS_FAILED},
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
	w := wait.WaiterHelper[mongodbflex.InstanceResponse, mongodbflex.InstanceStatus]{
		FetchInstance: a.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState:      getStateInstance,
		ActiveState:   []mongodbflex.InstanceStatus{},
		ErrorState:    []mongodbflex.InstanceStatus{INSTANCESTATUS_FAILED},
	}

	// adapter for adhering to the wait helper type schema
	genericCheck := w.Wait()
	adaptedCheck := func() (waitFinished bool, response *struct{}, err error) {
		finished, _, err := genericCheck()
		return finished, nil, err
	}

	handler := wait.New(adaptedCheck)
	handler.SetTimeout(15 * time.Minute)
	return handler
}

func getStateInstance(response *mongodbflex.InstanceResponse) (mongodbflex.InstanceStatus, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	if response.Item == nil {
		return "", errors.New("empty items")
	}
	if response.Item.Id == nil {
		return "", errors.New("empty item id")
	}
	if response.Item.Status == nil {
		return "", errors.New("empty item status")
	}
	return *response.Item.Status, nil
}
