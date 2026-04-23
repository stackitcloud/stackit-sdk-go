package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/uuid"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	scf "github.com/stackitcloud/stackit-sdk-go/services/scf/v1api"
)

var PROJECT_ID = uuid.New().String()
var INSTANCE_ID = uuid.New().String()

const REGION = "eu01"

type mockSettings struct {
	getFails       bool
	errorCode      int
	returnInstance bool
	projectId      string
	instanceId     string
	getSCFResponse *scf.Organization
}

func newAPIMock(settings *mockSettings) scf.DefaultAPI {
	return &scf.DefaultAPIServiceMock{
		GetOrganizationExecuteMock: utils.Ptr(func(_ scf.ApiGetOrganizationRequest) (*scf.Organization, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.errorCode,
				}
			}
			if !settings.returnInstance {
				return nil, nil
			}
			return settings.getSCFResponse, nil
		}),
	}
}

func TestDeleteOrganizationWaitHandler(t *testing.T) {
	statusDeletingFailed := "deleting_failed"
	tests := []struct {
		desc                 string
		wantErr              bool
		wantReturnedInstance bool
		getFails             bool
		errorCode            int
		returnInstance       bool
		getOrgResponse       *scf.Organization
	}{
		{
			desc:     "Instance deletion failed with error",
			wantErr:  true,
			getFails: true,
		},
		{
			desc:      "Instance is not found",
			wantErr:   false,
			getFails:  true,
			errorCode: 404,
		},
		{
			desc:           "Instance is in error state",
			wantErr:        true,
			returnInstance: true,
			getOrgResponse: &scf.Organization{
				Status: statusDeletingFailed,
			},
		},
		{
			desc:           "Instance is nil",
			wantErr:        true,
			returnInstance: true,
			getOrgResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					projectId:      PROJECT_ID,
					instanceId:     INSTANCE_ID,
					getFails:       tt.getFails,
					errorCode:      tt.errorCode,
					returnInstance: tt.returnInstance,
					getSCFResponse: tt.getOrgResponse,
				})

				handler := DeleteOrganizationWaitHandler(context.Background(), apiClient, PROJECT_ID, REGION, INSTANCE_ID)
				response, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if (response != nil) != tt.wantReturnedInstance {
					t.Fatalf("handler gotRes = %v, want nil", response)
				}
			})
		})
	}
}
