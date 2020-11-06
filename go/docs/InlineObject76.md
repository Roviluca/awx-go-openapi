# InlineObject76

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
**WorkflowJobTemplate** | **string** |  | 

## Methods

### NewInlineObject76

`func NewInlineObject76(workflowJobTemplate string, ) *InlineObject76`

NewInlineObject76 instantiates a new InlineObject76 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject76WithDefaults

`func NewInlineObject76WithDefaults() *InlineObject76`

NewInlineObject76WithDefaults instantiates a new InlineObject76 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllParentsMustConverge

`func (o *InlineObject76) GetAllParentsMustConverge() bool`

GetAllParentsMustConverge returns the AllParentsMustConverge field if non-nil, zero value otherwise.

### GetAllParentsMustConvergeOk

`func (o *InlineObject76) GetAllParentsMustConvergeOk() (*bool, bool)`

GetAllParentsMustConvergeOk returns a tuple with the AllParentsMustConverge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllParentsMustConverge

`func (o *InlineObject76) SetAllParentsMustConverge(v bool)`

SetAllParentsMustConverge sets AllParentsMustConverge field to given value.

### HasAllParentsMustConverge

`func (o *InlineObject76) HasAllParentsMustConverge() bool`

HasAllParentsMustConverge returns a boolean if a field has been set.

### GetDiffMode

`func (o *InlineObject76) GetDiffMode() bool`

GetDiffMode returns the DiffMode field if non-nil, zero value otherwise.

### GetDiffModeOk

`func (o *InlineObject76) GetDiffModeOk() (*bool, bool)`

GetDiffModeOk returns a tuple with the DiffMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffMode

`func (o *InlineObject76) SetDiffMode(v bool)`

SetDiffMode sets DiffMode field to given value.

### HasDiffMode

`func (o *InlineObject76) HasDiffMode() bool`

HasDiffMode returns a boolean if a field has been set.

### GetExtraData

`func (o *InlineObject76) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InlineObject76) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InlineObject76) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InlineObject76) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetIdentifier

`func (o *InlineObject76) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InlineObject76) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InlineObject76) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InlineObject76) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject76) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject76) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject76) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject76) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetJobTags

`func (o *InlineObject76) GetJobTags() string`

GetJobTags returns the JobTags field if non-nil, zero value otherwise.

### GetJobTagsOk

`func (o *InlineObject76) GetJobTagsOk() (*string, bool)`

GetJobTagsOk returns a tuple with the JobTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTags

`func (o *InlineObject76) SetJobTags(v string)`

SetJobTags sets JobTags field to given value.

### HasJobTags

`func (o *InlineObject76) HasJobTags() bool`

HasJobTags returns a boolean if a field has been set.

### GetJobType

`func (o *InlineObject76) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *InlineObject76) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *InlineObject76) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *InlineObject76) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject76) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject76) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject76) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject76) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetScmBranch

`func (o *InlineObject76) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject76) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject76) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject76) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetSkipTags

`func (o *InlineObject76) GetSkipTags() string`

GetSkipTags returns the SkipTags field if non-nil, zero value otherwise.

### GetSkipTagsOk

`func (o *InlineObject76) GetSkipTagsOk() (*string, bool)`

GetSkipTagsOk returns a tuple with the SkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTags

`func (o *InlineObject76) SetSkipTags(v string)`

SetSkipTags sets SkipTags field to given value.

### HasSkipTags

`func (o *InlineObject76) HasSkipTags() bool`

HasSkipTags returns a boolean if a field has been set.

### GetUnifiedJobTemplate

`func (o *InlineObject76) GetUnifiedJobTemplate() int32`

GetUnifiedJobTemplate returns the UnifiedJobTemplate field if non-nil, zero value otherwise.

### GetUnifiedJobTemplateOk

`func (o *InlineObject76) GetUnifiedJobTemplateOk() (*int32, bool)`

GetUnifiedJobTemplateOk returns a tuple with the UnifiedJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedJobTemplate

`func (o *InlineObject76) SetUnifiedJobTemplate(v int32)`

SetUnifiedJobTemplate sets UnifiedJobTemplate field to given value.

### HasUnifiedJobTemplate

`func (o *InlineObject76) HasUnifiedJobTemplate() bool`

HasUnifiedJobTemplate returns a boolean if a field has been set.

### GetVerbosity

`func (o *InlineObject76) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *InlineObject76) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *InlineObject76) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *InlineObject76) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.

### GetWorkflowJobTemplate

`func (o *InlineObject76) GetWorkflowJobTemplate() string`

GetWorkflowJobTemplate returns the WorkflowJobTemplate field if non-nil, zero value otherwise.

### GetWorkflowJobTemplateOk

`func (o *InlineObject76) GetWorkflowJobTemplateOk() (*string, bool)`

GetWorkflowJobTemplateOk returns a tuple with the WorkflowJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowJobTemplate

`func (o *InlineObject76) SetWorkflowJobTemplate(v string)`

SetWorkflowJobTemplate sets WorkflowJobTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


