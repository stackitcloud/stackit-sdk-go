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
	projectId := "PROJECT_ID"                    // the uuid of your STACKIT project
	startTime := time.Now().Add(-time.Hour * 24) // your start time
	endTime := time.Now()                        // your end time
	limit := float32(100)                        // set pagination limit to avoid rate limit

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

	// To fetch all audit log items within a specified time range, we must implement pagination, because the api returns only
	// up to 100 elements per request. We store the result of each request in `allItems`. The response includes a cursor,
	// if more elements are available. This cursor must be used to get the next set of elements.
	// The api has a rate limit, which can be reached when all elements will be fetched with pagination or if you do multiple request.
	// The rate limit should be taken care in this case.
	var allItems []auditlog.AuditLogEntryResponse
	for {
		if err != nil {
			var oapiErr *oapierror.GenericOpenAPIError
			if errors.As(err, &oapiErr) {
				// Check if rate limit is reached
				if oapiErr.StatusCode == http.StatusTooManyRequests {
					// In case you want to fetch all items, you may want to wait some time and retry the request.
					// In this example we just stop here the pagination.
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
			fmt.Printf("[Auditlog API] Successfully fetched all items.\n")
			break
		}

		// Paginate to the next set of items
		listProjectLogsReq = listProjectLogsReq.Cursor(*result.Cursor)
		result, err = listProjectLogsReq.Execute()
	}

	fmt.Printf("[Auditlog API] Number of project audit log entries: %v\n", len(allItems))
}
