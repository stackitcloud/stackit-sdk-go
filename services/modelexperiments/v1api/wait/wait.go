package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	modelexperiments "github.com/stackitcloud/stackit-sdk-go/services/modelexperiments/v1api"
)

// CreateInstanceWaitHandler will wait for creation of Model Experiments instance
func CreateInstanceWaitHandler(ctx context.Context, a modelexperiments.DefaultAPI, region, projectId, instanceId string) *wait.AsyncActionHandler[modelexperiments.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[modelexperiments.GetInstanceResponse, modelexperiments.InstanceState]{
		FetchInstance: a.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *modelexperiments.GetInstanceResponse) (modelexperiments.InstanceState, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.Instance.State, nil
		},
		ActiveState: []modelexperiments.InstanceState{modelexperiments.INSTANCESTATE_ACTIVE},
		ErrorState:  []modelexperiments.InstanceState{modelexperiments.INSTANCESTATE_IMPAIRED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceWaitHandler will wait for deletion of Model Experiments instance
func DeleteInstanceWaitHandler(ctx context.Context, a modelexperiments.DefaultAPI, region, projectId, instanceId string) *wait.AsyncActionHandler[modelexperiments.GetInstanceResponse] {
	waitConfig := wait.WaiterHelper[modelexperiments.GetInstanceResponse, modelexperiments.InstanceState]{
		FetchInstance: a.GetInstance(ctx, projectId, region, instanceId).Execute,
		GetState: func(response *modelexperiments.GetInstanceResponse) (modelexperiments.InstanceState, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.Instance.State, nil
		},
		ErrorState:                 []modelexperiments.InstanceState{modelexperiments.INSTANCESTATE_IMPAIRED},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// CreateInstanceTokenWaitHandler will wait for creation of Model Experiments instance token
func CreateInstanceTokenWaitHandler(ctx context.Context, a modelexperiments.DefaultAPI, region, projectId, instanceId, tokenId string) *wait.AsyncActionHandler[modelexperiments.GetInstanceTokenResponse] {
	waitConfig := wait.WaiterHelper[modelexperiments.GetInstanceTokenResponse, modelexperiments.TokenState]{
		FetchInstance: a.GetInstanceToken(ctx, projectId, region, tokenId, instanceId).Execute,
		GetState: func(response *modelexperiments.GetInstanceTokenResponse) (modelexperiments.TokenState, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.Token.State, nil
		},
		ActiveState: []modelexperiments.TokenState{modelexperiments.TOKENSTATE_ACTIVE},
		ErrorState:  []modelexperiments.TokenState{},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// DeleteInstanceTokenWaitHandler will wait for deletion of Model Experiments instance token
func DeleteInstanceTokenWaitHandler(ctx context.Context, a modelexperiments.DefaultAPI, region, projectId, instanceId, tokenId string) *wait.AsyncActionHandler[modelexperiments.GetInstanceTokenResponse] {
	waitConfig := wait.WaiterHelper[modelexperiments.GetInstanceTokenResponse, modelexperiments.TokenState]{
		FetchInstance: a.GetInstanceToken(ctx, projectId, region, tokenId, instanceId).Execute,
		GetState: func(response *modelexperiments.GetInstanceTokenResponse) (modelexperiments.TokenState, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.Token.State, nil
		},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}
