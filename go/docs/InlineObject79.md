# InlineObject79

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSimultaneous** | Pointer to **bool** |  | [optional] 
**AskInventoryOnLaunch** | Pointer to **bool** |  | [optional] 
**AskLimitOnLaunch** | Pointer to **bool** |  | [optional] 
**AskScmBranchOnLaunch** | Pointer to **bool** |  | [optional] 
**AskVariablesOnLaunch** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExtraVars** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Organization** | Pointer to **int32** | The organization used to determine access to this template. | [optional] 
**ScmBranch** | Pointer to **string** |  | [optional] 
**SurveyEnabled** | Pointer to **bool** |  | [optional] 
**WebhookCredential** | Pointer to **int32** | Personal Access Token for posting back the status to the service API | [optional] 
**WebhookService** | Pointer to **string** | Service that webhook requests will be accepted from | [optional] 

## Methods

### NewInlineObject79

`func NewInlineObject79(name string, ) *InlineObject79`

NewInlineObject79 instantiates a new InlineObject79 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject79WithDefaults

`func NewInlineObject79WithDefaults() *InlineObject79`

NewInlineObject79WithDefaults instantiates a new InlineObject79 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSimultaneous

`func (o *InlineObject79) GetAllowSimultaneous() bool`

GetAllowSimultaneous returns the AllowSimultaneous field if non-nil, zero value otherwise.

### GetAllowSimultaneousOk

`func (o *InlineObject79) GetAllowSimultaneousOk() (*bool, bool)`

GetAllowSimultaneousOk returns a tuple with the AllowSimultaneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneous

`func (o *InlineObject79) SetAllowSimultaneous(v bool)`

SetAllowSimultaneous sets AllowSimultaneous field to given value.

### HasAllowSimultaneous

`func (o *InlineObject79) HasAllowSimultaneous() bool`

HasAllowSimultaneous returns a boolean if a field has been set.

### GetAskInventoryOnLaunch

`func (o *InlineObject79) GetAskInventoryOnLaunch() bool`

GetAskInventoryOnLaunch returns the AskInventoryOnLaunch field if non-nil, zero value otherwise.

### GetAskInventoryOnLaunchOk

`func (o *InlineObject79) GetAskInventoryOnLaunchOk() (*bool, bool)`

GetAskInventoryOnLaunchOk returns a tuple with the AskInventoryOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskInventoryOnLaunch

`func (o *InlineObject79) SetAskInventoryOnLaunch(v bool)`

SetAskInventoryOnLaunch sets AskInventoryOnLaunch field to given value.

### HasAskInventoryOnLaunch

`func (o *InlineObject79) HasAskInventoryOnLaunch() bool`

HasAskInventoryOnLaunch returns a boolean if a field has been set.

### GetAskLimitOnLaunch

`func (o *InlineObject79) GetAskLimitOnLaunch() bool`

GetAskLimitOnLaunch returns the AskLimitOnLaunch field if non-nil, zero value otherwise.

### GetAskLimitOnLaunchOk

`func (o *InlineObject79) GetAskLimitOnLaunchOk() (*bool, bool)`

GetAskLimitOnLaunchOk returns a tuple with the AskLimitOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskLimitOnLaunch

`func (o *InlineObject79) SetAskLimitOnLaunch(v bool)`

SetAskLimitOnLaunch sets AskLimitOnLaunch field to given value.

### HasAskLimitOnLaunch

`func (o *InlineObject79) HasAskLimitOnLaunch() bool`

HasAskLimitOnLaunch returns a boolean if a field has been set.

### GetAskScmBranchOnLaunch

`func (o *InlineObject79) GetAskScmBranchOnLaunch() bool`

GetAskScmBranchOnLaunch returns the AskScmBranchOnLaunch field if non-nil, zero value otherwise.

### GetAskScmBranchOnLaunchOk

`func (o *InlineObject79) GetAskScmBranchOnLaunchOk() (*bool, bool)`

GetAskScmBranchOnLaunchOk returns a tuple with the AskScmBranchOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskScmBranchOnLaunch

