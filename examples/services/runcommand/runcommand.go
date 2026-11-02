package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v2api"
	"github.com/stackitcloud/stackit-sdk-go/services/runcommand/v2api/wait"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	serverId := "SERVER_ID"   // the uuid of the server to run the command on
	region := "eu01"          // the region of the server

	// Create a new API client, that uses default authentication and configuration.
	client, err := runcommand.NewAPIClient()
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

	// Submit the command.
	fmt.Printf("[Run Command API] Submitting command on server %q...\n", serverId)

	var createResp *runcommand.NewCommandResponse
	for attempt := range 60 {
		createResp, err = client.DefaultAPI.CreateCommand(ctx, projectId, serverId, region).CreateCommandPayload(*payload).Execute()
		if err == nil {
			break
		}
		var oapiErr *oapierror.GenericOpenAPIError
		ok := errors.As(err, &oapiErr)
		if !ok || oapiErr.StatusCode != http.StatusNotFound {
			fmt.Fprintf(os.Stderr, "[Run Command API] Error when calling `CreateCommand`: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[Run Command API] Agent not yet ready, retrying (%d/60)...\n", attempt+1)
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Run Command API] Agent did not become ready within timeout\n")
		os.Exit(1)
	}

	commandId := strconv.Itoa(int(createResp.GetId()))
	fmt.Printf("[Run Command API] Command submitted with ID %s.\n", commandId)

	fmt.Printf("[Run Command API] Waiting for command %s to finish...\n", commandId)

	details, err := wait.RunCommandWaitHandler(ctx, client.DefaultAPI, projectId, serverId, region, commandId).
		WaitWithContext(ctx)
	if err != nil {
		exitCode := int32(0)
		output := ""
		if details != nil {
			exitCode = details.GetExitCode()
			output = details.GetOutput()
		}
		fmt.Fprintf(os.Stderr, "[Run Command API] Command %s failed (exit code: %d).\nOutput:\n%s\nError: %v\n",
			commandId, exitCode, output, err)
		os.Exit(1)
	}

	fmt.Printf("[Run Command API] Command %s completed successfully.\n", commandId)
	fmt.Printf("[Run Command API] Output:\n%s\n", details.GetOutput())
}
