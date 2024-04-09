/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membership

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiAddMembersRequest struct {
	ctx               context.Context
	apiService        *DefaultApiService
	resourceId        string
	addMembersPayload *AddMembersPayload
}

func (r ApiAddMembersRequest) AddMembersPayload(addMembersPayload AddMembersPayload) ApiAddMembersRequest {
	r.addMembersPayload = &addMembersPayload
	return r
}

func (r ApiAddMembersRequest) Execute() (*MembersResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MembersResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AddMembers")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{resourceId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addMembersPayload == nil {
		return localVarReturnValue, fmt.Errorf("addMembersPayload is required and must be specified")
	}

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
	localVarPostBody = r.addMembersPayload
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
		if localVarHTTPResponse.StatusCode == 403 {
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
AddMembers Add members to a resource

Add members to the given resource with specified roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceId
	@return ApiAddMembersRequest
*/
func (a *APIClient) AddMembers(ctx context.Context, resourceId string) ApiAddMembersRequest {
	return ApiAddMembersRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		resourceId: resourceId,
	}
}

func (a *APIClient) AddMembersExecute(ctx context.Context, resourceId string) (*MembersResponse, error) {
	r := ApiAddMembersRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		resourceId: resourceId,
	}
	return r.Execute()
}

type ApiListMembersRequest struct {
	ctx          context.Context
	apiService   *DefaultApiService
	resourceType string
	resourceId   string
	subject      *string
}

func (r ApiListMembersRequest) Subject(subject string) ApiListMembersRequest {
	r.subject = &subject
	return r
}

