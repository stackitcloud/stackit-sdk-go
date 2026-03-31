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
	kms "github.com/stackitcloud/stackit-sdk-go/services/kms/v1api"
)

const (
	testRegion    = "eu01"
	testPublicKey = "i am an invalid public key"
)

var (
	testProject   = uuid.NewString()
	testDate      = time.Now()
	testKeyRingId = uuid.NewString()
	testKeyId     = uuid.NewString()
)

type keyResponse struct {
	key *kms.Key
	err error
}

type keyRingResponse struct {
	keyRing *kms.KeyRing
	err     error
}

type versionResponse struct {
	version *kms.Version
	err     error
}

type wrappingKeyResponse struct {
	wrappingKey *kms.WrappingKey
	err         error
}

type mockSettings struct {
	idxKeyResponse         int
	idxKeyRingResponse     int
	idxVersionResponse     int
	idxWrappingKeyResponse int
	keyResponses           []keyResponse
	keyRingResponses       []keyRingResponse
	versionResponses       []versionResponse
	wrappingKeyResponses   []wrappingKeyResponse
}

func newAPIMock(settings *mockSettings) kms.DefaultAPI {
	return &kms.DefaultAPIServiceMock{
		GetWrappingKeyExecuteMock: utils.Ptr(func(_ kms.ApiGetWrappingKeyRequest) (*kms.WrappingKey, error) {
			resp := settings.wrappingKeyResponses[settings.idxWrappingKeyResponse]
			settings.idxWrappingKeyResponse++
			settings.idxWrappingKeyResponse %= len(settings.wrappingKeyResponses)
			return resp.wrappingKey, resp.err
		}),
		GetVersionExecuteMock: utils.Ptr(func(_ kms.ApiGetVersionRequest) (*kms.Version, error) {
			resp := settings.versionResponses[settings.idxVersionResponse]
			settings.idxVersionResponse++
			settings.idxVersionResponse %= len(settings.versionResponses)
			return resp.version, resp.err
		}),
		GetKeyExecuteMock: utils.Ptr(func(_ kms.ApiGetKeyRequest) (*kms.Key, error) {
			resp := settings.keyResponses[settings.idxKeyResponse]
			settings.idxKeyResponse++
			settings.idxKeyResponse %= len(settings.keyResponses)
			return resp.key, resp.err
		}),
		GetKeyRingExecuteMock: utils.Ptr(func(_ kms.ApiGetKeyRingRequest) (*kms.KeyRing, error) {
			resp := settings.keyRingResponses[settings.idxKeyRingResponse]
			settings.idxKeyRingResponse++
			settings.idxKeyRingResponse %= len(settings.keyRingResponses)
			return resp.keyRing, resp.err
		}),
	}
}

func fixtureKey(state string) *kms.Key {
	return &kms.Key{
		Algorithm:    kms.ALGORITHM_AES_256_GCM,
		CreatedAt:    testDate,
		DeletionDate: &testDate,
		Description:  utils.Ptr("test-description"),
		DisplayName:  "test-displayname",
		Id:           testKeyId,
		ImportOnly:   false,
		KeyRingId:    testKeyRingId,
		Purpose:      kms.PURPOSE_SYMMETRIC_ENCRYPT_DECRYPT,
		State:        state,
	}
}

func fixtureKeyRing(state string) *kms.KeyRing {
	return &kms.KeyRing{
		CreatedAt:   testDate,
		Description: utils.Ptr("test-description"),
		DisplayName: "test-displayname",
		Id:          testKeyRingId,
		State:       state,
	}
}

func fixtureWrappingKey(state string) *kms.WrappingKey {
	return &kms.WrappingKey{
		Algorithm:   kms.WRAPPINGALGORITHM_RSA_2048_OAEP_SHA256,
		CreatedAt:   testDate,
		Description: utils.Ptr("test-description"),
		DisplayName: "test-displayname",
		Id:          testKeyId,
		KeyRingId:   testKeyRingId,
		Purpose:     kms.WRAPPINGPURPOSE_WRAP_SYMMETRIC_KEY,
		State:       state,
		ExpiresAt:   testDate,
		PublicKey:   utils.Ptr(testPublicKey),
	}
}

func fixtureVersion(version int, disabled bool, state string) *kms.Version {
	return &kms.Version{
		CreatedAt:   testDate,
		DestroyDate: &testDate,
		Disabled:    disabled,
		KeyId:       testKeyId,
		KeyRingId:   testKeyRingId,
		Number:      int64(version),
		PublicKey:   utils.Ptr(testPublicKey),
		State:       state,
	}
}

func TestCreateKeyRingWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []keyRingResponse
		want      *kms.KeyRing
		wantErr   bool
	}{
		{
			name: "create succeeded immediately",
			responses: []keyRingResponse{
				{fixtureKeyRing(KEYRINGSTATE_ACTIVE), nil},
			},
			want:    fixtureKeyRing(KEYRINGSTATE_ACTIVE),
			wantErr: false,
		},
		{
			name: "create succeeded delayed",
			responses: []keyRingResponse{
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
				{fixtureKeyRing(KEYRINGSTATE_ACTIVE), nil},
			},
			want:    fixtureKeyRing(KEYRINGSTATE_ACTIVE),
			wantErr: false,
		},
		{
			name: "create failed delayed",
			responses: []keyRingResponse{
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
				{fixtureKeyRing(KEYRINGSTATE_DELETED), nil},
			},
			want:    fixtureKeyRing(KEYRINGSTATE_DELETED),
			wantErr: false,
		},
		{
			name: "timeout",
			responses: []keyRingResponse{
				{fixtureKeyRing(KEYRINGSTATE_CREATING), nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "broken state",
			responses: []keyRingResponse{
				{fixtureKeyRing("bogus"), nil},
			},
			want:    fixtureKeyRing("bogus"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := newAPIMock(&mockSettings{
				keyRingResponses: tt.responses,
			})

			handler := CreateKeyRingWaitHandler(ctx, client, testProject, testRegion, testKeyRingId)
			got, err := handler.SetTimeout(250 * time.Millisecond).
				SetThrottle(1).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing key ring %s", diff)
			}
		})
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
				{fixtureKey(KEYSTATE_ACTIVE), nil},
			},
			fixtureKey(KEYSTATE_ACTIVE),
			false,
		},
		{
			"create succeeded delayed",
			[]keyResponse{
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_ACTIVE), nil},
			},
			fixtureKey(KEYSTATE_ACTIVE),
			false,
		},
		{
			"create failed delayed",
			[]keyResponse{
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_DELETED), nil},
			},
			fixtureKey(KEYSTATE_DELETED),
			false,
		},
		{
			"timeout",
			[]keyResponse{
				{fixtureKey(KEYSTATE_CREATING), nil},
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
			client := newAPIMock(&mockSettings{
				keyResponses: tt.responses,
			})

			handler := CreateOrUpdateKeyWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId)
			got, err := handler.SetTimeout(250 * time.Millisecond).
				SetThrottle(1).
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
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
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
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{nil, oapierror.NewError(http.StatusGone, "not found")},
			},
			false,
		},
		{
			"Delete with error delayed",
			[]keyResponse{
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},

				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_DELETED), oapierror.NewError(http.StatusInternalServerError, "kapow")},
			},
			true,
		},
		{
			"Cannot delete",
			[]keyResponse{
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_CREATING), nil},
				{fixtureKey(KEYSTATE_DELETED), oapierror.NewError(http.StatusOK, "ok")},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := newAPIMock(&mockSettings{
				keyResponses: tt.responses,
			})

			handler := DeleteKeyWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId)
			_, err := handler.SetTimeout(250 * time.Millisecond).
				SetThrottle(1).
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
				{fixtureVersion(1, false, VERSIONSTATE_ACTIVE), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_ACTIVE),
			false,
		},
		{
			"create succeeded delayed",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_ACTIVE), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_ACTIVE),
			false,
		},
		{
			"create failed with invalid key material",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_KEY_MATERIAL_INVALID), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_KEY_MATERIAL_INVALID),
			true,
		},
		{
			"timeout",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
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
			true,
		},
		{
			"version destroyed",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_DESTROYED), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_DESTROYED),
			true,
		},
		{
			"version disabled - unexpected state",
			[]versionResponse{
				{fixtureVersion(1, true, VERSIONSTATE_DISABLED), nil},
			},
			fixtureVersion(1, true, VERSIONSTATE_DISABLED),
			true,
		},
		{
			"version key material unavailable - unexpected state",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_KEY_MATERIAL_UNAVAILABLE),
			true,
		},
		{
			"version not found (404)",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			nil,
			true,
		},
		{
			"version gone (410)",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusGone, "gone")},
			},
			nil,
			true,
		},
		{
			"error fetching version",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusInternalServerError, "internal error")},
			},
			nil,
			true,
		},
		// no special update tests needed as the states are the same
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := newAPIMock(&mockSettings{
				versionResponses: tt.responses,
			})

			handler := EnableKeyVersionWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId, 1)
			got, err := handler.SetTimeout(250 * time.Millisecond).
				SetThrottle(1).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing version %s", diff)
			}
		})
	}
}

func TestDisableKeyVersionWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []versionResponse
		want      *kms.Version
		wantErr   bool
	}{
		{
			"disable succeeded immediately",
			[]versionResponse{
				{fixtureVersion(1, true, VERSIONSTATE_DISABLED), nil},
			},
			fixtureVersion(1, true, VERSIONSTATE_DISABLED),
			false,
		},
		{
			"disable succeeded delayed",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_ACTIVE), nil},
				{fixtureVersion(1, false, VERSIONSTATE_ACTIVE), nil},
				{fixtureVersion(1, false, VERSIONSTATE_ACTIVE), nil},
				{fixtureVersion(1, true, VERSIONSTATE_DISABLED), nil},
			},
			fixtureVersion(1, true, VERSIONSTATE_DISABLED),
			false,
		},
		{
			"disable succeeded from creating state",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, false, VERSIONSTATE_CREATING), nil},
				{fixtureVersion(1, true, VERSIONSTATE_DISABLED), nil},
			},
			fixtureVersion(1, true, VERSIONSTATE_DISABLED),
			false,
		},
		{
			"version already destroyed",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_DESTROYED), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_DESTROYED),
			true,
		},
		{
			"version key material invalid",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_KEY_MATERIAL_INVALID), nil},
			},
			fixtureVersion(1, false, VERSIONSTATE_KEY_MATERIAL_INVALID),
			true,
		},
		{
			"timeout waiting for disabled state",
			[]versionResponse{
				{fixtureVersion(1, false, VERSIONSTATE_ACTIVE), nil},
			},
			nil,
			true,
		},
		{
			"version not found (404)",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			nil,
			true,
		},
		{
			"version gone (410)",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusGone, "gone")},
			},
			nil,
			true,
		},
		{
			"error fetching version",
			[]versionResponse{
				{nil, oapierror.NewError(http.StatusInternalServerError, "internal error")},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := newAPIMock(&mockSettings{
				versionResponses: tt.responses,
			})

			handler := DisableKeyVersionWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId, 1)
			got, err := handler.SetTimeout(250 * time.Millisecond).
				SetThrottle(1).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing version %s", diff)
			}
		})
	}
}

func TestCreateWrappingWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []wrappingKeyResponse
		want      *kms.WrappingKey
		wantErr   bool
	}{
		{
			"create succeeded immediately",
			[]wrappingKeyResponse{
				{fixtureWrappingKey(WRAPPINGKEYSTATE_ACTIVE), nil},
			},
			fixtureWrappingKey(WRAPPINGKEYSTATE_ACTIVE),
			false,
		},
		{
			"create succeeded delayed",
			[]wrappingKeyResponse{
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
				{fixtureWrappingKey(WRAPPINGKEYSTATE_ACTIVE), nil},
			},
			fixtureWrappingKey(WRAPPINGKEYSTATE_ACTIVE),
			false,
		},
		{
			"create failed delayed",
			[]wrappingKeyResponse{
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
				{fixtureWrappingKey(WRAPPINGKEYSTATE_DELETED), nil},
			},
			fixtureWrappingKey(WRAPPINGKEYSTATE_DELETED),
			false,
		},
		{
			"timeout",
			[]wrappingKeyResponse{
				{fixtureWrappingKey(WRAPPINGKEYSTATE_CREATING), nil},
			},
			nil,
			true,
		},
		{
			"broken state",
			[]wrappingKeyResponse{
				{fixtureWrappingKey("bogus"), nil},
			},
			fixtureWrappingKey("bogus"),
			false,
		},
		// no special update tests needed as the states are the same
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := newAPIMock(&mockSettings{
				wrappingKeyResponses: tt.responses,
			})

			handler := CreateWrappingKeyWaitHandler(ctx, client, testProject, testRegion, testKeyRingId, testKeyId)
			got, err := handler.SetTimeout(250 * time.Millisecond).
				SetThrottle(1).
				WaitWithContext(ctx)

			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differing wrapping key %s", diff)
			}
		})
	}
}
