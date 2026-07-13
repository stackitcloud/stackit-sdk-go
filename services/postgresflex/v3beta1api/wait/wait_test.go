package wait

import (
	"context"
	"fmt"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"

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
	userId        int64
	userIsDeleted bool
	userState     string
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
					StatusCode: 423,
				}
			}

			if settings.userIsDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &postgresflex.GetUserResponse{
				Id:    settings.userId,
				State: settings.userState,
			}, nil
		}),
	}
}

func TestCreateOrUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    postgresflex.State
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "create_or_update_succeeded",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_or_update_failed",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_FAILURE,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_or_update_failed_2",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_UNKNOWN,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "instance_get_fails",
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
		{
			desc:             "timeout_2",
			instanceGetFails: false,
			instanceState:    postgresflex.STATE_PENDING,
			wantErr:          true,
			wantResp:         false,
		},
	}

	handlers := map[string]func(context.Context, postgresflex.DefaultAPI, string, string, string) *wait.AsyncActionHandler[postgresflex.GetInstanceResponse]{
		"common logic": createOrUpdateInstanceWaitHandler,
		"create":       CreateInstanceWaitHandler,
		"update":       PartialUpdateInstanceWaitHandler,
		"clone":        CloneInstanceWaitHandler,
	}

	for handlerDesc, handlerFn := range handlers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s - %s", handlerDesc, tt.desc), func(t *testing.T) {
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

					handler := handlerFn(context.Background(), apiClient, "", "", instanceId)
					gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

					if (err != nil) != tt.wantErr {
						t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
					}
					if !cmp.Equal(gotRes, wantRes, cmpopts.EquateComparable(postgresflex.NullableInt32{})) {
						t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
					}
				})
			})
		}
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

func TestCreateUserWaitHandler(t *testing.T) {
	tests := []struct {
		desc                string
		userGetFails        bool
		userState           string
		usersGetErrorStatus int
		wantErr             bool
		wantResp            bool
	}{
		{
			desc:         "create_succeeded",
			userGetFails: false,
			userState:    "PROCESSED",
			wantErr:      false,
			wantResp:     true,
		},
		{
			desc:         "user_get_fails",
			userGetFails: true,
			wantErr:      true,
			wantResp:     false,
		},
		{
			desc:                "users_get_fails",
			userGetFails:        true,
			usersGetErrorStatus: 423,
			wantErr:             true,
			wantResp:            false,
		},
		{
			desc:         "timeout",
			userGetFails: false,
			userState:    "",
			wantErr:      true,
			wantResp:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				userId := int64(34)

				apiClient := newAPIMock(&mockSettings{
					userGetFails: tt.userGetFails,
					userId:       userId,
					userState:    tt.userState,
				})

				var wantRes *postgresflex.GetUserResponse
				if tt.wantResp {
					wantRes = &postgresflex.GetUserResponse{
						Id:    userId,
						State: tt.userState,
					}
				}

				handler := CreateUserWaitHandler(context.Background(), apiClient, "", "", "", userId)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
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
				userId := int64(34)

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
