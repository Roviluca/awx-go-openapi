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
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// WorkflowApprovalTemplatesApiService WorkflowApprovalTemplatesApi service
type WorkflowApprovalTemplatesApiService service

// WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListOpts Optional parameters for the method 'WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList'
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListOpts struct {
    Page optional.Int32
    PageSize optional.Int32
    Search optional.String
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList  List Workflow Approvals for a Workflow Approval Template
 Make a GET request to this resource to retrieve a list of workflow approvals associated with the selected workflow approval template.  The resulting data structure contains:      {         \&quot;count\&quot;: 99,         \&quot;next\&quot;: null,         \&quot;previous\&quot;: null,         \&quot;results\&quot;: [             ...         ]     }  The &#x60;count&#x60; field indicates the total number of workflow approvals found for the given query.  The &#x60;next&#x60; and &#x60;previous&#x60; fields provides links to additional results if there are more than will fit on a single page.  The &#x60;results&#x60; list contains zero or more workflow approval records.    ## Results  Each workflow approval data structure includes the following fields:  * &#x60;id&#x60;: Database ID for this workflow approval. (integer) * &#x60;type&#x60;: Data type for this workflow approval. (choice) * &#x60;url&#x60;: URL for this workflow approval. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this workflow approval was created. (datetime) * &#x60;modified&#x60;: Timestamp when this workflow approval was last modified. (datetime) * &#x60;name&#x60;: Name of this workflow approval. (string) * &#x60;description&#x60;: Optional description of this workflow approval. (string) * &#x60;unified_job_template&#x60;:  (id) * &#x60;launch_type&#x60;:  (choice)     - &#x60;manual&#x60;: Manual     - &#x60;relaunch&#x60;: Relaunch     - &#x60;callback&#x60;: Callback     - &#x60;scheduled&#x60;: Scheduled     - &#x60;dependency&#x60;: Dependency     - &#x60;workflow&#x60;: Workflow     - &#x60;webhook&#x60;: Webhook     - &#x60;sync&#x60;: Sync     - &#x60;scm&#x60;: SCM Update * &#x60;status&#x60;:  (choice)     - &#x60;new&#x60;: New     - &#x60;pending&#x60;: Pending     - &#x60;waiting&#x60;: Waiting     - &#x60;running&#x60;: Running     - &#x60;successful&#x60;: Successful     - &#x60;failed&#x60;: Failed     - &#x60;error&#x60;: Error     - &#x60;canceled&#x60;: Canceled * &#x60;failed&#x60;:  (boolean) * &#x60;started&#x60;: The date and time the job was queued for starting. (datetime) * &#x60;finished&#x60;: The date and time the job finished execution. (datetime) * &#x60;canceled_on&#x60;: The date and time when the cancel request was sent. (datetime) * &#x60;elapsed&#x60;: Elapsed time in seconds that the job ran. (decimal) * &#x60;job_explanation&#x60;: A status field to indicate the state of the job if it wasn&amp;#39;t able to run and capture stdout (string) * &#x60;can_approve_or_deny&#x60;:  (field) * &#x60;approval_expiration&#x60;:  (field) * &#x60;timed_out&#x60;:  (boolean)    ## Sorting  To specify that workflow approvals are returned in a particular order, use the &#x60;order_by&#x60; query string parameter on the GET request.      ?order_by&#x3D;name  Prefix the field name with a dash &#x60;-&#x60; to sort in reverse:      ?order_by&#x3D;-name  Multiple sorting fields may be specified by separating the field names with a comma &#x60;,&#x60;:      ?order_by&#x3D;name,some_other_field  ## Pagination  Use the &#x60;page_size&#x60; query string parameter to change the number of results returned for each request.  Use the &#x60;page&#x60; query string parameter to retrieve a particular page of results.      ?page_size&#x3D;100&amp;page&#x3D;2  The &#x60;previous&#x60; and &#x60;next&#x60; links returned with the results will set these query string parameters automatically.  ## Searching  Use the &#x60;search&#x60; query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search&#x3D;findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search&#x3D;findme
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListOpts - Optional Parameters:
 * @param "Page" (optional.Int32) -  A page number within the paginated result set.
 * @param "PageSize" (optional.Int32) -  Number of results to return per page.
 * @param "Search" (optional.String) -  A search term.
*/
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList(ctx _context.Context, id string, localVarOptionals *WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/workflow_approval_templates/{id}/approvals/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteOpts Optional parameters for the method 'WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete'
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteOpts struct {
    Search optional.String
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete  Delete a Workflow Approval Template
 Make a DELETE request to this resource to delete this workflow approval template.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
*/
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete(ctx _context.Context, id string, localVarOptionals *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateOpts Optional parameters for the method 'WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate'
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateOpts struct {
    Search optional.String
    Data optional.Interface
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate  Update a Workflow Approval Template
 Make a PUT or PATCH request to this resource to update this workflow approval template.  The following fields may be modified:          * &#x60;name&#x60;: Name of this workflow approval template. (string, required) * &#x60;description&#x60;: Optional description of this workflow approval template. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;)     * &#x60;timeout&#x60;: The amount of time (in seconds) before the approval node expires and fails. (integer, default&#x3D;&#x60;0&#x60;)         For a PATCH request, include only the fields that are being modified.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
 * @param "Data" (optional.Interface of InlineObject73) - 
*/
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate(ctx _context.Context, id string, localVarOptionals *WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

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
		localVarOptionalData, localVarOptionalDataok := localVarOptionals.Data.Value().(InlineObject73)
		if !localVarOptionalDataok {
			return nil, reportError("data should be InlineObject73")
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// WorkflowApprovalTemplatesWorkflowApprovalTemplatesReadOpts Optional parameters for the method 'WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead'
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesReadOpts struct {
    Search optional.String
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead  Retrieve a Workflow Approval Template
 Make GET request to this resource to retrieve a single workflow approval template record containing the following fields:  * &#x60;id&#x60;: Database ID for this workflow approval template. (integer) * &#x60;type&#x60;: Data type for this workflow approval template. (choice) * &#x60;url&#x60;: URL for this workflow approval template. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this workflow approval template was created. (datetime) * &#x60;modified&#x60;: Timestamp when this workflow approval template was last modified. (datetime) * &#x60;name&#x60;: Name of this workflow approval template. (string) * &#x60;description&#x60;: Optional description of this workflow approval template. (string) * &#x60;last_job_run&#x60;:  (datetime) * &#x60;last_job_failed&#x60;:  (boolean) * &#x60;next_job_run&#x60;:  (datetime) * &#x60;status&#x60;:  (choice)     - &#x60;new&#x60;: New     - &#x60;pending&#x60;: Pending     - &#x60;waiting&#x60;: Waiting     - &#x60;running&#x60;: Running     - &#x60;successful&#x60;: Successful     - &#x60;failed&#x60;: Failed     - &#x60;error&#x60;: Error     - &#x60;canceled&#x60;: Canceled     - &#x60;never updated&#x60;: Never Updated     - &#x60;ok&#x60;: OK     - &#x60;missing&#x60;: Missing     - &#x60;none&#x60;: No External Source     - &#x60;updating&#x60;: Updating * &#x60;timeout&#x60;: The amount of time (in seconds) before the approval node expires and fails. (integer)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *WorkflowApprovalTemplatesWorkflowApprovalTemplatesReadOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
*/
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead(ctx _context.Context, id string, localVarOptionals *WorkflowApprovalTemplatesWorkflowApprovalTemplatesReadOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateOpts Optional parameters for the method 'WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate'
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateOpts struct {
    Search optional.String
    Data optional.Interface
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate  Update a Workflow Approval Template
 Make a PUT or PATCH request to this resource to update this workflow approval template.  The following fields may be modified:          * &#x60;name&#x60;: Name of this workflow approval template. (string, required) * &#x60;description&#x60;: Optional description of this workflow approval template. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;)     * &#x60;timeout&#x60;: The amount of time (in seconds) before the approval node expires and fails. (integer, default&#x3D;&#x60;0&#x60;)       For a PUT request, include **all** fields in the request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
 * @param "Data" (optional.Interface of InlineObject72) - 
*/
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate(ctx _context.Context, id string, localVarOptionals *WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

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
		localVarOptionalData, localVarOptionalDataok := localVarOptionals.Data.Value().(InlineObject72)
		if !localVarOptionalDataok {
			return nil, reportError("data should be InlineObject72")
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
