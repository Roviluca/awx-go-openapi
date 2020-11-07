# InlineObject46

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSimultaneous** | **bool** |  | [optional] 
**AskCredentialOnLaunch** | **bool** |  | [optional] 
**AskDiffModeOnLaunch** | **bool** |  | [optional] 
**AskInventoryOnLaunch** | **bool** |  | [optional] 
**AskJobTypeOnLaunch** | **bool** |  | [optional] 
**AskLimitOnLaunch** | **bool** |  | [optional] 
**AskScmBranchOnLaunch** | **bool** |  | [optional] 
**AskSkipTagsOnLaunch** | **bool** |  | [optional] 
**AskTagsOnLaunch** | **bool** |  | [optional] 
**AskVariablesOnLaunch** | **bool** |  | [optional] 
**AskVerbosityOnLaunch** | **bool** |  | [optional] 
**BecomeEnabled** | **bool** |  | [optional] 
**CustomVirtualenv** | **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] 
**Description** | **string** |  | [optional] 
**DiffMode** | **bool** | If enabled, textual changes made to any templated files on the host are shown in the standard output | [optional] 
**ExtraVars** | **string** |  | [optional] 
**ForceHandlers** | **bool** |  | [optional] 
**Forks** | **int32** |  | [optional] 
**HostConfigKey** | **string** |  | [optional] 
**Inventory** | **int32** |  | [optional] 
**JobSliceCount** | **int32** | The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1. | [optional] 
**JobTags** | **string** |  | [optional] 
**JobType** | **string** |  | [optional] 
**Limit** | **string** |  | [optional] 
**Name** | **string** |  | 
**Playbook** | **string** |  | [optional] 
**Project** | **string** |  | [optional] 
**ScmBranch** | **string** | Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true. | [optional] 
**SkipTags** | **string** |  | [optional] 
**StartAtTask** | **string** |  | [optional] 
**SurveyEnabled** | **bool** |  | [optional] 
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] 
**UseFactCache** | **bool** | If enabled, Tower will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible. | [optional] 
**Verbosity** | **string** |  | [optional] 
**WebhookCredential** | **int32** | Personal Access Token for posting back the status to the service API | [optional] 
**WebhookService** | **string** | Service that webhook requests will be accepted from | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


