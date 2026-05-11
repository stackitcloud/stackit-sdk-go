package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	dremio "github.com/stackitcloud/stackit-sdk-go/services/dremio/v1alphaapi"
)

type mockSettings struct {
	isDeleted     bool
	getFails      bool
	resourceState string
}

func newAPIMock(settings mockSettings) dremio.DefaultAPI {
	return &dremio.DefaultAPIServiceMock{
		GetDremioInstanceExecuteMock: utils.Ptr(
			func(_ dremio.ApiGetDremioInstanceRequest) (*dremio.DremioResponse, error) {
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

				return &dremio.DremioResponse{
					State: settings.resourceState,
				}, nil
			},
		),
		GetDremioUserExecuteMock: utils.Ptr(
			func(_ dremio.ApiGetDremioUserRequest) (*dremio.DremioUserResponse, error) {
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

				return &dremio.DremioUserResponse{
					State: settings.resourceState,
				}, nil
			},
		),
	}
}

func TestCreateDremioWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState string
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "create_succeeded",
			getFails:      false,
			resourceState: DREMIOSTATE_ACTIVE,
			wantError:     false,
			wantResponse:  true,
		},
		{
			description:   "create_failed",
			getFails:      false,
			resourceState: DREMIOSTATE_ERROR,
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

				var expectedResponse *dremio.DremioResponse
				if currentTest.wantResponse {
					expectedResponse = &dremio.DremioResponse{
						State: currentTest.resourceState,
					}
				}

				handler := CreateDremioWaitHandler(context.Background(), apiClient, "pid", "zid", "dremioId")

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

func TestUpdateDremioWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState string
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "update_succeeded",
			getFails:      false,
			resourceState: DREMIOSTATE_ACTIVE,
			wantError:     false,
			wantResponse:  true,
		},
		{
			description:   "update_failed",
			getFails:      false,
			resourceState: DREMIOSTATE_ERROR,
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

				var expectedResponse *dremio.DremioResponse
				if currentTest.wantResponse {
					expectedResponse = &dremio.DremioResponse{
						State: currentTest.resourceState,
					}
				}

				handler := UpdateDremioWaitHandler(context.Background(), apiClient, "pid", "zid", "dremioId")

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

func TestDeleteDremioWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		isDeleted     bool
		getFails      bool
		resourceState string
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
			resourceState: DREMIOSTATE_ERROR,
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

				var expectedResponse *dremio.DremioResponse
				if currentTest.wantResponse {
					expectedResponse = &dremio.DremioResponse{
						State: currentTest.resourceState,
					}
				}

				handler := DeleteDremioWaitHandler(context.Background(), apiClient, "pid", "zid", "dremioId")

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

func TestCreateDremioUserWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState string
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "create_succeeded",
			getFails:      false,
			resourceState: DREMIOUSERSTATE_ACTIVE,
			wantError:     false,
			wantResponse:  true,
		},
		{
			description:   "create_failed",
			getFails:      false,
			resourceState: DREMIOUSERSTATE_ERROR,
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

				var expectedResponse *dremio.DremioUserResponse
				if currentTest.wantResponse {
					expectedResponse = &dremio.DremioUserResponse{
						State: currentTest.resourceState,
					}
				}

				handler := CreateDremioUserWaitHandler(context.Background(), apiClient, "pid", "zid", "dremioId", "userId")

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

func TestUpdateDremioUserWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState string
		wantError     bool
		wantResponse  bool
	}{
		{
			description:   "update_succeeded",
			getFails:      false,
			resourceState: DREMIOUSERSTATE_ACTIVE,
			wantError:     false,
			wantResponse:  true,
		},
		{
			description:   "update_failed",
			getFails:      false,
			resourceState: DREMIOUSERSTATE_ERROR,
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

				var expectedResponse *dremio.DremioUserResponse
				if currentTest.wantResponse {
					expectedResponse = &dremio.DremioUserResponse{
						State: currentTest.resourceState,
					}
				}

				handler := UpdateDremioUserWaitHandler(context.Background(), apiClient, "pid", "zid", "dremioId", "userId")

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

func TestDeleteDremioUserWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		isDeleted     bool
		getFails      bool
		resourceState string
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
			resourceState: DREMIOUSERSTATE_ERROR,
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

				var expectedResponse *dremio.DremioUserResponse
				if currentTest.wantResponse {
					expectedResponse = &dremio.DremioUserResponse{
						State: currentTest.resourceState,
					}
				}

				handler := DeleteDremioUserWaitHandler(context.Background(), apiClient, "pid", "zid", "dremioId", "userId")

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
