package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	ske "github.com/stackitcloud/stackit-sdk-go/services/ske/v2api"
)

const testRegion = "eu01"

type mockSettings struct {
	getFails             bool
	name                 string
	resourceState        ske.ClusterStatusState
	invalidArgusInstance bool
}

// Used for testing cluster operations
func newAPIMock(settings mockSettings) ske.DefaultAPI {
	return &ske.DefaultAPIServiceMock{
		GetClusterExecuteMock: utils.Ptr(func(_ ske.ApiGetClusterRequest) (*ske.Cluster, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusInternalServerError,
				}
			}
			rs := settings.resourceState

			if settings.invalidArgusInstance {
				return &ske.Cluster{
					Name: utils.Ptr("cluster"),
					Status: &ske.ClusterStatus{
						Aggregated: &rs,
						Error: &ske.RuntimeError{
							Code:    utils.Ptr(RUNTIMEERRORCODE_OBSERVABILITY_INSTANCE_NOT_FOUND),
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
		}),
		ListClustersExecuteMock: utils.Ptr(func(_ ske.ApiListClustersRequest) (*ske.ListClustersResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusInternalServerError,
				}
			}
			rs := settings.resourceState
			return &ske.ListClustersResponse{
				Items: []ske.Cluster{
					{
						Name: utils.Ptr("cluster"),
						Status: &ske.ClusterStatus{
							Aggregated: &rs,
						},
					},
				},
			}, nil
		}),
	}
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
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:                 "unhealthy_cluster",
			getFails:             false,
			resourceState:        ske.CLUSTERSTATUSSTATE_STATE_UNHEALTHY,
			invalidArgusInstance: true,
			wantErr:              false,
			wantResp:             true,
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
			synctest.Test(t, func(t *testing.T) {
				name := "cluster"

				apiClient := newAPIMock(mockSettings{
					getFails:             tt.getFails,
					name:                 name,
					resourceState:        tt.resourceState,
					invalidArgusInstance: tt.invalidArgusInstance,
				})

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
							Code:    utils.Ptr(RUNTIMEERRORCODE_OBSERVABILITY_INSTANCE_NOT_FOUND),
							Message: utils.Ptr("invalid argus instance"),
						}
					}
				}

				handler := CreateOrUpdateClusterWaitHandler(context.Background(), apiClient, "", testRegion, name)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestTriggerClusterHibernationWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState ske.ClusterStatusState
		wantErr       bool
		wantResp      bool
	}{
		{
			description:   "hibernation ongoing (timeout)",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HIBERNATING,
			wantErr:       true,
			wantResp:      false,
		},
		{
			description:   "hibernation succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description:   "unexpected status",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description: "get_fails",
			getFails:    true,
			wantErr:     true,
			wantResp:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				name := "cluster"

				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					name:          name,
					resourceState: tt.resourceState,
				})

				var wantRes *ske.Cluster
				if tt.wantResp {
					wantRes = &ske.Cluster{
						Name: &name,
						Status: &ske.ClusterStatus{
							Aggregated: utils.Ptr(tt.resourceState),
						},
					}
				}

				handler := TriggerClusterHibernationWaitHandler(context.Background(), apiClient, "", testRegion, name)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestTriggerClusterMaintenanceWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState ske.ClusterStatusState
		wantErr       bool
		wantResp      bool
	}{
		{
			description:   "maintenance ongoing (timeout)",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_RECONCILING,
			wantErr:       true,
			wantResp:      false,
		},
		{
			description:   "maintenance succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description:   "unexpected status",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description: "get_fails",
			getFails:    true,
			wantErr:     true,
			wantResp:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				name := "cluster"

				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					name:          name,
					resourceState: tt.resourceState,
				})

				var wantRes *ske.Cluster
				if tt.wantResp {
					wantRes = &ske.Cluster{
						Name: &name,
						Status: &ske.ClusterStatus{
							Aggregated: utils.Ptr(tt.resourceState),
						},
					}
				}

				handler := TriggerClusterMaintenanceWaitHandler(context.Background(), apiClient, "", testRegion, name)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestTriggerClusterWakeupWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState ske.ClusterStatusState
		wantErr       bool
		wantResp      bool
	}{
		{
			description:   "wakeup ongoing (timeout)",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_WAKINGUP,
			wantErr:       true,
			wantResp:      false,
		},
		{
			description:   "wakeup succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description:   "unexpected status",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_DELETING,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description: "get_fails",
			getFails:    true,
			wantErr:     true,
			wantResp:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				name := "cluster"

				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					name:          name,
					resourceState: tt.resourceState,
				})

				var wantRes *ske.Cluster
				if tt.wantResp {
					wantRes = &ske.Cluster{
						Name: &name,
						Status: &ske.ClusterStatus{
							Aggregated: utils.Ptr(tt.resourceState),
						},
					}
				}

				handler := TriggerClusterWakeupWaitHandler(context.Background(), apiClient, "", testRegion, name)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestTriggerClusterReconciliationWaitHandler(t *testing.T) {
	tests := []struct {
		description   string
		getFails      bool
		resourceState ske.ClusterStatusState
		wantErr       bool
		wantResp      bool
	}{
		{
			description:   "reconciliation ongoing (timeout)",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_RECONCILING,
			wantErr:       true,
			wantResp:      false,
		},
		{
			description:   "reconciliation succeeded",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description:   "unexpected status",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_CREATING,
			wantErr:       false,
			wantResp:      true,
		},
		{
			description: "get_fails",
			getFails:    true,
			wantErr:     true,
			wantResp:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				name := "cluster"

				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					name:          name,
					resourceState: tt.resourceState,
				})

				var wantRes *ske.Cluster
				if tt.wantResp {
					wantRes = &ske.Cluster{
						Name: &name,
						Status: &ske.ClusterStatus{
							Aggregated: utils.Ptr(tt.resourceState),
						},
					}
				}

				handler := TriggerClusterReconciliationWaitHandler(context.Background(), apiClient, "", testRegion, name)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
				}
			})
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
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "reconciliation_succeeded_2",
			getFails:      false,
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED,
			wantErr:       false,
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
			resourceState: ske.CLUSTERSTATUSSTATE_STATE_RECONCILING,
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				name := "cluster"

				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					name:          name,
					resourceState: tt.resourceState,
				})

				var wantRes *ske.Cluster
				if tt.wantResp {
					wantRes = &ske.Cluster{
						Name: &name,
						Status: &ske.ClusterStatus{
							Aggregated: utils.Ptr(tt.resourceState),
						},
					}
				}

				handler := RotateCredentialsWaitHandler(context.Background(), apiClient, "", testRegion, name)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %+v, want %+v", gotRes, wantRes)
				}
			})
		})
	}
}
