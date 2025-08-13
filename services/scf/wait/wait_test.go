package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/services/scf"
)

var PROJECT_ID = uuid.New().String()
var INSTANCE_ID = uuid.New().String()

const REGION = "eu01"

type apiClientMocked struct {
	getFails       bool
	errorCode      int
	returnInstance bool
	projectId      string
	instanceId     string
	getGitResponse *scf.Organization
}

func (a *apiClientMocked) GetOrganizationExecute(_ context.Context, _, _, _ string) (*scf.Organization, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: a.errorCode,
		}
	}
	if !a.returnInstance {
		return nil, nil
	}
	return a.getGitResponse, nil
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
				Status: &statusDeletingFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				projectId:      PROJECT_ID,
				instanceId:     INSTANCE_ID,
				getFails:       tt.getFails,
				errorCode:      tt.errorCode,
				returnInstance: tt.returnInstance,
				getGitResponse: tt.getOrgResponse,
			}

			handler := DeleteOrganizationWaitHandler(context.Background(), apiClient, apiClient.projectId, REGION, apiClient.instanceId)
			response, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if (response != nil) != tt.wantReturnedInstance {
				t.Fatalf("handler gotRes = %v, want nil", response)
			}
		})
	}
}
