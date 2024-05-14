/*
IaaS-API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package iaas

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

func Test_iaas_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateNetworkAreaRange", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/network-ranges"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := NetworkRangeList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		createNetworkAreaRangePayload := CreateNetworkAreaRangePayload{}

		resp, reqErr := apiClient.CreateNetworkAreaRange(context.Background(), organizationId, areaId).CreateNetworkAreaRangePayload(createNetworkAreaRangePayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateNetworkAreaRoute", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/routes"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := RouteList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		createNetworkAreaRoutePayload := CreateNetworkAreaRoutePayload{}

		resp, reqErr := apiClient.CreateNetworkAreaRoute(context.Background(), organizationId, areaId).CreateNetworkAreaRoutePayload(createNetworkAreaRoutePayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateNetworkAreas", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := V1NetworkArea{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		createNetworkAreasPayload := CreateNetworkAreasPayload{}

		resp, reqErr := apiClient.CreateNetworkAreas(context.Background(), organizationId).CreateNetworkAreasPayload(createNetworkAreasPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteNetworkAreaRange", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/network-ranges/{networkRangeId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)
		networkRangeIdValue := "networkRangeId"
		path = strings.Replace(path, "{"+"networkRangeId"+"}", url.PathEscape(ParameterValueToString(networkRangeIdValue, "networkRangeId")), -1)

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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		networkRangeId := "networkRangeId"

		reqErr := apiClient.DeleteNetworkAreaRange(context.Background(), organizationId, areaId, networkRangeId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService DeleteNetworkAreaRoute", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/routes/{routeId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)
		routeIdValue := "routeId"
		path = strings.Replace(path, "{"+"routeId"+"}", url.PathEscape(ParameterValueToString(routeIdValue, "routeId")), -1)

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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		routeId := "routeId"

		reqErr := apiClient.DeleteNetworkAreaRoute(context.Background(), organizationId, areaId, routeId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService DeleteNetworkAreas", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"

		reqErr := apiClient.DeleteNetworkAreas(context.Background(), organizationId, areaId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService GetNetwork", func(t *testing.T) {
		path := "/v1beta1/projects/{projectId}/networks/{networkId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		networkIdValue := "networkId"
		path = strings.Replace(path, "{"+"networkId"+"}", url.PathEscape(ParameterValueToString(networkIdValue, "networkId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Network{}
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
					Description: "Localhost for iaas_DefaultApi",
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
		networkId := "networkId"

		resp, reqErr := apiClient.GetNetwork(context.Background(), projectId, networkId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetNetworkArea", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := V1NetworkArea{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"

		resp, reqErr := apiClient.GetNetworkArea(context.Background(), organizationId, areaId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetNetworkAreaRange", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/network-ranges/{networkRangeId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)
		networkRangeIdValue := "networkRangeId"
		path = strings.Replace(path, "{"+"networkRangeId"+"}", url.PathEscape(ParameterValueToString(networkRangeIdValue, "networkRangeId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := NetworkRange{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		networkRangeId := "networkRangeId"

		resp, reqErr := apiClient.GetNetworkAreaRange(context.Background(), organizationId, areaId, networkRangeId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetNetworkAreaRoute", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/routes/{routeId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)
		routeIdValue := "routeId"
		path = strings.Replace(path, "{"+"routeId"+"}", url.PathEscape(ParameterValueToString(routeIdValue, "routeId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Route{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		routeId := "routeId"

		resp, reqErr := apiClient.GetNetworkAreaRoute(context.Background(), organizationId, areaId, routeId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetOrganizationRequest", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/requests/{requestId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		requestIdValue := "requestId"
		path = strings.Replace(path, "{"+"requestId"+"}", url.PathEscape(ParameterValueToString(requestIdValue, "requestId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Request{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		requestId := "requestId"

		resp, reqErr := apiClient.GetOrganizationRequest(context.Background(), organizationId, requestId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetProjectRequest", func(t *testing.T) {
		path := "/v1beta1/projects/{projectId}/requests/{requestId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		requestIdValue := "requestId"
		path = strings.Replace(path, "{"+"requestId"+"}", url.PathEscape(ParameterValueToString(requestIdValue, "requestId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Request{}
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
					Description: "Localhost for iaas_DefaultApi",
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
		requestId := "requestId"

		resp, reqErr := apiClient.GetProjectRequest(context.Background(), projectId, requestId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListNetworkAreaProjects", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/projects"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ProjectList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"

		resp, reqErr := apiClient.ListNetworkAreaProjects(context.Background(), organizationId, areaId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListNetworkAreaRanges", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/network-ranges"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := NetworkRangeList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"

		resp, reqErr := apiClient.ListNetworkAreaRanges(context.Background(), organizationId, areaId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListNetworkAreaRoutes", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}/routes"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := RouteList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"

		resp, reqErr := apiClient.ListNetworkAreaRoutes(context.Background(), organizationId, areaId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListNetworkAreas", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := NetworkAreaList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"

		resp, reqErr := apiClient.ListNetworkAreas(context.Background(), organizationId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListNetworks", func(t *testing.T) {
		path := "/v1beta1/projects/{projectId}/networks"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := NetworkList{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		resp, reqErr := apiClient.ListNetworks(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService UpdateNetworkArea", func(t *testing.T) {
		path := "/v1beta1/organizations/{organizationId}/network-areas/{areaId}"
		organizationIdValue := "organizationId"
		path = strings.Replace(path, "{"+"organizationId"+"}", url.PathEscape(ParameterValueToString(organizationIdValue, "organizationId")), -1)
		areaIdValue := "areaId"
		path = strings.Replace(path, "{"+"areaId"+"}", url.PathEscape(ParameterValueToString(areaIdValue, "areaId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := V1NetworkArea{}
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
					Description: "Localhost for iaas_DefaultApi",
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

		organizationId := "organizationId"
		areaId := "areaId"
		updateNetworkAreaPayload := UpdateNetworkAreaPayload{}

		resp, reqErr := apiClient.UpdateNetworkArea(context.Background(), organizationId, areaId).UpdateNetworkAreaPayload(updateNetworkAreaPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

}