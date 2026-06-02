package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	modelserving "github.com/stackitcloud/stackit-sdk-go/services/modelserving/v1api"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	TOKENSTATE_CREATING = modelserving.TOKENSTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	TOKENSTATE_ACTIVE = modelserving.TOKENSTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	TOKENSTATE_DELETING = modelserving.TOKENSTATE_DELETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	TOKENSTATE_INACTIVE = modelserving.TOKENSTATE_INACTIVE
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