func (r ApiListMembersRequest) Execute() (*ListMembersResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListMembersResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListMembers")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{resourceType}/{resourceId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", url.PathEscape(ParameterValueToString(r.resourceType, "resourceType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.subject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject", r.subject, "")
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
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
ListMembers Get members to a resource

List members of the given resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceType
	@param resourceId
	@return ApiListMembersRequest
*/
func (a *APIClient) ListMembers(ctx context.Context, resourceType string, resourceId string) ApiListMembersRequest {
	return ApiListMembersRequest{
		apiService:   a.defaultApi,
		ctx:          ctx,
		resourceType: resourceType,
		resourceId:   resourceId,
	}
}

func (a *APIClient) ListMembersExecute(ctx context.Context, resourceType string, resourceId string) (*ListMembersResponse, error) {
	r := ApiListMembersRequest{
		apiService:   a.defaultApi,
		ctx:          ctx,
		resourceType: resourceType,
		resourceId:   resourceId,
	}
	return r.Execute()
}

type ApiListPermissionsRequest struct {
	ctx          context.Context
	apiService   *DefaultApiService
	resourceType *string
}

func (r ApiListPermissionsRequest) ResourceType(resourceType string) ApiListPermissionsRequest {
	r.resourceType = &resourceType
	return r
}

func (r ApiListPermissionsRequest) Execute() (*ListPermissionsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListPermissionsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListPermissions")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v2/permissions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resourceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resourceType", r.resourceType, "")
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
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
ListPermissions Get available permissions

Get available permissions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListPermissionsRequest
*/
func (a *APIClient) ListPermissions(ctx context.Context) ApiListPermissionsRequest {
	return ApiListPermissionsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
}

func (a *APIClient) ListPermissionsExecute(ctx context.Context) (*ListPermissionsResponse, error) {
	r := ApiListPermissionsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
	return r.Execute()
}

type ApiListRolesRequest struct {
	ctx          context.Context
	apiService   *DefaultApiService
	resourceType string
	resourceId   string
}

func (r ApiListRolesRequest) Execute() (*RolesResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RolesResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListRoles")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{resourceType}/{resourceId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", url.PathEscape(ParameterValueToString(r.resourceType, "resourceType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(r.resourceId, "resourceId")), -1)

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
		if localVarHTTPResponse.StatusCode == 403 {
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
ListRoles Get roles and permissions of a resource

Get roles and permissions of a resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceType
	@param resourceId
	@return ApiListRolesRequest
*/
func (a *APIClient) ListRoles(ctx context.Context, resourceType string, resourceId string) ApiListRolesRequest {
	return ApiListRolesRequest{
		apiService:   a.defaultApi,
		ctx:          ctx,
		resourceType: resourceType,
		resourceId:   resourceId,
	}
}

func (a *APIClient) ListRolesExecute(ctx context.Context, resourceType string, resourceId string) (*RolesResponse, error) {
	r := ApiListRolesRequest{
		apiService:   a.defaultApi,
		ctx:          ctx,
		resourceType: resourceType,
		resourceId:   resourceId,
	}
	return r.Execute()
}

type ApiListUserMembershipsRequest struct {
	ctx              context.Context
	apiService       *DefaultApiService
	email            string
	resourceType     *string
	resourceId       *string
	parentResourceId *string
}

func (r ApiListUserMembershipsRequest) ResourceType(resourceType string) ApiListUserMembershipsRequest {
	r.resourceType = &resourceType
	return r
}

func (r ApiListUserMembershipsRequest) ResourceId(resourceId string) ApiListUserMembershipsRequest {
	r.resourceId = &resourceId
	return r
}

func (r ApiListUserMembershipsRequest) ParentResourceId(parentResourceId string) ApiListUserMembershipsRequest {
	r.parentResourceId = &parentResourceId
	return r
}

func (r ApiListUserMembershipsRequest) Execute() (*ListUserMembershipsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListUserMembershipsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListUserMemberships")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{email}/memberships"
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(ParameterValueToString(r.email, "email")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resourceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resourceType", r.resourceType, "")
	}
	if r.resourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resourceId", r.resourceId, "")
	}
	if r.parentResourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentResourceId", r.parentResourceId, "")
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
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
ListUserMemberships List memberships of a user

List memberships of a user. An administrative access is needed to list any user's memberships, while the user can do it on his/her own email. You can use filters to scope the request to a project/folder/organization. In this case -if caller is not the subject-, owner permissions are required. Because of hierarchical role bindings, the user might have permissions on more resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param email
	@return ApiListUserMembershipsRequest
*/
func (a *APIClient) ListUserMemberships(ctx context.Context, email string) ApiListUserMembershipsRequest {
	return ApiListUserMembershipsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		email:      email,
	}
}

func (a *APIClient) ListUserMembershipsExecute(ctx context.Context, email string) (*ListUserMembershipsResponse, error) {
	r := ApiListUserMembershipsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		email:      email,
	}
	return r.Execute()
}

type ApiRemoveMembersRequest struct {
	ctx                  context.Context
	apiService           *DefaultApiService
	resourceId           string
	removeMembersPayload *RemoveMembersPayload
}

func (r ApiRemoveMembersRequest) RemoveMembersPayload(removeMembersPayload RemoveMembersPayload) ApiRemoveMembersRequest {
	r.removeMembersPayload = &removeMembersPayload
	return r
}

func (r ApiRemoveMembersRequest) Execute() (*MembersResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MembersResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RemoveMembers")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{resourceId}/members/remove"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(ParameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.removeMembersPayload == nil {
		return localVarReturnValue, fmt.Errorf("removeMembersPayload is required and must be specified")
	}

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
	localVarPostBody = r.removeMembersPayload
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
		if localVarHTTPResponse.StatusCode == 403 {
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
RemoveMembers Remove members from a resource

Remove members from the given resource with specified roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceId
	@return ApiRemoveMembersRequest
*/
func (a *APIClient) RemoveMembers(ctx context.Context, resourceId string) ApiRemoveMembersRequest {
	return ApiRemoveMembersRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		resourceId: resourceId,
	}
}

func (a *APIClient) RemoveMembersExecute(ctx context.Context, resourceId string) (*MembersResponse, error) {
	r := ApiRemoveMembersRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		resourceId: resourceId,
	}
	return r.Execute()
}
