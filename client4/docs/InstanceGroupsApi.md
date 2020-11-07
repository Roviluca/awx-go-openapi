# \InstanceGroupsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceGroupsInstanceGroupsCreate**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsCreate) | **Post** /api/v2/instance_groups/ |  Create an Instance Group
[**InstanceGroupsInstanceGroupsDelete**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsDelete) | **Delete** /api/v2/instance_groups/{id}/ |  Delete an Instance Group
[**InstanceGroupsInstanceGroupsInstancesCreate**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsInstancesCreate) | **Post** /api/v2/instance_groups/{id}/instances/ |  Create an Instance for an Instance Group
[**InstanceGroupsInstanceGroupsInstancesList**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsInstancesList) | **Get** /api/v2/instance_groups/{id}/instances/ |  List Instances for an Instance Group
[**InstanceGroupsInstanceGroupsJobsList**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsJobsList) | **Get** /api/v2/instance_groups/{id}/jobs/ |  List Unified Jobs for an Instance Group
[**InstanceGroupsInstanceGroupsList**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsList) | **Get** /api/v2/instance_groups/ |  List Instance Groups
[**InstanceGroupsInstanceGroupsPartialUpdate**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsPartialUpdate) | **Patch** /api/v2/instance_groups/{id}/ |  Update an Instance Group
[**InstanceGroupsInstanceGroupsRead**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsRead) | **Get** /api/v2/instance_groups/{id}/ |  Retrieve an Instance Group
[**InstanceGroupsInstanceGroupsUpdate**](InstanceGroupsApi.md#InstanceGroupsInstanceGroupsUpdate) | **Put** /api/v2/instance_groups/{id}/ |  Update an Instance Group



## InstanceGroupsInstanceGroupsCreate

> InstanceGroupsInstanceGroupsCreate(ctx, optional)

 Create an Instance Group

 Make a POST request to this resource with the following instance group fields to create a new instance group:       * `name`: Name of this instance group. (string, required)              * `credential`:  (id, default=``) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer, default=`0`) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer, default=`0`) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json, default=``) * `pod_spec_override`:  (string, default=`\"\"`)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceGroupsInstanceGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of InlineObject15**](InlineObject15.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsDelete

> InstanceGroupsInstanceGroupsDelete(ctx, id, optional)

 Delete an Instance Group

 Make a DELETE request to this resource to delete this instance group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsInstancesCreate

> InstanceGroupsInstanceGroupsInstancesCreate(ctx, id, optional)

 Create an Instance for an Instance Group

 Make a POST request to this resource with the following instance fields to create a new instance associated with this instance group.           * `capacity_adjustment`:  (decimal, default=`1`)           * `enabled`:  (boolean, default=`True`) * `managed_by_policy`:  (boolean, default=`True`)         # Add Instances for an Instance Group:  Make a POST request to this resource with only an `id` field to associate an existing instance with this instance group.  # Remove Instances from this Instance Group:  Make a POST request to this resource with `id` and `disassociate` fields to remove the instance from this instance group  without deleting the instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsInstancesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsInstancesCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of InlineObject17**](InlineObject17.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsInstancesList

> InstanceGroupsInstanceGroupsInstancesList(ctx, id, optional)

 List Instances for an Instance Group

 Make a GET request to this resource to retrieve a list of instances associated with the selected instance group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of instances found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more instance records.    ## Results  Each instance data structure includes the following fields:  * `id`: Database ID for this instance. (integer) * `type`: Data type for this instance. (choice) * `url`: URL for this instance. (string) * `related`: Data structure with URLs of related resources. (object) * `uuid`:  (string) * `hostname`:  (string) * `created`: Timestamp when this instance was created. (datetime) * `modified`: Timestamp when this instance was last modified. (datetime) * `capacity_adjustment`:  (decimal) * `version`:  (string) * `capacity`:  (integer) * `consumed_capacity`:  (field) * `percent_capacity_remaining`:  (field) * `jobs_running`: Count of jobs in the running or waiting state that are targeted for this instance (integer) * `jobs_total`: Count of all jobs that target this instance (integer) * `cpu`:  (integer) * `memory`:  (integer) * `cpu_capacity`:  (integer) * `mem_capacity`:  (integer) * `enabled`:  (boolean) * `managed_by_policy`:  (boolean)    ## Sorting  To specify that instances are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsInstancesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsInstancesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsJobsList

> InstanceGroupsInstanceGroupsJobsList(ctx, id, optional)

 List Unified Jobs for an Instance Group

 Make a GET request to this resource to retrieve a list of unified jobs associated with the selected instance group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of unified jobs found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more unified job records.    ## Results  Each unified job data structure includes the following fields:  * `id`: Database ID for this unified job. (integer) * `type`: Data type for this unified job. (choice) * `url`: URL for this unified job. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this unified job was created. (datetime) * `modified`: Timestamp when this unified job was last modified. (datetime) * `name`: Name of this unified job. (string) * `description`: Optional description of this unified job. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `controller_node`: The instance that managed the isolated execution environment. (string)    ## Sorting  To specify that unified jobs are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsJobsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsJobsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsList

> InstanceGroupsInstanceGroupsList(ctx, optional)

 List Instance Groups

 Make a GET request to this resource to retrieve the list of instance groups.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of instance groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more instance group records.    ## Results  Each instance group data structure includes the following fields:  * `id`: Database ID for this instance group. (integer) * `type`: Data type for this instance group. (choice) * `url`: URL for this instance group. (string) * `related`: Data structure with URLs of related resources. (object) * `name`: Name of this instance group. (string) * `created`: Timestamp when this instance group was created. (datetime) * `modified`: Timestamp when this instance group was last modified. (datetime) * `capacity`:  (field) * `committed_capacity`:  (field) * `consumed_capacity`:  (field) * `percent_capacity_remaining`:  (field) * `jobs_running`: Count of jobs in the running or waiting state that are targeted for this instance group (integer) * `jobs_total`: Count of all jobs that target this instance group (integer) * `instances`:  (field) * `controller`: Instance Group to remotely control this group. (id) * `is_controller`: Indicates whether instance group controls any other group (boolean) * `is_isolated`: Indicates whether instances in this group are isolated.Isolated groups have a designated controller group. (boolean) * `is_containerized`: Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster. (boolean) * `credential`:  (id) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json) * `pod_spec_override`:  (string) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)    ## Sorting  To specify that instance groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceGroupsInstanceGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsPartialUpdate

> InstanceGroupsInstanceGroupsPartialUpdate(ctx, id, optional)

 Update an Instance Group

 Make a PUT or PATCH request to this resource to update this instance group.  The following fields may be modified:       * `name`: Name of this instance group. (string, required)              * `credential`:  (id, default=``) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer, default=`0`) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer, default=`0`) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json, default=``) * `pod_spec_override`:  (string, default=`\"\"`)          For a PATCH request, include only the fields that are being modified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | **optional.Map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsRead

