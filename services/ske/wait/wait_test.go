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
	resourceState        ske.ClusterStatusState
	invalidArgusInstance bool
}

func (a *apiClientClusterMocked) GetClusterExecute(_ context.Context, _, _ string) (*ske.Cluster, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	rs := ske.ClusterStatusState(a.resourceState)

	if a.invalidArgusInstance {
		return &ske.Cluster{
			Name: utils.Ptr("cluster"),
			Status: &ske.ClusterStatus{
				Aggregated: &rs,
				Error: &ske.RuntimeError{
					Code:    utils.Ptr(ske.RUNTIMEERRORCODE_ARGUS_INSTANCE_NOT_FOUND),
					Message: utils.Ptr("invalid argus instance"),
				},
			},
		}, nil
	}
	return &ske.Cluster{
		Name: utils.Ptr("cluster"),
		Status: &ske.ClusterStatus{
			Aggregated: &rs,
		},
	}, nil
}

func (a *apiClientClusterMocked) ListClustersExecute(_ context.Context, _ string) (*ske.ListClustersResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}
	rs := ske.ClusterStatusState(a.resourceState)
	return &ske.ListClustersResponse{
		Items: &[]ske.Cluster{
			{
				Name: utils.Ptr("cluster"),
				Status: &ske.ClusterStatus{
					Aggregated: &rs,
				},
			},
		},
	}, nil
}

func TestCreateOrUpdateClusterWaitHandler(t *testing.T) {
	tests := []struct {
		desc                 string
		getFails             bool
		resourceState        ske.ClusterStatusState
		invalidArgusInstance bool
		wantErr              bool
		wantResp             bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_HIBERNATED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:                 "unhealthy_cluster",
			getFails:             false,
			resourceState:        ske.CLUSTERSTATUSSTATE_UNHEALTHY,
			invalidArgusInstance: true,
			wantErr:              false,
			wantResp:             true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: StateFailed,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
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
			name := "cluster"

			apiClient := &apiClientClusterMocked{
				getFails:             tt.getFails,
				name:                 name,
				resourceState:        tt.resourceState,
				invalidArgusInstance: tt.invalidArgusInstance,
			}
			var wantRes *ske.Cluster
			rs := ske.ClusterStatusState(tt.resourceState)
			if tt.wantResp {
				wantRes = &ske.Cluster{
					Name: &name,
					Status: &ske.ClusterStatus{
						Aggregated: &rs,
					},
				}

				if tt.invalidArgusInstance {
					wantRes.Status.Error = &ske.RuntimeError{
						Code:    utils.Ptr(ske.RUNTIMEERRORCODE_ARGUS_INSTANCE_NOT_FOUND),
						Message: utils.Ptr("invalid argus instance"),
					}
				}
			}

			handler := CreateOrUpdateClusterWaitHandler(context.Background(), apiClient, "", name)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
			}
		})
	}
}

func TestRotateCredentialsWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState ske.ClusterStatusState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "reconciliation_succeeded_1",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "reconciliation_succeeded_2",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_HIBERNATED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "reconciliation_failed",
			getFails:      false,
			resourceState: StateFailed,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_RECONCILING,
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			name := "cluster"

			apiClient := &apiClientClusterMocked{
				getFails:      tt.getFails,
				name:          name,
				resourceState: tt.resourceState,
			}
			var wantRes *ske.Cluster
			if tt.wantResp {
				wantRes = &ske.Cluster{
					Name: &name,
					Status: &ske.ClusterStatus{
						Aggregated: &tt.resourceState,
					},
				}
			}

			handler := RotateCredentialsWaitHandler(context.Background(), apiClient, "", name)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
			}
		})
	}
}
