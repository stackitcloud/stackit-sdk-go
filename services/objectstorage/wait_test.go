package objectstorage

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Used for testing bucket operations
type apiClientBucketMocked struct {
	bucketIsDeleted bool
	bucketGetFails  bool
}

func (a *apiClientBucketMocked) GetBucketExecute(_ context.Context, _, _ string) (*GetBucketResponse, error) {
	if a.bucketGetFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	if a.bucketIsDeleted {
		return nil, &GenericOpenAPIError{
			statusCode: 404,
		}
	}

	return &GetBucketResponse{}, nil
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		desc     string
		reqErr   error
		wantRes  interface{}
		wantDone bool
		wantErr  bool
	}{
		{
			desc: "handle_oapi_error",
			reqErr: &GenericOpenAPIError{
				statusCode: 500,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  true,
		},
		{
			desc:     "not_generic_oapi_error",
			reqErr:   fmt.Errorf("some error"),
			wantRes:  nil,
			wantDone: false,
			wantErr:  true,
		},
		{
			desc: "bad_gateway_error",
			reqErr: &GenericOpenAPIError{
				statusCode: http.StatusBadGateway,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  false,
		},
		{
			desc: "gateway_timeout_error",
			reqErr: &GenericOpenAPIError{
				statusCode: http.StatusBadGateway,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotRes, gotDone, err := handleError(tt.reqErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotRes, tt.wantRes) {
				t.Errorf("handleError() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotDone != tt.wantDone {
				t.Errorf("handleError() gotDone = %v, want %v", gotDone, tt.wantDone)
			}
		})
	}
}

func TestCreateBucketWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		bucketGetFails bool
		wantErr        bool
	}{
		{
			desc:           "create_succeeded",
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
				bucketGetFails: tt.bucketGetFails,
			}

			var wantRes *GetBucketResponse
			if !tt.bucketGetFails {
				wantRes = &GetBucketResponse{}
			}

			handler := CreateBucketWaitHandler(context.Background(), apiClient, "", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes) {
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

			handler := DeleteBucketWaitHandler(context.Background(), apiClient, "", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, nil)
			}
		})
	}
}