`func (o *InlineObject79) SetAskScmBranchOnLaunch(v bool)`

SetAskScmBranchOnLaunch sets AskScmBranchOnLaunch field to given value.

### HasAskScmBranchOnLaunch

`func (o *InlineObject79) HasAskScmBranchOnLaunch() bool`

HasAskScmBranchOnLaunch returns a boolean if a field has been set.

### GetAskVariablesOnLaunch

`func (o *InlineObject79) GetAskVariablesOnLaunch() bool`

GetAskVariablesOnLaunch returns the AskVariablesOnLaunch field if non-nil, zero value otherwise.

### GetAskVariablesOnLaunchOk

`func (o *InlineObject79) GetAskVariablesOnLaunchOk() (*bool, bool)`

GetAskVariablesOnLaunchOk returns a tuple with the AskVariablesOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskVariablesOnLaunch

`func (o *InlineObject79) SetAskVariablesOnLaunch(v bool)`

SetAskVariablesOnLaunch sets AskVariablesOnLaunch field to given value.

### HasAskVariablesOnLaunch

`func (o *InlineObject79) HasAskVariablesOnLaunch() bool`

HasAskVariablesOnLaunch returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject79) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject79) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject79) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject79) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraVars

`func (o *InlineObject79) GetExtraVars() string`

GetExtraVars returns the ExtraVars field if non-nil, zero value otherwise.

### GetExtraVarsOk

`func (o *InlineObject79) GetExtraVarsOk() (*string, bool)`

GetExtraVarsOk returns a tuple with the ExtraVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVars

`func (o *InlineObject79) SetExtraVars(v string)`

SetExtraVars sets ExtraVars field to given value.

### HasExtraVars

`func (o *InlineObject79) HasExtraVars() bool`

HasExtraVars returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject79) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject79) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject79) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject79) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject79) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject79) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject79) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject79) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *InlineObject79) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject79) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject79) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *InlineObject79) GetOrganization() int32`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InlineObject79) GetOrganizationOk() (*int32, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InlineObject79) SetOrganization(v int32)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InlineObject79) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScmBranch

`func (o *InlineObject79) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject79) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject79) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject79) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetSurveyEnabled

`func (o *InlineObject79) GetSurveyEnabled() bool`

GetSurveyEnabled returns the SurveyEnabled field if non-nil, zero value otherwise.

### GetSurveyEnabledOk

`func (o *InlineObject79) GetSurveyEnabledOk() (*bool, bool)`

GetSurveyEnabledOk returns a tuple with the SurveyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyEnabled

`func (o *InlineObject79) SetSurveyEnabled(v bool)`

SetSurveyEnabled sets SurveyEnabled field to given value.

### HasSurveyEnabled

`func (o *InlineObject79) HasSurveyEnabled() bool`

HasSurveyEnabled returns a boolean if a field has been set.

### GetWebhookCredential

`func (o *InlineObject79) GetWebhookCredential() int32`

GetWebhookCredential returns the WebhookCredential field if non-nil, zero value otherwise.

### GetWebhookCredentialOk

`func (o *InlineObject79) GetWebhookCredentialOk() (*int32, bool)`

GetWebhookCredentialOk returns a tuple with the WebhookCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCredential

`func (o *InlineObject79) SetWebhookCredential(v int32)`

SetWebhookCredential sets WebhookCredential field to given value.

### HasWebhookCredential

`func (o *InlineObject79) HasWebhookCredential() bool`

HasWebhookCredential returns a boolean if a field has been set.

### GetWebhookService

`func (o *InlineObject79) GetWebhookService() string`

GetWebhookService returns the WebhookService field if non-nil, zero value otherwise.

### GetWebhookServiceOk

`func (o *InlineObject79) GetWebhookServiceOk() (*string, bool)`

GetWebhookServiceOk returns a tuple with the WebhookService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookService

`func (o *InlineObject79) SetWebhookService(v string)`

SetWebhookService sets WebhookService field to given value.

### HasWebhookService

`func (o *InlineObject79) HasWebhookService() bool`

HasWebhookService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


