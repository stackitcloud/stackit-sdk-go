package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/cdn"
)

const (
	DistributionStateCreating = "CREATING"
	DistributionStateActive   = "ACTIVE"
	DistributionStateUpdating = "UPDATING"
	DistributionStateDeleting = "DELETING"
	DistributionStateError    = "ERROR"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetDistributionExecute(ctx context.Context, projectId string, distributionId string) (*cdn.GetDistributionResponse, error)
}

func CreateDistributionPoolWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistributionExecute(ctx, projectId, distributionId)
		if err != nil {
			return false, distribution, err
		}
		if distribution == nil ||
			distribution.Distribution == nil ||
			distribution.Distribution.Id == nil ||
			distribution.Distribution.Status == nil {
			return false, distribution, fmt.Errorf("create failed for distribution with id %s, the response is not valid (state missing)", distributionId)
		}
		if *distribution.Distribution.Id == distributionId {
			switch *distribution.Distribution.Status {
			case DistributionStateActive:
				return true, distribution, err
			default:
				return false, distribution, err
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateDistributionWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistributionExecute(ctx, projectId, distributionId)
		if err != nil {
			return false, distribution, err
		}
		if distribution == nil ||
			distribution.Distribution == nil ||
			distribution.Distribution.Id == nil ||
			distribution.Distribution.Status == nil {
			return false, distribution, fmt.Errorf("update failed for distribution with id %s, the response is not valid (state missing)", distributionId)
		}
		if *distribution.Distribution.Id == distributionId {
			switch *distribution.Distribution.Status {
			case DistributionStateActive:
				return true, distribution, err
			default:
				return false, distribution, err
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteDistributionWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistributionExecute(ctx, projectId, distributionId)
		if err != nil {
			var oapiError *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiError) {
				if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
					return true, distribution, nil
				}
			}
		}
		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

