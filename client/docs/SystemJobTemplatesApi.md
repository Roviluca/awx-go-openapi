# \SystemJobTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemJobTemplatesSystemJobTemplatesJobsList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesJobsList) | **Get** /api/v2/system_job_templates/{id}/jobs/ |  List System Jobs for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesLaunchCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesLaunchCreate) | **Post** /api/v2/system_job_templates/{id}/launch/ | Launch a Job Template
[**SystemJobTemplatesSystemJobTemplatesLaunchList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesLaunchList) | **Get** /api/v2/system_job_templates/{id}/launch/ | Launch a Job Template
[**SystemJobTemplatesSystemJobTemplatesList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesList) | **Get** /api/v2/system_job_templates/ |  List System Job Templates
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate) | **Post** /api/v2/system_job_templates/{id}/notification_templates_error/ |  Create a Notification Template for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList) | **Get** /api/v2/system_job_templates/{id}/notification_templates_error/ |  List Notification Templates for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate) | **Post** /api/v2/system_job_templates/{id}/notification_templates_started/ |  Create a Notification Template for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList) | **Get** /api/v2/system_job_templates/{id}/notification_templates_started/ |  List Notification Templates for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate) | **Post** /api/v2/system_job_templates/{id}/notification_templates_success/ |  Create a Notification Template for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList) | **Get** /api/v2/system_job_templates/{id}/notification_templates_success/ |  List Notification Templates for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesRead**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesRead) | **Get** /api/v2/system_job_templates/{id}/ |  Retrieve a System Job Template
[**SystemJobTemplatesSystemJobTemplatesSchedulesCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesSchedulesCreate) | **Post** /api/v2/system_job_templates/{id}/schedules/ |  Create a Schedule for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesSchedulesList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesSchedulesList) | **Get** /api/v2/system_job_templates/{id}/schedules/ |  List Schedules for a System Job Template



## SystemJobTemplatesSystemJobTemplatesJobsList

> SystemJobTemplatesSystemJobTemplatesJobsList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List System Jobs for a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesJobsList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesJobsList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesJobsListRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesLaunchCreate

> SystemJobTemplatesSystemJobTemplatesLaunchCreate(ctx, id).Execute()

Launch a Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesLaunchCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesLaunchCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesLaunchCreateRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesLaunchList

> SystemJobTemplatesSystemJobTemplatesLaunchList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

Launch a Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesLaunchList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesLaunchList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesLaunchListRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesList

> SystemJobTemplatesSystemJobTemplatesList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List System Job Templates



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesList(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesListRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate

> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate(ctx, id).Data(data).Execute()

 Create a Notification Template for a System Job Template



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
    data := *openapiclient.Newinline_object_61("Name_example", "NotificationType_example", 123) // InlineObject61 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject61**](InlineObject61.md) |  | 

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


## SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList

> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates for a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate

> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate(ctx, id).Data(data).Execute()

 Create a Notification Template for a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList

> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates for a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedListRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate

> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate(ctx, id).Data(data).Execute()

 Create a Notification Template for a System Job Template



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
    data := *openapiclient.Newinline_object_62("Name_example", "NotificationType_example", 123) // InlineObject62 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject62**](InlineObject62.md) |  | 

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


## SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList

> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Notification Templates for a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesRead

> SystemJobTemplatesSystemJobTemplatesRead(ctx, id).Search(search).Execute()

 Retrieve a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesReadRequest struct via the builder pattern


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


## SystemJobTemplatesSystemJobTemplatesSchedulesCreate

> SystemJobTemplatesSystemJobTemplatesSchedulesCreate(ctx, id).Data(data).Execute()

 Create a Schedule for a System Job Template



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
    data := *openapiclient.Newinline_object_63("Name_example", "Rrule_example", 123) // InlineObject63 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesSchedulesCreate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesSchedulesCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesSchedulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InlineObject63**](InlineObject63.md) |  | 

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


## SystemJobTemplatesSystemJobTemplatesSchedulesList

> SystemJobTemplatesSystemJobTemplatesSchedulesList(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()

 List Schedules for a System Job Template



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
    resp, r, err := api_client.SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesSchedulesList(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemJobTemplatesApi.SystemJobTemplatesSystemJobTemplatesSchedulesList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSystemJobTemplatesSystemJobTemplatesSchedulesListRequest struct via the builder pattern


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

