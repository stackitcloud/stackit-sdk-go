package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	objectstorage "github.com/stackitcloud/stackit-sdk-go/services/objectstorage/v2api"
)

type mockSettings struct {
	bucketIsDeleted bool
	bucketGetFails  bool
}

func newAPIMock(settings *mockSettings) objectstorage.DefaultAPI {
	return &objectstorage.DefaultAPIServiceMock{
		GetBucketExecuteMock: utils.Ptr(func(_ objectstorage.ApiGetBucketRequest) (*objectstorage.GetBucketResponse, error) {
			if settings.bucketGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.bucketIsDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &objectstorage.GetBucketResponse{}, nil
		}),
	}
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
			apiClient := newAPIMock(&mockSettings{
				bucketGetFails: tt.bucketGetFails,
			})

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
			apiClient := newAPIMock(&mockSettings{
				bucketIsDeleted: true,
				bucketGetFails:  tt.bucketGetFails,
			})

			handler := DeleteBucketWaitHandler(context.Background(), apiClient, "", "", "")

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
