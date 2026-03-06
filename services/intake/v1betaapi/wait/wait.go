package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
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

func CreateOrUpdateIntakeRunnerWaitHandler(ctx context.Context, a intake.DefaultAPI, projectId, region, intakeRunnerId string) *wait.AsyncActionHandler[intake.IntakeRunnerResponse] {
	handler := wait.New(func() (waitFinished bool, response *intake.IntakeRunnerResponse, err error) {
		runner, err := a.GetIntakeRunner(ctx, projectId, region, intakeRunnerId).Execute()
		if err != nil {
			return false, nil, err
		}

		if runner == nil {
			return false, nil, fmt.Errorf("API returned a nil response for Intake Runner %s", intakeRunnerId)
		}

		if runner.Id == intakeRunnerId && runner.State == INTAKERUNNERRESPONSESTATE_ACTIVE {
			return true, runner, nil
		}

		// The API does not have a dedicated failure state for this resource,
		// so we rely on the timeout for cases where it never becomes active.
		return false, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

func DeleteIntakeRunnerWaitHandler(ctx context.Context, a intake.DefaultAPI, projectId, region, intakeRunnerId string) *wait.AsyncActionHandler[intake.IntakeRunnerResponse] {
	handler := wait.New(func() (waitFinished bool, response *intake.IntakeRunnerResponse, err error) {
		_, err = a.GetIntakeRunner(ctx, projectId, region, intakeRunnerId).Execute()
		if err == nil {
			// Resource still exists
			return false, nil, nil
		}

		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if oapiError.StatusCode == http.StatusNotFound {
				// Success: Resource is gone
				return true, nil, nil
			}
		}
		// An unexpected error occurred
		return false, nil, err
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}

func CreateOrUpdateIntakeWaitHandler(ctx context.Context, a intake.DefaultAPI, projectId, region, intakeId string) *wait.AsyncActionHandler[intake.IntakeResponse] {
	handler := wait.New(func() (waitFinished bool, response *intake.IntakeResponse, err error) {
		ik, err := a.GetIntake(ctx, projectId, region, intakeId).Execute()
		if err != nil {
			return false, nil, err
		}

		if ik == nil {
			return false, nil, fmt.Errorf("API returned a nil response for Intake %s", intakeId)
		}

		if ik.Id == intakeId && ik.State == INTAKERESPONSESTATE_ACTIVE {
			return true, ik, nil
		}

		if ik.Id == intakeId && ik.State == INTAKERESPONSESTATE_FAILED {
			return true, ik, fmt.Errorf("create/update failed for Intake %s", intakeId)
		}

		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteIntakeWaitHandler(ctx context.Context, a intake.DefaultAPI, projectId, region, intakeId string) *wait.AsyncActionHandler[intake.IntakeResponse] {
	handler := wait.New(func() (waitFinished bool, response *intake.IntakeResponse, err error) {
		_, err = a.GetIntake(ctx, projectId, region, intakeId).Execute()
		if err == nil {
			return false, nil, nil
		}

		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if oapiError.StatusCode == http.StatusNotFound {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateOrUpdateIntakeUserWaitHandler(ctx context.Context, a intake.DefaultAPI, projectId, region, intakeId, intakeUserId string) *wait.AsyncActionHandler[intake.IntakeUserResponse] {
	handler := wait.New(func() (waitFinished bool, response *intake.IntakeUserResponse, err error) {
		user, err := a.GetIntakeUser(ctx, projectId, region, intakeId, intakeUserId).Execute()
		if err != nil {
			return false, nil, err
		}

		if user == nil {
			return false, nil, fmt.Errorf("API returned a nil response for Intake User %s", intakeUserId)
		}

		if user.Id == intakeUserId && user.State == INTAKEUSERRESPONSESTATE_ACTIVE {
			return true, user, nil
		}

		// The API does not have a dedicated failure state for this resource, we rely on the timeout for cases where
		// it never becomes active.
		return false, nil, nil
	})
	handler.SetTimeout(5 * time.Minute)
	return handler
}

func DeleteIntakeUserWaitHandler(ctx context.Context, a intake.DefaultAPI, projectId, region, intakeId, intakeUserId string) *wait.AsyncActionHandler[intake.IntakeUserResponse] {
	handler := wait.New(func() (waitFinished bool, response *intake.IntakeUserResponse, err error) {
		_, err = a.GetIntakeUser(ctx, projectId, region, intakeId, intakeUserId).Execute()
		if err == nil {
			return false, nil, nil
		}

		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if oapiError.StatusCode == http.StatusNotFound {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(5 * time.Minute)
	return handler
}
