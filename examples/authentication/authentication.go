package main

import (
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	// Create a new API client, that uses default authentication.
	// The SDK will first try to find a token in the STACKIT_SERVICE_ACCOUNT_TOKEN env var. If not present, it will
	// check the credentials file located in the path defined by the STACKIT_CREDENTIALS_PATH env var, if specified,
	// or in $HOME/.stackit/credentials.json as a fallback. If the token is found, all of the requests are authenticated using that token.
	_, err := dns.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create an unauthenticated API client
	_, err = dns.NewAPIClient(config.WithoutAuthentication())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a new API client, that will authenticate using the provided bearer token
	token := "TOKEN"
	_, err = dns.NewAPIClient(config.WithToken(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}
}
