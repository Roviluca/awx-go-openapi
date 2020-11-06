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

// JobEventsApiService JobEventsApi service
type JobEventsApiService service

type ApiJobEventsJobEventsChildrenListRequest struct {
	ctx _context.Context
	ApiService *JobEventsApiService
	id string
	page *int32
	pageSize *int32
	search *string
}

func (r ApiJobEventsJobEventsChildrenListRequest) Page(page int32) ApiJobEventsJobEventsChildrenListRequest {
	r.page = &page
	return r
}
func (r ApiJobEventsJobEventsChildrenListRequest) PageSize(pageSize int32) ApiJobEventsJobEventsChildrenListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiJobEventsJobEventsChildrenListRequest) Search(search string) ApiJobEventsJobEventsChildrenListRequest {
	r.search = &search
	return r
}

func (r ApiJobEventsJobEventsChildrenListRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.JobEventsJobEventsChildrenListExecute(r)
}

/*
 * JobEventsJobEventsChildrenList  List Job Events for a Job Event
 * 
Make a GET request to this resource to retrieve a list of
job events associated with the selected
job event.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of job events
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more job event records.  

## Results

Each job event data structure includes the following fields:

* `id`: Database ID for this job event. (integer)
* `type`: Data type for this job event. (choice)
* `url`: URL for this job event. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this job event was created. (datetime)
* `modified`: Timestamp when this job event was last modified. (datetime)
* `job`:  (id)
* `event`:  (choice)
    - `runner_on_failed`: Host Failed
    - `runner_on_start`: Host Started
    - `runner_on_ok`: Host OK
    - `runner_on_error`: Host Failure
    - `runner_on_skipped`: Host Skipped
    - `runner_on_unreachable`: Host Unreachable
    - `runner_on_no_hosts`: No Hosts Remaining
    - `runner_on_async_poll`: Host Polling
    - `runner_on_async_ok`: Host Async OK
    - `runner_on_async_failed`: Host Async Failure
    - `runner_item_on_ok`: Item OK
    - `runner_item_on_failed`: Item Failed
    - `runner_item_on_skipped`: Item Skipped
    - `runner_retry`: Host Retry
    - `runner_on_file_diff`: File Difference
    - `playbook_on_start`: Playbook Started
    - `playbook_on_notify`: Running Handlers
    - `playbook_on_include`: Including File
    - `playbook_on_no_hosts_matched`: No Hosts Matched
    - `playbook_on_no_hosts_remaining`: No Hosts Remaining
    - `playbook_on_task_start`: Task Started
    - `playbook_on_vars_prompt`: Variables Prompted
    - `playbook_on_setup`: Gathering Facts
    - `playbook_on_import_for_host`: internal: on Import for Host
    - `playbook_on_not_import_for_host`: internal: on Not Import for Host
    - `playbook_on_play_start`: Play Started
    - `playbook_on_stats`: Playbook Complete
    - `debug`: Debug
    - `verbose`: Verbose
    - `deprecated`: Deprecated
    - `warning`: Warning
    - `system_warning`: System Warning
    - `error`: Error
* `counter`:  (integer)
* `event_display`:  (string)
* `event_data`:  (json)
* `event_level`:  (integer)
* `failed`:  (boolean)
* `changed`:  (boolean)
* `uuid`:  (string)
* `parent_uuid`:  (string)
* `host`:  (id)
* `host_name`:  (string)
* `playbook`:  (string)
* `play`:  (string)
* `task`:  (string)
* `role`:  (string)
* `stdout`:  (string)
* `start_line`:  (integer)
* `end_line`:  (integer)
* `verbosity`:  (integer)



## Sorting

To specify that job events are returned in a particular
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
 * @return ApiJobEventsJobEventsChildrenListRequest
 */
func (a *JobEventsApiService) JobEventsJobEventsChildrenList(ctx _context.Context, id string) ApiJobEventsJobEventsChildrenListRequest {
	return ApiJobEventsJobEventsChildrenListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *JobEventsApiService) JobEventsJobEventsChildrenListExecute(r ApiJobEventsJobEventsChildrenListRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEventsApiService.JobEventsJobEventsChildrenList")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/job_events/{id}/children/"
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

type ApiJobEventsJobEventsHostsListRequest struct {
	ctx _context.Context
	ApiService *JobEventsApiService
	id string
	page *int32
	pageSize *int32
	search *string
}

func (r ApiJobEventsJobEventsHostsListRequest) Page(page int32) ApiJobEventsJobEventsHostsListRequest {
	r.page = &page
	return r
}
func (r ApiJobEventsJobEventsHostsListRequest) PageSize(pageSize int32) ApiJobEventsJobEventsHostsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiJobEventsJobEventsHostsListRequest) Search(search string) ApiJobEventsJobEventsHostsListRequest {
	r.search = &search
	return r
}

func (r ApiJobEventsJobEventsHostsListRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.JobEventsJobEventsHostsListExecute(r)
}

/*
 * JobEventsJobEventsHostsList  List Hosts for a Job Event
 * 
Make a GET request to this resource to retrieve a list of
hosts associated with the selected
job event.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of hosts
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more host records.  

## Results

Each host data structure includes the following fields:

* `id`: Database ID for this host. (integer)
* `type`: Data type for this host. (choice)
* `url`: URL for this host. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this host was created. (datetime)
* `modified`: Timestamp when this host was last modified. (datetime)
* `name`: Name of this host. (string)
* `description`: Optional description of this host. (string)
* `inventory`:  (id)
* `enabled`: Is this host online and available for running jobs? (boolean)
* `instance_id`: The value used by the remote inventory source to uniquely identify the host (string)
* `variables`: Host variables in JSON or YAML format. (json)
* `has_active_failures`:  (field)
* `has_inventory_sources`:  (field)
* `last_job`:  (id)
* `last_job_host_summary`:  (id)
* `insights_system_id`: Red Hat Insights host unique identifier. (string)
* `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)



## Sorting

To specify that hosts are returned in a particular
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
 * @return ApiJobEventsJobEventsHostsListRequest
 */
func (a *JobEventsApiService) JobEventsJobEventsHostsList(ctx _context.Context, id string) ApiJobEventsJobEventsHostsListRequest {
	return ApiJobEventsJobEventsHostsListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *JobEventsApiService) JobEventsJobEventsHostsListExecute(r ApiJobEventsJobEventsHostsListRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEventsApiService.JobEventsJobEventsHostsList")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/job_events/{id}/hosts/"
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

type ApiJobEventsJobEventsListRequest struct {
	ctx _context.Context
	ApiService *JobEventsApiService
	page *int32
	pageSize *int32
	search *string
}

func (r ApiJobEventsJobEventsListRequest) Page(page int32) ApiJobEventsJobEventsListRequest {
	r.page = &page
	return r
}
func (r ApiJobEventsJobEventsListRequest) PageSize(pageSize int32) ApiJobEventsJobEventsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiJobEventsJobEventsListRequest) Search(search string) ApiJobEventsJobEventsListRequest {
	r.search = &search
	return r
}

func (r ApiJobEventsJobEventsListRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.JobEventsJobEventsListExecute(r)
}

/*
 * JobEventsJobEventsList  List Job Events
 * 
Make a GET request to this resource to retrieve the list of
job events.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of job events
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more job event records.  

## Results

Each job event data structure includes the following fields:

* `id`: Database ID for this job event. (integer)
* `type`: Data type for this job event. (choice)
* `url`: URL for this job event. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this job event was created. (datetime)
* `modified`: Timestamp when this job event was last modified. (datetime)
* `job`:  (id)
* `event`:  (choice)
    - `runner_on_failed`: Host Failed
    - `runner_on_start`: Host Started
    - `runner_on_ok`: Host OK
    - `runner_on_error`: Host Failure
    - `runner_on_skipped`: Host Skipped
    - `runner_on_unreachable`: Host Unreachable
    - `runner_on_no_hosts`: No Hosts Remaining
    - `runner_on_async_poll`: Host Polling
    - `runner_on_async_ok`: Host Async OK
    - `runner_on_async_failed`: Host Async Failure
    - `runner_item_on_ok`: Item OK
    - `runner_item_on_failed`: Item Failed
    - `runner_item_on_skipped`: Item Skipped
    - `runner_retry`: Host Retry
    - `runner_on_file_diff`: File Difference
    - `playbook_on_start`: Playbook Started
    - `playbook_on_notify`: Running Handlers
    - `playbook_on_include`: Including File
    - `playbook_on_no_hosts_matched`: No Hosts Matched
    - `playbook_on_no_hosts_remaining`: No Hosts Remaining
    - `playbook_on_task_start`: Task Started
    - `playbook_on_vars_prompt`: Variables Prompted
    - `playbook_on_setup`: Gathering Facts
    - `playbook_on_import_for_host`: internal: on Import for Host
    - `playbook_on_not_import_for_host`: internal: on Not Import for Host
    - `playbook_on_play_start`: Play Started
    - `playbook_on_stats`: Playbook Complete
    - `debug`: Debug
    - `verbose`: Verbose
    - `deprecated`: Deprecated
    - `warning`: Warning
    - `system_warning`: System Warning
    - `error`: Error
* `counter`:  (integer)
* `event_display`:  (string)
* `event_data`:  (json)
* `event_level`:  (integer)
* `failed`:  (boolean)
* `changed`:  (boolean)
* `uuid`:  (string)
* `parent_uuid`:  (string)
* `host`:  (id)
* `host_name`:  (string)
* `playbook`:  (string)
* `play`:  (string)
* `task`:  (string)
* `role`:  (string)
* `stdout`:  (string)
* `start_line`:  (integer)
* `end_line`:  (integer)
* `verbosity`:  (integer)



## Sorting

To specify that job events are returned in a particular
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
 * @return ApiJobEventsJobEventsListRequest
 */
