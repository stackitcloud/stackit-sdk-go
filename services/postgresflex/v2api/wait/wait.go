package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	postgresflex "github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v2api"
)

const (
	InstanceStateEmpty       = ""
	InstanceStateProgressing = "Progressing"
	InstanceStateSuccess     = "Ready"
	InstanceStateFailed      = "Failure"
	InstanceStateDeleted     = "Deleted"
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.InstanceResponse] {
	instanceCreated := false
	var instanceGetResponse *postgresflex.InstanceResponse

	handler := wait.New(func() (waitFinished bool, response *postgresflex.InstanceResponse, err error) {
		if !instanceCreated {
			s, err := a.GetInstance(ctx, projectId, region, instanceId).Execute()
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
			case InstanceStateProgressing:
				return false, nil, nil
			case InstanceStateSuccess:
				instanceCreated = true
				instanceGetResponse = s
			case InstanceStateFailed:
				return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
			}
		}

		// User operations aren't available right after an instance is deemed successful
		// To check if they are, perform a users request
		_, err = a.ListUsers(ctx, projectId, region, instanceId).Execute()
		if err == nil {
			return true, instanceGetResponse, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode < 500 {
			return true, instanceGetResponse, fmt.Errorf("users request after instance creation returned %d status code", oapiErr.StatusCode)
		}
		return false, nil, nil
	})
	// Sleep before wait is set because sometimes API returns 404 right after creation request
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.InstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.InstanceResponse, string]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *postgresflex.InstanceResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Item == nil {
				return "", errors.New("empty instance")
			}
			if response.Item.Status == nil {
				return "", errors.New("status is missing")
			}
			return *response.Item.Status, nil
		},
		ActiveState: []string{InstanceStateSuccess},
		ErrorState:  []string{InstanceStateFailed},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.InstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.InstanceResponse, string]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState: func(response *postgresflex.InstanceResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Item.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Item.Status, nil
		},
		ActiveState: []string{InstanceStateDeleted},
		ErrorState:  []string{InstanceStateFailed},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// ForceDeleteInstanceWaitHandler will wait for instance deletion
func ForceDeleteInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.InstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.InstanceResponse, string]{
		FetchInstance: client.GetInstance(ctx, projectId, instanceId, region).Execute,
		GetState: func(response *postgresflex.InstanceResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Item.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *response.Item.Status, nil
		},
		ErrorState:                 []string{InstanceStateFailed},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteUserWaitHandler will wait for delete
func DeleteUserWaitHandler(ctx context.Context, a postgresflex.DefaultAPI, projectId, region, instanceId, userId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetUser(ctx, projectId, region, instanceId, userId).Execute()
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode != 404 {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}
