package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	ufw "github.com/stackitcloud/stackit-sdk-go/services/ufw/v1api"
)

type mockSettings struct {
	isDeleted     bool
	getFails      bool
	resourceState ufw.RuleResponseStatus
}

func newAPIMock(settings mockSettings) ufw.DefaultAPI {
	return &ufw.DefaultAPIServiceMock{
		GetRuleExecuteMock: utils.Ptr(
			func(_ ufw.ApiGetRuleRequest) (*ufw.RuleResponse, error) {
				if settings.isDeleted {
					return nil, &oapierror.GenericOpenAPIError{
						StatusCode: http.StatusNotFound,
					}
				}
				if settings.getFails {
					return nil, &oapierror.GenericOpenAPIError{
						StatusCode: 500,
					}
				}

				return &ufw.RuleResponse{
					Status: utils.Ptr(settings.resourceState),
				}, nil
			},
		),
	}
}

func TestCreateRuleWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState ufw.RuleResponseStatus
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "create_succeeded",
			getFails:      false,
			resourceState: ufw.RULERESPONSESTATUS_ACTIVE,
			wantError:     false,
			wantResponse:  true,
		},
		{
			description:   "create_failed",
			getFails:      false,
			resourceState: ufw.RULERESPONSESTATUS_ERROR,
			wantError:     true,
			wantResponse:  true,
		},
		{
			description:   "get_fails",
			getFails:      true,
			resourceState: "",
			wantError:     true,
			wantResponse:  false,
		},
		{
			description:   "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantError:     true,
			wantResponse:  false,
		},
	}
	for _, currentTest := range tests {
		t.Run(currentTest.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      currentTest.getFails,
					resourceState: currentTest.resourceState,
					isDeleted:     false,
				})

				var expectedResponse *ufw.RuleResponse
				if currentTest.wantResponse {
					expectedResponse = &ufw.RuleResponse{
						Status: utils.Ptr(currentTest.resourceState),
					}
				}

				handler := CreateRuleWaitHandler(context.Background(), apiClient, "pid", "region", "ruleId")

				gotResponse, err := handler.WaitWithContext(context.Background())

				if (err != nil) != currentTest.wantError {
					t.Fatalf("handler error = %v, wantErr %v", err, currentTest.wantError)
				}
				if !cmp.Equal(gotResponse, expectedResponse) {
					t.Fatalf("handler gotResponse = %v, expectedResponse = %v", gotResponse, expectedResponse)
				}
			})
		})
	}
}

func TestUpdateRuleWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState ufw.RuleResponseStatus
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "update_succeeded",
			getFails:      false,
			resourceState: ufw.RULERESPONSESTATUS_ACTIVE,
			wantError:     false,
			wantResponse:  true,
		},
		{
			description:   "update_failed",
			getFails:      false,
			resourceState: ufw.RULERESPONSESTATUS_ERROR,
			wantError:     true,
			wantResponse:  true,
		},
		{
			description:   "get_fails",
			getFails:      true,
			resourceState: "",
			wantError:     true,
			wantResponse:  false,
		},
		{
			description:   "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantError:     true,
			wantResponse:  false,
		},
	}
	for _, currentTest := range tests {
		t.Run(currentTest.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      currentTest.getFails,
					resourceState: currentTest.resourceState,
					isDeleted:     false,
				})

				var expectedResponse *ufw.RuleResponse
				if currentTest.wantResponse {
					expectedResponse = &ufw.RuleResponse{
						Status: utils.Ptr(currentTest.resourceState),
					}
				}

				handler := UpdateRuleWaitHandler(context.Background(), apiClient, "pid", "region", "ruleId")

				gotResponse, err := handler.WaitWithContext(context.Background())

				if (err != nil) != currentTest.wantError {
					t.Fatalf("handler error = %v, wantErr %v", err, currentTest.wantError)
				}
				if !cmp.Equal(gotResponse, expectedResponse) {
					t.Fatalf("handler gotResponse = %v, expectedResponse = %v", gotResponse, expectedResponse)
				}
			})
		})
	}
}

func TestDeleteRuleWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		isDeleted     bool
		getFails      bool
		resourceState ufw.RuleResponseStatus
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "delete_succeeded",
			isDeleted:     true,
			getFails:      false,
			resourceState: "",
			wantError:     false,
			wantResponse:  false,
		},
		{
			description:   "delete_failed",
			isDeleted:     false,
			getFails:      false,
			resourceState: ufw.RULERESPONSESTATUS_ERROR,
			wantError:     true,
			wantResponse:  true,
		},
		{
			description:   "get_fails",
			isDeleted:     false,
			getFails:      true,
			resourceState: "",
			wantError:     true,
			wantResponse:  false,
		},
		{
			description:   "timeout",
			isDeleted:     false,
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantError:     true,
			wantResponse:  false,
		},
	}
	for _, currentTest := range tests {
		t.Run(currentTest.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      currentTest.getFails,
					resourceState: currentTest.resourceState,
					isDeleted:     currentTest.isDeleted,
				})

				var expectedResponse *ufw.RuleResponse
				if currentTest.wantResponse {
					expectedResponse = &ufw.RuleResponse{
						Status: utils.Ptr(currentTest.resourceState),
					}
				}

				handler := DeleteRuleWaitHandler(context.Background(), apiClient, "pid", "region", "ruleId")

				gotResponse, err := handler.WaitWithContext(context.Background())

				if (err != nil) != currentTest.wantError {
					t.Fatalf("handler error = %v, wantErr %v", err, currentTest.wantError)
				}

				if !currentTest.wantResponse && gotResponse != nil {
					t.Fatalf("handler gotResponse = %v, expectedResponse = %v", gotResponse, expectedResponse)
				}
				if currentTest.wantResponse && !cmp.Equal(gotResponse, expectedResponse) {
					t.Fatalf("handler gotResponse = %v, expectedResponse = %v", gotResponse, expectedResponse)
				}
			})
		})
	}
}
