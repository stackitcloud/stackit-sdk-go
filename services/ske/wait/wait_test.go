package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/ske"
)

// Used for testing cluster operations
type apiClientClusterMocked struct {
	getFails             bool
	name                 string
	resourceState        string
	invalidArgusInstance bool
}

func (a *apiClientClusterMocked) GetClusterExecute(_ context.Context, _, _ string) (*ske.ClusterResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	rs := ske.ClusterStatusState(a.resourceState)

	if a.invalidArgusInstance {
		return &ske.ClusterResponse{
			Name: utils.Ptr("cluster"),
			Status: &ske.ClusterStatus{
				Aggregated: &rs,
				Error: &ske.RuntimeError{
					Code:    utils.Ptr(string(InvalidArgusInstanceErrorCode)),
					Message: utils.Ptr("invalid argus instance"),
				},
			},
		}, nil
	}
	return &ske.ClusterResponse{
		Name: utils.Ptr("cluster"),
		Status: &ske.ClusterStatus{
			Aggregated: &rs,
		},
	}, nil
}

func (a *apiClientClusterMocked) GetClustersExecute(_ context.Context, _ string) (*ske.ClustersResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	rs := ske.ClusterStatusState(a.resourceState)
	return &ske.ClustersResponse{
		Items: &[]ske.ClusterResponse{
			{
				Name: utils.Ptr("cluster"),
				Status: &ske.ClusterStatus{
					Aggregated: &rs,
				},
			},
		},
	}, nil
}

// Used for testing cluster operations
type apiClientProjectMocked struct {
	getFails      bool
	getNotFound   bool
	resourceState string
}

func (a *apiClientProjectMocked) GetProjectExecute(_ context.Context, _ string) (*ske.ProjectResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	if a.getNotFound {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusNotFound,
		}
	}
	rs := ske.ProjectState(a.resourceState)
	return &ske.ProjectResponse{
		ProjectId: utils.Ptr("pid"),
		State:     &rs,
	}, nil
}

func TestCreateOrUpdateClusterWaitHandler(t *testing.T) {
	tests := []struct {
		desc                 string
		getFails             bool
		resourceState        string
		invalidArgusInstance bool
		wantErr              bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: StateHealthy,
			wantErr:       false,
		},
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: StateHibernated,
			wantErr:       false,
		},
		{
			desc:                 "unhealthy_cluster",
			getFails:             false,
			resourceState:        StateUnhealthy,
			invalidArgusInstance: true,
			wantErr:              false,
		},
		{
			desc:     "create_failed",
			getFails: false,
			wantErr:  true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			name := "cluster"

			apiClient := &apiClientClusterMocked{
				getFails:             tt.getFails,
				name:                 name,
				resourceState:        tt.resourceState,
				invalidArgusInstance: tt.invalidArgusInstance,
			}
			var wantRes *ske.ClusterResponse
			rs := ske.ClusterStatusState(tt.resourceState)
			if !tt.getFails {
				wantRes = &ske.ClusterResponse{
					Name: &name,
					Status: &ske.ClusterStatus{
						Aggregated: &rs,
					},
				}
			} else {
				wantRes = nil
			}

			if tt.invalidArgusInstance {
				wantRes.Status.Error = &ske.RuntimeError{
					Code:    utils.Ptr(string(InvalidArgusInstanceErrorCode)),
					Message: utils.Ptr("invalid argus instance"),
				}
			}

			handler := CreateOrUpdateClusterWaitHandler(context.Background(), apiClient, "", name)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
			}
		})
	}
}

func TestCreateProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: StateCreated,
			wantErr:       false,
		},
		{
			desc:     "create_failed",
			getFails: false,
			wantErr:  true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientProjectMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}
			var wantRes *ske.ProjectResponse
			rs := ske.ProjectState(tt.resourceState)
			if !tt.getFails {
				wantRes = &ske.ProjectResponse{
					ProjectId: utils.Ptr("pid"),
					State:     &rs,
				}
			}

			handler := CreateProjectWaitHandler(context.Background(), apiClient, "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		getNotFound   bool
		wantErr       bool
		resourceState string
	}{
		{
			desc:        "delete_succeeded",
			getFails:    false,
			getNotFound: true,
			wantErr:     false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			wantErr:       true,
			resourceState: "ANOTHER STATE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientProjectMocked{
				getFails:      tt.getFails,
				getNotFound:   tt.getNotFound,
				resourceState: tt.resourceState,
			}

			var wantRes *ske.ProjectResponse
			if !tt.getFails && !tt.getNotFound {
				rs := ske.ProjectState(tt.resourceState)
				wantRes = &ske.ProjectResponse{
					ProjectId: utils.Ptr("pid"),
					State:     &rs,
				}
			} else {
				wantRes = nil
			}

			handler := DeleteProjectWaitHandler(context.Background(), apiClient, "")

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
	}
}
