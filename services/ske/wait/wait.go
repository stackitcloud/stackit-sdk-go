package wait

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/ske"
)

const (
	StateHealthy                  = "STATE_HEALTHY"
	StateHibernated               = "STATE_HIBERNATED"
	StateFailed                   = "STATE_FAILED"
	StateDeleting                 = "STATE_DELETING"
	StateCreated                  = "STATE_CREATED"
	StateUnhealthy                = "STATE_UNHEALTHY"
	InvalidArgusInstanceErrorCode = "SKE_ARGUS_INSTANCE_NOT_FOUND"
)

type APIClientProjectInterface interface {
	GetProjectExecute(ctx context.Context, projectId string) (*ske.ProjectResponse, error)
}

type APIClientClusterInterface interface {
	GetClusterExecute(ctx context.Context, projectId, name string) (*ske.ClusterResponse, error)
	GetClustersExecute(ctx context.Context, projectId string) (*ske.ClustersResponse, error)
}

type APIClientCredentialsInterface interface {
	GetCredentialsExecute(ctx context.Context, projectId, instanceId, credentialsId string) (*ske.CredentialsResponse, error)
}

// CreateOrUpdateClusterWaitHandler will wait for cluster creation or update
func CreateOrUpdateClusterWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, name string) *wait.AsyncActionHandler[ske.ClusterResponse] {
	return wait.New(func() (waitFinished bool, response *ske.ClusterResponse, err error) {
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
}

// DeleteClusterWaitHandler will wait for cluster deletion
func DeleteClusterWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, name string) *wait.AsyncActionHandler[ske.ClustersResponse] {
	return wait.New(func() (waitFinished bool, response *ske.ClustersResponse, err error) {
		s, err := a.GetClustersExecute(ctx, projectId)
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
}

// CreateOrUpdateClusterWaitHandler will wait for project creation
func CreateProjectWaitHandler(ctx context.Context, a APIClientProjectInterface, projectId string) *wait.AsyncActionHandler[ske.ProjectResponse] {
	return wait.New(func() (waitFinished bool, response *ske.ProjectResponse, err error) {
		s, err := a.GetProjectExecute(ctx, projectId)
		if err != nil {
			return false, nil, err
		}
		state := *s.State
		switch state {
		case StateDeleting, StateFailed:
			return false, nil, fmt.Errorf("received state: %s for project Id: %s", state, projectId)
		case StateCreated:
			return true, s, nil
		}
		return false, nil, nil
	})
}

// DeleteProjectWaitHandler will wait for project deletion
func DeleteProjectWaitHandler(ctx context.Context, a APIClientProjectInterface, projectId string) *wait.AsyncActionHandler[struct{}] {
	return wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetProjectExecute(ctx, projectId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError in delete wait.AsyncHandler, %w", err)
		}
		if oapiErr.StatusCode == http.StatusNotFound || oapiErr.StatusCode == http.StatusForbidden {
			return true, nil, nil
		}
		return false, nil, err
	})
}
