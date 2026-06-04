package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	workflows "github.com/stackitcloud/stackit-sdk-go/services/workflows/v1alphaapi"
	"github.com/stackitcloud/stackit-sdk-go/services/workflows/v1alphaapi/wait"
)

func main() {
	// Region is a path parameter on every workflows endpoint, not a client setting.
	region := "eu01"

	// Your STACKIT project ID
	projectId := "PROJECT_ID"

	// A version string from `GetProviderOptions` below. Airflow 3 is recommended:
	// it allows describing DAG sources via the dagBundles field. Airflow 2 instances
	// must supply a dagsRepository instead. See [provider options] in the API docs.
	version := "workflows-3.0-airflow-3.1"

	// Git DAG bundle config — at least one DAG bundle is required so the Airflow
	// cluster has something to schedule. Use NoAuth instead of BasicAuth for public repos.
	dagsGitUrl := "https://example.com/my-org/my-dags-repo.git"
	dagsGitBranch := "main"
	dagsGitUser := "git-user"
	dagsGitToken := "git-personal-access-token"

	// OAuth2 IdP — the STACKIT IdP variant is not yet supported by the API, so
	// every Workflows instance currently needs a custom OAuth2 IdP configured.
	idpName := "azure"
	idpClientId := "00000000-0000-0000-0000-000000000000"
	idpClientSecret := "client-secret"
	idpScope := "openid email"
	idpDiscoveryUrl := "https://login.microsoftonline.com/00000000-0000-0000-0000-000000000000/.well-known/openid-configuration"

	ctx := context.Background()

	workflowsClient, err := workflows.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Listing available provider options (versions, machine types, ...)
	optionsResp, err := workflowsClient.DefaultAPI.GetProviderOptions(ctx, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when listing provider options: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[Workflows] Available versions:")
	for _, v := range optionsResp.Versions {
		fmt.Printf("  %s (%s)\n", v.Version, v.State)
	}

	// Build the OAuth2 identity provider block
	idp := workflows.NewOAuth2IdentityProvider(
		idpClientId,
		idpClientSecret,
		idpDiscoveryUrl,
		idpName,
		idpScope,
		workflows.OAUTH2IDENTITYPROVIDERTYPE_OAUTH2,
	)

	// Build a Git-backed DAG bundle with HTTP basic auth
	gitAuth := workflows.NewBasicAuth()
	gitAuth.Type = utils.Ptr("basic")
	gitAuth.Username = utils.Ptr(dagsGitUser)
	gitAuth.Password = utils.Ptr(dagsGitToken)

	bundle := workflows.NewGitDagBundle(
		workflows.BasicAuthAsGitAuth(gitAuth),
		dagsGitBranch,
		"dags",
		workflows.GITDAGBUNDLETYPE_GIT,
		dagsGitUrl,
	)

	// Creating a Workflows instance
	createPayload := workflows.CreateInstancePayload{
		DisplayName:      "myExampleWorkflowsInstance",
		Description:      utils.Ptr("This is a Workflows instance."),
		Version:          version,
		IdentityProvider: utils.Ptr(workflows.OAuth2IdentityProviderAsIdentityProvider(idp)),
		DagBundles:       []workflows.DagBundle{workflows.GitDagBundleAsDagBundle(bundle)},
	}
	createResp, err := workflowsClient.DefaultAPI.CreateInstance(ctx, projectId, region).CreateInstancePayload(createPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when creating instance: %v\n", err)
		os.Exit(1)
	}
	instanceId := createResp.Id
	fmt.Printf("[Workflows] Triggered creation of instance with ID: %s. Waiting for it to become active...\n", instanceId)

	instance, err := wait.CreateInstanceWaitHandler(ctx, workflowsClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Workflows] Instance %s is active. Airflow UI: %s\n", instanceId, instance.Endpoints.Url)

	// Listing instances in the project
	listResp, err := workflowsClient.DefaultAPI.ListInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when listing instances: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[Workflows] Instances:")
	for _, item := range listResp.Instances {
		fmt.Printf("  %s - %s (%s)\n", item.Id, item.DisplayName, item.Status)
	}

	// Updating the instance description
	_, err = workflowsClient.DefaultAPI.UpdateInstance(ctx, projectId, region, instanceId).UpdateInstancePayload(workflows.UpdateInstancePayload{
		Description: utils.Ptr("Updated description."),
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when updating instance: %v\n", err)
		os.Exit(1)
	}
	_, err = wait.UpdateInstanceWaitHandler(ctx, workflowsClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when waiting for update: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Workflows] Instance %s updated.\n", instanceId)

	// Deleting the instance
	err = workflowsClient.DefaultAPI.DeleteInstance(ctx, projectId, region, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when deleting instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Workflows] Triggered deletion of instance %s. Waiting...\n", instanceId)

	_, err = wait.DeleteInstanceWaitHandler(ctx, workflowsClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Workflows] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Workflows] Instance %s deleted.\n", instanceId)
}
