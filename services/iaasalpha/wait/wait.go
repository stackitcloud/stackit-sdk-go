package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
)

const (
	DeleteSuccess = "DELETED"
	ErrorStatus   = "ERROR"

	VolumeAvailableStatus = "AVAILABLE"

	ServerActiveStatus   = "ACTIVE"
	ServerResizingStatus = "RESIZING"

	VirtualIpCreatedStatus = "CREATED"

	ImageAvailableStatus = "AVAILABLE"

	RequestCreateAction  = "CREATE"
	RequestUpdateAction  = "UPDATE"
	RequestDeleteAction  = "DELETE"
	RequestCreatedStatus = "CREATED"
	RequestUpdatedStatus = "UPDATED"
	RequestDeletedStatus = "DELETED"
	RequestFailedStatus  = "FAILED"

	XRequestIDHeader = "X-Request-Id"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetVolumeExecute(ctx context.Context, projectId string, volumeId string) (*iaasalpha.Volume, error)
	GetServerExecute(ctx context.Context, projectId string, serverId string) (*iaasalpha.Server, error)
	GetProjectRequestExecute(ctx context.Context, projectId string, requestId string) (*iaasalpha.Request, error)
	GetAttachedVolumeExecute(ctx context.Context, projectId string, serverId string, volumeId string) (*iaasalpha.VolumeAttachment, error)
	GetVirtualIPExecute(ctx context.Context, projectId string, networkId string, virtualIpId string) (*iaasalpha.VirtualIp, error)
	GetImageExecute(ctx context.Context, projectId string, imageId string) (*iaasalpha.Image, error)
}

// CreateVolumeWaitHandler will wait for volume creation
func CreateVolumeWaitHandler(ctx context.Context, a APIClientInterface, projectId, volumeId string) *wait.AsyncActionHandler[iaasalpha.Volume] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Volume, err error) {
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
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteVolumeWaitHandler will wait for volume deletion
func DeleteVolumeWaitHandler(ctx context.Context, a APIClientInterface, projectId, volumeId string) *wait.AsyncActionHandler[iaasalpha.Volume] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Volume, err error) {
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
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateServerWaitHandler will wait for server creation
func CreateServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaasalpha.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Server, err error) {
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
func ResizeServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) (h *wait.AsyncActionHandler[iaasalpha.Server]) {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Server, err error) {
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
func DeleteServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaasalpha.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Server, err error) {
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
//	err = iaasalphaClient.AddPublicIpToServer(ctxWithHTTPResp, projectId, serverId, publicIpId).Execute()
//
//	requestId := httpResp.Header[wait.XRequestIDHeader][0]
//	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasalphaClient, projectId, requestId).WaitWithContext(context.Background())
func ProjectRequestWaitHandler(ctx context.Context, a APIClientInterface, projectId, requestId string) *wait.AsyncActionHandler[iaasalpha.Request] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Request, err error) {
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
func AddVolumeToServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId, volumeId string) *wait.AsyncActionHandler[iaasalpha.VolumeAttachment] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.VolumeAttachment, err error) {
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
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// RemoveVolumeFromServerWaitHandler will wait for a volume to be attached to a server
func RemoveVolumeFromServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId, volumeId string) *wait.AsyncActionHandler[iaasalpha.VolumeAttachment] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.VolumeAttachment, err error) {
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
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateVirtualIPWaitHandler will wait for server creation
func CreateVirtualIPWaitHandler(ctx context.Context, a APIClientInterface, projectId, networkId, virtualIpId string) *wait.AsyncActionHandler[iaasalpha.VirtualIp] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.VirtualIp, err error) {
		virtualIp, err := a.GetVirtualIPExecute(ctx, projectId, networkId, virtualIpId)
		if err != nil {
			return false, virtualIp, err
		}
		if virtualIp.Id == nil || virtualIp.Status == nil {
			return false, virtualIp, fmt.Errorf("create failed for virtual ip with id %s, the response is not valid: the id or the status are missing", networkId)
		}
		if *virtualIp.Id == virtualIpId && *virtualIp.Status == VirtualIpCreatedStatus {
			return true, virtualIp, nil
		}
		if *virtualIp.Id == virtualIpId && *virtualIp.Status == ErrorStatus {
			return true, virtualIp, fmt.Errorf("create failed for virtual ip with id %s", networkId)
		}
		return false, virtualIp, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteVirtualIPWaitHandler will wait for volume deletion
func DeleteVirtualIPWaitHandler(ctx context.Context, a APIClientInterface, projectId, networkId, virtualIpId string) *wait.AsyncActionHandler[iaasalpha.VirtualIp] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.VirtualIp, err error) {
		virtualIp, err := a.GetVirtualIPExecute(ctx, projectId, networkId, virtualIpId)
		if err == nil {
			if virtualIp != nil {
				if virtualIp.Id == nil || virtualIp.Status == nil {
					return false, virtualIp, fmt.Errorf("delete failed for virtual ip with id %s, the response is not valid: the id or the status are missing", virtualIpId)
				}
				if *virtualIp.Id == virtualIpId && *virtualIp.Status == DeleteSuccess {
					return true, virtualIp, nil
				}
			}
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, virtualIp, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, virtualIp, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// ImageUploadWaitHandler will wait for server creation
func ImageUploadWaitHandler(ctx context.Context, a APIClientInterface, projectId, imageId string) *wait.AsyncActionHandler[iaasalpha.Image] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Image, err error) {
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
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// DeleteImageWaitHandler will wait for volume deletion
func DeleteImageWaitHandler(ctx context.Context, a APIClientInterface, projectId, imageId string) *wait.AsyncActionHandler[iaasalpha.Image] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Image, err error) {
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
