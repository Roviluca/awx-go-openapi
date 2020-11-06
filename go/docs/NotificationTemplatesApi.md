# \NotificationTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationTemplatesNotificationTemplatesCopyCreate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesCopyCreate) | **Post** /api/v2/notification_templates/{id}/copy/ | 
[**NotificationTemplatesNotificationTemplatesCopyList**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesCopyList) | **Get** /api/v2/notification_templates/{id}/copy/ | 
[**NotificationTemplatesNotificationTemplatesCreate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesCreate) | **Post** /api/v2/notification_templates/ |  Create a Notification Template
[**NotificationTemplatesNotificationTemplatesDelete**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesDelete) | **Delete** /api/v2/notification_templates/{id}/ |  Delete a Notification Template
[**NotificationTemplatesNotificationTemplatesList**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesList) | **Get** /api/v2/notification_templates/ |  List Notification Templates
[**NotificationTemplatesNotificationTemplatesNotificationsList**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesNotificationsList) | **Get** /api/v2/notification_templates/{id}/notifications/ |  List Notifications for a Notification Template
[**NotificationTemplatesNotificationTemplatesPartialUpdate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesPartialUpdate) | **Patch** /api/v2/notification_templates/{id}/ |  Update a Notification Template
[**NotificationTemplatesNotificationTemplatesRead**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesRead) | **Get** /api/v2/notification_templates/{id}/ |  Retrieve a Notification Template
[**NotificationTemplatesNotificationTemplatesTestCreate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesTestCreate) | **Post** /api/v2/notification_templates/{id}/test/ | Test a Notification Template
[**NotificationTemplatesNotificationTemplatesUpdate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesUpdate) | **Put** /api/v2/notification_templates/{id}/ |  Update a Notification Template



## NotificationTemplatesNotificationTemplatesCopyCreate

> NotificationTemplatesNotificationTemplatesCopyCreate(ctx, id).Data(data).Execute()



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
    data := *openapiclient.Newinline_object_43("Name_example") // InlineObject43 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesCopyCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesCopyCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesCopyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject43**](InlineObject43.md) |  | 

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


## NotificationTemplatesNotificationTemplatesCopyList

> NotificationTemplatesNotificationTemplatesCopyList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()



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
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesCopyList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesCopyList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesCopyListRequest struct via the builder pattern


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


## NotificationTemplatesNotificationTemplatesCreate

> NotificationTemplatesNotificationTemplatesCreate(ctx).Data(data).Execute()

 Create a Notification Template



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
    data := *openapiclient.Newinline_object_40("Name_example", "NotificationType_example", 123) // InlineObject40 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**InlineObject40**](InlineObject40.md) |  | 

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


## NotificationTemplatesNotificationTemplatesDelete

> NotificationTemplatesNotificationTemplatesDelete(ctx, id).Search(search).Execute()

 Delete a Notification Template



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
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesDelete(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesDeleteRequest struct via the builder pattern


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


## NotificationTemplatesNotificationTemplatesList

> NotificationTemplatesNotificationTemplatesList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates



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
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesList(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesListRequest struct via the builder pattern


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


## NotificationTemplatesNotificationTemplatesNotificationsList

> NotificationTemplatesNotificationTemplatesNotificationsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notifications for a Notification Template



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
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesNotificationsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesNotificationsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesNotificationsListRequest struct via the builder pattern


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


## NotificationTemplatesNotificationTemplatesPartialUpdate

> NotificationTemplatesNotificationTemplatesPartialUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Notification Template



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
    data := *openapiclient.Newinline_object_42() // InlineObject42 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesPartialUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject42**](InlineObject42.md) |  | 

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


## NotificationTemplatesNotificationTemplatesRead

> NotificationTemplatesNotificationTemplatesRead(ctx, id).Search(search).Execute()

 Retrieve a Notification Template



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
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesReadRequest struct via the builder pattern


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


## NotificationTemplatesNotificationTemplatesTestCreate

> NotificationTemplatesNotificationTemplatesTestCreate(ctx, id).Execute()

Test a Notification Template

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
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesTestCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesTestCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesTestCreateRequest struct via the builder pattern


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


## NotificationTemplatesNotificationTemplatesUpdate

> NotificationTemplatesNotificationTemplatesUpdate(ctx, id).Search(search).Data(data).Execute()

 Update a Notification Template



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
    data := *openapiclient.Newinline_object_41("Name_example", "NotificationType_example", 123) // InlineObject41 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesUpdate(context.Background(), id).Search(search).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplatesApi.NotificationTemplatesNotificationTemplatesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotificationTemplatesNotificationTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search term. | 
 **data** | [**InlineObject41**](InlineObject41.md) |  | 

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

