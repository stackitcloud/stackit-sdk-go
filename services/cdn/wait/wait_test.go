package wait

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/services/cdn"
)

type mockApiClient struct {
	distributions              []*cdn.Distribution
	getDistributionError       *oapierror.GenericOpenAPIError
	customDomains              []*cdn.CustomDomain
	customDomainProjectId      string
	customDomainDistributionId string
	getCustomDomainError       *oapierror.GenericOpenAPIError
}

func isNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected to get no error: %v", err)
	}
}

func (m *mockApiClient) GetDistributionExecute(_ context.Context, projectId, distributionId string) (*cdn.GetDistributionResponse, error) {
	if m.getDistributionError != nil {
		return nil, m.getDistributionError
	}
	for _, d := range m.distributions {
		if *d.ProjectId == projectId && *d.Id == distributionId {
			return &cdn.GetDistributionResponse{
				Distribution: d,
			}, nil
		}
	}
	return nil, &oapierror.GenericOpenAPIError{
		StatusCode: http.StatusNotFound,
	}
}

func (m *mockApiClient) GetCustomDomainExecute(_ context.Context, projectId, distributionId, domain string) (*cdn.GetCustomDomainResponse, error) {
	if m.getCustomDomainError != nil {
		return nil, m.getCustomDomainError
	}
	if projectId == m.customDomainProjectId && distributionId == m.customDomainDistributionId {
		for _, d := range m.customDomains {
			if *d.Name == domain {
				return &cdn.GetCustomDomainResponse{
					CustomDomain: d,
				}, nil
			}
		}
	}
	return nil, &oapierror.GenericOpenAPIError{
		StatusCode: http.StatusNotFound,
	}
}

func TestCreateDistributionWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	statusActive := DistributionStatusActive
	statusUpdating := DistributionStatusUpdating
	statusCreating := DistributionStatusCreating
	statusError := DistributionStatusError
	statusDeleting := DistributionStatusDeleting

	mockClientFixture := func(patches ...func(tc *mockApiClient)) *mockApiClient {
		client := &mockApiClient{
			distributions: []*cdn.Distribution{
				{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusActive,
				},
			},
			getDistributionError: nil,
		}
		for _, p := range patches {
			p(client)
		}
		return client
	}
	tests := map[string]struct {
		apiClient            *mockApiClient
		projectId            string
		distributionId       string
		expectedDistribution *cdn.GetDistributionResponse
		errCheck             func(*testing.T, error)
	}{
		"happyPath": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			distributionId: distributionId,
			expectedDistribution: &cdn.GetDistributionResponse{
				Distribution: &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusActive,
				},
			},
			errCheck: isNil,
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusUpdating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusCreating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusDeleting
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "creating CDN distribution failed") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusError
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "creating CDN distribution failed") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := CreateDistributionPoolWaitHandler(context.Background(), tc.apiClient, tc.projectId, tc.distributionId)
			dist, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())
			tc.errCheck(t, err)
			if diff := cmp.Diff(tc.expectedDistribution, dist); diff != "" {
				t.Fatalf("Unexpected response (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestDeleteDistributionWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	statusActive := DistributionStatusActive
	statusUpdating := DistributionStatusUpdating
	statusCreating := DistributionStatusCreating
	statusError := DistributionStatusError
	statusDeleting := DistributionStatusDeleting

	mockClientFixture := func(patches ...func(tc *mockApiClient)) *mockApiClient {
		client := &mockApiClient{
			distributions:        []*cdn.Distribution{},
			getDistributionError: nil,
		}
		for _, p := range patches {
			p(client)
		}
		return client
	}
	tests := map[string]struct {
		apiClient      *mockApiClient
		projectId      string
		distributionId string
		errCheck       func(*testing.T, error)
	}{
		"happyPath": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck:       isNil,
		},
		"statusActive": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions = append(tc.distributions, &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusActive,
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions = append(tc.distributions, &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusUpdating,
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions = append(tc.distributions, &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusCreating,
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions = append(tc.distributions, &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusError,
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions = append(tc.distributions, &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusDeleting,
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := DeleteDistributionWaitHandler(context.Background(), tc.apiClient, tc.projectId, tc.distributionId)
			dist, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())
			tc.errCheck(t, err)
			if dist != nil {
				t.Fatalf("Distribution not deleted")
			}
		})
	}
}

func TestUpdateDistributionWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	statusActive := DistributionStatusActive
	statusUpdating := DistributionStatusUpdating
	statusCreating := DistributionStatusCreating
	statusError := DistributionStatusError
	statusDeleting := DistributionStatusDeleting

	mockClientFixture := func(patches ...func(tc *mockApiClient)) *mockApiClient {
		client := &mockApiClient{
			distributions: []*cdn.Distribution{
				{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusActive,
				},
			},
			getDistributionError: nil,
		}
		for _, p := range patches {
			p(client)
		}
		return client
	}
	tests := map[string]struct {
		apiClient            *mockApiClient
		projectId            string
		distributionId       string
		expectedDistribution *cdn.GetDistributionResponse
		errCheck             func(*testing.T, error)
	}{
		"happyPath": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			distributionId: distributionId,
			expectedDistribution: &cdn.GetDistributionResponse{
				Distribution: &cdn.Distribution{
					Id:        &distributionId,
					ProjectId: &projectId,
					Status:    &statusActive,
				},
			},
			errCheck: isNil,
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusUpdating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusCreating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "unexpected status CREATING") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusDeleting
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "updating CDN distribution failed") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.distributions[0].Status = &statusError
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "updating CDN distribution failed") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := UpdateDistributionWaitHandler(context.Background(), tc.apiClient, tc.projectId, tc.distributionId)
			dist, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())
			tc.errCheck(t, err)
			if diff := cmp.Diff(tc.expectedDistribution, dist); diff != "" {
				t.Fatalf("Unexpected response (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestCreateCustomDomainWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	customDomain := "myCustomDomain.com"

	mockClientFixture := func(patches ...func(tc *mockApiClient)) *mockApiClient {
		client := &mockApiClient{
			customDomains: []*cdn.CustomDomain{
				{
					Name:   &customDomain,
					Status: cdn.DOMAINSTATUS_ACTIVE.Ptr(),
				},
			},
			customDomainProjectId:      projectId,
			customDomainDistributionId: distributionId,
			getCustomDomainError:       nil,
		}
		for _, p := range patches {
			p(client)
		}
		return client
	}
	tests := map[string]struct {
		apiClient            *mockApiClient
		projectId            string
		distributionId       string
		customDomain         string
		expectedCustomDomain *cdn.CustomDomain
		errCheck             func(*testing.T, error)
	}{
		"happyPath": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			expectedCustomDomain: &cdn.CustomDomain{
				Name:   &customDomain,
				Status: cdn.DOMAINSTATUS_ACTIVE.Ptr(),
			},
			errCheck: isNil,
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains[0].Status = cdn.DOMAINSTATUS_CREATING.Ptr()
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			customDomain:         customDomain,
			expectedCustomDomain: nil,
			errCheck: func(t *testing.T, err error) {
				if err != nil && err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains[0].Status = cdn.DOMAINSTATUS_DELETING.Ptr()
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			customDomain:         customDomain,
			expectedCustomDomain: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "creating CDN custom domain failed") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains[0].Status = cdn.DOMAINSTATUS_ERROR.Ptr()
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			customDomain:         customDomain,
			expectedCustomDomain: nil,
			errCheck: func(t *testing.T, err error) {
				if strings.Contains(err.Error(), "creating CDN custom domain failed") {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"CustomDomainMissing": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains = []*cdn.CustomDomain{}
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			customDomain:         customDomain,
			expectedCustomDomain: nil,
			errCheck: func(t *testing.T, err error) {
				oapiErr, ok := err.(*oapierror.GenericOpenAPIError) //nolint:errorlint //complaining that error.As should be used to catch wrapped errors, but this error should not be wrapped
				if ok && oapiErr.StatusCode == http.StatusNotFound {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := CreateCDNCustomDomainWaitHandler(context.Background(), tc.apiClient, tc.projectId, tc.distributionId, tc.customDomain)
			customDomain, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())
			tc.errCheck(t, err)
			if diff := cmp.Diff(tc.expectedCustomDomain, customDomain); diff != "" {
				t.Fatalf("Unexpected response (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestDeleteCustomDomainHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	customDomain := "myCustomDomain.com"

	mockClientFixture := func(patches ...func(tc *mockApiClient)) *mockApiClient {
		client := &mockApiClient{
			customDomains:              []*cdn.CustomDomain{},
			customDomainProjectId:      projectId,
			customDomainDistributionId: distributionId,
			getCustomDomainError:       nil,
		}
		for _, p := range patches {
			p(client)
		}
		return client
	}
	tests := map[string]struct {
		apiClient      *mockApiClient
		projectId      string
		distributionId string
		customDomain   string
		errCheck       func(*testing.T, error)
	}{
		"happyPath": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			customDomain:   customDomain,
			distributionId: distributionId,
			errCheck:       isNil,
		},
		"statusActive": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains = append(tc.customDomains, &cdn.CustomDomain{
					Name:   &customDomain,
					Status: cdn.DOMAINSTATUS_ACTIVE.Ptr(),
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains = append(tc.customDomains, &cdn.CustomDomain{
					Name:   &customDomain,
					Status: cdn.DOMAINSTATUS_UPDATING.Ptr(),
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains = append(tc.customDomains, &cdn.CustomDomain{
					Name:   &customDomain,
					Status: cdn.DOMAINSTATUS_CREATING.Ptr(),
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains = append(tc.customDomains, &cdn.CustomDomain{
					Name:   &customDomain,
					Status: cdn.DOMAINSTATUS_ERROR.Ptr(),
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(tc *mockApiClient) {
				tc.customDomains = append(tc.customDomains, &cdn.CustomDomain{
					Name:   &customDomain,
					Status: cdn.DOMAINSTATUS_DELETING.Ptr(),
				})
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && err.Error() == "WaitWithContext() has timed out" {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := DeleteCDNCustomDomainWaitHandler(context.Background(), tc.apiClient, tc.projectId, tc.distributionId, tc.customDomain)
			customDomain, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())
			tc.errCheck(t, err)
			if customDomain != nil {
				t.Fatalf("Custom domain not deleted")
			}
		})
	}
}
