// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
package wait

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex"
)

const (
	// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
	InstanceStateEmpty = ""
	// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
	InstanceStateProcessing = "Progressing"
	// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
	InstanceStateUnknown = "Unknown"
	// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
	InstanceStateSuccess = "Ready"
	// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
	InstanceStateFailed = "Failed"
)

// Interface needed for tests
//
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
type APIClientInstanceInterface interface {
	GetInstanceExecute(ctx context.Context, projectId, region, instanceId string) (*sqlserverflex.GetInstanceResponse, error)
}

// CreateInstanceWaitHandler will wait for instance creation
//
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func CreateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *sqlserverflex.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, region, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Id == nil || *s.Id != instanceId || s.State == nil {
			return false, nil, nil
		}
		switch strings.ToLower(string(*s.State)) {
		case strings.ToLower(string(sqlserverflex.STATE_READY)):
			return true, s, nil
		case strings.ToLower(string(sqlserverflex.STATE_UNKNOWN)), strings.ToLower(string(sqlserverflex.STATE_FAILURE)):
			return true, s, fmt.Errorf("create failed for instance with id %s", instanceId)
		default:
			return false, s, nil
		}
	})
	handler.SetTimeout(45 * time.Minute)
	handler.SetSleepBeforeWait(5 * time.Second)
	return handler
}

// UpdateInstanceWaitHandler will wait for instance update
//
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func UpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	handler := wait.New(func() (waitFinished bool, response *sqlserverflex.GetInstanceResponse, err error) {
		s, err := a.GetInstanceExecute(ctx, projectId, region, instanceId)
		if err != nil {
			return false, nil, err
		}
		if s == nil || s.Id == nil || *s.Id != instanceId || s.State == nil {
			return false, nil, nil
		}
		switch strings.ToLower(string(*s.State)) {
		case strings.ToLower(string(sqlserverflex.STATE_READY)):
			return true, s, nil
		case strings.ToLower(string(sqlserverflex.STATE_UNKNOWN)), strings.ToLower(string(sqlserverflex.STATE_FAILURE)):
			return true, s, fmt.Errorf("update failed for instance with id %s", instanceId)
		default:
			return false, s, nil
		}
	})
	handler.SetSleepBeforeWait(2 * time.Second)
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// PartialUpdateInstanceWaitHandler will wait for instance update
//
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func PartialUpdateInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, region, instanceId string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse] {
	return UpdateInstanceWaitHandler(ctx, a, projectId, region, instanceId)
}

// DeleteInstanceWaitHandler will wait for instance deletion
//
// Deprecated: Will be removed after 2026-09-30. Move to the packages generated for each available API version instead
func DeleteInstanceWaitHandler(ctx context.Context, a APIClientInstanceInterface, projectId, region, instanceId string) *wait.AsyncActionHandler[struct{}] {
	handler := wait.New(func() (waitFinished bool, response *struct{}, err error) {
		_, err = a.GetInstanceExecute(ctx, projectId, region, instanceId)
		if err == nil {
			return false, nil, nil
		}
		var oapiErr *oapierror.GenericOpenAPIError
		ok := errors.As(err, &oapiErr)
		if !ok {
			return false, nil, fmt.Errorf("could not convert error to oapierror.GenericOpenAPIError")
		}
		if oapiErr.StatusCode != http.StatusNotFound {
			return false, nil, err
		}
		return true, nil, nil
	})
	handler.SetTimeout(15 * time.Minute)
	return handler
}
