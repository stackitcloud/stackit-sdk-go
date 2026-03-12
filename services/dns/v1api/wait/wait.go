package wait

import (
	"context"
	"errors"
	"net/http"
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
	waitConfig := wait.WaiterHelper[dns.ZoneResponse, string]{
		FetchInstance: a.GetZone(ctx, projectId, instanceId).Execute,
		GetState: func(d *dns.ZoneResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Zone.State, nil
		},
		ActiveState: []string{ZONESTATE_CREATE_SUCCEEDED},
		ErrorState:  []string{ZONESTATE_CREATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateZoneWaitHandler will wait for zone update
func PartialUpdateZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	waitConfig := wait.WaiterHelper[dns.ZoneResponse, string]{
		FetchInstance: a.GetZone(ctx, projectId, instanceId).Execute,
		GetState: func(d *dns.ZoneResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Zone.State, nil
		},
		ActiveState: []string{ZONESTATE_UPDATE_SUCCEEDED},
		ErrorState:  []string{ZONESTATE_UPDATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteZoneWaitHandler will wait for zone deletion
// returned interface is nil or *ZoneResponseZone
func DeleteZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	waitConfig := wait.WaiterHelper[dns.ZoneResponse, string]{
		FetchInstance: a.GetZone(ctx, projectId, instanceId).Execute,
		GetState: func(d *dns.ZoneResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Zone.State, nil
		},
		ActiveState: []string{ZONESTATE_DELETE_SUCCEEDED},
		ErrorState:  []string{ZONESTATE_DELETE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateRecordSetWaitHandler will wait for recordset creation
func CreateRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	waitConfig := wait.WaiterHelper[dns.RecordSetResponse, string]{
		FetchInstance: a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute,
		GetState: func(d *dns.RecordSetResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Rrset.State, nil
		},
		ActiveState: []string{RECORDSETSTATE_CREATE_SUCCEEDED},
		ErrorState:  []string{RECORDSETSTATE_CREATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// PartialUpdateRecordSetWaitHandler will wait for recordset update
func PartialUpdateRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	waitConfig := wait.WaiterHelper[dns.RecordSetResponse, string]{
		FetchInstance: a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute,
		GetState: func(d *dns.RecordSetResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Rrset.State, nil
		},
		ActiveState: []string{RECORDSETSTATE_UPDATE_SUCCEEDED},
		ErrorState:  []string{RECORDSETSTATE_UPDATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteRecordSetWaitHandler will wait for deletion
// returned interface is nil or *RecordSetResponse
func DeleteRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	waitConfig := wait.WaiterHelper[dns.RecordSetResponse, string]{
		FetchInstance: a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute,
		GetState: func(d *dns.RecordSetResponse) (string, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Rrset.State, nil
		},
		ActiveState:                []string{RECORDSETSTATE_DELETE_SUCCEEDED},
		ErrorState:                 []string{RECORDSETSTATE_DELETE_FAILED},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(2 * time.Minute)
	return handler
}
