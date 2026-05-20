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

func CreateKeyRingWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region kms.ListKeyRingsRegionIdParameter, keyRingId string) *wait.AsyncActionHandler[kms.KeyRing] {
	waitConfig := wait.WaiterHelper[kms.KeyRing, kms.KeyRingState]{
		FetchInstance: client.GetKeyRing(ctx, projectId, region, keyRingId).Execute,
		GetState: func(d *kms.KeyRing) (kms.KeyRingState, error) {
			if d == nil {
				return "", errors.New("keyring is nil")
			}
			return d.State, nil
		},
		ActiveState: []kms.KeyRingState{kms.KEYRINGSTATE_ACTIVE, kms.KEYRINGSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region kms.ListKeyRingsRegionIdParameter, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	waitConfig := wait.WaiterHelper[kms.Key, kms.KeyState]{
		FetchInstance: client.GetKey(ctx, projectId, region, keyRingId, keyId).Execute,
		GetState: func(d *kms.Key) (kms.KeyState, error) {
			if d == nil {
				return "", errors.New("key is nil")
			}
			return d.State, nil
		},
		ActiveState: []kms.KeyState{
			kms.KEYSTATE_ACTIVE,
			kms.KEYSTATE_DELETED,
			kms.KEYSTATE_NO_VERSION,
			kms.KEYSTATE_ERRORS_EXIST,
			kms.KEYSTATE_NOT_AVAILABLE,
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region kms.ListKeyRingsRegionIdParameter, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	waitConfig := wait.WaiterHelper[kms.Key, kms.KeyState]{
		FetchInstance: client.GetKey(ctx, projectId, region, keyRingId, keyId).Execute,
		GetState: func(_ *kms.Key) (kms.KeyState, error) {
			return "", nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound, http.StatusGone},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func EnableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region kms.ListKeyRingsRegionIdParameter, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	waitConfig := wait.WaiterHelper[kms.Version, kms.VersionState]{
		FetchInstance: client.GetVersion(ctx, projectId, region, keyRingId, keyId, version).Execute,
		GetState: func(d *kms.Version) (kms.VersionState, error) {
			if d == nil {
				return "", errors.New("version is nil")
			}
			return d.State, nil
		},
		ActiveState: []kms.VersionState{kms.VERSIONSTATE_ACTIVE},
		ErrorState: []kms.VersionState{
			kms.VERSIONSTATE_DESTROYED,
			kms.VERSIONSTATE_KEY_MATERIAL_INVALID,
			kms.VERSIONSTATE_DISABLED,
			kms.VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE,
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DisableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region kms.ListKeyRingsRegionIdParameter, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	waitConfig := wait.WaiterHelper[kms.Version, kms.VersionState]{
		FetchInstance: client.GetVersion(ctx, projectId, region, keyRingId, keyId, version).Execute,
		GetState: func(d *kms.Version) (kms.VersionState, error) {
			if d == nil {
				return "", errors.New("version is nil")
			}
			return d.State, nil
		},
		ActiveState: []kms.VersionState{kms.VERSIONSTATE_DISABLED},
		ErrorState: []kms.VersionState{
			kms.VERSIONSTATE_DESTROYED,
			kms.VERSIONSTATE_KEY_MATERIAL_INVALID,
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateWrappingKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region kms.ListKeyRingsRegionIdParameter, keyRingId, wrappingKeyId string) *wait.AsyncActionHandler[kms.WrappingKey] {
	waitConfig := wait.WaiterHelper[kms.WrappingKey, kms.WrappingKeyState]{
		FetchInstance: client.GetWrappingKey(ctx, projectId, region, keyRingId, wrappingKeyId).Execute,
		GetState: func(d *kms.WrappingKey) (kms.WrappingKeyState, error) {
			if d == nil {
				return "", errors.New("wrappingkey is nil")
			}
			return d.State, nil
		},
		ActiveState: []kms.WrappingKeyState{kms.WRAPPINGKEYSTATE_ACTIVE, kms.WRAPPINGKEYSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
