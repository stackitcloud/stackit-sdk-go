package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/scf"
)

const statusDeletingFailed = "deleting_failed"

// Interfaces needed for tests
type APIClientInterface interface {
	GetOrganizationExecute(ctx context.Context, projectId, region, orgId string) (*scf.Organization, error)
}

// DeleteOrganizationWaitHandler will wait for Organization deletion
func DeleteOrganizationWaitHandler(ctx context.Context, a APIClientInterface, projectId, region, orgId string) *wait.AsyncActionHandler[scf.Organization] {
	handler := wait.New(func() (waitFinished bool, response *scf.Organization, err error) {
		s, err := a.GetOrganizationExecute(ctx, projectId, region, orgId)
		if err != nil {
			var oapiErr *oapierror.GenericOpenAPIError
			ok := errors.As(err, &oapiErr)
			if ok && oapiErr.StatusCode == http.StatusNotFound {
				return true, s, nil
			}
			return false, s, err
		}
		if s == nil {
			return false, nil, errors.New("organization is nil")
		}
		if *s.Status == statusDeletingFailed {
			return true, nil, fmt.Errorf("delete failed for Organization with id %s", orgId)
		}
		return false, s, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}
