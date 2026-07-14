package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2alpha1api"
)

const (
	CreateSuccess = "CREATED"
	UpdateSuccess = "UPDATED"
	FailedStatus  = "FAILED"
)

// CreateVPCRegionWaitHandler will wait for network area region creation
func CreateVPCRegionWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, vpcId, region string) *wait.AsyncActionHandler[iaas.RegionalVPC] {
	waitConfig := wait.WaiterHelper[iaas.RegionalVPC, string]{
		FetchInstance: a.GetVPCRegion(ctx, projectId, vpcId, region).Execute,
		GetState: func(i *iaas.RegionalVPC) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			if i.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *i.Status, nil
		},
		ActiveState: []string{CreateSuccess},
		ErrorState:  []string{FailedStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteVPCRegionWaitHandler will wait for network area region deletion
func DeleteVPCRegionWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, vpcId, region string) *wait.AsyncActionHandler[iaas.RegionalVPC] {
	waitConfig := wait.WaiterHelper[iaas.RegionalVPC, string]{
		FetchInstance: a.GetVPCRegion(ctx, projectId, vpcId, region).Execute,
		GetState: func(i *iaas.RegionalVPC) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			if i.Status == nil {
				return "", errors.New("status is missing in response")
			}
			return *i.Status, nil
		},
		ActiveState: []string{},
		ErrorState:  []string{FailedStatus},

		// The IaaS API response with a 400 if the regional network area configuration doesn't exist because of some compatible
		// issues to v1. When v1 is deprecated, they probably will respond with 404
		DeleteHttpErrorStatusCodes: []int{http.StatusBadRequest, http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

func createOrUpdateNetworkRangeWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, vpcId, region, networkRangeId string) *wait.AsyncActionHandler[iaas.VPCNetworkRange] {
	waitConfig := wait.WaiterHelper[iaas.VPCNetworkRange, string]{
		FetchInstance: a.GetVPCNetworkRange(ctx, projectId, vpcId, region, networkRangeId).Execute,
		GetState: func(i *iaas.VPCNetworkRange) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			if i.VPCNetworkRangeIPv4 != nil {
				if i.VPCNetworkRangeIPv4.Status == nil {
					return "", errors.New("status is missing in response")
				}
				return *i.VPCNetworkRangeIPv4.Status, nil
			}
			if i.VPCNetworkRangeIPv6 != nil {
				if i.VPCNetworkRangeIPv6.Status == nil {
					return "", errors.New("status is missing in response")
				}
				return *i.VPCNetworkRangeIPv6.Status, nil
			}
			return "", errors.New("unknown status in response")
		},
		// Usually "CREATED" should be enough as active state. The IaaS API uses "CREATED" always as active state, even after an update.
		// To prevent potential issues in the future, we set "UPDATED" also as active state.
		ActiveState: []string{CreateSuccess, UpdateSuccess},
		ErrorState:  []string{FailedStatus},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(5 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}

func CreateVPCNetworkRangeWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, vpcId, region, networkRangeId string) *wait.AsyncActionHandler[iaas.VPCNetworkRange] {
	return createOrUpdateNetworkRangeWaitHandler(ctx, a, projectId, vpcId, region, networkRangeId)
}

func UpdateVPCNetworkRangeWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, vpcId, region, networkRangeId string) *wait.AsyncActionHandler[iaas.VPCNetworkRange] {
	return createOrUpdateNetworkRangeWaitHandler(ctx, a, projectId, vpcId, region, networkRangeId)
}

func DeleteVPCNetworkRangeWaitHandler(ctx context.Context, a iaas.DefaultAPI, projectId, vpcId, region, networkRangeId string) *wait.AsyncActionHandler[iaas.VPCNetworkRange] {
	waitConfig := wait.WaiterHelper[iaas.VPCNetworkRange, string]{
		FetchInstance: a.GetVPCNetworkRange(ctx, projectId, vpcId, region, networkRangeId).Execute,
		GetState: func(i *iaas.VPCNetworkRange) (string, error) {
			if i == nil {
				return "", errors.New("empty response")
			}
			if i.VPCNetworkRangeIPv4 != nil {
				if i.VPCNetworkRangeIPv4.Status == nil {
					return "", errors.New("status is missing in response")
				}
				return *i.VPCNetworkRangeIPv4.Status, nil
			}
			if i.VPCNetworkRangeIPv6 != nil {
				if i.VPCNetworkRangeIPv6.Status == nil {
					return "", errors.New("status is missing in response")
				}
				return *i.VPCNetworkRangeIPv6.Status, nil
			}
			return "", errors.New("unknown status in response")
		},
		ErrorState:                 []string{FailedStatus},
		DeleteHttpErrorStatusCodes: nil, // Use defaults
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetSleepBeforeWait(5 * time.Second)
	handler.SetTimeout(30 * time.Minute)
	return handler
}
