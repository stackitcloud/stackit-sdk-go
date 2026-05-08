package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

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
	waitConfig := wait.WaiterHelper[kms.KeyRing, string]{
		FetchInstance: client.GetKeyRing(ctx, projectId, region, keyRingId).Execute,
		GetState: func(d *kms.KeyRing) (string, error) {
			if d == nil || d.State == "" {
				return "", errors.New("keyring or state is nil")
			}
			return d.State, nil
		},
		ActiveState: []string{KEYRINGSTATE_ACTIVE, KEYRINGSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	waitConfig := wait.WaiterHelper[kms.Key, string]{
		FetchInstance: client.GetKey(ctx, projectId, region, keyRingId, keyId).Execute,
		GetState: func(d *kms.Key) (string, error) {
			if d == nil || d.State == "" {
				return "", errors.New("key or state is nil")
			}
			return d.State, nil
		},
		ActiveState: []string{KEYSTATE_ACTIVE, KEYSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	waitConfig := wait.WaiterHelper[kms.Key, string]{
		FetchInstance: client.GetKey(ctx, projectId, region, keyRingId, keyId).Execute,
		GetState: func(_ *kms.Key) (string, error) {
			return "", nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound, http.StatusGone},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func EnableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	waitConfig := wait.WaiterHelper[kms.Version, string]{
		FetchInstance: client.GetVersion(ctx, projectId, region, keyRingId, keyId, version).Execute,
		GetState: func(d *kms.Version) (string, error) {
			if d == nil || d.State == "" {
				return "", errors.New("version or state is nil")
			}
			return d.State, nil
		},
		ActiveState: []string{VERSIONSTATE_ACTIVE},
		ErrorState:  []string{VERSIONSTATE_DESTROYED, VERSIONSTATE_KEY_MATERIAL_INVALID, VERSIONSTATE_DISABLED, VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DisableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	waitConfig := wait.WaiterHelper[kms.Version, string]{
		FetchInstance: client.GetVersion(ctx, projectId, region, keyRingId, keyId, version).Execute,
		GetState: func(d *kms.Version) (string, error) {
			if d == nil || d.State == "" {
				return "", errors.New("version or state is nil")
			}
			return d.State, nil
		},
		ActiveState: []string{VERSIONSTATE_DISABLED},
		ErrorState:  []string{VERSIONSTATE_DESTROYED, VERSIONSTATE_KEY_MATERIAL_INVALID},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateWrappingKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId, region, keyRingId, wrappingKeyId string) *wait.AsyncActionHandler[kms.WrappingKey] {
	waitConfig := wait.WaiterHelper[kms.WrappingKey, string]{
		FetchInstance: client.GetWrappingKey(ctx, projectId, region, keyRingId, wrappingKeyId).Execute,
		GetState: func(d *kms.WrappingKey) (string, error) {
			if d == nil || d.State == "" {
				return "", errors.New("wrappingkey or state is nil")
			}
			return d.State, nil
		},
		ActiveState: []string{WRAPPINGKEYSTATE_ACTIVE, WRAPPINGKEYSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
