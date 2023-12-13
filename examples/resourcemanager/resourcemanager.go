package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/resourcemanager"
)

func main() {
	// Specify the parent organization ID
	parentOrganizationId := "ORGANIZATION_ID"

	// Create a new API client, that uses default authentication and configuration
	client, err := resourcemanager.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the projects under a specific resource (organization)
	getProjectsResp, err := client.GetProjects(context.Background()).ContainerParentId(parentOrganizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetProjects`: %v\n", err)
	} else {
		fmt.Printf("Number of projects: %v\n", len(*getProjectsResp.Items))
	}

	// Create a project
	createProjectPayload := resourcemanager.CreateProjectPayload{
		ContainerParentId: utils.Ptr(parentOrganizationId),
		Name:              utils.Ptr("my-project"),
		Members: &[]resourcemanager.ProjectMember{
			{
				Role:    utils.Ptr("owner"),
				Subject: utils.Ptr("owner-email@example.com"),
			},
		},
	}
	createProjectResp, err := client.CreateProject(context.Background()).CreateProjectPayload(createProjectPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateProject`: %v\n", err)
	} else {
		fmt.Printf("Created project with id \"%s\".\n", *createProjectResp.ProjectId)
	}
}
