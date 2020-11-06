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

// InlineObject24 struct for InlineObject24
type InlineObject24 struct {
	//
	Description *string `json:"description,omitempty"`
	//
	Name string `json:"name"`
	// Organization owning this inventory script
	Organization int32 `json:"organization"`
	//
	Script string `json:"script"`
}

// NewInlineObject24 instantiates a new InlineObject24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject24(name string, organization int32, script string) *InlineObject24 {
	this := InlineObject24{}
	this.Name = name
	this.Organization = organization
	this.Script = script
	return &this
}

// NewInlineObject24WithDefaults instantiates a new InlineObject24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject24WithDefaults() *InlineObject24 {
	this := InlineObject24{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject24) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject24) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject24) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *InlineObject24) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject24) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value
func (o *InlineObject24) GetOrganization() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetOrganizationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *InlineObject24) SetOrganization(v int32) {
	o.Organization = v
}

// GetScript returns the Script field value
func (o *InlineObject24) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *InlineObject24) SetScript(v string) {
	o.Script = v
}

func (o InlineObject24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["script"] = o.Script
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject24 struct {
	value *InlineObject24
	isSet bool
}

func (v NullableInlineObject24) Get() *InlineObject24 {
	return v.value
}

func (v *NullableInlineObject24) Set(val *InlineObject24) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject24) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject24(val *InlineObject24) *NullableInlineObject24 {
	return &NullableInlineObject24{value: val, isSet: true}
}

func (v NullableInlineObject24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
