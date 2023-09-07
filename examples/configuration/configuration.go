package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
	"github.com/stackitcloud/stackit-sdk-go/services/postgresql"
)

func main() {
	// Create a new API client with a custom configuration
	configuration := &config.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: config.ServerConfigurations{
			{
				URL:         "http://localhost:8080",
				Description: "Localhost for dns_ZoneApi",
			},
		},
		OperationServers: map[string]config.ServerConfigurations{},
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	_, err := dns.NewAPIClient(config.WithCustomConfiguration(configuration))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a new API client with custom configuration for retries
	_, err = dns.NewAPIClient(
		config.WithMaxRetries(5),
		config.WithWaitBetweenCalls(10*time.Second),
		config.WithRetryTimeout(1*time.Minute))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a new API client with custom endpoint, e.g. for accessing the QA environment
	_, err = dns.NewAPIClient(config.WithEndpoint("www.api-qa-url.com"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a new API client for a regional API
	_, err = postgresql.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[PostgreSQL API] Creating API client: %v\n", err)
		os.Exit(1)
	}
}
