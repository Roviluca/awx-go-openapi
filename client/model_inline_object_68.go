/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InlineObject68 struct for InlineObject68
type InlineObject68 struct {
	//
	Description string `json:"description,omitempty"`
	// Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write'].
	Scope string `json:"scope,omitempty"`
}