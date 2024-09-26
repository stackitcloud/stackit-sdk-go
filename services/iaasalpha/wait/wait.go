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
	AvailableStatus = "AVAILABLE"
	DeleteSuccess   = "DELETED"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetVolumeExecute(ctx context.Context, projectId string, volumeId string) (*iaasalpha.Volume, error)
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
		if *volume.Id == volumeId && *volume.Status == AvailableStatus {
			return true, volume, nil
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
			return false, nil, nil
		}
		oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
		if !ok {
			return false, volume, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError: %w", err)
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			if volume != nil {
				if volume.Id == nil || volume.Status == nil {
					return false, nil, fmt.Errorf("delete failed for volume with id %s, the response is not valid: the id or the status are missing", volumeId)
				}
				if *volume.Id == volumeId && *volume.Status == DeleteSuccess {
					return true, nil, nil
				}
			}
			return false, volume, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
