# InlineObject46

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSimultaneous** | Pointer to **bool** |  | [optional] 
**AskCredentialOnLaunch** | Pointer to **bool** |  | [optional] 
**AskDiffModeOnLaunch** | Pointer to **bool** |  | [optional] 
**AskInventoryOnLaunch** | Pointer to **bool** |  | [optional] 
**AskJobTypeOnLaunch** | Pointer to **bool** |  | [optional] 
**AskLimitOnLaunch** | Pointer to **bool** |  | [optional] 
**AskScmBranchOnLaunch** | Pointer to **bool** |  | [optional] 
**AskSkipTagsOnLaunch** | Pointer to **bool** |  | [optional] 
**AskTagsOnLaunch** | Pointer to **bool** |  | [optional] 
**AskVariablesOnLaunch** | Pointer to **bool** |  | [optional] 
**AskVerbosityOnLaunch** | Pointer to **bool** |  | [optional] 
**BecomeEnabled** | Pointer to **bool** |  | [optional] 
**CustomVirtualenv** | Pointer to **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiffMode** | Pointer to **bool** | If enabled, textual changes made to any templated files on the host are shown in the standard output | [optional] 
**ExtraVars** | Pointer to **string** |  | [optional] 
**ForceHandlers** | Pointer to **bool** |  | [optional] 
**Forks** | Pointer to **int32** |  | [optional] 
**HostConfigKey** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to **int32** |  | [optional] 
**JobSliceCount** | Pointer to **int32** | The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1. | [optional] 
**JobTags** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Playbook** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ScmBranch** | Pointer to **string** | Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true. | [optional] 
**SkipTags** | Pointer to **string** |  | [optional] 
**StartAtTask** | Pointer to **string** |  | [optional] 
**SurveyEnabled** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] 
**UseFactCache** | Pointer to **bool** | If enabled, Tower will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible. | [optional] 
**Verbosity** | Pointer to **string** |  | [optional] 
**WebhookCredential** | Pointer to **int32** | Personal Access Token for posting back the status to the service API | [optional] 
**WebhookService** | Pointer to **string** | Service that webhook requests will be accepted from | [optional] 

## Methods

### NewInlineObject46

`func NewInlineObject46(name string, ) *InlineObject46`

NewInlineObject46 instantiates a new InlineObject46 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject46WithDefaults

`func NewInlineObject46WithDefaults() *InlineObject46`

NewInlineObject46WithDefaults instantiates a new InlineObject46 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSimultaneous

`func (o *InlineObject46) GetAllowSimultaneous() bool`

GetAllowSimultaneous returns the AllowSimultaneous field if non-nil, zero value otherwise.

### GetAllowSimultaneousOk

`func (o *InlineObject46) GetAllowSimultaneousOk() (*bool, bool)`

GetAllowSimultaneousOk returns a tuple with the AllowSimultaneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneous

`func (o *InlineObject46) SetAllowSimultaneous(v bool)`

SetAllowSimultaneous sets AllowSimultaneous field to given value.

### HasAllowSimultaneous

`func (o *InlineObject46) HasAllowSimultaneous() bool`

HasAllowSimultaneous returns a boolean if a field has been set.

### GetAskCredentialOnLaunch

`func (o *InlineObject46) GetAskCredentialOnLaunch() bool`

GetAskCredentialOnLaunch returns the AskCredentialOnLaunch field if non-nil, zero value otherwise.

### GetAskCredentialOnLaunchOk

`func (o *InlineObject46) GetAskCredentialOnLaunchOk() (*bool, bool)`

GetAskCredentialOnLaunchOk returns a tuple with the AskCredentialOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskCredentialOnLaunch

`func (o *InlineObject46) SetAskCredentialOnLaunch(v bool)`

SetAskCredentialOnLaunch sets AskCredentialOnLaunch field to given value.

### HasAskCredentialOnLaunch

`func (o *InlineObject46) HasAskCredentialOnLaunch() bool`

HasAskCredentialOnLaunch returns a boolean if a field has been set.

### GetAskDiffModeOnLaunch

`func (o *InlineObject46) GetAskDiffModeOnLaunch() bool`

GetAskDiffModeOnLaunch returns the AskDiffModeOnLaunch field if non-nil, zero value otherwise.

### GetAskDiffModeOnLaunchOk

`func (o *InlineObject46) GetAskDiffModeOnLaunchOk() (*bool, bool)`

