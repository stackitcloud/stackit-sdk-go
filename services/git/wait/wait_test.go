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
	desc           string
	getFails       bool
	returnInstance bool
	wantErr        bool
	created        time.Time
	id             string
	projectId      string
	name           string
	state          string
	url            string
	version        string
}

func (a *apiClientMocked) GetGitExecute(_ context.Context, _, _ string) (*git.Instance, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	if !a.returnInstance {
		return nil, nil
	}
	return &git.Instance{
		Created: utils.Ptr(a.created),
		Id:      utils.Ptr(a.id),
		Name:    utils.Ptr(a.name),
		State:   utils.Ptr(a.state),
		Url:     utils.Ptr(a.url),
		Version: utils.Ptr(a.version),
	}, nil
}

func TestCreateGitInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		getFails       bool
		wantErr        bool
		wantResp       bool
		created        time.Time
		id             string
		projectId      string
		name           string
		state          string
		url            string
		version        string
		returnInstance bool
	}{
		{
			desc:           "Creation of an instance succeeded",
			getFails:       false,
			wantErr:        false,
			wantResp:       true,
			created:        time.Now(),
			id:             uuid.New().String(),
			projectId:      uuid.New().String(),
			name:           "instance-test",
			state:          CreateSuccess,
			url:            "https://testing.git.onstackit.cloud",
			version:        "v1.6.0",
			returnInstance: true,
		},
		{
			desc:           "Creation of an instance Failed With Error",
			getFails:       true,
			wantErr:        true,
			wantResp:       false,
			created:        time.Now(),
			id:             uuid.New().String(),
			projectId:      uuid.New().String(),
			name:           "instance-test",
			state:          CreateFail,
			url:            "https://testing.git.onstackit.cloud",
			version:        "v1.6.0",
			returnInstance: true,
		},
		{
			desc:           "Creation of an instance with response failed and without error",
			getFails:       false,
			wantErr:        true,
			wantResp:       true,
			created:        time.Now(),
			id:             uuid.New().String(),
			projectId:      uuid.New().String(),
			name:           "instance-test",
			state:          CreateFail,
			url:            "https://testing.git.onstackit.cloud",
			version:        "v1.6.0",
			returnInstance: true,
		},
		{
			desc:           "Creation of an instance failed without id on the response",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			created:        time.Now(),
			projectId:      uuid.New().String(),
			name:           "instance-test",
			state:          CreateFail,
			url:            "https://testing.git.onstackit.cloud",
			version:        "v1.6.0",
			returnInstance: true,
		},

		{
			desc:           "Creation of an instance without state on the response",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			created:        time.Now(),
			id:             uuid.New().String(),
			projectId:      uuid.New().String(),
			name:           "instance-test",
			url:            "https://testing.git.onstackit.cloud",
			version:        "v1.6.0",
			returnInstance: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				desc:           tt.desc,
				getFails:       tt.getFails,
				wantErr:        tt.wantErr,
				created:        tt.created,
				id:             tt.id,
				projectId:      tt.projectId,
				name:           tt.name,
				state:          tt.state,
				url:            tt.url,
				version:        tt.version,
				returnInstance: tt.returnInstance,
			}
			var instanceWanted *git.Instance
			if tt.wantResp {
				instanceWanted = &git.Instance{
					Created: utils.Ptr(tt.created),
					Id:      utils.Ptr(tt.id),
					Name:    utils.Ptr(tt.name),
					State:   utils.Ptr(tt.state),
					Url:     utils.Ptr(tt.url),
					Version: utils.Ptr(tt.version),
				}
			}

			handler := CreateGitInstanceWaitHandler(context.Background(), apiClient, apiClient.projectId, apiClient.id)

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
		created              time.Time
		id                   string
		name                 string
		state                string
		url                  string
		version              string
	}{
		{
			desc:     "Instance deletion failed with error",
			wantErr:  true,
			getFails: true,
		},
		{
			desc:                 "Instance deletion failed returning existing instance",
			wantErr:              false,
			getFails:             false,
			wantReturnedInstance: true,
			created:              time.Now(),
			id:                   uuid.New().String(),
			name:                 "instance-test",
			state:                CreateSuccess,
			returnInstance:       true,
			url:                  "https://testing.git.onstackit.cloud",
			version:              "v1.6.0",
		},
		{
			desc:                 "Instance deletion succesfull",
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
				created:        tt.created,
				id:             tt.id,
				name:           tt.name,
				state:          tt.state,
				url:            tt.url,
				version:        tt.version,
				returnInstance: tt.returnInstance,
			}

			handler := DeleteGitInstanceWaitHandler(context.Background(), apiClient, apiClient.projectId, apiClient.id)
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
