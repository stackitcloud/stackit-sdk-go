package wait

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	cdn "github.com/stackitcloud/stackit-sdk-go/services/cdn/v1api"
)

type mockSettings struct {
	customDomain               *cdn.CustomDomain
	customDomainProjectId      string
	customDomainDistributionId string
	getCustomDomainError       *oapierror.GenericOpenAPIError

	distribution         *cdn.Distribution
	getDistributionError *oapierror.GenericOpenAPIError
}

func newAPIMock(settings *mockSettings) cdn.DefaultAPI {
	return &cdn.DefaultAPIServiceMock{
		GetDistributionExecuteMock: utils.Ptr(func(_ cdn.ApiGetDistributionRequest) (*cdn.GetDistributionResponse, error) {
			if settings.getDistributionError != nil {
				return nil, settings.getDistributionError
			}

			if settings.distribution != nil {
				return &cdn.GetDistributionResponse{
					Distribution: *settings.distribution,
				}, nil
			}

			return nil, &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusNotFound,
			}
		}),
		GetCustomDomainExecuteMock: utils.Ptr(func(_ cdn.ApiGetCustomDomainRequest) (*cdn.GetCustomDomainResponse, error) {
			if settings.getCustomDomainError != nil {
				return nil, settings.getCustomDomainError
			}

			if settings.customDomain != nil {
				return &cdn.GetCustomDomainResponse{
					CustomDomain: *settings.customDomain,
				}, nil
			}

			return nil, &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusNotFound,
			}
		}),
	}
}

func isNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected to get no error: %v", err)
	}
}

func TestCreateDistributionWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	statusActive := DISTRIBUTIONSTATUS_ACTIVE
	statusUpdating := DISTRIBUTIONSTATUS_UPDATING
	statusCreating := DISTRIBUTIONSTATUS_CREATING
	statusError := DISTRIBUTIONSTATUS_ERROR
	statusDeleting := DISTRIBUTIONSTATUS_DELETING

	mockClientFixture := func(patches ...func(settings *mockSettings)) cdn.DefaultAPI {
		settings := mockSettings{
			distribution: &cdn.Distribution{
				Id:        distributionId,
				ProjectId: projectId,
				Status:    statusActive,
			},
		}

		for _, p := range patches {
			p(&settings)
		}

		return newAPIMock(&settings)
	}

	tests := map[string]struct {
		apiClient            cdn.DefaultAPI
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
				Distribution: cdn.Distribution{
					Id:        distributionId,
					ProjectId: projectId,
					Status:    statusActive,
				},
			},
			errCheck: isNil,
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusUpdating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusCreating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusDeleting
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
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusError
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
			if diff := cmp.Diff(tc.expectedDistribution, dist, cmpopts.IgnoreUnexported(cdn.NullableString{}, cdn.NullableInt64{})); diff != "" {
				t.Fatalf("Unexpected response (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestDeleteDistributionWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	statusActive := DISTRIBUTIONSTATUS_ACTIVE
	statusUpdating := DISTRIBUTIONSTATUS_UPDATING
	statusCreating := DISTRIBUTIONSTATUS_CREATING
	statusError := DISTRIBUTIONSTATUS_ERROR
	statusDeleting := DISTRIBUTIONSTATUS_DELETING

	mockClientFixture := func(patches ...func(settings *mockSettings)) cdn.DefaultAPI {
		settings := mockSettings{
			distribution: &cdn.Distribution{
				Id:        distributionId,
				ProjectId: projectId,
				Status:    statusActive,
			},
		}

		for _, p := range patches {
			p(&settings)
		}

		return newAPIMock(&settings)
	}

	tests := map[string]struct {
		apiClient      cdn.DefaultAPI
		projectId      string
		distributionId string
		errCheck       func(*testing.T, error)
	}{
		"happyPath": {
			apiClient:      newAPIMock(&mockSettings{}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck:       isNil,
		},
		"statusActive": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusUpdating
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusCreating
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusError
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusDeleting
			}),
			projectId:      projectId,
			distributionId: distributionId,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
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
	statusActive := DISTRIBUTIONSTATUS_ACTIVE
	statusUpdating := DISTRIBUTIONSTATUS_UPDATING
	statusCreating := DISTRIBUTIONSTATUS_CREATING
	statusError := DISTRIBUTIONSTATUS_ERROR
	statusDeleting := DISTRIBUTIONSTATUS_DELETING

	mockClientFixture := func(patches ...func(settings *mockSettings)) cdn.DefaultAPI {
		settings := mockSettings{
			distribution: &cdn.Distribution{
				Id:        distributionId,
				ProjectId: projectId,
				Status:    statusActive,
			},
		}

		for _, p := range patches {
			p(&settings)
		}

		return newAPIMock(&settings)
	}

	tests := map[string]struct {
		apiClient            cdn.DefaultAPI
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
				Distribution: cdn.Distribution{
					Id:        distributionId,
					ProjectId: projectId,
					Status:    statusActive,
				},
			},
			errCheck: isNil,
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusUpdating
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			expectedDistribution: nil,
			errCheck: func(t *testing.T, err error) {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusCreating
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
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusDeleting
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
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.distribution.Status = statusError
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
			if diff := cmp.Diff(tc.expectedDistribution, dist, cmpopts.IgnoreUnexported(cdn.NullableString{}, cdn.NullableInt64{})); diff != "" {
				t.Fatalf("Unexpected response (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestCreateCustomDomainWaitHandler(t *testing.T) {
	projectId := "test-project-id"
	distributionId := "test-distribution-id"
	customDomain := "myCustomDomain.com"

	mockClientFixture := func(patches ...func(settings *mockSettings)) cdn.DefaultAPI {
		settings := mockSettings{
			customDomain: &cdn.CustomDomain{
				Name:   customDomain,
				Status: cdn.DOMAINSTATUS_ACTIVE,
			},
			customDomainProjectId:      projectId,
			customDomainDistributionId: distributionId,
			getCustomDomainError:       nil,
		}

		for _, p := range patches {
			p(&settings)
		}

		return newAPIMock(&settings)
	}

	tests := map[string]struct {
		apiClient            cdn.DefaultAPI
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
				Name:   customDomain,
				Status: cdn.DOMAINSTATUS_ACTIVE,
			},
			errCheck: isNil,
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_CREATING
			}),
			projectId:            projectId,
			distributionId:       distributionId,
			customDomain:         customDomain,
			expectedCustomDomain: nil,
			errCheck: func(t *testing.T, err error) {
				if err != nil && errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_DELETING
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
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_ERROR
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
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain = nil
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

	mockClientFixture := func(patches ...func(settings *mockSettings)) cdn.DefaultAPI {
		settings := mockSettings{
			customDomain: &cdn.CustomDomain{
				Name:   customDomain,
				Status: cdn.DOMAINSTATUS_ACTIVE,
			},
			customDomainProjectId:      projectId,
			customDomainDistributionId: distributionId,
			getCustomDomainError:       nil,
		}

		for _, p := range patches {
			p(&settings)
		}

		return newAPIMock(&settings)
	}

	tests := map[string]struct {
		apiClient      cdn.DefaultAPI
		projectId      string
		distributionId string
		customDomain   string
		errCheck       func(*testing.T, error)
	}{
		"happyPath": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain = nil
			}),
			projectId:      projectId,
			customDomain:   customDomain,
			distributionId: distributionId,
			errCheck:       isNil,
		},
		"statusActive": {
			apiClient:      mockClientFixture(),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusUpdating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_UPDATING
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusCreating": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_CREATING
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusError": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_ERROR
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && errors.Is(err, context.DeadlineExceeded) {
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			},
		},
		"statusDeleting": {
			apiClient: mockClientFixture(func(settings *mockSettings) {
				settings.customDomain.Status = cdn.DOMAINSTATUS_DELETING
			}),
			projectId:      projectId,
			distributionId: distributionId,
			customDomain:   customDomain,
			errCheck: func(t *testing.T, err error) {
				if err != nil && errors.Is(err, context.DeadlineExceeded) {
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
