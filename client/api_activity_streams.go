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

// ActivityStreamsApiService ActivityStreamsApi service
type ActivityStreamsApiService service

// ActivityStreamsActivityStreamListOpts Optional parameters for the method 'ActivityStreamsActivityStreamList'
type ActivityStreamsActivityStreamListOpts struct {
	Page     optional.Int32
	PageSize optional.Int32
	Search   optional.String
}

/*
ActivityStreamsActivityStreamList  List Activity Streams
 Make a GET request to this resource to retrieve the list of activity streams.  The resulting data structure contains:      {         \&quot;count\&quot;: 99,         \&quot;next\&quot;: null,         \&quot;previous\&quot;: null,         \&quot;results\&quot;: [             ...         ]     }  The &#x60;count&#x60; field indicates the total number of activity streams found for the given query.  The &#x60;next&#x60; and &#x60;previous&#x60; fields provides links to additional results if there are more than will fit on a single page.  The &#x60;results&#x60; list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * &#x60;id&#x60;: Database ID for this activity stream. (integer) * &#x60;type&#x60;: Data type for this activity stream. (choice) * &#x60;url&#x60;: URL for this activity stream. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;timestamp&#x60;:  (datetime) * &#x60;operation&#x60;: The action taken with respect to the given object(s). (choice)     - &#x60;create&#x60;: Entity Created     - &#x60;update&#x60;: Entity Updated     - &#x60;delete&#x60;: Entity Deleted     - &#x60;associate&#x60;: Entity Associated with another Entity     - &#x60;disassociate&#x60;: Entity was Disassociated with another Entity * &#x60;changes&#x60;: A summary of the new and changed values when an object is created, updated, or deleted (json) * &#x60;object1&#x60;: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * &#x60;object2&#x60;: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * &#x60;object_association&#x60;: When present, shows the field name of the role or relationship that changed. (field) * &#x60;action_node&#x60;: The cluster node the activity took place on. (string) * &#x60;object_type&#x60;: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the &#x60;order_by&#x60; query string parameter on the GET request.      ?order_by&#x3D;  Prefix the field name with a dash &#x60;-&#x60; to sort in reverse:      ?order_by&#x3D;-  Multiple sorting fields may be specified by separating the field names with a comma &#x60;,&#x60;:      ?order_by&#x3D;,some_other_field  ## Pagination  Use the &#x60;page_size&#x60; query string parameter to change the number of results returned for each request.  Use the &#x60;page&#x60; query string parameter to retrieve a particular page of results.      ?page_size&#x3D;100&amp;page&#x3D;2  The &#x60;previous&#x60; and &#x60;next&#x60; links returned with the results will set these query string parameters automatically.  ## Searching  Use the &#x60;search&#x60; query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search&#x3D;findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search&#x3D;findme
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ActivityStreamsActivityStreamListOpts - Optional Parameters:
 * @param "Page" (optional.Int32) -  A page number within the paginated result set.
 * @param "PageSize" (optional.Int32) -  Number of results to return per page.
 * @param "Search" (optional.String) -  A search term.
*/
func (a *ActivityStreamsApiService) ActivityStreamsActivityStreamList(ctx _context.Context, localVarOptionals *ActivityStreamsActivityStreamListOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/activity_stream/"
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

// ActivityStreamsActivityStreamReadOpts Optional parameters for the method 'ActivityStreamsActivityStreamRead'
type ActivityStreamsActivityStreamReadOpts struct {
	Search optional.String
}

/*
ActivityStreamsActivityStreamRead  Retrieve an Activity Stream
 Make GET request to this resource to retrieve a single activity stream record containing the following fields:  * &#x60;id&#x60;: Database ID for this activity stream. (integer) * &#x60;type&#x60;: Data type for this activity stream. (choice) * &#x60;url&#x60;: URL for this activity stream. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;timestamp&#x60;:  (datetime) * &#x60;operation&#x60;: The action taken with respect to the given object(s). (choice)     - &#x60;create&#x60;: Entity Created     - &#x60;update&#x60;: Entity Updated     - &#x60;delete&#x60;: Entity Deleted     - &#x60;associate&#x60;: Entity Associated with another Entity     - &#x60;disassociate&#x60;: Entity was Disassociated with another Entity * &#x60;changes&#x60;: A summary of the new and changed values when an object is created, updated, or deleted (json) * &#x60;object1&#x60;: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * &#x60;object2&#x60;: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * &#x60;object_association&#x60;: When present, shows the field name of the role or relationship that changed. (field) * &#x60;action_node&#x60;: The cluster node the activity took place on. (string) * &#x60;object_type&#x60;: When present, shows the model on which the role or relationship was defined. (field)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *ActivityStreamsActivityStreamReadOpts - Optional Parameters:
 * @param "Search" (optional.String) -  A search term.
*/
func (a *ActivityStreamsApiService) ActivityStreamsActivityStreamRead(ctx _context.Context, id string, localVarOptionals *ActivityStreamsActivityStreamReadOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/activity_stream/{id}/"
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
