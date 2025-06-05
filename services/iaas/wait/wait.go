package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"errors"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
)

const (
	CreateSuccess         = "CREATED"
	VolumeAvailableStatus = "AVAILABLE"
	DeleteSuccess         = "DELETED"

	ErrorStatus = "ERROR"

	ServerActiveStatus      = "ACTIVE"
	ServerResizingStatus    = "RESIZING"
	ServerInactiveStatus    = "INACTIVE"
	ServerDeallocatedStatus = "DEALLOCATED"
	ServerRescueStatus      = "RESCUE"

	ImageAvailableStatus = "AVAILABLE"

	RequestCreateAction  = "CREATE"
	RequestUpdateAction  = "UPDATE"
	RequestDeleteAction  = "DELETE"
	RequestCreatedStatus = "CREATED"
	RequestUpdatedStatus = "UPDATED"
	RequestDeletedStatus = "DELETED"
	RequestFailedStatus  = "FAILED"

	XRequestIDHeader = "X-Request-Id"

	BackupAvailableStatus = "AVAILABLE"
	BackupRestoringStatus = "RESTORING"
	BackupDeletingStatus  = "DELETING"

	SnapshotAvailableStatus = "AVAILABLE"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetNetworkAreaExecute(ctx context.Context, organizationId, areaId string) (*iaas.NetworkArea, error)
	GetProjectRequestExecute(ctx context.Context, projectId string, requestId string) (*iaas.Request, error)
	GetNetworkExecute(ctx context.Context, projectId, networkId string) (*iaas.Network, error)
	GetVolumeExecute(ctx context.Context, projectId string, volumeId string) (*iaas.Volume, error)
	GetServerExecute(ctx context.Context, projectId string, serverId string) (*iaas.Server, error)
	GetAttachedVolumeExecute(ctx context.Context, projectId string, serverId string, volumeId string) (*iaas.VolumeAttachment, error)
	GetImageExecute(ctx context.Context, projectId string, imageId string) (*iaas.Image, error)
	GetBackupExecute(ctx context.Context, projectId string, backupId string) (*iaas.Backup, error)
	GetSnapshotExecute(ctx context.Context, projectId string, snapshotId string) (*iaas.Snapshot, error)
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
	handler.SetTimeout(15 * time.Minute)
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
	handler.SetTimeout(15 * time.Minute)
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
	handler.SetTimeout(15 * time.Minute)
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
	handler.SetTimeout(15 * time.Minute)
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
	handler.SetTimeout(15 * time.Minute)
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
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateVolumeWaitHandler will wait for volume creation
func CreateVolumeWaitHandler(ctx context.Context, a APIClientInterface, projectId, volumeId string) *wait.AsyncActionHandler[iaas.Volume] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Volume, err error) {
		volume, err := a.GetVolumeExecute(ctx, projectId, volumeId)
		if err != nil {
			return false, volume, err
		}
		if volume.Id == nil || volume.Status == nil {
			return false, volume, fmt.Errorf("create failed for volume with id %s, the response is not valid: the id or the status are missing", volumeId)
		}
		if *volume.Id == volumeId && *volume.Status == VolumeAvailableStatus {
			return true, volume, nil
		}
		if *volume.Id == volumeId && *volume.Status == ErrorStatus {
			return true, volume, fmt.Errorf("create failed for volume with id %s", volumeId)
		}
		return false, volume, nil
	})
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteVolumeWaitHandler will wait for volume deletion
func DeleteVolumeWaitHandler(ctx context.Context, a APIClientInterface, projectId, volumeId string) *wait.AsyncActionHandler[iaas.Volume] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Volume, err error) {
		volume, err := a.GetVolumeExecute(ctx, projectId, volumeId)
		if err == nil {
			if volume != nil {
				if volume.Id == nil || volume.Status == nil {
					return false, volume, fmt.Errorf("delete failed for volume with id %s, the response is not valid: the id or the status are missing", volumeId)
				}
				if *volume.Id == volumeId && *volume.Status == DeleteSuccess {
					return true, volume, nil
				}
			}
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, volume, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, volume, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// CreateServerWaitHandler will wait for server creation
func CreateServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("create failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerActiveStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("create failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("create failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// ResizeServerWaitHandler will wait for server resize
// It checks for an intermediate resizing status and only then waits for the server to become active
func ResizeServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) (h *wait.AsyncActionHandler[iaas.Server]) {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}

		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("resizing failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}

		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("resizing failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("resizing failed for server with id %s", serverId)
		}

		if !h.IntermediateStateReached {
			if *server.Id == serverId && *server.Status == ServerResizingStatus {
				h.IntermediateStateReached = true
				return false, server, nil
			}
			return false, server, nil
		}

		if *server.Id == serverId && *server.Status == ServerActiveStatus {
			return true, server, nil
		}

		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// DeleteServerWaitHandler will wait for volume deletion
func DeleteServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err == nil {
			if server != nil {
				if server.Id == nil || server.Status == nil {
					return false, server, fmt.Errorf("delete failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
				}
				if *server.Id == serverId && *server.Status == DeleteSuccess {
					return true, server, nil
				}
			}
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, server, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, server, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// StartServerWaitHandler will wait for server start
func StartServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("start failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerActiveStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("start failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("start failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// StopServerWaitHandler will wait for server stop
func StopServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("stop failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerInactiveStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("stop failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("stop failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// DeallocateServerWaitHandler will wait for server deallocation
func DeallocateServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("deallocate failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerDeallocatedStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("deallocate failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("deallocate failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// RescueServerWaitHandler will wait for server rescue
func RescueServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("rescue failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerRescueStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("rescue failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("rescue failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// UnrescueServerWaitHandler will wait for server unrescue
func UnrescueServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("unrescue failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerActiveStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			if server.ErrorMessage != nil {
				return true, server, fmt.Errorf("unrescue failed for server with id %s: %s", serverId, *server.ErrorMessage)
			}
			return true, server, fmt.Errorf("unrescue failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// ProjectRequestWaitHandler will wait for a request to succeed.
//
// It receives a request ID that can be obtained from the "X-Request-Id" header in the HTTP response of any operation in the IaaS API.
// To get this response header, use the "runtime.WithCaptureHTTPResponse" method from the "core" packaghe to get the raw HTTP response of an SDK operation.
// Then, the value of the request ID can be obtained by accessing the header key which is defined in the constant "XRequestIDHeader" of this package.
//
// Example usage:
//
//	var httpResp *http.Response
//	ctxWithHTTPResp := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)
//
//	err = iaasClient.AddPublicIpToServer(ctxWithHTTPResp, projectId, serverId, publicIpId).Execute()
//
//	requestId := httpResp.Header[wait.XRequestIDHeader][0]
//	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasClient, projectId, requestId).WaitWithContext(context.Background())
func ProjectRequestWaitHandler(ctx context.Context, a APIClientInterface, projectId, requestId string) *wait.AsyncActionHandler[iaas.Request] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Request, err error) {
		request, err := a.GetProjectRequestExecute(ctx, projectId, requestId)
		if err != nil {
			return false, request, err
		}

		if request == nil {
			return false, nil, fmt.Errorf("request failed for request with id %s: nil response from GetProjectRequestExecute", requestId)
		}

		if request.RequestId == nil || request.RequestAction == nil || request.Status == nil {
			return false, request, fmt.Errorf("request failed for request with id %s, the response is not valid: the id, the request action or the status are missing", requestId)
		}

		if *request.RequestId != requestId {
			return false, request, fmt.Errorf("request failed for request with id %s: the response id doesn't match the request id", requestId)
		}

		switch *request.RequestAction {
		case RequestCreateAction:
			if *request.Status == RequestCreatedStatus {
				return true, request, nil
			}
		case RequestUpdateAction:
			if *request.Status == RequestUpdatedStatus {
				return true, request, nil
			}
		case RequestDeleteAction:
			if *request.Status == RequestDeletedStatus {
				return true, request, nil
			}
		default:
			return false, request, fmt.Errorf("request failed for request with id %s, the request action %s is not supported", requestId, *request.RequestAction)
		}

		if *request.Status == RequestFailedStatus {
			return true, request, fmt.Errorf("request failed for request with id %s", requestId)
		}

		return false, request, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// AddVolumeToServerWaitHandler will wait for a volume to be attached to a server
func AddVolumeToServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId, volumeId string) *wait.AsyncActionHandler[iaas.VolumeAttachment] {
	handler := wait.New(func() (waitFinished bool, response *iaas.VolumeAttachment, err error) {
		volumeAttachment, err := a.GetAttachedVolumeExecute(ctx, projectId, serverId, volumeId)
		if err == nil {
			if volumeAttachment != nil {
				if volumeAttachment.VolumeId == nil {
					return false, volumeAttachment, fmt.Errorf("attachment failed for server with id %s and volume with id %s, the response is not valid: the volume id is missing", serverId, volumeId)
				}
				if *volumeAttachment.VolumeId == volumeId {
					return true, volumeAttachment, nil
				}
			}
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, volumeAttachment, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, volumeAttachment, err
		}
		return false, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// RemoveVolumeFromServerWaitHandler will wait for a volume to be attached to a server
func RemoveVolumeFromServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId, volumeId string) *wait.AsyncActionHandler[iaas.VolumeAttachment] {
	handler := wait.New(func() (waitFinished bool, response *iaas.VolumeAttachment, err error) {
		volumeAttachment, err := a.GetAttachedVolumeExecute(ctx, projectId, serverId, volumeId)
		if err == nil {
			if volumeAttachment != nil {
				if volumeAttachment.VolumeId == nil {
					return false, volumeAttachment, fmt.Errorf("remove volume failed for server with id %s and volume with id %s, the response is not valid: the volume id is missing", serverId, volumeId)
				}
			}
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, volumeAttachment, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, volumeAttachment, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// UploadImageWaitHandler will wait for the status image to become AVAILABLE, which indicates the upload of the image has been completed successfully
func UploadImageWaitHandler(ctx context.Context, a APIClientInterface, projectId, imageId string) *wait.AsyncActionHandler[iaas.Image] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Image, err error) {
		image, err := a.GetImageExecute(ctx, projectId, imageId)
		if err != nil {
			return false, image, err
		}
		if image.Id == nil || image.Status == nil {
			return false, image, fmt.Errorf("upload failed for image with id %s, the response is not valid: the id or the status are missing", imageId)
		}
		if *image.Id == imageId && *image.Status == ImageAvailableStatus {
			return true, image, nil
		}
		if *image.Id == imageId && *image.Status == ErrorStatus {
			return true, image, fmt.Errorf("upload failed for image with id %s", imageId)
		}
		return false, image, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteImageWaitHandler will wait for image deletion
func DeleteImageWaitHandler(ctx context.Context, a APIClientInterface, projectId, imageId string) *wait.AsyncActionHandler[iaas.Image] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Image, err error) {
		image, err := a.GetImageExecute(ctx, projectId, imageId)
		if err == nil {
			if image != nil {
				if image.Id == nil || image.Status == nil {
					return false, image, fmt.Errorf("delete failed for image with id %s, the response is not valid: the id or the status are missing", imageId)
				}
				if *image.Id == imageId && *image.Status == DeleteSuccess {
					return true, image, nil
				}
			}
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, image, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, image, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateBackupWaitHandler will wait for backup creation
func CreateBackupWaitHandler(ctx context.Context, a APIClientInterface, projectId, backupId string) *wait.AsyncActionHandler[iaas.Backup] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Backup, err error) {
		backup, err := a.GetBackupExecute(ctx, projectId, backupId)
		if err == nil {
			if backup != nil {
				if backup.Id == nil || backup.Status == nil {
					return false, backup, fmt.Errorf("create failed for backup with id %s, the response is not valid: the id or the status are missing", backupId)
				}
				if *backup.Id == backupId && *backup.Status == BackupAvailableStatus {
					return true, backup, nil
				}
				if *backup.Id == backupId && *backup.Status == ErrorStatus {
					return true, backup, fmt.Errorf("create failed for backup with id %s", backupId)
				}
			}
			return false, nil, nil
		}
		return false, nil, err
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteBackupWaitHandler will wait for backup deletion
func DeleteBackupWaitHandler(ctx context.Context, a APIClientInterface, projectId, backupId string) *wait.AsyncActionHandler[iaas.Backup] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Backup, err error) {
		backup, err := a.GetBackupExecute(ctx, projectId, backupId)
		if err == nil {
			if backup != nil {
				if backup.Id == nil || backup.Status == nil {
					return false, backup, fmt.Errorf("delete failed for backup with id %s, the response is not valid: the id or the status are missing", backupId)
				}
				if *backup.Id == backupId && *backup.Status == DeleteSuccess {
					return true, backup, nil
				}
			}
			return false, nil, nil
		}
		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// RestoreBackupWaitHandler will wait for backup restoration
func RestoreBackupWaitHandler(ctx context.Context, a APIClientInterface, projectId, backupId string) *wait.AsyncActionHandler[iaas.Backup] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Backup, err error) {
		backup, err := a.GetBackupExecute(ctx, projectId, backupId)
		if err == nil {
			if backup != nil {
				if backup.Id == nil || backup.Status == nil {
					return false, backup, fmt.Errorf("restore failed for backup with id %s, the response is not valid: the id or the status are missing", backupId)
				}
				if *backup.Id == backupId && *backup.Status == BackupAvailableStatus {
					return true, backup, nil
				}
				if *backup.Id == backupId && *backup.Status == ErrorStatus {
					return true, backup, fmt.Errorf("restore failed for backup with id %s", backupId)
				}
			}
			return false, nil, nil
		}
		return false, nil, err
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateSnapshotWaitHandler will wait for snapshot creation
func CreateSnapshotWaitHandler(ctx context.Context, a APIClientInterface, projectId, snapshotId string) *wait.AsyncActionHandler[iaas.Snapshot] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Snapshot, err error) {
		snapshot, err := a.GetSnapshotExecute(ctx, projectId, snapshotId)
		if err == nil {
			if snapshot != nil {
				if snapshot.Id == nil || snapshot.Status == nil {
					return false, snapshot, fmt.Errorf("create failed for snapshot with id %s, the response is not valid: the id or the status are missing", snapshotId)
				}
				if *snapshot.Id == snapshotId && *snapshot.Status == SnapshotAvailableStatus {
					return true, snapshot, nil
				}
				if *snapshot.Id == snapshotId && *snapshot.Status == ErrorStatus {
					return true, snapshot, fmt.Errorf("create failed for snapshot with id %s", snapshotId)
				}
			}
			return false, nil, nil
		}
		return false, nil, err
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteSnapshotWaitHandler will wait for snapshot deletion
func DeleteSnapshotWaitHandler(ctx context.Context, a APIClientInterface, projectId, snapshotId string) *wait.AsyncActionHandler[iaas.Snapshot] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Snapshot, err error) {
		snapshot, err := a.GetSnapshotExecute(ctx, projectId, snapshotId)
		if err == nil {
			if snapshot != nil {
				if snapshot.Id == nil || snapshot.Status == nil {
					return false, snapshot, fmt.Errorf("delete failed for snapshot with id %s, the response is not valid: the id or the status are missing", snapshotId)
				}
				if *snapshot.Id == snapshotId && *snapshot.Status == DeleteSuccess {
					return true, snapshot, nil
				}
			}
			return false, nil, nil
		}
		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}
