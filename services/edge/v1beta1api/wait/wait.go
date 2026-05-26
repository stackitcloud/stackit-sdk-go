package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	edge "github.com/stackitcloud/stackit-sdk-go/services/edge/v1beta1api"
)

const timeoutMinutes time.Duration = 10

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_ERROR = edge.INSTANCESTATUS_ERROR
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_RECONCILING = edge.INSTANCESTATUS_RECONCILING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_ACTIVE = edge.INSTANCESTATUS_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	INSTANCESTATUS_DELETING = edge.INSTANCESTATUS_DELETING
)

var (
	ErrInstanceNotFound = errors.New("instance not found")
	// Deprecated: ErrInstanceStatusUndefined is no longer used and will be removed after 2026-11-07
	ErrInstanceStatusUndefined = errors.New("instance status undefined")
	// Deprecated: ErrInstanceCreationFailed is no longer used and will be removed after 2026-11-07
	ErrInstanceCreationFailed = errors.New("instance creation failed")
	// Deprecated: ErrInstanceIsBeingDeleted is no longer used and will be removed after 2026-11-07
	ErrInstanceIsBeingDeleted = errors.New("instance is being deleted")
)

// createOrUpdateInstanceWaitHandler contains the shared logic for waiting on instance creation or updates.
func createOrUpdateInstanceWaitHandler(ctx context.Context, getInstance func(ctx context.Context) (*edge.Instance, error)) *wait.AsyncActionHandler[edge.Instance] {
	waitConfig := wait.WaiterHelper[edge.Instance, edge.InstanceStatus]{
		FetchInstance: func() (*edge.Instance, error) {
			return getInstance(ctx)
		},
		GetState: func(e *edge.Instance) (edge.InstanceStatus, error) {
			if e == nil {
				return "", ErrInstanceNotFound
			}
			return e.Status, nil
		},
		ActiveState: []edge.InstanceStatus{edge.INSTANCESTATUS_ACTIVE},
		ErrorState:  []edge.InstanceStatus{edge.INSTANCESTATUS_ERROR, edge.INSTANCESTATUS_DELETING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(timeoutMinutes * time.Minute)
	return handler
}

// CreateOrUpdateInstanceWaitHandler waits for instance creation or update by ID to complete.
func CreateOrUpdateInstanceWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[edge.Instance] {
	return createOrUpdateInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstance(ctx, projectId, regionId, instanceId).Execute()
	})
}

// CreateOrUpdateInstanceByNameWaitHandler waits for instance creation or update by name to complete.
func CreateOrUpdateInstanceByNameWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, displayName string) *wait.AsyncActionHandler[edge.Instance] {
	return createOrUpdateInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstanceByName(ctx, projectId, regionId, displayName).Execute()
	})
}

// deleteInstanceWaitHandler contains the shared logic for waiting on instance deletion.
func deleteInstanceWaitHandler(ctx context.Context, getInstance func(ctx context.Context) (*edge.Instance, error)) *wait.AsyncActionHandler[edge.Instance] {
	waitConfig := wait.WaiterHelper[edge.Instance, edge.InstanceStatus]{
		FetchInstance: func() (*edge.Instance, error) {
			return getInstance(ctx)
		},
		GetState: func(e *edge.Instance) (edge.InstanceStatus, error) {
			if e == nil {
				return "", ErrInstanceNotFound
			}
			return e.Status, nil
		},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(timeoutMinutes * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler waits for instance deletion by ID.
func DeleteInstanceWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[edge.Instance] {
	return deleteInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstance(ctx, projectId, regionId, instanceId).Execute()
	})
}

// DeleteInstanceByNameWaitHandler waits for instance deletion by name.
func DeleteInstanceByNameWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, displayName string) *wait.AsyncActionHandler[edge.Instance] {
	return deleteInstanceWaitHandler(ctx, func(ctx context.Context) (*edge.Instance, error) {
		return a.GetInstanceByName(ctx, projectId, regionId, displayName).Execute()
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
		GetState: func(e *edge.Kubeconfig) (string, error) {
			if e == nil {
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
func KubeconfigWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, instanceId string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Kubeconfig] {
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
func KubeconfigByInstanceNameWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, displayName string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Kubeconfig] {
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
		GetState: func(e *edge.Token) (string, error) {
			if e == nil {
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
func TokenWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, instanceId string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Token] {
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
func TokenByInstanceNameWaitHandler(ctx context.Context, a edge.DefaultAPI, projectId, regionId, displayName string, expirationSeconds *int64) *wait.AsyncActionHandler[edge.Token] {
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
	if instance.Status == edge.INSTANCESTATUS_ACTIVE || instance.Status == edge.INSTANCESTATUS_RECONCILING {
		return nil
	}
	return fmt.Errorf("cannot use instance with %s '%s' with status '%s'", identifierType, identifierValue, instance.Status)
}

// checkInstanceExistsWithUsableStatus checks if the instance with the given instanceId exists and has a usable status.
func checkInstanceExistsWithUsableStatus(ctx context.Context, a edge.DefaultAPI, projectId, regionId, instanceId string) error {
	return checkInstanceUsableStatus(
		ctx,
		func(ctx context.Context) (*edge.Instance, error) {
			return a.GetInstance(ctx, projectId, regionId, instanceId).Execute()
		},
		"ID",
		instanceId,
	)
}

// checkInstanceNameExistsWithUsableStatus checks if the instance with the given displayName exists and has a usable status.
func checkInstanceNameExistsWithUsableStatus(ctx context.Context, a edge.DefaultAPI, projectId, regionId, displayName string) error {
	return checkInstanceUsableStatus(
		ctx,
		func(ctx context.Context) (*edge.Instance, error) {
			return a.GetInstanceByName(ctx, projectId, regionId, displayName).Execute()
		},
		"name",
		displayName,
	)
}
