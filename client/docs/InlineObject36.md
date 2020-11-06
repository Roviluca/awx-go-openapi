# InlineObject36

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSimultaneous** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExtraVars** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] 
**IsSlicedJob** | Pointer to **bool** |  | [optional] 
**JobTemplate** | Pointer to **string** | If automatically created for a sliced job run, the job template the workflow job was created from. | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ScmBranch** | Pointer to **string** |  | [optional] 
**WebhookCredential** | Pointer to **int32** | Personal Access Token for posting back the status to the service API | [optional] 
**WebhookGuid** | Pointer to **string** | Unique identifier of the event that triggered this webhook | [optional] 
**WebhookService** | Pointer to **string** | Service that webhook requests will be accepted from | [optional] 
**WorkflowJobTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineObject36

`func NewInlineObject36(name string, ) *InlineObject36`

NewInlineObject36 instantiates a new InlineObject36 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject36WithDefaults

`func NewInlineObject36WithDefaults() *InlineObject36`

NewInlineObject36WithDefaults instantiates a new InlineObject36 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSimultaneous

`func (o *InlineObject36) GetAllowSimultaneous() bool`

GetAllowSimultaneous returns the AllowSimultaneous field if non-nil, zero value otherwise.

### GetAllowSimultaneousOk

`func (o *InlineObject36) GetAllowSimultaneousOk() (*bool, bool)`

GetAllowSimultaneousOk returns a tuple with the AllowSimultaneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneous

`func (o *InlineObject36) SetAllowSimultaneous(v bool)`

SetAllowSimultaneous sets AllowSimultaneous field to given value.

### HasAllowSimultaneous

`func (o *InlineObject36) HasAllowSimultaneous() bool`

HasAllowSimultaneous returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject36) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject36) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject36) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject36) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraVars

`func (o *InlineObject36) GetExtraVars() string`

GetExtraVars returns the ExtraVars field if non-nil, zero value otherwise.

### GetExtraVarsOk

`func (o *InlineObject36) GetExtraVarsOk() (*string, bool)`

GetExtraVarsOk returns a tuple with the ExtraVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVars

`func (o *InlineObject36) SetExtraVars(v string)`

SetExtraVars sets ExtraVars field to given value.

### HasExtraVars

`func (o *InlineObject36) HasExtraVars() bool`

HasExtraVars returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject36) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject36) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject36) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject36) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetIsSlicedJob

`func (o *InlineObject36) GetIsSlicedJob() bool`

GetIsSlicedJob returns the IsSlicedJob field if non-nil, zero value otherwise.

### GetIsSlicedJobOk

`func (o *InlineObject36) GetIsSlicedJobOk() (*bool, bool)`

GetIsSlicedJobOk returns a tuple with the IsSlicedJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlicedJob

`func (o *InlineObject36) SetIsSlicedJob(v bool)`

SetIsSlicedJob sets IsSlicedJob field to given value.

### HasIsSlicedJob

`func (o *InlineObject36) HasIsSlicedJob() bool`

HasIsSlicedJob returns a boolean if a field has been set.

### GetJobTemplate

`func (o *InlineObject36) GetJobTemplate() string`

GetJobTemplate returns the JobTemplate field if non-nil, zero value otherwise.

### GetJobTemplateOk

`func (o *InlineObject36) GetJobTemplateOk() (*string, bool)`

GetJobTemplateOk returns a tuple with the JobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplate

`func (o *InlineObject36) SetJobTemplate(v string)`

SetJobTemplate sets JobTemplate field to given value.

### HasJobTemplate

`func (o *InlineObject36) HasJobTemplate() bool`

HasJobTemplate returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject36) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject36) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject36) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject36) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *InlineObject36) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject36) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject36) SetName(v string)`

SetName sets Name field to given value.


### GetScmBranch

`func (o *InlineObject36) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject36) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject36) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject36) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetWebhookCredential

`func (o *InlineObject36) GetWebhookCredential() int32`

GetWebhookCredential returns the WebhookCredential field if non-nil, zero value otherwise.

### GetWebhookCredentialOk

`func (o *InlineObject36) GetWebhookCredentialOk() (*int32, bool)`

GetWebhookCredentialOk returns a tuple with the WebhookCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCredential

`func (o *InlineObject36) SetWebhookCredential(v int32)`

SetWebhookCredential sets WebhookCredential field to given value.

### HasWebhookCredential

`func (o *InlineObject36) HasWebhookCredential() bool`

HasWebhookCredential returns a boolean if a field has been set.

### GetWebhookGuid

`func (o *InlineObject36) GetWebhookGuid() string`

GetWebhookGuid returns the WebhookGuid field if non-nil, zero value otherwise.

### GetWebhookGuidOk

`func (o *InlineObject36) GetWebhookGuidOk() (*string, bool)`

GetWebhookGuidOk returns a tuple with the WebhookGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookGuid

`func (o *InlineObject36) SetWebhookGuid(v string)`

SetWebhookGuid sets WebhookGuid field to given value.

### HasWebhookGuid

`func (o *InlineObject36) HasWebhookGuid() bool`

HasWebhookGuid returns a boolean if a field has been set.

### GetWebhookService

`func (o *InlineObject36) GetWebhookService() string`

GetWebhookService returns the WebhookService field if non-nil, zero value otherwise.

### GetWebhookServiceOk

`func (o *InlineObject36) GetWebhookServiceOk() (*string, bool)`

GetWebhookServiceOk returns a tuple with the WebhookService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookService

`func (o *InlineObject36) SetWebhookService(v string)`

SetWebhookService sets WebhookService field to given value.

### HasWebhookService

`func (o *InlineObject36) HasWebhookService() bool`

HasWebhookService returns a boolean if a field has been set.

### GetWorkflowJobTemplate

`func (o *InlineObject36) GetWorkflowJobTemplate() string`

GetWorkflowJobTemplate returns the WorkflowJobTemplate field if non-nil, zero value otherwise.

### GetWorkflowJobTemplateOk

`func (o *InlineObject36) GetWorkflowJobTemplateOk() (*string, bool)`

GetWorkflowJobTemplateOk returns a tuple with the WorkflowJobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowJobTemplate

`func (o *InlineObject36) SetWorkflowJobTemplate(v string)`

SetWorkflowJobTemplate sets WorkflowJobTemplate field to given value.

### HasWorkflowJobTemplate

`func (o *InlineObject36) HasWorkflowJobTemplate() bool`

HasWorkflowJobTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


