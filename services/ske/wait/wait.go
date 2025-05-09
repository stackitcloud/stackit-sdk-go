package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/ske"
)

const (
	// Deprecated: StateHealthy is deprecated and will be removed after 9th November 2025. Use [ske.CLUSTERSTATUSSTATE_HEALTHY] instead.
	StateHealthy = "STATE_HEALTHY"
	// Deprecated: StateHibernated is deprecated and will be removed after 9th November 2025. Use [ske.CLUSTERSTATUSSTATE_HIBERNATED] instead.
	StateHibernated = "STATE_HIBERNATED"
	StateFailed     = "STATE_FAILED"
	// Deprecated: StateDeleting is deprecated and will be removed after 9th November 2025. Use [ske.CLUSTERSTATUSSTATE_DELETING] instead.
	StateDeleting = "STATE_DELETING"
	// Deprecated: StateCreated is deprecated and will be removed after 9th November 2025.
	StateCreated = "STATE_CREATED"
	// Deprecated: StateUnhealthy is deprecated and will be removed after 9th November 2025. Use [ske.CLUSTERSTATUSSTATE_UNHEALTHY] instead.
	StateUnhealthy = "STATE_UNHEALTHY"
	// Deprecated: StateReconciling is deprecated and will be removed after 9th November 2025. Use [ske.CLUSTERSTATUSSTATE_RECONCILING] instead.
	StateReconciling = "STATE_RECONCILING"
	// Deprecated: CredentialsRotationStatePreparing is deprecated and will be removed after 9th November 2025. Use [ske.CREDENTIALSROTATIONSTATEPHASE_PREPARING] instead.
	CredentialsRotationStatePreparing = "PREPARING"
	// Deprecated: CredentialsRotationStatePrepared is deprecated and will be removed after 9th November 2025. Use [ske.CREDENTIALSROTATIONSTATEPHASE_PREPARED] instead.
	CredentialsRotationStatePrepared = "PREPARED"
	// Deprecated: CredentialsRotationStateCompleting is deprecated and will be removed after 9th November 2025. Use [ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETING] instead.
	CredentialsRotationStateCompleting = "COMPLETING"
	// Deprecated: CredentialsRotationStateCompleted is deprecated and will be removed after 9th November 2025. Use [ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETED] instead.
	CredentialsRotationStateCompleted = "COMPLETED"
	// Deprecated: InvalidArgusInstanceErrorCode is deprecated and will be removed after 9th November 2025. Use [ske.RUNTIMEERRORCODE_ARGUS_INSTANCE_NOT_FOUND] instead.
	InvalidArgusInstanceErrorCode = "SKE_ARGUS_INSTANCE_NOT_FOUND"
)

type APIClientClusterInterface interface {
	GetClusterExecute(ctx context.Context, projectId, name string) (*ske.Cluster, error)
	ListClustersExecute(ctx context.Context, projectId string) (*ske.ListClustersResponse, error)
}

// CreateOrUpdateClusterWaitHandler will wait for cluster creation or update
func CreateOrUpdateClusterWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, name string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetClusterExecute(ctx, projectId, name)
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.Aggregated

		// The state "STATE_UNHEALTHY" (aka "Impaired" in the portal) could be temporarily occur during cluster creation and the system is recovering usually, so it is not considered as a failed state here.
		// -- alignment meeting with SKE team on 4.8.23
		// The exception is when providing an invalid argus instance id, in that case the cluster will stay as "Impaired" until the SKE team solves it, but it is still usable.
		if state == ske.CLUSTERSTATUSSTATE_UNHEALTHY && s.Status.Error != nil && s.Status.Error.Message != nil && *s.Status.Error.Code == ske.RUNTIMEERRORCODE_ARGUS_INSTANCE_NOT_FOUND {
			return true, s, nil
		}

		if state == ske.CLUSTERSTATUSSTATE_HEALTHY || state == ske.CLUSTERSTATUSSTATE_HIBERNATED {
			return true, s, nil
		}

		if state == StateFailed {
			return true, s, fmt.Errorf("create failed")
		}

		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteClusterWaitHandler will wait for cluster deletion
func DeleteClusterWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, name string) *wait.AsyncActionHandler[ske.ListClustersResponse] {
	handler := wait.New(func() (waitFinished bool, response *ske.ListClustersResponse, err error) {
		s, err := a.ListClustersExecute(ctx, projectId)
		if err != nil {
			return false, nil, err
		}
		items := *s.Items
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

// RotateCredentialsWaitHandler will wait for credentials rotation
func RotateCredentialsWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetClusterExecute(ctx, projectId, clusterName)
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.Aggregated

		if state == ske.CLUSTERSTATUSSTATE_HEALTHY || state == ske.CLUSTERSTATUSSTATE_HIBERNATED {
			return true, s, nil
		}

		if state == ske.CLUSTERSTATUSSTATE_RECONCILING {
			return false, nil, nil
		}

		if state == StateFailed {
			return true, s, fmt.Errorf("credentials rotation failed")
		}

		return true, s, fmt.Errorf("unexpected state %s while waiting for cluster reconciliation", state)
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

// StartCredentialsRotationWaitHandler will wait for credentials rotation
func StartCredentialsRotationWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetClusterExecute(ctx, projectId, clusterName)
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.CredentialsRotation.Phase

		if state == ske.CREDENTIALSROTATIONSTATEPHASE_PREPARED {
			return true, s, nil
		}

		if state == ske.CREDENTIALSROTATIONSTATEPHASE_PREPARING {
			return false, nil, nil
		}

		return true, s, fmt.Errorf("unexpected status %s while waiting for cluster credentials rotation to be prepared", state)
	})

	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CompleteCredentialsRotationWaitHandler will wait for credentials rotation
func CompleteCredentialsRotationWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, clusterName string) *wait.AsyncActionHandler[ske.Cluster] {
	handler := wait.New(func() (waitFinished bool, response *ske.Cluster, err error) {
		s, err := a.GetClusterExecute(ctx, projectId, clusterName)
		if err != nil {
			return false, nil, err
		}
		state := *s.Status.CredentialsRotation.Phase

		if state == ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETED {
			return true, s, nil
		}

		if state == ske.CREDENTIALSROTATIONSTATEPHASE_COMPLETING {
			return false, nil, nil
		}

		return true, s, fmt.Errorf("unexpected status %s while waiting for cluster credentials rotation to be completed", state)
	})

	handler.SetTimeout(45 * time.Minute)
	return handler
}
