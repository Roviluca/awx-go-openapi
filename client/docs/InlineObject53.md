# InlineObject53

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

### NewInlineObject53

`func NewInlineObject53(name string, ) *InlineObject53`

NewInlineObject53 instantiates a new InlineObject53 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject53WithDefaults

`func NewInlineObject53WithDefaults() *InlineObject53`

NewInlineObject53WithDefaults instantiates a new InlineObject53 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowOverride

`func (o *InlineObject53) GetAllowOverride() bool`

GetAllowOverride returns the AllowOverride field if non-nil, zero value otherwise.

### GetAllowOverrideOk

`func (o *InlineObject53) GetAllowOverrideOk() (*bool, bool)`

GetAllowOverrideOk returns a tuple with the AllowOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverride

`func (o *InlineObject53) SetAllowOverride(v bool)`

SetAllowOverride sets AllowOverride field to given value.

### HasAllowOverride

`func (o *InlineObject53) HasAllowOverride() bool`

HasAllowOverride returns a boolean if a field has been set.

### GetCredential

`func (o *InlineObject53) GetCredential() int32`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InlineObject53) GetCredentialOk() (*int32, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InlineObject53) SetCredential(v int32)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InlineObject53) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetCustomVirtualenv

`func (o *InlineObject53) GetCustomVirtualenv() string`

GetCustomVirtualenv returns the CustomVirtualenv field if non-nil, zero value otherwise.

### GetCustomVirtualenvOk

`func (o *InlineObject53) GetCustomVirtualenvOk() (*string, bool)`

GetCustomVirtualenvOk returns a tuple with the CustomVirtualenv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVirtualenv

`func (o *InlineObject53) SetCustomVirtualenv(v string)`

SetCustomVirtualenv sets CustomVirtualenv field to given value.

### HasCustomVirtualenv

`func (o *InlineObject53) HasCustomVirtualenv() bool`

HasCustomVirtualenv returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject53) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject53) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject53) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject53) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalPath

`func (o *InlineObject53) GetLocalPath() string`

GetLocalPath returns the LocalPath field if non-nil, zero value otherwise.

### GetLocalPathOk

`func (o *InlineObject53) GetLocalPathOk() (*string, bool)`

GetLocalPathOk returns a tuple with the LocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPath

`func (o *InlineObject53) SetLocalPath(v string)`

SetLocalPath sets LocalPath field to given value.

### HasLocalPath

`func (o *InlineObject53) HasLocalPath() bool`

HasLocalPath returns a boolean if a field has been set.

### GetName

`func (o *InlineObject53) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject53) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject53) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *InlineObject53) GetOrganization() int32`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InlineObject53) GetOrganizationOk() (*int32, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InlineObject53) SetOrganization(v int32)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InlineObject53) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScmBranch

