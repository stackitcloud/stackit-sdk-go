/*
STACKIT Object Storage API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package objectstorage

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

func Test_objectstorage_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateAccessKey", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/access-key"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CreateAccessKeyResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		createAccessKeyPayload := CreateAccessKeyPayload{}

		resp, reqErr := apiClient.CreateAccessKey(context.Background(), projectId, region).CreateAccessKeyPayload(createAccessKeyPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateBucket", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/bucket/{bucketName}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		bucketNameValue := "bucketName"
		path = strings.Replace(path, "{"+"bucketName"+"}", url.PathEscape(ParameterValueToString(bucketNameValue, "bucketName")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CreateBucketResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		bucketName := "bucketName"

		resp, reqErr := apiClient.CreateBucket(context.Background(), projectId, region, bucketName).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateCredentialsGroup", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/credentials-group"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := CreateCredentialsGroupResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		createCredentialsGroupPayload := CreateCredentialsGroupPayload{}

		resp, reqErr := apiClient.CreateCredentialsGroup(context.Background(), projectId, region).CreateCredentialsGroupPayload(createCredentialsGroupPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteAccessKey", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/access-key/{keyId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		keyIdValue := "keyId"
		path = strings.Replace(path, "{"+"keyId"+"}", url.PathEscape(ParameterValueToString(keyIdValue, "keyId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := DeleteAccessKeyResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		keyId := "keyId"

		resp, reqErr := apiClient.DeleteAccessKey(context.Background(), projectId, region, keyId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteBucket", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/bucket/{bucketName}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		bucketNameValue := "bucketName"
		path = strings.Replace(path, "{"+"bucketName"+"}", url.PathEscape(ParameterValueToString(bucketNameValue, "bucketName")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := DeleteBucketResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		bucketName := "bucketName"

		resp, reqErr := apiClient.DeleteBucket(context.Background(), projectId, region, bucketName).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteCredentialsGroup", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/credentials-group/{groupId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		groupIdValue := "groupId"
		path = strings.Replace(path, "{"+"groupId"+"}", url.PathEscape(ParameterValueToString(groupIdValue, "groupId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := DeleteCredentialsGroupResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		groupId := "groupId"

		resp, reqErr := apiClient.DeleteCredentialsGroup(context.Background(), projectId, region, groupId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DisableService", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ProjectStatus{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"

		resp, reqErr := apiClient.DisableService(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService EnableService", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ProjectStatus{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"

		resp, reqErr := apiClient.EnableService(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetBucket", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/bucket/{bucketName}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		bucketNameValue := "bucketName"
		path = strings.Replace(path, "{"+"bucketName"+"}", url.PathEscape(ParameterValueToString(bucketNameValue, "bucketName")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetBucketResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"
		bucketName := "bucketName"

		resp, reqErr := apiClient.GetBucket(context.Background(), projectId, region, bucketName).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetServiceStatus", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ProjectStatus{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"

		resp, reqErr := apiClient.GetServiceStatus(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListAccessKeys", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/access-keys"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListAccessKeysResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"

		resp, reqErr := apiClient.ListAccessKeys(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListBuckets", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/buckets"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListBucketsResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"

		resp, reqErr := apiClient.ListBuckets(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListCredentialsGroups", func(t *testing.T) {
		path := "/v2/project/{projectId}/regions/{region}/credentials-groups"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region"
		path = strings.Replace(path, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListCredentialsGroupsResponse{}
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
					Description: "Localhost for objectstorage_DefaultApi",
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
		region := "region"

		resp, reqErr := apiClient.ListCredentialsGroups(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

}
