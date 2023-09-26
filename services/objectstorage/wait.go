package objectstorage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

// Interface needed for tests
type APIClientBucketInterface interface {
	GetBucketExecute(ctx context.Context, projectId string, bucketName string) (*GetBucketResponse, error)
}

func handleError(reqErr error) (res interface{}, done bool, err error) {
	oapiErr, ok := reqErr.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
	if !ok {
		return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
	}
	// Some APIs may return temporary errors and the request should be retried
	if utils.Contains(wait.RetryHttpErrorStatusCodes, oapiErr.statusCode) {
		return nil, false, nil
	}
	return nil, false, reqErr
}

// CreateBucketWaitHandler will wait for creation
func CreateBucketWaitHandler(ctx context.Context, a APIClientBucketInterface, projectId, bucketName string) *wait.Handler {
	waitHandler := wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetBucketExecute(ctx, projectId, bucketName)
		if err != nil {
			return handleError(err)
		}
		return s, true, nil
	})
	return waitHandler
}

// DeleteBucketWaitHandler will wait for delete
func DeleteBucketWaitHandler(ctx context.Context, a APIClientBucketInterface, projectId, bucketName string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetBucketExecute(ctx, projectId, bucketName)
		if err == nil {
			return s, false, nil
		}
		oapiErr, ok := err.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
		}
		if oapiErr.statusCode != http.StatusNotFound {
			return handleError(err)
		}
		return nil, true, nil
	})
}
