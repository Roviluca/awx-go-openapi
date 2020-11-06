/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// WorkflowApprovalTemplatesApiService WorkflowApprovalTemplatesApi service
type WorkflowApprovalTemplatesApiService service

type ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest struct {
	ctx _context.Context
	ApiService *WorkflowApprovalTemplatesApiService
	id string
	page *int32
	pageSize *int32
	search *string
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest) Page(page int32) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest {
	r.page = &page
	return r
}
func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest) PageSize(pageSize int32) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest) Search(search string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest {
	r.search = &search
	return r
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListExecute(r)
}

/*
 * WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList  List Workflow Approvals for a Workflow Approval Template
 * 
Make a GET request to this resource to retrieve a list of
workflow approvals associated with the selected
workflow approval template.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of workflow approvals
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more workflow approval records.  

## Results

Each workflow approval data structure includes the following fields:

* `id`: Database ID for this workflow approval. (integer)
* `type`: Data type for this workflow approval. (choice)
* `url`: URL for this workflow approval. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow approval was created. (datetime)
* `modified`: Timestamp when this workflow approval was last modified. (datetime)
* `name`: Name of this workflow approval. (string)
* `description`: Optional description of this workflow approval. (string)
* `unified_job_template`:  (id)
* `launch_type`:  (choice)
    - `manual`: Manual
    - `relaunch`: Relaunch
    - `callback`: Callback
    - `scheduled`: Scheduled
    - `dependency`: Dependency
    - `workflow`: Workflow
    - `webhook`: Webhook
    - `sync`: Sync
    - `scm`: SCM Update
* `status`:  (choice)
    - `new`: New
    - `pending`: Pending
    - `waiting`: Waiting
    - `running`: Running
    - `successful`: Successful
    - `failed`: Failed
    - `error`: Error
    - `canceled`: Canceled
* `failed`:  (boolean)
* `started`: The date and time the job was queued for starting. (datetime)
* `finished`: The date and time the job finished execution. (datetime)
* `canceled_on`: The date and time when the cancel request was sent. (datetime)
* `elapsed`: Elapsed time in seconds that the job ran. (decimal)
* `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string)
* `can_approve_or_deny`:  (field)
* `approval_expiration`:  (field)
* `timed_out`:  (boolean)



## Sorting

To specify that workflow approvals are returned in a particular
order, use the `order_by` query string parameter on the GET request.

    ?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

    ?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

    ?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

    ?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

    ?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

    ?related__search=findme
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList(ctx _context.Context, id string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest {
	return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListExecute(r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowApprovalTemplatesApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/workflow_approval_templates/{id}/approvals/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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

type ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest struct {
	ctx _context.Context
	ApiService *WorkflowApprovalTemplatesApiService
	id string
	search *string
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest) Search(search string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest {
	r.search = &search
	return r
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteExecute(r)
}

/*
 * WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete  Delete a Workflow Approval Template
 * 
Make a DELETE request to this resource to delete this workflow approval template.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete(ctx _context.Context, id string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest {
	return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteExecute(r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowApprovalTemplatesApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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

type ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest struct {
	ctx _context.Context
	ApiService *WorkflowApprovalTemplatesApiService
	id string
	search *string
	data *InlineObject73
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest) Search(search string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest {
	r.search = &search
	return r
}
func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest) Data(data InlineObject73) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest {
	r.data = &data
	return r
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateExecute(r)
}

/*
 * WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate  Update a Workflow Approval Template
 * 
Make a PUT or PATCH request to this resource to update this
workflow approval template.  The following fields may be modified:









* `name`: Name of this workflow approval template. (string, required)
* `description`: Optional description of this workflow approval template. (string, default=`""`)




* `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer, default=`0`)








For a PATCH request, include only the fields that are being modified.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate(ctx _context.Context, id string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest {
	return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateExecute(r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowApprovalTemplatesApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
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
	localVarPostBody = r.data
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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

type ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest struct {
	ctx _context.Context
	ApiService *WorkflowApprovalTemplatesApiService
	id string
	search *string
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest) Search(search string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest {
	r.search = &search
	return r
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesReadExecute(r)
}

/*
 * WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead  Retrieve a Workflow Approval Template
 * 
Make GET request to this resource to retrieve a single workflow approval template
record containing the following fields:

* `id`: Database ID for this workflow approval template. (integer)
* `type`: Data type for this workflow approval template. (choice)
* `url`: URL for this workflow approval template. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow approval template was created. (datetime)
* `modified`: Timestamp when this workflow approval template was last modified. (datetime)
* `name`: Name of this workflow approval template. (string)
* `description`: Optional description of this workflow approval template. (string)
* `last_job_run`:  (datetime)
* `last_job_failed`:  (boolean)
* `next_job_run`:  (datetime)
* `status`:  (choice)
    - `new`: New
    - `pending`: Pending
    - `waiting`: Waiting
    - `running`: Running
    - `successful`: Successful
    - `failed`: Failed
    - `error`: Error
    - `canceled`: Canceled
    - `never updated`: Never Updated
    - `ok`: OK
    - `missing`: Missing
    - `none`: No External Source
    - `updating`: Updating
* `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead(ctx _context.Context, id string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest {
	return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesReadExecute(r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowApprovalTemplatesApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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

type ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest struct {
	ctx _context.Context
	ApiService *WorkflowApprovalTemplatesApiService
	id string
	search *string
	data *InlineObject72
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest) Search(search string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest {
	r.search = &search
	return r
}
func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest) Data(data InlineObject72) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest {
	r.data = &data
	return r
}

func (r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateExecute(r)
}

/*
 * WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate  Update a Workflow Approval Template
 * 
Make a PUT or PATCH request to this resource to update this
workflow approval template.  The following fields may be modified:









* `name`: Name of this workflow approval template. (string, required)
* `description`: Optional description of this workflow approval template. (string, default=`""`)




* `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer, default=`0`)






For a PUT request, include **all** fields in the request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate(ctx _context.Context, id string) ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest {
	return ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *WorkflowApprovalTemplatesApiService) WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateExecute(r ApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowApprovalTemplatesApiService.WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/workflow_approval_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
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
	localVarPostBody = r.data
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
