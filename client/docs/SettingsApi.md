# \SettingsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsSettingsDelete**](SettingsApi.md#SettingsSettingsDelete) | **Delete** /api/v2/settings/{category_slug}/ |  Delete a Setting
[**SettingsSettingsList**](SettingsApi.md#SettingsSettingsList) | **Get** /api/v2/settings/ |  List Settings
[**SettingsSettingsLoggingTestCreate**](SettingsApi.md#SettingsSettingsLoggingTestCreate) | **Post** /api/v2/settings/logging/test/ |  Test Logging Configuration
[**SettingsSettingsPartialUpdate**](SettingsApi.md#SettingsSettingsPartialUpdate) | **Patch** /api/v2/settings/{category_slug}/ |  Update a Setting
[**SettingsSettingsRead**](SettingsApi.md#SettingsSettingsRead) | **Get** /api/v2/settings/{category_slug}/ |  Retrieve a Setting
[**SettingsSettingsUpdate**](SettingsApi.md#SettingsSettingsUpdate) | **Put** /api/v2/settings/{category_slug}/ |  Update a Setting



## SettingsSettingsDelete

> SettingsSettingsDelete(ctx, categorySlug).Execute()

 Delete a Setting



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
    categorySlug := "categorySlug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsDelete(context.Background(), categorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SettingsSettingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSettingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SettingsSettingsList

> SettingsSettingsList(ctx).Page(page).PageSize(pageSize).Execute()

 List Settings



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SettingsSettingsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSettingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

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


## SettingsSettingsLoggingTestCreate

> SettingsSettingsLoggingTestCreate(ctx).Data(data).Execute()

 Test Logging Configuration

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
    data := *openapiclient.Newinline_object_60(false, false, "ALLOW_JINJA_IN_EXTRA_VARS_example", false, "AUTOMATION_ANALYTICS_LAST_GATHER_example", 123, 123, "AWX_PROOT_BASE_PATH_example", false, 123, false, 123, false, false, []string{"PROXY_IP_ALLOWED_LIST_example"), []string{"REMOTE_HOST_HEADERS_example"), 123, 123, 123, 123, "SOCIAL_AUTH_SAML_SP_PRIVATE_KEY_example", "SOCIAL_AUTH_SAML_SP_PUBLIC_CERT_example", 123, 123, 123, "TOWER_URL_BASE_example", false) // InlineObject60 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsLoggingTestCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SettingsSettingsLoggingTestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSettingsLoggingTestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**InlineObject60**](InlineObject60.md) |  | 

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


## SettingsSettingsPartialUpdate

> SettingsSettingsPartialUpdate(ctx, categorySlug).Data(data).Execute()

 Update a Setting



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
    categorySlug := "categorySlug_example" // string | 
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsPartialUpdate(context.Background(), categorySlug).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SettingsSettingsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSettingsPartialUpdateRequest struct via the builder pattern


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


## SettingsSettingsRead

> SettingsSettingsRead(ctx, categorySlug).Execute()

 Retrieve a Setting



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
    categorySlug := "categorySlug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsRead(context.Background(), categorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SettingsSettingsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSettingsReadRequest struct via the builder pattern


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


## SettingsSettingsUpdate

> SettingsSettingsUpdate(ctx, categorySlug).Data(data).Execute()

 Update a Setting



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
    categorySlug := "categorySlug_example" // string | 
    data := 987 // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsUpdate(context.Background(), categorySlug).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SettingsSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSettingsUpdateRequest struct via the builder pattern


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

