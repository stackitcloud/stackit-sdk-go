package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	intake "github.com/stackitcloud/stackit-sdk-go/services/intake/v1betaapi"
)

const (
	INTAKERESPONSESTATE_RECONCILING = "reconciling"
	INTAKERESPONSESTATE_ACTIVE      = "active"
	INTAKERESPONSESTATE_DELETING    = "deleting"
	INTAKERESPONSESTATE_FAILED      = "failed"

	INTAKERUNNERRESPONSESTATE_RECONCILING = "reconciling"
	INTAKERUNNERRESPONSESTATE_ACTIVE      = "active"
	INTAKERUNNERRESPONSESTATE_DELETING    = "deleting"

	INTAKEUSERRESPONSESTATE_RECONCILING = "reconciling"
	INTAKEUSERRESPONSESTATE_ACTIVE      = "active"
	INTAKEUSERRESPONSESTATE_DELETING    = "deleting"
)

// Deprecated: Will be removed after 2026-11-13. Use the CreateIntakeRunnerWaitHandler or UpdateIntakeRunnerWaitHandler instead
func CreateOrUpdateIntakeRunnerWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeRunnerId string) *wait.AsyncActionHandler[intake.IntakeRunnerResponse] {
	// TODO: mark function as private after deprecation period
	waitConfig := wait.WaiterHelper[intake.IntakeRunnerResponse, string]{
		FetchInstance: client.GetIntakeRunner(ctx, projectId, region, intakeRunnerId).Execute,
		GetState: func(response *intake.IntakeRunnerResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState: []string{INTAKERUNNERRESPONSESTATE_ACTIVE},
		ErrorState:  []string{},
		// The API does not have a dedicated failure state for this resource,
		// so we rely on the timeout for cases where it never becomes active.
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

func CreateIntakeRunnerWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeRunnerId string) *wait.AsyncActionHandler[intake.IntakeRunnerResponse] {
	return CreateOrUpdateIntakeRunnerWaitHandler(ctx, client, projectId, region, intakeRunnerId)
}

func UpdateIntakeRunnerWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeRunnerId string) *wait.AsyncActionHandler[intake.IntakeRunnerResponse] {
	return CreateOrUpdateIntakeRunnerWaitHandler(ctx, client, projectId, region, intakeRunnerId)
}

func DeleteIntakeRunnerWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeRunnerId string) *wait.AsyncActionHandler[intake.IntakeRunnerResponse] {
	waitConfig := wait.WaiterHelper[intake.IntakeRunnerResponse, string]{
		FetchInstance: client.GetIntakeRunner(ctx, projectId, region, intakeRunnerId).Execute,
		GetState: func(response *intake.IntakeRunnerResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState:                []string{},
		ErrorState:                 []string{},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(15 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-11-13. Use the CreateIntakeWaitHandler or UpdateIntakeWaitHandler instead
func CreateOrUpdateIntakeWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId string) *wait.AsyncActionHandler[intake.IntakeResponse] {
	// TODO: mark function as private after deprecation period
	waitConfig := wait.WaiterHelper[intake.IntakeResponse, string]{
		FetchInstance: client.GetIntake(ctx, projectId, region, intakeId).Execute,
		GetState: func(response *intake.IntakeResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState: []string{INTAKERUNNERRESPONSESTATE_ACTIVE},
		ErrorState:  []string{INTAKERESPONSESTATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateIntakeWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId string) *wait.AsyncActionHandler[intake.IntakeResponse] {
	return CreateOrUpdateIntakeWaitHandler(ctx, client, projectId, region, intakeId)
}

func UpdateIntakeWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId string) *wait.AsyncActionHandler[intake.IntakeResponse] {
	return CreateOrUpdateIntakeWaitHandler(ctx, client, projectId, region, intakeId)
}

func DeleteIntakeWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId string) *wait.AsyncActionHandler[intake.IntakeResponse] {
	waitConfig := wait.WaiterHelper[intake.IntakeResponse, string]{
		FetchInstance: client.GetIntake(ctx, projectId, region, intakeId).Execute,
		GetState: func(response *intake.IntakeResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState:                []string{},
		ErrorState:                 []string{},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-11-13. Use the CreateIntakeUserWaitHandler or UpdateIntakeUserWaitHandler instead
func CreateOrUpdateIntakeUserWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId, intakeUserId string) *wait.AsyncActionHandler[intake.IntakeUserResponse] {
	// TODO: mark function as private after deprecation period
	waitConfig := wait.WaiterHelper[intake.IntakeUserResponse, string]{
		FetchInstance: client.GetIntakeUser(ctx, projectId, region, intakeId, intakeUserId).Execute,
		GetState: func(response *intake.IntakeUserResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState: []string{INTAKEUSERRESPONSESTATE_ACTIVE},
		ErrorState:  []string{},
		// The API does not have a dedicated failure state for this resource,
		// so we rely on the timeout for cases where it never becomes active.
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

func CreateIntakeUserWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId, intakeUserId string) *wait.AsyncActionHandler[intake.IntakeUserResponse] {
	return CreateOrUpdateIntakeUserWaitHandler(ctx, client, projectId, region, intakeId, intakeUserId)
}

func UpdateIntakeUserWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId, intakeUserId string) *wait.AsyncActionHandler[intake.IntakeUserResponse] {
	return CreateOrUpdateIntakeUserWaitHandler(ctx, client, projectId, region, intakeId, intakeUserId)
}

func DeleteIntakeUserWaitHandler(ctx context.Context, client intake.DefaultAPI, projectId, region, intakeId, intakeUserId string) *wait.AsyncActionHandler[intake.IntakeUserResponse] {
	waitConfig := wait.WaiterHelper[intake.IntakeUserResponse, string]{
		FetchInstance: client.GetIntakeUser(ctx, projectId, region, intakeId, intakeUserId).Execute,
		GetState: func(response *intake.IntakeUserResponse) (string, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			return response.State, nil
		},
		ActiveState:                []string{},
		ErrorState:                 []string{},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}
