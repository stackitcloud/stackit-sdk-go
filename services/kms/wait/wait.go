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
	StatusKeyActive   = "active"
	StatusKeyNotReady = "version_not_ready"
	StatusKeyDeleted  = "deleted"
)

const (
	StatusKeyVersionActive              = "active"
	StatusKeyVersionKeyMaterialNotReady = "key_material_not_ready"
	StatusKeyVersionKeyMaterialInvalid  = "key_material_invalid"
	StatusKeyVersionDisabled            = "disabled"
	StatusKeyVersionDestroyed           = "destroyed"
)

type ApiKmsClient interface {
	GetKeyExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string) (*kms.Key, error)
	GetVersionExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string, versionNumber int64) (*kms.Version, error)
}

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		response, err := client.GetKeyExecute(ctx, projectId, region, keyRingId, keyId)
		if err != nil {
			return false, nil, err
		}

		if response.State != nil {
			switch *response.State {
			case StatusKeyNotReady:
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

func DeleteKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		_, err := client.GetKeyExecute(ctx, projectId, region, keyRingId, keyId)
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

func EnableKeyVersionWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	handler := wait.New(func() (bool, *kms.Version, error) {
		response, err := client.GetVersionExecute(ctx, projectId, region, keyRingId, keyId, version)
		if err != nil {
			return false, nil, err
		}

		if response.State != nil {
			switch *response.State {
			case StatusKeyVersionDestroyed:
				return true, response, nil
			case StatusKeyVersionKeyMaterialInvalid:
				return true, response, nil
			case StatusKeyVersionDisabled:
				return true, response, nil
			case StatusKeyVersionKeyMaterialNotReady:
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

func DisableKeyVersionWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	handler := wait.New(func() (bool, *kms.Version, error) {
		_, err := client.GetVersionExecute(ctx, projectId, region, keyRingId, keyId, version)
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
