/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// AuthAPIService AuthAPI service
type AuthAPIService service

type ApiLoginApiApiV2AuthApiTokenPostRequest struct {
	ctx        context.Context
	ApiService *AuthAPIService
}

func (r ApiLoginApiApiV2AuthApiTokenPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.LoginApiApiV2AuthApiTokenPostExecute(r)
}

/*
LoginApiApiV2AuthApiTokenPost Login Api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLoginApiApiV2AuthApiTokenPostRequest
*/
func (a *AuthAPIService) LoginApiApiV2AuthApiTokenPost(ctx context.Context) ApiLoginApiApiV2AuthApiTokenPostRequest {
	return ApiLoginApiApiV2AuthApiTokenPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *AuthAPIService) LoginApiApiV2AuthApiTokenPostExecute(r ApiLoginApiApiV2AuthApiTokenPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.LoginApiApiV2AuthApiTokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/auth/api-token"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-yeti-apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiLoginApiV2AuthTokenPostRequest struct {
	ctx          context.Context
	ApiService   *AuthAPIService
	username     *string
	password     *string
	grantType    *string
	scope        *string
	clientId     *string
	clientSecret *string
}

func (r ApiLoginApiV2AuthTokenPostRequest) Username(username string) ApiLoginApiV2AuthTokenPostRequest {
	r.username = &username
	return r
}

func (r ApiLoginApiV2AuthTokenPostRequest) Password(password string) ApiLoginApiV2AuthTokenPostRequest {
	r.password = &password
	return r
}

func (r ApiLoginApiV2AuthTokenPostRequest) GrantType(grantType string) ApiLoginApiV2AuthTokenPostRequest {
	r.grantType = &grantType
	return r
}

func (r ApiLoginApiV2AuthTokenPostRequest) Scope(scope string) ApiLoginApiV2AuthTokenPostRequest {
	r.scope = &scope
	return r
}

func (r ApiLoginApiV2AuthTokenPostRequest) ClientId(clientId string) ApiLoginApiV2AuthTokenPostRequest {
	r.clientId = &clientId
	return r
}

func (r ApiLoginApiV2AuthTokenPostRequest) ClientSecret(clientSecret string) ApiLoginApiV2AuthTokenPostRequest {
	r.clientSecret = &clientSecret
	return r
}

func (r ApiLoginApiV2AuthTokenPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.LoginApiV2AuthTokenPostExecute(r)
}

/*
LoginApiV2AuthTokenPost Login

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLoginApiV2AuthTokenPostRequest
*/
func (a *AuthAPIService) LoginApiV2AuthTokenPost(ctx context.Context) ApiLoginApiV2AuthTokenPostRequest {
	return ApiLoginApiV2AuthTokenPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *AuthAPIService) LoginApiV2AuthTokenPostExecute(r ApiLoginApiV2AuthTokenPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.LoginApiV2AuthTokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/auth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.username == nil {
		return localVarReturnValue, nil, reportError("username is required and must be specified")
	}
	if r.password == nil {
		return localVarReturnValue, nil, reportError("password is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "username", r.username, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scope", r.scope, "", "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_id", r.clientId, "", "")
	}
	if r.clientSecret != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_secret", r.clientSecret, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiLogoutApiV2AuthLogoutPostRequest struct {
	ctx        context.Context
	ApiService *AuthAPIService
}

func (r ApiLogoutApiV2AuthLogoutPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.LogoutApiV2AuthLogoutPostExecute(r)
}

/*
LogoutApiV2AuthLogoutPost Logout

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLogoutApiV2AuthLogoutPostRequest
*/
func (a *AuthAPIService) LogoutApiV2AuthLogoutPost(ctx context.Context) ApiLogoutApiV2AuthLogoutPostRequest {
	return ApiLogoutApiV2AuthLogoutPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *AuthAPIService) LogoutApiV2AuthLogoutPostExecute(r ApiLogoutApiV2AuthLogoutPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.LogoutApiV2AuthLogoutPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/auth/logout"

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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMeApiV2AuthMeGetRequest struct {
	ctx        context.Context
	ApiService *AuthAPIService
}

func (r ApiMeApiV2AuthMeGetRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.MeApiV2AuthMeGetExecute(r)
}

/*
MeApiV2AuthMeGet Me

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiMeApiV2AuthMeGetRequest
*/
func (a *AuthAPIService) MeApiV2AuthMeGet(ctx context.Context) ApiMeApiV2AuthMeGetRequest {
	return ApiMeApiV2AuthMeGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return User
func (a *AuthAPIService) MeApiV2AuthMeGetExecute(r ApiMeApiV2AuthMeGetRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.MeApiV2AuthMeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/auth/me"

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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
