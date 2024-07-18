/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

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

type ApiCreateCommandRequest struct {
	ctx                  context.Context
	apiService           *DefaultApiService
	projectId            string
	serverId             string
	createCommandPayload *CreateCommandPayload
}

// Command to post

func (r ApiCreateCommandRequest) CreateCommandPayload(createCommandPayload CreateCommandPayload) ApiCreateCommandRequest {
	r.createCommandPayload = &createCommandPayload
	return r
}

func (r ApiCreateCommandRequest) Execute() (*NewCommandResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *NewCommandResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateCommand")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/servers/{serverId}/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(r.serverId, "serverId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createCommandPayload
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
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
			var v ErrorResponse
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
CreateCommand: Method for CreateCommand

Creates a new command for execution

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId ID of the project
	@param serverId Server ID of the machine
	@return ApiCreateCommandRequest
*/
func (a *APIClient) CreateCommand(ctx context.Context, projectId string, serverId string) ApiCreateCommandRequest {
	return ApiCreateCommandRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serverId:   serverId,
	}
}

func (a *APIClient) CreateCommandExecute(ctx context.Context, projectId string, serverId string) (*NewCommandResponse, error) {
	r := ApiCreateCommandRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serverId:   serverId,
	}
	return r.Execute()
}

type ApiGetCommandRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	serverId   string
	commandId  string
}

func (r ApiGetCommandRequest) Execute() (*CommandDetails, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommandDetails
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetCommand")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/servers/{serverId}/commands/{commandId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(r.serverId, "serverId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commandId"+"}", url.PathEscape(ParameterValueToString(r.commandId, "commandId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
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
			var v ErrorResponse
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
GetCommand: Method for GetCommand

Returns details about a command

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId ID of the project
	@param serverId Server ID of the machine
	@param commandId ID of the command
	@return ApiGetCommandRequest
*/
func (a *APIClient) GetCommand(ctx context.Context, projectId string, serverId string, commandId string) ApiGetCommandRequest {
	return ApiGetCommandRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serverId:   serverId,
		commandId:  commandId,
	}
}

func (a *APIClient) GetCommandExecute(ctx context.Context, projectId string, serverId string, commandId string) (*CommandDetails, error) {
	r := ApiGetCommandRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serverId:   serverId,
		commandId:  commandId,
	}
	return r.Execute()
}

type ApiGetCommandTemplateRequest struct {
	ctx                 context.Context
	apiService          *DefaultApiService
	projectId           string
	serverId            string
	commandTemplateName string
}

func (r ApiGetCommandTemplateRequest) Execute() (*CommandTemplateSchema, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommandTemplateSchema
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetCommandTemplate")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/servers/{serverId}/command-templates/{commandTemplateName}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(r.serverId, "serverId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commandTemplateName"+"}", url.PathEscape(ParameterValueToString(r.commandTemplateName, "commandTemplateName")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
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
			var v ErrorResponse
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
GetCommandTemplate: Method for GetCommandTemplate

Returns details about a command template

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId ID of the project
	@param serverId Server ID of the machine
	@param commandTemplateName Name of the template
	@return ApiGetCommandTemplateRequest
*/
func (a *APIClient) GetCommandTemplate(ctx context.Context, projectId string, serverId string, commandTemplateName string) ApiGetCommandTemplateRequest {
	return ApiGetCommandTemplateRequest{
		apiService:          a.defaultApi,
		ctx:                 ctx,
		projectId:           projectId,
		serverId:            serverId,
		commandTemplateName: commandTemplateName,
	}
}

func (a *APIClient) GetCommandTemplateExecute(ctx context.Context, projectId string, serverId string, commandTemplateName string) (*CommandTemplateSchema, error) {
	r := ApiGetCommandTemplateRequest{
		apiService:          a.defaultApi,
		ctx:                 ctx,
		projectId:           projectId,
		serverId:            serverId,
		commandTemplateName: commandTemplateName,
	}
	return r.Execute()
}

type ApiListCommandTemplatesRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	osType     *string
}

// The type of the Operating System (windows or linux). If not provided will return data for all OS types.

func (r ApiListCommandTemplatesRequest) OsType(osType string) ApiListCommandTemplatesRequest {
	r.osType = &osType
	return r
}

func (r ApiListCommandTemplatesRequest) Execute() (*CommandTemplateResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommandTemplateResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListCommandTemplates")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/command-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.osType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "osType", r.osType, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
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
ListCommandTemplates: Method for ListCommandTemplates

Returns a list of command templates

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCommandTemplatesRequest
*/
func (a *APIClient) ListCommandTemplates(ctx context.Context) ApiListCommandTemplatesRequest {
	return ApiListCommandTemplatesRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
}

func (a *APIClient) ListCommandTemplatesExecute(ctx context.Context) (*CommandTemplateResponse, error) {
	r := ApiListCommandTemplatesRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
	return r.Execute()
}

type ApiListCommandsRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	serverId   string
}

func (r ApiListCommandsRequest) Execute() (*GetCommandsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCommandsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListCommands")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/servers/{serverId}/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverId"+"}", url.PathEscape(ParameterValueToString(r.serverId, "serverId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
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
			var v ErrorResponse
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
ListCommands: Method for ListCommands

Returns a list of commands

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId ID of the project
	@param serverId Server ID of the machine
	@return ApiListCommandsRequest
*/
func (a *APIClient) ListCommands(ctx context.Context, projectId string, serverId string) ApiListCommandsRequest {
	return ApiListCommandsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serverId:   serverId,
	}
}

func (a *APIClient) ListCommandsExecute(ctx context.Context, projectId string, serverId string) (*GetCommandsResponse, error) {
	r := ApiListCommandsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		serverId:   serverId,
	}
	return r.Execute()
}
