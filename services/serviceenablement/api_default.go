/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

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

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiDisableServiceRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	serviceId  string
}

func (r ApiDisableServiceRequest) Execute() error {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DisableService")
	if err != nil {
		return &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/services/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(ParameterValueToString(r.serviceId, "serviceId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return err
	}

	contextHTTPRequest, ok := r.ctx.Value(config.ContextHTTPRequest).(**http.Request)
	if ok {
		*contextHTTPRequest = req
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	contextHTTPResponse, ok := r.ctx.Value(config.ContextHTTPResponse).(**http.Response)
	if ok {
		*contextHTTPResponse = localVarHTTPResponse
	}
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapierror.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		return newErr
	}

	return nil
}

/*
DisableService: Method for DisableService

disables the service in a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param serviceId
	@return ApiDisableServiceRequest
*/
func (a *APIClient) DisableService(ctx context.Context, projectId string, serviceId string) ApiDisableServiceRequest {
	return ApiDisableServiceRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serviceId:  serviceId,
	}
}

func (a *APIClient) DisableServiceExecute(ctx context.Context, projectId string, serviceId string) error {
	r := ApiDisableServiceRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serviceId:  serviceId,
	}
	return r.Execute()
}

type ApiEnableServiceRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	serviceId  string
}

func (r ApiEnableServiceRequest) Execute() error {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.EnableService")
	if err != nil {
		return &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/services/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(ParameterValueToString(r.serviceId, "serviceId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return err
	}

	contextHTTPRequest, ok := r.ctx.Value(config.ContextHTTPRequest).(**http.Request)
	if ok {
		*contextHTTPRequest = req
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	contextHTTPResponse, ok := r.ctx.Value(config.ContextHTTPResponse).(**http.Response)
	if ok {
		*contextHTTPResponse = localVarHTTPResponse
	}
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapierror.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		return newErr
	}

	return nil
}

/*
EnableService: Method for EnableService

enables the service in a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param serviceId
	@return ApiEnableServiceRequest
*/
func (a *APIClient) EnableService(ctx context.Context, projectId string, serviceId string) ApiEnableServiceRequest {
	return ApiEnableServiceRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serviceId:  serviceId,
	}
}

func (a *APIClient) EnableServiceExecute(ctx context.Context, projectId string, serviceId string) error {
	r := ApiEnableServiceRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serviceId:  serviceId,
	}
	return r.Execute()
}

type ApiGetServiceStatusRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	serviceId  string
}

func (r ApiGetServiceStatusRequest) Execute() (*ServiceStatus, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceStatus
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetServiceStatus")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/services/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(ParameterValueToString(r.serviceId, "serviceId")), -1)

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

	contextHTTPRequest, ok := r.ctx.Value(config.ContextHTTPRequest).(**http.Request)
	if ok {
		*contextHTTPRequest = req
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
GetServiceStatus: Method for GetServiceStatus

returns the current status of a service in a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param serviceId
	@return ApiGetServiceStatusRequest
*/
func (a *APIClient) GetServiceStatus(ctx context.Context, projectId string, serviceId string) ApiGetServiceStatusRequest {
	return ApiGetServiceStatusRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serviceId:  serviceId,
	}
}

func (a *APIClient) GetServiceStatusExecute(ctx context.Context, projectId string, serviceId string) (*ServiceStatus, error) {
	r := ApiGetServiceStatusRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serviceId:  serviceId,
	}
	return r.Execute()
}

type ApiListServiceStatusRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	cursor     *string
}

func (r ApiListServiceStatusRequest) Cursor(cursor string) ApiListServiceStatusRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListServiceStatusRequest) Execute() (*ListServiceStatus200Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListServiceStatus200Response
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListServiceStatus")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/services"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
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

	contextHTTPRequest, ok := r.ctx.Value(config.ContextHTTPRequest).(**http.Request)
	if ok {
		*contextHTTPRequest = req
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
ListServiceStatus: Method for ListServiceStatus

returns a list of all available services for a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@return ApiListServiceStatusRequest
*/
func (a *APIClient) ListServiceStatus(ctx context.Context, projectId string) ApiListServiceStatusRequest {
	return ApiListServiceStatusRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *APIClient) ListServiceStatusExecute(ctx context.Context, projectId string) (*ListServiceStatus200Response, error) {
	r := ApiListServiceStatusRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
	return r.Execute()
}
