package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/services/objectstorage"
)

// Used for testing bucket operations
type apiClientBucketMocked struct {
	bucketIsDeleted bool
	bucketGetFails  bool
}

func (a *apiClientBucketMocked) GetBucketExecute(_ context.Context, _, _, _ string) (*objectstorage.GetBucketResponse, error) {
	if a.bucketGetFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if a.bucketIsDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &objectstorage.GetBucketResponse{}, nil
}

func TestCreateBucketWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		bucketGetFails bool
		wantErr        bool
		wantResp       bool
	}{
		{
			desc:           "create_succeeded",
			bucketGetFails: false,
			wantErr:        false,
			wantResp:       true,
		},
		{
			desc:           "get_fails",
			bucketGetFails: true,
			wantErr:        true,
			wantResp:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientBucketMocked{
				bucketGetFails: tt.bucketGetFails,
			}

			var wantRes *objectstorage.GetBucketResponse
			if tt.wantResp {
				wantRes = &objectstorage.GetBucketResponse{}
			}

			handler := CreateBucketWaitHandler(context.Background(), apiClient, "", "", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteBucketWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		bucketGetFails bool
		wantErr        bool
	}{
		{
			desc:           "delete_succeeded",
			bucketGetFails: false,
			wantErr:        false,
		},
		{
			desc:           "get_fails",
			bucketGetFails: true,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientBucketMocked{
				bucketIsDeleted: true,
				bucketGetFails:  tt.bucketGetFails,
			}

			handler := DeleteBucketWaitHandler(context.Background(), apiClient, "", "", "")

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
