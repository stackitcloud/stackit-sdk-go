package main

import (
	"context"
	"log"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	telemetrylink "github.com/stackitcloud/stackit-sdk-go/services/telemetrylink/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/telemetrylink/v1api/wait"
)

func main() {
	ctx := context.Background()

	organizationId := "ORGANIZATION_ID" // the uuid of your STACKIT organization
	folderId := "FOLDER_ID"             // the uuid of your STACKIT folder
	projectId := "PROJECT_ID"           // the uuid of your STACKIT project
	regionId := "eu01"
	telemetryRouterId := "TELEMETRY_ROUTER_ID"                    // the uuid of your STACKIT TelemetryRouter
	telemetryRouterAccessToken := "TELEMETRY_ROUTER_ACCESS_TOKEN" // the access token of your TelemetryRouter

	client, err := telemetrylink.NewAPIClient()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Creating API client: %v\n", err)
	}

	// --- Organization TelemetryLink Examples ---

	// Create an organization TelemetryLink
	log.Printf("[TelemetryLink API] Creating Organization TelemetryLink.\n")
	var createdOrgLink string
	createOrgLinkPayload := telemetrylink.CreateOrUpdateOrganizationTelemetryLinkPayload{
		DisplayName:       "my-org-telemetry-link",
		TelemetryRouterId: telemetryRouterId,
		AccessToken:       telemetryRouterAccessToken,
	}
	createResp, err := client.DefaultAPI.CreateOrUpdateOrganizationTelemetryLink(ctx, organizationId, regionId).
		CreateOrUpdateOrganizationTelemetryLinkPayload(createOrgLinkPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrUpdateOrganizationTelemetryLink`: %v\n", err)
	}
	createdOrgLink = createResp.Id

	// Wait for the TelemetryLink to be ready
	log.Printf("[TelemetryLink API] Waiting for organization TelemetryLink to be created.\n")
	_, err = wait.CreateOrganizationTelemetryLinkWaitHandler(ctx, client.DefaultAPI, organizationId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Organization TelemetryLink \"%s\" has been successfully created.\n", createdOrgLink)

	// Get the created organization TelemetryLink Instance
	log.Printf("[TelemetryLink API] Retrieving organization TelemetryLink.\n")
	getResp, err := client.DefaultAPI.GetOrganizationTelemetryLink(ctx, organizationId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `GetOrganizationTelemetryLink`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Retrieved organization TelemetryLink with ID \"%s\" and Display Name \"%s\".\n", getResp.Id, getResp.DisplayName)

	// Update the created organization TelemetryLink
	log.Printf("[TelemetryLink API] Updating organization TelemetryLink.\n")
	updatePayload := telemetrylink.CreateOrUpdateOrganizationTelemetryLinkPayload{
		DisplayName:       "my-upd-org-telemetry-link",
		TelemetryRouterId: telemetryRouterId,
		AccessToken:       telemetryRouterAccessToken,
	}
	updateResp, err := client.DefaultAPI.CreateOrUpdateOrganizationTelemetryLink(ctx, organizationId, regionId).
		CreateOrUpdateOrganizationTelemetryLinkPayload(updatePayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrUpdateOrganizationTelemetryLink`: %v\n", err)
	}

	// Wait for the organization TelemetryLink to be updated
	log.Printf("[TelemetryLink API] Waiting for organization TelemetryLink to be updated.\n")
	_, err = wait.UpdateOrganizationTelemetryLinkWaitHandler(ctx, client.DefaultAPI, organizationId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Organization TelemetryLink \"%s\" has been successfully updated.\n", updateResp.Id)

	// Partially update organization TelemetryLink
	log.Printf("[TelemetryLink API] Partially updating organization TelemetryLink.\n")
	partialUpdatePayload := telemetrylink.PartialUpdateOrganizationTelemetryLinkPayload{
		Description: utils.Ptr("description"),
	}
	partialUpdateResp, err := client.DefaultAPI.PartialUpdateOrganizationTelemetryLink(ctx, organizationId, regionId).
		PartialUpdateOrganizationTelemetryLinkPayload(partialUpdatePayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `PartialUpdateOrganizationTelemetryLink`: %v\n", err)
	}

	// Wait for the organization TelemetryLink to be partiallyupdated
	log.Printf("[TelemetryLink API] Waiting for organization TelemetryLink to be partially updated.\n")
	_, err = wait.PartialUpdateOrganizationTelemetryLinkWaitHandler(ctx, client.DefaultAPI, organizationId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for partial update: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Organization TelemetryLink \"%s\" has been successfully partially updated.\n", partialUpdateResp.Id)

	// Create a organization backflush job
	log.Printf("[TelemetryLink API] Creating organization backflush job.\n")
	orgCreateBackflushResp, err := client.DefaultAPI.CreateOrganizationBackflushJob(ctx, organizationId, regionId).
		CreateOrganizationBackflushJobPayload(telemetrylink.CreateOrganizationBackflushJobPayload{
			DisplayName: utils.Ptr("my-organization-backflush-job"),
		}).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrganizationBackflushJob`: %v\n", err)
	}

	// Wait for the organization backflush job to be running or complete
	log.Printf("[TelemetryLink API] Waiting for organization backflush job to be created.\n")
	_, err = wait.CreateOrganizationBackflushJobWaitHandler(ctx, client.DefaultAPI, organizationId, regionId, orgCreateBackflushResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for backflush job creation: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Organization backflush job \"%s\" has been successfully created.\n", orgCreateBackflushResp.Id)

	// List the organization backflush jobs
	orgListBackflushResp, err := client.DefaultAPI.ListOrganizationBackflushJobs(ctx, organizationId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `ListOrganizationBackflushJobs`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Found %d organization backflush job(s).\n", len(orgListBackflushResp.BackflushJobs))

	// Get the organization backflush job
	orgGetBackflushResp, err := client.DefaultAPI.GetOrganizationBackflushJob(ctx, organizationId, regionId, orgCreateBackflushResp.Id).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `GetOrganizationBackflushJob`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Organization backflush job \"%s\" has status \"%s\".\n", orgGetBackflushResp.Id, orgGetBackflushResp.Status)

	// Delete the organization TelemetryLink
	log.Printf("[TelemetryLink API] Deleting organization TelemetryLink.\n")
	err = client.DefaultAPI.DeleteOrganizationTelemetryLink(ctx, organizationId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `DeleteOrganizationTelemetryLink`: %v\n", err)
	}

	// Wait for the organization TelemetryLink to be deleted
	log.Printf("[TelemetryLink API] Waiting for organization TelemetryLink to be deleted.\n")
	_, err = wait.DeleteOrganizationTelemetryLinkWaitHandler(ctx, client.DefaultAPI, organizationId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Organization TelemetryLink \"%s\" has been successfully deleted.\n", createdOrgLink)

	// --- Folder TelemetryLink Examples ---

	// Create a folder TelemetryLink
	log.Printf("[TelemetryLink API] Creating Folder TelemetryLink.\n")
	var createdFolderLink string
	createFolderLinkPayload := telemetrylink.CreateOrUpdateFolderTelemetryLinkPayload{
		DisplayName:       "my-folder-telemetry-link",
		TelemetryRouterId: telemetryRouterId,
		AccessToken:       telemetryRouterAccessToken,
	}
	createFolderResp, err := client.DefaultAPI.CreateOrUpdateFolderTelemetryLink(ctx, folderId, regionId).
		CreateOrUpdateFolderTelemetryLinkPayload(createFolderLinkPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrUpdateFolderTelemetryLink`: %v\n", err)
	}
	createdFolderLink = createFolderResp.Id

	// Wait for the Folder TelemetryLink to be ready
	log.Printf("[TelemetryLink API] Waiting for folder TelemetryLink to be created.\n")
	_, err = wait.CreateFolderTelemetryLinkWaitHandler(ctx, client.DefaultAPI, folderId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Folder TelemetryLink \"%s\" has been successfully created.\n", createdFolderLink)

	// Get the created folder TelemetryLink Instance
	log.Printf("[TelemetryLink API] Retrieving folder TelemetryLink.\n")
	getFolderResp, err := client.DefaultAPI.GetFolderTelemetryLink(ctx, folderId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `GetFolderTelemetryLink`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Retrieved folder TelemetryLink with ID \"%s\" and Display Name \"%s\".\n", getFolderResp.Id, getFolderResp.DisplayName)

	// Update the created folder TelemetryLink
	log.Printf("[TelemetryLink API] Updating folder TelemetryLink.\n")
	updateFolderPayload := telemetrylink.CreateOrUpdateFolderTelemetryLinkPayload{
		DisplayName:       "my-upd-folder-telemetry-link",
		TelemetryRouterId: telemetryRouterId,
		AccessToken:       telemetryRouterAccessToken,
	}
	updateFolderResp, err := client.DefaultAPI.CreateOrUpdateFolderTelemetryLink(ctx, folderId, regionId).
		CreateOrUpdateFolderTelemetryLinkPayload(updateFolderPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrUpdateFolderTelemetryLink`: %v\n", err)
	}

	// Wait for the folder TelemetryLink to be updated
	log.Printf("[TelemetryLink API] Waiting for folder TelemetryLink to be updated.\n")
	_, err = wait.UpdateFolderTelemetryLinkWaitHandler(ctx, client.DefaultAPI, folderId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Folder TelemetryLink \"%s\" has been successfully updated.\n", updateFolderResp.Id)

	// Partially update folder TelemetryLink
	log.Printf("[TelemetryLink API] Partially updating folder TelemetryLink.\n")
	partialUpdateFolderPayload := telemetrylink.PartialUpdateFolderTelemetryLinkPayload{
		Description: utils.Ptr("folder description"),
	}
	partialUpdateFolderResp, err := client.DefaultAPI.PartialUpdateFolderTelemetryLink(ctx, folderId, regionId).
		PartialUpdateFolderTelemetryLinkPayload(partialUpdateFolderPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `PartialUpdateFolderTelemetryLink`: %v\n", err)
	}

	// Wait for the folder TelemetryLink to be partially updated
	log.Printf("[TelemetryLink API] Waiting for folder TelemetryLink to be partially updated.\n")
	_, err = wait.PartialUpdateFolderTelemetryLinkWaitHandler(ctx, client.DefaultAPI, folderId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for partial update: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Folder TelemetryLink \"%s\" has been successfully partially updated.\n", partialUpdateFolderResp.Id)

	// Create a folder backflush job
	log.Printf("[TelemetryLink API] Creating folder backflush job.\n")
	folderCreateBackflushResp, err := client.DefaultAPI.CreateFolderBackflushJob(ctx, folderId, regionId).
		CreateFolderBackflushJobPayload(telemetrylink.CreateFolderBackflushJobPayload{
			DisplayName: utils.Ptr("my-folder-backflush-job"),
		}).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateFolderBackflushJob`: %v\n", err)
	}

	// Wait for the folder backflush job to be running or complete
	log.Printf("[TelemetryLink API] Waiting for folder backflush job to be created.\n")
	_, err = wait.CreateFolderBackflushJobWaitHandler(ctx, client.DefaultAPI, folderId, regionId, folderCreateBackflushResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for backflush job creation: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Folder backflush job \"%s\" has been successfully created.\n", folderCreateBackflushResp.Id)

	// List the folder backflush jobs
	folderListBackflushResp, err := client.DefaultAPI.ListFolderBackflushJobs(ctx, folderId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `ListFolderBackflushJobs`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Found %d folder backflush job(s).\n", len(folderListBackflushResp.BackflushJobs))

	// Get the folder backflush job
	folderGetBackflushResp, err := client.DefaultAPI.GetFolderBackflushJob(ctx, folderId, regionId, folderCreateBackflushResp.Id).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `GetFolderBackflushJob`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Folder backflush job \"%s\" has status \"%s\".\n", folderGetBackflushResp.Id, folderGetBackflushResp.Status)

	// Delete the folder TelemetryLink
	log.Printf("[TelemetryLink API] Deleting folder TelemetryLink.\n")
	err = client.DefaultAPI.DeleteFolderTelemetryLink(ctx, folderId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `DeleteFolderTelemetryLink`: %v\n", err)
	}

	// Wait for the folder TelemetryLink to be deleted
	log.Printf("[TelemetryLink API] Waiting for folder TelemetryLink to be deleted.\n")
	_, err = wait.DeleteFolderTelemetryLinkWaitHandler(ctx, client.DefaultAPI, folderId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Folder TelemetryLink \"%s\" has been successfully deleted.\n", createdFolderLink)

	// --- Project TelemetryLink Examples ---

	// Create a project TelemetryLink
	log.Printf("[TelemetryLink API] Creating Project TelemetryLink.\n")
	var createdProjectLink string
	createProjectLinkPayload := telemetrylink.CreateOrUpdateProjectTelemetryLinkPayload{
		DisplayName:       "my-project-telemetry-link",
		TelemetryRouterId: telemetryRouterId,
		AccessToken:       telemetryRouterAccessToken,
	}
	createProjectResp, err := client.DefaultAPI.CreateOrUpdateProjectTelemetryLink(ctx, projectId, regionId).
		CreateOrUpdateProjectTelemetryLinkPayload(createProjectLinkPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrUpdateProjectTelemetryLink`: %v\n", err)
	}
	createdProjectLink = createProjectResp.Id

	// Wait for the Project TelemetryLink to be ready
	log.Printf("[TelemetryLink API] Waiting for project TelemetryLink to be created.\n")
	_, err = wait.CreateProjectTelemetryLinkWaitHandler(ctx, client.DefaultAPI, projectId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Project TelemetryLink \"%s\" has been successfully created.\n", createdProjectLink)

	// Get the created project TelemetryLink Instance
	log.Printf("[TelemetryLink API] Retrieving project TelemetryLink.\n")
	getProjectResp, err := client.DefaultAPI.GetProjectTelemetryLink(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `GetProjectTelemetryLink`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Retrieved project TelemetryLink with ID \"%s\" and Display Name \"%s\".\n", getProjectResp.Id, getProjectResp.DisplayName)

	// Update the created project TelemetryLink
	log.Printf("[TelemetryLink API] Updating project TelemetryLink.\n")
	updateProjectPayload := telemetrylink.CreateOrUpdateProjectTelemetryLinkPayload{
		DisplayName:       "my-upd-project-telemetry-link",
		TelemetryRouterId: telemetryRouterId,
		AccessToken:       telemetryRouterAccessToken,
	}
	updateProjectResp, err := client.DefaultAPI.CreateOrUpdateProjectTelemetryLink(ctx, projectId, regionId).
		CreateOrUpdateProjectTelemetryLinkPayload(updateProjectPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateOrUpdateProjectTelemetryLink`: %v\n", err)
	}

	// Wait for the project TelemetryLink to be updated
	log.Printf("[TelemetryLink API] Waiting for project TelemetryLink to be updated.\n")
	_, err = wait.UpdateProjectTelemetryLinkWaitHandler(ctx, client.DefaultAPI, projectId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Project TelemetryLink \"%s\" has been successfully updated.\n", updateProjectResp.Id)

	// Partially update project TelemetryLink
	log.Printf("[TelemetryLink API] Partially updating project TelemetryLink.\n")
	partialUpdateProjectPayload := telemetrylink.PartialUpdateProjectTelemetryLinkPayload{
		Description: utils.Ptr("project description"),
	}
	partialUpdateProjectResp, err := client.DefaultAPI.PartialUpdateProjectTelemetryLink(ctx, projectId, regionId).
		PartialUpdateProjectTelemetryLinkPayload(partialUpdateProjectPayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `PartialUpdateProjectTelemetryLink`: %v\n", err)
	}

	// Wait for the project TelemetryLink to be partially updated
	log.Printf("[TelemetryLink API] Waiting for project TelemetryLink to be partially updated.\n")
	_, err = wait.PartialUpdateProjectTelemetryLinkWaitHandler(ctx, client.DefaultAPI, projectId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for partial update: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Project TelemetryLink \"%s\" has been successfully partially updated.\n", partialUpdateProjectResp.Id)

	// Create a project backflush job
	log.Printf("[TelemetryLink API] Creating project backflush job.\n")
	projectCreateBackflushResp, err := client.DefaultAPI.CreateProjectBackflushJob(ctx, projectId, regionId).
		CreateProjectBackflushJobPayload(telemetrylink.CreateProjectBackflushJobPayload{
			DisplayName: utils.Ptr("my-project-backflush-job"),
		}).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `CreateProjectBackflushJob`: %v\n", err)
	}

	// Wait for the project backflush job to be running or complete
	log.Printf("[TelemetryLink API] Waiting for project backflush job to be created.\n")
	_, err = wait.CreateProjectBackflushJobWaitHandler(ctx, client.DefaultAPI, projectId, regionId, projectCreateBackflushResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for backflush job creation: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Project backflush job \"%s\" has been successfully created.\n", projectCreateBackflushResp.Id)

	// List the project backflush jobs
	projectListBackflushResp, err := client.DefaultAPI.ListProjectBackflushJobs(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `ListProjectBackflushJobs`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Found %d project backflush job(s).\n", len(projectListBackflushResp.BackflushJobs))

	// Get the project backflush job
	projectGetBackflushResp, err := client.DefaultAPI.GetProjectBackflushJob(ctx, projectId, regionId, projectCreateBackflushResp.Id).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `GetProjectBackflushJob`: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Project backflush job \"%s\" has status \"%s\".\n", projectGetBackflushResp.Id, projectGetBackflushResp.Status)

	// Delete the project TelemetryLink
	log.Printf("[TelemetryLink API] Deleting project TelemetryLink.\n")
	err = client.DefaultAPI.DeleteProjectTelemetryLink(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when calling `DeleteProjectTelemetryLink`: %v\n", err)
	}

	// Wait for the project TelemetryLink to be deleted
	log.Printf("[TelemetryLink API] Waiting for project TelemetryLink to be deleted.\n")
	_, err = wait.DeleteProjectTelemetryLinkWaitHandler(ctx, client.DefaultAPI, projectId, regionId).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryLink API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[TelemetryLink API] Project TelemetryLink \"%s\" has been successfully deleted.\n", createdProjectLink)
}
