package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	modelserving "github.com/stackitcloud/stackit-sdk-go/services/modelserving/v1api"
)

const (
	TOKENSTATE_CREATING = "creating"
	TOKENSTATE_ACTIVE   = "active"
	TOKENSTATE_DELETING = "deleting"
	TOKENSTATE_INACTIVE = "inactive"
)

func CreateModelServingWaitHandler(ctx context.Context, a modelserving.DefaultAPI, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	waitConfig := wait.WaiterHelper[modelserving.GetTokenResponse, modelserving.TokenState]{
		FetchInstance: a.GetToken(ctx, region, projectId, tokenId).Execute,
		GetState: func(response *modelserving.GetTokenResponse) (modelserving.TokenState, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.Token.State, nil
		},
		ActiveState: []modelserving.TokenState{modelserving.TOKENSTATE_ACTIVE},
		ErrorState:  []modelserving.TokenState{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)

	return handler
}

// UpdateModelServingWaitHandler will wait for the model serving auth token to be updated.
// Eventually it will have a different implementation, but for now it's the same as the create handler.
func UpdateModelServingWaitHandler(ctx context.Context, a modelserving.DefaultAPI, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	return CreateModelServingWaitHandler(ctx, a, region, projectId, tokenId)
}

func DeleteModelServingWaitHandler(ctx context.Context, a modelserving.DefaultAPI, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	waitConfig := wait.WaiterHelper[modelserving.GetTokenResponse, modelserving.TokenState]{
		FetchInstance: a.GetToken(ctx, region, projectId, tokenId).Execute,
		GetState: func(response *modelserving.GetTokenResponse) (modelserving.TokenState, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.Token.State, nil
		},
		ActiveState: []modelserving.TokenState{},
		ErrorState:  []modelserving.TokenState{},
	}

	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)

	return handler
}
