package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	postgresflex "github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v3beta1api"
)

type mockSettings struct {
	instanceId        string
	instanceState     postgresflex.State
	instanceIsDeleted bool
	instanceGetFails  bool

	usersGetErrorStatus int

	userGetFails  bool
	userId        int32
	userIsDeleted bool
}

// Used for testing instance operations
func newAPIMock(settings *mockSettings) postgresflex.DefaultAPI {
	return &postgresflex.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ postgresflex.ApiGetInstanceRequest) (*postgresflex.GetInstanceResponse, error) {
			if settings.instanceGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.instanceIsDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &postgresflex.GetInstanceResponse{
				Id:    settings.instanceId,
				State: settings.instanceState,
			}, nil
		}),
		ListUsersExecuteMock: utils.Ptr(func(_ postgresflex.ApiListUsersRequest) (*postgresflex.ListUserResponse, error) {
			if settings.usersGetErrorStatus != 0 {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.usersGetErrorStatus,
				}
			}

			return &postgresflex.ListUserResponse{
				Users: []postgresflex.ListUser{},
			}, nil
		}),
		GetUserExecuteMock: utils.Ptr(func(_ postgresflex.ApiGetUserRequest) (*postgresflex.GetUserResponse, error) {
			if settings.userGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.userIsDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &postgresflex.GetUserResponse{
				Id: settings.userId,
			}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                string
		instanceGetFails    bool
		instanceState       postgresflex.State
		usersGetErrorStatus int
		wantErr             bool
		wantResp            bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_FAILURE,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceState:    InstanceStateEmpty,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "instance_get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:                "users_get_fails",
			instanceGetFails:    false,
			instanceState:       postgresflex.STATE_READY,
			usersGetErrorStatus: 500,
			wantErr:             true,
			wantResp:            false,
		},
		{
			desc:                "users_get_fails_2",
			instanceGetFails:    false,
			instanceState:       postgresflex.STATE_READY,
			usersGetErrorStatus: 400,
			wantErr:             true,
			wantResp:            true,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_PROGRESSING,
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := newAPIMock(&mockSettings{
					instanceId:          instanceId,
					instanceState:       tt.instanceState,
					instanceGetFails:    tt.instanceGetFails,
					usersGetErrorStatus: tt.usersGetErrorStatus,
				})

				var wantRes *postgresflex.GetInstanceResponse
				if tt.wantResp {
					wantRes = &postgresflex.GetInstanceResponse{
						Id:    instanceId,
						State: tt.instanceState,
					}
				}

				handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmp.AllowUnexported(postgresflex.NullableString{}, postgresflex.NullableBool{}, postgresflex.NullableInt32{}), cmpopts.EquateComparable(apiClient)) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    postgresflex.State
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "update_succeeded",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "update_failed",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_FAILURE,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "update_failed_2",
			instanceGetFails: false,
			instanceState:    InstanceStateEmpty,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_PROGRESSING,
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := newAPIMock(&mockSettings{
					instanceId:       instanceId,
					instanceState:    tt.instanceState,
					instanceGetFails: tt.instanceGetFails,
				})

				var wantRes *postgresflex.GetInstanceResponse
				if tt.wantResp {
					wantRes = &postgresflex.GetInstanceResponse{
						Id:    instanceId,
						State: tt.instanceState,
					}
				}

				handler := PartialUpdateInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmp.AllowUnexported(postgresflex.NullableString{}, postgresflex.NullableBool{}, postgresflex.NullableInt32{}), cmpopts.EquateComparable(apiClient)) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		isDeleted        bool
		instanceState    postgresflex.State
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			isDeleted:        true,
			instanceGetFails: false,
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_FAILURE,
			wantErr:          true,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := newAPIMock(&mockSettings{
					instanceGetFails:  tt.instanceGetFails,
					instanceIsDeleted: tt.isDeleted,
					instanceId:        instanceId,
					instanceState:     tt.instanceState,
				})

				handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}

func TestDeleteUserWaitHandler(t *testing.T) {
	tests := []struct {
		desc        string
		deleteFails bool
		getFails    bool
		wantErr     bool
	}{
		{
			desc:        "delete_succeeded",
			deleteFails: false,
			getFails:    false,
			wantErr:     false,
		},
		{
			desc:        "delete_failed",
			deleteFails: true,
			getFails:    false,
			wantErr:     true,
		},
		{
			desc:        "get_fails",
			deleteFails: false,
			getFails:    true,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				userId := int32(34)

				apiClient := newAPIMock(&mockSettings{
					userGetFails:  tt.getFails,
					userId:        userId,
					userIsDeleted: !tt.deleteFails,
				})

				handler := DeleteUserWaitHandler(context.Background(), apiClient, "", "", "", userId)

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
