package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	redis "github.com/stackitcloud/stackit-sdk-go/services/redis/v1api"
)

type mockSettings struct {
	instanceGetFails                   bool
	instanceDeletionSucceedsWithErrors bool
	instanceResourceId                 string
	instanceResourceOperation          string
	instanceResourceState              *string
	instanceResourceDescription        string

	credentialGetFails          bool
	credentialResourceId        string
	credentialOperationSucceeds bool
	credentialDeletionSucceeds  bool
}

func newAPIMock(settings *mockSettings) redis.DefaultAPI {
	return &redis.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ redis.ApiGetInstanceRequest) (*redis.Instance, error) {
			if settings.instanceGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}
			if settings.instanceResourceOperation == INSTANCELASTOPERATIONTYPE_DELETE && settings.instanceResourceState != nil && *settings.instanceResourceState == INSTANCESTATUS_ACTIVE {
				if settings.instanceDeletionSucceedsWithErrors {
					return &redis.Instance{
						InstanceId: &settings.instanceResourceId,
						Status:     settings.instanceResourceState,
						LastOperation: redis.InstanceLastOperation{
							Description: settings.instanceResourceDescription,
							Type:        settings.instanceResourceOperation,
						},
					}, nil
				}
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 410,
				}
			}

			return &redis.Instance{
				InstanceId: &settings.instanceResourceId,
				Status:     settings.instanceResourceState,
			}, nil
		}),
		GetCredentialsExecuteMock: utils.Ptr(func(_ redis.ApiGetCredentialsRequest) (*redis.CredentialsResponse, error) {
			if settings.credentialGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if !settings.credentialOperationSucceeds || settings.credentialDeletionSucceeds {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &redis.CredentialsResponse{
				Id: settings.credentialResourceId,
			}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(INSTANCESTATUS_ACTIVE),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: utils.Ptr(INSTANCESTATUS_FAILED),
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
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := newAPIMock(&mockSettings{
				instanceGetFails:      tt.getFails,
				instanceResourceId:    instanceId,
				instanceResourceState: tt.resourceState,
			})

			var wantRes *redis.Instance
			if tt.wantResp {
				wantRes = &redis.Instance{
					InstanceId: &instanceId,
					Status:     tt.resourceState,
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
		resourceState *string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(INSTANCESTATUS_ACTIVE),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: utils.Ptr(INSTANCESTATUS_FAILED),
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
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := newAPIMock(&mockSettings{
				instanceGetFails:      tt.getFails,
				instanceResourceId:    instanceId,
				instanceResourceState: tt.resourceState,
			})

			var wantRes *redis.Instance
			if tt.wantResp {
				wantRes = &redis.Instance{
					InstanceId: &instanceId,
					Status:     tt.resourceState,
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
		resourceState             *string
		resourceDescription       string
		wantErr                   bool
	}{
		{
			desc:                      "delete_succeeded",
			getFails:                  false,
			deleteSucceeedsWithErrors: false,
			resourceState:             utils.Ptr(INSTANCESTATUS_ACTIVE),
			wantErr:                   false,
		},
		{
			desc:                      "delete_failed",
			getFails:                  false,
			deleteSucceeedsWithErrors: false,
			resourceState:             utils.Ptr(INSTANCESTATUS_FAILED),
			wantErr:                   true,
		},
		{
			desc:                      "delete_succeeds_with_errors",
			getFails:                  false,
			resourceState:             utils.Ptr(INSTANCESTATUS_ACTIVE),
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

			apiClient := newAPIMock(&mockSettings{
				instanceGetFails:                   tt.getFails,
				instanceDeletionSucceedsWithErrors: tt.deleteSucceeedsWithErrors,
				instanceResourceId:                 instanceId,
				instanceResourceOperation:          INSTANCELASTOPERATIONTYPE_DELETE,
				instanceResourceDescription:        tt.resourceDescription,
				instanceResourceState:              tt.resourceState,
			})

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

			apiClient := newAPIMock(&mockSettings{
				credentialGetFails:          tt.getFails,
				credentialResourceId:        credentialsId,
				credentialOperationSucceeds: tt.operationSucceeds,
			})

			var wantRes *redis.CredentialsResponse
			if tt.wantResp {
				wantRes = &redis.CredentialsResponse{
					Id: credentialsId,
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

			apiClient := newAPIMock(&mockSettings{
				credentialGetFails:          tt.getFails,
				credentialResourceId:        credentialsId,
				credentialOperationSucceeds: true,
				credentialDeletionSucceeds:  tt.deletionSucceeds,
			})

			handler := DeleteCredentialsWaitHandler(context.Background(), apiClient, "", "", credentialsId)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
