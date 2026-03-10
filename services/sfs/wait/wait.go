package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/sfs"
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

// Interfaces needed for tests
type APIClientInterface interface {
	GetResourcePoolExecute(ctx context.Context, projectId string, region string, resourcePoolId string) (*sfs.GetResourcePoolResponse, error)
	GetShareExecute(ctx context.Context, projectId string, region string, resourcePoolId string, shareId string) (*sfs.GetShareResponse, error)
}

func CreateResourcePoolWaitHandler(ctx context.Context, api APIClientInterface, projectId, region, resourcePoolId string) *wait.AsyncActionHandler[sfs.GetResourcePoolResponse] {
	handler := wait.New(func() (waitFinished bool, resourcePool *sfs.GetResourcePoolResponse, err error) {
		resourcePool, err = api.GetResourcePoolExecute(ctx, projectId, region, resourcePoolId)
		if err != nil {
			return false, resourcePool, err
		}
		if resourcePool == nil ||
			resourcePool.ResourcePool == nil ||
			resourcePool.ResourcePool.Id == nil ||
			resourcePool.ResourcePool.State == nil {
			return false, resourcePool, fmt.Errorf("create failed for resourcepool with id %s, the response is not valid (state missing)", resourcePoolId)
		}
		if *resourcePool.ResourcePool.Id == resourcePoolId {
			switch *resourcePool.ResourcePool.State {
			case ResourcePoolStateCreated:
				return true, resourcePool, err
			default:
				return false, resourcePool, err
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateResourcePoolWaitHandler(ctx context.Context, api APIClientInterface, projectId, region, resourcePoolId string) *wait.AsyncActionHandler[sfs.GetResourcePoolResponse] {
	handler := wait.New(func() (waitFinished bool, resourcePool *sfs.GetResourcePoolResponse, err error) {
		resourcePool, err = api.GetResourcePoolExecute(ctx, projectId, region, resourcePoolId)
		if err != nil {
			return false, resourcePool, err
		}
		if resourcePool == nil ||
			resourcePool.ResourcePool == nil ||
			resourcePool.ResourcePool.Id == nil ||
			resourcePool.ResourcePool.State == nil {
			return false, resourcePool, fmt.Errorf("update failed for resourcepool with id %s, the response is not valid (state missing)", resourcePoolId)
		}
		if *resourcePool.ResourcePool.Id == resourcePoolId {
			switch *resourcePool.ResourcePool.State {
			case ResourcePoolStateCreated:
				return true, resourcePool, err
			default:
				return false, resourcePool, err
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteResourcePoolWaitHandler(ctx context.Context, api APIClientInterface, projectId, region, resourcePoolId string) *wait.AsyncActionHandler[sfs.GetResourcePoolResponse] {
	handler := wait.New(func() (waitFinished bool, resourcePool *sfs.GetResourcePoolResponse, err error) {
		resourcePool, err = api.GetResourcePoolExecute(ctx, projectId, region, resourcePoolId)
		if err != nil {
			var oapiError *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiError) {
				if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, resourcePool, nil
				}
			}
		}
		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateShareWaitHandler(ctx context.Context, api APIClientInterface, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	handler := wait.New(func() (waitFinished bool, share *sfs.GetShareResponse, err error) {
		share, err = api.GetShareExecute(ctx, projectId, region, resourcePoolId, shareId)
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

func UpdateShareWaitHandler(ctx context.Context, api APIClientInterface, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	handler := wait.New(func() (waitFinished bool, share *sfs.GetShareResponse, err error) {
		share, err = api.GetShareExecute(ctx, projectId, region, resourcePoolId, shareId)
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

func DeleteShareWaitHandler(ctx context.Context, api APIClientInterface, projectId, region, resourcePoolId, shareId string) *wait.AsyncActionHandler[sfs.GetShareResponse] {
	handler := wait.New(func() (waitFinished bool, share *sfs.GetShareResponse, err error) {
		share, err = api.GetShareExecute(ctx, projectId, region, resourcePoolId, shareId)
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
