# InlineObject50

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOverride** | Pointer to **bool** | Allow changing the SCM branch or revision in a job template that uses this project. | [optional] 
**Credential** | Pointer to **int32** |  | [optional] 
**CustomVirtualenv** | Pointer to **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LocalPath** | Pointer to **string** | Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. | [optional] 
**Name** | **string** |  | 
**Organization** | Pointer to **int32** | The organization used to determine access to this template. | [optional] 
**ScmBranch** | Pointer to **string** | Specific branch, tag or commit to checkout. | [optional] 
**ScmClean** | Pointer to **bool** | Discard any local changes before syncing the project. | [optional] 
**ScmDeleteOnUpdate** | Pointer to **bool** | Delete the project before syncing. | [optional] 
**ScmRefspec** | Pointer to **string** | For git projects, an additional refspec to fetch. | [optional] 
**ScmType** | Pointer to **string** | Specifies the source control system used to store the project. | [optional] 
**ScmUpdateCacheTimeout** | Pointer to **int32** | The number of seconds after the last project update ran that a new project update will be launched as a job dependency. | [optional] 
**ScmUpdateOnLaunch** | Pointer to **bool** | Update the project when a job is launched that uses the project. | [optional] 
**ScmUrl** | Pointer to **string** | The location where the project is stored. | [optional] 
**Timeout** | Pointer to **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] 

## Methods

### NewInlineObject50

`func NewInlineObject50(name string, ) *InlineObject50`

NewInlineObject50 instantiates a new InlineObject50 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject50WithDefaults

`func NewInlineObject50WithDefaults() *InlineObject50`

NewInlineObject50WithDefaults instantiates a new InlineObject50 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowOverride

`func (o *InlineObject50) GetAllowOverride() bool`

GetAllowOverride returns the AllowOverride field if non-nil, zero value otherwise.

### GetAllowOverrideOk

`func (o *InlineObject50) GetAllowOverrideOk() (*bool, bool)`

GetAllowOverrideOk returns a tuple with the AllowOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverride

`func (o *InlineObject50) SetAllowOverride(v bool)`

SetAllowOverride sets AllowOverride field to given value.

### HasAllowOverride

`func (o *InlineObject50) HasAllowOverride() bool`

HasAllowOverride returns a boolean if a field has been set.

### GetCredential

`func (o *InlineObject50) GetCredential() int32`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InlineObject50) GetCredentialOk() (*int32, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InlineObject50) SetCredential(v int32)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InlineObject50) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetCustomVirtualenv

`func (o *InlineObject50) GetCustomVirtualenv() string`

GetCustomVirtualenv returns the CustomVirtualenv field if non-nil, zero value otherwise.

### GetCustomVirtualenvOk

`func (o *InlineObject50) GetCustomVirtualenvOk() (*string, bool)`

GetCustomVirtualenvOk returns a tuple with the CustomVirtualenv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVirtualenv

`func (o *InlineObject50) SetCustomVirtualenv(v string)`

SetCustomVirtualenv sets CustomVirtualenv field to given value.

### HasCustomVirtualenv

`func (o *InlineObject50) HasCustomVirtualenv() bool`

HasCustomVirtualenv returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject50) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject50) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject50) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject50) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalPath

`func (o *InlineObject50) GetLocalPath() string`

GetLocalPath returns the LocalPath field if non-nil, zero value otherwise.

### GetLocalPathOk

`func (o *InlineObject50) GetLocalPathOk() (*string, bool)`

GetLocalPathOk returns a tuple with the LocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPath

`func (o *InlineObject50) SetLocalPath(v string)`

SetLocalPath sets LocalPath field to given value.

### HasLocalPath

`func (o *InlineObject50) HasLocalPath() bool`

HasLocalPath returns a boolean if a field has been set.

### GetName

`func (o *InlineObject50) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject50) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject50) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *InlineObject50) GetOrganization() int32`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InlineObject50) GetOrganizationOk() (*int32, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InlineObject50) SetOrganization(v int32)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InlineObject50) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScmBranch

