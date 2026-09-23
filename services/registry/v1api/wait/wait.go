package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	registry "github.com/stackitcloud/stackit-sdk-go/services/registry/v1api"
)

const (
	ArtifactoryReconcilling = "Reconciling"
	ArtifactoryActive       = "Active"
	ArtifactoryDeleted      = "Deleted"
	ArtifactoryError        = "Error"
)

// CreateArtifactoryWaitHandler will wait for the artifactory to be created/ready.
func CreateArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	waitConfig := wait.WaiterHelper[registry.Artifactory, string]{
		FetchInstance: a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute,
		GetState:      getStateArtifactory,
		ActiveState:   []string{ArtifactoryActive, ArtifactoryActive},
		ErrorState:    []string{ArtifactoryError},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// UpdateArtifactoryWaitHandler will wait for the artifactory update/patch to complete.
func UpdateArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	waitConfig := wait.WaiterHelper[registry.Artifactory, string]{
		FetchInstance: a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute,
		GetState:      getStateArtifactory,
		ActiveState:   []string{ArtifactoryActive, ArtifactoryActive},
		ErrorState:    []string{ArtifactoryError},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

func getStateArtifactory(response *registry.Artifactory) (string, error) {
	if response == nil {
		return "", errors.New("empty response")
	}
	return response.GetState(), nil
}
