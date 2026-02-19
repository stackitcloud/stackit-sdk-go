package wait

import (
	"context"
	"testing"
	"time"
	"net/http"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	postgresflex "github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v2api"
)

type mockSettings struct {
	instanceId             string
	instanceState          string
	instanceIsForceDeleted bool
	instanceGetFails       bool
	usersGetErrorStatus    int

	userGetFails  bool
	userId        string
	isUserDeleted bool
}

func newApiClientMock(settings mockSettings) postgresflex.DefaultAPI {
	return &postgresflex.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(r postgresflex.ApiGetInstanceRequest) (*postgresflex.InstanceResponse, *http.Response, error) {
			if settings.instanceGetFails {
				return nil, nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.instanceIsForceDeleted {
				return nil, nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &postgresflex.InstanceResponse{
				Item: &postgresflex.Instance{
					Id:     &settings.instanceId,
					Status: &settings.instanceState,
				},
			}, &http.Response{}, nil
		}),

		ListUsersExecuteMock: utils.Ptr(func(r postgresflex.ApiListUsersRequest) (*postgresflex.ListUsersResponse, *http.Response, error) {
			if settings.usersGetErrorStatus != 0 {
				return nil, nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.usersGetErrorStatus,
				}
			}

			aux := int64(0)
			return &postgresflex.ListUsersResponse{
				Count: &aux,
				Items: []postgresflex.ListUsersResponseItem{},
			}, &http.Response{}, nil
		}),

		GetUserExecuteMock: utils.Ptr(func(r postgresflex.ApiGetUserRequest) (*postgresflex.GetUserResponse, *http.Response, error) {
			if settings.userGetFails {
				return nil, nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.isUserDeleted {
				return nil, nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &postgresflex.GetUserResponse{
				Item: &postgresflex.UserResponse{
					Id: &settings.userId,
				},
			}, &http.Response{}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                string
		instanceGetFails    bool
		instanceState       string
		usersGetErrorStatus int
		wantErr             bool
		wantResp            bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateSuccess,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceState:    InstanceStateFailed,
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
			instanceState:       InstanceStateSuccess,
			usersGetErrorStatus: 500,
			wantErr:             true,
			wantResp:            false,
		},
		{
			desc:                "users_get_fails_2",
			instanceGetFails:    false,
			instanceState:       InstanceStateSuccess,
			usersGetErrorStatus: 400,
			wantErr:             true,
			wantResp:            true,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    InstanceStateProgressing,
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := newApiClientMock(mockSettings{
				instanceId:          instanceId,
				instanceState:       tt.instanceState,
				instanceGetFails:    tt.instanceGetFails,
				usersGetErrorStatus: tt.usersGetErrorStatus,
			})

			var wantRes *postgresflex.InstanceResponse
			if tt.wantResp {
				wantRes = &postgresflex.InstanceResponse{
					Item: &postgresflex.Instance{
						Id:     &instanceId,
						Status: utils.Ptr(tt.instanceState),
					},
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    string
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "update_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateSuccess,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "update_failed",
			instanceGetFails: false,
			instanceState:    InstanceStateFailed,
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
			instanceState:    InstanceStateProgressing,
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := newApiClientMock(mockSettings{
				instanceId:       instanceId,
				instanceState:    tt.instanceState,
				instanceGetFails: tt.instanceGetFails,
			})

			var wantRes *postgresflex.InstanceResponse
			if tt.wantResp {
				wantRes = &postgresflex.InstanceResponse{
					Item: &postgresflex.Instance{
						Id:     &instanceId,
						Status: utils.Ptr(tt.instanceState),
					},
				}
			}

			handler := PartialUpdateInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

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
		desc             string
		instanceGetFails bool
		instanceState    string
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateDeleted,
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			instanceGetFails: false,
			instanceState:    InstanceStateFailed,
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
			instanceId := "foo-bar"

			apiClient := newApiClientMock(mockSettings{
				instanceGetFails: tt.instanceGetFails,
				instanceId:       instanceId,
				instanceState:    tt.instanceState,
			})

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestForceDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    string
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateDeleted,
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			instanceGetFails: false,
			instanceState:    InstanceStateFailed,
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
			instanceId := "foo-bar"

			apiClient := newApiClientMock(mockSettings{
				instanceGetFails:       tt.instanceGetFails,
				instanceIsForceDeleted: tt.instanceState == InstanceStateDeleted,
				instanceId:             instanceId,
				instanceState:          tt.instanceState,
			})

			handler := ForceDeleteInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
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
			userId := "foo-bar"

			apiClient := newApiClientMock(mockSettings{
				userGetFails:      tt.getFails,
				userId:        userId,
				isUserDeleted: !tt.deleteFails,
			})

			handler := DeleteUserWaitHandler(context.Background(), apiClient, "", "", "", userId)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
