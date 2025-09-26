package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/services/auditlog"
)

func main() {
	// Specify the project ID, startTime and endTIme
	projectId := "PROJECT_ID"
	startTime := time.Now().Add(-time.Hour * 24)
	endTime := time.Now()
	limit := float32(100) // set pagination limit to avoid rate limit

	// Create a new API client, that uses default authentication and configuration
	auditlogClient, err := auditlog.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Auditlog API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List all audit logs of a project
	listProjectLogsReq := auditlogClient.ListProjectAuditLogEntries(context.Background(), projectId).
		StartTimeRange(startTime).
		EndTimeRange(endTime).
		Limit(limit)
	result, err := listProjectLogsReq.Execute()

	var allItems []auditlog.AuditLogEntryResponse
	for {
		if err != nil {
			var oapiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiErr) {
				// reached rate limit
				if oapiErr.StatusCode == http.StatusTooManyRequests {
					fmt.Fprintf(os.Stderr, "[Auditlog API] Too Many Requests: %v\n", string(oapiErr.Body))
					break
				}
			}
			fmt.Fprintf(os.Stderr, "[Auditlog API] List project audit log entries: %v\n", err)
			os.Exit(1)
		}
		// Break loop when response has no items
		if result == nil || result.Items == nil || len(*result.Items) == 0 {
			break
		}

		// Append items to allItems
		allItems = append(allItems, *result.Items...)

		// If cursor is not set, end of logs is reached
		if result.Cursor == nil {
			fmt.Printf("[Auditlog API] Successfully all items.\n")
			break
		}

		// Paginate to the next set of items
		listProjectLogsReq = listProjectLogsReq.Cursor(*result.Cursor)
		result, err = listProjectLogsReq.Execute()
	}

	fmt.Printf("[Auditlog API] Number of project audit log entries: %v\n", len(allItems))
}
