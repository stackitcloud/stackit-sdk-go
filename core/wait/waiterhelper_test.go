package wait

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// Mock types for testing
type MockResource struct {
	Status string
	Error  error
}

func TestWaiterHelper_Wait(t *testing.T) {
	tests := []struct {
		name                       string
		fetchResult                *MockResource
		fetchErr                   error
		activeStates               []string
		errorStates                []string
		deleteHttpErrorStatusCodes []int
		wantFinished               bool
		wantErr                    bool
		wantResponse               *MockResource
	}{
		{
			name:         "Success - Reaches Active State",
			fetchResult:  &MockResource{Status: "READY"},
			activeStates: []string{"READY", "ACTIVE"},
			wantFinished: true,
			wantErr:      false,
			wantResponse: &MockResource{Status: "READY"},
		},
		{
			name:         "Failure - Reaches Error State",
			fetchResult:  &MockResource{Status: "FAILED"},
			errorStates:  []string{"FAILED", "ERROR"},
			wantFinished: true,
			wantErr:      true,
			wantResponse: &MockResource{Status: "FAILED"},
		},
		{
			name:         "Pending - Still In Progress",
			fetchResult:  &MockResource{Status: "CREATING"},
			activeStates: []string{"READY"},
			wantFinished: false,
			wantErr:      false,
			wantResponse: nil,
		},
		{
			name:     "Success - Deletion (403 Forbidden)",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusForbidden},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates: nil,
			wantFinished: true,
			wantErr:      false,
			wantResponse: nil,
		},
		{
			name:     "Success - Deletion (404 Not Found)",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates: nil,
			wantFinished: true,
			wantErr:      false,
			wantResponse: nil,
		},
		{
			name:     "Success - Deletion (410 Gone)",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusGone},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates: nil,
			wantFinished: true,
			wantErr:      false,
			wantResponse: nil,
		},
		{
			name:     "Failure - Set explicit 404 as deletion status code. Get 403 Forbidden",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusForbidden},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates:               nil,
			deleteHttpErrorStatusCodes: []int{http.StatusNotFound},
			wantFinished:               true,
			wantErr:                    true,
			wantResponse:               nil,
		},
		{
			name:     "Failure - Set explicit 404 as deletion status code. Get 410 Gone",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusGone},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates:               nil,
			deleteHttpErrorStatusCodes: []int{http.StatusNotFound},
			wantFinished:               true,
			wantErr:                    true,
			wantResponse:               nil,
		},
		{
			name:     "Failure - Error on fetch instance without active states (400 Bad Request)",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadRequest},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates: nil,
			wantFinished: true,
			wantErr:      true,
		},
		{
			name:         "Success - Error on fetch instance (400 Bad Request)",
			fetchErr:     &oapierror.GenericOpenAPIError{StatusCode: http.StatusForbidden},
			activeStates: []string{"READY"},
			wantFinished: true,
			wantErr:      true,
			wantResponse: nil,
		},
		{
			name:         "Failure - GetState failed",
			fetchResult:  &MockResource{Error: fmt.Errorf("can not read state")},
			activeStates: []string{"READY"},
			wantFinished: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WaiterHelper[MockResource, string]{
				FetchInstance: func() (*MockResource, error) {
					return tt.fetchResult, tt.fetchErr
				},
				GetState: func(m *MockResource) (string, error) {
					return m.Status, m.Error
				},
				DeleteHttpErrorStatusCodes: tt.deleteHttpErrorStatusCodes,
				ActiveState:                tt.activeStates,
				ErrorState:                 tt.errorStates,
			}

			// Get the check function
			checkFunc := w.Wait()

			// Execute the check
			gotFinished, gotResp, gotErr := checkFunc()

			// Check Finished Status
			if gotFinished != tt.wantFinished {
				t.Errorf("Wait() finished = %v, want %v", gotFinished, tt.wantFinished)
			}

			// Check Error Presence
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Wait() error = %v, wantErr %v", gotErr, tt.wantErr)
			}

			// Check Response Data
			if diff := cmp.Diff(tt.wantResponse, gotResp); diff != "" {
				t.Errorf("Wait() response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
