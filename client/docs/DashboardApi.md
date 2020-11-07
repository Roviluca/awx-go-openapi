# \DashboardApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardDashboardGraphsJobsList**](DashboardApi.md#DashboardDashboardGraphsJobsList) | **Get** /api/v2/dashboard/graphs/jobs/ |  View Statistics for Job Runs
[**DashboardDashboardList**](DashboardApi.md#DashboardDashboardList) | **Get** /api/v2/dashboard/ | Show Dashboard Details



## DashboardDashboardGraphsJobsList

> DashboardDashboardGraphsJobsList(ctx).Execute()

 View Statistics for Job Runs



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardDashboardGraphsJobsList(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardDashboardGraphsJobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardDashboardGraphsJobsListRequest struct via the builder pattern


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


## DashboardDashboardList

> DashboardDashboardList(ctx).Execute()

Show Dashboard Details

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.DashboardDashboardList(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.DashboardDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardDashboardListRequest struct via the builder pattern


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

