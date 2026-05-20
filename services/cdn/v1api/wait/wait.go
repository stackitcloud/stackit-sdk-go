package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	cdn "github.com/stackitcloud/stackit-sdk-go/services/cdn/v1api"
)

const (
	DISTRIBUTIONSTATUS_CREATING = "CREATING"
	DISTRIBUTIONSTATUS_ACTIVE   = "ACTIVE"
	DISTRIBUTIONSTATUS_UPDATING = "UPDATING"
	DISTRIBUTIONSTATUS_DELETING = "DELETING"
	DISTRIBUTIONSTATUS_ERROR    = "ERROR"
)

func CreateDistributionPoolWaitHandler(ctx context.Context, api cdn.DefaultAPI, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	waitConfig := wait.WaiterHelper[cdn.GetDistributionResponse, string]{
		FetchInstance: api.GetDistribution(ctx, projectId, distributionId).Execute,
		GetState: func(c *cdn.GetDistributionResponse) (string, error) {
			if c == nil {
				return "", errors.New("empty response")
			}
			return c.Distribution.Status, nil
		},
		ActiveState: []string{DISTRIBUTIONSTATUS_ACTIVE},
		ErrorState:  []string{DISTRIBUTIONSTATUS_ERROR, DISTRIBUTIONSTATUS_DELETING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateDistributionWaitHandler(ctx context.Context, api cdn.DefaultAPI, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	waitConfig := wait.WaiterHelper[cdn.GetDistributionResponse, string]{
		FetchInstance: api.GetDistribution(ctx, projectId, distributionId).Execute,
		GetState: func(c *cdn.GetDistributionResponse) (string, error) {
			if c == nil {
				return "", errors.New("empty response")
			}
			return c.Distribution.Status, nil
		},
		ActiveState: []string{DISTRIBUTIONSTATUS_ACTIVE},
		ErrorState:  []string{DISTRIBUTIONSTATUS_ERROR, DISTRIBUTIONSTATUS_DELETING, DISTRIBUTIONSTATUS_CREATING},
	}
	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteDistributionWaitHandler(ctx context.Context, api cdn.DefaultAPI, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	waitConfig := wait.WaiterHelper[cdn.GetDistributionResponse, string]{
		FetchInstance: api.GetDistribution(ctx, projectId, distributionId).Execute,
		GetState: func(c *cdn.GetDistributionResponse) (string, error) {
			if c == nil {
				return "", errors.New("empty response")
			}
			return c.Distribution.Status, nil
		},
		ActiveState: []string{},
		ErrorState:  []string{},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateCDNCustomDomainWaitHandler(ctx context.Context, a cdn.DefaultAPI, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	waitConfig := wait.WaiterHelper[cdn.CustomDomain, cdn.DomainStatus]{
		FetchInstance: func() (*cdn.CustomDomain, error) {
			resp, err := a.GetCustomDomain(ctx, projectId, distributionId, domain).Execute()
			if err != nil {
				return nil, err
			}
			if resp == nil {
				return nil, errors.New("CDNDistributionWaitHandler: response is nil")
			}
			return &resp.CustomDomain, err
		},
		GetState: func(c *cdn.CustomDomain) (cdn.DomainStatus, error) {
			if c == nil {
				return "", errors.New("empty response")
			}
			return c.Status, nil
		},
		ActiveState: []cdn.DomainStatus{cdn.DOMAINSTATUS_ACTIVE},
		ErrorState:  []cdn.DomainStatus{cdn.DOMAINSTATUS_ERROR, cdn.DOMAINSTATUS_DELETING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteCDNCustomDomainWaitHandler(ctx context.Context, a cdn.DefaultAPI, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	waitConfig := wait.WaiterHelper[cdn.CustomDomain, cdn.DomainStatus]{
		FetchInstance: func() (*cdn.CustomDomain, error) {
			resp, err := a.GetCustomDomain(ctx, projectId, distributionId, domain).Execute()
			if err != nil {
				return nil, err
			}
			if resp == nil {
				return nil, errors.New("CDNDistributionWaitHandler: response is nil")
			}
			return &resp.CustomDomain, err
		},
		GetState: func(c *cdn.CustomDomain) (cdn.DomainStatus, error) {
			if c == nil {
				return "", errors.New("empty response")
			}
			return c.Status, nil
		},
		ActiveState: []cdn.DomainStatus{},
		ErrorState:  []cdn.DomainStatus{},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
