package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/membership"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	yourEmail := "your-email@example.com"
	emailToBeAdded := "email@example.com"

	// Create a new API client, that uses default authentication and configuration
	client, err := membership.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the available permissions for the project resource type
	getPermissionsResp, err := client.GetPermissions(context.Background()).ResourceType("project").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetPermissions`: %v\n", err)
	} else {
		availablePermissions := *getPermissionsResp.Permissions
		if len(availablePermissions) > 0 {
			fmt.Printf("Example of available permission: %v\n", *availablePermissions[0].Name)
		}
	}

	// Get the memberships of your user
	getMembershipsResp, err := client.GetMemberships(context.Background(), yourEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetMemberships`: %v\n", err)
	} else {
		userMemberships := *getMembershipsResp.Items
		fmt.Printf("Number of memberships: %v\n", len(userMemberships))

		if len(userMemberships) > 0 {
			fmt.Printf("Example of a membership of user %s: \nResource type - %s\nResource id - %s\nRole - %s\n",
				*userMemberships[0].Subject,
				*userMemberships[0].ResourceType,
				*userMemberships[0].ResourceId,
				*userMemberships[0].Role,
			)
		}
	}

	// Get the members of your project
	getMembersResp, err := client.GetMembers(context.Background(), "project", projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetMembers`: %v\n", err)
	} else {
		fmt.Printf("Number of members: %v\n", len(*getMembersResp.Members))
	}

	// Add a member to your project or add an additional role to an existing member
	updateMemberPayload := membership.UpdateMembersPayload{
		Members: &[]membership.Member{
			{
				Role:    utils.Ptr("project.member"),
				Subject: utils.Ptr(emailToBeAdded),
			},
		},
		ResourceType: utils.Ptr("project"),
	}
	_, err = client.UpdateMembers(context.Background(), projectId).UpdateMembersPayload(updateMemberPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateMembers`: %v\n", err)
	} else {
		fmt.Printf("Added member successfully.\n")
	}

	// Remove a role from a member of your project
	deleteMemberPayload := membership.DeleteMembersPayload{
		Members: &[]membership.Member{
			{
				Role:    utils.Ptr("project.member"),
				Subject: utils.Ptr(emailToBeAdded),
			},
		},
		ResourceType: utils.Ptr("project"),
	}
	_, err = client.DeleteMembers(context.Background(), projectId).DeleteMembersPayload(deleteMemberPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteMembers`: %v\n", err)
	} else {
		fmt.Printf("Removed member successfully.\n")
	}
}
