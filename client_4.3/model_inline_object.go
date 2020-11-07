/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// The Grant type the user must use for acquire tokens for this application.
	AuthorizationGrantType string `json:"authorization_grant_type"`
	// Set to Public or Confidential depending on how secure the client device is.
	ClientType string `json:"client_type"`
	// 
	Description string `json:"description,omitempty"`
	// 
	Name string `json:"name"`
	// Organization containing this application.
	Organization int32 `json:"organization"`
	// Allowed URIs list, space separated
	RedirectUris string `json:"redirect_uris,omitempty"`
	// Set True to skip authorization step for completely trusted applications.
	SkipAuthorization bool `json:"skip_authorization,omitempty"`
}