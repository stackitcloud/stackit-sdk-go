package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	ske "github.com/stackitcloud/stackit-sdk-go/services/ske/v2api"
)

const (
	CREDENTIALSROTATIONSTATEPHASE_NEVER      = "NEVER"
	CREDENTIALSROTATIONSTATEPHASE_PREPARING  = "PREPARING"
	CREDENTIALSROTATIONSTATEPHASE_PREPARED   = "PREPARED"
	CREDENTIALSROTATIONSTATEPHASE_COMPLETING = "COMPLETING"
	CREDENTIALSROTATIONSTATEPHASE_COMPLETED  = "COMPLETED"

	RUNTIMEERRORCODE_OBSERVABILITY_INSTANCE_NOT_FOUND = "SKE_OBSERVABILITY_INSTANCE_NOT_FOUND"
)

// CreateOrUpdateClusterWaitHandler will wait for cluster creation or update
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
		if state == ske.CLUSTERSTATUSSTATE_STATE_UNHEALTHY && s.Status.Error != nil && s.Status.Error.Message != nil && *s.Status.Error.Code == RUNTIMEERRORCODE_OBSERVABILITY_INSTANCE_NOT_FOUND {
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
	handler.SetTimeout(45 * time.Minute)
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
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.Aggregated

		if state == ske.CLUSTERSTATUSSTATE_STATE_HEALTHY || state == ske.CLUSTERSTATUSSTATE_STATE_HIBERNATED {
			return true, s, nil
		}

		if state == ske.CLUSTERSTATUSSTATE_STATE_RECONCILING {
			return false, nil, nil
		}

		return true, s, fmt.Errorf("unexpected state %s while waiting for cluster reconciliation", state)
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

// StartCredentialsRotationWaitHandler will wait for credentials rotation
func StartCredentialsRotationWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.CredentialsRotation.Phase

		if state == CREDENTIALSROTATIONSTATEPHASE_PREPARED {
			return true, s, nil
		}

		if state == CREDENTIALSROTATIONSTATEPHASE_PREPARING {
			return false, nil, nil
		}

		return true, s, fmt.Errorf("unexpected status %s while waiting for cluster credentials rotation to be prepared", state)
	})

	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CompleteCredentialsRotationWaitHandler will wait for credentials rotation
func CompleteCredentialsRotationWaitHandler(ctx context.Context, a ske.DefaultAPI, projectId, region, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetCluster(ctx, projectId, region, clusterName).Execute()
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.CredentialsRotation.Phase

		if state == CREDENTIALSROTATIONSTATEPHASE_COMPLETED {
			return true, s, nil
		}

		if state == CREDENTIALSROTATIONSTATEPHASE_COMPLETING {
			return false, nil, nil
		}

		return true, s, fmt.Errorf("unexpected status %s while waiting for cluster credentials rotation to be completed", state)
	})

	handler.SetTimeout(45 * time.Minute)
	return handler
}
