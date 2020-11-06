# InlineObject57

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DiffMode** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** | Enables processing of this schedule. | [optional] 
**ExtraData** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] 
**JobTags** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Rrule** | **string** | A value representing the schedules iCal recurrence rule. | 
**ScmBranch** | Pointer to **string** |  | [optional] 
**SkipTags** | Pointer to **string** |  | [optional] 
**UnifiedJobTemplate** | **int32** |  | 
**Verbosity** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineObject57

`func NewInlineObject57(name string, rrule string, unifiedJobTemplate int32, ) *InlineObject57`

NewInlineObject57 instantiates a new InlineObject57 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject57WithDefaults

`func NewInlineObject57WithDefaults() *InlineObject57`

NewInlineObject57WithDefaults instantiates a new InlineObject57 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InlineObject57) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject57) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject57) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject57) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiffMode

`func (o *InlineObject57) GetDiffMode() bool`

GetDiffMode returns the DiffMode field if non-nil, zero value otherwise.

### GetDiffModeOk

`func (o *InlineObject57) GetDiffModeOk() (*bool, bool)`

GetDiffModeOk returns a tuple with the DiffMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffMode

`func (o *InlineObject57) SetDiffMode(v bool)`

SetDiffMode sets DiffMode field to given value.

### HasDiffMode

`func (o *InlineObject57) HasDiffMode() bool`

HasDiffMode returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject57) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject57) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject57) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject57) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraData

`func (o *InlineObject57) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InlineObject57) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InlineObject57) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InlineObject57) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject57) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject57) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject57) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject57) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetJobTags

`func (o *InlineObject57) GetJobTags() string`

GetJobTags returns the JobTags field if non-nil, zero value otherwise.

### GetJobTagsOk

`func (o *InlineObject57) GetJobTagsOk() (*string, bool)`

GetJobTagsOk returns a tuple with the JobTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTags

`func (o *InlineObject57) SetJobTags(v string)`

SetJobTags sets JobTags field to given value.

### HasJobTags

`func (o *InlineObject57) HasJobTags() bool`

HasJobTags returns a boolean if a field has been set.

### GetJobType

`func (o *InlineObject57) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *InlineObject57) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *InlineObject57) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *InlineObject57) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject57) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject57) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject57) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject57) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *InlineObject57) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject57) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject57) SetName(v string)`

SetName sets Name field to given value.


### GetRrule

`func (o *InlineObject57) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *InlineObject57) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *InlineObject57) SetRrule(v string)`

SetRrule sets Rrule field to given value.


### GetScmBranch

`func (o *InlineObject57) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject57) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject57) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject57) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetSkipTags

`func (o *InlineObject57) GetSkipTags() string`

GetSkipTags returns the SkipTags field if non-nil, zero value otherwise.

### GetSkipTagsOk

`func (o *InlineObject57) GetSkipTagsOk() (*string, bool)`

GetSkipTagsOk returns a tuple with the SkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTags

`func (o *InlineObject57) SetSkipTags(v string)`

SetSkipTags sets SkipTags field to given value.

### HasSkipTags

`func (o *InlineObject57) HasSkipTags() bool`

HasSkipTags returns a boolean if a field has been set.

### GetUnifiedJobTemplate

`func (o *InlineObject57) GetUnifiedJobTemplate() int32`

GetUnifiedJobTemplate returns the UnifiedJobTemplate field if non-nil, zero value otherwise.

### GetUnifiedJobTemplateOk

`func (o *InlineObject57) GetUnifiedJobTemplateOk() (*int32, bool)`

GetUnifiedJobTemplateOk returns a tuple with the UnifiedJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedJobTemplate

`func (o *InlineObject57) SetUnifiedJobTemplate(v int32)`

SetUnifiedJobTemplate sets UnifiedJobTemplate field to given value.


### GetVerbosity

`func (o *InlineObject57) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *InlineObject57) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *InlineObject57) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *InlineObject57) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


