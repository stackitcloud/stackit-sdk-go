package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2api"
	resourcemanager "github.com/stackitcloud/stackit-sdk-go/services/resourcemanager/v0api"
)

const (
	CreateSuccess         = "CREATED"
	VolumeAvailableStatus = "AVAILABLE"
	DeleteSuccess         = "DELETED"

	FailedStatus = "FAILED"
	ErrorStatus  = "ERROR"

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

// CreateNetworkAreaRegionWaitHandler will wait for network area region creation
func CreateNetworkAreaRegionWaitHandler(ctx context.Context, a iaas.DefaultAPI, organizationId, areaId, region string) *wait.AsyncActionHandler[iaas.RegionalArea] {
	waitConfig := wait.WaiterHelper[iaas.RegionalArea, string]{
		FetchInstance: a.GetNetworkAreaRegion(ctx, organizationId, areaId, region).Execute,
		GetState: func(i *iaas.RegionalArea) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			if i.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *i.Status, nil
		},
		ActiveState: []string{CreateSuccess},
		ErrorState:  []string{FailedStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteNetworkAreaRegionWaitHandler will wait for network area region deletion
func DeleteNetworkAreaRegionWaitHandler(ctx context.Context, a iaas.DefaultAPI, organizationId, areaId, region string) *wait.AsyncActionHandler[iaas.RegionalArea] {
	waitConfig := wait.WaiterHelper[iaas.RegionalArea, string]{
		FetchInstance: a.GetNetworkAreaRegion(ctx, organizationId, areaId, region).Execute,
		GetState: func(i *iaas.RegionalArea) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			if i.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *i.Status, nil
		},
		ActiveState: []string{},
		ErrorState:  []string{FailedStatus},

		// The IaaS API response with a 400 if the regional network area configuration doesn't exist because of some compatible
		// issues to v1. When v1 is deprecated, they may will respond with 404
		DeleteHttpErrorStatusCodes: []int{http.StatusBadRequest, http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// ReadyForNetworkAreaDeletionWaitHandler will wait until a deletion of network area is possible
// Workaround for https://github.com/stackitcloud/terraform-provider-stackit/issues/907.
// When the deletion for a project is triggered, the backend starts a workflow in the background which cleans up all resources
// within a project and deletes the project in each service. When the project is attached to an SNA, the SNA can't be
// deleted until the workflow inform the IaaS-API that the project is deleted.
func ReadyForNetworkAreaDeletionWaitHandler(ctx context.Context, a iaas.DefaultAPI, r resourcemanager.DefaultAPI, organizationId, areaId string) *wait.AsyncActionHandler[iaas.ProjectListResponse] {
	handler := wait.New(func() (waitFinished bool, response *iaas.ProjectListResponse, err error) {
		projectList, err := a.ListNetworkAreaProjects(ctx, organizationId, areaId).Execute()
		if err != nil {
			return false, projectList, err
		}
		if projectList == nil || projectList.Items == nil {
			return false, nil, fmt.Errorf("read failed for projects in network area with id %s, the response is not valid: the items are missing", areaId)
		}
		if len(projectList.Items) == 0 {
			return true, projectList, nil
		}
		var activeProjects, forbiddenProjects []string
		for _, projectId := range projectList.Items {
			_, err := r.GetProject(ctx, projectId).Execute()
			if err == nil {
				activeProjects = append(activeProjects, projectId)
				continue
			}
			var oapiErr *oapierror.GenericOpenAPIError
			ok := errors.As(err, &oapiErr)
			if !ok {
				return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
			}
			// The resource manager api responds with StatusForbidden(=403) when a project is deleted or if the project does not exist
			if oapiErr.StatusCode == http.StatusNotFound || oapiErr.StatusCode == http.StatusForbidden {
				forbiddenProjects = append(forbiddenProjects, projectId)
			}
		}
		if len(activeProjects) > 0 {
			return false, nil, fmt.Errorf("network area with id %s has still active projects: %s", areaId, strings.Join(activeProjects, ","))
		}
		if len(forbiddenProjects) > 0 {
			return false, nil, nil
		}
		return true, projectList, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// CreateNetworkWaitHandler will wait for network creation using network id
func CreateNetworkWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, networkId string) *wait.AsyncActionHandler[iaas.Network] {
	waitConfig := wait.WaiterHelper[iaas.Network, string]{
		FetchInstance: a.GetNetwork(ctx, projectId, region, networkId).Execute,
		GetState: func(i *iaas.Network) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			return i.Status, nil
		},
		ActiveState: []string{CreateSuccess},
		ErrorState:  []string{FailedStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// UpdateNetworkWaitHandler will wait for network update
func UpdateNetworkWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, networkId string) *wait.AsyncActionHandler[iaas.Network] {
	waitConfig := wait.WaiterHelper[iaas.Network, string]{
		FetchInstance: a.GetNetwork(ctx, projectId, region, networkId).Execute,
		GetState: func(i *iaas.Network) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			return i.Status, nil
		},
		ActiveState: []string{CreateSuccess},
		ErrorState:  []string{FailedStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteNetworkWaitHandler will wait for network deletion
func DeleteNetworkWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, networkId string) *wait.AsyncActionHandler[iaas.Network] {
	waitConfig := wait.WaiterHelper[iaas.Network, string]{
		FetchInstance: a.GetNetwork(ctx, projectId, region, networkId).Execute,
		GetState: func(i *iaas.Network) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			return i.Status, nil
		},
		ErrorState: []string{FailedStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// CreateVolumeWaitHandler will wait for volume creation
func CreateVolumeWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, volumeId string) *wait.AsyncActionHandler[iaas.Volume] {
	waitConfig := wait.WaiterHelper[iaas.Volume, string]{
		FetchInstance: a.GetVolume(ctx, projectId, region, volumeId).Execute,
		GetState: func(i *iaas.Volume) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("create failed for volume with id %s, the response is not valid: the id or the status are missing", volumeId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{VolumeAvailableStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteVolumeWaitHandler will wait for volume deletion
func DeleteVolumeWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, volumeId string) *wait.AsyncActionHandler[iaas.Volume] {
	waitConfig := wait.WaiterHelper[iaas.Volume, string]{
		FetchInstance: a.GetVolume(ctx, projectId, region, volumeId).Execute,
		GetState: func(i *iaas.Volume) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("create failed for volume with id %s, the response is not valid: the id or the status are missing", volumeId)
			}
			return *i.Status, nil
		},
		ErrorState: []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// CreateServerWaitHandler will wait for server creation
func CreateServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("create failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ServerActiveStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// ResizeServerWaitHandler will wait for server resize
// It checks for an intermediate resizing status and only then waits for the server to become active
func ResizeServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) (h *wait.AsyncActionHandler[iaas.Server]) {
	handler := wait.New(func() (waitFinished bool, response *iaas.Server, err error) {
		server, err := a.GetServer(ctx, projectId, region, serverId).Execute()
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
func DeleteServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("create failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ErrorState: []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// StartServerWaitHandler will wait for server start
func StartServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("start failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ServerActiveStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// StopServerWaitHandler will wait for server stop
func StopServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("stop failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ServerInactiveStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// DeallocateServerWaitHandler will wait for server deallocation
func DeallocateServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("deallocate failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ServerDeallocatedStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// RescueServerWaitHandler will wait for server rescue
func RescueServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("rescue failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ServerRescueStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// UnrescueServerWaitHandler will wait for server unrescue
func UnrescueServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId string) *wait.AsyncActionHandler[iaas.Server] {
	waitConfig := wait.WaiterHelper[iaas.Server, string]{
		FetchInstance: a.GetServer(ctx, projectId, region, serverId).Execute,
		GetState: func(i *iaas.Server) (string, error) {
			if i.Id == nil || i.Status == nil {
				return "", fmt.Errorf("unrescue failed for server with id %s, the response is not valid: the id or the status are missing", serverId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ServerActiveStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// ProjectRequestWaitHandler will wait for a request to succeed.
//
// It receives a request ID that can be obtained from the "X-Request-Id" header in the HTTP response of any operation in the IaaS API.
// To get this response header, use the "runtime.WithCaptureHTTPResponse" method from the "core" package to get the raw HTTP response of an SDK operation.
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
func ProjectRequestWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, requestId string) *wait.AsyncActionHandler[iaas.Request] {
	handler := wait.New(func() (waitFinished bool, response *iaas.Request, err error) {
		request, err := a.GetProjectRequest(ctx, projectId, region, requestId).Execute()
		if err != nil {
			return false, request, err
		}

		if request == nil {
			return false, nil, fmt.Errorf("request failed for request with id %s: nil response from GetProjectRequestExecute", requestId)
		}

		if request.RequestId != requestId {
			return false, request, fmt.Errorf("request failed for request with id %s: the response id doesn't match the request id", requestId)
		}

		switch request.RequestAction {
		case RequestCreateAction:
			if request.Status == RequestCreatedStatus {
				return true, request, nil
			}
		case RequestUpdateAction:
			if request.Status == RequestUpdatedStatus {
				return true, request, nil
			}
		case RequestDeleteAction:
			if request.Status == RequestDeletedStatus {
				return true, request, nil
			}
		default:
			return false, request, fmt.Errorf("request failed for request with id %s, the request action %s is not supported", requestId, request.RequestAction)
		}

		if request.Status == RequestFailedStatus {
			return true, request, fmt.Errorf("request failed for request with id %s", requestId)
		}

		return false, request, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// AddVolumeToServerWaitHandler will wait for a volume to be attached to a server
// Deprecated: AddVolumeToServerWaitHandler is deprecated and will be removed after October 2026. Please use instead ProjectRequestWaitHandler
func AddVolumeToServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId, volumeId string) *wait.AsyncActionHandler[iaas.VolumeAttachment] {
	waitConfig := wait.WaiterHelper[iaas.VolumeAttachment, string]{
		FetchInstance: a.GetAttachedVolume(ctx, projectId, region, serverId, volumeId).Execute,
		GetState: func(i *iaas.VolumeAttachment) (string, error) {
			if i.VolumeId == nil {
				return "", fmt.Errorf("attachment failed for server with id %s and volume with id %s, the response is not valid: the volume id is missing", serverId, volumeId)
			}
			return *i.VolumeId, nil
		},
		ActiveState: []string{volumeId},
		ErrorState:  []string{},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// RemoveVolumeFromServerWaitHandler will wait for a volume to be attached to a server
// Deprecated: RemoveVolumeFromServerWaitHandler is deprecated and will be removed after October 2026. Please use instead ProjectRequestWaitHandler
func RemoveVolumeFromServerWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, serverId, volumeId string) *wait.AsyncActionHandler[iaas.VolumeAttachment] {
	waitConfig := wait.WaiterHelper[iaas.VolumeAttachment, string]{
		FetchInstance: a.GetAttachedVolume(ctx, projectId, region, serverId, volumeId).Execute,
		GetState: func(i *iaas.VolumeAttachment) (string, error) {
			if i.VolumeId == nil {
				return "", fmt.Errorf("remove volume failed for server with id %s and volume with id %s, the response is not valid: the volume id is missing", serverId, volumeId)
			}
			return *i.VolumeId, nil
		},
		ActiveState: []string{},
		ErrorState:  []string{},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// UploadImageWaitHandler will wait for the status image to become AVAILABLE, which indicates the upload of the image has been completed successfully
func UploadImageWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, imageId string) *wait.AsyncActionHandler[iaas.Image] {
	waitConfig := wait.WaiterHelper[iaas.Image, string]{
		FetchInstance: a.GetImage(ctx, projectId, region, imageId).Execute,
		GetState: func(i *iaas.Image) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("upload failed for image with id %s, the response is not valid: the id or the status are missing", imageId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{ImageAvailableStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteImageWaitHandler will wait for image deletion
func DeleteImageWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, imageId string) *wait.AsyncActionHandler[iaas.Image] {
	waitConfig := wait.WaiterHelper[iaas.Image, string]{
		FetchInstance: a.GetImage(ctx, projectId, region, imageId).Execute,
		GetState: func(i *iaas.Image) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("delete failed for image with id %s, the response is not valid: the id or the status are missing", imageId)
			}
			return *i.Status, nil
		},
		ErrorState: []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// CreateBackupWaitHandler will wait for backup creation
func CreateBackupWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, backupId string) *wait.AsyncActionHandler[iaas.Backup] {
	waitConfig := wait.WaiterHelper[iaas.Backup, string]{
		FetchInstance: a.GetBackup(ctx, projectId, region, backupId).Execute,
		GetState: func(i *iaas.Backup) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("create failed for backup with id %s, the response is not valid: the id or the status are missing", backupId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{BackupAvailableStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteBackupWaitHandler will wait for backup deletion
func DeleteBackupWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, backupId string) *wait.AsyncActionHandler[iaas.Backup] {
	waitConfig := wait.WaiterHelper[iaas.Backup, string]{
		FetchInstance: a.GetBackup(ctx, projectId, region, backupId).Execute,
		GetState: func(i *iaas.Backup) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("delete failed for backup with id %s, the response is not valid: the id or the status are missing", backupId)
			}
			return *i.Status, nil
		},
		ErrorState: []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}

// RestoreBackupWaitHandler will wait for backup restoration
func RestoreBackupWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, backupId string) *wait.AsyncActionHandler[iaas.Backup] {
	waitConfig := wait.WaiterHelper[iaas.Backup, string]{
		FetchInstance: a.GetBackup(ctx, projectId, region, backupId).Execute,
		GetState: func(i *iaas.Backup) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("delete failed for backup with id %s, the response is not valid: the id or the status are missing", backupId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{BackupAvailableStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateSnapshotWaitHandler will wait for snapshot creation
func CreateSnapshotWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, snapshotId string) *wait.AsyncActionHandler[iaas.Snapshot] {
	waitConfig := wait.WaiterHelper[iaas.Snapshot, string]{
		FetchInstance: a.GetSnapshot(ctx, projectId, region, snapshotId).Execute,
		GetState: func(i *iaas.Snapshot) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("create failed for snapshot with id %s, the response is not valid: the id or the status are missing", snapshotId)
			}
			return *i.Status, nil
		},
		ActiveState: []string{SnapshotAvailableStatus},
		ErrorState:  []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteSnapshotWaitHandler will wait for snapshot deletion
func DeleteSnapshotWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, region, snapshotId string) *wait.AsyncActionHandler[iaas.Snapshot] {
	waitConfig := wait.WaiterHelper[iaas.Snapshot, string]{
		FetchInstance: a.GetSnapshot(ctx, projectId, region, snapshotId).Execute,
		GetState: func(i *iaas.Snapshot) (string, error) {
			if i.Status == nil {
				return "", fmt.Errorf("create failed for snapshot with id %s, the response is not valid: the id or the status are missing", snapshotId)
			}
			return *i.Status, nil
		},
		ErrorState: []string{ErrorStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}
