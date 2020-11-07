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
    resp, r, err := api_client.SettingsApi.SettingsSettingsList(context.Background(), ).Page(page).PageSize(pageSize).Execute()
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
    data := openapiclient.inline_object_60{ACTIVITY_STREAM_ENABLED: false, ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC: false, AD_HOC_COMMANDS: []string{"AD_HOC_COMMANDS_example"), ALLOW_JINJA_IN_EXTRA_VARS: "ALLOW_JINJA_IN_EXTRA_VARS_example", ALLOWOAUTH2FOREXTERNALUSERS: false, ANSIBLE_FACT_CACHE_TIMEOUT: 123, AUTH_BASIC_ENABLED: false, AUTHLDAP1BINDDN: "AUTHLDAP1BINDDN_example", AUTHLDAP1BINDPASSWORD: "AUTHLDAP1BINDPASSWORD_example", AUTHLDAP1CONNECTIONOPTIONS: 123, AUTHLDAP1DENYGROUP: "AUTHLDAP1DENYGROUP_example", AUTHLDAP1GROUPSEARCH: []string{"AUTHLDAP1GROUPSEARCH_example"), AUTHLDAP1GROUPTYPE: "AUTHLDAP1GROUPTYPE_example", AUTHLDAP1GROUPTYPEPARAMS: 123, AUTHLDAP1ORGANIZATIONMAP: 123, AUTHLDAP1REQUIREGROUP: "AUTHLDAP1REQUIREGROUP_example", AUTHLDAP1SERVERURI: "AUTHLDAP1SERVERURI_example", AUTHLDAP1STARTTLS: false, AUTHLDAP1TEAMMAP: 123, AUTHLDAP1USERATTRMAP: 123, AUTHLDAP1USERDNTEMPLATE: "AUTHLDAP1USERDNTEMPLATE_example", AUTHLDAP1USERFLAGSBYGROUP: 123, AUTHLDAP1USERSEARCH: []string{"AUTHLDAP1USERSEARCH_example"), AUTHLDAP2BINDDN: "AUTHLDAP2BINDDN_example", AUTHLDAP2BINDPASSWORD: "AUTHLDAP2BINDPASSWORD_example", AUTHLDAP2CONNECTIONOPTIONS: 123, AUTHLDAP2DENYGROUP: "AUTHLDAP2DENYGROUP_example", AUTHLDAP2GROUPSEARCH: []string{"AUTHLDAP2GROUPSEARCH_example"), AUTHLDAP2GROUPTYPE: "AUTHLDAP2GROUPTYPE_example", AUTHLDAP2GROUPTYPEPARAMS: 123, AUTHLDAP2ORGANIZATIONMAP: 123, AUTHLDAP2REQUIREGROUP: "AUTHLDAP2REQUIREGROUP_example", AUTHLDAP2SERVERURI: "AUTHLDAP2SERVERURI_example", AUTHLDAP2STARTTLS: false, AUTHLDAP2TEAMMAP: 123, AUTHLDAP2USERATTRMAP: 123, AUTHLDAP2USERDNTEMPLATE: "AUTHLDAP2USERDNTEMPLATE_example", AUTHLDAP2USERFLAGSBYGROUP: 123, AUTHLDAP2USERSEARCH: []string{"AUTHLDAP2USERSEARCH_example"), AUTHLDAP3BINDDN: "AUTHLDAP3BINDDN_example", AUTHLDAP3BINDPASSWORD: "AUTHLDAP3BINDPASSWORD_example", AUTHLDAP3CONNECTIONOPTIONS: 123, AUTHLDAP3DENYGROUP: "AUTHLDAP3DENYGROUP_example", AUTHLDAP3GROUPSEARCH: []string{"AUTHLDAP3GROUPSEARCH_example"), AUTHLDAP3GROUPTYPE: "AUTHLDAP3GROUPTYPE_example", AUTHLDAP3GROUPTYPEPARAMS: 123, AUTHLDAP3ORGANIZATIONMAP: 123, AUTHLDAP3REQUIREGROUP: "AUTHLDAP3REQUIREGROUP_example", AUTHLDAP3SERVERURI: "AUTHLDAP3SERVERURI_example", AUTHLDAP3STARTTLS: false, AUTHLDAP3TEAMMAP: 123, AUTHLDAP3USERATTRMAP: 123, AUTHLDAP3USERDNTEMPLATE: "AUTHLDAP3USERDNTEMPLATE_example", AUTHLDAP3USERFLAGSBYGROUP: 123, AUTHLDAP3USERSEARCH: []string{"AUTHLDAP3USERSEARCH_example"), AUTHLDAP4BINDDN: "AUTHLDAP4BINDDN_example", AUTHLDAP4BINDPASSWORD: "AUTHLDAP4BINDPASSWORD_example", AUTHLDAP4CONNECTIONOPTIONS: 123, AUTHLDAP4DENYGROUP: "AUTHLDAP4DENYGROUP_example", AUTHLDAP4GROUPSEARCH: []string{"AUTHLDAP4GROUPSEARCH_example"), AUTHLDAP4GROUPTYPE: "AUTHLDAP4GROUPTYPE_example", AUTHLDAP4GROUPTYPEPARAMS: 123, AUTHLDAP4ORGANIZATIONMAP: 123, AUTHLDAP4REQUIREGROUP: "AUTHLDAP4REQUIREGROUP_example", AUTHLDAP4SERVERURI: "AUTHLDAP4SERVERURI_example", AUTHLDAP4STARTTLS: false, AUTHLDAP4TEAMMAP: 123, AUTHLDAP4USERATTRMAP: 123, AUTHLDAP4USERDNTEMPLATE: "AUTHLDAP4USERDNTEMPLATE_example", AUTHLDAP4USERFLAGSBYGROUP: 123, AUTHLDAP4USERSEARCH: []string{"AUTHLDAP4USERSEARCH_example"), AUTHLDAP5BINDDN: "AUTHLDAP5BINDDN_example", AUTHLDAP5BINDPASSWORD: "AUTHLDAP5BINDPASSWORD_example", AUTHLDAP5CONNECTIONOPTIONS: 123, AUTHLDAP5DENYGROUP: "AUTHLDAP5DENYGROUP_example", AUTHLDAP5GROUPSEARCH: []string{"AUTHLDAP5GROUPSEARCH_example"), AUTHLDAP5GROUPTYPE: "AUTHLDAP5GROUPTYPE_example", AUTHLDAP5GROUPTYPEPARAMS: 123, AUTHLDAP5ORGANIZATIONMAP: 123, AUTHLDAP5REQUIREGROUP: "AUTHLDAP5REQUIREGROUP_example", AUTHLDAP5SERVERURI: "AUTHLDAP5SERVERURI_example", AUTHLDAP5STARTTLS: false, AUTHLDAP5TEAMMAP: 123, AUTHLDAP5USERATTRMAP: 123, AUTHLDAP5USERDNTEMPLATE: "AUTHLDAP5USERDNTEMPLATE_example", AUTHLDAP5USERFLAGSBYGROUP: 123, AUTHLDAP5USERSEARCH: []string{"AUTHLDAP5USERSEARCH_example"), AUTH_LDAP_BIND_DN: "AUTH_LDAP_BIND_DN_example", AUTH_LDAP_BIND_PASSWORD: "AUTH_LDAP_BIND_PASSWORD_example", AUTH_LDAP_CONNECTION_OPTIONS: 123, AUTH_LDAP_DENY_GROUP: "AUTH_LDAP_DENY_GROUP_example", AUTH_LDAP_GROUP_SEARCH: []string{"AUTH_LDAP_GROUP_SEARCH_example"), AUTH_LDAP_GROUP_TYPE: "AUTH_LDAP_GROUP_TYPE_example", AUTH_LDAP_GROUP_TYPE_PARAMS: 123, AUTH_LDAP_ORGANIZATION_MAP: 123, AUTH_LDAP_REQUIRE_GROUP: "AUTH_LDAP_REQUIRE_GROUP_example", AUTH_LDAP_SERVER_URI: "AUTH_LDAP_SERVER_URI_example", AUTH_LDAP_START_TLS: false, AUTH_LDAP_TEAM_MAP: 123, AUTH_LDAP_USER_ATTR_MAP: 123, AUTH_LDAP_USER_DN_TEMPLATE: "AUTH_LDAP_USER_DN_TEMPLATE_example", AUTH_LDAP_USER_FLAGS_BY_GROUP: 123, AUTH_LDAP_USER_SEARCH: []string{"AUTH_LDAP_USER_SEARCH_example"), AUTOMATION_ANALYTICS_GATHER_INTERVAL: 123, AUTOMATION_ANALYTICS_LAST_GATHER: "AUTOMATION_ANALYTICS_LAST_GATHER_example", AUTOMATION_ANALYTICS_URL: "AUTOMATION_ANALYTICS_URL_example", AWX_ANSIBLE_CALLBACK_PLUGINS: []string{"AWX_ANSIBLE_CALLBACK_PLUGINS_example"), AWX_COLLECTIONS_ENABLED: false, AWX_ISOLATED_CHECK_INTERVAL: 123, AWX_ISOLATED_CONNECTION_TIMEOUT: 123, AWX_ISOLATED_HOST_KEY_CHECKING: false, AWX_ISOLATED_LAUNCH_TIMEOUT: 123, AWX_PROOT_BASE_PATH: "AWX_PROOT_BASE_PATH_example", AWX_PROOT_ENABLED: false, AWX_PROOT_HIDE_PATHS: []string{"AWX_PROOT_HIDE_PATHS_example"), AWX_PROOT_SHOW_PATHS: []string{"AWX_PROOT_SHOW_PATHS_example"), AWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL: 123, AWX_RESOURCE_PROFILING_ENABLED: false, AWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL: 123, AWX_RESOURCE_PROFILING_PID_POLL_INTERVAL: 123, AWX_ROLES_ENABLED: false, AWX_SHOW_PLAYBOOK_LINKS: false, AWX_TASK_ENV: 123, CUSTOM_LOGIN_INFO: "CUSTOM_LOGIN_INFO_example", CUSTOM_LOGO: "CUSTOM_LOGO_example", CUSTOM_VENV_PATHS: []string{"CUSTOM_VENV_PATHS_example"), DEFAULT_INVENTORY_UPDATE_TIMEOUT: 123, DEFAULT_JOB_TIMEOUT: 123, DEFAULT_PROJECT_UPDATE_TIMEOUT: 123, EVENT_STDOUT_MAX_BYTES_DISPLAY: 123, GALAXY_IGNORE_CERTS: false, INSIGHTS_TRACKING_STATE: false, LOGIN_REDIRECT_OVERRIDE: "LOGIN_REDIRECT_OVERRIDE_example", LOG_AGGREGATOR_ENABLED: false, LOG_AGGREGATOR_HOST: "LOG_AGGREGATOR_HOST_example", LOG_AGGREGATOR_INDIVIDUAL_FACTS: false, LOG_AGGREGATOR_LEVEL: "LOG_AGGREGATOR_LEVEL_example", LOG_AGGREGATOR_LOGGERS: []string{"LOG_AGGREGATOR_LOGGERS_example"), LOG_AGGREGATOR_MAX_DISK_USAGE_GB: 123, LOG_AGGREGATOR_MAX_DISK_USAGE_PATH: "LOG_AGGREGATOR_MAX_DISK_USAGE_PATH_example", LOG_AGGREGATOR_PASSWORD: "LOG_AGGREGATOR_PASSWORD_example", LOG_AGGREGATOR_PORT: 123, LOG_AGGREGATOR_PROTOCOL: "LOG_AGGREGATOR_PROTOCOL_example", LOG_AGGREGATOR_RSYSLOGD_DEBUG: false, LOG_AGGREGATOR_TCP_TIMEOUT: 123, LOG_AGGREGATOR_TOWER_UUID: "LOG_AGGREGATOR_TOWER_UUID_example", LOG_AGGREGATOR_TYPE: "LOG_AGGREGATOR_TYPE_example", LOG_AGGREGATOR_USERNAME: "LOG_AGGREGATOR_USERNAME_example", LOG_AGGREGATOR_VERIFY_CERT: false, MANAGE_ORGANIZATION_AUTH: false, MAX_FORKS: 123, MAX_UI_JOB_EVENTS: 123, OAUTH2PROVIDER: 123, ORG_ADMINS_CAN_SEE_ALL_USERS: false, PROJECT_UPDATE_VVV: false, PROXY_IP_ALLOWED_LIST: []string{"PROXY_IP_ALLOWED_LIST_example"), RADIUS_PORT: 123, RADIUS_SECRET: "RADIUS_SECRET_example", RADIUS_SERVER: "RADIUS_SERVER_example", REDHAT_PASSWORD: "REDHAT_PASSWORD_example", REDHAT_USERNAME: "REDHAT_USERNAME_example", REMOTE_HOST_HEADERS: []string{"REMOTE_HOST_HEADERS_example"), SAML_AUTO_CREATE_OBJECTS: false, SCHEDULE_MAX_JOBS: 123, SESSIONS_PER_USER: 123, SESSION_COOKIE_AGE: 123, SOCIALAUTHAZUREADOAUTH2KEY: "SOCIALAUTHAZUREADOAUTH2KEY_example", SOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP: 123, SOCIALAUTHAZUREADOAUTH2SECRET: "SOCIALAUTHAZUREADOAUTH2SECRET_example", SOCIALAUTHAZUREADOAUTH2TEAMMAP: 123, SOCIAL_AUTH_GITHUB_KEY: "SOCIAL_AUTH_GITHUB_KEY_example", SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP: 123, SOCIAL_AUTH_GITHUB_ORG_KEY: "SOCIAL_AUTH_GITHUB_ORG_KEY_example", SOCIAL_AUTH_GITHUB_ORG_NAME: "SOCIAL_AUTH_GITHUB_ORG_NAME_example", SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP: 123, SOCIAL_AUTH_GITHUB_ORG_SECRET: "SOCIAL_AUTH_GITHUB_ORG_SECRET_example", SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP: 123, SOCIAL_AUTH_GITHUB_SECRET: "SOCIAL_AUTH_GITHUB_SECRET_example", SOCIAL_AUTH_GITHUB_TEAM_ID: "SOCIAL_AUTH_GITHUB_TEAM_ID_example", SOCIAL_AUTH_GITHUB_TEAM_KEY: "SOCIAL_AUTH_GITHUB_TEAM_KEY_example", SOCIAL_AUTH_GITHUB_TEAM_MAP: 123, SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP: 123, SOCIAL_AUTH_GITHUB_TEAM_SECRET: "SOCIAL_AUTH_GITHUB_TEAM_SECRET_example", SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP: 123, SOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS: 123, SOCIALAUTHGOOGLEOAUTH2KEY: "SOCIALAUTHGOOGLEOAUTH2KEY_example", SOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP: 123, SOCIALAUTHGOOGLEOAUTH2SECRET: "SOCIALAUTHGOOGLEOAUTH2SECRET_example", SOCIALAUTHGOOGLEOAUTH2TEAMMAP: 123, SOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS: []string{"SOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS_example"), SOCIAL_AUTH_ORGANIZATION_MAP: 123, SOCIAL_AUTH_SAML_ENABLED_IDPS: 123, SOCIAL_AUTH_SAML_EXTRA_DATA: []string{"SOCIAL_AUTH_SAML_EXTRA_DATA_example"), SOCIAL_AUTH_SAML_ORGANIZATION_ATTR: 123, SOCIAL_AUTH_SAML_ORGANIZATION_MAP: 123, SOCIAL_AUTH_SAML_ORG_INFO: 123, SOCIAL_AUTH_SAML_SECURITY_CONFIG: 123, SOCIAL_AUTH_SAML_SP_ENTITY_ID: "SOCIAL_AUTH_SAML_SP_ENTITY_ID_example", SOCIAL_AUTH_SAML_SP_EXTRA: 123, SOCIAL_AUTH_SAML_SP_PRIVATE_KEY: "SOCIAL_AUTH_SAML_SP_PRIVATE_KEY_example", SOCIAL_AUTH_SAML_SP_PUBLIC_CERT: "SOCIAL_AUTH_SAML_SP_PUBLIC_CERT_example", SOCIAL_AUTH_SAML_SUPPORT_CONTACT: 123, SOCIAL_AUTH_SAML_TEAM_ATTR: 123, SOCIAL_AUTH_SAML_TEAM_MAP: 123, SOCIAL_AUTH_SAML_TECHNICAL_CONTACT: 123, SOCIAL_AUTH_TEAM_MAP: 123, SOCIAL_AUTH_USER_FIELDS: []string{"SOCIAL_AUTH_USER_FIELDS_example"), STDOUT_MAX_BYTES_DISPLAY: 123, TACACSPLUS_AUTH_PROTOCOL: "TACACSPLUS_AUTH_PROTOCOL_example", TACACSPLUS_HOST: "TACACSPLUS_HOST_example", TACACSPLUS_PORT: 123, TACACSPLUS_SECRET: "TACACSPLUS_SECRET_example", TACACSPLUS_SESSION_TIMEOUT: 123, TOWER_URL_BASE: "TOWER_URL_BASE_example", UI_LIVE_UPDATES_ENABLED: false} // InlineObject60 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SettingsSettingsLoggingTestCreate(context.Background(), ).Data(data).Execute()
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

