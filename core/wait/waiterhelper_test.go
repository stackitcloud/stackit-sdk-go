package wait

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

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
			wantFinished:               false,
			wantErr:                    true,
			wantResponse:               nil,
		},
		{
			name:     "Failure - Set explicit 404 as deletion status code. Get 410 Gone",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusGone},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates:               nil,
			deleteHttpErrorStatusCodes: []int{http.StatusNotFound},
			wantFinished:               false,
			wantErr:                    true,
			wantResponse:               nil,
		},
		{
			name:     "Failure - Error on fetch instance without active states (400 Bad Request)",
			fetchErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadRequest},
			// If ActiveState is empty, it assumes we are waiting for a deletion
			activeStates: nil,
			wantFinished: false,
			wantErr:      true,
		},
		{
			name:         "Failure - Error on fetch instance (403 Forbidden)",
			fetchErr:     &oapierror.GenericOpenAPIError{StatusCode: http.StatusForbidden},
			activeStates: []string{"READY"},
			wantFinished: false,
			wantErr:      true,
			wantResponse: nil,
		},
		{
			name:         "Failure - GetState failed",
			fetchResult:  &MockResource{Error: fmt.Errorf("can not read state")},
			activeStates: []string{"READY"},
			wantFinished: false,
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

func TestWaiterHelper_WaitWithContext(t *testing.T) {
	type fetchResponse struct {
		res *MockResource
		err error
	}

	tests := []struct {
		name string
		// fetchResponses is the array of responses the mocked endpoint will answer in sequence
		fetchResponses             []fetchResponse
		activeStates               []string
		errorStates                []string
		deleteHttpErrorStatusCodes []int
		wantCalls                  int
		wantErr                    bool
		wantResponse               *MockResource
	}{
		{
			name: "Success - Retryable 502 Gateway Error followed by Active State",
			fetchResponses: []fetchResponse{
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
				{res: &MockResource{Status: "ACTIVE"}, err: nil},
			},
			activeStates: []string{"ACTIVE"},
			errorStates:  []string{"ERROR"},
			wantCalls:    2,
			wantErr:      false,
			wantResponse: &MockResource{Status: "ACTIVE"},
		},
		{
			name: "Success - Retryable 504 Error followed by Active State",
			fetchResponses: []fetchResponse{
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusGatewayTimeout}},
				{res: &MockResource{Status: "ACTIVE"}, err: nil},
			},
			activeStates: []string{"ACTIVE"},
			errorStates:  []string{"ERROR"},
			wantCalls:    2,
			wantErr:      false,
			wantResponse: &MockResource{Status: "ACTIVE"},
		},
		{
			name: "Success - Retryable 502 Gateway Error during Deletion followed by 404",
			fetchResponses: []fetchResponse{
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound}},
			},
			activeStates: nil,
			wantCalls:    2,
			wantErr:      false,
			wantResponse: nil,
		},
		{
			name: "Success - Immediate Active State",
			fetchResponses: []fetchResponse{
				{res: &MockResource{Status: "ACTIVE"}, err: nil},
			},
			activeStates: []string{"ACTIVE"},
			wantCalls:    1,
			wantErr:      false,
			wantResponse: &MockResource{Status: "ACTIVE"},
		},
		{
			name: "Success - Pending State transitioned to Active State",
			fetchResponses: []fetchResponse{
				{res: &MockResource{Status: "CREATING"}, err: nil},
				{res: &MockResource{Status: "ACTIVE"}, err: nil},
			},
			activeStates: []string{"ACTIVE"},
			wantCalls:    2,
			wantErr:      false,
			wantResponse: &MockResource{Status: "ACTIVE"},
		},
		{
			name: "Success - Deletion (404 Not Found)",
			fetchResponses: []fetchResponse{
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound}},
			},
			activeStates: nil,
			wantCalls:    1,
			wantErr:      false,
			wantResponse: nil,
		},
		{
			name: "Failure - Non-retryable HTTP Error (400 Bad Request)",
			fetchResponses: []fetchResponse{
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadRequest}},
			},
			activeStates: []string{"ACTIVE"},
			wantCalls:    1,
			wantErr:      true,
			wantResponse: nil,
		},
		{
			name: "Failure - Retry limit reached for temporary error",
			fetchResponses: []fetchResponse{
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
				{res: nil, err: &oapierror.GenericOpenAPIError{StatusCode: http.StatusBadGateway}},
			},
			activeStates: []string{"ACTIVE"},
			wantCalls:    5,
			wantErr:      true,
			wantResponse: nil,
		},
		{
			name: "Failure - Pending State transitioned to Error State",
			fetchResponses: []fetchResponse{
				{res: &MockResource{Status: "CREATING"}, err: nil},
				{res: &MockResource{Status: "FAILED"}, err: nil},
			},
			activeStates: []string{"ACTIVE"},
			errorStates:  []string{"FAILED"},
			wantCalls:    2,
			wantErr:      true,
			wantResponse: &MockResource{Status: "FAILED"},
		},
		{
			name: "Failure - Non-GenericOpenAPIError on fetch",
			fetchResponses: []fetchResponse{
				{res: nil, err: fmt.Errorf("network connection failure")},
			},
			activeStates: []string{"ACTIVE"},
			wantCalls:    1,
			wantErr:      true,
			wantResponse: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// synctest for the fake clock
			synctest.Test(t, func(t *testing.T) {
				calls := 0
				w := &WaiterHelper[MockResource, string]{
					FetchInstance: func() (*MockResource, error) {
						calls++
						if calls <= len(tt.fetchResponses) {
							resp := tt.fetchResponses[calls-1]
							return resp.res, resp.err
						}
						return nil, fmt.Errorf("unexpected fetch call %d", calls)
					},
					GetState: func(m *MockResource) (string, error) {
						return m.Status, m.Error
					},
					DeleteHttpErrorStatusCodes: tt.deleteHttpErrorStatusCodes,
					ActiveState:                tt.activeStates,
					ErrorState:                 tt.errorStates,
				}

				handler := New(w.Wait()).SetThrottle(10 * time.Millisecond)
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				res, err := handler.WaitWithContext(ctx)

				if (err != nil) != tt.wantErr {
					t.Fatalf("WaitWithContext() error = %v, wantErr %v", err, tt.wantErr)
				}

				if tt.wantCalls > 0 && calls != tt.wantCalls {
					t.Errorf("FetchInstance calls = %d, want %d", calls, tt.wantCalls)
				}

				if diff := cmp.Diff(tt.wantResponse, res); diff != "" {
					t.Errorf("WaitWithContext() response mismatch (-want +got):\n%s", diff)
				}
			})
		})
	}
}
