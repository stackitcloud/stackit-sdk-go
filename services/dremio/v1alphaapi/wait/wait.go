package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	dremio "github.com/stackitcloud/stackit-sdk-go/services/dremio/v1alphaapi"
)

const (
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	DREMIOSTATE_ACTIVE = dremio.DREMIORESPONSESTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	DREMIOSTATE_ERROR = dremio.DREMIORESPONSESTATE_ERROR

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	DREMIOUSERSTATE_ACTIVE = dremio.DREMIOUSERRESPONSESTATE_ACTIVE
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	DREMIOUSERSTATE_ERROR = dremio.DREMIOUSERRESPONSESTATE_ERROR
)

// CreateDremioWaitHandler will wait for the creation of a Dremio instance
func CreateDremioWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId string, regionId dremio.ListDremioInstancesRegionIdParameter, dremioId string) *wait.AsyncActionHandler[dremio.DremioResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioResponse, dremio.DremioResponseState]{
		FetchInstance: a.GetDremioInstance(ctx, projectId, regionId, dremioId).Execute,
		GetState: func(dremioResp *dremio.DremioResponse) (dremio.DremioResponseState, error) {
			if dremioResp == nil {
				return "", errors.New("empty response")
			}
			return dremioResp.State, nil
		},
		ActiveState: []dremio.DremioResponseState{dremio.DREMIORESPONSESTATE_ACTIVE},
		ErrorState:  []dremio.DremioResponseState{dremio.DREMIORESPONSESTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// UpdateDremioWaitHandler will wait an update of a Dremio instance
func UpdateDremioWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId string, regionId dremio.ListDremioInstancesRegionIdParameter, dremioId string) *wait.AsyncActionHandler[dremio.DremioResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioResponse, dremio.DremioResponseState]{
		FetchInstance: a.GetDremioInstance(ctx, projectId, regionId, dremioId).Execute,
		GetState: func(dremioResp *dremio.DremioResponse) (dremio.DremioResponseState, error) {
			if dremioResp == nil {
				return "", errors.New("empty response")
			}
			return dremioResp.State, nil
		},
		ActiveState: []dremio.DremioResponseState{dremio.DREMIORESPONSESTATE_ACTIVE},
		ErrorState:  []dremio.DremioResponseState{dremio.DREMIORESPONSESTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteDremioWaitHandler will wait for the deletion of a Dremio instance
func DeleteDremioWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId string, regionId dremio.ListDremioInstancesRegionIdParameter, dremioId string) *wait.AsyncActionHandler[dremio.DremioResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioResponse, dremio.DremioResponseState]{
		FetchInstance: a.GetDremioInstance(ctx, projectId, regionId, dremioId).Execute,
		GetState: func(dremioResp *dremio.DremioResponse) (dremio.DremioResponseState, error) {
			if dremioResp == nil {
				return "", errors.New("empty response")
			}
			return dremioResp.State, nil
		},
		ErrorState:                 []dremio.DremioResponseState{dremio.DREMIORESPONSESTATE_ERROR},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// CreateDremioUserWaitHandler will wait for the creation of a Dremio user
func CreateDremioUserWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId string, regionId dremio.ListDremioInstancesRegionIdParameter, dremioId, dremioUserId string) *wait.AsyncActionHandler[dremio.DremioUserResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioUserResponse, dremio.DremioUserResponseState]{
		FetchInstance: a.GetDremioUser(ctx, projectId, regionId, dremioId, dremioUserId).Execute,
		GetState: func(user *dremio.DremioUserResponse) (dremio.DremioUserResponseState, error) {
			if user == nil {
				return "", errors.New("empty response")
			}
			return user.State, nil
		},
		ActiveState: []dremio.DremioUserResponseState{dremio.DREMIOUSERRESPONSESTATE_ACTIVE},
		ErrorState:  []dremio.DremioUserResponseState{dremio.DREMIOUSERRESPONSESTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateDremioUserWaitHandler will wait for an update to a Dremio user
func UpdateDremioUserWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId string, regionId dremio.ListDremioInstancesRegionIdParameter, dremioId, dremioUserId string) *wait.AsyncActionHandler[dremio.DremioUserResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioUserResponse, dremio.DremioUserResponseState]{
		FetchInstance: a.GetDremioUser(ctx, projectId, regionId, dremioId, dremioUserId).Execute,
		GetState: func(user *dremio.DremioUserResponse) (dremio.DremioUserResponseState, error) {
			if user == nil {
				return "", errors.New("empty response")
			}
			return user.State, nil
		},
		ActiveState: []dremio.DremioUserResponseState{dremio.DREMIOUSERRESPONSESTATE_ACTIVE},
		ErrorState:  []dremio.DremioUserResponseState{dremio.DREMIOUSERRESPONSESTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteDremioUserWaitHandler will wait for a deletion of a Dremio user
func DeleteDremioUserWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId string, regionId dremio.ListDremioInstancesRegionIdParameter, dremioId, dremioUserId string) *wait.AsyncActionHandler[dremio.DremioUserResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioUserResponse, dremio.DremioUserResponseState]{
		FetchInstance: a.GetDremioUser(ctx, projectId, regionId, dremioId, dremioUserId).Execute,
		GetState: func(user *dremio.DremioUserResponse) (dremio.DremioUserResponseState, error) {
			if user == nil {
				return "", errors.New("empty response")
			}
			return user.State, nil
		},
		ErrorState:                 []dremio.DremioUserResponseState{dremio.DREMIOUSERRESPONSESTATE_ERROR},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
