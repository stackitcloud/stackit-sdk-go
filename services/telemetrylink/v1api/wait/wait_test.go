package wait

import (
	"context"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
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

type backflushMockSettings struct {
	getFails      bool
	resourceState telemetrylink.BackflushJobResponseStatus
}

func newBackflushAPIMock(settings backflushMockSettings) telemetrylink.DefaultAPI {
	get := func() (*telemetrylink.BackflushJobResponse, error) {
		if settings.getFails {
			return nil, &oapierror.GenericOpenAPIError{StatusCode: 500}
		}
		return &telemetrylink.BackflushJobResponse{Status: settings.resourceState}, nil
	}
	return &telemetrylink.DefaultAPIServiceMock{
		GetOrganizationBackflushJobExecuteMock: utils.Ptr(func(_ telemetrylink.ApiGetOrganizationBackflushJobRequest) (*telemetrylink.BackflushJobResponse, error) {
			return get()
		}),
		GetFolderBackflushJobExecuteMock: utils.Ptr(func(_ telemetrylink.ApiGetFolderBackflushJobRequest) (*telemetrylink.BackflushJobResponse, error) {
			return get()
		}),
		GetProjectBackflushJobExecuteMock: utils.Ptr(func(_ telemetrylink.ApiGetProjectBackflushJobRequest) (*telemetrylink.BackflushJobResponse, error) {
			return get()
		}),
	}
}

type createBackflushJobHandler func(ctx context.Context, a telemetrylink.DefaultAPI, id, region, backflushJobId string) *wait.AsyncActionHandler[telemetrylink.BackflushJobResponse]

func testCreateBackflushJobWaitHandler(t *testing.T, newHandler createBackflushJobHandler) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState telemetrylink.BackflushJobResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{desc: "active", resourceState: telemetrylink.BACKFLUSHJOBRESPONSESTATUS_ACTIVE, wantResp: true},
		{desc: "complete", resourceState: telemetrylink.BACKFLUSHJOBRESPONSESTATUS_COMPLETE, wantResp: true},
		{desc: "failed", resourceState: telemetrylink.BACKFLUSHJOBRESPONSESTATUS_FAILED, wantErr: true, wantResp: true},
		{desc: "deleting", resourceState: telemetrylink.BACKFLUSHJOBRESPONSESTATUS_DELETING, wantErr: true, wantResp: true},
		{desc: "get_fails", getFails: true, wantErr: true},
		{desc: "timeout", resourceState: telemetrylink.BACKFLUSHJOBRESPONSESTATUS_RECONCILING, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newBackflushAPIMock(backflushMockSettings{getFails: tt.getFails, resourceState: tt.resourceState})

				handler := newHandler(context.Background(), apiClient, "id", "eu01", "job-id")
				gotRes, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr {
					want := &telemetrylink.BackflushJobResponse{Status: tt.resourceState}
					if !cmp.Equal(gotRes, want) {
						t.Fatalf("handler gotRes = %v, want %v", gotRes, want)
					}
				}
			})
		})
	}
}

func TestCreateOrganizationBackflushJobWaitHandler(t *testing.T) {
	testCreateBackflushJobWaitHandler(t, CreateOrganizationBackflushJobWaitHandler)
}

func TestCreateFolderBackflushJobWaitHandler(t *testing.T) {
	testCreateBackflushJobWaitHandler(t, CreateFolderBackflushJobWaitHandler)
}

func TestCreateProjectBackflushJobWaitHandler(t *testing.T) {
	testCreateBackflushJobWaitHandler(t, CreateProjectBackflushJobWaitHandler)
}
