# \UnifiedJobTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnifiedJobTemplatesUnifiedJobTemplatesList**](UnifiedJobTemplatesApi.md#UnifiedJobTemplatesUnifiedJobTemplatesList) | **Get** /api/v2/unified_job_templates/ |  List Unified Job Templates



## UnifiedJobTemplatesUnifiedJobTemplatesList

> UnifiedJobTemplatesUnifiedJobTemplatesList(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

 List Unified Job Templates



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
    resp, r, err := api_client.UnifiedJobTemplatesApi.UnifiedJobTemplatesUnifiedJobTemplatesList(context.Background(), ).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnifiedJobTemplatesApi.UnifiedJobTemplatesUnifiedJobTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnifiedJobTemplatesUnifiedJobTemplatesListRequest struct via the builder pattern


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