> InstanceGroupsInstanceGroupsRead(ctx, id, optional)

 Retrieve an Instance Group

 Make GET request to this resource to retrieve a single instance group record containing the following fields:  * `id`: Database ID for this instance group. (integer) * `type`: Data type for this instance group. (choice) * `url`: URL for this instance group. (string) * `related`: Data structure with URLs of related resources. (object) * `name`: Name of this instance group. (string) * `created`: Timestamp when this instance group was created. (datetime) * `modified`: Timestamp when this instance group was last modified. (datetime) * `capacity`:  (field) * `committed_capacity`:  (field) * `consumed_capacity`:  (field) * `percent_capacity_remaining`:  (field) * `jobs_running`: Count of jobs in the running or waiting state that are targeted for this instance group (integer) * `jobs_total`: Count of all jobs that target this instance group (integer) * `instances`:  (field) * `controller`: Instance Group to remotely control this group. (id) * `is_controller`: Indicates whether instance group controls any other group (boolean) * `is_isolated`: Indicates whether instances in this group are isolated.Isolated groups have a designated controller group. (boolean) * `is_containerized`: Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster. (boolean) * `credential`:  (id) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json) * `pod_spec_override`:  (string) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceGroupsInstanceGroupsUpdate

> InstanceGroupsInstanceGroupsUpdate(ctx, id, optional)

 Update an Instance Group

 Make a PUT or PATCH request to this resource to update this instance group.  The following fields may be modified:       * `name`: Name of this instance group. (string, required)              * `credential`:  (id, default=``) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer, default=`0`) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer, default=`0`) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json, default=``) * `pod_spec_override`:  (string, default=`\"\"`)        For a PUT request, include **all** fields in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***InstanceGroupsInstanceGroupsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstanceGroupsInstanceGroupsUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of InlineObject16**](InlineObject16.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

