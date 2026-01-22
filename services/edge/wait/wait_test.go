package wait

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/edge"
)

const (
	testRegion     = "eu01"
	testProjectId  = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	testInstanceId = "xxxxxxxx-xxxxxxx"
	testName       = "xxxxxxxx"
)

type mockApiClient struct {
	instanceStatus edge.InstanceStatus
	instanceError  error

	deleteErrors    []error
	deleteCallCount int

	resourceErrors    []error
	resourceCallCount int
	mockShouldTimeout bool
}

func (a *mockApiClient) getDeleteError() error {
	var err error
	if a.deleteCallCount < len(a.deleteErrors) {
		err = a.deleteErrors[a.deleteCallCount]
	}
	a.deleteCallCount++
	return err
}

func (a *mockApiClient) getResourceError() error {
	if a.mockShouldTimeout {
		if len(a.resourceErrors) > 0 {
			return a.resourceErrors[0]
		}
		return nil
	}
	var err error
	if a.resourceCallCount < len(a.resourceErrors) {
		err = a.resourceErrors[a.resourceCallCount]
	}
	a.resourceCallCount++
	return err
}

func (a *mockApiClient) GetInstanceExecute(_ context.Context, _, _, _ string) (*edge.Instance, error) {
	if a.deleteErrors != nil {
		err := a.getDeleteError()
		if err != nil {
			return nil, err
		}
		return &edge.Instance{}, nil
	}

	if a.instanceError != nil {
		return nil, a.instanceError
	}
	return &edge.Instance{Status: &a.instanceStatus}, nil
}

func (a *mockApiClient) GetInstanceByNameExecute(ctx context.Context, _, _, _ string) (*edge.Instance, error) {
	return a.GetInstanceExecute(ctx, "", "", "")
}

type mockGetKubeconfigRequest struct {
	client *mockApiClient
}

func (r *mockGetKubeconfigRequest) ExpirationSeconds(_ int64) edge.ApiGetKubeconfigByInstanceIdRequest {
	return r
}

func (r *mockGetKubeconfigRequest) Execute() (*edge.Kubeconfig, error) {
	if err := r.client.getResourceError(); err != nil {
		return nil, err
	}
	return &edge.Kubeconfig{}, nil
}

func (a *mockApiClient) GetKubeconfigByInstanceId(context.Context, string, string, string) edge.ApiGetKubeconfigByInstanceIdRequest {
	return &mockGetKubeconfigRequest{client: a}
}

type mockGetKubeconfigByNameRequest struct {
	client *mockApiClient
}

func (r *mockGetKubeconfigByNameRequest) ExpirationSeconds(_ int64) edge.ApiGetKubeconfigByInstanceNameRequest {
	return r
}

func (r *mockGetKubeconfigByNameRequest) Execute() (*edge.Kubeconfig, error) {
	if err := r.client.getResourceError(); err != nil {
		return nil, err
	}
	return &edge.Kubeconfig{}, nil
}

func (a *mockApiClient) GetKubeconfigByInstanceName(context.Context, string, string, string) edge.ApiGetKubeconfigByInstanceNameRequest {
	return &mockGetKubeconfigByNameRequest{client: a}
}

type mockGetTokenRequest struct {
	client *mockApiClient
}

func (r *mockGetTokenRequest) ExpirationSeconds(_ int64) edge.ApiGetTokenByInstanceIdRequest {
	return r
}

func (r *mockGetTokenRequest) Execute() (*edge.Token, error) {
	if err := r.client.getResourceError(); err != nil {
		return nil, err
	}
	return &edge.Token{}, nil
}

func (a *mockApiClient) GetTokenByInstanceId(context.Context, string, string, string) edge.ApiGetTokenByInstanceIdRequest {
	return &mockGetTokenRequest{client: a}
}

type mockGetTokenByNameRequest struct {
	client *mockApiClient
}

func (r *mockGetTokenByNameRequest) ExpirationSeconds(_ int64) edge.ApiGetTokenByInstanceNameRequest {
	return r
}

