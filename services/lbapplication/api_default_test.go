/*
Application Load Balancer API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package lbapplication

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

func Test_lbapplication_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateCredentials", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/credentials"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CreateCredentialsResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		var xRequestID string
		createCredentialsPayload := CreateCredentialsPayload{}

		resp, reqErr := apiClient.CreateCredentials(context.Background(), projectId).XRequestID(xRequestID).CreateCredentialsPayload(createCredentialsPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateLoadBalancer", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/load-balancers"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := LoadBalancer{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		var xRequestID string
		createLoadBalancerPayload := CreateLoadBalancerPayload{}

		resp, reqErr := apiClient.CreateLoadBalancer(context.Background(), projectId).XRequestID(xRequestID).CreateLoadBalancerPayload(createLoadBalancerPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteCredentials", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/credentials/{credentialsRef}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		credentialsRefValue := "credentialsRef"
		path = strings.Replace(path, "{"+"credentialsRef"+"}", url.PathEscape(ParameterValueToString(credentialsRefValue, "credentialsRef")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := map[string]interface{}{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		credentialsRef := "credentialsRef"

		resp, reqErr := apiClient.DeleteCredentials(context.Background(), projectId, credentialsRef).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteLoadBalancer", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/load-balancers/{name}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		nameValue := "name"
		path = strings.Replace(path, "{"+"name"+"}", url.PathEscape(ParameterValueToString(nameValue, "name")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := map[string]interface{}{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		name := "name"

		resp, reqErr := apiClient.DeleteLoadBalancer(context.Background(), projectId, name).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DisableService", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := map[string]interface{}{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"

		resp, reqErr := apiClient.DisableService(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService EnableService", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := map[string]interface{}{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		var xRequestID string

		resp, reqErr := apiClient.EnableService(context.Background(), projectId).XRequestID(xRequestID).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetCredentials", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/credentials/{credentialsRef}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		credentialsRefValue := "credentialsRef"
		path = strings.Replace(path, "{"+"credentialsRef"+"}", url.PathEscape(ParameterValueToString(credentialsRefValue, "credentialsRef")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetCredentialsResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		credentialsRef := "credentialsRef"

		resp, reqErr := apiClient.GetCredentials(context.Background(), projectId, credentialsRef).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetLoadBalancer", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/load-balancers/{name}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		nameValue := "name"
		path = strings.Replace(path, "{"+"name"+"}", url.PathEscape(ParameterValueToString(nameValue, "name")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := LoadBalancer{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		name := "name"

		resp, reqErr := apiClient.GetLoadBalancer(context.Background(), projectId, name).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetQuota", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/quota"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetQuotaResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"

		resp, reqErr := apiClient.GetQuota(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetServiceStatus", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetServiceStatusResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"

		resp, reqErr := apiClient.GetServiceStatus(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListCredentials", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/credentials"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListCredentialsResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"

		resp, reqErr := apiClient.ListCredentials(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListLoadBalancers", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/load-balancers"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListLoadBalancersResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"

		resp, reqErr := apiClient.ListLoadBalancers(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListPlans", func(t *testing.T) {
		path := "/v1beta/plans"

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListPlansResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		resp, reqErr := apiClient.ListPlans(context.Background()).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService UpdateCredentials", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/credentials/{credentialsRef}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		credentialsRefValue := "credentialsRef"
		path = strings.Replace(path, "{"+"credentialsRef"+"}", url.PathEscape(ParameterValueToString(credentialsRefValue, "credentialsRef")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := UpdateCredentialsResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		credentialsRef := "credentialsRef"
		updateCredentialsPayload := UpdateCredentialsPayload{}

		resp, reqErr := apiClient.UpdateCredentials(context.Background(), projectId, credentialsRef).UpdateCredentialsPayload(updateCredentialsPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService UpdateLoadBalancer", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/load-balancers/{name}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		nameValue := "name"
		path = strings.Replace(path, "{"+"name"+"}", url.PathEscape(ParameterValueToString(nameValue, "name")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := LoadBalancer{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		name := "name"
		updateLoadBalancerPayload := UpdateLoadBalancerPayload{}

		resp, reqErr := apiClient.UpdateLoadBalancer(context.Background(), projectId, name).UpdateLoadBalancerPayload(updateLoadBalancerPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService UpdateTargetPool", func(t *testing.T) {
		path := "/v1beta/projects/{projectId}/load-balancers/{name}/target-pools/{targetPoolName}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		nameValue := "name"
		path = strings.Replace(path, "{"+"name"+"}", url.PathEscape(ParameterValueToString(nameValue, "name")), -1)
		targetPoolNameValue := "targetPoolName"
		path = strings.Replace(path, "{"+"targetPoolName"+"}", url.PathEscape(ParameterValueToString(targetPoolNameValue, "targetPoolName")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := TargetPool{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for lbapplication_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		name := "name"
		targetPoolName := "targetPoolName"
		updateTargetPoolPayload := UpdateTargetPoolPayload{}

		resp, reqErr := apiClient.UpdateTargetPool(context.Background(), projectId, name, targetPoolName).UpdateTargetPoolPayload(updateTargetPoolPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

}