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
	DREMIOSTATE_ACTIVE = "active"
	DREMIOSTATE_ERROR  = "error"

	DREMIOUSERSTATE_ACTIVE = "active"
	DREMIOUSERSTATE_ERROR  = "error"
)

// CreateDremioWaitHandler will wait for the creation of a Dremio instance
func CreateDremioWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId, regionId, dremioId string) *wait.AsyncActionHandler[dremio.DremioResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioResponse, string]{
		FetchInstance: a.GetDremio(ctx, projectId, regionId, dremioId).Execute,
		GetState: func(dremio *dremio.DremioResponse) (string, error) {
			if dremio == nil {
				return "", errors.New("empty response")
			}
			return dremio.State, nil
		},
		ActiveState: []string{DREMIOSTATE_ACTIVE},
		ErrorState:  []string{DREMIOSTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// UpdateDremioWaitHandler will wait an update of a Dremio instance
func UpdateDremioWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId, regionId, dremioId string) *wait.AsyncActionHandler[dremio.DremioResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioResponse, string]{
		FetchInstance: a.GetDremio(ctx, projectId, regionId, dremioId).Execute,
		GetState: func(dremio *dremio.DremioResponse) (string, error) {
			if dremio == nil {
				return "", errors.New("empty response")
			}
			return dremio.State, nil
		},
		ActiveState: []string{DREMIOSTATE_ACTIVE},
		ErrorState:  []string{DREMIOSTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteDremioWaitHandler will wait for the deletion of a Dremio instance
func DeleteDremioWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId, regionId, dremioId string) *wait.AsyncActionHandler[dremio.DremioResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioResponse, string]{
		FetchInstance: a.GetDremio(ctx, projectId, regionId, dremioId).Execute,
		GetState: func(dremio *dremio.DremioResponse) (string, error) {
			if dremio == nil {
				return "", errors.New("empty response")
			}
			return dremio.State, nil
		},
		ErrorState:                 []string{DREMIOSTATE_ERROR},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// CreateDremioUserWaitHandler will wait for the creation of a Dremio user
func CreateDremioUserWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId, regionId, dremioId, dremioUserId string) *wait.AsyncActionHandler[dremio.DremioUserResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioUserResponse, string]{
		FetchInstance: a.GetDremioUser(ctx, projectId, regionId, dremioId, dremioUserId).Execute,
		GetState: func(user *dremio.DremioUserResponse) (string, error) {
			if user == nil {
				return "", errors.New("empty response")
			}
			return user.State, nil
		},
		ActiveState: []string{DREMIOUSERSTATE_ACTIVE},
		ErrorState:  []string{DREMIOUSERSTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// UpdateDremioUserWaitHandler will wait for an update to a Dremio user
func UpdateDremioUserWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId, regionId, dremioId, dremioUserId string) *wait.AsyncActionHandler[dremio.DremioUserResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioUserResponse, string]{
		FetchInstance: a.GetDremioUser(ctx, projectId, regionId, dremioId, dremioUserId).Execute,
		GetState: func(user *dremio.DremioUserResponse) (string, error) {
			if user == nil {
				return "", errors.New("empty response")
			}
			return user.State, nil
		},
		ActiveState: []string{DREMIOUSERSTATE_ACTIVE},
		ErrorState:  []string{DREMIOUSERSTATE_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteDremioUserWaitHandler will wait for a deletion of a Dremio user
func DeleteDremioUserWaitHandler(ctx context.Context, a dremio.DefaultAPI, projectId, regionId, dremioId, dremioUserId string) *wait.AsyncActionHandler[dremio.DremioUserResponse] {
	waitConfig := wait.WaiterHelper[dremio.DremioUserResponse, string]{
		FetchInstance: a.GetDremioUser(ctx, projectId, regionId, dremioId, dremioUserId).Execute,
		GetState: func(user *dremio.DremioUserResponse) (string, error) {
			if user == nil {
				return "", errors.New("empty response")
			}
			return user.State, nil
		},
		ErrorState:                 []string{DREMIOUSERSTATE_ERROR},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