`func (o *InlineObject50) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject50) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject50) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject50) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetScmClean

`func (o *InlineObject50) GetScmClean() bool`

GetScmClean returns the ScmClean field if non-nil, zero value otherwise.

### GetScmCleanOk

`func (o *InlineObject50) GetScmCleanOk() (*bool, bool)`

GetScmCleanOk returns a tuple with the ScmClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmClean

`func (o *InlineObject50) SetScmClean(v bool)`

SetScmClean sets ScmClean field to given value.

### HasScmClean

`func (o *InlineObject50) HasScmClean() bool`

HasScmClean returns a boolean if a field has been set.

### GetScmDeleteOnUpdate

`func (o *InlineObject50) GetScmDeleteOnUpdate() bool`

GetScmDeleteOnUpdate returns the ScmDeleteOnUpdate field if non-nil, zero value otherwise.

### GetScmDeleteOnUpdateOk

`func (o *InlineObject50) GetScmDeleteOnUpdateOk() (*bool, bool)`

GetScmDeleteOnUpdateOk returns a tuple with the ScmDeleteOnUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmDeleteOnUpdate

`func (o *InlineObject50) SetScmDeleteOnUpdate(v bool)`

SetScmDeleteOnUpdate sets ScmDeleteOnUpdate field to given value.

### HasScmDeleteOnUpdate

`func (o *InlineObject50) HasScmDeleteOnUpdate() bool`

HasScmDeleteOnUpdate returns a boolean if a field has been set.

### GetScmRefspec

`func (o *InlineObject50) GetScmRefspec() string`

GetScmRefspec returns the ScmRefspec field if non-nil, zero value otherwise.

### GetScmRefspecOk

`func (o *InlineObject50) GetScmRefspecOk() (*string, bool)`

GetScmRefspecOk returns a tuple with the ScmRefspec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmRefspec

`func (o *InlineObject50) SetScmRefspec(v string)`

SetScmRefspec sets ScmRefspec field to given value.

### HasScmRefspec

`func (o *InlineObject50) HasScmRefspec() bool`

HasScmRefspec returns a boolean if a field has been set.

### GetScmType

`func (o *InlineObject50) GetScmType() string`

GetScmType returns the ScmType field if non-nil, zero value otherwise.

### GetScmTypeOk

`func (o *InlineObject50) GetScmTypeOk() (*string, bool)`

GetScmTypeOk returns a tuple with the ScmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmType

`func (o *InlineObject50) SetScmType(v string)`

SetScmType sets ScmType field to given value.

### HasScmType

`func (o *InlineObject50) HasScmType() bool`

HasScmType returns a boolean if a field has been set.

### GetScmUpdateCacheTimeout

`func (o *InlineObject50) GetScmUpdateCacheTimeout() int32`

GetScmUpdateCacheTimeout returns the ScmUpdateCacheTimeout field if non-nil, zero value otherwise.

### GetScmUpdateCacheTimeoutOk

`func (o *InlineObject50) GetScmUpdateCacheTimeoutOk() (*int32, bool)`

GetScmUpdateCacheTimeoutOk returns a tuple with the ScmUpdateCacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUpdateCacheTimeout

`func (o *InlineObject50) SetScmUpdateCacheTimeout(v int32)`

SetScmUpdateCacheTimeout sets ScmUpdateCacheTimeout field to given value.

### HasScmUpdateCacheTimeout

`func (o *InlineObject50) HasScmUpdateCacheTimeout() bool`

HasScmUpdateCacheTimeout returns a boolean if a field has been set.

### GetScmUpdateOnLaunch

`func (o *InlineObject50) GetScmUpdateOnLaunch() bool`

GetScmUpdateOnLaunch returns the ScmUpdateOnLaunch field if non-nil, zero value otherwise.

### GetScmUpdateOnLaunchOk

`func (o *InlineObject50) GetScmUpdateOnLaunchOk() (*bool, bool)`

GetScmUpdateOnLaunchOk returns a tuple with the ScmUpdateOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUpdateOnLaunch

`func (o *InlineObject50) SetScmUpdateOnLaunch(v bool)`

SetScmUpdateOnLaunch sets ScmUpdateOnLaunch field to given value.

### HasScmUpdateOnLaunch

`func (o *InlineObject50) HasScmUpdateOnLaunch() bool`

HasScmUpdateOnLaunch returns a boolean if a field has been set.

### GetScmUrl

`func (o *InlineObject50) GetScmUrl() string`

GetScmUrl returns the ScmUrl field if non-nil, zero value otherwise.

### GetScmUrlOk

`func (o *InlineObject50) GetScmUrlOk() (*string, bool)`

GetScmUrlOk returns a tuple with the ScmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUrl

`func (o *InlineObject50) SetScmUrl(v string)`

SetScmUrl sets ScmUrl field to given value.

### HasScmUrl

`func (o *InlineObject50) HasScmUrl() bool`

HasScmUrl returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineObject50) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject50) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject50) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject50) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


