package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	ske "github.com/stackitcloud/stackit-sdk-go/services/ske/v2api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	CREDENTIALSROTATIONSTATEPHASE_NEVER = ske.CREDENTIALSROTATIONSTATEPHASE_NEVER
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	CREDENTIALSROTATIONSTATEPHASE_PREPARING = ske.CREDENTIALSROTATIONSTATEPHASE_PREPARING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	CREDENTIALSROTATIONSTATEPHASE_PREPARED = ske.CREDENTIALSROTATIONSTATEPHASE_PREPARED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	CREDENTIALSROTATIONSTATEPHASE_COMPLETING = ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	CREDENTIALSROTATIONSTATEPHASE_COMPLETED = ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETED

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RUNTIMEERRORCODE_OBSERVABILITY_INSTANCE_NOT_FOUND = ske.RUNTIMEERRORCODE_SKE_OBSERVABILITY_INSTANCE_NOT_FOUND
)

// CreateClusterWaitHandler will wait for cluster creation
func CreateClusterWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[ske.Cluster] {
	return CreateOrUpdateClusterWaitHandler(ctx, a, projectId, region, name)
}

// UpdateClusterWaitHandler will wait for cluster update
func UpdateClusterWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[ske.Cluster] {
	return CreateOrUpdateClusterWaitHandler(ctx, a, projectId, region, name)
}

// Deprecated: Will be removed after 2026-12-08. Use the CreateClusterWaitHandler or UpdateClusterWaitHandler instead.
func CreateOrUpdateClusterWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetCluster(ctx, projectId, region, name).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.Aggregated

		// The state "STATE_UNHEALTHY" (aka "Impaired" in the portal) could be temporarily occur during cluster creation and the system is recovering usually, so it is not considered as a failed state here.
		// -- alignment meeting with SKE team on 4.8.23
		// The exception is when providing an invalid argus instance id, in that case the cluster will stay as "Impaired" until the SKE team solves it, but it is still usable.
		if state == ske.CLUSTERSTATUSSTATE_STATE_UNHEALTHY && s.Status.Error != nil && s.Status.Error.Message != nil && *s.Status.Error.Code == ske.RUNTIMEERRORCODE_SKE_OBSERVABILITY_INSTANCE_NOT_FOUND {
			return true, s, nil
		}

		if state == ske.CLUSTERSTATUSSTATE_STATE_HEALTHY || state == ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED {
			return true, s, nil
		}

		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteClusterWaitHandler will wait for cluster deletion
func DeleteClusterWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, name string) *wait.AsyncActionHandler[ske.ListClustersResponse] {
	handler := wait.New(func() (waitFinished bool, response *ske.ListClustersResponse, err error) {
		s, err := a.ListClusters(ctx, projectId, region).Execute()
		if err != nil {
			return false, nil, err
		}
		items := s.Items
		for i := range items {
			n := items[i].Name
			if n != nil && *n == name {
				return false, nil, nil
			}
		}
		return true, s, nil
	})
	// Timeout was aligned with the service team
	handler.SetTimeout(90 * time.Minute)
	return handler
}

func TriggerClusterHibernationWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		cluster, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *cluster.Status.Aggregated

		if state == ske.CLUSTERSTATUSSTATE_STATE_HIBERNATING {
			return false, nil, nil
		}

		return true, cluster, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func TriggerClusterMaintenanceWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		cluster, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *cluster.Status.Aggregated

		if state == ske.CLUSTERSTATUSSTATE_STATE_RECONCILING {
			return false, nil, nil
		}

		return true, cluster, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func TriggerClusterReconciliationWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		cluster, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *cluster.Status.Aggregated

		if state == ske.CLUSTERSTATUSSTATE_STATE_RECONCILING {
			return false, nil, nil
		}

		return true, cluster, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func TriggerClusterWakeupWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		cluster, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *cluster.Status.Aggregated

		if state == ske.CLUSTERSTATUSSTATE_STATE_WAKINGUP {
			return false, nil, nil
		}

		return true, cluster, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// RotateCredentialsWaitHandler will wait for credentials rotation
func RotateCredentialsWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	waitConfig := wait.WaiterHelper[ske.Cluster, ske.ClusterStatusState]{
		FetchInstance: a.GetCluster(ctx, projectId, region, clusterName).Execute,
		GetState: func(s *ske.Cluster) (ske.ClusterStatusState, error) {
			if s == nil || s.Status == nil || s.Status.Aggregated == nil {
				return "", errors.New("empty response or status")
			}
			return *s.Status.Aggregated, nil
		},
		ActiveState: []ske.ClusterStatusState{
			ske.CLUSTERSTATUSSTATE_STATE_HEALTHY,
			ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED,
		},
		ErrorState: []ske.ClusterStatusState{
			ske.CLUSTERSTATUSSTATE_STATE_UNHEALTHY,
			ske.CLUSTERSTATUSSTATE_STATE_DELETING,
			ske.CLUSTERSTATUSSTATE_STATE_HIBERNATING,
			ske.CLUSTERSTATUSSTATE_STATE_WAKINGUP,
		},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// StartCredentialsRotationWaitHandler will wait for credentials rotation
func StartCredentialsRotationWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	waitConfig := wait.WaiterHelper[ske.Cluster, ske.CredentialsRotationStatePhase]{
		FetchInstance: a.GetCluster(ctx, projectId, region, clusterName).Execute,
		GetState: func(s *ske.Cluster) (ske.CredentialsRotationStatePhase, error) {
			if s == nil || s.Status == nil || s.Status.CredentialsRotation == nil || s.Status.CredentialsRotation.Phase == nil {
				return "", errors.New("empty response or credentials rotation phase")
			}
			return *s.Status.CredentialsRotation.Phase, nil
		},
		ActiveState: []ske.CredentialsRotationStatePhase{ske.CREDENTIALSROTATIONSTATEPHASE_PREPARED},
		ErrorState: []ske.CredentialsRotationStatePhase{
			ske.CREDENTIALSROTATIONSTATEPHASE_NEVER,
			ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETING,
			ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETED,
		},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CompleteCredentialsRotationWaitHandler will wait for credentials rotation
func CompleteCredentialsRotationWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	waitConfig := wait.WaiterHelper[ske.Cluster, ske.CredentialsRotationStatePhase]{
		FetchInstance: a.GetCluster(ctx, projectId, region, clusterName).Execute,
		GetState: func(s *ske.Cluster) (ske.CredentialsRotationStatePhase, error) {
			if s == nil || s.Status == nil || s.Status.CredentialsRotation == nil || s.Status.CredentialsRotation.Phase == nil {
				return "", errors.New("empty response or credentials rotation phase")
			}
			return *s.Status.CredentialsRotation.Phase, nil
		},
		ActiveState: []ske.CredentialsRotationStatePhase{ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETED},
		ErrorState: []ske.CredentialsRotationStatePhase{
			ske.CREDENTIALSROTATIONSTATEPHASE_NEVER,
			ske.CREDENTIALSROTATIONSTATEPHASE_PREPARING,
			ske.CREDENTIALSROTATIONSTATEPHASE_PREPARED,
		},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}