func (r *mockGetTokenByNameRequest) Execute() (*edge.Token, error) {
	if err := r.client.getResourceError(); err != nil {
		return nil, err
	}
	return &edge.Token{}, nil
}

func (a *mockApiClient) GetTokenByInstanceName(context.Context, string, string, string) edge.ApiGetTokenByInstanceNameRequest {
	return &mockGetTokenByNameRequest{client: a}
}

var createOrUpdateInstanceTests = []struct {
	desc           string
	shouldFail     bool // This will be used to set instanceError
	instanceStatus edge.InstanceStatus
	wantErr        error
}{
	{
		desc:           "successful creation",
		shouldFail:     false,
		instanceStatus: edge.INSTANCESTATUS_ACTIVE,
		wantErr:        nil,
	},
	{
		desc:           "timeout during reconciliation",
		shouldFail:     false,
		instanceStatus: edge.INSTANCESTATUS_RECONCILING,
		wantErr:        errors.New("WaitWithContext() has timed out"),
	},
	{
		desc:           "failed creation",
		shouldFail:     false,
		instanceStatus: edge.INSTANCESTATUS_ERROR,
		wantErr:        errors.New("instance creation failed"),
	},
	{
		desc:       "API fails",
		shouldFail: true,
		wantErr:    &oapierror.GenericOpenAPIError{StatusCode: http.StatusInternalServerError},
	},
}

var deleteInstanceTests = []struct {
	desc           string
	errorsToReturn []error
	wantErr        error
}{
	{
		desc: "successful deletion",
		errorsToReturn: []error{
			nil,
			&oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
		},
		wantErr: nil,
	},
	{
		desc:           "timeout during deletion",
		errorsToReturn: []error{},
		wantErr:        errors.New("WaitWithContext() has timed out"),
	},
	{
		desc: "API fails with status 500 error",
		errorsToReturn: []error{
			&oapierror.GenericOpenAPIError{StatusCode: http.StatusInternalServerError},
		},
		wantErr: &oapierror.GenericOpenAPIError{StatusCode: http.StatusInternalServerError},
	},
}

// This test table is for handlers that retry on 404, which is currently required due to incorrect active instance status
var kubeconfigOrTokenTests = []struct {
	desc              string
	instanceStatus    edge.InstanceStatus
	instanceError     error
	errorsToReturn    []error
	mockShouldTimeout bool
	wantErr           error
}{
	{
		desc:           "success",
		instanceStatus: edge.INSTANCESTATUS_ACTIVE,
		errorsToReturn: []error{nil},
		wantErr:        nil,
	},
	{
		desc:           "success_on_reconciling",
		instanceStatus: edge.INSTANCESTATUS_RECONCILING,
		errorsToReturn: []error{nil},
		wantErr:        nil,
	},
	{
		desc:           "retry_on_404",
		instanceStatus: edge.INSTANCESTATUS_ACTIVE,
		errorsToReturn: []error{
			&oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
			&oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
			&oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
			nil,
		},
		wantErr: nil,
	},
	{
		desc:              "timeout",
		instanceStatus:    edge.INSTANCESTATUS_ACTIVE,
		errorsToReturn:    []error{&oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound}},
		mockShouldTimeout: true,
		wantErr:           errors.New("WaitWithContext() has timed out"),
	},
	{
		desc:          "instance_not_found",
		instanceError: &oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
		wantErr:       &oapierror.GenericOpenAPIError{StatusCode: http.StatusNotFound},
	},
	{
		desc:           "unusable_instance_status",
		instanceStatus: edge.INSTANCESTATUS_ERROR,
		wantErr:        errors.New("cannot use instance"),
	},
}

