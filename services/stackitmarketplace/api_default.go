/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

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

type ApiApproveSubscriptionRequest struct {
	ctx            context.Context
	apiService     *DefaultApiService
	projectId      string
	subscriptionId string
}

func (r ApiApproveSubscriptionRequest) Execute() error {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ApproveSubscription")
	if err != nil {
		return &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vendors/projects/{projectId}/subscriptions/{subscriptionId}/approve"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(ParameterValueToString(r.subscriptionId, "subscriptionId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return newErr
			}
			newErr.ErrorMessage = oapierror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
		}
		return newErr
	}

	return nil
}

/*
ApproveSubscription: Approve a subscription

Approve a pending subscription.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId The project ID.
	@param subscriptionId The subscription ID.
	@return ApiApproveSubscriptionRequest
*/
func (a *APIClient) ApproveSubscription(ctx context.Context, projectId string, subscriptionId string) ApiApproveSubscriptionRequest {
	return ApiApproveSubscriptionRequest{
		apiService:     a.defaultApi,
		ctx:            ctx,
		projectId:      projectId,
		subscriptionId: subscriptionId,
	}
}

func (a *APIClient) ApproveSubscriptionExecute(ctx context.Context, projectId string, subscriptionId string) error {
	r := ApiApproveSubscriptionRequest{
		apiService:     a.defaultApi,
		ctx:            ctx,
		projectId:      projectId,
		subscriptionId: subscriptionId,
	}
	return r.Execute()
}

type ApiGetCatalogProductRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	productId  string
	locale     *string
}

// The language of the response.

func (r ApiGetCatalogProductRequest) Locale(locale string) ApiGetCatalogProductRequest {
	r.locale = &locale
	return r
}

