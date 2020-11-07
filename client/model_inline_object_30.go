/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject30 struct for InlineObject30
type InlineObject30 struct {
	// 
	Description string `json:"description,omitempty"`
	// Optional custom messages for notification template.
	Messages string `json:"messages,omitempty"`
	// 
	Name string `json:"name"`
	// 
	NotificationConfiguration string `json:"notification_configuration,omitempty"`
	// 
	NotificationType string `json:"notification_type"`
	// 
	Organization int32 `json:"organization"`
}
