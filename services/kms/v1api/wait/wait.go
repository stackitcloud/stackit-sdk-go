package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	kms "github.com/stackitcloud/stackit-sdk-go/services/kms/v1api"
)

const (
	VERSIONSTATE_ACTIVE                   = "active"
	VERSIONSTATE_CREATING                 = "creating"
	VERSIONSTATE_KEY_MATERIAL_INVALID     = "key_material_invalid"
	VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE = "key_material_unavailable"
	VERSIONSTATE_DISABLED                 = "disabled"
	VERSIONSTATE_DESTROYED                = "destroyed"

	KEYRINGSTATE_CREATING = "creating"
	KEYRINGSTATE_ACTIVE   = "active"
	KEYRINGSTATE_DELETED  = "deleted"

	WRAPPINGKEYSTATE_ACTIVE                   = "active"
	WRAPPINGKEYSTATE_CREATING                 = "creating"
	WRAPPINGKEYSTATE_EXPIRED                  = "expired"
	WRAPPINGKEYSTATE_DELETED                  = "deleted"
	WRAPPINGKEYSTATE_KEY_MATERIAL_UNAVAILABLE = "key_material_unavailable"

	KEYSTATE_ACTIVE        = "active"
	KEYSTATE_DELETED       = "deleted"
	KEYSTATE_NOT_AVAILABLE = "not_available"
	KEYSTATE_ERRORS_EXIST  = "errors_exist"
	KEYSTATE_CREATING      = "creating"
	KEYSTATE_NO_VERSION    = "no_version"
)

func CreateKeyRingWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId string) *wait.AsyncActionHandler[kms.KeyRing] {
	handler := wait.New(func() (bool, *kms.KeyRing, error) {
		response, err := client.GetKeyRing(ctx, projectId, region, keyRingId).Execute()
		if err != nil {
			return false, nil, err
		}

		if response != nil {
			switch response.State {
			case KEYRINGSTATE_CREATING:
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

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		response, err := client.GetKey(ctx, projectId, region, keyRingId, keyId).Execute()
		if err != nil {
			return false, nil, err
		}

		if response != nil {
			switch response.State {
			case KEYSTATE_CREATING:
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

func DeleteKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	handler := wait.New(func() (bool, *kms.Key, error) {
		_, err := client.GetKey(ctx, projectId, region, keyRingId, keyId).Execute()
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

func EnableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	handler := wait.New(func() (bool, *kms.Version, error) {
		response, err := client.GetVersion(ctx, projectId, region, keyRingId, keyId, version).Execute()
		if err != nil {
			var apiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &apiErr) {
				if statusCode := apiErr.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, nil, fmt.Errorf("enabling failed for key %s version %d: version or key not found", keyId, version)
				}
			}
			return false, nil, err
		}

		if response != nil {
			switch response.State {
			case VERSIONSTATE_ACTIVE:
				return true, response, nil
			case VERSIONSTATE_CREATING:
				return false, nil, nil
			case VERSIONSTATE_DESTROYED, VERSIONSTATE_KEY_MATERIAL_INVALID:
				return true, response, fmt.Errorf("enabling failed for key %s version %d: state %s", keyId, version, response.State)
			default:
				return true, response, fmt.Errorf("key version %d for key %s has unexpected state %s", version, keyId, response.State)
			}
		}

		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DisableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	handler := wait.New(func() (bool, *kms.Version, error) {
		response, err := client.GetVersion(ctx, projectId, region, keyRingId, keyId, version).Execute()
		if err != nil {
			var apiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &apiErr) {
				if statusCode := apiErr.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, nil, fmt.Errorf("disabling failed for key %s version %d: version or key not found", keyId, version)
				}
			}
			return false, nil, err
		}

		if response != nil {
			switch response.State {
			case VERSIONSTATE_DISABLED:
				return true, response, nil
			case VERSIONSTATE_ACTIVE, VERSIONSTATE_CREATING, VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE:
				return false, nil, nil
			case VERSIONSTATE_DESTROYED, VERSIONSTATE_KEY_MATERIAL_INVALID:
				return true, response, fmt.Errorf("disabling failed for key %s version %d: state %s", keyId, version, response.State)
			default:
				return true, response, fmt.Errorf("key version %d for key %s has unexpected state %s", version, keyId, response.State)
			}
		}

		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateWrappingKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, wrappingKeyId string) *wait.AsyncActionHandler[kms.WrappingKey] {
	handler := wait.New(func() (bool, *kms.WrappingKey, error) {
		response, err := client.GetWrappingKey(ctx, projectId, region, keyRingId, wrappingKeyId).Execute()
		if err != nil {
			return false, nil, err
		}

		if response != nil {
			switch response.State {
			case WRAPPINGKEYSTATE_CREATING:
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
