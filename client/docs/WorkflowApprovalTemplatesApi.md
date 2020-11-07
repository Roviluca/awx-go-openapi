# \WorkflowApprovalTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList) | **Get** /api/v2/workflow_approval_templates/{id}/approvals/ |  List Workflow Approvals for a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete) | **Delete** /api/v2/workflow_approval_templates/{id}/ |  Delete a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate) | **Patch** /api/v2/workflow_approval_templates/{id}/ |  Update a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead) | **Get** /api/v2/workflow_approval_templates/{id}/ |  Retrieve a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate) | **Put** /api/v2/workflow_approval_templates/{id}/ |  Update a Workflow Approval Template



## WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList

> WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Workflow Approvals for a Workflow Approval Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    page := 987 // int32 | A page number within the paginated result set. (optional)
    pageSize := 987 // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

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


## WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete

> WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete(ctx, id).Search(search).Execute()

 Delete a Workflow Approval Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 

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


## WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate

> WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Workflow Approval Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    search := "search_example" // string | A search term. (optional)
    data := openapiclient.inline_object_73{Description: "Description_example", Name: "Name_example", Timeout: 123} // InlineObject73 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject73**](InlineObject73.md) |  | 

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


## WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead

> WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead(ctx, id).Search(search).Execute()

 Retrieve a Workflow Approval Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 

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


## WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate

> WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Workflow Approval Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    search := "search_example" // string | A search term. (optional)
    data := openapiclient.inline_object_72{Description: "Description_example", Name: "Name_example", Timeout: 123} // InlineObject72 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalTemplatesApi.WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject72**](InlineObject72.md) |  | 

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

