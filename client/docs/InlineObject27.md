# InlineObject27

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to **int32** | Cloud credential to use for inventory updates. | [optional] 
**CustomVirtualenv** | Pointer to **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnabledValue** | Pointer to **string** | Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var&#x3D;\&quot;status.power_state\&quot;and enabled_value&#x3D;\&quot;powered_on\&quot; with host variables:{   \&quot;status\&quot;: {     \&quot;power_state\&quot;: \&quot;powered_on\&quot;,     \&quot;created\&quot;: \&quot;2018-02-01T08:00:00.000000Z:00\&quot;,     \&quot;healthy\&quot;: true    },    \&quot;name\&quot;: \&quot;foobar\&quot;,    \&quot;ip_address\&quot;: \&quot;192.168.2.1\&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported into Tower. If the key is not found then the host will be enabled | [optional] 
**EnabledVar** | Pointer to **string** | Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \&quot;foo.bar\&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\&quot;foo\&quot;, {}).get(\&quot;bar\&quot;, default) | [optional] 
**HostFilter** | Pointer to **string** | Regex where only matching hosts will be imported into Tower. | [optional] 
**Inventory** | **int32** |  | 
**Name** | **string** |  | 
**Overwrite** | Pointer to **bool** | Overwrite local groups and hosts from remote inventory source. | [optional] 
**OverwriteVars** | Pointer to **bool** | Overwrite local variables from remote inventory source. | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourcePath** | Pointer to **string** |  | [optional] 
**SourceProject** | Pointer to **string** | Project containing inventory file used as source. | [optional] 
**SourceScript** | Pointer to **int32** |  | [optional] 
**SourceVars** | Pointer to **string** | Inventory source variables in YAML or JSON format. | [optional] 
**Timeout** | Pointer to **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] 
**UpdateCacheTimeout** | Pointer to **int32** |  | [optional] 
**UpdateOnLaunch** | Pointer to **bool** |  | [optional] 
**UpdateOnProjectUpdate** | Pointer to **bool** |  | [optional] 
**Verbosity** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineObject27

`func NewInlineObject27(inventory int32, name string, ) *InlineObject27`

NewInlineObject27 instantiates a new InlineObject27 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject27WithDefaults

`func NewInlineObject27WithDefaults() *InlineObject27`

NewInlineObject27WithDefaults instantiates a new InlineObject27 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *InlineObject27) GetCredential() int32`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InlineObject27) GetCredentialOk() (*int32, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InlineObject27) SetCredential(v int32)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InlineObject27) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetCustomVirtualenv

`func (o *InlineObject27) GetCustomVirtualenv() string`

GetCustomVirtualenv returns the CustomVirtualenv field if non-nil, zero value otherwise.

### GetCustomVirtualenvOk

`func (o *InlineObject27) GetCustomVirtualenvOk() (*string, bool)`

GetCustomVirtualenvOk returns a tuple with the CustomVirtualenv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVirtualenv

`func (o *InlineObject27) SetCustomVirtualenv(v string)`

SetCustomVirtualenv sets CustomVirtualenv field to given value.

### HasCustomVirtualenv

`func (o *InlineObject27) HasCustomVirtualenv() bool`

HasCustomVirtualenv returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject27) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject27) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject27) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject27) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabledValue

`func (o *InlineObject27) GetEnabledValue() string`

GetEnabledValue returns the EnabledValue field if non-nil, zero value otherwise.

### GetEnabledValueOk

`func (o *InlineObject27) GetEnabledValueOk() (*string, bool)`

GetEnabledValueOk returns a tuple with the EnabledValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledValue

`func (o *InlineObject27) SetEnabledValue(v string)`

SetEnabledValue sets EnabledValue field to given value.

### HasEnabledValue

`func (o *InlineObject27) HasEnabledValue() bool`

HasEnabledValue returns a boolean if a field has been set.

### GetEnabledVar

`func (o *InlineObject27) GetEnabledVar() string`

GetEnabledVar returns the EnabledVar field if non-nil, zero value otherwise.

### GetEnabledVarOk

`func (o *InlineObject27) GetEnabledVarOk() (*string, bool)`

GetEnabledVarOk returns a tuple with the EnabledVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledVar

`func (o *InlineObject27) SetEnabledVar(v string)`

SetEnabledVar sets EnabledVar field to given value.

### HasEnabledVar

`func (o *InlineObject27) HasEnabledVar() bool`

HasEnabledVar returns a boolean if a field has been set.

### GetHostFilter

`func (o *InlineObject27) GetHostFilter() string`

GetHostFilter returns the HostFilter field if non-nil, zero value otherwise.

### GetHostFilterOk

`func (o *InlineObject27) GetHostFilterOk() (*string, bool)`

GetHostFilterOk returns a tuple with the HostFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFilter

`func (o *InlineObject27) SetHostFilter(v string)`

SetHostFilter sets HostFilter field to given value.

### HasHostFilter

`func (o *InlineObject27) HasHostFilter() bool`

HasHostFilter returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject27) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject27) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject27) SetInventory(v int32)`

SetInventory sets Inventory field to given value.


### GetName

`func (o *InlineObject27) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject27) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject27) SetName(v string)`

SetName sets Name field to given value.


### GetOverwrite

