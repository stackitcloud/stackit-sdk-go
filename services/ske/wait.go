package ske

import (
	"context"
	"fmt"
	"net/http"

	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

const (
	StateHealthy    = "STATE_HEALTHY"
	StateHibernated = "STATE_HIBERNATED"
	StateFailed     = "STATE_FAILED"
	StateDeleting   = "STATE_DELETING"
	stateCreated    = "STATE_CREATED"
)

type APIClientProjectInterface interface {
	GetProjectExecute(ctx context.Context, projectId string) (*ProjectResponse, error)
}

type APIClientClusterInterface interface {
	GetClusterExecute(ctx context.Context, projectId, name string) (*ClusterResponse, error)
	GetClustersExecute(ctx context.Context, projectId string) (*ClustersResponse, error)
}

type APIClientCredentialsInterface interface {
	GetCredentialsExecute(ctx context.Context, projectId, instanceId, credentialsId string) (*CredentialsResponse, error)
}

// CreateOrUpdateClusterWaitHandler will wait for creation
func CreateOrUpdateClusterWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, name string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetClusterExecute(ctx, projectId, name)
		if err != nil {
			return nil, false, err
		}
		state := *s.Status.Aggregated
		if state == StateHealthy || state == StateHibernated {
			return s, true, nil
		}
		// The state "STATE_UNHEALTHY" (aka "Impaired" in the portal) could be temporarily occur during cluster creation and the system is recovering usually, so it is not considered as a failed state here.
		// -- alignment meeting with SKE team on 4.8.23
		return s, false, nil
	})
}

// DeleteClusterWaitHandler will wait for delete
func DeleteClusterWaitHandler(ctx context.Context, a APIClientClusterInterface, projectId, name string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetClustersExecute(ctx, projectId)
		if err != nil {
			return nil, false, err
		}
		items := *s.Items
		for i := range items {
			n := items[i].Name
			if n != nil && *n == name {
				return s, false, nil
			}
		}
		return s, true, nil
	})
}

func CreateProjectWaitHandler(ctx context.Context, a APIClientProjectInterface, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetProjectExecute(ctx, projectId)
		if err != nil {
			return nil, false, err
		}
		state := *s.State
		switch state {
		case StateDeleting, StateFailed:
			return nil, false, fmt.Errorf("received state: %s for project Id: %s", state, projectId)
		case stateCreated:
			return s, true, nil
		}
		return s, false, nil
	})
}

// DeleteProjectWaitHandler will wait for delete
func DeleteProjectWaitHandler(ctx context.Context, a APIClientProjectInterface, projectId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetProjectExecute(ctx, projectId)
		if err != nil {
			oapiErr, ok := err.(*oapiError.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
			if !ok {
				return nil, false, fmt.Errorf("could not convert error to oapiError.GenericOpenAPIError in delete wait handler, %w", err)
			}
			if oapiErr.StatusCode == http.StatusNotFound || oapiErr.StatusCode == http.StatusForbidden {
				return nil, true, nil
			}
			return nil, false, err
		}
		return s, false, nil
	})
}
