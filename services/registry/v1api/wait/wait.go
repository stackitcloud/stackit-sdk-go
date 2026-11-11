package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	registry "github.com/stackitcloud/stackit-sdk-go/services/registry/v1api"
)

const (
	ArtifactoryReconciling = registry.ARTIFACTORYSTATE_RECONCILING
	ArtifactoryActive      = registry.ARTIFACTORYSTATE_ACTIVE
	ArtifactoryDeleted     = registry.ARTIFACTORYSTATE_DELETED
	ArtifactoryError       = registry.ARTIFACTORYSTATE_ERROR
)

// CreateArtifactoryWaitHandler will wait for the artifactory to be created/ready.
func CreateArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	waitConfig := wait.WaiterHelper[registry.Artifactory, registry.ArtifactoryState]{
		FetchInstance: a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute,
		GetState:      getStateArtifactory,
		ActiveState:   []registry.ArtifactoryState{registry.ARTIFACTORYSTATE_ACTIVE},
		ErrorState:    []registry.ArtifactoryState{registry.ARTIFACTORYSTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// UpdateArtifactoryWaitHandler will wait for the artifactory update/patch to complete.
func UpdateArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	waitConfig := wait.WaiterHelper[registry.Artifactory, registry.ArtifactoryState]{
		FetchInstance: a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute,
		GetState:      getStateArtifactory,
		ActiveState:   []registry.ArtifactoryState{registry.ARTIFACTORYSTATE_ACTIVE},
		ErrorState:    []registry.ArtifactoryState{registry.ARTIFACTORYSTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// DeleteArtifactoryWaitHandler will wait for the artifactory to be deleted.
// It supports two deletion flows based on the initial state:
// 1. Soft delete (initial status is Active): waits through Active/Reconciling until state is Deleted.
// 2. Hard delete (initial status is Error): waits through Error/Reconciling until the resource is removed (HTTP 404).
func DeleteArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	var initialChecked bool
	var isHardDelete bool

	handler := wait.New(func() (waitFinished bool, response *registry.Artifactory, err error) {
		instance, err := a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute()
		if err != nil {
			var oapiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiErr) && oapiErr.StatusCode == http.StatusNotFound {
				return true, nil, nil
			}
			return false, nil, err
		}

		state, err := getStateArtifactory(instance)
		if err != nil {
			return false, nil, err
		}

		if !initialChecked {
			initialChecked = true
			switch state {
			case registry.ARTIFACTORYSTATE_ERROR:
				isHardDelete = true
			case registry.ARTIFACTORYSTATE_ACTIVE, registry.ARTIFACTORYSTATE_RECONCILING:
				isHardDelete = false
			case registry.ARTIFACTORYSTATE_DELETED:
				return true, instance, nil
			default:
				return true, instance, fmt.Errorf("cannot wait for delete: artifactory is in invalid initial state %v", state)
			}
		}

		if isHardDelete {
			return false, nil, nil
		}

		switch state {
		case registry.ARTIFACTORYSTATE_DELETED:
			return true, instance, nil
		case registry.ARTIFACTORYSTATE_ACTIVE, registry.ARTIFACTORYSTATE_RECONCILING:
			return false, nil, nil
		case registry.ARTIFACTORYSTATE_ERROR:
			return true, instance, fmt.Errorf("artifactory soft deletion failed: state is %v", state)
		default:
			return false, nil, nil
		}
	})

	handler.SetTimeout(5 * time.Minute)
	return handler
}

func getStateArtifactory(response *registry.Artifactory) (registry.ArtifactoryState, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	return response.GetState(), nil
}
