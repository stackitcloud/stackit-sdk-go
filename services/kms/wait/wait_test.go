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
	testName      = "testlb"
	testDate      = time.Now()
	testKeyRingId = uuid.NewString()
	testKeyId     = uuid.NewString()
)

var _ ApiKmsClient = (*apiKmsMocked)(nil)

type response struct {
	key *kms.Key
	err error
}

type apiKmsMocked struct {
	n         int
	responses []response
}

// GetKeyExecute implements ApiKmsClient.
func (a *apiKmsMocked) GetKeyExecute(ctx context.Context, projectId string, region string, keyRingId string, keyId string) (*kms.Key, error) {
	resp := a.responses[a.n]
	a.n++
	a.n %= len(a.responses)
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

func TestCreateOrUpdateKeyWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []response
		want      *kms.Key
		wantErr   bool
	}{
		{
			"create succeeded immediately",
			[]response{
				{fixtureKey(StatusKeyActive), nil},
			},
			fixtureKey(StatusKeyActive),
			false,
		},
		{
			"create succeeded delayed",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyActive), nil},
			},
			fixtureKey(StatusKeyActive),
			false,
		},
		{
			"create failed delayed",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyDeleted), nil},
			},
			fixtureKey(StatusKeyDeleted),
			false,
		},
		{
			"timeout",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
			},
			nil,
			true,
		},
		{
			"broken state",
			[]response{
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
				responses: tt.responses,
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
		responses []response
		wantErr   bool
	}{
		{
			"Delete with '404' succeeded immediately",
			[]response{
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with '404' delayed",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{nil, oapierror.NewError(http.StatusNotFound, "not found")},
			},
			false,
		},
		{
			"Delete with 'gone' succeeded immediately",
			[]response{
				{nil, oapierror.NewError(http.StatusGone, "gone")},
			},
			false,
		},
		{
			"Delete with 'gone' delayed",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{nil, oapierror.NewError(http.StatusGone, "not found")},
			},
			false,
		},
		{
			"Delete with error delayed",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},

				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyDeleted), oapierror.NewError(http.StatusInternalServerError, "kapow")},
			},
			true,
		},
		{
			"Cannot delete",
			[]response{
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyVersionNotReady), nil},
				{fixtureKey(StatusKeyDeleted), oapierror.NewError(http.StatusOK, "ok")},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := &apiKmsMocked{
				responses: tt.responses,
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
