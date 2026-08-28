package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/stackitcloud/stackit-sdk-go/core/paginate"
	albwaf "github.com/stackitcloud/stackit-sdk-go/services/albwaf/v1api"
	automation "github.com/stackitcloud/stackit-sdk-go/services/automation/v1betaapi"
)

const region = "eu01"

func main() {
	ctx := context.Background()
	projectId := ""
	if projectId == "" {
		projectId = os.Getenv("PROJECT_ID")
	}
	if projectId == "" {
		log.Fatal("projectId is required")
	}
	aipCompliantExample(ctx, projectId)
	fmt.Print("\n\n")
	adapterExample(ctx, projectId)
}

func aipCompliantExample(ctx context.Context, projectId string) {
	client, err := automation.NewAPIClient()
	if err != nil {
		log.Fatal("create automation client: ", err)
	}
	req := client.DefaultAPI.ListVolumeTemplates(ctx, projectId, region)
	items, err := paginate.All(
		req,
		paginate.WithLimit(2),
	)
	if err != nil {
		log.Fatal("list items: ", err)
	}
	fmt.Println("listing automation templates")
	fmt.Printf("%+v\n", items)
}

func adapterExample(ctx context.Context, projectId string) {
	// setup
	client, err := albwaf.NewAPIClient()
	if err != nil {
		log.Fatal("create albwaf client: ", err)
	}
	teardown, err := setup(ctx, client, projectId)
	defer func() {
		err := teardown()
		if err != nil {
			log.Fatal("teardown: ", err)
		}
	}()
	if err != nil {
		log.Fatal("setup: ", err)
	}

	// list items
	fmt.Println("listing albwaf managed rule sets")
	req := client.DefaultAPI.ListManagedRuleSets(ctx, projectId, region)
	for item, err := range paginate.Items(
		listManagedRuleSetsAdapter{request: req},
		paginate.WithLimit(5),
		paginate.WithPageSize(2),
	) {
		if err != nil {
			log.Fatal("list items: ", err)
		}
		fmt.Printf("%+v\n", item)
	}
}

func setup(ctx context.Context, client *albwaf.APIClient, projectId string) (teardown func() error, err error) {
	req := client.DefaultAPI.CreateManagedRuleSet(ctx, projectId, region)
	var created []string
	teardown = func() error {
		for _, name := range created {
			_, teardownErr := client.DefaultAPI.DeleteManagedRuleSet(ctx, projectId, region, name).Execute()
			if teardownErr != nil {
				return teardownErr
			}
		}
		return nil
	}
	for n := range 10 {
		payload := albwaf.CreateManagedRuleSetPayload{
			Name: fmt.Sprintf("%s-%d", projectId, n),
			Type: albwaf.TYPE_TYPE_OWASP_CRS,
		}
		ruleSet, err := req.CreateManagedRuleSetPayload(payload).Execute()
		if err != nil {
			return teardown, err
		}
		created = append(created, ruleSet.Name)
	}
	return teardown, nil
}

// adapter
type listManagedRuleSetsAdapter struct {
	request albwaf.ApiListManagedRuleSetsRequest
}

func (a listManagedRuleSetsAdapter) PageSize(pageSize int32) listManagedRuleSetsAdapter {
	a.request = a.request.PageSize(strconv.Itoa(int(pageSize)))
	return a
}

func (a listManagedRuleSetsAdapter) PageToken(pageToken string) listManagedRuleSetsAdapter {
	a.request = a.request.PageId(pageToken)
	return a
}

func (a listManagedRuleSetsAdapter) Execute() (listManagedRuleSetsResponse, error) {
	resp, err := a.request.Execute()
	if err != nil {
		return listManagedRuleSetsResponse{}, err
	}
	return listManagedRuleSetsResponse{ListManagedRuleSetResponse: resp}, nil
}

type listManagedRuleSetsResponse struct {
	*albwaf.ListManagedRuleSetResponse
}

func (r listManagedRuleSetsResponse) GetNextPageToken() string {
	return r.GetNextPageId()
}
