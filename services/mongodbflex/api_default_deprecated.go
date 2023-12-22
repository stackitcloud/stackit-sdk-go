// Will be deleted next time SDK generation happens

package mongodbflex

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// Deprecated: this struct will be removed in the next minor upgrade, use ApiListStoragesRequest instead
type ApiGetStorageRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	flavor     string
}

func (r ApiGetStorageRequest) Execute() (*ListStoragesResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListStoragesResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetStorage")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/storages/{flavor}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flavor"+"}", url.PathEscape(ParameterValueToString(r.flavor, "flavor")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	contextHTTPResponse, ok := r.ctx.Value(config.ContextHTTPResponse).(**http.Response)
	if ok {
		*contextHTTPResponse = localVarHTTPResponse
	}
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapierror.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &oapierror.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
GetStorage Get storage
returns the storage for a certain flavor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId project id
	@param flavor flavor id
	@return ApiGetStorageRequest
*/
// Deprecated: this method will be removed in the next minor upgrade, use ListStorage instead
func (a *APIClient) GetStorage(ctx context.Context, projectId string, flavor string) ApiGetStorageRequest {
	return ApiGetStorageRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		flavor:     flavor,
	}
}

// Deprecated: this method will be removed in the next minor upgrade,  use ListStorageExecute instead
func (a *APIClient) GetStorageExecute(ctx context.Context, projectId string, flavor string) (*ListStoragesResponse, error) {
	r := ApiGetStorageRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		flavor:     flavor,
	}
	return r.Execute()
}