`func (o *InlineObject53) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject53) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject53) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject53) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetScmClean

`func (o *InlineObject53) GetScmClean() bool`

GetScmClean returns the ScmClean field if non-nil, zero value otherwise.

### GetScmCleanOk

`func (o *InlineObject53) GetScmCleanOk() (*bool, bool)`

GetScmCleanOk returns a tuple with the ScmClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmClean

`func (o *InlineObject53) SetScmClean(v bool)`

SetScmClean sets ScmClean field to given value.

### HasScmClean

`func (o *InlineObject53) HasScmClean() bool`

HasScmClean returns a boolean if a field has been set.

### GetScmDeleteOnUpdate

`func (o *InlineObject53) GetScmDeleteOnUpdate() bool`

GetScmDeleteOnUpdate returns the ScmDeleteOnUpdate field if non-nil, zero value otherwise.

### GetScmDeleteOnUpdateOk

`func (o *InlineObject53) GetScmDeleteOnUpdateOk() (*bool, bool)`

GetScmDeleteOnUpdateOk returns a tuple with the ScmDeleteOnUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmDeleteOnUpdate

`func (o *InlineObject53) SetScmDeleteOnUpdate(v bool)`

SetScmDeleteOnUpdate sets ScmDeleteOnUpdate field to given value.

### HasScmDeleteOnUpdate

`func (o *InlineObject53) HasScmDeleteOnUpdate() bool`

HasScmDeleteOnUpdate returns a boolean if a field has been set.

### GetScmRefspec

`func (o *InlineObject53) GetScmRefspec() string`

GetScmRefspec returns the ScmRefspec field if non-nil, zero value otherwise.

### GetScmRefspecOk

`func (o *InlineObject53) GetScmRefspecOk() (*string, bool)`

GetScmRefspecOk returns a tuple with the ScmRefspec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmRefspec

`func (o *InlineObject53) SetScmRefspec(v string)`

SetScmRefspec sets ScmRefspec field to given value.

### HasScmRefspec

`func (o *InlineObject53) HasScmRefspec() bool`

HasScmRefspec returns a boolean if a field has been set.

### GetScmType

`func (o *InlineObject53) GetScmType() string`

GetScmType returns the ScmType field if non-nil, zero value otherwise.

### GetScmTypeOk

`func (o *InlineObject53) GetScmTypeOk() (*string, bool)`

GetScmTypeOk returns a tuple with the ScmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmType

`func (o *InlineObject53) SetScmType(v string)`

SetScmType sets ScmType field to given value.

### HasScmType

`func (o *InlineObject53) HasScmType() bool`

HasScmType returns a boolean if a field has been set.

### GetScmUpdateCacheTimeout

`func (o *InlineObject53) GetScmUpdateCacheTimeout() int32`

GetScmUpdateCacheTimeout returns the ScmUpdateCacheTimeout field if non-nil, zero value otherwise.

### GetScmUpdateCacheTimeoutOk

`func (o *InlineObject53) GetScmUpdateCacheTimeoutOk() (*int32, bool)`

GetScmUpdateCacheTimeoutOk returns a tuple with the ScmUpdateCacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUpdateCacheTimeout

`func (o *InlineObject53) SetScmUpdateCacheTimeout(v int32)`

SetScmUpdateCacheTimeout sets ScmUpdateCacheTimeout field to given value.

### HasScmUpdateCacheTimeout

`func (o *InlineObject53) HasScmUpdateCacheTimeout() bool`

HasScmUpdateCacheTimeout returns a boolean if a field has been set.

### GetScmUpdateOnLaunch

`func (o *InlineObject53) GetScmUpdateOnLaunch() bool`

GetScmUpdateOnLaunch returns the ScmUpdateOnLaunch field if non-nil, zero value otherwise.

### GetScmUpdateOnLaunchOk

`func (o *InlineObject53) GetScmUpdateOnLaunchOk() (*bool, bool)`

GetScmUpdateOnLaunchOk returns a tuple with the ScmUpdateOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUpdateOnLaunch

`func (o *InlineObject53) SetScmUpdateOnLaunch(v bool)`

SetScmUpdateOnLaunch sets ScmUpdateOnLaunch field to given value.

### HasScmUpdateOnLaunch

`func (o *InlineObject53) HasScmUpdateOnLaunch() bool`

HasScmUpdateOnLaunch returns a boolean if a field has been set.

### GetScmUrl

`func (o *InlineObject53) GetScmUrl() string`

GetScmUrl returns the ScmUrl field if non-nil, zero value otherwise.

### GetScmUrlOk

`func (o *InlineObject53) GetScmUrlOk() (*string, bool)`

GetScmUrlOk returns a tuple with the ScmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUrl

`func (o *InlineObject53) SetScmUrl(v string)`

SetScmUrl sets ScmUrl field to given value.

### HasScmUrl

`func (o *InlineObject53) HasScmUrl() bool`

HasScmUrl returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineObject53) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject53) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject53) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject53) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


