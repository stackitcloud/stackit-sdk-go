package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	cdn "github.com/stackitcloud/stackit-sdk-go/services/cdn/v1api"
)

func CreateDistributionPoolWaitHandler(ctx context.Context, api cdn.DefaultAPI, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistribution(ctx, projectId, distributionId).Execute()
		if err != nil {
			return false, distribution, err
		}
		if distribution == nil {
			return false, nil, fmt.Errorf("create failed for distribution with id %s, the response is nil", distributionId)
		}
		if distribution.Distribution.Id == distributionId {
			switch distribution.Distribution.Status {
			case cdn.DISTRIBUTIONSTATUS_ACTIVE:
				return true, distribution, nil
			case cdn.DISTRIBUTIONSTATUS_CREATING, cdn.DISTRIBUTIONSTATUS_UPDATING:
				return false, nil, nil
			case cdn.DISTRIBUTIONSTATUS_DELETING:
				return true, nil, fmt.Errorf("creating CDN distribution failed")
			case cdn.DISTRIBUTIONSTATUS_ERROR:
				return true, nil, fmt.Errorf("creating CDN distribution failed")
			default:
				return true, nil, fmt.Errorf("CDNDistributionWaitHandler: unexpected status %s", distribution.Distribution.Status)
			}
		}
		return false, nil, nil
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func UpdateDistributionWaitHandler(ctx context.Context, api cdn.DefaultAPI, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		distribution, err = api.GetDistribution(ctx, projectId, distributionId).Execute()
		if err != nil {
			return false, distribution, err
		}
		if distribution == nil {
			return false, nil, fmt.Errorf("update failed for distribution with id %s, the response is nil", distributionId)
		}
		if distribution.Distribution.Id == distributionId {
			switch distribution.Distribution.Status {
			case cdn.DISTRIBUTIONSTATUS_ACTIVE:
				return true, distribution, err
			case cdn.DISTRIBUTIONSTATUS_UPDATING:
				return false, nil, nil
			case cdn.DISTRIBUTIONSTATUS_DELETING:
				return true, nil, fmt.Errorf("updating CDN distribution failed")
			case cdn.DISTRIBUTIONSTATUS_ERROR:
				return true, nil, fmt.Errorf("updating CDN distribution failed")
			default:
				return true, nil, fmt.Errorf("UpdateDistributionWaitHandler: unexpected status %s", distribution.Distribution.Status)
			}
		}

		return false, nil, nil
	})

	handler.SetTimeout(10 * time.Minute)
	return handler
}
func DeleteDistributionWaitHandler(ctx context.Context, api cdn.DefaultAPI, projectId, distributionId string) *wait.AsyncActionHandler[cdn.GetDistributionResponse] {
	handler := wait.New(func() (waitFinished bool, distribution *cdn.GetDistributionResponse, err error) {
		_, err = api.GetDistribution(ctx, projectId, distributionId).Execute()

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

func CreateCDNCustomDomainWaitHandler(ctx context.Context, a cdn.DefaultAPI, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	handler := wait.New(func() (waitFinished bool, response *cdn.CustomDomain, err error) {
		resp, err := a.GetCustomDomain(ctx, projectId, distributionId, domain).Execute()
		if err != nil {
			return false, nil, err
		}
		if resp == nil {
			return false, nil, errors.New("CDNDistributionWaitHandler: response is nil")
		}

		switch resp.CustomDomain.Status {
		case cdn.DOMAINSTATUS_ACTIVE:
			return true, &resp.CustomDomain, nil
		case cdn.DOMAINSTATUS_CREATING, cdn.DOMAINSTATUS_UPDATING:
			return false, nil, nil
		case cdn.DOMAINSTATUS_DELETING:
			return true, nil, fmt.Errorf("creating CDN custom domain failed")
		case cdn.DOMAINSTATUS_ERROR:
			return true, nil, fmt.Errorf("creating CDN custom domain failed")
		default:
			return true, nil, fmt.Errorf("CDNCustomDomainWaitHandler: unexpected status %s", resp.CustomDomain.Status)
		}
	})
	handler.SetTimeout(10 * time.Minute)
	return handler
}

func DeleteCDNCustomDomainWaitHandler(ctx context.Context, a cdn.DefaultAPI, projectId, distributionId, domain string) *wait.AsyncActionHandler[cdn.CustomDomain] {
	handler := wait.New(func() (waitFinished bool, response *cdn.CustomDomain, err error) {
		_, err = a.GetCustomDomain(ctx, projectId, distributionId, domain).Execute()

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
