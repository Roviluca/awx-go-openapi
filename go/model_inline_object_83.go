/*
 * Ansible Tower API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineObject83 struct for InlineObject83
type InlineObject83 struct {
	// 
	Description *string `json:"description,omitempty"`
	// Optional custom messages for notification template.
	Messages *string `json:"messages,omitempty"`
	// 
	Name string `json:"name"`
	// 
	NotificationConfiguration *string `json:"notification_configuration,omitempty"`
	// 
	NotificationType string `json:"notification_type"`
	// 
	Organization int32 `json:"organization"`
}

// NewInlineObject83 instantiates a new InlineObject83 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject83(name string, notificationType string, organization int32, ) *InlineObject83 {
	this := InlineObject83{}
	this.Name = name
	this.NotificationType = notificationType
	this.Organization = organization
	return &this
}

// NewInlineObject83WithDefaults instantiates a new InlineObject83 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject83WithDefaults() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject83) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject83) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject83) SetDescription(v string) {
	o.Description = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *InlineObject83) GetMessages() string {
	if o == nil || o.Messages == nil {
		var ret string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetMessagesOk() (*string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *InlineObject83) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given string and assigns it to the Messages field.
func (o *InlineObject83) SetMessages(v string) {
	o.Messages = &v
}

// GetName returns the Name field value
func (o *InlineObject83) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject83) SetName(v string) {
	o.Name = v
}

// GetNotificationConfiguration returns the NotificationConfiguration field value if set, zero value otherwise.
func (o *InlineObject83) GetNotificationConfiguration() string {
	if o == nil || o.NotificationConfiguration == nil {
		var ret string
		return ret
	}
	return *o.NotificationConfiguration
}

// GetNotificationConfigurationOk returns a tuple with the NotificationConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetNotificationConfigurationOk() (*string, bool) {
	if o == nil || o.NotificationConfiguration == nil {
		return nil, false
	}
	return o.NotificationConfiguration, true
}

// HasNotificationConfiguration returns a boolean if a field has been set.
func (o *InlineObject83) HasNotificationConfiguration() bool {
	if o != nil && o.NotificationConfiguration != nil {
		return true
	}

	return false
}

// SetNotificationConfiguration gets a reference to the given string and assigns it to the NotificationConfiguration field.
func (o *InlineObject83) SetNotificationConfiguration(v string) {
	o.NotificationConfiguration = &v
}

// GetNotificationType returns the NotificationType field value
func (o *InlineObject83) GetNotificationType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetNotificationTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *InlineObject83) SetNotificationType(v string) {
	o.NotificationType = v
}

// GetOrganization returns the Organization field value
func (o *InlineObject83) GetOrganization() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetOrganizationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *InlineObject83) SetOrganization(v int32) {
	o.Organization = v
}

func (o InlineObject83) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NotificationConfiguration != nil {
		toSerialize["notification_configuration"] = o.NotificationConfiguration
	}
	if true {
		toSerialize["notification_type"] = o.NotificationType
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject83 struct {
	value *InlineObject83
	isSet bool
}

func (v NullableInlineObject83) Get() *InlineObject83 {
	return v.value
}

func (v *NullableInlineObject83) Set(val *InlineObject83) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject83) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject83) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject83(val *InlineObject83) *NullableInlineObject83 {
	return &NullableInlineObject83{value: val, isSet: true}
}

func (v NullableInlineObject83) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject83) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


