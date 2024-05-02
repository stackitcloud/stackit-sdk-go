/*
STACKIT LogMe API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package logme

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

func Test_logme_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateCredentials", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/credentials"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CredentialsResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"

		resp, reqErr := apiClient.CreateCredentials(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CreateInstanceResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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
		createInstancePayload := CreateInstancePayload{}

		resp, reqErr := apiClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteCredentials", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/credentials/{credentialsId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		credentialsIdValue := "credentialsId"
		path = strings.Replace(path, "{"+"credentialsId"+"}", url.PathEscape(ParameterValueToString(credentialsIdValue, "credentialsId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"
		credentialsId := "credentialsId"

		reqErr := apiClient.DeleteCredentials(context.Background(), projectId, instanceId, credentialsId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

	t.Run("Test DefaultApiService DeleteInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"

		reqErr := apiClient.DeleteInstance(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

	t.Run("Test DefaultApiService GetCredentials", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/credentials/{credentialsId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		credentialsIdValue := "credentialsId"
		path = strings.Replace(path, "{"+"credentialsId"+"}", url.PathEscape(ParameterValueToString(credentialsIdValue, "credentialsId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CredentialsResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"
		credentialsId := "credentialsId"

		resp, reqErr := apiClient.GetCredentials(context.Background(), projectId, instanceId, credentialsId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Instance{}
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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"

		resp, reqErr := apiClient.GetInstance(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService InstanceBackupsGet", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/backups"
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListBackupsResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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

		instanceId := "instanceId"
		projectId := "projectId"

		resp, reqErr := apiClient.InstanceBackupsGet(context.Background(), instanceId, projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService InstanceMetricsGet", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/metrics"
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListMetricsResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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

		instanceId := "instanceId"
		projectId := "projectId"

		resp, reqErr := apiClient.InstanceMetricsGet(context.Background(), instanceId, projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListCredentials", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/credentials"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"

		resp, reqErr := apiClient.ListCredentials(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListInstances", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListInstancesResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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

		resp, reqErr := apiClient.ListInstances(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListOfferings", func(t *testing.T) {
		path := "/v1/projects/{projectId}/offerings"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListOfferingsResponse{}
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
					Description: "Localhost for logme_DefaultApi",
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

		resp, reqErr := apiClient.ListOfferings(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService PartialUpdateInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
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
					Description: "Localhost for logme_DefaultApi",
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
		instanceId := "instanceId"
		partialUpdateInstancePayload := PartialUpdateInstancePayload{}

		reqErr := apiClient.PartialUpdateInstance(context.Background(), projectId, instanceId).PartialUpdateInstancePayload(partialUpdateInstancePayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

}
