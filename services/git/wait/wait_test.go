package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/git"
)

type apiClientMocked struct {
	getFails       bool
	returnInstance bool
	projectId      string
	instanceId     string
	getGitResponse *git.Instance
}

func (a *apiClientMocked) GetInstanceExecute(_ context.Context, _, _ string) (*git.Instance, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	if !a.returnInstance {
		return nil, nil
	}
	return a.getGitResponse, nil
}

var PROJECT_ID = uuid.New().String()
var INSTANCE_ID = uuid.New().String()

func TestCreateGitInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		getFails       bool
		wantErr        bool
		wantResp       bool
		projectId      string
		instanceId     string
		returnInstance bool
		getGitResponse *git.Instance
	}{
		{
			desc:           "Creation of an instance succeeded",
			getFails:       false,
			wantErr:        false,
			wantResp:       true,
			projectId:      uuid.New().String(),
			instanceId:     INSTANCE_ID,
			returnInstance: true,
			getGitResponse: &git.Instance{
				Created: utils.Ptr(time.Now()),
				Id:      utils.Ptr(INSTANCE_ID),
				Name:    utils.Ptr("instance-test"),
				State:   utils.Ptr(InstanceStateReady),
				Url:     utils.Ptr("https://testing.git.onstackit.cloud"),
				Version: utils.Ptr("v1.6.0"),
			},
		},
		{
			desc:           "Creation of an instance failed with error",
			getFails:       true,
			wantErr:        true,
			wantResp:       false,
			projectId:      uuid.New().String(),
			instanceId:     INSTANCE_ID,
			returnInstance: true,
			getGitResponse: &git.Instance{
				Created: utils.Ptr(time.Now()),
				Id:      utils.Ptr(INSTANCE_ID),
				Name:    utils.Ptr("instance-test"),
				State:   utils.Ptr(InstanceStateReady),
				Url:     utils.Ptr("https://testing.git.onstackit.cloud"),
				Version: utils.Ptr("v1.6.0"),
			},
		},
		{
			desc:           "Creation of an instance with response failed and without error",
			getFails:       false,
			wantErr:        true,
			wantResp:       true,
			projectId:      uuid.New().String(),
			instanceId:     INSTANCE_ID,
			returnInstance: true,
			getGitResponse: &git.Instance{
				Created: utils.Ptr(time.Now()),
				Id:      utils.Ptr(INSTANCE_ID),
				Name:    utils.Ptr("instance-test"),
				State:   utils.Ptr(InstanceStateError),
				Url:     utils.Ptr("https://testing.git.onstackit.cloud"),
				Version: utils.Ptr("v1.6.0"),
			},
		},
		{
			desc:           "Creation of an instance failed without id on the response",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			projectId:      PROJECT_ID,
			instanceId:     INSTANCE_ID,
			returnInstance: true,
			getGitResponse: &git.Instance{
				Created: utils.Ptr(time.Now()),
				Name:    utils.Ptr("instance-test"),
				State:   utils.Ptr(InstanceStateError),
				Url:     utils.Ptr("https://testing.git.onstackit.cloud"),
				Version: utils.Ptr("v1.6.0"),
			},
		},

		{
			desc:           "Creation of an instance without state on the response",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			projectId:      PROJECT_ID,
			instanceId:     INSTANCE_ID,
			returnInstance: true,
			getGitResponse: &git.Instance{
				Created: utils.Ptr(time.Now()),
				Id:      utils.Ptr(INSTANCE_ID),
				Name:    utils.Ptr("instance-test"),
				Url:     utils.Ptr("https://testing.git.onstackit.cloud"),
				Version: utils.Ptr("v1.6.0"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:       tt.getFails,
				projectId:      tt.projectId,
				instanceId:     tt.instanceId,
				getGitResponse: tt.getGitResponse,
				returnInstance: tt.returnInstance,
			}
			var instanceWanted *git.Instance
			if tt.wantResp {
				instanceWanted = tt.getGitResponse
			}

			handler := CreateGitInstanceWaitHandler(context.Background(), apiClient, apiClient.projectId, apiClient.instanceId)

			response, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(response, instanceWanted) {
				t.Fatalf("handler gotRes = %v, want %v", response, instanceWanted)
			}
		})
	}
}

func TestDeleteGitInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                 string
		wantErr              bool
		wantReturnedInstance bool
		getFails             bool
		returnInstance       bool
		getGitResponse       *git.Instance
	}{
		{
			desc:     "Instance deletion failed with error",
			wantErr:  true,
			getFails: true,
		},
		{
			desc:                 "Instance deletion failed returning existing instance",
			wantErr:              true,
			getFails:             false,
			wantReturnedInstance: true,
			returnInstance:       true,
			getGitResponse: &git.Instance{
				Created: utils.Ptr(time.Now()),
				Id:      utils.Ptr(INSTANCE_ID),
				Name:    utils.Ptr("instance-test"),
				State:   utils.Ptr(InstanceStateReady),
				Url:     utils.Ptr("https://testing.git.onstackit.cloud"),
				Version: utils.Ptr("v1.6.0"),
			},
		},
		{
			desc:                 "Instance deletion succesful",
			wantErr:              false,
			getFails:             false,
			wantReturnedInstance: false,
			returnInstance:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{

				projectId:      uuid.New().String(),
				getFails:       tt.getFails,
				returnInstance: tt.returnInstance,
				getGitResponse: tt.getGitResponse,
			}

			handler := DeleteGitInstanceWaitHandler(context.Background(), apiClient, apiClient.projectId, apiClient.instanceId)
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
