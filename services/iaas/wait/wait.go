package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
)

const (
	CreateSuccess = "CREATED"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetNetworkAreaExecute(ctx context.Context, organizationId, areaId string) (*iaas.NetworkArea, error)
	GetProjectRequestExecute(ctx context.Context, projectId string, requestId string) (*iaas.Request, error)
	GetNetworkExecute(ctx context.Context, projectId, networkId string) (*iaas.Network, error)
}

// CreateNetworkAreaWaitHandler will wait for network area creation
func CreateNetworkAreaWaitHandler(ctx context.Context, a APIClientInterface, organizationId, areaId string) *wait.AsyncActionHandler[iaas.NetworkArea] {
	handler := wait.New(func() (waitFinished bool, response *iaas.NetworkArea, err error) {
		area, err := a.GetNetworkAreaExecute(ctx, organizationId, areaId)
		if err != nil {
			return false, area, err
		}
		if area.AreaId == nil || area.State == nil {
			return false, area, fmt.Errorf("create failed for network area with id %s, the response is not valid: the id or the state are missing", areaId)
		}
		if *area.AreaId == areaId && *area.State == CreateSuccess {
			return true, area, nil
		}
		return false, area, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateNetworkAreaWaitHandler will wait for network area update
func UpdateNetworkAreaWaitHandler(ctx context.Context, a APIClientInterface, organizationId, areaId string) *wait.AsyncActionHandler[iaas.NetworkArea] {
	handler := wait.New(func() (waitFinished bool, response *iaas.NetworkArea, err error) {
		area, err := a.GetNetworkAreaExecute(ctx, organizationId, areaId)
		if err != nil {
			return false, area, err
		}
		if area.AreaId == nil || area.State == nil {
			return false, nil, fmt.Errorf("update failed for network area with id %s, the response is not valid: the id or the state are missing", areaId)
		}
		// The state returns to "CREATED" after a successful update is completed
		if *area.AreaId == areaId && *area.State == CreateSuccess {
			return true, area, nil
		}
		return false, area, nil
	})
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteNetworkAreaWaitHandler will wait for network area deletion
func DeleteNetworkAreaWaitHandler(ctx context.Context, a APIClientInterface, organizationId, areaId string) *wait.AsyncActionHandler[iaas.NetworkArea] {
	handler := wait.New(func() (waitFinished bool, response *iaas.NetworkArea, err error) {
		area, err := a.GetNetworkAreaExecute(ctx, organizationId, areaId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, area, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, area, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateNetworkWaitHandler will wait for network creation using network id
func CreateNetworkWaitHandler(ctx context.Context, a APIClientInterface, projectId, networkId string) *wait.AsyncActionHandler[iaas.Network] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Network, err error) {
		network, err := a.GetNetworkExecute(ctx, projectId, networkId)
		if err != nil {
			return false, network, err
		}
		if network.NetworkId == nil || network.State == nil {
			return false, network, fmt.Errorf("crate failed for network with id %s, the response is not valid: the id or the state are missing", networkId)
		}
		// The state returns to "CREATED" after a successful creation is completed
		if *network.NetworkId == networkId && *network.State == CreateSuccess {
			return true, network, nil
		}
		return false, network, nil
	})
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateNetworkWaitHandler will wait for network update
func UpdateNetworkWaitHandler(ctx context.Context, a APIClientInterface, projectId, networkId string) *wait.AsyncActionHandler[iaas.Network] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Network, err error) {
		network, err := a.GetNetworkExecute(ctx, projectId, networkId)
		if err != nil {
			return false, network, err
		}
		if network.NetworkId == nil || network.State == nil {
			return false, network, fmt.Errorf("update failed for network with id %s, the response is not valid: the id or the state are missing", networkId)
		}
		// The state returns to "CREATED" after a successful update is completed
		if *network.NetworkId == networkId && *network.State == CreateSuccess {
			return true, network, nil
		}
		return false, network, nil
	})
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteNetworkWaitHandler will wait for network deletion
func DeleteNetworkWaitHandler(ctx context.Context, a APIClientInterface, projectId, networkId string) *wait.AsyncActionHandler[iaas.Network] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Network, err error) {
		network, err := a.GetNetworkExecute(ctx, projectId, networkId)
		if err == nil {
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, network, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, network, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
