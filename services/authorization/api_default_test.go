/*
STACKIT Membership API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authorization

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

func Test_authorization_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService AddMembers", func(t *testing.T) {
		path := "/v2/{resourceId}/members"
		resourceIdValue := "resourceId"
		path = strings.Replace(path, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(resourceIdValue, "resourceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := MembersResponse{}
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
					Description: "Localhost for authorization_DefaultApi",
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

		resourceId := "resourceId"
		addMembersPayload := AddMembersPayload{}

		resp, reqErr := apiClient.AddMembers(context.Background(), resourceId).AddMembersPayload(addMembersPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListMembers", func(t *testing.T) {
		path := "/v2/{resourceType}/{resourceId}/members"
		resourceTypeValue := "resourceType"
		path = strings.Replace(path, "{"+"resourceType"+"}", url.PathEscape(ParameterValueToString(resourceTypeValue, "resourceType")), -1)
		resourceIdValue := "resourceId"
		path = strings.Replace(path, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(resourceIdValue, "resourceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListMembersResponse{}
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
					Description: "Localhost for authorization_DefaultApi",
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

		resourceType := "resourceType"
		resourceId := "resourceId"

		resp, reqErr := apiClient.ListMembers(context.Background(), resourceType, resourceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListPermissions", func(t *testing.T) {
		path := "/v2/permissions"

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListPermissionsResponse{}
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
					Description: "Localhost for authorization_DefaultApi",
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

		resp, reqErr := apiClient.ListPermissions(context.Background()).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListRoles", func(t *testing.T) {
		path := "/v2/{resourceType}/{resourceId}/roles"
		resourceTypeValue := "resourceType"
		path = strings.Replace(path, "{"+"resourceType"+"}", url.PathEscape(ParameterValueToString(resourceTypeValue, "resourceType")), -1)
		resourceIdValue := "resourceId"
		path = strings.Replace(path, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(resourceIdValue, "resourceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := RolesResponse{}
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
					Description: "Localhost for authorization_DefaultApi",
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

		resourceType := "resourceType"
		resourceId := "resourceId"

		resp, reqErr := apiClient.ListRoles(context.Background(), resourceType, resourceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListUserMemberships", func(t *testing.T) {
		path := "/v2/users/{email}/memberships"
		emailValue := "email"
		path = strings.Replace(path, "{"+"email"+"}", url.PathEscape(ParameterValueToString(emailValue, "email")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := ListUserMembershipsResponse{}
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
					Description: "Localhost for authorization_DefaultApi",
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

		email := "email"

		resp, reqErr := apiClient.ListUserMemberships(context.Background(), email).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService RemoveMembers", func(t *testing.T) {
		path := "/v2/{resourceId}/members/remove"
		resourceIdValue := "resourceId"
		path = strings.Replace(path, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(resourceIdValue, "resourceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := MembersResponse{}
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
					Description: "Localhost for authorization_DefaultApi",
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

		resourceId := "resourceId"
		removeMembersPayload := RemoveMembersPayload{}

		resp, reqErr := apiClient.RemoveMembers(context.Background(), resourceId).RemoveMembersPayload(removeMembersPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

}
