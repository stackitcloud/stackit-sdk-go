package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

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
func DeleteArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	waitConfig := wait.WaiterHelper[registry.Artifactory, registry.ArtifactoryState]{
		FetchInstance:              a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute,
		GetState:                   getStateArtifactory,
		ActiveState:                nil,
		ErrorState:                 []registry.ArtifactoryState{registry.ARTIFACTORYSTATE_ERROR},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

func getStateArtifactory(response *registry.Artifactory) (registry.ArtifactoryState, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	return response.GetState(), nil
}
