package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/edge"
)

const timeoutMinutes time.Duration = 10

var (
	ErrInstanceNotFound        = errors.New("instance not found")
	ErrInstanceStatusUndefined = errors.New("instance status undefined")
	ErrInstanceCreationFailed  = errors.New("instance creation failed")
	ErrInstanceIsBeingDeleted  = errors.New("instance is being deleted")
)

// EdgeCloudApiClient is the interface for Edge Cloud API calls which require a waiter.
type EdgeCloudApiClient interface {
	GetInstanceExecute(ctx context.Context, projectId, regionId, instanceId string) (*edge.Instance, error)
	GetInstanceByNameExecute(ctx context.Context, projectId, regionId, displayName string) (*edge.Instance, error)
	GetTokenByInstanceId(ctx context.Context, projectId string, regionId string, instanceId string) edge.ApiGetTokenByInstanceIdRequest
	GetTokenByInstanceName(ctx context.Context, projectId string, regionId string, displayName string) edge.ApiGetTokenByInstanceNameRequest
	GetKubeconfigByInstanceId(ctx context.Context, projectId string, regionId string, instanceId string) edge.ApiGetKubeconfigByInstanceIdRequest
	GetKubeconfigByInstanceName(ctx context.Context, projectId string, regionId string, displayName string) edge.ApiGetKubeconfigByInstanceNameRequest
}

// createOrUpdateInstanceWaitHandler contains the shared logic for waiting on instance creation or updates.
func createOrUpdateInstanceWaitHandler(ctx context.Context, getInstance func(ctx context.Context) (*edge.Instance, error)) *wait.AsyncActionHandler[edge.Instance] {
	waitConfig := wait.WaiterHelper[edge.Instance, edge.InstanceStatus]{
		FetchInstance: func() (*edge.Instance, error) {
			return getInstance(ctx)
		},
		GetState: func(instance *edge.Instance) (edge.InstanceStatus, error) {
			if instance == nil || instance.Status == nil {
				return "", ErrInstanceNotFound
			}
			return *instance.Status, nil
		},
		ActiveState: []edge.InstanceStatus{edge.INSTANCESTATUS_ACTIVE},
		ErrorState:  []edge.InstanceStatus{edge.INSTANCESTATUS_ERROR, edge.INSTANCESTATUS_DELETING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(timeoutMinutes * time.Minute)
	return handler
}

// CreateOrUpdateInstanceWaitHandler waits for instance creation or update by ID to complete.
func CreateOrUpdateInstanceWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, instanceId string) *wait.AsyncActionHandler[edge.Instance] {
	return createOrUpdateInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstanceExecute(ctx, projectId, regionId, instanceId)
	})
}

// CreateOrUpdateInstanceByNameWaitHandler waits for instance creation or update by name to complete.
func CreateOrUpdateInstanceByNameWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, displayName string) *wait.AsyncActionHandler[edge.Instance] {
	return createOrUpdateInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstanceByNameExecute(ctx, projectId, regionId, displayName)
	})
}

// deleteInstanceWaitHandler contains the shared logic for waiting on instance deletion.
func deleteInstanceWaitHandler(ctx context.Context, getInstance func(ctx context.Context) (*edge.Instance, error)) *wait.AsyncActionHandler[edge.Instance] {
	waitConfig := wait.WaiterHelper[edge.Instance, edge.InstanceStatus]{
		FetchInstance: func() (*edge.Instance, error) {
			return getInstance(ctx)
		},
		GetState: func(i *edge.Instance) (edge.InstanceStatus, error) {
			return "", nil
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(timeoutMinutes * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler waits for instance deletion by ID.
func DeleteInstanceWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, instanceId string) *wait.AsyncActionHandler[edge.Instance] {
	return deleteInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstanceExecute(ctx, projectId, regionId, instanceId)
	})
}

// DeleteInstanceByNameWaitHandler waits for instance deletion by name.
func DeleteInstanceByNameWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, displayName string) *wait.AsyncActionHandler[edge.Instance] {
	return deleteInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstanceByNameExecute(ctx, projectId, regionId, displayName)
	})
}

// kubeconfigWaitHandlerHelper contains the shared logic for waiting for the instance to become ready before retrieving the kubeconfig.
func kubeconfigWaitHandlerHelper(ctx context.Context, checkInstance func(ctx context.Context) error, getKubeconfig func(ctx context.Context) (*edge.Kubeconfig, error)) *wait.AsyncActionHandler[edge.Kubeconfig] {
	waitConfig := wait.WaiterHelper[edge.Kubeconfig, string]{
		FetchInstance: func() (*edge.Kubeconfig, error) {
			err := checkInstance(ctx)
			if err != nil {
				return nil, err
			}
			config, err := getKubeconfig(ctx)
			if err != nil {
				var apiErr *oapierror.GenericOpenAPIError
				if ok := errors.As(err, &apiErr); ok && apiErr.StatusCode == http.StatusNotFound {
					return nil, nil
				}
			}
			return config, err
		},
		GetState: func(k *edge.Kubeconfig) (string, error) {
			if k == nil {
				return "NOT_READY", nil
			}
			return "READY", nil
		},
		ActiveState: []string{"READY"},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(timeoutMinutes * time.Minute)
	return handler
}

// KubeconfigWaitHandler waits the instance to become ready before retrieving the kubeconfig by instance ID.
func KubeconfigWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, instanceId string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Kubeconfig] {
	return kubeconfigWaitHandlerHelper(ctx,
		func(ctx context.Context) error {
			return checkInstanceExistsWithUsableStatus(ctx, a, projectId, regionId, instanceId)
		},
		func(ctx context.Context) (*edge.Kubeconfig, error) {
			req := a.GetKubeconfigByInstanceId(ctx, projectId, regionId, instanceId)
			if expirationSeconds != nil {
				req = req.ExpirationSeconds(*expirationSeconds)
			}
			return req.Execute()
		},
	)
}

