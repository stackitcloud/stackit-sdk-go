package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"

	functions "github.com/stackitcloud/stackit-sdk-go/services/functions/v1alphaapi"
)

type mockSettings struct {
	createFails    bool
	resourceStates []functions.RevisionState
}

func newAPIMock(settings mockSettings) functions.DefaultAPI {
	return &functions.DefaultAPIServiceMock{
		GetRevisionExecuteMock: utils.Ptr(
			func(_ functions.ApiGetRevisionRequest) (*functions.Revision, error) {
				if settings.createFails {
					return nil, &oapierror.GenericOpenAPIError{
						StatusCode: 500,
					}
				}

				state := settings.resourceStates[len(settings.resourceStates)-1]
				return &functions.Revision{
					Id:         utils.Ptr("rid"),
					FunctionId: utils.Ptr("fid"),
					State:      utils.Ptr(state),
				}, nil
			},
		),
	}
}

func TestCreateRevisionWaitHandler(t *testing.T) {
	tests := []struct {
		desc           string
		resourceStates []functions.RevisionState
		wantErr        bool
		wantResp       bool
		wantState      functions.RevisionState
	}{
		{
			desc: "create_succeeded",
			resourceStates: []functions.RevisionState{
				functions.REVISIONSTATE_ACTIVE,
			},
			wantErr:   false,
			wantResp:  true,
			wantState: functions.REVISIONSTATE_ACTIVE,
		},
		{
			desc: "revision_failed",
			resourceStates: []functions.RevisionState{
				functions.REVISIONSTATE_FAILED,
			},
			wantErr:   true,
			wantResp:  false,
			wantState: functions.REVISIONSTATE_FAILED,
		},
		{
			desc: "creating_then_succeeded",
			resourceStates: []functions.RevisionState{
				functions.REVISIONSTATE_CREATING,
				functions.REVISIONSTATE_ACTIVE,
			},
			wantErr:   false,
			wantResp:  true,
			wantState: functions.REVISIONSTATE_ACTIVE,
		},
		{
			desc: "unknown",
			resourceStates: []functions.RevisionState{
				functions.REVISIONSTATE_UNKNOWN_DEFAULT_OPEN_API,
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:     "create_api_error",
			wantErr:  true,
			wantResp: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					createFails:    tt.wantErr,
					resourceStates: tt.resourceStates,
				})

				var wantRes *functions.Revision
				if tt.wantResp {
					wantRes = &functions.Revision{
						Id:         utils.Ptr("rid"),
						FunctionId: utils.Ptr("fid"),
						State:      utils.Ptr(tt.wantState),
					}
				}

				handler := CreateRevisionWaitHandler(
					context.Background(),
					apiClient,
					"pid",
					"fid",
					"rid",
				)

				gotRes, err := handler.
					SetTimeout(10 * time.Millisecond).
					WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf(
						"handler error = %v, wantErr %v",
						err,
						tt.wantErr,
					)
				}

				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf(
						"handler gotRes = %v, want %v",
						gotRes,
						wantRes,
					)
				}
			})
		})
	}
}
