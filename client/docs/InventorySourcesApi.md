# \InventorySourcesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventorySourcesInventorySourcesActivityStreamList**](InventorySourcesApi.md#InventorySourcesInventorySourcesActivityStreamList) | **Get** /api/v2/inventory_sources/{id}/activity_stream/ |  List Activity Streams for an Inventory Source
[**InventorySourcesInventorySourcesCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesCreate) | **Post** /api/v2/inventory_sources/ |  Create an Inventory Source
[**InventorySourcesInventorySourcesCredentialsCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesCredentialsCreate) | **Post** /api/v2/inventory_sources/{id}/credentials/ |  Create a Credential for an Inventory Source
[**InventorySourcesInventorySourcesCredentialsList**](InventorySourcesApi.md#InventorySourcesInventorySourcesCredentialsList) | **Get** /api/v2/inventory_sources/{id}/credentials/ |  List Credentials for an Inventory Source
[**InventorySourcesInventorySourcesDelete**](InventorySourcesApi.md#InventorySourcesInventorySourcesDelete) | **Delete** /api/v2/inventory_sources/{id}/ |  Delete an Inventory Source
[**InventorySourcesInventorySourcesGroupsDelete**](InventorySourcesApi.md#InventorySourcesInventorySourcesGroupsDelete) | **Delete** /api/v2/inventory_sources/{id}/groups/ |  Create a Group for an Inventory Source
[**InventorySourcesInventorySourcesGroupsList**](InventorySourcesApi.md#InventorySourcesInventorySourcesGroupsList) | **Get** /api/v2/inventory_sources/{id}/groups/ |  List Groups for an Inventory Source
[**InventorySourcesInventorySourcesHostsDelete**](InventorySourcesApi.md#InventorySourcesInventorySourcesHostsDelete) | **Delete** /api/v2/inventory_sources/{id}/hosts/ |  Create a Host for an Inventory Source
[**InventorySourcesInventorySourcesHostsList**](InventorySourcesApi.md#InventorySourcesInventorySourcesHostsList) | **Get** /api/v2/inventory_sources/{id}/hosts/ |  List Hosts for an Inventory Source
[**InventorySourcesInventorySourcesInventoryUpdatesList**](InventorySourcesApi.md#InventorySourcesInventorySourcesInventoryUpdatesList) | **Get** /api/v2/inventory_sources/{id}/inventory_updates/ |  List Inventory Updates for an Inventory Source
[**InventorySourcesInventorySourcesList**](InventorySourcesApi.md#InventorySourcesInventorySourcesList) | **Get** /api/v2/inventory_sources/ |  List Inventory Sources
[**InventorySourcesInventorySourcesNotificationTemplatesErrorCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesErrorCreate) | **Post** /api/v2/inventory_sources/{id}/notification_templates_error/ |  Create a Notification Template for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesErrorList**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesErrorList) | **Get** /api/v2/inventory_sources/{id}/notification_templates_error/ |  List Notification Templates for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesStartedCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesStartedCreate) | **Post** /api/v2/inventory_sources/{id}/notification_templates_started/ |  Create a Notification Template for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesStartedList**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesStartedList) | **Get** /api/v2/inventory_sources/{id}/notification_templates_started/ |  List Notification Templates for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate) | **Post** /api/v2/inventory_sources/{id}/notification_templates_success/ |  Create a Notification Template for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesSuccessList**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesSuccessList) | **Get** /api/v2/inventory_sources/{id}/notification_templates_success/ |  List Notification Templates for an Inventory Source
[**InventorySourcesInventorySourcesPartialUpdate**](InventorySourcesApi.md#InventorySourcesInventorySourcesPartialUpdate) | **Patch** /api/v2/inventory_sources/{id}/ |  Update an Inventory Source
[**InventorySourcesInventorySourcesRead**](InventorySourcesApi.md#InventorySourcesInventorySourcesRead) | **Get** /api/v2/inventory_sources/{id}/ |  Retrieve an Inventory Source
[**InventorySourcesInventorySourcesSchedulesCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesSchedulesCreate) | **Post** /api/v2/inventory_sources/{id}/schedules/ |  Create a Schedule for an Inventory Source
[**InventorySourcesInventorySourcesSchedulesList**](InventorySourcesApi.md#InventorySourcesInventorySourcesSchedulesList) | **Get** /api/v2/inventory_sources/{id}/schedules/ |  List Schedules for an Inventory Source
[**InventorySourcesInventorySourcesUpdate0**](InventorySourcesApi.md#InventorySourcesInventorySourcesUpdate0) | **Put** /api/v2/inventory_sources/{id}/ |  Update an Inventory Source
[**InventorySourcesInventorySourcesUpdateCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesUpdateCreate) | **Post** /api/v2/inventory_sources/{id}/update/ |  Update Inventory Source
[**InventorySourcesInventorySourcesUpdateRead**](InventorySourcesApi.md#InventorySourcesInventorySourcesUpdateRead) | **Get** /api/v2/inventory_sources/{id}/update/ |  Update Inventory Source



## InventorySourcesInventorySourcesActivityStreamList

> InventorySourcesInventorySourcesActivityStreamList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Activity Streams for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesActivityStreamList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesActivityStreamList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesActivityStreamListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesCreate

> InventorySourcesInventorySourcesCreate(ctx).Data(data).Execute()

 Create an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesCreate(context.Background(), ).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesCreateRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesCredentialsCreate

> InventorySourcesInventorySourcesCredentialsCreate(ctx, id).Data(data).Execute()

 Create a Credential for an Inventory Source



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
    data := openapiclient.inline_object_28{CredentialType: 123, Description: "Description_example", Inputs: 123, Name: "Name_example", Organization: 123} // InlineObject28 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesCredentialsCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesCredentialsCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesCredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject28**](InlineObject28.md) |  | 

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


## InventorySourcesInventorySourcesCredentialsList

> InventorySourcesInventorySourcesCredentialsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Credentials for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesCredentialsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesCredentialsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesCredentialsListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesDelete

> InventorySourcesInventorySourcesDelete(ctx, id).Search(search).Execute()

 Delete an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesDeleteRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesGroupsDelete

> InventorySourcesInventorySourcesGroupsDelete(ctx, id).Search(search).Execute()

 Create a Group for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesGroupsDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesGroupsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesGroupsDeleteRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesGroupsList

> InventorySourcesInventorySourcesGroupsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Groups for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesGroupsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesGroupsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesGroupsListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesHostsDelete

> InventorySourcesInventorySourcesHostsDelete(ctx, id).Search(search).Execute()

 Create a Host for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesHostsDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesHostsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesHostsDeleteRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesHostsList

> InventorySourcesInventorySourcesHostsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Hosts for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesHostsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesHostsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesHostsListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesInventoryUpdatesList

> InventorySourcesInventorySourcesInventoryUpdatesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Inventory Updates for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesInventoryUpdatesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesInventoryUpdatesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesInventoryUpdatesListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesList

> InventorySourcesInventorySourcesList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Inventory Sources



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesList(context.Background(), ).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesNotificationTemplatesErrorCreate

> InventorySourcesInventorySourcesNotificationTemplatesErrorCreate(ctx, id).Data(data).Execute()

 Create a Notification Template for an Inventory Source



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
    data := openapiclient.inline_object_29{Description: "Description_example", Messages: "Messages_example", Name: "Name_example", NotificationConfiguration: "NotificationConfiguration_example", NotificationType: "NotificationType_example", Organization: 123} // InlineObject29 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesErrorCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesErrorCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesNotificationTemplatesErrorCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject29**](InlineObject29.md) |  | 

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


## InventorySourcesInventorySourcesNotificationTemplatesErrorList

> InventorySourcesInventorySourcesNotificationTemplatesErrorList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesErrorList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesErrorList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesNotificationTemplatesErrorListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesNotificationTemplatesStartedCreate

> InventorySourcesInventorySourcesNotificationTemplatesStartedCreate(ctx, id).Data(data).Execute()

 Create a Notification Template for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesStartedCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesStartedCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesNotificationTemplatesStartedCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **map[string]interface{}** |  | 

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


## InventorySourcesInventorySourcesNotificationTemplatesStartedList

> InventorySourcesInventorySourcesNotificationTemplatesStartedList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesStartedList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesStartedList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesNotificationTemplatesStartedListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate

> InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate(ctx, id).Data(data).Execute()

 Create a Notification Template for an Inventory Source



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
    data := openapiclient.inline_object_30{Description: "Description_example", Messages: "Messages_example", Name: "Name_example", NotificationConfiguration: "NotificationConfiguration_example", NotificationType: "NotificationType_example", Organization: 123} // InlineObject30 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesNotificationTemplatesSuccessCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject30**](InlineObject30.md) |  | 

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


## InventorySourcesInventorySourcesNotificationTemplatesSuccessList

> InventorySourcesInventorySourcesNotificationTemplatesSuccessList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesSuccessList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesNotificationTemplatesSuccessList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesNotificationTemplatesSuccessListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesPartialUpdate

> InventorySourcesInventorySourcesPartialUpdate(ctx, id).Search(search).Data(data).Execute()

 Update an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesPartialUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesPartialUpdateRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesRead

> InventorySourcesInventorySourcesRead(ctx, id).Search(search).Execute()

 Retrieve an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesReadRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesSchedulesCreate

> InventorySourcesInventorySourcesSchedulesCreate(ctx, id).Data(data).Execute()

 Create a Schedule for an Inventory Source



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
    data := openapiclient.inline_object_31{Description: "Description_example", DiffMode: false, Enabled: false, ExtraData: "ExtraData_example", Inventory: 123, JobTags: "JobTags_example", JobType: "JobType_example", Limit: "Limit_example", Name: "Name_example", Rrule: "Rrule_example", ScmBranch: "ScmBranch_example", SkipTags: "SkipTags_example", UnifiedJobTemplate: 123, Verbosity: "Verbosity_example"} // InlineObject31 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesSchedulesCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesSchedulesCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesSchedulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject31**](InlineObject31.md) |  | 

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


## InventorySourcesInventorySourcesSchedulesList

> InventorySourcesInventorySourcesSchedulesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Schedules for an Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesSchedulesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesSchedulesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesSchedulesListRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesUpdate0

> InventorySourcesInventorySourcesUpdate0(ctx, id).Search(search).Data(data).Execute()

 Update an Inventory Source



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
    data := openapiclient.inline_object_27{Credential: 123, CustomVirtualenv: "CustomVirtualenv_example", Description: "Description_example", EnabledValue: "EnabledValue_example", EnabledVar: "EnabledVar_example", HostFilter: "HostFilter_example", Inventory: 123, Name: "Name_example", Overwrite: false, OverwriteVars: false, Source: "Source_example", SourcePath: "SourcePath_example", SourceProject: "SourceProject_example", SourceScript: 123, SourceVars: "SourceVars_example", Timeout: 123, UpdateCacheTimeout: 123, UpdateOnLaunch: false, UpdateOnProjectUpdate: false, Verbosity: "Verbosity_example"} // InlineObject27 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesUpdate0(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesUpdate0``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesUpdate0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject27**](InlineObject27.md) |  | 

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


## InventorySourcesInventorySourcesUpdateCreate

> InventorySourcesInventorySourcesUpdateCreate(ctx, id).Execute()

 Update Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesUpdateCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesUpdateCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesUpdateCreateRequest struct via the builder pattern


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


## InventorySourcesInventorySourcesUpdateRead

> InventorySourcesInventorySourcesUpdateRead(ctx, id).Search(search).Execute()

 Update Inventory Source



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
    resp, r, err := api_client.InventorySourcesApi.InventorySourcesInventorySourcesUpdateRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesApi.InventorySourcesInventorySourcesUpdateRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventorySourcesInventorySourcesUpdateReadRequest struct via the builder pattern


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

