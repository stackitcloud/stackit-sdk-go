package wait

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/kms"
)

var (
	testProject   = uuid.NewString()
	testRegion    = "eu01"
	testDate      = time.Now()
	testKeyRingId = uuid.NewString()
	testKeyId     = uuid.NewString()
	testPublicKey = "i am an invalid public key"
)

var _ ApiKmsClient = (*apiKmsMocked)(nil)

type keyResponse struct {
	key *kms.Key
	err error
}

type versionResponse struct {
	version *kms.Version
	err     error
}

type apiKmsMocked struct {
	idxKeyResponse     int
	idxVersionResponse int
	keyResponses       []keyResponse
	versionResponses   []versionResponse
}

// GetVersionExecute implements ApiKmsClient.
func (a *apiKmsMocked) GetVersionExecute(_ context.Context, _, _, _, _ string, _ int64) (*kms.Version, error) {
	resp := a.versionResponses[a.idxVersionResponse]
	a.idxVersionResponse++
	a.idxVersionResponse %= len(a.versionResponses)
	return resp.version, resp.err
}

// GetKeyExecute implements ApiKmsClient.
func (a *apiKmsMocked) GetKeyExecute(_ context.Context, _, _, _, _ string) (*kms.Key, error) {
	resp := a.keyResponses[a.idxKeyResponse]
	a.idxKeyResponse++
	a.idxKeyResponse %= len(a.keyResponses)
	return resp.key, resp.err
}

func fixtureKey(state string) *kms.Key {
	return &kms.Key{
		Algorithm:    kms.ALGORITHM_AES_256_GCM.Ptr(),
		Backend:      kms.BACKEND_SOFTWARE.Ptr(),
		CreatedAt:    &testDate,
		DeletionDate: &testDate,
		Description:  utils.Ptr("test-description"),
		DisplayName:  utils.Ptr("test-displayname"),
		Id:           &testKeyId,
		ImportOnly:   utils.Ptr(false),
		KeyRingId:    &testKeyRingId,
		Purpose:      kms.PURPOSE_SYMMETRIC_ENCRYPT_DECRYPT.Ptr(),
		State:        &state,
	}
}

func fixtureVersion(version int, disabled bool, state string) *kms.Version {
	return &kms.Version{
		CreatedAt:   &testDate,
		DestroyDate: &testDate,
		Disabled:    &disabled,
		KeyId:       &testKeyId,
		KeyRingId:   &testKeyRingId,
		Number:      utils.Ptr(int64(version)),
		PublicKey:   &testPublicKey,
		State:       &state,
	}
}

func TestCreateOrUpdateKeyWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []keyResponse
		want      *kms.Key
		wantErr   bool
	}{
		{
			"create succeeded immediately",
			[]keyResponse{
				{fixtureKey(StatusKeyActive), nil},
			},
			fixtureKey(StatusKeyActive),
			false,
		},
		{
			"create succeeded delayed",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyActive), nil},
			},
			fixtureKey(StatusKeyActive),
			false,
		},
		{
			"create failed delayed",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyDeleted), nil},
			},
			fixtureKey(StatusKeyDeleted),
			false,
		},
		{
			"timeout",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
			},
			nil,
			true,
		},
		{
			"broken state",
			[]keyResponse{
				{fixtureKey("bogus"), nil},
			},
			fixtureKey("bogus"),
			false,
		},
		// no special update tests needed as the states are the same
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiKmsMocked{
				keyResponses: tt.responses,
			}

			handler := CreateOrUpdateKeyWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId)
			got, err := handler.SetTimeout(1 * time.Second).
				SetThrottle(250 * time.Millisecond).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing key %s", diff)
			}
		})
	}
}

func TestDeleteKeyWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []keyResponse
		wantErr   bool
	}{
		{
			"Delete with '404' succeeded immediately",
			[]keyResponse{
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with '404' delayed",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with 'gone' succeeded immediately",
			[]keyResponse{
				{nil, oapierror.NewError(http.StatusGone, "gone")},
			},
			false,
		},
		{
			"Delete with 'gone' delayed",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{nil, oapierror.NewError(http.StatusGone, "not found")},
			},
			false,
		},
		{
			"Delete with error delayed",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},

				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyDeleted), oapierror.NewError(http.StatusInternalServerError, "kapow")},
			},
			true,
		},
		{
			"Cannot delete",
			[]keyResponse{
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyNotReady), nil},
				{fixtureKey(StatusKeyDeleted), oapierror.NewError(http.StatusOK, "ok")},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiKmsMocked{
				keyResponses: tt.responses,
			}
			handler := DeleteKeyWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId)
			_, err := handler.SetTimeout(1 * time.Second).
				SetThrottle(250 * time.Millisecond).
				WaitWithContext(ctx)

			if tt.wantErr != (err != nil) {
				t.Fatalf("wrong error result. want err: %v got %v", tt.wantErr, err)
			}
			if tt.wantErr {
				var apiErr *oapierror.GenericOpenAPIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("expected openapi error, got %v", err)
				}
			}
		})
	}
}

func TestEnableKeyVersionWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []versionResponse
		want      *kms.Version
		wantErr   bool
	}{
		{
			"create succeeded immediately",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionActive), nil},
			},
			fixtureVersion(1, false, StatusKeyVersionActive),
			false,
		},
		{
			"create succeeded delayed",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionActive), nil},
			},
			fixtureVersion(1, false, StatusKeyVersionActive),
			false,
		},
		{
			"create failed delayed",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialInvalid), nil},
			},
			fixtureVersion(1, false, StatusKeyVersionKeyMaterialInvalid),
			false,
		},
		{
			"timeout",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
			},
			nil,
			true,
		},
		{
			"broken state",
			[]versionResponse{
				{fixtureVersion(1, false, "bogus"), nil},
			},
			fixtureVersion(1, false, "bogus"),
			false,
		},
		// no special update tests needed as the states are the same
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiKmsMocked{
				versionResponses: tt.responses,
			}

			handler := EnableKeyVersionWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId, 1)
			got, err := handler.SetTimeout(1 * time.Second).
				SetThrottle(250 * time.Millisecond).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing key %s", diff)
			}
		})
	}
}

func TestDisableKeyVersionWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []versionResponse
		wantErr   bool
	}{
		{
			"Delete with '404' succeeded immediately",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with '404' delayed",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with 'gone' succeeded immediately",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusGone, "gone")},
			},
			false,
		},
		{
			"Delete with 'gone' delayed",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{nil, oapierror.NewError(http.StatusGone, "not found")},
			},
			false,
		},
		{
			"Delete with error delayed",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},

				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionDestroyed), oapierror.NewError(http.StatusInternalServerError, "kapow")},
			},
			true,
		},
		{
			"Cannot delete",
			[]versionResponse{
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionKeyMaterialNotReady), nil},
				{fixtureVersion(1, false, StatusKeyVersionDestroyed), oapierror.NewError(http.StatusOK, "ok")},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiKmsMocked{
				versionResponses: tt.responses,
			}
			handler := DisableKeyVersionWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId, 1)
			_, err := handler.SetTimeout(1 * time.Second).
				SetThrottle(250 * time.Millisecond).
				WaitWithContext(ctx)

			if tt.wantErr != (err != nil) {
				t.Fatalf("wrong error result. want err: %v got %v", tt.wantErr, err)
			}
			if tt.wantErr {
				var apiErr *oapierror.GenericOpenAPIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("expected openapi error, got %v", err)
				}
			}
		})
	}
}
