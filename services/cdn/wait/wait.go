package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/cdn"
)

const (
	// Deprecated: DistributionStatusCreating is deprecated and will be removed after 14th November 2025. Use [cdn.DISTRIBUTIONSTATUS_CREATING] instead.
	DistributionStatusCreating = "CREATING"
	// Deprecated: DistributionStatusActive is deprecated and will be removed after 14th November 2025. Use [cdn.DISTRIBUTIONSTATUS_ACTIVE] instead.
	DistributionStatusActive = "ACTIVE"
	// Deprecated: DistributionStatusUpdating is deprecated and will be removed after 14th November 2025. Use [cdn.DISTRIBUTIONSTATUS_UPDATING] instead.
	DistributionStatusUpdating = "UPDATING"
	// Deprecated: DistributionStatusDeleting is deprecated and will be removed after 14th November 2025. Use [cdn.DISTRIBUTIONSTATUS_DELETING] instead.
	DistributionStatusDeleting = "DELETING"
	// Deprecated: DistributionStatusError is deprecated and will be removed after 14th November 2025. Use [cdn.DISTRIBUTIONSTATUS_ERROR] instead.
	DistributionStatusError = "ERROR"
)

// Interfaces needed for tests
type APIClientInterface interface {
	GetDistributionExecute(ctx context.Context, projectId string, distributionId string) (*cdn.GetDistributionResponse, error)
	GetCustomDomainExecute(ctx context.Context, projectId string, distributionId string, domain string) (*cdn.GetCustomDomainResponse, error)
}

func CreateDistributionPoolWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistributionExecute(ctx, projectId, distributionId)
		if err != nil {
			return false, distribution, err
		}
		if distribution == nil ||
			distribution.Distribution == nil ||
			distribution.Distribution.Id == nil ||
			distribution.Distribution.Status == nil {
			return false, distribution, fmt.Errorf("create failed for distribution with id %s, the response is not valid (state missing)", distributionId)
		}
		if *distribution.Distribution.Id == distributionId {
			switch *distribution.Distribution.Status {
			case cdn.DISTRIBUTIONSTATUS_ACTIVE:
				return true, distribution, nil
			case cdn.DISTRIBUTIONSTATUS_CREATING, cdn.DISTRIBUTIONSTATUS_UPDATING:
				return false, nil, nil
			case cdn.DISTRIBUTIONSTATUS_DELETING:
				return true, nil, fmt.Errorf("creating CDN distribution failed")
			case cdn.DISTRIBUTIONSTATUS_ERROR:
				return true, nil, fmt.Errorf("creating CDN distribution failed")
			default:
				return true, nil, fmt.Errorf("CDNDistributionWaitHandler: unexpected status %s", *distribution.Distribution.Status)
			}
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateDistributionWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistributionExecute(ctx, projectId, distributionId)
		if err != nil {
			return false, distribution, err
		}
		if distribution == nil ||
			distribution.Distribution == nil ||
			distribution.Distribution.Id == nil ||
			distribution.Distribution.Status == nil {
			return false, distribution, fmt.Errorf("update failed for distribution with id %s, the response is not valid (state missing)", distributionId)
		}
		if *distribution.Distribution.Id == distributionId {
			switch *distribution.Distribution.Status {
			case cdn.DISTRIBUTIONSTATUS_ACTIVE:
				return true, distribution, err
			case cdn.DISTRIBUTIONSTATUS_UPDATING:
				return false, nil, nil
			case cdn.DISTRIBUTIONSTATUS_DELETING:
				return true, nil, fmt.Errorf("updating CDN distribution failed")
			case cdn.DISTRIBUTIONSTATUS_ERROR:
				return true, nil, fmt.Errorf("updating CDN distribution failed")
			default:
				return true, nil, fmt.Errorf("UpdateDistributionWaitHandler: unexpected status %s", *distribution.Distribution.Status)
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteDistributionWaitHandler(ctx context.Context, api APIClientInterface, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		_, err = api.GetDistributionExecute(ctx, projectId, distributionId)

		// the distribution is still gettable, e.g. not deleted
		if err == nil {
			return false, nil, nil
		}
		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
				return true, nil, nil
			}
		}

		return false, nil, err
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func CreateCDNCustomDomainWaitHandler(ctx context.Context, a APIClientInterface, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	handler := wait.New(func() (waitFinished bool, response *cdn.CustomDomain, err error) {
		resp, err := a.GetCustomDomainExecute(ctx, projectId, distributionId, domain)
		if err != nil {
			return false, nil, err
		}
		if resp == nil || resp.CustomDomain == nil || resp.CustomDomain.Status == nil {
			return false, nil, errors.New("CDNDistributionWaitHandler: status or custom domain missing in response")
		}

		switch *resp.CustomDomain.Status {
		case cdn.DOMAINSTATUS_ACTIVE:
			return true, resp.CustomDomain, nil
		case cdn.DOMAINSTATUS_CREATING, cdn.DOMAINSTATUS_UPDATING:
			return false, nil, nil
		case cdn.DOMAINSTATUS_DELETING:
			return true, nil, fmt.Errorf("creating CDN custom domain failed")
		case cdn.DOMAINSTATUS_ERROR:
			return true, nil, fmt.Errorf("creating CDN custom domain failed")
		default:
			return true, nil, fmt.Errorf("CDNCustomDomainWaitHandler: unexpected status %s", *resp.CustomDomain.Status)
		}
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteCDNCustomDomainWaitHandler(ctx context.Context, a APIClientInterface, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	handler := wait.New(func() (waitFinished bool, response *cdn.CustomDomain, err error) {
		_, err = a.GetCustomDomainExecute(ctx, projectId, distributionId, domain)

		// the custom domain is still gettable, e.g. not deleted
		if err == nil {
			return false, nil, nil
		}
		var oapiError *oapierror.GenericOpenAPIError
		if errors.As(err, &oapiError) {
			if statusCode := oapiError.StatusCode; statusCode == http.StatusNotFound || statusCode == http.StatusGone {
				return true, nil, nil
			}
		}
		return false, nil, err
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}
