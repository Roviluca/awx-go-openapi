# \CredentialInputSourcesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CredentialInputSourcesCredentialInputSourcesCreate**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesCreate) | **Post** /api/v2/credential_input_sources/ |  Create a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesDelete**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesDelete) | **Delete** /api/v2/credential_input_sources/{id}/ |  Delete a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesList**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesList) | **Get** /api/v2/credential_input_sources/ |  List Credential Input Sources
[**CredentialInputSourcesCredentialInputSourcesPartialUpdate**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesPartialUpdate) | **Patch** /api/v2/credential_input_sources/{id}/ |  Update a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesRead**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesRead) | **Get** /api/v2/credential_input_sources/{id}/ |  Retrieve a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesUpdate**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesUpdate) | **Put** /api/v2/credential_input_sources/{id}/ |  Update a Credential Input Source



## CredentialInputSourcesCredentialInputSourcesCreate

> CredentialInputSourcesCredentialInputSourcesCreate(ctx).Data(data).Execute()

 Create a Credential Input Source



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialInputSourcesCredentialInputSourcesCreateRequest struct via the builder pattern


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


## CredentialInputSourcesCredentialInputSourcesDelete

> CredentialInputSourcesCredentialInputSourcesDelete(ctx, id).Search(search).Execute()

 Delete a Credential Input Source



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
    resp, r, err := api_client.CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCredentialInputSourcesCredentialInputSourcesDeleteRequest struct via the builder pattern


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


## CredentialInputSourcesCredentialInputSourcesList

> CredentialInputSourcesCredentialInputSourcesList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Credential Input Sources



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
    resp, r, err := api_client.CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesList(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialInputSourcesCredentialInputSourcesListRequest struct via the builder pattern


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


## CredentialInputSourcesCredentialInputSourcesPartialUpdate

> CredentialInputSourcesCredentialInputSourcesPartialUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Credential Input Source



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
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesPartialUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCredentialInputSourcesCredentialInputSourcesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
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


## CredentialInputSourcesCredentialInputSourcesRead

> CredentialInputSourcesCredentialInputSourcesRead(ctx, id).Search(search).Execute()

 Retrieve a Credential Input Source



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
    resp, r, err := api_client.CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCredentialInputSourcesCredentialInputSourcesReadRequest struct via the builder pattern


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


## CredentialInputSourcesCredentialInputSourcesUpdate

> CredentialInputSourcesCredentialInputSourcesUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Credential Input Source



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
    data := *openapiclient.Newinline_object_1("InputFieldName_example", 123, 123) // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialInputSourcesApi.CredentialInputSourcesCredentialInputSourcesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCredentialInputSourcesCredentialInputSourcesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject1**](InlineObject1.md) |  | 

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

