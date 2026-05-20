package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	scf "github.com/stackitcloud/stackit-sdk-go/services/scf/v1api"
)

const statusDeletingFailed = "deleting_failed"

// DeleteOrganizationWaitHandler will wait for Organization deletion
func DeleteOrganizationWaitHandler(ctx context.Context, a scf.DefaultAPI, projectId, region, orgId string) *wait.AsyncActionHandler[scf.Organization] {
	waitConfig := wait.WaiterHelper[scf.Organization, string]{
		FetchInstance: a.GetOrganization(ctx, projectId, region, orgId).Execute,
		GetState: func(s *scf.Organization) (string, error) {
			if s == nil {
				return "", errors.New("organization is nil")
			}
			return s.Status, nil
		},
		ActiveState:                []string{},
		ErrorState:                 []string{statusDeletingFailed},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(20 * time.Minute)
	return handler
}
