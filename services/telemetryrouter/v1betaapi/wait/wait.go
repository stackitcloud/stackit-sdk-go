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
	TELEMETRYROUTER_ACTIVE = "active"
	DESTINATION_ACTIVE     = "active"
	ACCESSTOKEN_ACTIVE     = "active"
)

// CreateTelemetryRouterWaitHandler will wait for TelemetryRouter creation
func CreateTelemetryRouterWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[telemetryrouter.TelemetryRouterResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.TelemetryRouterResponse, string]{
		FetchInstance: a.GetTelemetryRouter(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *telemetryrouter.TelemetryRouterResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []string{TELEMETRYROUTER_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateTelemetryRouterWaitHandler will wait for TelemetryRouter update
func UpdateTelemetryRouterWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[telemetryrouter.TelemetryRouterResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.TelemetryRouterResponse, string]{
		FetchInstance: a.GetTelemetryRouter(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *telemetryrouter.TelemetryRouterResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []string{TELEMETRYROUTER_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteTelemetryRouterWaitHandler will wait for TelemetryRouter deletion
func DeleteTelemetryRouterWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId string) *wait.AsyncActionHandler[telemetryrouter.TelemetryRouterResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.TelemetryRouterResponse, string]{
		FetchInstance: a.GetTelemetryRouter(ctx, projectId, regionId, instanceId).Execute,
		GetState: func(d *telemetryrouter.TelemetryRouterResponse) (string, error) {
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
func CreateDestinationWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId, destinationId string) *wait.AsyncActionHandler[telemetryrouter.DestinationResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.DestinationResponse, string]{
		FetchInstance: a.GetDestination(ctx, projectId, regionId, instanceId, destinationId).Execute,
		GetState: func(d *telemetryrouter.DestinationResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []string{DESTINATION_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateDestinationWaitHandler will wait for Destination update
func UpdateDestinationWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId, destinationId string) *wait.AsyncActionHandler[telemetryrouter.DestinationResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.DestinationResponse, string]{
		FetchInstance: a.GetDestination(ctx, projectId, regionId, instanceId, destinationId).Execute,
		GetState: func(d *telemetryrouter.DestinationResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []string{DESTINATION_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteDestinationWaitHandler will wait for Destination deletion
func DeleteDestinationWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId, destinationId string) *wait.AsyncActionHandler[telemetryrouter.DestinationResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.DestinationResponse, string]{
		FetchInstance: a.GetDestination(ctx, projectId, regionId, instanceId, destinationId).Execute,
		GetState: func(d *telemetryrouter.DestinationResponse) (string, error) {
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
func CreateAccessTokenWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId, accessTokenId string) *wait.AsyncActionHandler[telemetryrouter.GetAccessTokenResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.GetAccessTokenResponse, string]{
		FetchInstance: a.GetAccessToken(ctx, projectId, regionId, instanceId, accessTokenId).Execute,
		GetState: func(d *telemetryrouter.GetAccessTokenResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []string{ACCESSTOKEN_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateAccessTokenWaitHandler will wait for AccessToken update
func UpdateAccessTokenWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId, accessTokenId string) *wait.AsyncActionHandler[telemetryrouter.GetAccessTokenResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.GetAccessTokenResponse, string]{
		FetchInstance: a.GetAccessToken(ctx, projectId, regionId, instanceId, accessTokenId).Execute,
		GetState: func(d *telemetryrouter.GetAccessTokenResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []string{ACCESSTOKEN_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteAccessTokenWaitHandler will wait for AccessToken deletion
func DeleteAccessTokenWaitHandler(ctx context.Context, a telemetryrouter.DefaultAPI, projectId, regionId, instanceId, accessTokenId string) *wait.AsyncActionHandler[telemetryrouter.GetAccessTokenResponse] {
	waitConfig := wait.WaiterHelper[telemetryrouter.GetAccessTokenResponse, string]{
		FetchInstance: a.GetAccessToken(ctx, projectId, regionId, instanceId, accessTokenId).Execute,
		GetState: func(d *telemetryrouter.GetAccessTokenResponse) (string, error) {
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
