package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	telemetrylink "github.com/stackitcloud/stackit-sdk-go/services/telemetrylink/v1betaapi"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	TELEMETRYLINK_ACTIVE = telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE
)

// CreateOrUpdateOrganizationTelemetryLinkWaitHandler will wait for organization TelemetryLink creation or update
func CreateOrUpdateOrganizationTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, organizationId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetOrganizationTelemetryLink(ctx, organizationId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetrylink.TelemetryLinkResponseStatus{telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateOrganizationTelemetryLinkWaitHandler will wait for organization TelemetryLink partial update
func PartialUpdateOrganizationTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, organizationId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetOrganizationTelemetryLink(ctx, organizationId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetrylink.TelemetryLinkResponseStatus{telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteOrganizationTelemetryLinkWaitHandler will wait for organization TelemetryLink deletion
func DeleteOrganizationTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, organizationId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetOrganizationTelemetryLink(ctx, organizationId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
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

// CreateOrUpdateFolderTelemetryLinkWaitHandler will wait for folder TelemetryLink creation or update
func CreateOrUpdateFolderTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, folderId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetFolderTelemetryLink(ctx, folderId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetrylink.TelemetryLinkResponseStatus{telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateFolderTelemetryLinkWaitHandler will wait for folder TelemetryLink partial update
func PartialUpdateFolderTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, folderId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetFolderTelemetryLink(ctx, folderId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetrylink.TelemetryLinkResponseStatus{telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteFolderTelemetryLinkWaitHandler will wait for folder TelemetryLink deletion
func DeleteFolderTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, folderId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetFolderTelemetryLink(ctx, folderId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
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

// CreateOrUpdateProjectTelemetryLinkWaitHandler will wait for project TelemetryLink creation or update
func CreateOrUpdateProjectTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, projectId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetProjectTelemetryLink(ctx, projectId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetrylink.TelemetryLinkResponseStatus{telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateProjectTelemetryLinkWaitHandler will wait for project TelemetryLink partial update
func PartialUpdateProjectTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, projectId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetProjectTelemetryLink(ctx, projectId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Status, nil
		},
		ActiveState: []telemetrylink.TelemetryLinkResponseStatus{telemetrylink.TELEMETRYLINKRESPONSESTATUS_ACTIVE},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteProjectTelemetryLinkWaitHandler will wait for project TelemetryLink deletion
func DeleteProjectTelemetryLinkWaitHandler(ctx context.Context, a telemetrylink.DefaultAPI, projectId string, region telemetrylink.GetFolderTelemetryLinkRegionIdParameter) *wait.AsyncActionHandler[telemetrylink.TelemetryLinkResponse] {
	waitConfig := wait.WaiterHelper[telemetrylink.TelemetryLinkResponse, telemetrylink.TelemetryLinkResponseStatus]{
		FetchInstance: a.GetProjectTelemetryLink(ctx, projectId, region).Execute,
		GetState: func(d *telemetrylink.TelemetryLinkResponse) (telemetrylink.TelemetryLinkResponseStatus, error) {
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
