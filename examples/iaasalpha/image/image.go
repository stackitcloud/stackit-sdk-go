package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha/wait"
)

func main() {
	// Specify the and project ID
	projectId := "PROJECT_ID"
	imageFilePath := "PATH/TO/IMAGE" // Should be a path to a valid image file, e.g. "./my-image.qcow2"
	imageDiskFormat := "DISK_FORMAT" // E.g. "qcow2", "raw", "iso"

	// Create a new API client, that uses default authentication and configuration
	iaasalphaClient, err := iaasalpha.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Creating API client: %v\n", err)
		os.Exit(1)
	}
	ctx := context.Background()

	// Create an image
	createImagePayload := iaasalpha.CreateImagePayload{
		Name:       utils.Ptr("my-image"),
		DiskFormat: utils.Ptr(imageDiskFormat),
	}
	imageCreateResp, err := iaasalphaClient.CreateImage(ctx, projectId).CreateImagePayload(createImagePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `CreateImage`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[iaasalpha API] Image %q has been successfully created.\n", *imageCreateResp.Id)

	// Upload the image by making a PUT request to upload URL
	fileContents, err := os.ReadFile(imageFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when reading file: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodPut, *imageCreateResp.UploadUrl, bytes.NewReader(fileContents))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when creating request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	fmt.Printf("[iaasalpha API] Uploading image %q...\n", *imageCreateResp.Id)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when making request: %v\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when uploading image: %v\n", resp.Status)
		_ = resp.Body.Close()
		os.Exit(1)
	}
	_ = resp.Body.Close()
	fmt.Printf("[iaasalpha API] Image %q has been successfully uploaded.\n", *imageCreateResp.Id)

	// Wait for image to become available
	image, err := wait.ImageUploadWaitHandler(ctx, iaasalphaClient, projectId, *imageCreateResp.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Image %q is available.\n", *image.Id)

	// Delete the image
	err = iaasalphaClient.DeleteImage(ctx, projectId, *imageCreateResp.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `DeleteImage`: %v\n", err)
	}
	fmt.Printf("[iaasalpha API] Triggered deletion of image with ID %q.\n", *image.Id)

	// Wait for image to be deleted
	_, err = wait.DeleteImageWaitHandler(ctx, iaasalphaClient, projectId, *imageCreateResp.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Image %q has been successfully deleted.\n", *imageCreateResp.Id)
}
