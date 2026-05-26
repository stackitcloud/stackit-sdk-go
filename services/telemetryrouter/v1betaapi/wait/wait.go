package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	telemetryrouter "github.com/stackitcloud/stackit-sdk-go/services/telemetryrouter/v1betaapi"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	TELEMETRYROUTER_ACTIVE = telemetryrouter.TELEMETRYROUTERRESPONSESTATUS_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	DESTINATION_ACTIVE = telemetryrouter.DESTINATIONRESPONSESTATUS_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ACCESSTOKEN_ACTIVE = telemetryrouter.ACCESSTOKENBASERESPONSESTATUS_ACTIVE
)

// CreateTelemetryRouterWaitHandler will wait for TelemetryRouter creation
func CreateTelemetryRouterWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId string) *wait.AsyncActionHandler[telemetryrouter.TelemetryRouterResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.TelemetryRouterResponse, telemetryrouter.TelemetryRouterResponseStatus]{
		FetchInstance: a.GetTelemetryRouter(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *telemetryrouter.TelemetryRouterResponse) (telemetryrouter.TelemetryRouterResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetryrouter.TelemetryRouterResponseStatus{telemetryrouter.TELEMETRYROUTERRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateTelemetryRouterWaitHandler will wait for TelemetryRouter update
func UpdateTelemetryRouterWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId string) *wait.AsyncActionHandler[telemetryrouter.TelemetryRouterResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.TelemetryRouterResponse, telemetryrouter.TelemetryRouterResponseStatus]{
		FetchInstance: a.GetTelemetryRouter(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *telemetryrouter.TelemetryRouterResponse) (telemetryrouter.TelemetryRouterResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetryrouter.TelemetryRouterResponseStatus{telemetryrouter.TELEMETRYROUTERRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteTelemetryRouterWaitHandler will wait for TelemetryRouter deletion
func DeleteTelemetryRouterWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId string) *wait.AsyncActionHandler[telemetryrouter.TelemetryRouterResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.TelemetryRouterResponse, telemetryrouter.TelemetryRouterResponseStatus]{
		FetchInstance: a.GetTelemetryRouter(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *telemetryrouter.TelemetryRouterResponse) (telemetryrouter.TelemetryRouterResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateDestinationWaitHandler will wait for Destination creation
func CreateDestinationWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId, destinationId string) *wait.AsyncActionHandler[telemetryrouter.DestinationResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.DestinationResponse, telemetryrouter.DestinationResponseStatus]{
		FetchInstance: a.GetDestination(ctx, projectId, regionId, instanceId, destinationId).Execute,
		GetState: func(d *telemetryrouter.DestinationResponse) (telemetryrouter.DestinationResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetryrouter.DestinationResponseStatus{telemetryrouter.DESTINATIONRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateDestinationWaitHandler will wait for Destination update
func UpdateDestinationWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId, destinationId string) *wait.AsyncActionHandler[telemetryrouter.DestinationResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.DestinationResponse, telemetryrouter.DestinationResponseStatus]{
		FetchInstance: a.GetDestination(ctx, projectId, regionId, instanceId, destinationId).Execute,
		GetState: func(d *telemetryrouter.DestinationResponse) (telemetryrouter.DestinationResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetryrouter.DestinationResponseStatus{telemetryrouter.DESTINATIONRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteDestinationWaitHandler will wait for Destination deletion
func DeleteDestinationWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId, destinationId string) *wait.AsyncActionHandler[telemetryrouter.DestinationResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.DestinationResponse, telemetryrouter.DestinationResponseStatus]{
		FetchInstance: a.GetDestination(ctx, projectId, regionId, instanceId, destinationId).Execute,
		GetState: func(d *telemetryrouter.DestinationResponse) (telemetryrouter.DestinationResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateAccessTokenWaitHandler will wait for AccessToken creation
func CreateAccessTokenWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId, accessTokenId string) *wait.AsyncActionHandler[telemetryrouter.GetAccessTokenResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.GetAccessTokenResponse, telemetryrouter.AccessTokenBaseResponseStatus]{
		FetchInstance: a.GetAccessToken(ctx, projectId, regionId, instanceId, accessTokenId).Execute,
		GetState: func(d *telemetryrouter.GetAccessTokenResponse) (telemetryrouter.AccessTokenBaseResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetryrouter.AccessTokenBaseResponseStatus{telemetryrouter.ACCESSTOKENBASERESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateAccessTokenWaitHandler will wait for AccessToken update
func UpdateAccessTokenWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId, accessTokenId string) *wait.AsyncActionHandler[telemetryrouter.GetAccessTokenResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.GetAccessTokenResponse, telemetryrouter.AccessTokenBaseResponseStatus]{
		FetchInstance: a.GetAccessToken(ctx, projectId, regionId, instanceId, accessTokenId).Execute,
		GetState: func(d *telemetryrouter.GetAccessTokenResponse) (telemetryrouter.AccessTokenBaseResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetryrouter.AccessTokenBaseResponseStatus{telemetryrouter.ACCESSTOKENBASERESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteAccessTokenWaitHandler will wait for AccessToken deletion
func DeleteAccessTokenWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId string, regionId string, instanceId, accessTokenId string) *wait.AsyncActionHandler[telemetryrouter.GetAccessTokenResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.GetAccessTokenResponse, telemetryrouter.AccessTokenBaseResponseStatus]{
		FetchInstance: a.GetAccessToken(ctx, projectId, regionId, instanceId, accessTokenId).Execute,
		GetState: func(d *telemetryrouter.GetAccessTokenResponse) (telemetryrouter.AccessTokenBaseResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
