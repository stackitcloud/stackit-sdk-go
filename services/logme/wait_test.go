package logme

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

// Used for testing instance operations
type apiClientInstanceMocked struct {
	getFails                   bool
	deletionSucceedsWithErrors bool
	resourceId                 string
	resourceOperation          *string
	resourceState              string
	resourceDescription        string
}

var (
	instanceTypeCreate = InstanceTypeCreate
	instanceTypeUpdate = InstanceTypeUpdate
	instanceTypeDelete = InstanceTypeDelete
)

func (a *apiClientInstanceMocked) GetInstanceExecute(_ context.Context, _, _ string) (*Instance, error) {
	if a.getFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}
	if *a.resourceOperation == InstanceTypeDelete && a.resourceState == InstanceStateSuccess {
		if a.deletionSucceedsWithErrors {
			return &Instance{
				InstanceId: &a.resourceId,
				LastOperation: &LastOperation{
					Description: &a.resourceDescription,
					Type:        a.resourceOperation,
					State:       &a.resourceState,
				},
			}, nil
		}
		return nil, &GenericOpenAPIError{
			statusCode: 410,
		}
	}

	return &Instance{
		InstanceId: &a.resourceId,
		LastOperation: &LastOperation{
			Description: &a.resourceDescription,
			Type:        a.resourceOperation,
			State:       &a.resourceState,
		},
	}, nil
}

// Used for testing credentials operations
type apiClientCredentialsMocked struct {
	getFails          bool
	resourceId        string
	operationSucceeds bool
	deletionSucceeds  bool
}

func (a *apiClientCredentialsMocked) GetCredentialsExecute(_ context.Context, _, _, _ string) (*CredentialsResponse, error) {
	if a.getFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	if !a.operationSucceeds || a.deletionSucceeds {
		return nil, &GenericOpenAPIError{
			statusCode: 404,
		}
	}

	return &CredentialsResponse{
		Id: &a.resourceId,
	}, nil
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

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: InstanceStateSuccess,
			wantErr:       false,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: InstanceStateFailed,
			wantErr:       true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				getFails:          tt.getFails,
				resourceId:        instanceId,
				resourceOperation: &instanceTypeCreate,
				resourceState:     tt.resourceState,
			}

			var wantRes *Instance
			if !tt.getFails {
				wantRes = &Instance{
					InstanceId: &instanceId,
					LastOperation: &LastOperation{
						Type:        &instanceTypeCreate,
						State:       &tt.resourceState,
						Description: utils.Ptr(""),
					},
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "pid", instanceId)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			diff := cmp.Diff(gotRes, wantRes)
			if wantRes != nil && diff != "" {
				t.Fatalf("handler gotRes = %+v\n want %+v\n diff = %s", gotRes, wantRes, diff)
			}
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: InstanceStateSuccess,
			wantErr:       false,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: InstanceStateFailed,
			wantErr:       true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				getFails:          tt.getFails,
				resourceId:        instanceId,
				resourceOperation: &instanceTypeUpdate,
				resourceState:     tt.resourceState,
			}

			var wantRes *Instance
			if !tt.getFails {
				wantRes = &Instance{
					InstanceId: &instanceId,
					LastOperation: &LastOperation{
						Type:        &instanceTypeUpdate,
						State:       &tt.resourceState,
						Description: utils.Ptr(""),
					},
				}
			}

			handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "", instanceId)

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

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                      string
		getFails                  bool
		deleteSucceeedsWithErrors bool
		resourceState             string
		resourceDescription       string
		wantErr                   bool
	}{
		{
			desc:                      "delete_succeeded",
			getFails:                  false,
			deleteSucceeedsWithErrors: false,
			resourceState:             InstanceStateSuccess,
			wantErr:                   false,
		},
		{
			desc:                      "delete_failed",
			getFails:                  false,
			deleteSucceeedsWithErrors: false,
			resourceState:             InstanceStateFailed,
			wantErr:                   true,
		},
		{
			desc:                      "delete_succeeds_with_errors",
			getFails:                  false,
			resourceState:             InstanceStateSuccess,
			deleteSucceeedsWithErrors: true,
			resourceDescription:       "Deleting resource: cf failed with error: DeleteFailed",
			wantErr:                   true,
		},
		{
			desc:                      "get_fails",
			deleteSucceeedsWithErrors: false,
			getFails:                  true,
			wantErr:                   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				getFails:                   tt.getFails,
				deletionSucceedsWithErrors: tt.deleteSucceeedsWithErrors,
				resourceId:                 instanceId,
				resourceOperation:          &instanceTypeDelete,
				resourceDescription:        tt.resourceDescription,
				resourceState:              tt.resourceState,
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", instanceId)

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

func TestCreateCredentialsWaitHandler(t *testing.T) {
	tests := []struct {
		desc              string
		getFails          bool
		operationSucceeds bool
		wantErr           bool
	}{
		{
			desc:              "create_succeeded",
			getFails:          false,
			operationSucceeds: true,
			wantErr:           false,
		},
		{
			desc:              "create_failed",
			getFails:          false,
			operationSucceeds: false,
			wantErr:           true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			credentialsId := "foo-bar"

			apiClient := &apiClientCredentialsMocked{
				getFails:          tt.getFails,
				resourceId:        credentialsId,
				operationSucceeds: tt.operationSucceeds,
			}

			var wantRes *CredentialsResponse
			if !tt.getFails && tt.operationSucceeds {
				wantRes = &CredentialsResponse{
					Id: &credentialsId,
				}
			} else if !tt.getFails && !tt.operationSucceeds {
				wantRes = nil
			}

			handler := CreateCredentialsWaitHandler(context.Background(), apiClient, "", "", credentialsId)

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

func TestDeleteCredentialsWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		getFails         bool
		deletionSucceeds bool
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			getFails:         false,
			deletionSucceeds: true,
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			getFails:         false,
			deletionSucceeds: false,
			wantErr:          true,
		},
		{
			desc:             "get_fails",
			getFails:         true,
			deletionSucceeds: false,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			credentialsId := "foo-bar"

			apiClient := &apiClientCredentialsMocked{
				getFails:          tt.getFails,
				resourceId:        credentialsId,
				operationSucceeds: true,
				deletionSucceeds:  tt.deletionSucceeds,
			}

			var wantRes *CredentialsResponse
			if !tt.getFails && !tt.deletionSucceeds {
				wantRes = &CredentialsResponse{
					Id: &credentialsId,
				}
			} else if !tt.getFails && tt.deletionSucceeds {
				wantRes = nil
			}

			handler := DeleteCredentialsWaitHandler(context.Background(), apiClient, "", "", credentialsId)

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
