/*
STACKIT Server Backup Management API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package serverbackup

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

func Test_serverbackup_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateBackup", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := BackupJob{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		resp, reqErr := apiClient.CreateBackup(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateBackupSchedule", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backup-schedules"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := BackupSchedule{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		resp, reqErr := apiClient.CreateBackupSchedule(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteBackup", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups/{backupId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupIdValue := "backupId"
		path = strings.Replace(path, "{"+"backupId"+"}", url.PathEscape(ParameterValueToString(backupIdValue, "backupId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupId := "backupId"

		reqErr := apiClient.DeleteBackup(context.Background(), projectId, serverId, backupId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService DeleteBackupSchedule", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backup-schedules/{backupScheduleId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupScheduleIdValue := "backupScheduleId"
		path = strings.Replace(path, "{"+"backupScheduleId"+"}", url.PathEscape(ParameterValueToString(backupScheduleIdValue, "backupScheduleId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupScheduleId := "backupScheduleId"

		reqErr := apiClient.DeleteBackupSchedule(context.Background(), projectId, serverId, backupScheduleId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService DeleteVolumeBackup", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups/{backupId}/volume-backups/{volumeBackupId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupIdValue := "backupId"
		path = strings.Replace(path, "{"+"backupId"+"}", url.PathEscape(ParameterValueToString(backupIdValue, "backupId")), -1)
		volumeBackupIdValue := "volumeBackupId"
		path = strings.Replace(path, "{"+"volumeBackupId"+"}", url.PathEscape(ParameterValueToString(volumeBackupIdValue, "volumeBackupId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupId := "backupId"
		volumeBackupId := "volumeBackupId"

		reqErr := apiClient.DeleteVolumeBackup(context.Background(), projectId, serverId, backupId, volumeBackupId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService DisableService", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		reqErr := apiClient.DisableService(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService DisableServiceResource", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/service"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		reqErr := apiClient.DisableServiceResource(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService EnableService", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		reqErr := apiClient.EnableService(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService EnableServiceResource", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/service"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		reqErr := apiClient.EnableServiceResource(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService GetBackup", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups/{backupId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupIdValue := "backupId"
		path = strings.Replace(path, "{"+"backupId"+"}", url.PathEscape(ParameterValueToString(backupIdValue, "backupId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Backup{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupId := "backupId"

		resp, reqErr := apiClient.GetBackup(context.Background(), projectId, serverId, backupId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetBackupSchedule", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backup-schedules/{backupScheduleId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupScheduleIdValue := "backupScheduleId"
		path = strings.Replace(path, "{"+"backupScheduleId"+"}", url.PathEscape(ParameterValueToString(backupScheduleIdValue, "backupScheduleId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := BackupSchedule{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupScheduleId := "backupScheduleId"

		resp, reqErr := apiClient.GetBackupSchedule(context.Background(), projectId, serverId, backupScheduleId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetServiceResource", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/service"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetBackupServiceResponse{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		resp, reqErr := apiClient.GetServiceResource(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListBackupPolicies", func(t *testing.T) {
		path := "/v1/projects/{projectId}/backup-policies"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetBackupPoliciesResponse{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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

		resp, reqErr := apiClient.ListBackupPolicies(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListBackupSchedules", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backup-schedules"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetBackupSchedulesResponse{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		resp, reqErr := apiClient.ListBackupSchedules(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListBackups", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := GetBackupsListResponse{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"

		resp, reqErr := apiClient.ListBackups(context.Background(), projectId, serverId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService RestoreBackup", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups/{backupId}/restore"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupIdValue := "backupId"
		path = strings.Replace(path, "{"+"backupId"+"}", url.PathEscape(ParameterValueToString(backupIdValue, "backupId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupId := "backupId"

		reqErr := apiClient.RestoreBackup(context.Background(), projectId, serverId, backupId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService RestoreVolumeBackup", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backups/{backupId}/volume-backups/{volumeBackupId}/restore"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupIdValue := "backupId"
		path = strings.Replace(path, "{"+"backupId"+"}", url.PathEscape(ParameterValueToString(backupIdValue, "backupId")), -1)
		volumeBackupIdValue := "volumeBackupId"
		path = strings.Replace(path, "{"+"volumeBackupId"+"}", url.PathEscape(ParameterValueToString(volumeBackupIdValue, "volumeBackupId")), -1)

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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupId := "backupId"
		volumeBackupId := "volumeBackupId"

		reqErr := apiClient.RestoreVolumeBackup(context.Background(), projectId, serverId, backupId, volumeBackupId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
	})

	t.Run("Test DefaultApiService UpdateBackupSchedule", func(t *testing.T) {
		path := "/v1/projects/{projectId}/servers/{serverId}/backup-schedules/{backupScheduleId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		serverIdValue := "serverId"
		path = strings.Replace(path, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(serverIdValue, "serverId")), -1)
		backupScheduleIdValue := "backupScheduleId"
		path = strings.Replace(path, "{"+"backupScheduleId"+"}", url.PathEscape(ParameterValueToString(backupScheduleIdValue, "backupScheduleId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := BackupSchedule{}
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
					Description: "Localhost for serverbackup_DefaultApi",
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
		serverId := "serverId"
		backupScheduleId := "backupScheduleId"

		resp, reqErr := apiClient.UpdateBackupSchedule(context.Background(), projectId, serverId, backupScheduleId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

}
