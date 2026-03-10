package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	scf "github.com/stackitcloud/stackit-sdk-go/services/scf/v1api"
)

const statusDeletingFailed = "deleting_failed"

// DeleteOrganizationWaitHandler will wait for Organization deletion
func DeleteOrganizationWaitHandler(ctx context.Context, a scf.DefaultAPI, projectId, region, orgId string) *wait.AsyncActionHandler[scf.Organization] {
	handler := wait.New(func() (waitFinished bool, response *scf.Organization, err error) {
		s, err := a.GetOrganization(ctx, projectId, region, orgId).Execute()
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
		if s.Status == statusDeletingFailed {
			return true, nil, fmt.Errorf("delete failed for Organization with id %s", orgId)
		}
		return false, s, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}
