package wait_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	sca "github.com/stackitcloud/stackit-sdk-go/services/sca/v1alphaapi"
	scawait "github.com/stackitcloud/stackit-sdk-go/services/sca/v1alphaapi/wait"
)

type mockSettings struct {
	getFails      bool
	getErrorCode  int
	resourceState sca.CurrentStatus
}

func newAPIMock(settings mockSettings) sca.DefaultAPI {
	return &sca.DefaultAPIServiceMock{
		GetApplicationExecuteMock: utils.Ptr(func(r sca.ApiGetApplicationRequest) (*sca.Application, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.getErrorCode,
				}
			}

			return &sca.Application{
				DisplayName:   "application",
				EnvironmentId: sca.PtrString("environment"),
				RuntimeStatus: &sca.RuntimeStatus{
					CurrentStatus: settings.resourceState.Ptr(),
				},
			}, nil
		}),
	}
}

func TestCreateOrUpdateApplicationWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		mockSettings mockSettings
		wantErr      bool
	}{
		{
			desc: "create_succeeded",
			mockSettings: mockSettings{
				resourceState: sca.CURRENTSTATUS_CURRENT_STATUS_RUNNING,
				getFails:      false,
			},
		},
		{
			desc: "create_idle_succeeded",
			mockSettings: mockSettings{
				resourceState: sca.CURRENTSTATUS_CURRENT_STATUS_IDLE,
				getFails:      false,
			},
		},
		{
			desc: "get_fails",
			mockSettings: mockSettings{
				getFails: true,
			},
			wantErr: true,
		},
		{
			desc: "timeout",
			mockSettings: mockSettings{
				getFails:      false,
				resourceState: sca.CURRENTSTATUS_CURRENT_STATUS_PROGRESSING,
			},
			wantErr: true,
		},
	}

	handlers := map[string]func(context.Context, sca.DefaultAPI, string, string, string) *wait.AsyncActionHandler[sca.Application]{
		"create": scawait.CreateApplicationWaitHandler,
		"update": scawait.UpdateApplicationWaitHandler,
	}

	for handlerKey, handlerFn := range handlers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s - %s", handlerKey, tt.desc), func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					apiClient := newAPIMock(tt.mockSettings)

					handler := handlerFn(t.Context(), apiClient, "pid", "environment", "application")

					_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(t.Context())
					if err != nil != tt.wantErr {
						t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
					}
				})
			})
		}
	}
}

func TestDeleteApplicationWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		mockSettings mockSettings
		wantErr      bool
	}{
		{
			desc: "delete_succeeded",
			mockSettings: mockSettings{
				getFails:     true,
				getErrorCode: http.StatusNotFound,
			},
		},
		{
			desc: "get_fails",
			mockSettings: mockSettings{
				getFails:     true,
				getErrorCode: http.StatusInternalServerError,
			},
			wantErr: true,
		},
		{
			desc: "timeout",
			mockSettings: mockSettings{
				getFails:      false,
				resourceState: sca.CURRENTSTATUS_CURRENT_STATUS_RUNNING,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(tt.mockSettings)

				handler := scawait.DeleteApplicationWaitHandler(t.Context(), apiClient, "pid", "environment", "application")

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(t.Context())
				if err != nil != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
