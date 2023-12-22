// Will be deleted next time SDK generation happens

package mongodbflex

import (
	"context"
)

// Deprecated: this struct will be removed in the next minor update, use ApiListStoragesRequest instead
type ApiGetStorageRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	flavor     string
}

func (r ApiGetStorageRequest) Execute() (*ListStoragesResponse, error) {
	return ApiListStoragesRequest(r).Execute()
}

/*
GetStorage Get storage
returns the storage for a certain flavor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId project id
	@param flavor flavor id
	@return ApiGetStorageRequest
*/
// Deprecated: this method will be removed in the next minor update, use ListStorage instead
func (a *APIClient) GetStorage(ctx context.Context, projectId, flavor string) ApiGetStorageRequest {
	return ApiGetStorageRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		flavor:     flavor,
	}
}

// Deprecated: this method will be removed in the next minor update,  use ListStorageExecute instead
func (a *APIClient) GetStorageExecute(ctx context.Context, projectId, flavor string) (*ListStoragesResponse, error) {
	r := ApiGetStorageRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		flavor:     flavor,
	}
	return r.Execute()
}