func (r ApiGetCatalogProductRequest) Execute() (*CatalogProductDetail, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CatalogProductDetail
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetCatalogProduct")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/catalog/products/{productId}"
	localVarPath = strings.Replace(localVarPath, "{"+"productId"+"}", url.PathEscape(ParameterValueToString(r.productId, "productId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "")
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
GetCatalogProduct: Get a product

Get a product.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param productId The product ID.
	@return ApiGetCatalogProductRequest
*/
func (a *APIClient) GetCatalogProduct(ctx context.Context, productId string) ApiGetCatalogProductRequest {
	return ApiGetCatalogProductRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		productId:  productId,
	}
}

func (a *APIClient) GetCatalogProductExecute(ctx context.Context, productId string) (*CatalogProductDetail, error) {
	r := ApiGetCatalogProductRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		productId:  productId,
	}
	return r.Execute()
}

type ApiGetVendorSubscriptionRequest struct {
	ctx            context.Context
	apiService     *DefaultApiService
	projectId      string
	subscriptionId string
}

func (r ApiGetVendorSubscriptionRequest) Execute() (*VendorSubscription, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VendorSubscription
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetVendorSubscription")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vendors/projects/{projectId}/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(ParameterValueToString(r.subscriptionId, "subscriptionId")), -1)

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
GetVendorSubscription: Get a subscription

Get a subscription.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId The project ID.
	@param subscriptionId The subscription ID.
	@return ApiGetVendorSubscriptionRequest
*/
func (a *APIClient) GetVendorSubscription(ctx context.Context, projectId string, subscriptionId string) ApiGetVendorSubscriptionRequest {
	return ApiGetVendorSubscriptionRequest{
		apiService:     a.defaultApi,
		ctx:            ctx,
		projectId:      projectId,
		subscriptionId: subscriptionId,
	}
}

func (a *APIClient) GetVendorSubscriptionExecute(ctx context.Context, projectId string, subscriptionId string) (*VendorSubscription, error) {
	r := ApiGetVendorSubscriptionRequest{
		apiService:     a.defaultApi,
		ctx:            ctx,
		projectId:      projectId,
		subscriptionId: subscriptionId,
	}
	return r.Execute()
}

type ApiListCatalogProductsRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	cursor     *string
	limit      *float32
	locale     *string
	filter     *string
}

// A pagination cursor that represents a position in the dataset. If given, results will be returned from the item after the cursor. If not given, results will be returned from the beginning.

func (r ApiListCatalogProductsRequest) Cursor(cursor string) ApiListCatalogProductsRequest {
	r.cursor = &cursor
	return r
}

// The maximum number of items to return in the response. If maximum is exceeded, maximum is used.

func (r ApiListCatalogProductsRequest) Limit(limit float32) ApiListCatalogProductsRequest {
	r.limit = &limit
	return r
}

// The language of the response.

func (r ApiListCatalogProductsRequest) Locale(locale string) ApiListCatalogProductsRequest {
	r.locale = &locale
	return r
}

// Filter the products based on attributes. E.g &#x60;deliveryMethod eq \&quot;SAAS\&quot;&#x60;. The supported attributes are &#x60;deliveryMethod&#x60;, &#x60;priceType&#x60;, &#x60;category&#x60;, &#x60;vendorId&#x60;, &#x60;vendorName&#x60;, and &#x60;name&#x60;. The supported operators are &#x60;eq&#x60;. Filters can be joined with &#x60;and&#x60; or &#x60;or&#x60;.

func (r ApiListCatalogProductsRequest) Filter(filter string) ApiListCatalogProductsRequest {
	r.filter = &filter
	return r
}

func (r ApiListCatalogProductsRequest) Execute() (*ListCatalogProductsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListCatalogProductsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListCatalogProducts")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/catalog/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
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
ListCatalogProducts: List all products

List all products.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCatalogProductsRequest
*/
func (a *APIClient) ListCatalogProducts(ctx context.Context) ApiListCatalogProductsRequest {
	return ApiListCatalogProductsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
}

func (a *APIClient) ListCatalogProductsExecute(ctx context.Context) (*ListCatalogProductsResponse, error) {
	r := ApiListCatalogProductsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
	return r.Execute()
}

type ApiListVendorSubscriptionsRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	cursor     *string
	limit      *float32
	productId  *string
}

// A pagination cursor that represents a position in the dataset. If given, results will be returned from the item after the cursor. If not given, results will be returned from the beginning.

func (r ApiListVendorSubscriptionsRequest) Cursor(cursor string) ApiListVendorSubscriptionsRequest {
	r.cursor = &cursor
	return r
}

// The maximum number of items to return in the response. If maximum is exceeded, maximum is used.

func (r ApiListVendorSubscriptionsRequest) Limit(limit float32) ApiListVendorSubscriptionsRequest {
	r.limit = &limit
	return r
}

// The product ID.

func (r ApiListVendorSubscriptionsRequest) ProductId(productId string) ApiListVendorSubscriptionsRequest {
	r.productId = &productId
	return r
}

func (r ApiListVendorSubscriptionsRequest) Execute() (*ListVendorSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListVendorSubscriptionsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListVendorSubscriptions")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vendors/projects/{projectId}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "")
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
ListVendorSubscriptions: List all subscriptions

List all subscriptions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId The project ID.
	@return ApiListVendorSubscriptionsRequest
*/
func (a *APIClient) ListVendorSubscriptions(ctx context.Context, projectId string) ApiListVendorSubscriptionsRequest {
	return ApiListVendorSubscriptionsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *APIClient) ListVendorSubscriptionsExecute(ctx context.Context, projectId string) (*ListVendorSubscriptionsResponse, error) {
	r := ApiListVendorSubscriptionsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
	return r.Execute()
}

type ApiResolveCustomerRequest struct {
	ctx                    context.Context
	apiService             *DefaultApiService
	projectId              string
	resolveCustomerPayload *ResolveCustomerPayload
}

func (r ApiResolveCustomerRequest) ResolveCustomerPayload(resolveCustomerPayload ResolveCustomerPayload) ApiResolveCustomerRequest {
	r.resolveCustomerPayload = &resolveCustomerPayload
	return r
}

func (r ApiResolveCustomerRequest) Execute() (*VendorSubscription, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VendorSubscription
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ResolveCustomer")
	if err != nil {
		return localVarReturnValue, &oapierror.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vendors/projects/{projectId}/resolve-customer"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resolveCustomerPayload == nil {
		return localVarReturnValue, fmt.Errorf("resolveCustomerPayload is required and must be specified")
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
	localVarPostBody = r.resolveCustomerPayload
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
ResolveCustomer: Resolve customer

Get details about the requested subscription.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId The project ID.
	@return ApiResolveCustomerRequest
*/
func (a *APIClient) ResolveCustomer(ctx context.Context, projectId string) ApiResolveCustomerRequest {
	return ApiResolveCustomerRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *APIClient) ResolveCustomerExecute(ctx context.Context, projectId string) (*VendorSubscription, error) {
	r := ApiResolveCustomerRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
	return r.Execute()
}
