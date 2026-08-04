package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/runcommand/v1api/wait"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	serverId := "SERVER_ID"   // the uuid of the server to run the command on

	// Create a new API client, that uses default authentication and configuration
	client, err := runcommand.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Run Command API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List available command templates
	templates, err := client.DefaultAPI.ListCommandTemplates(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Run Command API] Error when calling `ListCommandTemplates`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[Run Command API] Available command templates:\n")
	for _, t := range templates.GetItems() {
		fmt.Printf("  %s\n", t.GetName())
	}

	// Build the command payload
	payload := runcommand.NewCreateCommandPayload("RunShellScript")
	payload.SetParameters(map[string]string{
		"script": "echo 'Hello from STACKIT Run Commands!'",
	})

	// AgentReadyWaitHandler submits the command and retries until the server agent
	// has registered. The API returns 404 while the agent is still booting after
	// server creation. The returned response already contains the command ID.
	fmt.Printf("[Run Command API] Waiting for agent on server %q and submitting command...\n", serverId)

	createResp, err := wait.AgentReadyWaitHandler(ctx, client.DefaultAPI, projectId, serverId, *payload).
		WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Run Command API] Error when submitting command: %v\n", err)
		os.Exit(1)
	}

	commandId := strconv.Itoa(int(createResp.GetId()))
	fmt.Printf("[Run Command API] Command submitted with ID %s.\n", commandId)

	// RunCommandWaitHandler polls until the command reaches a terminal state.
	// Both COMPLETED and FAILED are terminal; inspect the status to distinguish them.
	fmt.Printf("[Run Command API] Waiting for command %s to finish...\n", commandId)

	details, err := wait.RunCommandWaitHandler(ctx, client.DefaultAPI, projectId, serverId, commandId).
		WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Run Command API] Error when waiting for command: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[Run Command API] Command %s finished with status %q (exit code: %d).\n",
		commandId, details.GetStatus(), details.GetExitCode())
	fmt.Printf("[Run Command API] Output:\n%s\n", details.GetOutput())
}
