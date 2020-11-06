# InlineObject77

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllParentsMustConverge** | Pointer to **bool** | If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node | [optional] 
**DiffMode** | Pointer to **bool** |  | [optional] 
**ExtraData** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** | An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. | [optional] 
**Inventory** | Pointer to **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] 
**JobTags** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**ScmBranch** | Pointer to **string** |  | [optional] 
**SkipTags** | Pointer to **string** |  | [optional] 
**UnifiedJobTemplate** | Pointer to **int32** |  | [optional] 
**Verbosity** | Pointer to **string** |  | [optional] 
**WorkflowJobTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineObject77

`func NewInlineObject77() *InlineObject77`

NewInlineObject77 instantiates a new InlineObject77 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject77WithDefaults

`func NewInlineObject77WithDefaults() *InlineObject77`

NewInlineObject77WithDefaults instantiates a new InlineObject77 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllParentsMustConverge

`func (o *InlineObject77) GetAllParentsMustConverge() bool`

GetAllParentsMustConverge returns the AllParentsMustConverge field if non-nil, zero value otherwise.

### GetAllParentsMustConvergeOk

`func (o *InlineObject77) GetAllParentsMustConvergeOk() (*bool, bool)`

GetAllParentsMustConvergeOk returns a tuple with the AllParentsMustConverge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllParentsMustConverge

`func (o *InlineObject77) SetAllParentsMustConverge(v bool)`

SetAllParentsMustConverge sets AllParentsMustConverge field to given value.

### HasAllParentsMustConverge

`func (o *InlineObject77) HasAllParentsMustConverge() bool`

HasAllParentsMustConverge returns a boolean if a field has been set.

### GetDiffMode

`func (o *InlineObject77) GetDiffMode() bool`

GetDiffMode returns the DiffMode field if non-nil, zero value otherwise.

### GetDiffModeOk

`func (o *InlineObject77) GetDiffModeOk() (*bool, bool)`

GetDiffModeOk returns a tuple with the DiffMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffMode

`func (o *InlineObject77) SetDiffMode(v bool)`

SetDiffMode sets DiffMode field to given value.

### HasDiffMode

`func (o *InlineObject77) HasDiffMode() bool`

HasDiffMode returns a boolean if a field has been set.

### GetExtraData

`func (o *InlineObject77) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InlineObject77) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InlineObject77) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InlineObject77) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetIdentifier

`func (o *InlineObject77) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InlineObject77) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InlineObject77) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InlineObject77) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject77) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject77) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject77) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject77) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetJobTags

`func (o *InlineObject77) GetJobTags() string`

GetJobTags returns the JobTags field if non-nil, zero value otherwise.

### GetJobTagsOk

`func (o *InlineObject77) GetJobTagsOk() (*string, bool)`

GetJobTagsOk returns a tuple with the JobTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTags

`func (o *InlineObject77) SetJobTags(v string)`

SetJobTags sets JobTags field to given value.

### HasJobTags

`func (o *InlineObject77) HasJobTags() bool`

HasJobTags returns a boolean if a field has been set.

### GetJobType

`func (o *InlineObject77) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *InlineObject77) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *InlineObject77) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *InlineObject77) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject77) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject77) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject77) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject77) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetScmBranch

`func (o *InlineObject77) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject77) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject77) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject77) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetSkipTags

`func (o *InlineObject77) GetSkipTags() string`

GetSkipTags returns the SkipTags field if non-nil, zero value otherwise.

### GetSkipTagsOk

`func (o *InlineObject77) GetSkipTagsOk() (*string, bool)`

GetSkipTagsOk returns a tuple with the SkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTags

`func (o *InlineObject77) SetSkipTags(v string)`

SetSkipTags sets SkipTags field to given value.

### HasSkipTags

`func (o *InlineObject77) HasSkipTags() bool`

HasSkipTags returns a boolean if a field has been set.

### GetUnifiedJobTemplate

`func (o *InlineObject77) GetUnifiedJobTemplate() int32`

GetUnifiedJobTemplate returns the UnifiedJobTemplate field if non-nil, zero value otherwise.

### GetUnifiedJobTemplateOk

`func (o *InlineObject77) GetUnifiedJobTemplateOk() (*int32, bool)`

GetUnifiedJobTemplateOk returns a tuple with the UnifiedJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedJobTemplate

`func (o *InlineObject77) SetUnifiedJobTemplate(v int32)`

SetUnifiedJobTemplate sets UnifiedJobTemplate field to given value.

### HasUnifiedJobTemplate

`func (o *InlineObject77) HasUnifiedJobTemplate() bool`

HasUnifiedJobTemplate returns a boolean if a field has been set.

### GetVerbosity

`func (o *InlineObject77) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *InlineObject77) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *InlineObject77) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *InlineObject77) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.

### GetWorkflowJobTemplate

`func (o *InlineObject77) GetWorkflowJobTemplate() string`

GetWorkflowJobTemplate returns the WorkflowJobTemplate field if non-nil, zero value otherwise.

### GetWorkflowJobTemplateOk

`func (o *InlineObject77) GetWorkflowJobTemplateOk() (*string, bool)`

GetWorkflowJobTemplateOk returns a tuple with the WorkflowJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowJobTemplate

`func (o *InlineObject77) SetWorkflowJobTemplate(v string)`

SetWorkflowJobTemplate sets WorkflowJobTemplate field to given value.

### HasWorkflowJobTemplate

`func (o *InlineObject77) HasWorkflowJobTemplate() bool`

HasWorkflowJobTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


