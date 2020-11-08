# InlineObject27

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | **int32** | Cloud credential to use for inventory updates. | [optional] 
**CustomVirtualenv** | **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] 
**Description** | **string** |  | [optional] 
**EnabledValue** | **string** | Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var&#x3D;\&quot;status.power_state\&quot;and enabled_value&#x3D;\&quot;powered_on\&quot; with host variables:{   \&quot;status\&quot;: {     \&quot;power_state\&quot;: \&quot;powered_on\&quot;,     \&quot;created\&quot;: \&quot;2018-02-01T08:00:00.000000Z:00\&quot;,     \&quot;healthy\&quot;: true    },    \&quot;name\&quot;: \&quot;foobar\&quot;,    \&quot;ip_address\&quot;: \&quot;192.168.2.1\&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported into Tower. If the key is not found then the host will be enabled | [optional] 
**EnabledVar** | **string** | Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \&quot;foo.bar\&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\&quot;foo\&quot;, {}).get(\&quot;bar\&quot;, default) | [optional] 
**HostFilter** | **string** | Regex where only matching hosts will be imported into Tower. | [optional] 
**Inventory** | **int32** |  | 
**Name** | **string** |  | 
**Overwrite** | **bool** | Overwrite local groups and hosts from remote inventory source. | [optional] 
**OverwriteVars** | **bool** | Overwrite local variables from remote inventory source. | [optional] 
**Source** | **string** |  | [optional] 
**SourcePath** | **string** |  | [optional] 
**SourceProject** | **string** | Project containing inventory file used as source. | [optional] 
**SourceScript** | **int32** |  | [optional] 
**SourceVars** | **string** | Inventory source variables in YAML or JSON format. | [optional] 
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] 
**UpdateCacheTimeout** | **int32** |  | [optional] 
**UpdateOnLaunch** | **bool** |  | [optional] 
**UpdateOnProjectUpdate** | **bool** |  | [optional] 
**Verbosity** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


