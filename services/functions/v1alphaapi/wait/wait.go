package wait

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	functions "github.com/stackitcloud/stackit-sdk-go/services/functions/v1alphaapi"
)

// CreateRevisionWaitHandler will wait for Revision creation
func CreateRevisionWaitHandler(ctx context.Context, client functions.DefaultAPI, projectId, functionId, revisionId string) *wait.AsyncActionHandler[functions.Revision] {
	waitConfig := wait.WaiterHelper[functions.Revision, functions.RevisionState]{
		FetchInstance: client.GetRevision(ctx, projectId, functionId, revisionId).Execute,
		GetState: func(revision *functions.Revision) (functions.RevisionState, error) {
			if revision.Id == nil || revision.State == nil {
				return functions.REVISIONSTATE_UNKNOWN_DEFAULT_OPEN_API, fmt.Errorf("could not get revision id or state from response")
			}
			if *revision.Id == revisionId && *revision.State == functions.REVISIONSTATE_FAILED {
				if revision.StateError != nil {
					return *revision.State, fmt.Errorf("revision deployment failed: %s", *revision.StateError)
				}

				return *revision.State, fmt.Errorf("revision deployment failed")
			}
			return *revision.State, nil
		},
		ActiveState:                []functions.RevisionState{functions.REVISIONSTATE_ACTIVE},
		ErrorState:                 []functions.RevisionState{functions.REVISIONSTATE_FAILED, functions.REVISIONSTATE_CANCELLED}, //nolint:misspell // british wording
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}
	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(10 * time.Minute)
	return handler
}
