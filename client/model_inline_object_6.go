/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	//
	Description string `json:"description,omitempty"`
	//
	Inventory int32 `json:"inventory,omitempty"`
	//
	Name string `json:"name,omitempty"`
	// Group variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}