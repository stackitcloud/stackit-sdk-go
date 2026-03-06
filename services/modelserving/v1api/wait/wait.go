package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
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
	handler := wait.New(func() (waitFinished bool, response *modelserving.GetTokenResponse, err error) {
		getTokenResp, err := a.GetToken(ctx, region, projectId, tokenId).Execute()
		if err != nil {
			return false, nil, err
		}
		if getTokenResp.Token.State == TOKENSTATE_ACTIVE {
			return true, getTokenResp, nil
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)

	return handler
}

// UpdateModelServingWaitHandler will wait for the model serving auth token to be updated.
// Eventually it will have a different implementation, but for now it's the same as the create handler.
func UpdateModelServingWaitHandler(ctx context.Context, a modelserving.DefaultAPI, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	return CreateModelServingWaitHandler(ctx, a, region, projectId, tokenId)
}

func DeleteModelServingWaitHandler(ctx context.Context, a modelserving.DefaultAPI, region, projectId, tokenId string) *wait.AsyncActionHandler[modelserving.GetTokenResponse] {
	handler := wait.New(
		func() (waitFinished bool, response *modelserving.GetTokenResponse, err error) {
			_, err = a.GetToken(ctx, region, projectId, tokenId).Execute()
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
