package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	vpn "github.com/stackitcloud/stackit-sdk-go/services/vpn/v1api"
)

func createOrUpdateGatewayWaitHandler(ctx context.Context, a vpn.DefaultAPI, projectId string, region vpn.Region, gatewayId string) *wait.AsyncActionHandler[vpn.GatewayResponse] {
	waitConfig := wait.WaiterHelper[vpn.GatewayResponse, vpn.GatewayStatus]{
		FetchInstance: a.GetGateway(ctx, projectId, region, gatewayId).Execute,
		GetState: func(resp *vpn.GatewayResponse) (vpn.GatewayStatus, error) {
			if resp == nil {
				return "", errors.New("could not get gateway status: response is nil")
			}
			if resp.State == nil {
				return "", errors.New("could not get gateway status: state is nil")
			}
			return *resp.State, nil
		},
		ActiveState: []vpn.GatewayStatus{vpn.GATEWAYSTATUS_READY},
		ErrorState:  []vpn.GatewayStatus{vpn.GATEWAYSTATUS_ERROR, vpn.GATEWAYSTATUS_DELETING},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(45 * time.Minute)
	return handler
}

func CreateGatewayWaitHandler(ctx context.Context, a vpn.DefaultAPI, projectId string, region vpn.Region, gatewayId string) *wait.AsyncActionHandler[vpn.GatewayResponse] {
	return createOrUpdateGatewayWaitHandler(ctx, a, projectId, region, gatewayId)
}

func UpdateGatewayWaitHandler(ctx context.Context, a vpn.DefaultAPI, projectId string, region vpn.Region, gatewayId string) *wait.AsyncActionHandler[vpn.GatewayResponse] {
	return createOrUpdateGatewayWaitHandler(ctx, a, projectId, region, gatewayId)
}

func DeleteGatewayWaitHandler(ctx context.Context, a vpn.DefaultAPI, projectId string, region vpn.Region, gatewayId string) *wait.AsyncActionHandler[vpn.GatewayResponse] {
	waitConfig := wait.WaiterHelper[vpn.GatewayResponse, vpn.GatewayStatus]{
		FetchInstance: a.GetGateway(ctx, projectId, region, gatewayId).Execute,
		GetState: func(resp *vpn.GatewayResponse) (vpn.GatewayStatus, error) {
			if resp == nil {
				return "", errors.New("could not get gateway status: response is nil")
			}
			if resp.State == nil {
				return "", errors.New("could not get gateway status: state is nil")
			}
			return *resp.State, nil
		},
		ErrorState: []vpn.GatewayStatus{vpn.GATEWAYSTATUS_ERROR},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}
