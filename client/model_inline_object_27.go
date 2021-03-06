/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InlineObject27 struct for InlineObject27
type InlineObject27 struct {
	// Cloud credential to use for inventory updates.
	Credential int32 `json:"credential,omitempty"`
	// Local absolute file path containing a custom Python virtualenv to use
	CustomVirtualenv string `json:"custom_virtualenv,omitempty"`
	//
	Description string `json:"description,omitempty"`
	// Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2018-02-01T08:00:00.000000Z:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported into Tower. If the key is not found then the host will be enabled
	EnabledValue string `json:"enabled_value,omitempty"`
	// Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)
	EnabledVar string `json:"enabled_var,omitempty"`
	// Regex where only matching hosts will be imported into Tower.
	HostFilter string `json:"host_filter,omitempty"`
	//
	Inventory int32 `json:"inventory"`
	//
	Name string `json:"name"`
	// Overwrite local groups and hosts from remote inventory source.
	Overwrite bool `json:"overwrite,omitempty"`
	// Overwrite local variables from remote inventory source.
	OverwriteVars bool `json:"overwrite_vars,omitempty"`
	//
	Source string `json:"source,omitempty"`
	//
	SourcePath string `json:"source_path,omitempty"`
	// Project containing inventory file used as source.
	SourceProject string `json:"source_project,omitempty"`
	//
	SourceScript int32 `json:"source_script,omitempty"`
	// Inventory source variables in YAML or JSON format.
	SourceVars string `json:"source_vars,omitempty"`
	// The amount of time (in seconds) to run before the task is canceled.
	Timeout int32 `json:"timeout,omitempty"`
	//
	UpdateCacheTimeout int32 `json:"update_cache_timeout,omitempty"`
	//
	UpdateOnLaunch bool `json:"update_on_launch,omitempty"`
	//
	UpdateOnProjectUpdate bool `json:"update_on_project_update,omitempty"`
	//
	Verbosity string `json:"verbosity,omitempty"`
}
