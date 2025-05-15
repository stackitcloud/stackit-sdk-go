package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/opensearch"
)

// Used for testing instance operations
type apiClientInstanceMocked struct {
	getFails                   bool
	deletionSucceedsWithErrors bool
	resourceId                 string
	resourceOperation          *opensearch.InstanceLastOperationTypes
	resourceState              opensearch.InstanceStatus
	resourceDescription        string
}

func (a *apiClientInstanceMocked) GetInstanceExecute(_ context.Context, _, _ string) (*opensearch.Instance, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}
	if a.resourceOperation != nil && *a.resourceOperation == opensearch.INSTANCELASTOPERATIONTYPE_DELETE && a.resourceState == opensearch.INSTANCESTATUS_ACTIVE {
		if a.deletionSucceedsWithErrors {
			return &opensearch.Instance{
				InstanceId: &a.resourceId,
				Status:     &a.resourceState,
				LastOperation: &opensearch.InstanceLastOperation{
					Description: &a.resourceDescription,
					Type:        a.resourceOperation,
				},
			}, nil
		}
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 410,
		}
	}

	return &opensearch.Instance{
		InstanceId: &a.resourceId,
		Status:     utils.Ptr(a.resourceState),
	}, nil
}

// Used for testing credentials operations
type apiClientCredentialsMocked struct {
	getFails          bool
	resourceId        string
	operationSucceeds bool
	deletionSucceeds  bool
}

func (a *apiClientCredentialsMocked) GetCredentialsExecute(_ context.Context, _, _, _ string) (*opensearch.CredentialsResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if !a.operationSucceeds || a.deletionSucceeds {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &opensearch.CredentialsResponse{
		Id: &a.resourceId,
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState opensearch.InstanceStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: opensearch.INSTANCESTATUS_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: opensearch.INSTANCESTATUS_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				getFails:      tt.getFails,
				resourceId:    instanceId,
				resourceState: tt.resourceState,
			}

			var wantRes *opensearch.Instance
			if tt.wantResp {
				wantRes = &opensearch.Instance{
					InstanceId: &instanceId,
					Status:     utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "pid", instanceId)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			diff := cmp.Diff(gotRes, wantRes)
			if diff != "" {
				t.Fatalf("handler gotRes = %+v\n want %+v\n diff = %s", gotRes, wantRes, diff)
			}
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState opensearch.InstanceStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: opensearch.INSTANCESTATUS_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: opensearch.INSTANCESTATUS_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				getFails:      tt.getFails,
				resourceId:    instanceId,
				resourceState: tt.resourceState,
			}

			var wantRes *opensearch.Instance
			if tt.wantResp {
				wantRes = &opensearch.Instance{
					InstanceId: &instanceId,
					Status:     utils.Ptr(tt.resourceState),
				}
			}

			handler := PartialUpdateInstanceWaitHandler(context.Background(), apiClient, "", instanceId)

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

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                      string
		getFails                  bool
		deleteSucceeedsWithErrors bool
		resourceState             opensearch.InstanceStatus
		resourceDescription       string
		wantErr                   bool
	}{
		{
			desc:                      "delete_succeeded",
			getFails:                  false,
			deleteSucceeedsWithErrors: false,
			resourceState:             opensearch.INSTANCESTATUS_ACTIVE,
			wantErr:                   false,
		},
		{
			desc:                      "delete_failed",
			getFails:                  false,
			deleteSucceeedsWithErrors: false,
			resourceState:             opensearch.INSTANCESTATUS_FAILED,
			wantErr:                   true,
		},
		{
			desc:                      "delete_succeeds_with_errors",
			getFails:                  false,
			resourceState:             opensearch.INSTANCESTATUS_ACTIVE,
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
				resourceOperation:          utils.Ptr(opensearch.INSTANCELASTOPERATIONTYPE_DELETE),
				resourceDescription:        tt.resourceDescription,
				resourceState:              tt.resourceState,
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", instanceId)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
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
		wantResp          bool
	}{
		{
			desc:              "create_succeeded",
			getFails:          false,
			operationSucceeds: true,
			wantErr:           false,
			wantResp:          true,
		},
		{
			desc:              "create_failed",
			getFails:          false,
			operationSucceeds: false,
			wantErr:           true,
			wantResp:          false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
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

			var wantRes *opensearch.CredentialsResponse
			if tt.wantResp {
				wantRes = &opensearch.CredentialsResponse{
					Id: &credentialsId,
				}
			}

			handler := CreateCredentialsWaitHandler(context.Background(), apiClient, "", "", credentialsId)

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

			handler := DeleteCredentialsWaitHandler(context.Background(), apiClient, "", "", credentialsId)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
