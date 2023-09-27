package objectstorage

import (
	"context"
	"fmt"
	"net/http"

	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

// Interface needed for tests
type APIClientBucketInterface interface {
	GetBucketExecute(ctx context.Context, projectId string, bucketName string) (*GetBucketResponse, error)
}

// CreateBucketWaitHandler will wait for creation
func CreateBucketWaitHandler(ctx context.Context, a APIClientBucketInterface, projectId, bucketName string) *wait.Handler {
	waitHandler := wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetBucketExecute(ctx, projectId, bucketName)
		if err != nil {
			return nil, false, err
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
		oapiErr, ok := err.(*oapiError.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return nil, false, err
		}
		return nil, true, nil
	})
}
