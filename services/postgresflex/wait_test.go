package postgresflex

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Used for testing instance operations
type apiClientInstanceMocked struct {
	instanceId          string
	instanceState       string
	instanceIsDeleted   bool
	instanceGetFails    bool
	usersGetErrorStatus int
}

func (a *apiClientInstanceMocked) GetInstanceExecute(_ context.Context, _, _ string) (*InstanceResponse, error) {
	if a.instanceGetFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	if a.instanceIsDeleted {
		return nil, &GenericOpenAPIError{
			statusCode: 404,
		}
	}

	return &InstanceResponse{
		Item: &InstanceSingleInstance{
			Id:     &a.instanceId,
			Status: &a.instanceState,
		},
	}, nil
}

func (a *apiClientInstanceMocked) GetUsersExecute(_ context.Context, _, _ string) (*UsersResponse, error) {
	if a.usersGetErrorStatus != 0 {
		return nil, &GenericOpenAPIError{
			statusCode: a.usersGetErrorStatus,
		}
	}

	aux := int32(0)
	return &UsersResponse{
		Count: &aux,
		Items: &[]InstanceListUser{},
	}, nil
}

// Used for testing user operations
type apiClientUserMocked struct {
	getFails      bool
	userId        string
	isUserDeleted bool
}

func (a *apiClientUserMocked) GetUserExecute(_ context.Context, _, _, _ string) (*UserResponse, error) {
	if a.getFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	if a.isUserDeleted {
		return nil, &GenericOpenAPIError{
			statusCode: 404,
		}
	}

	return &UserResponse{
		Item: &UserResponseUser{
			Id: &a.userId,
		},
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
		desc                string
		instanceGetFails    bool
		instanceState       string
		usersGetErrorStatus int
		wantErr             bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateSuccess,
			wantErr:          false,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceState:    InstanceStateFailed,
			wantErr:          true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceState:    InstanceStateEmpty,
			wantErr:          true,
		},
		{
			desc:             "instance_get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
		{
			desc:                "users_get_fails",
			instanceGetFails:    false,
			instanceState:       InstanceStateSuccess,
			usersGetErrorStatus: 500,
			wantErr:             true,
		},
		{
			desc:                "users_get_fails_2",
			instanceGetFails:    false,
			instanceState:       InstanceStateSuccess,
			usersGetErrorStatus: 400,
			wantErr:             true,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    "ANOTHER STATE",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				instanceId:          instanceId,
				instanceState:       tt.instanceState,
				instanceGetFails:    tt.instanceGetFails,
				usersGetErrorStatus: tt.usersGetErrorStatus,
			}

			var wantRes *InstanceResponse
			if (tt.instanceState == InstanceStateSuccess) && !tt.instanceGetFails && (tt.usersGetErrorStatus == 0) {
				wantRes = &InstanceResponse{
					Item: &InstanceSingleInstance{
						Id:     &instanceId,
						Status: &tt.instanceState,
					},
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", instanceId)

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

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    string
		wantErr          bool
	}{
		{
			desc:             "update_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateSuccess,
			wantErr:          false,
		},
		{
			desc:             "update_failed",
			instanceGetFails: false,
			instanceState:    InstanceStateFailed,
			wantErr:          true,
		},
		{
			desc:             "update_failed_2",
			instanceGetFails: false,
			instanceState:    InstanceStateEmpty,
			wantErr:          true,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    "ANOTHER STATE",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				instanceId:       instanceId,
				instanceState:    tt.instanceState,
				instanceGetFails: tt.instanceGetFails,
			}

			var wantRes *InstanceResponse
			if !tt.instanceGetFails {
				wantRes = &InstanceResponse{
					Item: &InstanceSingleInstance{
						Id:     &instanceId,
						Status: &tt.instanceState,
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
		desc             string
		instanceGetFails bool
		instanceState    string
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			instanceGetFails: false,
			instanceState:    InstanceStateSuccess,
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

			apiClient := &apiClientInstanceMocked{
				instanceGetFails:  tt.instanceGetFails,
				instanceIsDeleted: tt.instanceState == InstanceStateSuccess,
				instanceId:        instanceId,
				instanceState:     tt.instanceState,
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

			apiClient := &apiClientUserMocked{
				getFails:      tt.getFails,
				userId:        userId,
				isUserDeleted: !tt.deleteFails,
			}

			handler := DeleteUserWaitHandler(context.Background(), apiClient, "", "", userId)

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
