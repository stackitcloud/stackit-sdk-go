package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	postgresflex "github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v3beta1api"
)

// CreateInstanceWaitHandler will wait for instance creation
func CreateInstanceWaitHandler(ctx context.Context, a postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	instanceCreated := false
	var instanceGetResponse *postgresflex.GetInstanceResponse

	handler := wait.New(func() (waitFinished bool, response *postgresflex.GetInstanceResponse, err error) {
		if !instanceCreated {
			s, err := a.GetInstance(ctx, projectId, region, instanceId).Execute()
			if err != nil {
				return false, nil, err
			}
			if s == nil || s.Id == "" || s.Id != instanceId {
				return false, nil, nil
			}
			switch s.State {
			default:
				return false, nil, nil
			case postgresflex.STATE_READY:
				instanceCreated = true
				instanceGetResponse = s
			case postgresflex.STATE_FAILURE:
				return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
			}
		}

		// User operations aren't available right after an instance is deemed successful
		// To check if they are, perform a users request
		_, err = a.ListUsers(ctx, projectId, region, instanceId).Execute()
		if err == nil {
			return true, instanceGetResponse, nil
		}
		var oapiErr *oapierror.GenericOpenAPIError
		ok := errors.As(err, &oapiErr)
		if !ok {
			return false, nil, err
		}
		if oapiErr.StatusCode != http.StatusLocked {
			return true, instanceGetResponse, fmt.Errorf("users request after instance creation returned %d status code", oapiErr.StatusCode)
		}
		return false, nil, nil
	})
	// Sleep before wait is set because sometimes API returns 404 right after creation request
	handler.SetTimeout(45 * time.Minute).SetSleepBeforeWait(15 * time.Second)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
func PartialUpdateInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.GetInstanceResponse, postgresflex.State]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *postgresflex.GetInstanceResponse) (postgresflex.State, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.State == "" {
				return "", errors.New("state is missing")
			}
			return response.State, nil
		},
		ActiveState: []postgresflex.State{postgresflex.STATE_READY},
		ErrorState:  []postgresflex.State{postgresflex.STATE_FAILURE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for instance deletion
func DeleteInstanceWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.GetInstanceResponse, postgresflex.State]{
		FetchInstance: client.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *postgresflex.GetInstanceResponse) (postgresflex.State, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.State == "" {
				return "", errors.New("status is missing in response")
			}
			return response.State, nil
		},
		ErrorState:                 []postgresflex.State{postgresflex.STATE_FAILURE},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateUserWaitHandler will wait for user creation
func CreateUserWaitHandler(ctx context.Context, client postgresflex.DefaultAPI, projectId, region, instanceId string, userId int64) *wait.AsyncActionHandler[postgresflex.GetUserResponse] {
	waitConfig := wait.WaiterHelper[postgresflex.GetUserResponse, string]{
		FetchInstance: client.GetUser(ctx, projectId, region, instanceId, userId).Execute,
		GetState: func(response *postgresflex.GetUserResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState: []string{"PROCESSED"},
		ErrorState:  []string{},
		// The API does not have a dedicated failure state for this resource,
		// so we rely on the timeout for cases where it never becomes active.
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteUserWaitHandler will wait for user deletion
func DeleteUserWaitHandler(ctx context.Context, a postgresflex.DefaultAPI, projectId, region, instanceId string, userId int64) *wait.AsyncActionHandler[struct{}] {
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
