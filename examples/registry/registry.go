package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	registry "github.com/stackitcloud/stackit-sdk-go/services/registry/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/registry/v1api/wait"
)

func main() {
	projectId := "<project-uuid>" // the UUID of your STACKIT project
	region := "eu01"              // STACKIT region, e.g. "eu01"

	// Create a new API client using default authentication and configuration
	client, err := registry.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Alternatively, you can create a new API client using your own authentication and configuration
	// Create a new API client, that will authenticate using the key flow
	// If you created a service account key and provided your own RSA key pair,
	// you need to add the path to a PEM encoded file including the private key
	// using config.WithPrivateKeyPath("path/to/private_key.pem")
	// saKeyPath := "/path/to/service_account_key.json"
	// 	client, err := registry.NewAPIClient(
	// 	config.WithServiceAccountKeyPath(saKeyPath), config.WithEndpoint("https://my-alternative-enpoint.com"),
	// )
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "[Registry] Creating API client: %v\n", err)
	// 	os.Exit(1)
	// }

	ctx := context.Background()

	// 1. List existing artifactories (registries) in the project
	listResp, err := client.DefaultAPI.ListArtifactories(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing artifactories: %v\n", err)
	} else {
		fmt.Printf("Found %d artifactory(ies)\n", len(listResp.Artifactories))
		for _, artifactory := range listResp.Artifactories {
			fmt.Printf("- %s (ID: %s, State: %s)\n", artifactory.Name, artifactory.Id, artifactory.State)
		}
	}

	// 2. Create a new artifactory
	artifactoryName := "my-registry"
	createPayload := registry.CreateArtifactoryPayload{
		Name: artifactoryName,
	}
	createResp, err := client.DefaultAPI.CreateArtifactory(ctx, projectId, region).
		CreateArtifactoryPayload(createPayload).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating artifactory: %v\n", err)
		return
	}

	artifactoryId := createResp.Id
	fmt.Printf("Created artifactory with ID %q. Waiting for it to become ready...\n", artifactoryId)

	// 3. Wait for the artifactory to be ready
	_, err = wait.CreateArtifactoryWaitHandler(ctx, client.DefaultAPI, projectId, region, artifactoryId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for artifactory to be ready: %v\n", err)
		return
	}
	fmt.Println("Artifactory is ready.")

	// 4. Retrieve artifactory details
	artifactory, err := client.DefaultAPI.GetArtifactory(ctx, projectId, region, artifactoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting artifactory details: %v\n", err)
		return
	}
	fmt.Printf("Artifactory details: Name=%s, State=%s, Url=%s\n", artifactory.Name, artifactory.State, artifactory.Url)

	// 5. Delete the artifactory
	err = client.DefaultAPI.DeleteArtifactory(ctx, projectId, region, artifactoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting artifactory: %v\n", err)
		return
	}
	fmt.Println("Artifactory deleted successfully.")

	// ==========================================
	// Proxy Artifactory Example
	// ==========================================

	// 6. Create a proxy artifactory (e.g. proxying Docker Hub)
	proxyArtifactoryName := "my-proxy-registry2"
	createProxyPayload := registry.CreateArtifactoryPayload{
		Name: proxyArtifactoryName,
		Proxy: &registry.CreateProxyRequest{
			RegistryType: "docker-hub",
			Url:          "https://hub.docker.com",
			Description:  utils.Ptr("Docker Hub Proxy Cache"),
			Credentials: &registry.Credentials{
				AccessKey:    utils.Ptr("my-proxy_access_key"),
				AccessSecret: utils.Ptr("my-proxy_access_secret"),
			},
		},
	}
	createProxyResp, err := client.DefaultAPI.CreateArtifactory(ctx, projectId, region).
		CreateArtifactoryPayload(createProxyPayload).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating proxy artifactory: %v\n", err)
		return
	}

	proxyArtifactoryId := createProxyResp.Id
	fmt.Printf("Created proxy artifactory with ID %q. Waiting for it to become ready...\n", proxyArtifactoryId)

	// Wait for the proxy artifactory to be ready

	_, err = wait.CreateArtifactoryWaitHandler(ctx, client.DefaultAPI, projectId, region, proxyArtifactoryId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for proxy artifactory to be ready: %v\n", err)
		return
	}
	fmt.Println("Proxy artifactory is ready.")

	// 7. Update proxy credentials
	patchPayload := registry.PatchArtifactoryPayload{
		Proxy: &registry.UpdateProxyRequest{
			Credentials: &registry.Credentials{
				AccessKey:    utils.Ptr("my-proxy_access_key_updated"),
				AccessSecret: utils.Ptr("my-proxy_access_secret_updated"),
			},
			Description: utils.Ptr("Updated Docker Hub Proxy Cache"),
		},
	}
	_, err = client.DefaultAPI.PatchArtifactory(ctx, projectId, region, proxyArtifactoryId).
		PatchArtifactoryPayload(patchPayload).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating proxy credentials: %v\n", err)
		return
	}
	fmt.Printf("Updating proxy credentials for artifactory %q. Waiting for update to complete...\n", proxyArtifactoryId)

	// Wait for the update to complete
	_, err = wait.UpdateArtifactoryWaitHandler(ctx, client.DefaultAPI, projectId, region, proxyArtifactoryId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for proxy artifactory update: %v\n", err)
		return
	}
	fmt.Println("Proxy artifactory credentials updated successfully.")

	// 8. Delete the proxy artifactory
	err = client.DefaultAPI.DeleteArtifactory(ctx, projectId, region, proxyArtifactoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting proxy artifactory: %v\n", err)
		return
	}
	fmt.Println("Proxy artifactory deleted successfully.")
}
