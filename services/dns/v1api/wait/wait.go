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
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_CREATING = dns.ZONESTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_CREATE_SUCCEEDED = dns.ZONESTATE_CREATE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_CREATE_FAILED = dns.ZONESTATE_CREATE_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_DELETING = dns.ZONESTATE_DELETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_DELETE_SUCCEEDED = dns.ZONESTATE_DELETE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_DELETE_FAILED = dns.ZONESTATE_DELETE_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_UPDATING = dns.ZONESTATE_UPDATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_UPDATE_SUCCEEDED = dns.ZONESTATE_UPDATE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	ZONESTATE_UPDATE_FAILED = dns.ZONESTATE_UPDATE_FAILED

	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_CREATING = dns.RECORDSETSTATE_CREATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_CREATE_SUCCEEDED = dns.RECORDSETSTATE_CREATE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_CREATE_FAILED = dns.RECORDSETSTATE_CREATE_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_DELETING = dns.RECORDSETSTATE_DELETING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_DELETE_SUCCEEDED = dns.RECORDSETSTATE_DELETE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_DELETE_FAILED = dns.RECORDSETSTATE_DELETE_FAILED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_UPDATING = dns.RECORDSETSTATE_UPDATING
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_UPDATE_SUCCEEDED = dns.RECORDSETSTATE_UPDATE_SUCCEEDED
	// Deprecated: symbol is not used anymore, use the packages enum instead, will be removed 2026-12, use `go fix` for automatic fixing
	//go:fix inline
	RECORDSETSTATE_UPDATE_FAILED = dns.RECORDSETSTATE_CREATE_FAILED
)

// CreateZoneWaitHandler will wait for zone creation
func CreateZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	waitConfig := wait.WaiterHelper[dns.ZoneResponse, dns.ZoneState]{
		FetchInstance: a.GetZone(ctx, projectId, instanceId).Execute,
		GetState: func(d *dns.ZoneResponse) (dns.ZoneState, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Zone.State, nil
		},
		ActiveState: []dns.ZoneState{dns.ZONESTATE_CREATE_SUCCEEDED},
		ErrorState:  []dns.ZoneState{dns.ZONESTATE_CREATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// PartialUpdateZoneWaitHandler will wait for zone update
func PartialUpdateZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	waitConfig := wait.WaiterHelper[dns.ZoneResponse, dns.ZoneState]{
		FetchInstance: a.GetZone(ctx, projectId, instanceId).Execute,
		GetState: func(d *dns.ZoneResponse) (dns.ZoneState, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Zone.State, nil
		},
		ActiveState: []dns.ZoneState{dns.ZONESTATE_UPDATE_SUCCEEDED},
		ErrorState:  []dns.ZoneState{dns.ZONESTATE_UPDATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// DeleteZoneWaitHandler will wait for zone deletion
// returned interface is nil or *ZoneResponseZone
func DeleteZoneWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId string) *wait.AsyncActionHandler[dns.ZoneResponse] {
	waitConfig := wait.WaiterHelper[dns.ZoneResponse, dns.ZoneState]{
		FetchInstance: a.GetZone(ctx, projectId, instanceId).Execute,
		GetState: func(d *dns.ZoneResponse) (dns.ZoneState, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Zone.State, nil
		},
		ActiveState: []dns.ZoneState{dns.ZONESTATE_DELETE_SUCCEEDED},
		ErrorState:  []dns.ZoneState{dns.ZONESTATE_DELETE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// CreateRecordSetWaitHandler will wait for recordset creation
func CreateRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	waitConfig := wait.WaiterHelper[dns.RecordSetResponse, dns.RecordSetState]{
		FetchInstance: a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute,
		GetState: func(d *dns.RecordSetResponse) (dns.RecordSetState, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Rrset.State, nil
		},
		ActiveState: []dns.RecordSetState{dns.RECORDSETSTATE_CREATE_SUCCEEDED},
		ErrorState:  []dns.RecordSetState{dns.RECORDSETSTATE_CREATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// PartialUpdateRecordSetWaitHandler will wait for recordset update
func PartialUpdateRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	waitConfig := wait.WaiterHelper[dns.RecordSetResponse, dns.RecordSetState]{
		FetchInstance: a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute,
		GetState: func(d *dns.RecordSetResponse) (dns.RecordSetState, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Rrset.State, nil
		},
		ActiveState: []dns.RecordSetState{dns.RECORDSETSTATE_UPDATE_SUCCEEDED},
		ErrorState:  []dns.RecordSetState{dns.RECORDSETSTATE_UPDATE_FAILED},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(1 * time.Minute)
	return handler
}

// DeleteRecordSetWaitHandler will wait for deletion
// returned interface is nil or *RecordSetResponse
func DeleteRecordSetWaitHandler(ctx context.Context, a dns.DefaultAPI, projectId, instanceId, rrSetId string) *wait.AsyncActionHandler[dns.RecordSetResponse] {
	waitConfig := wait.WaiterHelper[dns.RecordSetResponse, dns.RecordSetState]{
		FetchInstance: a.GetRecordSet(ctx, projectId, instanceId, rrSetId).Execute,
		GetState: func(d *dns.RecordSetResponse) (dns.RecordSetState, error) {
			if d == nil {
				return "", errors.New("empty response")
			}
			return d.Rrset.State, nil
		},
		ActiveState:                []dns.RecordSetState{dns.RECORDSETSTATE_DELETE_SUCCEEDED},
		ErrorState:                 []dns.RecordSetState{dns.RECORDSETSTATE_DELETE_FAILED},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(2 * time.Minute)
	return handler
}
