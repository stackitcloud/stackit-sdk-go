package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceaccount"
)

func main() {
	// Specify the project ID
	projectId := "16f49d71-37ad-4137-8b97-44d9c55c4094"
	token := "eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vYXV0aC4wMS5pZHAuZXUwMS5zdGFja2l0LmNsb3VkL3Rva2VuX2tleXMiLCJraWQiOiJrZXktMSIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiODEzYjc3MGU5NGU0YjhhODEzMzhkYjlkMzEzNzAwZCIsIm5vbmNlIjoiY1c5R1VVUnVVMWhYVG5WWE1UVklkR3N6Y0ROU2RuWkxNRWxDVUU1dVJGVkVVVTFQTUdKcWQzRnFMa2hIIiwic3ViIjoiY2E1YjcwZDAtYjlkZi00NWZmLWJkNGYtOGZjOWNhZjhhOWI1Iiwic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJjbGllbnRfaWQiOiJzdGFja2l0LXBvcnRhbC1sb2dpbi1xYS1jbGllbnQtaWQiLCJjaWQiOiJzdGFja2l0LXBvcnRhbC1sb2dpbi1xYS1jbGllbnQtaWQiLCJhenAiOiJzdGFja2l0LXBvcnRhbC1sb2dpbi1xYS1jbGllbnQtaWQiLCJncmFudF90eXBlIjoiaW1wbGljaXQiLCJ1c2VyX2lkIjoiY2E1YjcwZDAtYjlkZi00NWZmLWJkNGYtOGZjOWNhZjhhOWI1Iiwib3JpZ2luIjoic2Nod2Fyei1mZWRlcmF0aW9uIiwidXNlcl9uYW1lIjoiaGVucmlxdWUuc2FudG9zQGZyZWloZWl0LmNvbSIsImVtYWlsIjoiaGVucmlxdWUuc2FudG9zQGZyZWloZWl0LmNvbSIsImF1dGhfdGltZSI6MTY5ODE2MDg4NSwicmV2X3NpZyI6IjJkMjliZjBkIiwiaWF0IjoxNjk4MTYwODg1LCJleHAiOjE2OTgxNjg0ODUsImlzcyI6Imh0dHBzOi8vYXV0aC4wMS5pZHAuZXUwMS5zdGFja2l0LmNsb3VkL29hdXRoL3Rva2VuIiwiemlkIjoidWFhIiwiYXVkIjpbInN0YWNraXQtcG9ydGFsLWxvZ2luLXFhLWNsaWVudC1pZCIsIm9wZW5pZCJdfQ.BFCAAwpomvO4SjsIC3RAdZjtTkiu93zSLzsaSM8JRMqqxKE4jx_qjAlxwQhT58grBvn5CDCkIlGjgNwlnfMzNp9r_ls0VC7SRiptfovLzjft9-RfH4ekcUtfI7DCPndNzANCbZbc7Jpc807ciO9BDdpml5ip1Oz707yJmQGGNOzDso83ivp-LJ6NQjCkDiui4H65qzdeUn7CLkzIg19kBFNgCTn7BFWy1kpdJcLwcqbSuoCjhOaA1vv2HPKGaHH2zrEdd23V_0rxTAhPmlGxLEilMSDPQoUHy4LSw_UvCj-GsFxt7cY9wFkR8aXHGqme1qDFKbQ25hqyS3wUWOrIPg"

	// Create a new API client, that uses default authentication and configuration
	client, err := serviceaccount.NewAPIClient(
		config.WithToken(token),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the service accounts for your project
	getAccountsResp, err := client.GetServiceAccounts(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetServiceAccounts`: %v\n", err)
	} else {
		fmt.Printf("Number of service accounts: %v\n", len(*getAccountsResp.Items))
	}

	// Create a service account
	createAccountPayload := serviceaccount.CreateServiceAccountPayload{
		Name: utils.Ptr("my-service-account"),
	}
	createAccountResp, err := client.CreateServiceAccount(context.Background(), projectId).CreateServiceAccountPayload(createAccountPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created service account with email \"%s\".\n", *createAccountResp.Email)
	}
}