func checkError(t *testing.T, gotErr, wantErr error) {
	t.Helper()
	if wantErr == nil {
		if gotErr != nil {
			t.Fatalf("handler returned an unexpected error: %v", gotErr)
		}
		return
	}

	if gotErr == nil {
		t.Fatalf("handler expected error %q, but got nil", wantErr)
	}

	var wantOapiErr *oapierror.GenericOpenAPIError
	if errors.As(wantErr, &wantOapiErr) {
		var gotOapiErr *oapierror.GenericOpenAPIError
		if !errors.As(gotErr, &gotOapiErr) {
			t.Fatalf("handler error type = %T, want %T", gotErr, wantErr)
		}
		if wantOapiErr.StatusCode != 0 && gotOapiErr.StatusCode != wantOapiErr.StatusCode {
			t.Fatalf("handler error status code = %d, want %d", gotOapiErr.StatusCode, wantOapiErr.StatusCode)
		}
		return
	}

	if !strings.Contains(gotErr.Error(), wantErr.Error()) {
		t.Fatalf("handler error = %q, want to contain %q", gotErr.Error(), wantErr.Error())
	}
}

func executeHandlerAndGetError[T any](handler *wait.AsyncActionHandler[T]) error {
	_, err := handler.
		SetTimeout(20 * time.Millisecond).
		SetThrottle(1 * time.Millisecond).
		WaitWithContext(context.Background())
	return err
}

func TestCreateOrUpdateInstanceWaitHandler(t *testing.T) {
	for _, tc := range createOrUpdateInstanceTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				instanceStatus: tc.instanceStatus,
			}
			if tc.shouldFail {
				apiClient.instanceError = &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusInternalServerError,
				}
			}
			handler := CreateOrUpdateInstanceWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testInstanceId)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestCreateOrUpdateInstanceByNameWaitHandler(t *testing.T) {
	for _, tc := range createOrUpdateInstanceTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				instanceStatus: tc.instanceStatus,
			}
			if tc.shouldFail {
				apiClient.instanceError = &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusInternalServerError,
				}
			}
			handler := CreateOrUpdateInstanceByNameWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testName)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	for _, tc := range deleteInstanceTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				deleteErrors: tc.errorsToReturn,
			}
			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testInstanceId)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestDeleteInstanceByNameWaitHandler(t *testing.T) {
	for _, tc := range deleteInstanceTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				deleteErrors: tc.errorsToReturn,
			}
			handler := DeleteInstanceByNameWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testName)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestKubeconfigWaitHandler(t *testing.T) {
	for _, tc := range kubeconfigOrTokenTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				instanceStatus:    tc.instanceStatus,
				instanceError:     tc.instanceError,
				resourceErrors:    tc.errorsToReturn,
				mockShouldTimeout: tc.mockShouldTimeout,
			}
			handler := KubeconfigWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testInstanceId, nil)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestKubeconfigByInstanceNameWaitHandler(t *testing.T) {
	for _, tc := range kubeconfigOrTokenTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				instanceStatus:    tc.instanceStatus,
				instanceError:     tc.instanceError,
				resourceErrors:    tc.errorsToReturn,
				mockShouldTimeout: tc.mockShouldTimeout,
			}
			handler := KubeconfigByInstanceNameWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testName, nil)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestTokenWaitHandler(t *testing.T) {
	for _, tc := range kubeconfigOrTokenTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				instanceStatus:    tc.instanceStatus,
				instanceError:     tc.instanceError,
				resourceErrors:    tc.errorsToReturn,
				mockShouldTimeout: tc.mockShouldTimeout,
			}
			handler := TokenWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testInstanceId, nil)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}

func TestTokenByInstanceNameWaitHandler(t *testing.T) {
	for _, tc := range kubeconfigOrTokenTests {
		t.Run(tc.desc, func(t *testing.T) {
			apiClient := &mockApiClient{
				instanceStatus:    tc.instanceStatus,
				instanceError:     tc.instanceError,
				resourceErrors:    tc.errorsToReturn,
				mockShouldTimeout: tc.mockShouldTimeout,
			}
			handler := TokenByInstanceNameWaitHandler(context.Background(), apiClient, testProjectId, testRegion, testName, nil)
			err := executeHandlerAndGetError(handler)
			checkError(t, err, tc.wantErr)
		})
	}
}
