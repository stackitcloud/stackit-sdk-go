package mongodbflex

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// Used for testing instance operations
type apiClientInstanceMocked struct {
	instanceId        string
	instanceState     string
	instanceIsDeleted bool
	instanceGetFails  bool
}

func (a *apiClientInstanceMocked) GetInstanceExecute(_ context.Context, _, _ string) (*GetInstanceResponse, error) {
	if a.instanceGetFails {
		return nil, &oapiError.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if a.instanceIsDeleted {
		return nil, &oapiError.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &GetInstanceResponse{
		Item: &InstanceSingleInstance{
			Id:     &a.instanceId,
			Status: &a.instanceState,
		},
	}, nil
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

			var wantRes *GetInstanceResponse
			if (tt.instanceState == InstanceStateSuccess) && !tt.instanceGetFails {
				wantRes = &GetInstanceResponse{
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

			var wantRes *GetInstanceResponse
			if !tt.instanceGetFails {
				wantRes = &GetInstanceResponse{
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
