# \WorkflowJobTemplateNodesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/always_nodes/ |  Create a Workflow Job Template Node for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList) | **Get** /api/v2/workflow_job_template_nodes/{id}/always_nodes/ |  List Workflow Job Template Nodes for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/ |  Create a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/create_approval_template/ |  Retrieve a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead) | **Get** /api/v2/workflow_job_template_nodes/{id}/create_approval_template/ |  Retrieve a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/credentials/ |  Create a Credential for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList) | **Get** /api/v2/workflow_job_template_nodes/{id}/credentials/ |  List Credentials for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete) | **Delete** /api/v2/workflow_job_template_nodes/{id}/ |  Delete a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/failure_nodes/ |  Create a Workflow Job Template Node for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList) | **Get** /api/v2/workflow_job_template_nodes/{id}/failure_nodes/ |  List Workflow Job Template Nodes for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesList) | **Get** /api/v2/workflow_job_template_nodes/ |  List Workflow Job Template Nodes
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate) | **Patch** /api/v2/workflow_job_template_nodes/{id}/ |  Update a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead) | **Get** /api/v2/workflow_job_template_nodes/{id}/ |  Retrieve a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/success_nodes/ |  Create a Workflow Job Template Node for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList) | **Get** /api/v2/workflow_job_template_nodes/{id}/success_nodes/ |  List Workflow Job Template Nodes for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate) | **Put** /api/v2/workflow_job_template_nodes/{id}/ |  Update a Workflow Job Template Node



## WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate(ctx, id).Data(data).Execute()

 Create a Workflow Job Template Node for a Workflow Job Template Node



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **map[string]interface{}** |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Workflow Job Template Nodes for a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListRequest struct via the builder pattern


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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate(ctx).Data(data).Execute()

 Create a Workflow Job Template Node



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
    data := openapiclient.inline_object_75{AllParentsMustConverge: false, DiffMode: false, ExtraData: "ExtraData_example", Identifier: "Identifier_example", Inventory: 123, JobTags: "JobTags_example", JobType: "JobType_example", Limit: "Limit_example", ScmBranch: "ScmBranch_example", SkipTags: "SkipTags_example", UnifiedJobTemplate: 123, Verbosity: "Verbosity_example", WorkflowJobTemplate: "WorkflowJobTemplate_example"} // InlineObject75 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate(context.Background(), ).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**InlineObject75**](InlineObject75.md) |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate(ctx, id).Data(data).Execute()

 Retrieve a Workflow Job Template Node



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **map[string]interface{}** |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead(ctx, id).Search(search).Execute()

 Retrieve a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate(ctx, id).Data(data).Execute()

 Create a Credential for a Workflow Job Template Node



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **map[string]interface{}** |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Credentials for a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsListRequest struct via the builder pattern


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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete(ctx, id).Search(search).Execute()

 Delete a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesDeleteRequest struct via the builder pattern


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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate(ctx, id).Data(data).Execute()

 Create a Workflow Job Template Node for a Workflow Job Template Node



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **map[string]interface{}** |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Workflow Job Template Nodes for a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListRequest struct via the builder pattern


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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesList

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Workflow Job Template Nodes



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesList(context.Background(), ).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesListRequest struct via the builder pattern


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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Workflow Job Template Node



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
    data := openapiclient.inline_object_77{AllParentsMustConverge: false, DiffMode: false, ExtraData: "ExtraData_example", Identifier: "Identifier_example", Inventory: 123, JobTags: "JobTags_example", JobType: "JobType_example", Limit: "Limit_example", ScmBranch: "ScmBranch_example", SkipTags: "SkipTags_example", UnifiedJobTemplate: 123, Verbosity: "Verbosity_example", WorkflowJobTemplate: "WorkflowJobTemplate_example"} // InlineObject77 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject77**](InlineObject77.md) |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead(ctx, id).Search(search).Execute()

 Retrieve a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate(ctx, id).Data(data).Execute()

 Create a Workflow Job Template Node for a Workflow Job Template Node



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **map[string]interface{}** |  | 

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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Workflow Job Template Nodes for a Workflow Job Template Node



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
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesListRequest struct via the builder pattern


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


## WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate

> WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Workflow Job Template Node



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
    data := openapiclient.inline_object_76{AllParentsMustConverge: false, DiffMode: false, ExtraData: "ExtraData_example", Identifier: "Identifier_example", Inventory: 123, JobTags: "JobTags_example", JobType: "JobType_example", Limit: "Limit_example", ScmBranch: "ScmBranch_example", SkipTags: "SkipTags_example", UnifiedJobTemplate: 123, Verbosity: "Verbosity_example", WorkflowJobTemplate: "WorkflowJobTemplate_example"} // InlineObject76 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobTemplateNodesApi.WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject76**](InlineObject76.md) |  | 

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

