/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// CredentialInputSourcesApiService CredentialInputSourcesApi service
type CredentialInputSourcesApiService service

// CredentialInputSourcesCredentialInputSourcesCreateOpts Optional parameters for the method 'CredentialInputSourcesCredentialInputSourcesCreate'
type CredentialInputSourcesCredentialInputSourcesCreateOpts struct {
	Data optional.Interface
}

/*
CredentialInputSourcesCredentialInputSourcesCreate  Create a Credential Input Source
 Make a POST request to this resource with the following credential input source fields to create a new credential input source:          * &#x60;description&#x60;: Optional description of this credential input source. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;) * &#x60;input_field_name&#x60;:  (string, required) * &#x60;metadata&#x60;:  (json, default&#x3D;&#x60;{}&#x60;) * &#x60;target_credential&#x60;:  (id, required) * &#x60;source_credential&#x60;:  (id, required)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CredentialInputSourcesCredentialInputSourcesCreateOpts - Optional Parameters:
 * @param "Data" (optional.Interface) -
*/
func (a *CredentialInputSourcesApiService) CredentialInputSourcesCredentialInputSourcesCreate(ctx _context.Context, localVarOptionals *CredentialInputSourcesCredentialInputSourcesCreateOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/credential_input_sources/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.Data.IsSet() {
		localVarPostBody = localVarOptionals.Data.Value()
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		if err != nil {
			return localVarHTTPResponse, err
		}

		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CredentialInputSourcesCredentialInputSourcesDeleteOpts Optional parameters for the method 'CredentialInputSourcesCredentialInputSourcesDelete'
type CredentialInputSourcesCredentialInputSourcesDeleteOpts struct {
	Search optional.String
}

/*
CredentialInputSourcesCredentialInputSourcesDelete  Delete a Credential Input Source
 Make a DELETE request to this resource to delete this credential input source.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CredentialInputSourcesCredentialInputSourcesDeleteOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
*/
func (a *CredentialInputSourcesApiService) CredentialInputSourcesCredentialInputSourcesDelete(ctx _context.Context, id string, localVarOptionals *CredentialInputSourcesCredentialInputSourcesDeleteOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/credential_input_sources/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		if err != nil {
			return localVarHTTPResponse, err
		}

		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CredentialInputSourcesCredentialInputSourcesListOpts Optional parameters for the method 'CredentialInputSourcesCredentialInputSourcesList'
type CredentialInputSourcesCredentialInputSourcesListOpts struct {
	Page     optional.Int32
	PageSize optional.Int32
	Search   optional.String
}

/*
CredentialInputSourcesCredentialInputSourcesList  List Credential Input Sources
 Make a GET request to this resource to retrieve the list of credential input sources.  The resulting data structure contains:      {         \&quot;count\&quot;: 99,         \&quot;next\&quot;: null,         \&quot;previous\&quot;: null,         \&quot;results\&quot;: [             ...         ]     }  The &#x60;count&#x60; field indicates the total number of credential input sources found for the given query.  The &#x60;next&#x60; and &#x60;previous&#x60; fields provides links to additional results if there are more than will fit on a single page.  The &#x60;results&#x60; list contains zero or more credential input source records.    ## Results  Each credential input source data structure includes the following fields:  * &#x60;id&#x60;: Database ID for this credential input source. (integer) * &#x60;type&#x60;: Data type for this credential input source. (choice) * &#x60;url&#x60;: URL for this credential input source. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this credential input source was created. (datetime) * &#x60;modified&#x60;: Timestamp when this credential input source was last modified. (datetime) * &#x60;description&#x60;: Optional description of this credential input source. (string) * &#x60;input_field_name&#x60;:  (string) * &#x60;metadata&#x60;:  (json) * &#x60;target_credential&#x60;:  (id) * &#x60;source_credential&#x60;:  (id)    ## Sorting  To specify that credential input sources are returned in a particular order, use the &#x60;order_by&#x60; query string parameter on the GET request.      ?order_by&#x3D;name  Prefix the field name with a dash &#x60;-&#x60; to sort in reverse:      ?order_by&#x3D;-name  Multiple sorting fields may be specified by separating the field names with a comma &#x60;,&#x60;:      ?order_by&#x3D;name,some_other_field  ## Pagination  Use the &#x60;page_size&#x60; query string parameter to change the number of results returned for each request.  Use the &#x60;page&#x60; query string parameter to retrieve a particular page of results.      ?page_size&#x3D;100&amp;page&#x3D;2  The &#x60;previous&#x60; and &#x60;next&#x60; links returned with the results will set these query string parameters automatically.  ## Searching  Use the &#x60;search&#x60; query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search&#x3D;findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search&#x3D;findme
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CredentialInputSourcesCredentialInputSourcesListOpts - Optional Parameters:
 * @param "Page" (optional.Int32) -  A page number within the paginated result set.
 * @param "PageSize" (optional.Int32) -  Number of results to return per page.
 * @param "Search" (optional.String) -  A search term.
*/
func (a *CredentialInputSourcesApiService) CredentialInputSourcesCredentialInputSourcesList(ctx _context.Context, localVarOptionals *CredentialInputSourcesCredentialInputSourcesListOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/credential_input_sources/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		if err != nil {
			return localVarHTTPResponse, err
		}

		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CredentialInputSourcesCredentialInputSourcesPartialUpdateOpts Optional parameters for the method 'CredentialInputSourcesCredentialInputSourcesPartialUpdate'
type CredentialInputSourcesCredentialInputSourcesPartialUpdateOpts struct {
	Search optional.String
	Data   optional.Interface
}

/*
CredentialInputSourcesCredentialInputSourcesPartialUpdate  Update a Credential Input Source
 Make a PUT or PATCH request to this resource to update this credential input source.  The following fields may be modified:          * &#x60;description&#x60;: Optional description of this credential input source. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;) * &#x60;input_field_name&#x60;:  (string, required) * &#x60;metadata&#x60;:  (json, default&#x3D;&#x60;{}&#x60;) * &#x60;target_credential&#x60;:  (id, required) * &#x60;source_credential&#x60;:  (id, required)         For a PATCH request, include only the fields that are being modified.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CredentialInputSourcesCredentialInputSourcesPartialUpdateOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
 * @param "Data" (optional.Interface) -
*/
func (a *CredentialInputSourcesApiService) CredentialInputSourcesCredentialInputSourcesPartialUpdate(ctx _context.Context, id string, localVarOptionals *CredentialInputSourcesCredentialInputSourcesPartialUpdateOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/credential_input_sources/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.Data.IsSet() {
		localVarPostBody = localVarOptionals.Data.Value()
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		if err != nil {
			return localVarHTTPResponse, err
		}

		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CredentialInputSourcesCredentialInputSourcesReadOpts Optional parameters for the method 'CredentialInputSourcesCredentialInputSourcesRead'
type CredentialInputSourcesCredentialInputSourcesReadOpts struct {
	Search optional.String
}

/*
CredentialInputSourcesCredentialInputSourcesRead  Retrieve a Credential Input Source
 Make GET request to this resource to retrieve a single credential input source record containing the following fields:  * &#x60;id&#x60;: Database ID for this credential input source. (integer) * &#x60;type&#x60;: Data type for this credential input source. (choice) * &#x60;url&#x60;: URL for this credential input source. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this credential input source was created. (datetime) * &#x60;modified&#x60;: Timestamp when this credential input source was last modified. (datetime) * &#x60;description&#x60;: Optional description of this credential input source. (string) * &#x60;input_field_name&#x60;:  (string) * &#x60;metadata&#x60;:  (json) * &#x60;target_credential&#x60;:  (id) * &#x60;source_credential&#x60;:  (id)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CredentialInputSourcesCredentialInputSourcesReadOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
*/
func (a *CredentialInputSourcesApiService) CredentialInputSourcesCredentialInputSourcesRead(ctx _context.Context, id string, localVarOptionals *CredentialInputSourcesCredentialInputSourcesReadOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/credential_input_sources/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		if err != nil {
			return localVarHTTPResponse, err
		}

		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CredentialInputSourcesCredentialInputSourcesUpdateOpts Optional parameters for the method 'CredentialInputSourcesCredentialInputSourcesUpdate'
type CredentialInputSourcesCredentialInputSourcesUpdateOpts struct {
	Search optional.String
	Data   optional.Interface
}

/*
CredentialInputSourcesCredentialInputSourcesUpdate  Update a Credential Input Source
 Make a PUT or PATCH request to this resource to update this credential input source.  The following fields may be modified:          * &#x60;description&#x60;: Optional description of this credential input source. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;) * &#x60;input_field_name&#x60;:  (string, required) * &#x60;metadata&#x60;:  (json, default&#x3D;&#x60;{}&#x60;) * &#x60;target_credential&#x60;:  (id, required) * &#x60;source_credential&#x60;:  (id, required)       For a PUT request, include **all** fields in the request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CredentialInputSourcesCredentialInputSourcesUpdateOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
 * @param "Data" (optional.Interface of InlineObject1) -
*/
func (a *CredentialInputSourcesApiService) CredentialInputSourcesCredentialInputSourcesUpdate(ctx _context.Context, id string, localVarOptionals *CredentialInputSourcesCredentialInputSourcesUpdateOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/credential_input_sources/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Data.IsSet() {
		localVarOptionalData, localVarOptionalDataok := localVarOptionals.Data.Value().(InlineObject1)
		if !localVarOptionalDataok {
			return nil, reportError("data should be InlineObject1")
		}
		localVarPostBody = &localVarOptionalData
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		if err != nil {
			return localVarHTTPResponse, err
		}

		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}