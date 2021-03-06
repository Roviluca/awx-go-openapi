/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	//
	BecomeEnabled bool `json:"become_enabled,omitempty"`
	//
	Credential int32 `json:"credential,omitempty"`
	//
	DiffMode bool `json:"diff_mode,omitempty"`
	//
	ExtraVars string `json:"extra_vars,omitempty"`
	//
	Forks int32 `json:"forks,omitempty"`
	//
	Inventory int32 `json:"inventory,omitempty"`
	//
	JobType string `json:"job_type,omitempty"`
	//
	Limit string `json:"limit,omitempty"`
	//
	ModuleArgs string `json:"module_args,omitempty"`
	//
	ModuleName string `json:"module_name,omitempty"`
	//
	Verbosity string `json:"verbosity,omitempty"`
}
