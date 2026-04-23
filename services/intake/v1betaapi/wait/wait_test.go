package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	intake "github.com/stackitcloud/stackit-sdk-go/services/intake/v1betaapi"
)

type mockSettings struct {
	getRunnerFails       bool
	getIntakeFails       bool
	getUserFails         bool
	getErrorCode         int
	returnRunner         bool
	returnIntake         bool
	returnUser           bool
	intakeRunnerResponse *intake.IntakeRunnerResponse
	intakeResponse       *intake.IntakeResponse
	intakeUserResponse   *intake.IntakeUserResponse
}

func newAPIMock(settings *mockSettings) intake.DefaultAPI {
	return &intake.DefaultAPIServiceMock{
		GetIntakeRunnerExecuteMock: utils.Ptr(func(_ intake.ApiGetIntakeRunnerRequest) (*intake.IntakeRunnerResponse, error) {
			if settings.getRunnerFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.getErrorCode,
				}
			}
			if !settings.returnRunner {
				return nil, nil
			}
			return settings.intakeRunnerResponse, nil
		}),
		GetIntakeExecuteMock: utils.Ptr(func(_ intake.ApiGetIntakeRequest) (*intake.IntakeResponse, error) {
			if settings.getIntakeFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.getErrorCode,
				}
			}
			if !settings.returnIntake {
				return nil, nil
			}
			return settings.intakeResponse, nil
		}),
		GetIntakeUserExecuteMock: utils.Ptr(func(_ intake.ApiGetIntakeUserRequest) (*intake.IntakeUserResponse, error) {
			if settings.getUserFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.getErrorCode,
				}
			}
			if !settings.returnUser {
				return nil, nil
			}
			return settings.intakeUserResponse, nil
		}),
	}
}

const region = "eu01"

var (
	projectId      = uuid.NewString()
	intakeRunnerId = uuid.NewString()
	intakeId       = uuid.NewString()
	intakeUserId   = uuid.NewString()
)

func TestCreateOrUpdateIntakeRunnerWaitHandler(t *testing.T) {
	tests := []struct {
		desc                 string
		getFails             bool
		getErrorCode         int
		wantErr              bool
		wantResp             bool
		returnRunner         bool
		intakeRunnerResponse *intake.IntakeRunnerResponse
	}{
		{
			desc:         "succeeded",
			getFails:     false,
			wantErr:      false,
			wantResp:     true,
			returnRunner: true,
			intakeRunnerResponse: &intake.IntakeRunnerResponse{
				Id:    intakeRunnerId,
				State: INTAKERUNNERRESPONSESTATE_ACTIVE,
			},
		},
		{
			desc:         "get fails",
			getFails:     true,
			getErrorCode: http.StatusInternalServerError,
			wantErr:      true,
			wantResp:     false,
			returnRunner: false,
		},
		{
			desc:         "timeout",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnRunner: true,
			intakeRunnerResponse: &intake.IntakeRunnerResponse{
				Id:    intakeRunnerId,
				State: INTAKERUNNERRESPONSESTATE_RECONCILING,
			},
		},
		{
			desc:         "nil response",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnRunner: false,
		},
		{
			desc:         "nil id in response",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnRunner: true,
			intakeRunnerResponse: &intake.IntakeRunnerResponse{
				State: INTAKERUNNERRESPONSESTATE_RECONCILING,
			},
		},
		{
			desc:         "nil state in response",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnRunner: true,
			intakeRunnerResponse: &intake.IntakeRunnerResponse{
				Id: intakeRunnerId,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getRunnerFails:       tt.getFails,
					getErrorCode:         tt.getErrorCode,
					returnRunner:         tt.returnRunner,
					intakeRunnerResponse: tt.intakeRunnerResponse,
				})

				var wantResp *intake.IntakeRunnerResponse
				if tt.wantResp {
					wantResp = tt.intakeRunnerResponse
				}

				handler := CreateOrUpdateIntakeRunnerWaitHandler(context.Background(), apiClient, projectId, region, intakeRunnerId)
				got, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(got, wantResp) {
					t.Fatalf("handler got = %v, want %v", got, wantResp)
				}
			})
		})
	}
}

func TestDeleteIntakeRunnerWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		getFails     bool
		getErrorCode int
		wantErr      bool
		returnRunner bool
	}{
		{
			desc:         "succeeded",
			getFails:     true,
			getErrorCode: http.StatusNotFound,
			wantErr:      false,
			returnRunner: false,
		},
		{
			desc:         "get fails",
			getFails:     true,
			getErrorCode: http.StatusInternalServerError,
			wantErr:      true,
			returnRunner: false,
		},
		{
			desc:         "timeout",
			getFails:     false,
			wantErr:      true,
			returnRunner: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getRunnerFails: tt.getFails,
					getErrorCode:   tt.getErrorCode,
					returnRunner:   tt.returnRunner,
					intakeRunnerResponse: &intake.IntakeRunnerResponse{ // This is only used in the timeout case
						Id: intakeRunnerId,
					},
				})

				handler := DeleteIntakeRunnerWaitHandler(context.Background(), apiClient, projectId, region, intakeRunnerId)
				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}

func TestCreateOrUpdateIntakeWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		getFails       bool
		getErrorCode   int
		wantErr        bool
		wantResp       bool
		returnIntake   bool
		intakeResponse *intake.IntakeResponse
	}{
		{
			desc:         "succeeded",
			getFails:     false,
			wantErr:      false,
			wantResp:     true,
			returnIntake: true,
			intakeResponse: &intake.IntakeResponse{
				Id:    intakeId,
				State: INTAKERESPONSESTATE_ACTIVE,
			},
		},
		{
			desc:         "failed state",
			getFails:     false,
			wantErr:      true,
			wantResp:     true,
			returnIntake: true,
			intakeResponse: &intake.IntakeResponse{
				Id:    intakeId,
				State: INTAKERESPONSESTATE_FAILED,
			},
		},
		{
			desc:         "get fails",
			getFails:     true,
			getErrorCode: http.StatusInternalServerError,
			wantErr:      true,
			wantResp:     false,
			returnIntake: false,
		},
		{
			desc:         "timeout",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnIntake: true,
			intakeResponse: &intake.IntakeResponse{
				Id:    intakeId,
				State: INTAKERESPONSESTATE_RECONCILING,
			},
		},
		{
			desc:         "nil response",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnIntake: false,
		},
		{
			desc:         "nil id in response",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnIntake: true,
			intakeResponse: &intake.IntakeResponse{
				State: INTAKERESPONSESTATE_RECONCILING,
			},
		},
		{
			desc:         "nil state in response",
			getFails:     false,
			wantErr:      true,
			wantResp:     false,
			returnIntake: true,
			intakeResponse: &intake.IntakeResponse{
				Id: intakeId,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getIntakeFails: tt.getFails,
					getErrorCode:   tt.getErrorCode,
					returnIntake:   tt.returnIntake,
					intakeResponse: tt.intakeResponse,
				})

				var wantResp *intake.IntakeResponse
				if tt.wantResp {
					wantResp = tt.intakeResponse
				}

				handler := CreateOrUpdateIntakeWaitHandler(context.Background(), apiClient, projectId, region, intakeId)
				got, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(got, wantResp) {
					t.Fatalf("handler got = %v, want %v", got, wantResp)
				}
			})
		})
	}
}

func TestDeleteIntakeWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		getFails     bool
		getErrorCode int
		wantErr      bool
		returnIntake bool
	}{
		{
			desc:         "succeeded",
			getFails:     true,
			getErrorCode: http.StatusNotFound,
			wantErr:      false,
			returnIntake: false,
		},
		{
			desc:         "get fails",
			getFails:     true,
			getErrorCode: http.StatusInternalServerError,
			wantErr:      true,
			returnIntake: false,
		},
		{
			desc:         "timeout",
			getFails:     false,
			wantErr:      true,
			returnIntake: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getIntakeFails: tt.getFails,
					getErrorCode:   tt.getErrorCode,
					returnIntake:   tt.returnIntake,
					intakeResponse: &intake.IntakeResponse{
						Id: intakeId,
					},
				})

				handler := DeleteIntakeWaitHandler(context.Background(), apiClient, projectId, region, intakeId)
				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}

func TestCreateOrUpdateIntakeUserWaitHandler(t *testing.T) {
	tests := []struct {
		desc               string
		getFails           bool
		getErrorCode       int
		wantErr            bool
		wantResp           bool
		returnUser         bool
		intakeUserResponse *intake.IntakeUserResponse
	}{
		{
			desc:       "succeeded",
			getFails:   false,
			wantErr:    false,
			wantResp:   true,
			returnUser: true,
			intakeUserResponse: &intake.IntakeUserResponse{
				Id:    intakeUserId,
				State: INTAKEUSERRESPONSESTATE_ACTIVE,
			},
		},
		{
			desc:         "get fails",
			getFails:     true,
			getErrorCode: http.StatusInternalServerError,
			wantErr:      true,
			wantResp:     false,
			returnUser:   false,
		},
		{
			desc:       "timeout",
			getFails:   false,
			wantErr:    true,
			wantResp:   false,
			returnUser: true,
			intakeUserResponse: &intake.IntakeUserResponse{
				Id:    intakeUserId,
				State: INTAKEUSERRESPONSESTATE_RECONCILING,
			},
		},
		{
			desc:       "nil response",
			getFails:   false,
			wantErr:    true,
			wantResp:   false,
			returnUser: false,
		},
		{
			desc:       "nil id in response",
			getFails:   false,
			wantErr:    true,
			wantResp:   false,
			returnUser: true,
			intakeUserResponse: &intake.IntakeUserResponse{
				State: INTAKEUSERRESPONSESTATE_RECONCILING,
			},
		},
		{
			desc:       "nil state in response",
			getFails:   false,
			wantErr:    true,
			wantResp:   false,
			returnUser: true,
			intakeUserResponse: &intake.IntakeUserResponse{
				Id: intakeUserId,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getUserFails:       tt.getFails,
					getErrorCode:       tt.getErrorCode,
					returnUser:         tt.returnUser,
					intakeUserResponse: tt.intakeUserResponse,
				})

				var wantResp *intake.IntakeUserResponse
				if tt.wantResp {
					wantResp = tt.intakeUserResponse
				}

				handler := CreateOrUpdateIntakeUserWaitHandler(context.Background(), apiClient, projectId, region, intakeId, intakeUserId)
				got, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(got, wantResp) {
					t.Fatalf("handler got = %v, want %v", got, wantResp)
				}
			})
		})
	}
}

func TestDeleteIntakeUserWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		getFails     bool
		getErrorCode int
		wantErr      bool
		returnUser   bool
	}{
		{
			desc:         "succeeded",
			getFails:     true,
			getErrorCode: http.StatusNotFound,
			wantErr:      false,
			returnUser:   false,
		},
		{
			desc:         "get fails",
			getFails:     true,
			getErrorCode: http.StatusInternalServerError,
			wantErr:      true,
			returnUser:   false,
		},
		{
			desc:       "timeout",
			getFails:   false,
			wantErr:    true,
			returnUser: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getUserFails: tt.getFails,
					getErrorCode: tt.getErrorCode,
					returnUser:   tt.returnUser,
					intakeUserResponse: &intake.IntakeUserResponse{
						Id: intakeUserId,
					},
				})

				handler := DeleteIntakeUserWaitHandler(context.Background(), apiClient, projectId, region, intakeId, intakeUserId)
				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