func (a *JobEventsApiService) JobEventsJobEventsList(ctx _context.Context) ApiJobEventsJobEventsListRequest {
	return ApiJobEventsJobEventsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *JobEventsApiService) JobEventsJobEventsListExecute(r ApiJobEventsJobEventsListRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEventsApiService.JobEventsJobEventsList")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/job_events/"

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

type ApiJobEventsJobEventsReadRequest struct {
	ctx _context.Context
	ApiService *JobEventsApiService
	id string
	search *string
}

func (r ApiJobEventsJobEventsReadRequest) Search(search string) ApiJobEventsJobEventsReadRequest {
	r.search = &search
	return r
}

func (r ApiJobEventsJobEventsReadRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.JobEventsJobEventsReadExecute(r)
}

/*
 * JobEventsJobEventsRead  Retrieve a Job Event
 * 
Make GET request to this resource to retrieve a single job event
record containing the following fields:

* `id`: Database ID for this job event. (integer)
* `type`: Data type for this job event. (choice)
* `url`: URL for this job event. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this job event was created. (datetime)
* `modified`: Timestamp when this job event was last modified. (datetime)
* `job`:  (id)
* `event`:  (choice)
    - `runner_on_failed`: Host Failed
    - `runner_on_start`: Host Started
    - `runner_on_ok`: Host OK
    - `runner_on_error`: Host Failure
    - `runner_on_skipped`: Host Skipped
    - `runner_on_unreachable`: Host Unreachable
    - `runner_on_no_hosts`: No Hosts Remaining
    - `runner_on_async_poll`: Host Polling
    - `runner_on_async_ok`: Host Async OK
    - `runner_on_async_failed`: Host Async Failure
    - `runner_item_on_ok`: Item OK
    - `runner_item_on_failed`: Item Failed
    - `runner_item_on_skipped`: Item Skipped
    - `runner_retry`: Host Retry
    - `runner_on_file_diff`: File Difference
    - `playbook_on_start`: Playbook Started
    - `playbook_on_notify`: Running Handlers
    - `playbook_on_include`: Including File
    - `playbook_on_no_hosts_matched`: No Hosts Matched
    - `playbook_on_no_hosts_remaining`: No Hosts Remaining
    - `playbook_on_task_start`: Task Started
    - `playbook_on_vars_prompt`: Variables Prompted
    - `playbook_on_setup`: Gathering Facts
    - `playbook_on_import_for_host`: internal: on Import for Host
    - `playbook_on_not_import_for_host`: internal: on Not Import for Host
    - `playbook_on_play_start`: Play Started
    - `playbook_on_stats`: Playbook Complete
    - `debug`: Debug
    - `verbose`: Verbose
    - `deprecated`: Deprecated
    - `warning`: Warning
    - `system_warning`: System Warning
    - `error`: Error
* `counter`:  (integer)
* `event_display`:  (string)
* `event_data`:  (json)
* `event_level`:  (integer)
* `failed`:  (boolean)
* `changed`:  (boolean)
* `uuid`:  (string)
* `parent_uuid`:  (string)
* `host`:  (id)
* `host_name`:  (string)
* `playbook`:  (string)
* `play`:  (string)
* `task`:  (string)
* `role`:  (string)
* `stdout`:  (string)
* `start_line`:  (integer)
* `end_line`:  (integer)
* `verbosity`:  (integer)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiJobEventsJobEventsReadRequest
 */
func (a *JobEventsApiService) JobEventsJobEventsRead(ctx _context.Context, id string) ApiJobEventsJobEventsReadRequest {
	return ApiJobEventsJobEventsReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *JobEventsApiService) JobEventsJobEventsReadExecute(r ApiJobEventsJobEventsReadRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEventsApiService.JobEventsJobEventsRead")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/job_events/{id}/"
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
