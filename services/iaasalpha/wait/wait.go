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
	VolumeAvailableStatus = "AVAILABLE"
	DeleteSuccess         = "DELETED"
	ErrorStatus           = "ERROR"
	ServerActiveStatus    = "ACTIVE"
	ServerResizingStatus  = "RESIZING"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetVolumeExecute(ctx context.Context, projectId string, volumeId string) (*iaasalpha.Volume, error)
	GetServerExecute(ctx context.Context, projectId string, serverId string) (*iaasalpha.Server, error)
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
			return true, server, fmt.Errorf("create failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// ResizingServerWaitHandler will wait for a server to be resizing
// Ii is an intermediate step inside the ResizeServerWaitHandler
func resizingServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaasalpha.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Server, err error) {
		server, err := a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("resizing failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerResizingStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			return true, server, fmt.Errorf("resizing failed for server with id %s", serverId)
		}
		return false, server, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// ResizeServerWaitHandler will wait for server resize
// It checks for an intermediate resizing status and only then waits for the server to become active
func ResizeServerWaitHandler(ctx context.Context, a APIClientInterface, projectId, serverId string) *wait.AsyncActionHandler[iaasalpha.Server] {
	handler := wait.New(func() (waitFinished bool, response *iaasalpha.Server, err error) {
		server, err := resizingServerWaitHandler(ctx, a, projectId, serverId).WaitWithContext(ctx)
		if err != nil {
			return false, server, err
		}
		server, err = a.GetServerExecute(ctx, projectId, serverId)
		if err != nil {
			return false, server, err
		}
		if server.Id == nil || server.Status == nil {
			return false, server, fmt.Errorf("resizing failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
		}
		if *server.Id == serverId && *server.Status == ServerActiveStatus {
			return true, server, nil
		}
		if *server.Id == serverId && *server.Status == ErrorStatus {
			return true, server, fmt.Errorf("resizing failed for server with id %s", serverId)
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
