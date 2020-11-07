# \AdHocCommandsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdHocCommandsAdHocCommandsActivityStreamList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsActivityStreamList) | **Get** /api/v2/ad_hoc_commands/{id}/activity_stream/ |  List Activity Streams for an Ad Hoc Command
[**AdHocCommandsAdHocCommandsCancelCreate**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsCancelCreate) | **Post** /api/v2/ad_hoc_commands/{id}/cancel/ |  Retrieve an Ad Hoc Command
[**AdHocCommandsAdHocCommandsCancelRead**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsCancelRead) | **Get** /api/v2/ad_hoc_commands/{id}/cancel/ |  Retrieve an Ad Hoc Command
[**AdHocCommandsAdHocCommandsCreate**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsCreate) | **Post** /api/v2/ad_hoc_commands/ |  Create an Ad Hoc Command
[**AdHocCommandsAdHocCommandsDelete**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsDelete) | **Delete** /api/v2/ad_hoc_commands/{id}/ |  Delete an Ad Hoc Command
[**AdHocCommandsAdHocCommandsEventsList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsEventsList) | **Get** /api/v2/ad_hoc_commands/{id}/events/ |  List Ad Hoc Command Events for an Ad Hoc Command
[**AdHocCommandsAdHocCommandsList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsList) | **Get** /api/v2/ad_hoc_commands/ |  List Ad Hoc Commands
[**AdHocCommandsAdHocCommandsNotificationsList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsNotificationsList) | **Get** /api/v2/ad_hoc_commands/{id}/notifications/ |  List Notifications for an Ad Hoc Command
[**AdHocCommandsAdHocCommandsRead**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsRead) | **Get** /api/v2/ad_hoc_commands/{id}/ |  Retrieve an Ad Hoc Command
[**AdHocCommandsAdHocCommandsRelaunchCreate**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsRelaunchCreate) | **Post** /api/v2/ad_hoc_commands/{id}/relaunch/ | Relaunch an Ad Hoc Command
[**AdHocCommandsAdHocCommandsRelaunchList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsRelaunchList) | **Get** /api/v2/ad_hoc_commands/{id}/relaunch/ | Relaunch an Ad Hoc Command
[**AdHocCommandsAdHocCommandsStdoutRead**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsStdoutRead) | **Get** /api/v2/ad_hoc_commands/{id}/stdout/ |  Retrieve Ad Hoc Command Stdout



## AdHocCommandsAdHocCommandsActivityStreamList

> AdHocCommandsAdHocCommandsActivityStreamList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Activity Streams for an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsActivityStreamList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsActivityStreamList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsActivityStreamListRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsCancelCreate

> AdHocCommandsAdHocCommandsCancelCreate(ctx, id).Execute()

 Retrieve an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsCancelCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsCancelCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsCancelCreateRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsCancelRead

> AdHocCommandsAdHocCommandsCancelRead(ctx, id).Search(search).Execute()

 Retrieve an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsCancelRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsCancelRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsCancelReadRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsCreate

> AdHocCommandsAdHocCommandsCreate(ctx).Data(data).Execute()

 Create an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsCreate(context.Background(), ).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsCreateRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsDelete

> AdHocCommandsAdHocCommandsDelete(ctx, id).Search(search).Execute()

 Delete an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsDeleteRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsEventsList

> AdHocCommandsAdHocCommandsEventsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Ad Hoc Command Events for an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsEventsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsEventsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsEventsListRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsList

> AdHocCommandsAdHocCommandsList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Ad Hoc Commands



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsList(context.Background(), ).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsListRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsNotificationsList

> AdHocCommandsAdHocCommandsNotificationsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notifications for an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsNotificationsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsNotificationsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsNotificationsListRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsRead

> AdHocCommandsAdHocCommandsRead(ctx, id).Search(search).Execute()

 Retrieve an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsReadRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsRelaunchCreate

> AdHocCommandsAdHocCommandsRelaunchCreate(ctx, id).Execute()

Relaunch an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsRelaunchCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsRelaunchCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsRelaunchCreateRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsRelaunchList

> AdHocCommandsAdHocCommandsRelaunchList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

Relaunch an Ad Hoc Command



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsRelaunchList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsRelaunchList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsRelaunchListRequest struct via the builder pattern


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


## AdHocCommandsAdHocCommandsStdoutRead

> AdHocCommandsAdHocCommandsStdoutRead(ctx, id).Execute()

 Retrieve Ad Hoc Command Stdout



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
    resp, r, err := api_client.AdHocCommandsApi.AdHocCommandsAdHocCommandsStdoutRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandsApi.AdHocCommandsAdHocCommandsStdoutRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandsAdHocCommandsStdoutReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

