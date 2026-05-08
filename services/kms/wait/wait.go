package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/kms"
)

const (
	// Deprecated: StatusKeyActive is deprecated and will be removed after 14th November 2025. Use [kms.KEYSTATE_ACTIVE] instead.
	StatusKeyActive = "active"
	// Deprecated: StatusKeyNotReady is deprecated and will be removed after 14th November 2025. Use [kms.KEYSTATE_CREATING] instead.
	StatusKeyNotReady = "version_not_ready"
	// Deprecated: StatusKeyDeleted is deprecated and will be removed after 14th November 2025. Use [kms.KEYSTATE_DELETED] instead.
	StatusKeyDeleted = "deleted"
)

const (
	// Deprecated: StatusKeyVersionActive is deprecated and will be removed after 14th November 2025. Use [kms.VERSIONSTATE_ACTIVE] instead.
	StatusKeyVersionActive = "active"
	// Deprecated: StatusKeyVersionKeyMaterialNotReady is deprecated and will be removed after 14th November 2025. Use [kms.KEYSTATE_CREATING] instead.
	StatusKeyVersionKeyMaterialNotReady = "key_material_not_ready"
	// Deprecated: StatusKeyVersionKeyMaterialInvalid is deprecated and will be removed after 14th November 2025. Use [kms.VERSIONSTATE_KEY_MATERIAL_INVALID] instead.
	StatusKeyVersionKeyMaterialInvalid = "key_material_invalid"
	// Deprecated: StatusKeyVersionDisabled is deprecated and will be removed after 14th November 2025. Use [kms.VERSIONSTATE_DISABLED] instead.
	StatusKeyVersionDisabled = "disabled"
	// Deprecated: StatusKeyVersionDestroyed is deprecated and will be removed after 14th November 2025. Use [kms.VERSIONSTATE_DESTROYED] instead.
	StatusKeyVersionDestroyed = "destroyed"
)

const (
	// Deprecated: StatusWrappingKeyActive is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_ACTIVE] instead.
	StatusWrappingKeyActive = "active"
	// Deprecated: StatusWrappingKeyKeyMaterialNotReady is deprecated and will be removed after 12th November 2025. Use [kms.VERSIONSTATE_CREATING] instead.
	StatusWrappingKeyKeyMaterialNotReady = "key_material_not_ready"
	// Deprecated: StatusWrappingKeyExpired is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_EXPIRED] instead.
	StatusWrappingKeyExpired = "expired"
	// Deprecated: StatusWrappingKeyDeleting is deprecated and will be removed after 12th November 2025. Use [kms.WRAPPINGKEYSTATE_CREATING] instead.
	StatusWrappingKeyDeleting = "deleting"
)

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
type ApiKmsClient interface {
	GetKeyExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string) (*kms.Key, error)
	GetKeyRingExecute(ctx context.Context, projectId string, regionId string, keyRingId string) (*kms.KeyRing, error)
	GetVersionExecute(ctx context.Context, projectId string, regionId string, keyRingId string, keyId string, versionNumber int64) (*kms.Version, error)
	GetWrappingKeyExecute(ctx context.Context, projectId string, regionId string, keyRingId string, wrappingKeyId string) (*kms.WrappingKey, error)
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateKeyRingWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId string) *wait.AsyncActionHandler[kms.KeyRing] {
	waitConfig := wait.WaiterHelper[kms.KeyRing, kms.KeyRingState]{
		FetchInstance: func() (*kms.KeyRing, error) {
			return client.GetKeyRingExecute(ctx, projectId, region, keyRingId)
		},
		GetState: func(d *kms.KeyRing) (kms.KeyRingState, error) {
			if d == nil || d.State == nil {
				return "", errors.New("keyring or state is nil")
			}
			return *d.State, nil
		},
		ActiveState: []kms.KeyRingState{kms.KEYRINGSTATE_ACTIVE, kms.KEYRINGSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateOrUpdateKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	waitConfig := wait.WaiterHelper[kms.Key, kms.KeyState]{
		FetchInstance: func() (*kms.Key, error) {
			return client.GetKeyExecute(ctx, projectId, region, keyRingId, keyId)
		},
		GetState: func(d *kms.Key) (kms.KeyState, error) {
			if d == nil || d.State == nil {
				return "", errors.New("key or state is nil")
			}
			return *d.State, nil
		},
		ActiveState: []kms.KeyState{kms.KEYSTATE_ACTIVE, kms.KEYSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func DeleteKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string) *wait.AsyncActionHandler[kms.Key] {
	waitConfig := wait.WaiterHelper[kms.Key, kms.KeyState]{
		FetchInstance: func() (*kms.Key, error) {
			return client.GetKeyExecute(ctx, projectId, region, keyRingId, keyId)
		},
		GetState: func(_ *kms.Key) (kms.KeyState, error) {
			return "", nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound, http.StatusGone},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func EnableKeyVersionWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	waitConfig := wait.WaiterHelper[kms.Version, kms.VersionState]{
		FetchInstance: func() (*kms.Version, error) {
			return client.GetVersionExecute(ctx, projectId, region, keyRingId, keyId, version)
		},
		GetState: func(d *kms.Version) (kms.VersionState, error) {
			if d == nil || d.State == nil {
				return "", errors.New("version or state is nil")
			}
			return *d.State, nil
		},
		ActiveState: []kms.VersionState{kms.VERSIONSTATE_ACTIVE},
		ErrorState:  []kms.VersionState{kms.VERSIONSTATE_DESTROYED, kms.VERSIONSTATE_KEY_MATERIAL_INVALID},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func DisableKeyVersionWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, keyId string, version int64) *wait.AsyncActionHandler[kms.Version] {
	waitConfig := wait.WaiterHelper[kms.Version, kms.VersionState]{
		FetchInstance: func() (*kms.Version, error) {
			return client.GetVersionExecute(ctx, projectId, region, keyRingId, keyId, version)
		},
		GetState: func(d *kms.Version) (kms.VersionState, error) {
			if d == nil || d.State == nil {
				return "", errors.New("version or state is nil")
			}
			return *d.State, nil
		},
		ActiveState: []kms.VersionState{kms.VERSIONSTATE_DISABLED},
		ErrorState:  []kms.VersionState{kms.VERSIONSTATE_DESTROYED, kms.VERSIONSTATE_KEY_MATERIAL_INVALID},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateWrappingKeyWaitHandler(ctx context.Context, client ApiKmsClient, projectId, region, keyRingId, wrappingKeyId string) *wait.AsyncActionHandler[kms.WrappingKey] {
	waitConfig := wait.WaiterHelper[kms.WrappingKey, kms.WrappingKeyState]{
		FetchInstance: func() (*kms.WrappingKey, error) {
			return client.GetWrappingKeyExecute(ctx, projectId, region, keyRingId, wrappingKeyId)
		},
		GetState: func(d *kms.WrappingKey) (kms.WrappingKeyState, error) {
			if d == nil || d.State == nil {
				return "", errors.New("wrappingkey or state is nil")
			}
			return *d.State, nil
		},
		ActiveState: []kms.WrappingKeyState{kms.WRAPPINGKEYSTATE_ACTIVE, kms.WRAPPINGKEYSTATE_DELETED},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