// KubeconfigByInstanceNameWaitHandler waits the instance to become ready before retrieving the kubeconfig by instance displayName.
func KubeconfigByInstanceNameWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, displayName string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Kubeconfig] {
	return kubeconfigWaitHandlerHelper(ctx,
		func(ctx context.Context) error {
			return checkInstanceNameExistsWithUsableStatus(ctx, a, projectId, regionId, displayName)
		},
		func(ctx context.Context) (*edge.Kubeconfig, error) {
			req := a.GetKubeconfigByInstanceName(ctx, projectId, regionId, displayName)
			if expirationSeconds != nil {
				req = req.ExpirationSeconds(*expirationSeconds)
			}
			return req.Execute()
		},
	)
}

// tokenWaitHandlerHelper contains the shared logic for waiting for the instance to become ready before retrieving the service token.
func tokenWaitHandlerHelper(ctx context.Context, checkInstance func(ctx context.Context) error, getToken func(ctx context.Context) (*edge.Token, error)) *wait.AsyncActionHandler[edge.Token] {
	waitConfig := wait.WaiterHelper[edge.Token, string]{
		FetchInstance: func() (*edge.Token, error) {
			err := checkInstance(ctx)
			if err != nil {
				return nil, err
			}
			token, err := getToken(ctx)
			if err != nil {
				var apiErr *oapierror.GenericOpenAPIError
				if ok := errors.As(err, &apiErr); ok && apiErr.StatusCode == http.StatusNotFound {
					return nil, nil
				}
			}
			return token, err
		},
		GetState: func(t *edge.Token) (string, error) {
			if t == nil {
				return "NOT_READY", nil
			}
			return "READY", nil
		},
		ActiveState: []string{"READY"},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(timeoutMinutes * time.Minute)
	return handler
}

// TokenWaitHandler waits for the instance to become ready before retrieving the service token by instance ID.
func TokenWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, instanceId string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Token] {
	return tokenWaitHandlerHelper(ctx,
		func(ctx context.Context) error {
			return checkInstanceExistsWithUsableStatus(ctx, a, projectId, regionId, instanceId)
		},
		func(ctx context.Context) (*edge.Token, error) {
			req := a.GetTokenByInstanceId(ctx, projectId, regionId, instanceId)
			if expirationSeconds != nil {
				req = req.ExpirationSeconds(*expirationSeconds)
			}
			return req.Execute()
		},
	)
}

// TokenByInstanceNameWaitHandler waits for the instance to become ready before retrieving the service token by instance displayName.
func TokenByInstanceNameWaitHandler(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, displayName string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Token] {
	return tokenWaitHandlerHelper(ctx,
		func(ctx context.Context) error {
			return checkInstanceNameExistsWithUsableStatus(ctx, a, projectId, regionId, displayName)
		},
		func(ctx context.Context) (*edge.Token, error) {
			req := a.GetTokenByInstanceName(ctx, projectId, regionId, displayName)
			if expirationSeconds != nil {
				req = req.ExpirationSeconds(*expirationSeconds)
			}
			return req.Execute()
		},
	)
}

// checkInstanceUsableStatus contains the shared logic for checking instance status.
func checkInstanceUsableStatus(ctx context.Context, getInstance func(ctx context.Context) (*edge.Instance, error), identifierType, identifierValue string) error {
	instance, err := getInstance(ctx)
	if err != nil {
		return err
	}
	if instance == nil {
		return ErrInstanceNotFound
	}
	if instance.Status == nil {
		return ErrInstanceStatusUndefined
	}
	if *instance.Status == edge.INSTANCESTATUS_ACTIVE || *instance.Status == edge.INSTANCESTATUS_RECONCILING {
		return nil
	}
	return fmt.Errorf("cannot use instance with %s '%s' with status '%s'", identifierType, identifierValue, *instance.Status)
}

// checkInstanceExistsWithUsableStatus checks if the instance with the given instanceId exists and has a usable status.
func checkInstanceExistsWithUsableStatus(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, instanceId string) error {
	return checkInstanceUsableStatus(
		ctx,
		func(ctx context.Context) (*edge.Instance, error) {
			return a.GetInstanceExecute(ctx, projectId, regionId, instanceId)
		},
		"ID",
		instanceId,
	)
}

// checkInstanceNameExistsWithUsableStatus checks if the instance with the given displayName exists and has a usable status.
func checkInstanceNameExistsWithUsableStatus(ctx context.Context, a EdgeCloudApiClient, projectId, regionId, displayName string) error {
	return checkInstanceUsableStatus(
		ctx,
		func(ctx context.Context) (*edge.Instance, error) {
			return a.GetInstanceByNameExecute(ctx, projectId, regionId, displayName)
		},
		"name",
		displayName,
	)
}
