package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/kms"
)

const (
	StatusKeyActive          = "active"
	StatusKeyVersionNotReady = "version_not_ready"
	StatusKeyDeleted         = "deleted"
)

type ApiKmsClient interface {
	GetKeyExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string) (*kms.Key, error)
}

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		response, err := client.GetKeyExecute(ctx, projectId, region, keyRingId, keyId)
		if err != nil {
			return false, nil, err
		}

		if response.State != nil {
			switch *response.State {
			case StatusKeyVersionNotReady:
				return false, nil, nil
			default:
				return true, response, nil
			}
		}

		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId,keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		_, err := client.GetKeyExecute(ctx,projectId,region,keyRingId, keyId)
		if err != nil {
			var apiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &apiErr) {
				if statusCode := apiErr.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, nil, nil
				}
			}
			return true, nil, err
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
