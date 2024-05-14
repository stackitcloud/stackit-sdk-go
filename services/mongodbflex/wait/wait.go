package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/mongodbflex"
)

const (
	InstanceStateEmpty      = ""
	InstanceStateProcessing = "PROCESSING"
	InstanceStateUnknown    = "UNKNOWN"
	InstanceStateSuccess    = "READY"
	InstanceStateFailed     = "FAILED"

	RestoreJobProcessing = "IN_PROGRESS"
	RestoreJobFinished   = "FINISHED"
	RestoreJobBroken     = "BROKEN"
	RestoreJobKilled     = "KILLED"
)

// Interface needed for tests
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, instanceId string) (*mongodbflex.GetInstanceResponse, error)
	ListRestoreJobsExecute(ctx context.Context, projectId, instanceId string) (*mongodbflex.ListRestoreJobsResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[mongodbflex.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *mongodbflex.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, s, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateEmpty:
			return false, nil, nil
		case InstanceStateProcessing:
			return false, nil, nil
		case InstanceStateUnknown:
			return false, nil, nil
		case InstanceStateSuccess:
			return true, s, nil
		case InstanceStateFailed:
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	handler.SetSleepBeforeWait(5 * time.Second)
	return handler
}

// CloneInstanceWaitHandler will wait for instance clone to be created
func CloneInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[mongodbflex.GetInstanceResponse] {
	return CreateInstanceWaitHandler(ctx, a, projectId, instanceId)
}

func RestoreInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId, backupId string) *wait.AsyncActionHandler[mongodbflex.ListRestoreJobsResponse] {
	handler := wait.New(func() (waitFinished bool, response *mongodbflex.ListRestoreJobsResponse, err error) {
		s, err := a.ListRestoreJobsExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Items == nil {
			return false, nil, nil
		}

		var status string
		for _, restoreJob := range *s.Items {
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
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[mongodbflex.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *mongodbflex.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Item == nil || s.Item.Id == nil || *s.Item.Id != instanceId || s.Item.Status == nil {
			return false, nil, nil
		}
		switch *s.Item.Status {
		default:
			return true, s, fmt.Errorf("instance with id %s has unexpected status %s", instanceId, *s.Item.Status)
		case InstanceStateEmpty:
			return false, nil, nil
		case InstanceStateProcessing:
			return false, nil, nil
		case InstanceStateUnknown:
			return false, nil, nil
		case InstanceStateSuccess:
			return true, s, nil
		case InstanceStateFailed:
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		}
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[mongodbflex.GetInstanceResponse] {
	return UpdateInstanceWaitHandler(ctx, a, projectId, instanceId)
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetInstanceExecute(ctx, projectId, instanceId)
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
