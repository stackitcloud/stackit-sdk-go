package wait

import (
	"context"
	"errors"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/cdn"
)

// Interfaces needed for tests
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
type APIClientInterface interface {
	GetDistributionExecute(ctx context.Context, projectId string, distributionId string) (*cdn.GetDistributionResponse, error)
	GetCustomDomainExecute(ctx context.Context, projectId string, distributionId string, domain string) (*cdn.GetCustomDomainResponse, error)
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateDistributionPoolWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	waitConfig := wait.WaiterHelper[cdn.GetDistributionResponse, cdn.DistributionStatus]{
		FetchInstance: func() (*cdn.GetDistributionResponse, error) {
			return api.GetDistributionExecute(ctx, projectId, distributionId)
		},
		GetState: func(response *cdn.GetDistributionResponse) (cdn.DistributionStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Distribution == nil {
				return "", errors.New("empty distribution")
			}
			if response.Distribution.Status == nil {
				return "", errors.New("empty status")
			}
			return *response.Distribution.Status, nil

		},
		ActiveState: []cdn.DistributionStatus{cdn.DISTRIBUTIONSTATUS_ACTIVE},
		ErrorState:  []cdn.DistributionStatus{cdn.DISTRIBUTIONSTATUS_ERROR, cdn.DISTRIBUTIONSTATUS_DELETING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func UpdateDistributionWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	waitConfig := wait.WaiterHelper[cdn.GetDistributionResponse, cdn.DistributionStatus]{
		FetchInstance: func() (*cdn.GetDistributionResponse, error) {
			return api.GetDistributionExecute(ctx, projectId, distributionId)
		},
		GetState: func(response *cdn.GetDistributionResponse) (cdn.DistributionStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Distribution == nil {
				return "", errors.New("empty distribution")
			}
			if response.Distribution.Status == nil {
				return "", errors.New("empty status")
			}
			return *response.Distribution.Status, nil
		},
		ActiveState: []cdn.DistributionStatus{cdn.DISTRIBUTIONSTATUS_ACTIVE},
		ErrorState:  []cdn.DistributionStatus{cdn.DISTRIBUTIONSTATUS_ERROR, cdn.DISTRIBUTIONSTATUS_DELETING, cdn.DISTRIBUTIONSTATUS_CREATING},
	}
	handler := wait.New(waitConfig.Wait())

	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func DeleteDistributionWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	waitConfig := wait.WaiterHelper[cdn.GetDistributionResponse, cdn.DistributionStatus]{
		FetchInstance: func() (*cdn.GetDistributionResponse, error) {
			return api.GetDistributionExecute(ctx, projectId, distributionId)
		},
		GetState: func(response *cdn.GetDistributionResponse) (cdn.DistributionStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Distribution == nil {
				return "", errors.New("empty distribution")
			}
			if response.Distribution.Status == nil {
				return "", errors.New("empty status")
			}
			return *response.Distribution.Status, nil
		},
		ActiveState: []cdn.DistributionStatus{},
		ErrorState:  []cdn.DistributionStatus{},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateCDNCustomDomainWaitHandler(ctx context.Context, a APIClientInterface, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	waitConfig := wait.WaiterHelper[cdn.CustomDomain, cdn.DomainStatus]{
		FetchInstance: func() (*cdn.CustomDomain, error) {
			resp, err := a.GetCustomDomainExecute(ctx, projectId, distributionId, domain)
			if err != nil || resp == nil || resp.CustomDomain == nil {
				return nil, err
			}
			return resp.CustomDomain, nil
		},
		GetState: func(response *cdn.CustomDomain) (cdn.DomainStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Status == nil {
				return "", errors.New("empty status")
			}
			return *response.Status, nil
		},
		ActiveState: []cdn.DomainStatus{cdn.DOMAINSTATUS_ACTIVE},
		ErrorState:  []cdn.DomainStatus{cdn.DOMAINSTATUS_ERROR, cdn.DOMAINSTATUS_DELETING},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}

// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func DeleteCDNCustomDomainWaitHandler(ctx context.Context, a APIClientInterface, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	waitConfig := wait.WaiterHelper[cdn.CustomDomain, cdn.DomainStatus]{
		FetchInstance: func() (*cdn.CustomDomain, error) {
			resp, err := a.GetCustomDomainExecute(ctx, projectId, distributionId, domain)
			if err != nil || resp == nil || resp.CustomDomain == nil {
				return nil, err
			}
			return resp.CustomDomain, nil
		},
		GetState: func(response *cdn.CustomDomain) (cdn.DomainStatus, error) {
			if response == nil {
				return "", errors.New("empty response")
			}
			if response.Status == nil {
				return "", errors.New("empty status")
			}
			return *response.Status, nil
		},
		ActiveState: []cdn.DomainStatus{},
		ErrorState:  []cdn.DomainStatus{},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
