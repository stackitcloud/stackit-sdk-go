package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sfs "github.com/stackitcloud/stackit-sdk-go/services/sfs/v1api"
)

const (
	ResourcePoolStatePending  = "pending"
	ResourcePoolStateCreating = "creating"
	ResourcePoolStateCreated  = "created"
	ResourcePoolStateUpdating = "updating"
	ResourcePoolStateError    = "error"
	ResourcePoolStateDeleting = "deleting"
)

const (
	ShareStatePending  = "pending"
	ShareStateCreating = "creating"
	ShareStateCreated  = "created"
	ShareStateError    = "error"
	ShareStateDeleting = "deleting"
)

func CreateResourcePoolWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId string) *wait.AsyncActionHandler[sfs.GetResourcePoolResponse] {
	waitConfig := wait.WaiterHelper[sfs.GetResourcePoolResponse, string]{
		FetchInstance: api.GetResourcePool(ctx, projectId, region, resourcePoolId).Execute,
		GetState:      GetStateResourcePool,
		ActiveState:   []string{ResourcePoolStateCreated},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateResourcePoolWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId string) *wait.AsyncActionHandler[sfs.GetResourcePoolResponse] {
	waitConfig := wait.WaiterHelper[sfs.GetResourcePoolResponse, string]{
		FetchInstance: api.GetResourcePool(ctx, projectId, region, resourcePoolId).Execute,
		GetState:      GetStateResourcePool,
		ActiveState:   []string{ResourcePoolStateCreated},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteResourcePoolWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId string) *wait.AsyncActionHandler[sfs.GetResourcePoolResponse] {
	waitConfig := wait.WaiterHelper[sfs.GetResourcePoolResponse, string]{
		FetchInstance: api.GetResourcePool(ctx, projectId, region, resourcePoolId).Execute,
		GetState:      GetStateResourcePool,
		ActiveState:   []string{},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func GetStateResourcePool(response *sfs.GetResourcePoolResponse) (string, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	if response.ResourcePool == nil {
		return "", errors.New("resource pool is nil")
	}
	if response.ResourcePool.State == nil {
		return "", errors.New("resource pool state is nil")
	}
	return *response.ResourcePool.State, nil
}

func CreateShareWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	handler := wait.New(func() (waitFinished bool, share *sfs.GetShareResponse, err error) {
		share, err = api.GetShare(ctx, projectId, region, resourcePoolId, shareId).Execute()
		if err != nil {
			return false, share, err
		}
		if share == nil ||
			share.Share == nil ||
			share.Share.Id == nil ||
			share.Share.State == nil {
			return false, share, fmt.Errorf("create failed for  share with id %s %s, the response is not valid (state missing)", resourcePoolId, shareId)
		}
		if *share.Share.Id == shareId {
			switch *share.Share.State {
			case ShareStateCreated:
				return true, share, err
			default:
				return false, share, err
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateShareWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	handler := wait.New(func() (waitFinished bool, share *sfs.GetShareResponse, err error) {
		share, err = api.GetShare(ctx, projectId, region, resourcePoolId, shareId).Execute()
		if err != nil {
			return false, share, err
		}
		if share == nil ||
			share.Share == nil ||
			share.Share.Id == nil ||
			share.Share.State == nil {
			return false, share, fmt.Errorf("update failed for resourcepool with id %s, the response is not valid (state missing)", resourcePoolId)
		}
		if *share.Share.Id == shareId {
			switch *share.Share.State {
			case ResourcePoolStateCreated:
				return true, share, err
			default:
				return false, share, err
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteShareWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	handler := wait.New(func() (waitFinished bool, share *sfs.GetShareResponse, err error) {
		share, err = api.GetShare(ctx, projectId, region, resourcePoolId, shareId).Execute()
		if err != nil {
			var oapiError *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiError) {
				if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, share, nil
				}
			}
		}
		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}
