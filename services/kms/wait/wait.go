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
	// Deprecated: StatusKeyActive is deprecated and will be removed after 12th November 2025. Use [kms.KEYSTATE_ACTIVE] instead.
	StatusKeyActive = "active"
	// Deprecated: StatusKeyNotReady is deprecated and will be removed after 12th November 2025. Use [kms.KEYSTATE_VERSION_NOT_READY] instead.
	StatusKeyNotReady = "version_not_ready"
	// Deprecated: StatusKeyDeleted is deprecated and will be removed after 12th November 2025. Use [kms.KEYSTATE_DELETED] instead.
	StatusKeyDeleted = "deleted"
)

const (
	// Deprecated: StatusKeyVersionActive is deprecated and will be removed after 12th November 2025. Use [kms.VERSIONSTATE_ACTIVE] instead.
	StatusKeyVersionActive = "active"
	// Deprecated: StatusKeyVersionKeyMaterialNotReady is deprecated and will be removed after 12th November 2025. Use [kms.VERSIONSTATE_KEY_MATERIAL_NOT_READY] instead.
	StatusKeyVersionKeyMaterialNotReady = "key_material_not_ready"
	// Deprecated: StatusKeyVersionKeyMaterialInvalid is deprecated and will be removed after 12th November 2025. Use [kms.VERSIONSTATE_KEY_MATERIAL_INVALID] instead.
	StatusKeyVersionKeyMaterialInvalid = "key_material_invalid"
	// Deprecated: StatusKeyVersionDisabled is deprecated and will be removed after 12th November 2025. Use [kms.VERSIONSTATE_DISABLED] instead.
	StatusKeyVersionDisabled = "disabled"
	// Deprecated: StatusKeyVersionDestroyed is deprecated and will be removed after 12th November 2025. Use [kms.VERSIONSTATE_DESTROYED] instead.
	StatusKeyVersionDestroyed = "destroyed"
)

const (
	// Deprecated: StatusWrappingKeyActive is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_ACTIVE] instead.
	StatusWrappingKeyActive = "active"
	// Deprecated: StatusWrappingKeyKeyMaterialNotReady is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_KEY_MATERIAL_NOT_READY] instead.
	StatusWrappingKeyKeyMaterialNotReady = "key_material_not_ready"
	// Deprecated: StatusWrappingKeyExpired is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_EXPIRED] instead.
	StatusWrappingKeyExpired = "expired"
	// Deprecated: StatusWrappingKeyDeleting is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_DELETING] instead.
	StatusWrappingKeyDeleting = "deleting"
)

type ApiKmsClient interface {
	GetKeyExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string) (*kms.Key, error)
	GetVersionExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string, versionNumber int64) (*kms.Version, error)
	GetWrappingKeyExecute(ctx context.Context, projectId string, regionId string, keyRingId string, wrappingKeyId string) (*kms.WrappingKey, error)
}

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		response, err := client.GetKeyExecute(ctx, projectId, region, keyRingId, keyId)
		if err != nil {
			return false, nil, err
		}

		if response.State != nil {
			switch *response.State {
			case kms.KEYSTATE_VERSION_NOT_READY:
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
			case kms.VERSIONSTATE_DESTROYED:
				return true, response, nil
			case kms.VERSIONSTATE_KEY_MATERIAL_INVALID:
				return true, response, nil
			case kms.VERSIONSTATE_DISABLED:
				return true, response, nil
			case kms.VERSIONSTATE_KEY_MATERIAL_NOT_READY:
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

func CreateWrappingKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, wrappingKeyId string) *wait.AsyncActionHandler[kms.WrappingKey] {
	handler := wait.New(func() (bool, *kms.WrappingKey, error) {
		response, err := client.GetWrappingKeyExecute(ctx, projectId, region, keyRingId, wrappingKeyId)
		if err != nil {
			return false, nil, err
		}

		if state := response.State; state != nil {
			switch *state {
			case kms.WRAPPINGKEYSTATE_KEY_MATERIAL_NOT_READY:
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