GetAskDiffModeOnLaunchOk returns a tuple with the AskDiffModeOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskDiffModeOnLaunch

`func (o *InlineObject46) SetAskDiffModeOnLaunch(v bool)`

SetAskDiffModeOnLaunch sets AskDiffModeOnLaunch field to given value.

### HasAskDiffModeOnLaunch

`func (o *InlineObject46) HasAskDiffModeOnLaunch() bool`

HasAskDiffModeOnLaunch returns a boolean if a field has been set.

### GetAskInventoryOnLaunch

`func (o *InlineObject46) GetAskInventoryOnLaunch() bool`

GetAskInventoryOnLaunch returns the AskInventoryOnLaunch field if non-nil, zero value otherwise.

### GetAskInventoryOnLaunchOk

`func (o *InlineObject46) GetAskInventoryOnLaunchOk() (*bool, bool)`

GetAskInventoryOnLaunchOk returns a tuple with the AskInventoryOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskInventoryOnLaunch

`func (o *InlineObject46) SetAskInventoryOnLaunch(v bool)`

SetAskInventoryOnLaunch sets AskInventoryOnLaunch field to given value.

### HasAskInventoryOnLaunch

`func (o *InlineObject46) HasAskInventoryOnLaunch() bool`

HasAskInventoryOnLaunch returns a boolean if a field has been set.

### GetAskJobTypeOnLaunch

`func (o *InlineObject46) GetAskJobTypeOnLaunch() bool`

GetAskJobTypeOnLaunch returns the AskJobTypeOnLaunch field if non-nil, zero value otherwise.

### GetAskJobTypeOnLaunchOk

`func (o *InlineObject46) GetAskJobTypeOnLaunchOk() (*bool, bool)`

GetAskJobTypeOnLaunchOk returns a tuple with the AskJobTypeOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskJobTypeOnLaunch

`func (o *InlineObject46) SetAskJobTypeOnLaunch(v bool)`

SetAskJobTypeOnLaunch sets AskJobTypeOnLaunch field to given value.

### HasAskJobTypeOnLaunch

`func (o *InlineObject46) HasAskJobTypeOnLaunch() bool`

HasAskJobTypeOnLaunch returns a boolean if a field has been set.

### GetAskLimitOnLaunch

`func (o *InlineObject46) GetAskLimitOnLaunch() bool`

GetAskLimitOnLaunch returns the AskLimitOnLaunch field if non-nil, zero value otherwise.

### GetAskLimitOnLaunchOk

`func (o *InlineObject46) GetAskLimitOnLaunchOk() (*bool, bool)`

GetAskLimitOnLaunchOk returns a tuple with the AskLimitOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskLimitOnLaunch

`func (o *InlineObject46) SetAskLimitOnLaunch(v bool)`

SetAskLimitOnLaunch sets AskLimitOnLaunch field to given value.

### HasAskLimitOnLaunch

`func (o *InlineObject46) HasAskLimitOnLaunch() bool`

HasAskLimitOnLaunch returns a boolean if a field has been set.

### GetAskScmBranchOnLaunch

`func (o *InlineObject46) GetAskScmBranchOnLaunch() bool`

GetAskScmBranchOnLaunch returns the AskScmBranchOnLaunch field if non-nil, zero value otherwise.

### GetAskScmBranchOnLaunchOk

`func (o *InlineObject46) GetAskScmBranchOnLaunchOk() (*bool, bool)`

GetAskScmBranchOnLaunchOk returns a tuple with the AskScmBranchOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskScmBranchOnLaunch

`func (o *InlineObject46) SetAskScmBranchOnLaunch(v bool)`

SetAskScmBranchOnLaunch sets AskScmBranchOnLaunch field to given value.

### HasAskScmBranchOnLaunch

`func (o *InlineObject46) HasAskScmBranchOnLaunch() bool`

HasAskScmBranchOnLaunch returns a boolean if a field has been set.

### GetAskSkipTagsOnLaunch

`func (o *InlineObject46) GetAskSkipTagsOnLaunch() bool`

GetAskSkipTagsOnLaunch returns the AskSkipTagsOnLaunch field if non-nil, zero value otherwise.

### GetAskSkipTagsOnLaunchOk

`func (o *InlineObject46) GetAskSkipTagsOnLaunchOk() (*bool, bool)`

GetAskSkipTagsOnLaunchOk returns a tuple with the AskSkipTagsOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSkipTagsOnLaunch

