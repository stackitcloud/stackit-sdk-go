package wait

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/objectstorage"
)

// Interface needed for tests
type APIClientBucketInterface interface {
	GetBucketExecute(ctx context.Context, projectId string, bucketName string) (*objectstorage.GetBucketResponse, error)
}

// CreateBucketWaitHandler will wait for bucket creation
func CreateBucketWaitHandler(ctx context.Context, a APIClientBucketInterface, projectId, bucketName string) *wait.AsyncActionHandler[objectstorage.GetBucketResponse] {
	waitHandler := wait.New(func() (waitFinished bool, response *objectstorage.GetBucketResponse, err error) {
		s, err := a.GetBucketExecute(ctx, projectId, bucketName)
		if err != nil {
			return false, nil, err
		}
		return true, s, nil
	})
	return waitHandler
}

// DeleteBucketWaitHandler will wait for bucket deletion
func DeleteBucketWaitHandler(ctx context.Context, a APIClientBucketInterface, projectId, bucketName string) *wait.AsyncActionHandler[struct{}] {
	return wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetBucketExecute(ctx, projectId, bucketName)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to GenericOpenApiError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, nil, err
		}
		return true, nil, nil
	})
}