`func (o *InlineObject27) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *InlineObject27) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *InlineObject27) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *InlineObject27) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetOverwriteVars

`func (o *InlineObject27) GetOverwriteVars() bool`

GetOverwriteVars returns the OverwriteVars field if non-nil, zero value otherwise.

### GetOverwriteVarsOk

`func (o *InlineObject27) GetOverwriteVarsOk() (*bool, bool)`

GetOverwriteVarsOk returns a tuple with the OverwriteVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteVars

`func (o *InlineObject27) SetOverwriteVars(v bool)`

SetOverwriteVars sets OverwriteVars field to given value.

### HasOverwriteVars

`func (o *InlineObject27) HasOverwriteVars() bool`

HasOverwriteVars returns a boolean if a field has been set.

### GetSource

`func (o *InlineObject27) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineObject27) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineObject27) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InlineObject27) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourcePath

`func (o *InlineObject27) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *InlineObject27) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *InlineObject27) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *InlineObject27) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.

### GetSourceProject

`func (o *InlineObject27) GetSourceProject() string`

GetSourceProject returns the SourceProject field if non-nil, zero value otherwise.

### GetSourceProjectOk

`func (o *InlineObject27) GetSourceProjectOk() (*string, bool)`

GetSourceProjectOk returns a tuple with the SourceProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProject

`func (o *InlineObject27) SetSourceProject(v string)`

SetSourceProject sets SourceProject field to given value.

### HasSourceProject

`func (o *InlineObject27) HasSourceProject() bool`

HasSourceProject returns a boolean if a field has been set.

### GetSourceScript

`func (o *InlineObject27) GetSourceScript() int32`

GetSourceScript returns the SourceScript field if non-nil, zero value otherwise.

### GetSourceScriptOk

`func (o *InlineObject27) GetSourceScriptOk() (*int32, bool)`

GetSourceScriptOk returns a tuple with the SourceScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceScript

`func (o *InlineObject27) SetSourceScript(v int32)`

SetSourceScript sets SourceScript field to given value.

### HasSourceScript

`func (o *InlineObject27) HasSourceScript() bool`

HasSourceScript returns a boolean if a field has been set.

### GetSourceVars

`func (o *InlineObject27) GetSourceVars() string`

GetSourceVars returns the SourceVars field if non-nil, zero value otherwise.

### GetSourceVarsOk

`func (o *InlineObject27) GetSourceVarsOk() (*string, bool)`

GetSourceVarsOk returns a tuple with the SourceVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVars

`func (o *InlineObject27) SetSourceVars(v string)`

SetSourceVars sets SourceVars field to given value.

### HasSourceVars

`func (o *InlineObject27) HasSourceVars() bool`

HasSourceVars returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineObject27) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject27) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject27) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject27) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUpdateCacheTimeout

`func (o *InlineObject27) GetUpdateCacheTimeout() int32`

GetUpdateCacheTimeout returns the UpdateCacheTimeout field if non-nil, zero value otherwise.

### GetUpdateCacheTimeoutOk

`func (o *InlineObject27) GetUpdateCacheTimeoutOk() (*int32, bool)`

GetUpdateCacheTimeoutOk returns a tuple with the UpdateCacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCacheTimeout

`func (o *InlineObject27) SetUpdateCacheTimeout(v int32)`

SetUpdateCacheTimeout sets UpdateCacheTimeout field to given value.

### HasUpdateCacheTimeout

`func (o *InlineObject27) HasUpdateCacheTimeout() bool`

HasUpdateCacheTimeout returns a boolean if a field has been set.

### GetUpdateOnLaunch

`func (o *InlineObject27) GetUpdateOnLaunch() bool`

GetUpdateOnLaunch returns the UpdateOnLaunch field if non-nil, zero value otherwise.

### GetUpdateOnLaunchOk

`func (o *InlineObject27) GetUpdateOnLaunchOk() (*bool, bool)`

GetUpdateOnLaunchOk returns a tuple with the UpdateOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnLaunch

`func (o *InlineObject27) SetUpdateOnLaunch(v bool)`

SetUpdateOnLaunch sets UpdateOnLaunch field to given value.

### HasUpdateOnLaunch

`func (o *InlineObject27) HasUpdateOnLaunch() bool`

HasUpdateOnLaunch returns a boolean if a field has been set.

### GetUpdateOnProjectUpdate

`func (o *InlineObject27) GetUpdateOnProjectUpdate() bool`

GetUpdateOnProjectUpdate returns the UpdateOnProjectUpdate field if non-nil, zero value otherwise.

### GetUpdateOnProjectUpdateOk

`func (o *InlineObject27) GetUpdateOnProjectUpdateOk() (*bool, bool)`

GetUpdateOnProjectUpdateOk returns a tuple with the UpdateOnProjectUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnProjectUpdate

`func (o *InlineObject27) SetUpdateOnProjectUpdate(v bool)`

SetUpdateOnProjectUpdate sets UpdateOnProjectUpdate field to given value.

### HasUpdateOnProjectUpdate

`func (o *InlineObject27) HasUpdateOnProjectUpdate() bool`

HasUpdateOnProjectUpdate returns a boolean if a field has been set.

### GetVerbosity

`func (o *InlineObject27) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *InlineObject27) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *InlineObject27) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *InlineObject27) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


