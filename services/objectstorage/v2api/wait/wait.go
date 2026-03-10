package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	objectstorage "github.com/stackitcloud/stackit-sdk-go/services/objectstorage/v2api"
)

// CreateBucketWaitHandler will wait for bucket creation
func CreateBucketWaitHandler(ctx context.Context, a objectstorage.DefaultAPI, projectId, region, bucketName string) *wait.AsyncActionHandler[objectstorage.GetBucketResponse] {
	handler := wait.New(func() (waitFinished bool, response *objectstorage.GetBucketResponse, err error) {
		s, err := a.GetBucket(ctx, projectId, region, bucketName).Execute()
		if err != nil {
			return false, nil, err
		}
		return true, s, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteBucketWaitHandler will wait for bucket deletion
func DeleteBucketWaitHandler(ctx context.Context, a objectstorage.DefaultAPI, projectId, region, bucketName string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetBucket(ctx, projectId, region, bucketName).Execute()
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
	handler.SetTimeout(1 * time.Minute)
	return handler
}
