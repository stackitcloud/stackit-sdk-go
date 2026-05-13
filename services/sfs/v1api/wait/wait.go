package wait

import (
	"context"
	"errors"
	"time"

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
		GetState:      getStateResourcePool,
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
		GetState:      getStateResourcePool,
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
		GetState:      getStateResourcePool,
		ActiveState:   []string{},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateShareWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	waitConfig := wait.WaiterHelper[sfs.GetShareResponse, string]{
		FetchInstance: api.GetShare(ctx, projectId, region, resourcePoolId, shareId).Execute,
		GetState:      getStateShare,
		ActiveState:   []string{ShareStateCreated},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateShareWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	waitConfig := wait.WaiterHelper[sfs.GetShareResponse, string]{
		FetchInstance: api.GetShare(ctx, projectId, region, resourcePoolId, shareId).Execute,
		GetState:      getStateShare,
		ActiveState:   []string{ShareStateCreated},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteShareWaitHandler(ctx context.Context, api sfs.DefaultAPI, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	waitConfig := wait.WaiterHelper[sfs.GetShareResponse, string]{
		FetchInstance: api.GetShare(ctx, projectId, region, resourcePoolId, shareId).Execute,
		GetState:      getStateShare,
		ActiveState:   []string{},
		ErrorState:    []string{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func getStateResourcePool(response *sfs.GetResourcePoolResponse) (string, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	if response.ResourcePool == nil {
		return "", errors.New("resource pool is nil")
	}
	if response.ResourcePool.Id == nil {
		return "", errors.New("resource pool id is nil")
	}
	if response.ResourcePool.State == nil {
		return "", errors.New("resource pool state is nil")
	}
	return *response.ResourcePool.State, nil
}

func getStateShare(response *sfs.GetShareResponse) (string, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	if response.Share == nil {
		return "", errors.New("share is nil")
	}
	if response.Share.Id == nil {
		return "", errors.New("share id is nil")
	}
	if response.Share.State == nil {
		return "", errors.New("share state is nil")
	}
	return *response.Share.State, nil
}
