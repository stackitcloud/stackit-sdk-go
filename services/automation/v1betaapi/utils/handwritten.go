package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	automation "github.com/stackitcloud/stackit-sdk-go/services/automation/v1betaapi"
)

func main() {
	ctx := context.Background()

	projectID := os.Getenv("TF_ACC_PROJECT_ID") // the uuid of your STACKIT project
	region := "eu01"
	serviceAccKeyPath := os.Getenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH") // path to your service account key


	client, err := automation.NewAPIClient(
		config.WithServiceAccountKeyPath(serviceAccKeyPath),
	)
	if err != nil {
		log.Fatalf("Creating API client: %v\n", err)
	}

	client.DefaultAPI.ListVolumeAutomations(ctx, projectID, region).Execute()

	var _ PaginatedListResponse[automation.ListAutomationsItem] = &automation.ListAutomationsResponse{}
	
	var _ ListRequestWithPagination[automation.ApiListVolumeAutomationsRequest, *automation.ListAutomationsResponse] = client.DefaultAPI.ListVolumeAutomations(ctx, "", "")

	//PaginationHandler[automation.ApiListVolumeAutomationsRequest, automation.ListAutomationsResponse, automation.ListAutomationsItem](client.DefaultAPI.ListVolumeAutomations(ctx, projectID, region))

	items, err := PaginationHandler(client.DefaultAPI.ListVolumeAutomations(ctx, projectID, region))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(items))
	
	items2, err := PaginationHandler(client.DefaultAPI.ListVolumeExecutions(ctx, projectID, region, ""))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(items2))
}

func PaginationHandler[T ListRequestWithPagination[T, RESP], RESP PaginatedListResponse[ITEM], ITEM any](listReq T) ([]ITEM, error) {
	var combinedResponses []ITEM
	
	nextPageToken := ""

	for {
		// from the second API call on we have a page token from the previous call
		if nextPageToken != "" {
			listReq = listReq.PageToken(nextPageToken)
		} 

		apiResp, err := listReq.Execute()
		if err != nil {
			return nil, err
		}

		items, _ := apiResp.GetItemsOk()
		if items != nil {
			combinedResponses = append(combinedResponses, items...)
		}

		// if there is no further page: stop
		apiRespToken, _ := apiResp.GetNextPageTokenOk(); 
		if apiRespToken == nil {
			break
		}

		nextPageToken = *apiRespToken
	}

	return combinedResponses, nil
}

type PaginatedListResponse[ITEM any] interface {
	GetItemsOk() ([]ITEM, bool)
	GetNextPageTokenOk() (*string, bool)
}

type ListRequestWithPagination[REQ, RESP any] interface {
	PageSize(pageSize int32) REQ
	PageToken(pageToken string) REQ
	Execute() (RESP, error)
}
