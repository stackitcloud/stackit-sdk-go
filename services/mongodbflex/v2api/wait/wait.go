package wait

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	mongodbflex "github.com/stackitcloud/stackit-sdk-go/services/mongodbflex/v2api"
)

const (
	INSTANCESTATUS_READY      = "READY"
	INSTANCESTATUS_PENDING    = "PENDING"
	INSTANCESTATUS_PROCESSING = "PROCESSING"
	INSTANCESTATUS_FAILED     = "FAILED"
	INSTANCESTATUS_UNKNOWN    = "UNKNOWN"

	RestoreJobProcessing = "IN_PROGRESS"
	RestoreJobFinished   = "FINISHED"
	RestoreJobBroken     = "BROKEN"
	RestoreJobKilled     = "KILLED"
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *mongodbflex.InstanceResponse, err error) {
		s, err := a.GetInstance(ctx, projectId, instanceId, region).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, s, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case "":
			return false, nil, nil
		case INSTANCESTATUS_PROCESSING:
			return false, nil, nil
		case INSTANCESTATUS_UNKNOWN:
			return false, nil, nil
		case INSTANCESTATUS_READY:
			return true, s, nil
		case INSTANCESTATUS_FAILED:
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
	})
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
	handler := wait.New(func() (waitFinished bool, response *mongodbflex.InstanceResponse, err error) {
		s, err := a.GetInstance(ctx, projectId, instanceId, region).Execute()
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, s, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case "":
			return false, nil, nil
		case INSTANCESTATUS_PROCESSING:
			return false, nil, nil
		case INSTANCESTATUS_UNKNOWN:
			return false, nil, nil
		case INSTANCESTATUS_READY:
			return true, s, nil
		case INSTANCESTATUS_FAILED:
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[mongodbflex.InstanceResponse] {
	return UpdateInstanceWaitHandler(ctx, a, projectId, instanceId, region)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a mongodbflex.DefaultAPI, projectId, instanceId, region string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetInstance(ctx, projectId, instanceId, region).Execute()
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}
