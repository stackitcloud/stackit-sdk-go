package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	registry "github.com/stackitcloud/stackit-sdk-go/services/registry/v1api"
)

const (
	ArtifactoryReconciling = registry.ARTIFACTORYSTATE_RECONCILING
	ArtifactoryActive      = registry.ARTIFACTORYSTATE_ACTIVE
	ArtifactoryError       = registry.ARTIFACTORYSTATE_ERROR
)

// CreateArtifactoryWaitHandler will wait for the artifactory to be created/ready.
// If there is an error state, it will return the error immediately.
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
//
// The API provides a single delete call, but the internal deletion behavior
// depends on the status of the backend custom resource (CR) when delete is called:
//  1. Soft delete (CR is in Active state):
//     Active -> Reconciling -> HTTP 404 (Not Found).
//     The artifactory is soft-deleted; the API returns HTTP 404 (Not Found),
//     even though the underlying custom resource still remains internally.
//  2. Hard delete (CR is in Error state):
//     Error -> HTTP 404 (Not Found).
//     The custom resource is permanently purged without transitioning through Reconciling.
//
// In both flows, the wait handler polls until GetArtifactory returns HTTP 404 (Not Found).
func DeleteArtifactoryWaitHandler(ctx context.Context, a registry.DefaultAPI, projectId, regionId, artifactoryId string) *wait.AsyncActionHandler[registry.Artifactory] {
	waitConfig := wait.WaiterHelper[registry.Artifactory, registry.ArtifactoryState]{
		FetchInstance:              a.GetArtifactory(ctx, projectId, regionId, artifactoryId).Execute,
		GetState:                   getStateArtifactoryDelete,
		ActiveState:                []registry.ArtifactoryState{},
		ErrorState:                 nil,
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

// For insert and update the error state is an error and the error is enriched with the information got from the container registry API.
func getStateArtifactory(response *registry.Artifactory) (registry.ArtifactoryState, error) {
	if response == nil {
		return "", errors.New("empty response")
	}

	state := response.GetState()
	// If the state is error, enrich the description with the information got from the container registry API
	if state == registry.ARTIFACTORYSTATE_ERROR {
		if response.ErrorDescription != nil && *response.ErrorDescription != "" {
			return state, fmt.Errorf("%s", *response.ErrorDescription)
		}
	}
	return state, nil
}

// Special case of getState, in delete action the error state is not an error itself but a normal state.
func getStateArtifactoryDelete(response *registry.Artifactory) (registry.ArtifactoryState, error) {
	if response == nil {
		return "", errors.New("empty response")
	}

	state := response.GetState()
	return state, nil
}
