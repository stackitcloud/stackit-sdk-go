package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/ske"
)

const (
	StateHealthy                       = "STATE_HEALTHY"
	StateHibernated                    = "STATE_HIBERNATED"
	StateFailed                        = "STATE_FAILED"
	StateDeleting                      = "STATE_DELETING"
	StateCreated                       = "STATE_CREATED"
	StateUnhealthy                     = "STATE_UNHEALTHY"
	StateReconciling                   = "STATE_RECONCILING"
	CredentialsRotationStatePreparing  = "PREPARING"
	CredentialsRotationStatePrepared   = "PREPARED"
	CredentialsRotationStateCompleting = "COMPLETING"
	CredentialsRotationStateCompleted  = "COMPLETED"
	InvalidArgusInstanceErrorCode      = "SKE_ARGUS_INSTANCE_NOT_FOUND"
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
		if state == StateUnhealthy && s.Status.Error != nil && s.Status.Error.Message != nil && *s.Status.Error.Code == InvalidArgusInstanceErrorCode {
			return true, s, nil
		}

		if state == StateHealthy || state == StateHibernated {
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

		if state == StateHealthy || state == StateHibernated {
			return true, s, nil
		}

		if state == StateReconciling {
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

		if state == CredentialsRotationStatePrepared {
			return true, s, nil
		}

		if state == CredentialsRotationStatePreparing {
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

		if state == CredentialsRotationStateCompleted {
			return true, s, nil
		}

		if state == CredentialsRotationStateCompleting {
			return false, nil, nil
		}

		return true, s, fmt.Errorf("unexpected status %s while waiting for cluster credentials rotation to be completed", state)
	})

	handler.SetTimeout(45 * time.Minute)
	return handler
}
