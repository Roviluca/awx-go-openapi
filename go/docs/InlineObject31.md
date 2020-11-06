# InlineObject31

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

### NewInlineObject31

`func NewInlineObject31(name string, rrule string, unifiedJobTemplate int32, ) *InlineObject31`

NewInlineObject31 instantiates a new InlineObject31 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject31WithDefaults

`func NewInlineObject31WithDefaults() *InlineObject31`

NewInlineObject31WithDefaults instantiates a new InlineObject31 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InlineObject31) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject31) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject31) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject31) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiffMode

`func (o *InlineObject31) GetDiffMode() bool`

GetDiffMode returns the DiffMode field if non-nil, zero value otherwise.

### GetDiffModeOk

`func (o *InlineObject31) GetDiffModeOk() (*bool, bool)`

GetDiffModeOk returns a tuple with the DiffMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffMode

`func (o *InlineObject31) SetDiffMode(v bool)`

SetDiffMode sets DiffMode field to given value.

### HasDiffMode

`func (o *InlineObject31) HasDiffMode() bool`

HasDiffMode returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject31) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject31) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject31) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject31) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraData

`func (o *InlineObject31) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InlineObject31) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InlineObject31) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InlineObject31) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject31) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject31) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject31) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject31) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetJobTags

`func (o *InlineObject31) GetJobTags() string`

GetJobTags returns the JobTags field if non-nil, zero value otherwise.

### GetJobTagsOk

`func (o *InlineObject31) GetJobTagsOk() (*string, bool)`

GetJobTagsOk returns a tuple with the JobTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTags

`func (o *InlineObject31) SetJobTags(v string)`

SetJobTags sets JobTags field to given value.

### HasJobTags

`func (o *InlineObject31) HasJobTags() bool`

HasJobTags returns a boolean if a field has been set.

### GetJobType

`func (o *InlineObject31) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *InlineObject31) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *InlineObject31) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *InlineObject31) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject31) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject31) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject31) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject31) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *InlineObject31) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject31) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject31) SetName(v string)`

SetName sets Name field to given value.


### GetRrule

`func (o *InlineObject31) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *InlineObject31) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *InlineObject31) SetRrule(v string)`

SetRrule sets Rrule field to given value.


### GetScmBranch

`func (o *InlineObject31) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject31) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject31) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject31) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetSkipTags

`func (o *InlineObject31) GetSkipTags() string`

GetSkipTags returns the SkipTags field if non-nil, zero value otherwise.

### GetSkipTagsOk

`func (o *InlineObject31) GetSkipTagsOk() (*string, bool)`

GetSkipTagsOk returns a tuple with the SkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTags

`func (o *InlineObject31) SetSkipTags(v string)`

SetSkipTags sets SkipTags field to given value.

### HasSkipTags

`func (o *InlineObject31) HasSkipTags() bool`

HasSkipTags returns a boolean if a field has been set.

### GetUnifiedJobTemplate

`func (o *InlineObject31) GetUnifiedJobTemplate() int32`

GetUnifiedJobTemplate returns the UnifiedJobTemplate field if non-nil, zero value otherwise.

### GetUnifiedJobTemplateOk

`func (o *InlineObject31) GetUnifiedJobTemplateOk() (*int32, bool)`

GetUnifiedJobTemplateOk returns a tuple with the UnifiedJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedJobTemplate

`func (o *InlineObject31) SetUnifiedJobTemplate(v int32)`

SetUnifiedJobTemplate sets UnifiedJobTemplate field to given value.


### GetVerbosity

`func (o *InlineObject31) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *InlineObject31) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *InlineObject31) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *InlineObject31) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


