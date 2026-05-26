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
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	VERSIONSTATE_ACTIVE = kms.VERSIONSTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	VERSIONSTATE_CREATING = kms.VERSIONSTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	VERSIONSTATE_KEY_MATERIAL_INVALID = kms.VERSIONSTATE_KEY_MATERIAL_INVALID
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE = kms.VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	VERSIONSTATE_DISABLED = kms.VERSIONSTATE_DISABLED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	VERSIONSTATE_DESTROYED = kms.VERSIONSTATE_DESTROYED

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYRINGSTATE_CREATING = kms.KEYRINGSTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYRINGSTATE_ACTIVE = kms.KEYRINGSTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYRINGSTATE_DELETED = kms.KEYRINGSTATE_DELETED

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	WRAPPINGKEYSTATE_ACTIVE = kms.WRAPPINGKEYSTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	WRAPPINGKEYSTATE_CREATING = kms.WRAPPINGKEYSTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	WRAPPINGKEYSTATE_EXPIRED = kms.WRAPPINGKEYSTATE_EXPIRED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	WRAPPINGKEYSTATE_DELETED = kms.WRAPPINGKEYSTATE_DELETED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	WRAPPINGKEYSTATE_KEY_MATERIAL_UNAVAILABLE = kms.WRAPPINGKEYSTATE_KEY_MATERIAL_UNAVAILABLE

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYSTATE_ACTIVE = kms.KEYSTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYSTATE_DELETED = kms.KEYSTATE_DELETED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYSTATE_NOT_AVAILABLE = kms.KEYSTATE_NOT_AVAILABLE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYSTATE_ERRORS_EXIST = kms.KEYSTATE_ERRORS_EXIST
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYSTATE_CREATING = kms.KEYSTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	KEYSTATE_NO_VERSION = kms.KEYSTATE_NO_VERSION
)

func CreateKeyRingWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region string, keyRingId string) *wait.AsyncActionHandler[kms.KeyRing] {
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

func CreateOrUpdateKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region string, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
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

func DeleteKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region string, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
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

func EnableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region string, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
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

func DisableKeyVersionWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region string, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
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

func CreateWrappingKeyWaitHandler(ctx context.Context, client kms.DefaultAPI, projectId string, region string, keyRingId, wrappingKeyId string) *wait.AsyncActionHandler[kms.WrappingKey] {
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
