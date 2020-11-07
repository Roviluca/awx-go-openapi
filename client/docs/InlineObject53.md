# InlineObject53

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOverride** | **bool** | Allow changing the SCM branch or revision in a job template that uses this project. | [optional] 
**Credential** | **int32** |  | [optional] 
**CustomVirtualenv** | **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] 
**Description** | **string** |  | [optional] 
**LocalPath** | **string** | Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. | [optional] 
**Name** | **string** |  | 
**Organization** | **int32** | The organization used to determine access to this template. | [optional] 
**ScmBranch** | **string** | Specific branch, tag or commit to checkout. | [optional] 
**ScmClean** | **bool** | Discard any local changes before syncing the project. | [optional] 
**ScmDeleteOnUpdate** | **bool** | Delete the project before syncing. | [optional] 
**ScmRefspec** | **string** | For git projects, an additional refspec to fetch. | [optional] 
**ScmType** | **string** | Specifies the source control system used to store the project. | [optional] 
**ScmUpdateCacheTimeout** | **int32** | The number of seconds after the last project update ran that a new project update will be launched as a job dependency. | [optional] 
**ScmUpdateOnLaunch** | **bool** | Update the project when a job is launched that uses the project. | [optional] 
**ScmUrl** | **string** | The location where the project is stored. | [optional] 
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


