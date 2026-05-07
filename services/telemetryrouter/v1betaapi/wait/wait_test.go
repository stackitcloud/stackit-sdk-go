package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	telemetryrouter "github.com/stackitcloud/stackit-sdk-go/services/telemetryrouter/v1betaapi"
)

type mockSettings struct {
	getFails      bool
	resourceState string
}

func newAPIMock(settings mockSettings) telemetryrouter.DefaultAPI {
	return &telemetryrouter.DefaultAPIServiceMock{
		GetTelemetryRouterExecuteMock: utils.Ptr(func(_ telemetryrouter.ApiGetTelemetryRouterRequest) (*telemetryrouter.TelemetryRouterResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &telemetryrouter.TelemetryRouterResponse{
				Id:     "trid",
				Status: settings.resourceState,
			}, nil
		}),
		GetDestinationExecuteMock: utils.Ptr(func(_ telemetryrouter.ApiGetDestinationRequest) (*telemetryrouter.DestinationResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &telemetryrouter.DestinationResponse{
				Id:     "did",
				Status: settings.resourceState,
			}, nil
		}),
		GetAccessTokenExecuteMock: utils.Ptr(func(_ telemetryrouter.ApiGetAccessTokenRequest) (*telemetryrouter.GetAccessTokenResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &telemetryrouter.GetAccessTokenResponse{
				Id:             "atid",
				ExpirationTime: telemetryrouter.NullableTime{},
				Status:         settings.resourceState,
			}, nil
		}),
	}
}

func TestCreateTelemetryRouterWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: TELEMETRYROUTER_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.TelemetryRouterResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.TelemetryRouterResponse{
						Status: tt.resourceState,
						Id:     "trid",
					}
				}

				handler := CreateTelemetryRouterWaitHandler(context.Background(), apiClient, "pid", "eu01", "trid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

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

func TestUpdateTelemetryRouterWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: TELEMETRYROUTER_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.TelemetryRouterResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.TelemetryRouterResponse{
						Status: tt.resourceState,
						Id:     "trid",
					}
				}

				handler := UpdateTelemetryRouterWaitHandler(context.Background(), apiClient, "pid", "eu01", "trid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

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

func TestDeleteTelemetryRouterWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: string(tt.resourceState),
				})

				var wantRes *telemetryrouter.TelemetryRouterResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.TelemetryRouterResponse{
						Status: tt.resourceState,
						Id:     "trid",
					}
				} else {
					wantRes = nil
				}

				handler := DeleteTelemetryRouterWaitHandler(context.Background(), apiClient, "pid", "zid", "trid")

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
		})
	}
}

func TestCreateDestinationWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: DESTINATION_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.DestinationResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.DestinationResponse{
						Status: tt.resourceState,
						Id:     "did",
					}
				}

				handler := CreateDestinationWaitHandler(context.Background(), apiClient, "pid", "eu01", "trid", "did")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

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

func TestUpdateDestinationWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: DESTINATION_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.DestinationResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.DestinationResponse{
						Status: tt.resourceState,
						Id:     "did",
					}
				}

				handler := UpdateDestinationWaitHandler(context.Background(), apiClient, "pid", "eu01", "trid", "did")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

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

func TestDeleteDestinationWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: string(tt.resourceState),
				})

				var wantRes *telemetryrouter.DestinationResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.DestinationResponse{
						Status: tt.resourceState,
						Id:     "did",
					}
				} else {
					wantRes = nil
				}

				handler := DeleteDestinationWaitHandler(context.Background(), apiClient, "pid", "zid", "trid", "did")

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
		})
	}
}

func TestCreateAccessTokenWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: ACCESSTOKEN_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.GetAccessTokenResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.GetAccessTokenResponse{
						Status: tt.resourceState,
						Id:     "atid",
					}
				}

				handler := CreateAccessTokenWaitHandler(context.Background(), apiClient, "pid", "eu01", "trid", "atid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmp.Comparer(compareNullableTime)) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestUpdateAccessTokenWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: ACCESSTOKEN_ACTIVE,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.GetAccessTokenResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.GetAccessTokenResponse{
						Status: tt.resourceState,
						Id:     "atid",
					}
				}

				handler := UpdateAccessTokenWaitHandler(context.Background(), apiClient, "pid", "eu01", "trid", "atid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmp.Comparer(compareNullableTime)) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestDeleteAccessTokenWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetryrouter.GetAccessTokenResponse
				if tt.wantResp {
					wantRes = &telemetryrouter.GetAccessTokenResponse{
						Status: tt.resourceState,
						Id:     "atid",
					}
				} else {
					wantRes = nil
				}

				handler := DeleteAccessTokenWaitHandler(context.Background(), apiClient, "pid", "zid", "trid", "atid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if wantRes == nil && gotRes != nil {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
				if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmp.Comparer(compareNullableTime)) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func compareNullableTime(x, y telemetryrouter.NullableTime) bool {
	if x.IsSet() != y.IsSet() {
		return false
	}

	if !x.IsSet() && !y.IsSet() {
		return true
	}

	valX := x.Get()
	valY := y.Get()

	if valX == nil || valY == nil {
		return valX == valY
	}

	return *valX == *valY
}
