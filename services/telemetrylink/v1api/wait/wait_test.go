package wait

import (
	"context"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	telemetrylink "github.com/stackitcloud/stackit-sdk-go/services/telemetrylink/v1api"
)

type mockSettings struct {
	getFails      bool
	notFound      bool
	resourceState telemetrylink.TelemetryLinkResponseStatus
}

func newAPIMock(settings mockSettings) telemetrylink.DefaultAPI {
	return &telemetrylink.DefaultAPIServiceMock{
		GetOrganizationTelemetryLinkExecuteMock: utils.Ptr(func(_ telemetrylink.ApiGetOrganizationTelemetryLinkRequest) (*telemetrylink.TelemetryLinkResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.notFound {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &telemetrylink.TelemetryLinkResponse{
				Status: settings.resourceState,
			}, nil
		}),
		GetFolderTelemetryLinkExecuteMock: utils.Ptr(func(_ telemetrylink.ApiGetFolderTelemetryLinkRequest) (*telemetrylink.TelemetryLinkResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.notFound {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &telemetrylink.TelemetryLinkResponse{
				Status: settings.resourceState,
			}, nil
		}),
		GetProjectTelemetryLinkExecuteMock: utils.Ptr(func(_ telemetrylink.ApiGetProjectTelemetryLinkRequest) (*telemetrylink.TelemetryLinkResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.notFound {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &telemetrylink.TelemetryLinkResponse{
				Status: settings.resourceState,
			}, nil
		}),
	}
}

func TestCreateOrganizationTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := CreateOrganizationTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestUpdateOrganizationTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := UpdateOrganizationTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestPartialUpdateOrganizationTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_or_update_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := PartialUpdateOrganizationTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestDeleteOrganizationTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		notFound      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "not_found",
			getFails:      false,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
					notFound:      tt.notFound,
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				} else {
					wantRes = nil
				}

				handler := DeleteOrganizationTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestCreateFolderTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := CreateFolderTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestUpdateFolderTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := UpdateFolderTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestPartialUpdateFolderTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_or_update_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := PartialUpdateFolderTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestDeleteFolderTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		notFound      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "not_found",
			getFails:      false,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
					notFound:      tt.notFound,
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				} else {
					wantRes = nil
				}

				handler := DeleteFolderTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestCreateProjectTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := CreateProjectTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestUpdateProjectTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := UpdateProjectTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestPartialUpdateProjectTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_or_update_succeeded",
			getFails:      false,
			resourceState: telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE,
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

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				}

				handler := PartialUpdateProjectTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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

func TestDeleteProjectTelemetryLinkWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		notFound      bool
		resourceState telemetrylink.TelemetryLinkResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "not_found",
			getFails:      false,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
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
					notFound:      tt.notFound,
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *telemetrylink.TelemetryLinkResponse
				if tt.wantResp {
					wantRes = &telemetrylink.TelemetryLinkResponse{
						Status: tt.resourceState,
					}
				} else {
					wantRes = nil
				}

				handler := DeleteProjectTelemetryLinkWaitHandler(context.Background(), apiClient, "pid", "eu01")

				gotRes, err := handler.WaitWithContext(context.Background())

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