`func (o *InlineObject46) SetAskSkipTagsOnLaunch(v bool)`

SetAskSkipTagsOnLaunch sets AskSkipTagsOnLaunch field to given value.

### HasAskSkipTagsOnLaunch

`func (o *InlineObject46) HasAskSkipTagsOnLaunch() bool`

HasAskSkipTagsOnLaunch returns a boolean if a field has been set.

### GetAskTagsOnLaunch

`func (o *InlineObject46) GetAskTagsOnLaunch() bool`

GetAskTagsOnLaunch returns the AskTagsOnLaunch field if non-nil, zero value otherwise.

### GetAskTagsOnLaunchOk

`func (o *InlineObject46) GetAskTagsOnLaunchOk() (*bool, bool)`

GetAskTagsOnLaunchOk returns a tuple with the AskTagsOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskTagsOnLaunch

`func (o *InlineObject46) SetAskTagsOnLaunch(v bool)`

SetAskTagsOnLaunch sets AskTagsOnLaunch field to given value.

### HasAskTagsOnLaunch

`func (o *InlineObject46) HasAskTagsOnLaunch() bool`

HasAskTagsOnLaunch returns a boolean if a field has been set.

### GetAskVariablesOnLaunch

`func (o *InlineObject46) GetAskVariablesOnLaunch() bool`

GetAskVariablesOnLaunch returns the AskVariablesOnLaunch field if non-nil, zero value otherwise.

### GetAskVariablesOnLaunchOk

`func (o *InlineObject46) GetAskVariablesOnLaunchOk() (*bool, bool)`

GetAskVariablesOnLaunchOk returns a tuple with the AskVariablesOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskVariablesOnLaunch

`func (o *InlineObject46) SetAskVariablesOnLaunch(v bool)`

SetAskVariablesOnLaunch sets AskVariablesOnLaunch field to given value.

### HasAskVariablesOnLaunch

`func (o *InlineObject46) HasAskVariablesOnLaunch() bool`

HasAskVariablesOnLaunch returns a boolean if a field has been set.

### GetAskVerbosityOnLaunch

`func (o *InlineObject46) GetAskVerbosityOnLaunch() bool`

GetAskVerbosityOnLaunch returns the AskVerbosityOnLaunch field if non-nil, zero value otherwise.

### GetAskVerbosityOnLaunchOk

`func (o *InlineObject46) GetAskVerbosityOnLaunchOk() (*bool, bool)`

GetAskVerbosityOnLaunchOk returns a tuple with the AskVerbosityOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskVerbosityOnLaunch

`func (o *InlineObject46) SetAskVerbosityOnLaunch(v bool)`

SetAskVerbosityOnLaunch sets AskVerbosityOnLaunch field to given value.

### HasAskVerbosityOnLaunch

`func (o *InlineObject46) HasAskVerbosityOnLaunch() bool`

HasAskVerbosityOnLaunch returns a boolean if a field has been set.

### GetBecomeEnabled

`func (o *InlineObject46) GetBecomeEnabled() bool`

GetBecomeEnabled returns the BecomeEnabled field if non-nil, zero value otherwise.

### GetBecomeEnabledOk

`func (o *InlineObject46) GetBecomeEnabledOk() (*bool, bool)`

GetBecomeEnabledOk returns a tuple with the BecomeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBecomeEnabled

`func (o *InlineObject46) SetBecomeEnabled(v bool)`

SetBecomeEnabled sets BecomeEnabled field to given value.

### HasBecomeEnabled

`func (o *InlineObject46) HasBecomeEnabled() bool`

HasBecomeEnabled returns a boolean if a field has been set.

### GetCustomVirtualenv

`func (o *InlineObject46) GetCustomVirtualenv() string`

GetCustomVirtualenv returns the CustomVirtualenv field if non-nil, zero value otherwise.

### GetCustomVirtualenvOk

`func (o *InlineObject46) GetCustomVirtualenvOk() (*string, bool)`

GetCustomVirtualenvOk returns a tuple with the CustomVirtualenv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVirtualenv

`func (o *InlineObject46) SetCustomVirtualenv(v string)`

SetCustomVirtualenv sets CustomVirtualenv field to given value.

### HasCustomVirtualenv

`func (o *InlineObject46) HasCustomVirtualenv() bool`

HasCustomVirtualenv returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject46) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject46) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject46) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject46) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiffMode

`func (o *InlineObject46) GetDiffMode() bool`

