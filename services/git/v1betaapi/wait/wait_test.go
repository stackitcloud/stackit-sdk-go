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
	git "github.com/stackitcloud/stackit-sdk-go/services/git/v1betaapi"
)

type mockSettings struct {
	getFails       bool
	errorCode      int
	returnInstance bool
	projectId      string
	instanceId     string
	getGitResponse *git.Instance
}

func newAPIMock(settings *mockSettings) git.DefaultAPI {
	return &git.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ git.ApiGetInstanceRequest) (*git.Instance, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.errorCode,
				}
			}

			if !settings.returnInstance {
				return nil, nil
			}

			return settings.getGitResponse, nil
		}),
	}
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
				Created: time.Now(),
				Id:      INSTANCE_ID,
				Name:    "instance-test",
				State:   INSTANCESTATE_READY,
				Url:     "https://testing.git.onstackit.cloud",
				Version: "v1.6.0",
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
				Created: time.Now(),
				Id:      INSTANCE_ID,
				Name:    "instance-test",
				State:   INSTANCESTATE_READY,
				Url:     "https://testing.git.onstackit.cloud",
				Version: "v1.6.0",
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
				Created: time.Now(),
				Id:      INSTANCE_ID,
				Name:    "instance-test",
				State:   INSTANCESTATE_ERROR,
				Url:     "https://testing.git.onstackit.cloud",
				Version: "v1.6.0",
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
				Created: time.Now(),
				Name:    "instance-test",
				State:   INSTANCESTATE_ERROR,
				Url:     "https://testing.git.onstackit.cloud",
				Version: "v1.6.0",
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
				Created: time.Now(),
				Id:      INSTANCE_ID,
				Name:    "instance-test",
				Url:     "https://testing.git.onstackit.cloud",
				Version: "v1.6.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := newAPIMock(&mockSettings{
				getFails:       tt.getFails,
				projectId:      tt.projectId,
				instanceId:     tt.instanceId,
				getGitResponse: tt.getGitResponse,
				returnInstance: tt.returnInstance,
			})

			var instanceWanted *git.Instance
			if tt.wantResp {
				instanceWanted = tt.getGitResponse
			}

			handler := CreateGitInstanceWaitHandler(context.Background(), apiClient, tt.projectId, tt.instanceId)

			response, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}

			if !cmp.Equal(response, instanceWanted, cmp.AllowUnexported(git.NullableString{}, git.NullableBool{})) {
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
		errorCode            int
		returnInstance       bool
		getGitResponse       *git.Instance
	}{
		{
			desc:     "Instance deletion failed with error",
			wantErr:  true,
			getFails: true,
		},
		{
			desc:     "Instance deletion failed returning existing instance",
			wantErr:  true,
			getFails: false,

			wantReturnedInstance: false,
			returnInstance:       true,
			getGitResponse: &git.Instance{
				Created: time.Now(),
				Id:      INSTANCE_ID,
				Name:    "instance-test",
				State:   INSTANCESTATE_READY,
				Url:     "https://testing.git.onstackit.cloud",
				Version: "v1.6.0",
			},
		},
		{
			desc:                 "Instance deletion successful",
			wantErr:              false,
			getFails:             true,
			errorCode:            http.StatusNotFound,
			wantReturnedInstance: false,
			returnInstance:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			projectId := uuid.New().String()
			instanceId := uuid.New().String()

			apiClient := newAPIMock(&mockSettings{
				projectId:      projectId,
				getFails:       tt.getFails,
				errorCode:      tt.errorCode,
				returnInstance: tt.returnInstance,
				getGitResponse: tt.getGitResponse,
			})

			handler := DeleteGitInstanceWaitHandler(context.Background(), apiClient, projectId, instanceId)
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
