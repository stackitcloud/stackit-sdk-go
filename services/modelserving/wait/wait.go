package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/modelserving"
)

const (
	activeState = "active"
)

type APIClientInterface interface {
	GetTokenExecute(ctx context.Context, region, projectId, tokenId string) (*modelserving.GetTokenResponse, error)
}

func CreateModelServingWaitHandler(ctx context.Context, a APIClientInterface, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	handler := wait.New(func() (waitFinished bool, response *modelserving.GetTokenResponse, err error) {
		getTokenResp, err := a.GetTokenExecute(ctx, region, projectId, tokenId)
		if err != nil {
			return false, nil, err
		}
		if getTokenResp.Token.State == nil {
			return false, nil, fmt.Errorf(
				"token state is missing for token with id %s",
				tokenId,
			)
		}
		if *getTokenResp.Token.State == activeState {
			return true, getTokenResp, nil
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)

	return handler
}

// UpdateModelServingWaitHandler will wait for the model serving auth token to be updated.
// Eventually it will have a different implementation, but for now it's the same as the create handler.
func UpdateModelServingWaitHandler(ctx context.Context, a APIClientInterface, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	return CreateModelServingWaitHandler(ctx, a, region, projectId, tokenId)
}

func DeleteModelServingWaitHandler(ctx context.Context, a APIClientInterface, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	handler := wait.New(
		func() (waitFinished bool, response *modelserving.GetTokenResponse, err error) {
			_, err = a.GetTokenExecute(ctx, region, projectId, tokenId)
			if err != nil {
				var oapiErr *oapierror.GenericOpenAPIError
				if errors.As(err, &oapiErr) {
					if oapiErr.StatusCode == http.StatusNotFound {
						return true, nil, nil
					}
				}

				return false, nil, err
			}

			return false, nil, nil
		},
	)

	handler.SetTimeout(10 * time.Minute)

	return handler
}