GetDiffMode returns the DiffMode field if non-nil, zero value otherwise.

### GetDiffModeOk

`func (o *InlineObject46) GetDiffModeOk() (*bool, bool)`

GetDiffModeOk returns a tuple with the DiffMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffMode

`func (o *InlineObject46) SetDiffMode(v bool)`

SetDiffMode sets DiffMode field to given value.

### HasDiffMode

`func (o *InlineObject46) HasDiffMode() bool`

HasDiffMode returns a boolean if a field has been set.

### GetExtraVars

`func (o *InlineObject46) GetExtraVars() string`

GetExtraVars returns the ExtraVars field if non-nil, zero value otherwise.

### GetExtraVarsOk

`func (o *InlineObject46) GetExtraVarsOk() (*string, bool)`

GetExtraVarsOk returns a tuple with the ExtraVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVars

`func (o *InlineObject46) SetExtraVars(v string)`

SetExtraVars sets ExtraVars field to given value.

### HasExtraVars

`func (o *InlineObject46) HasExtraVars() bool`

HasExtraVars returns a boolean if a field has been set.

### GetForceHandlers

`func (o *InlineObject46) GetForceHandlers() bool`

GetForceHandlers returns the ForceHandlers field if non-nil, zero value otherwise.

### GetForceHandlersOk

`func (o *InlineObject46) GetForceHandlersOk() (*bool, bool)`

GetForceHandlersOk returns a tuple with the ForceHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHandlers

`func (o *InlineObject46) SetForceHandlers(v bool)`

SetForceHandlers sets ForceHandlers field to given value.

### HasForceHandlers

`func (o *InlineObject46) HasForceHandlers() bool`

HasForceHandlers returns a boolean if a field has been set.

### GetForks

`func (o *InlineObject46) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *InlineObject46) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *InlineObject46) SetForks(v int32)`

SetForks sets Forks field to given value.

### HasForks

`func (o *InlineObject46) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetHostConfigKey

`func (o *InlineObject46) GetHostConfigKey() string`

GetHostConfigKey returns the HostConfigKey field if non-nil, zero value otherwise.

### GetHostConfigKeyOk

`func (o *InlineObject46) GetHostConfigKeyOk() (*string, bool)`

GetHostConfigKeyOk returns a tuple with the HostConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostConfigKey

`func (o *InlineObject46) SetHostConfigKey(v string)`

SetHostConfigKey sets HostConfigKey field to given value.

### HasHostConfigKey

`func (o *InlineObject46) HasHostConfigKey() bool`

HasHostConfigKey returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject46) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject46) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject46) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject46) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetJobSliceCount

`func (o *InlineObject46) GetJobSliceCount() int32`

GetJobSliceCount returns the JobSliceCount field if non-nil, zero value otherwise.

### GetJobSliceCountOk

`func (o *InlineObject46) GetJobSliceCountOk() (*int32, bool)`

GetJobSliceCountOk returns a tuple with the JobSliceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSliceCount

`func (o *InlineObject46) SetJobSliceCount(v int32)`

SetJobSliceCount sets JobSliceCount field to given value.

### HasJobSliceCount

`func (o *InlineObject46) HasJobSliceCount() bool`

HasJobSliceCount returns a boolean if a field has been set.

### GetJobTags

`func (o *InlineObject46) GetJobTags() string`

GetJobTags returns the JobTags field if non-nil, zero value otherwise.

### GetJobTagsOk

`func (o *InlineObject46) GetJobTagsOk() (*string, bool)`

GetJobTagsOk returns a tuple with the JobTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTags

`func (o *InlineObject46) SetJobTags(v string)`

SetJobTags sets JobTags field to given value.

### HasJobTags

`func (o *InlineObject46) HasJobTags() bool`

HasJobTags returns a boolean if a field has been set.

### GetJobType

`func (o *InlineObject46) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *InlineObject46) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *InlineObject46) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *InlineObject46) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLimit

`func (o *InlineObject46) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineObject46) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineObject46) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineObject46) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *InlineObject46) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject46) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject46) SetName(v string)`

SetName sets Name field to given value.


### GetPlaybook

`func (o *InlineObject46) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *InlineObject46) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *InlineObject46) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *InlineObject46) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetProject

`func (o *InlineObject46) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InlineObject46) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InlineObject46) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *InlineObject46) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetScmBranch

`func (o *InlineObject46) GetScmBranch() string`

