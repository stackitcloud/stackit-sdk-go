package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

const (
	// Deprecated: CreateSuccess is deprecated and will be removed after 9th November 2025. Use [dns.ZONESTATE_CREATE_SUCCEEDED] or [dns.RECORDSETSTATE_CREATE_SUCCEEDED] instead.
	CreateSuccess = "CREATE_SUCCEEDED"
	// Deprecated: CreateFail is deprecated and will be removed after 9th November 2025. Use [dns.ZONESTATE_CREATE_FAILED] or [dns.RECORDSETSTATE_CREATE_FAILED] instead.
	CreateFail = "CREATE_FAILED"
	// Deprecated: UpdateSuccess is deprecated and will be removed after 9th November 2025. Use [dns.ZONESTATE_UPDATE_SUCCEEDED] or [dns.RECORDSETSTATE_UPDATE_SUCCEEDED] instead.
	UpdateSuccess = "UPDATE_SUCCEEDED"
	// Deprecated: UpdateSuccess is deprecated and will be removed after 9th November 2025. Use [dns.ZONESTATE_UPDATE_FAILED] or [dns.RECORDSETSTATE_UPDATE_FAILED] instead.
	UpdateFail = "UPDATE_FAILED"
	// Deprecated: UpdateSuccess is deprecated and will be removed after 9th November 2025. Use [dns.ZONESTATE_DELETE_SUCCEEDED] or [dns.RECORDSETSTATE_DELETE_SUCCEEDED] instead.
	DeleteSuccess = "DELETE_SUCCEEDED"
	// Deprecated: UpdateSuccess is deprecated and will be removed after 9th November 2025. Use [dns.ZONESTATE_DELETE_FAILED] or [dns.RECORDSETSTATE_DELETE_FAILED] instead.
	DeleteFail = "DELETE_FAILED"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetZoneExecute(ctx context.Context, projectId, zoneId string) (*dns.ZoneResponse, error)
	GetRecordSetExecute(ctx context.Context, projectId, zoneId, rrSetId string) (*dns.RecordSetResponse, error)
}

// CreateZoneWaitHandler will wait for zone creation
func CreateZoneWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.ZoneResponse, err error) {
		s, err := a.GetZoneExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.Zone.Id == nil || s.Zone.State == nil {
			return false, nil, fmt.Errorf("create failed for instance with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == dns.ZONESTATE_CREATE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == dns.ZONESTATE_CREATE_FAILED {
			return true, s, fmt.Errorf("create failed for zone with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateZoneWaitHandler will wait for zone update
func PartialUpdateZoneWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.ZoneResponse, err error) {
		s, err := a.GetZoneExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.Zone.Id == nil || s.Zone.State == nil {
			return false, nil, fmt.Errorf("update failed for instance with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == dns.ZONESTATE_UPDATE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == dns.ZONESTATE_UPDATE_FAILED {
			return true, s, fmt.Errorf("update failed for zone with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteZoneWaitHandler will wait for zone deletion
// returned interface is nil or *ZoneResponseZone
func DeleteZoneWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.ZoneResponse, err error) {
		s, err := a.GetZoneExecute(ctx, projectId, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s.Zone.Id == nil || s.Zone.State == nil {
			return false, nil, fmt.Errorf("delete failed for instance with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == dns.ZONESTATE_DELETE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Zone.Id == instanceId && *s.Zone.State == dns.ZONESTATE_DELETE_FAILED {
			return true, s, fmt.Errorf("delete failed for zone with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateRecordWaitHandler will wait for recordset creation
func CreateRecordSetWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.RecordSetResponse, err error) {
		s, err := a.GetRecordSetExecute(ctx, projectId, instanceId, rrSetId)
		if err != nil {
			return false, nil, err
		}
		if s.Rrset.Id == nil || s.Rrset.State == nil {
			return false, nil, fmt.Errorf("create failed for record set with id %s. The response is not valid: the id or the state are missing", instanceId)
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == dns.RECORDSETSTATE_CREATE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == dns.RECORDSETSTATE_CREATE_FAILED {
			return true, s, fmt.Errorf("create failed for record with id %s", rrSetId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// UpdateRecordWaitHandler will wait for recordset update
func PartialUpdateRecordSetWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.RecordSetResponse, err error) {
		s, err := a.GetRecordSetExecute(ctx, projectId, instanceId, rrSetId)
		if err != nil {
			return false, nil, err
		}
		if s.Rrset.Id == nil || s.Rrset.State == nil {
			return false, nil, fmt.Errorf("update failed for record set with id %s. The response is not valid: the id or the state are missing", rrSetId)
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == dns.RECORDSETSTATE_UPDATE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == dns.RECORDSETSTATE_UPDATE_FAILED {
			return true, s, fmt.Errorf("update failed for record with id %s", rrSetId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteRecordWaitHandler will wait for deletion
// returned interface is nil or *RecordSetResponse
func DeleteRecordSetWaitHandler(ctx context.Context, a APIClientInterface, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.RecordSetResponse, err error) {
		s, err := a.GetRecordSetExecute(ctx, projectId, instanceId, rrSetId)
		if err != nil {
			return false, nil, err
		}
		if s.Rrset.Id == nil || s.Rrset.State == nil {
			return false, nil, fmt.Errorf("delete failed for record set with id %s. The response is not valid: the id or the state are missing", rrSetId)
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == dns.RECORDSETSTATE_DELETE_SUCCEEDED {
			return true, s, nil
		}
		if *s.Rrset.Id == rrSetId && *s.Rrset.State == dns.RECORDSETSTATE_DELETE_FAILED {
			return true, s, fmt.Errorf("delete failed for record with id %s", rrSetId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(2 * time.Minute)
	return handler
}
