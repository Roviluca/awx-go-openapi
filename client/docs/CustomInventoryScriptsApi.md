# \CustomInventoryScriptsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomInventoryScriptsInventoryScriptsCopyCreate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsCopyCreate) | **Post** /api/v2/inventory_scripts/{id}/copy/ | 
[**CustomInventoryScriptsInventoryScriptsCopyList**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsCopyList) | **Get** /api/v2/inventory_scripts/{id}/copy/ | 
[**CustomInventoryScriptsInventoryScriptsCreate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsCreate) | **Post** /api/v2/inventory_scripts/ |  Create a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsDelete**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsDelete) | **Delete** /api/v2/inventory_scripts/{id}/ |  Delete a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsList**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsList) | **Get** /api/v2/inventory_scripts/ |  List Custom Inventory Scripts
[**CustomInventoryScriptsInventoryScriptsObjectRolesList**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsObjectRolesList) | **Get** /api/v2/inventory_scripts/{id}/object_roles/ |  List Roles for a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsPartialUpdate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsPartialUpdate) | **Patch** /api/v2/inventory_scripts/{id}/ |  Update a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsRead**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsRead) | **Get** /api/v2/inventory_scripts/{id}/ |  Retrieve a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsUpdate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsUpdate) | **Put** /api/v2/inventory_scripts/{id}/ |  Update a Custom Inventory Script



## CustomInventoryScriptsInventoryScriptsCopyCreate

> CustomInventoryScriptsInventoryScriptsCopyCreate(ctx, id).Data(data).Execute()



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
    data := openapiclient.inline_object_26{Name: "Name_example"} // InlineObject26 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsCopyCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsCopyCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsCopyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject26**](InlineObject26.md) |  | 

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


## CustomInventoryScriptsInventoryScriptsCopyList

> CustomInventoryScriptsInventoryScriptsCopyList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()



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
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsCopyList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsCopyList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsCopyListRequest struct via the builder pattern


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


## CustomInventoryScriptsInventoryScriptsCreate

> CustomInventoryScriptsInventoryScriptsCreate(ctx).Data(data).Execute()

 Create a Custom Inventory Script



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
    data := openapiclient.inline_object_23{Description: "Description_example", Name: "Name_example", Organization: 123, Script: "Script_example"} // InlineObject23 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsCreate(context.Background(), ).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**InlineObject23**](InlineObject23.md) |  | 

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


## CustomInventoryScriptsInventoryScriptsDelete

> CustomInventoryScriptsInventoryScriptsDelete(ctx, id).Search(search).Execute()

 Delete a Custom Inventory Script



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
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsDeleteRequest struct via the builder pattern


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


## CustomInventoryScriptsInventoryScriptsList

> CustomInventoryScriptsInventoryScriptsList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Custom Inventory Scripts



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
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsList(context.Background(), ).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsListRequest struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomInventoryScriptsInventoryScriptsObjectRolesList

> CustomInventoryScriptsInventoryScriptsObjectRolesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Roles for a Custom Inventory Script



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
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsObjectRolesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsObjectRolesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsObjectRolesListRequest struct via the builder pattern


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


## CustomInventoryScriptsInventoryScriptsPartialUpdate

> CustomInventoryScriptsInventoryScriptsPartialUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Custom Inventory Script



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
    data := openapiclient.inline_object_25{Description: "Description_example", Name: "Name_example", Organization: 123, Script: "Script_example"} // InlineObject25 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsPartialUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject25**](InlineObject25.md) |  | 

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


## CustomInventoryScriptsInventoryScriptsRead

> CustomInventoryScriptsInventoryScriptsRead(ctx, id).Search(search).Execute()

 Retrieve a Custom Inventory Script



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
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsReadRequest struct via the builder pattern


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


## CustomInventoryScriptsInventoryScriptsUpdate

> CustomInventoryScriptsInventoryScriptsUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Custom Inventory Script



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
    data := openapiclient.inline_object_24{Description: "Description_example", Name: "Name_example", Organization: 123, Script: "Script_example"} // InlineObject24 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomInventoryScriptsApi.CustomInventoryScriptsInventoryScriptsUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomInventoryScriptsInventoryScriptsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject24**](InlineObject24.md) |  | 

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