GetScmBranch returns the ScmBranch field if non-nil, zero value otherwise.

### GetScmBranchOk

`func (o *InlineObject46) GetScmBranchOk() (*string, bool)`

GetScmBranchOk returns a tuple with the ScmBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmBranch

`func (o *InlineObject46) SetScmBranch(v string)`

SetScmBranch sets ScmBranch field to given value.

### HasScmBranch

`func (o *InlineObject46) HasScmBranch() bool`

HasScmBranch returns a boolean if a field has been set.

### GetSkipTags

`func (o *InlineObject46) GetSkipTags() string`

GetSkipTags returns the SkipTags field if non-nil, zero value otherwise.

### GetSkipTagsOk

`func (o *InlineObject46) GetSkipTagsOk() (*string, bool)`

GetSkipTagsOk returns a tuple with the SkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTags

`func (o *InlineObject46) SetSkipTags(v string)`

SetSkipTags sets SkipTags field to given value.

### HasSkipTags

`func (o *InlineObject46) HasSkipTags() bool`

HasSkipTags returns a boolean if a field has been set.

### GetStartAtTask

`func (o *InlineObject46) GetStartAtTask() string`

GetStartAtTask returns the StartAtTask field if non-nil, zero value otherwise.

### GetStartAtTaskOk

`func (o *InlineObject46) GetStartAtTaskOk() (*string, bool)`

GetStartAtTaskOk returns a tuple with the StartAtTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAtTask

`func (o *InlineObject46) SetStartAtTask(v string)`

SetStartAtTask sets StartAtTask field to given value.

### HasStartAtTask

`func (o *InlineObject46) HasStartAtTask() bool`

HasStartAtTask returns a boolean if a field has been set.

### GetSurveyEnabled

`func (o *InlineObject46) GetSurveyEnabled() bool`

GetSurveyEnabled returns the SurveyEnabled field if non-nil, zero value otherwise.

### GetSurveyEnabledOk

`func (o *InlineObject46) GetSurveyEnabledOk() (*bool, bool)`

GetSurveyEnabledOk returns a tuple with the SurveyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyEnabled

`func (o *InlineObject46) SetSurveyEnabled(v bool)`

SetSurveyEnabled sets SurveyEnabled field to given value.

### HasSurveyEnabled

`func (o *InlineObject46) HasSurveyEnabled() bool`

HasSurveyEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineObject46) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject46) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject46) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject46) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUseFactCache

`func (o *InlineObject46) GetUseFactCache() bool`

GetUseFactCache returns the UseFactCache field if non-nil, zero value otherwise.

### GetUseFactCacheOk

`func (o *InlineObject46) GetUseFactCacheOk() (*bool, bool)`

GetUseFactCacheOk returns a tuple with the UseFactCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFactCache

`func (o *InlineObject46) SetUseFactCache(v bool)`

SetUseFactCache sets UseFactCache field to given value.

### HasUseFactCache

`func (o *InlineObject46) HasUseFactCache() bool`

HasUseFactCache returns a boolean if a field has been set.

### GetVerbosity

`func (o *InlineObject46) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *InlineObject46) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *InlineObject46) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *InlineObject46) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.

### GetWebhookCredential

`func (o *InlineObject46) GetWebhookCredential() int32`

GetWebhookCredential returns the WebhookCredential field if non-nil, zero value otherwise.

### GetWebhookCredentialOk

`func (o *InlineObject46) GetWebhookCredentialOk() (*int32, bool)`

GetWebhookCredentialOk returns a tuple with the WebhookCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCredential

`func (o *InlineObject46) SetWebhookCredential(v int32)`

SetWebhookCredential sets WebhookCredential field to given value.

### HasWebhookCredential

`func (o *InlineObject46) HasWebhookCredential() bool`

HasWebhookCredential returns a boolean if a field has been set.

### GetWebhookService

`func (o *InlineObject46) GetWebhookService() string`

GetWebhookService returns the WebhookService field if non-nil, zero value otherwise.

### GetWebhookServiceOk

`func (o *InlineObject46) GetWebhookServiceOk() (*string, bool)`

GetWebhookServiceOk returns a tuple with the WebhookService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookService

`func (o *InlineObject46) SetWebhookService(v string)`

SetWebhookService sets WebhookService field to given value.

### HasWebhookService

`func (o *InlineObject46) HasWebhookService() bool`

HasWebhookService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


