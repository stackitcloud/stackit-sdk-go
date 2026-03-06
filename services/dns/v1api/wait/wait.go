package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	dns "github.com/stackitcloud/stackit-sdk-go/services/dns/v1api"
)

const (
	ZONESTATE_CREATING         = "CREATING"
	ZONESTATE_CREATE_SUCCEEDED = "CREATE_SUCCEEDED"
	ZONESTATE_CREATE_FAILED    = "CREATE_FAILED"
	ZONESTATE_DELETING         = "DELETING"
	ZONESTATE_DELETE_SUCCEEDED = "DELETE_SUCCEEDED"
	ZONESTATE_DELETE_FAILED    = "DELETE_FAILED"
	ZONESTATE_UPDATING         = "UPDATING"
	ZONESTATE_UPDATE_SUCCEEDED = "UPDATE_SUCCEEDED"
	ZONESTATE_UPDATE_FAILED    = "UPDATE_FAILED"

	RECORDSETSTATE_CREATING         = "CREATING"
	RECORDSETSTATE_CREATE_SUCCEEDED = "CREATE_SUCCEEDED"
	RECORDSETSTATE_CREATE_FAILED    = "CREATE_FAILED"
	RECORDSETSTATE_DELETING         = "DELETING"
	RECORDSETSTATE_DELETE_SUCCEEDED = "DELETE_SUCCEEDED"
	RECORDSETSTATE_DELETE_FAILED    = "DELETE_FAILED"
	RECORDSETSTATE_UPDATING         = "UPDATING"
	RECORDSETSTATE_UPDATE_SUCCEEDED = "UPDATE_SUCCEEDED"
	RECORDSETSTATE_UPDATE_FAILED    = "UPDATE_FAILED"
)

// CreateZoneWaitHandler will wait for zone creation
func CreateZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.ZoneResponse, err error) {
		s, err := a.GetZone(ctx, projectId, instanceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Zone.Id == instanceId && s.Zone.State == ZONESTATE_CREATE_SUCCEEDED {
			return true, s, nil
		}
		if s.Zone.Id == instanceId && s.Zone.State == ZONESTATE_CREATE_FAILED {
			return true, s, fmt.Errorf("create failed for zone with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateZoneWaitHandler will wait for zone update
func PartialUpdateZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.ZoneResponse, err error) {
		s, err := a.GetZone(ctx, projectId, instanceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Zone.Id == instanceId && s.Zone.State == ZONESTATE_UPDATE_SUCCEEDED {
			return true, s, nil
		}
		if s.Zone.Id == instanceId && s.Zone.State == ZONESTATE_UPDATE_FAILED {
			return true, s, fmt.Errorf("update failed for zone with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteZoneWaitHandler will wait for zone deletion
// returned interface is nil or *ZoneResponseZone
func DeleteZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.ZoneResponse, err error) {
		s, err := a.GetZone(ctx, projectId, instanceId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Zone.Id == instanceId && s.Zone.State == ZONESTATE_DELETE_SUCCEEDED {
			return true, s, nil
		}
		if s.Zone.Id == instanceId && s.Zone.State == ZONESTATE_DELETE_FAILED {
			return true, s, fmt.Errorf("delete failed for zone with id %s", instanceId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateRecordWaitHandler will wait for recordset creation
func CreateRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.RecordSetResponse, err error) {
		s, err := a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Rrset.Id == rrSetId && s.Rrset.State == RECORDSETSTATE_CREATE_SUCCEEDED {
			return true, s, nil
		}
		if s.Rrset.Id == rrSetId && s.Rrset.State == RECORDSETSTATE_CREATE_FAILED {
			return true, s, fmt.Errorf("create failed for record with id %s", rrSetId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// UpdateRecordWaitHandler will wait for recordset update
func PartialUpdateRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.RecordSetResponse, err error) {
		s, err := a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Rrset.Id == rrSetId && s.Rrset.State == RECORDSETSTATE_UPDATE_SUCCEEDED {
			return true, s, nil
		}
		if s.Rrset.Id == rrSetId && s.Rrset.State == RECORDSETSTATE_UPDATE_FAILED {
			return true, s, fmt.Errorf("update failed for record with id %s", rrSetId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteRecordWaitHandler will wait for deletion
// returned interface is nil or *RecordSetResponse
func DeleteRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	handler := wait.New(func() (waitFinished bool, response *dns.RecordSetResponse, err error) {
		s, err := a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute()
		if err != nil {
			return false, nil, err
		}
		if s.Rrset.Id == rrSetId && s.Rrset.State == RECORDSETSTATE_DELETE_SUCCEEDED {
			return true, s, nil
		}
		if s.Rrset.Id == rrSetId && s.Rrset.State == RECORDSETSTATE_DELETE_FAILED {
			return true, s, fmt.Errorf("delete failed for record with id %s", rrSetId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(2 * time.Minute)
	return handler
}
