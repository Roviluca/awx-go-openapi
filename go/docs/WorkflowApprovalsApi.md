# \WorkflowApprovalsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowApprovalsWorkflowApprovalsApproveCreate**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsApproveCreate) | **Post** /api/v2/workflow_approvals/{id}/approve/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsApproveRead**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsApproveRead) | **Get** /api/v2/workflow_approvals/{id}/approve/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsCreate**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsCreate) | **Post** /api/v2/workflow_approvals/ |  Create a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsDelete**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsDelete) | **Delete** /api/v2/workflow_approvals/{id}/ |  Delete a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsDenyCreate**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsDenyCreate) | **Post** /api/v2/workflow_approvals/{id}/deny/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsDenyRead**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsDenyRead) | **Get** /api/v2/workflow_approvals/{id}/deny/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsList**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsList) | **Get** /api/v2/workflow_approvals/ |  List Workflow Approvals
[**WorkflowApprovalsWorkflowApprovalsRead**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsRead) | **Get** /api/v2/workflow_approvals/{id}/ |  Retrieve a Workflow Approval



## WorkflowApprovalsWorkflowApprovalsApproveCreate

> WorkflowApprovalsWorkflowApprovalsApproveCreate(ctx, id).Execute()

 Retrieve a Workflow Approval



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsApproveCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsApproveCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsApproveCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## WorkflowApprovalsWorkflowApprovalsApproveRead

> WorkflowApprovalsWorkflowApprovalsApproveRead(ctx, id).Search(search).Execute()

 Retrieve a Workflow Approval



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
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsApproveRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsApproveRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsApproveReadRequest struct via the builder pattern


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


## WorkflowApprovalsWorkflowApprovalsCreate

> WorkflowApprovalsWorkflowApprovalsCreate(ctx).Data(data).Execute()

 Create a Workflow Approval



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
    data := *openapiclient.Newinline_object_74("Name_example") // InlineObject74 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**InlineObject74**](InlineObject74.md) |  | 

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


## WorkflowApprovalsWorkflowApprovalsDelete

> WorkflowApprovalsWorkflowApprovalsDelete(ctx, id).Search(search).Execute()

 Delete a Workflow Approval



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
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsDeleteRequest struct via the builder pattern


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


## WorkflowApprovalsWorkflowApprovalsDenyCreate

> WorkflowApprovalsWorkflowApprovalsDenyCreate(ctx, id).Execute()

 Retrieve a Workflow Approval



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsDenyCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsDenyCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsDenyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## WorkflowApprovalsWorkflowApprovalsDenyRead

> WorkflowApprovalsWorkflowApprovalsDenyRead(ctx, id).Search(search).Execute()

 Retrieve a Workflow Approval



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
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsDenyRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsDenyRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsDenyReadRequest struct via the builder pattern


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


## WorkflowApprovalsWorkflowApprovalsList

> WorkflowApprovalsWorkflowApprovalsList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Workflow Approvals



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
    page := 987 // int32 | A page number within the paginated result set. (optional)
    pageSize := 987 // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsList(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsListRequest struct via the builder pattern


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


## WorkflowApprovalsWorkflowApprovalsRead

> WorkflowApprovalsWorkflowApprovalsRead(ctx, id).Search(search).Execute()

 Retrieve a Workflow Approval



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
    resp, r, err := api_client.WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApprovalsApi.WorkflowApprovalsWorkflowApprovalsRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowApprovalsWorkflowApprovalsReadRequest struct via the builder pattern


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

