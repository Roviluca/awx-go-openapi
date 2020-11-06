# \AdHocCommandEventsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdHocCommandEventsAdHocCommandEventsList**](AdHocCommandEventsApi.md#AdHocCommandEventsAdHocCommandEventsList) | **Get** /api/v2/ad_hoc_command_events/ |  List Ad Hoc Command Events
[**AdHocCommandEventsAdHocCommandEventsRead**](AdHocCommandEventsApi.md#AdHocCommandEventsAdHocCommandEventsRead) | **Get** /api/v2/ad_hoc_command_events/{id}/ |  Retrieve an Ad Hoc Command Event



## AdHocCommandEventsAdHocCommandEventsList

> AdHocCommandEventsAdHocCommandEventsList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Ad Hoc Command Events



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
    resp, r, err := api_client.AdHocCommandEventsApi.AdHocCommandEventsAdHocCommandEventsList(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandEventsApi.AdHocCommandEventsAdHocCommandEventsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdHocCommandEventsAdHocCommandEventsListRequest struct via the builder pattern


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


## AdHocCommandEventsAdHocCommandEventsRead

> AdHocCommandEventsAdHocCommandEventsRead(ctx, id).Search(search).Execute()

 Retrieve an Ad Hoc Command Event



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
    resp, r, err := api_client.AdHocCommandEventsApi.AdHocCommandEventsAdHocCommandEventsRead(context.Background(), id).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdHocCommandEventsApi.AdHocCommandEventsAdHocCommandEventsRead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdHocCommandEventsAdHocCommandEventsReadRequest struct via the builder pattern


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

