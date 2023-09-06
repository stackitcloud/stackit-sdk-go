package dns

import (
	"context"
	"fmt"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
)

const (
	CreateSuccess = "CREATE_SUCCEEDED"
	CreateFail    = "CREATE_FAILED"
	UpdateSuccess = "UPDATE_SUCCEEDED"
	UpdateFail    = "UPDATE_FAILED"
	DeleteSuccess = "DELETE_SUCCEEDED"
	DeleteFail    = "DELETE_FAILED"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetZoneExecute(ctx context.Context, projectId, zoneId string) (*ZoneResponse, error)
	GetRecordSetExecute(ctx context.Context, projectId, zoneId, rrSetId string) (*RecordSetResponse, error)
}

func handleError(reqErr error) (res interface{}, done bool, err error) {
	oapiErr, ok := reqErr.(*GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
	if !ok {
		return nil, false, fmt.Errorf("could not convert error to GenericOpenApiError")
	}
	// Some APIs may return temporary errors and the request should be retried
	if utils.Contains(wait.RetryHttpErrorStatusCodes, oapiErr.statusCode) {
		return nil, false, nil
	}
	return nil, false, reqErr
}

// CreateZoneWaitHandler will wait for creation
// returned interface is nil or *ZoneResponseZone
func CreateZoneWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetZoneExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s.Zone.Id == nil || s.Zone.State == nil {
			return s, false, fmt.Errorf("create failed for instance with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == CreateSuccess {
			return s, true, nil
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == CreateFail {
			return s, true, fmt.Errorf("create failed for zone with id %s", instanceId)
		}
		return s, false, nil
	})
}

// UpdateZoneWaitHandler will wait for update
// returned interface is nil or *ZoneResponseZone
func UpdateZoneWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetZoneExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s.Zone.Id == nil || s.Zone.State == nil {
			return s, false, fmt.Errorf("update failed for instance with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == UpdateSuccess {
			return s, true, nil
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == UpdateFail {
			return s, true, fmt.Errorf("update failed for zone with id %s", instanceId)
		}
		return s, false, nil
	})
}

// DeleteZoneWaitHandler will wait for delete
// returned interface is nil or *ZoneResponseZone
func DeleteZoneWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetZoneExecute(ctx, projectId, instanceId)
		if err != nil {
			return handleError(err)
		}
		if s.Zone.Id == nil || s.Zone.State == nil {
			return s, false, fmt.Errorf("delete failed for instance with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == DeleteSuccess {
			return s, true, nil
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == DeleteFail {
			return s, true, fmt.Errorf("delete failed for zone with id %s", instanceId)
		}
		return s, false, nil
	})
}

// CreateRecordWaitHandler will wait for creation
// returned interface is nil or *RecordSetResponse
func CreateRecordSetWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId, rrSetId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetRecordSetExecute(ctx, projectId, instanceId, rrSetId)
		if err != nil {
			return handleError(err)
		}
		if s.Rrset.Id == nil || s.Rrset.State == nil {
			return s, false, fmt.Errorf("create failed for record set with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == CreateSuccess {
			return s, true, nil
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == CreateFail {
			return s, true, fmt.Errorf("create failed for record with id %s", rrSetId)
		}
		return s, false, nil
	})
}

// UpdateRecordWaitHandler will wait for update
// returned interface is nil or *RecordSetResponse
func UpdateRecordSetWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId, rrSetId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetRecordSetExecute(ctx, projectId, instanceId, rrSetId)
		if err != nil {
			return handleError(err)
		}
		if s.Rrset.Id == nil || s.Rrset.State == nil {
			return s, false, fmt.Errorf("update failed for record set with id %s. The response is not valid: the id or the state are missing", rrSetId)
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == UpdateSuccess {
			return s, true, nil
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == UpdateFail {
			return s, true, fmt.Errorf("update failed for record with id %s", rrSetId)
		}
		return s, false, nil
	})
}

// DeleteRecordWaitHandler will wait for deletion
// returned interface is nil or *RecordSetResponse
func DeleteRecordSetWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId, rrSetId string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := a.GetRecordSetExecute(ctx, projectId, instanceId, rrSetId)
		if err != nil {
			return handleError(err)
		}
		if s.Rrset.Id == nil || s.Rrset.State == nil {
			return s, false, fmt.Errorf("delete failed for record set with id %s. The response is not valid: the id or the state are missing", rrSetId)
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == DeleteSuccess {
			return s, true, nil
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == DeleteFail {
			return s, true, fmt.Errorf("delete failed for record with id %s", rrSetId)
		}
		return s, false, nil
	})
}
