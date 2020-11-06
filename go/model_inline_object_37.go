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

// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	// 
	Name string `json:"name"`
	// Organization this label belongs to.
	Organization int32 `json:"organization"`
}

// NewInlineObject37 instantiates a new InlineObject37 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject37(name string, organization int32, ) *InlineObject37 {
	this := InlineObject37{}
	this.Name = name
	this.Organization = organization
	return &this
}

// NewInlineObject37WithDefaults instantiates a new InlineObject37 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject37WithDefaults() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject37) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject37) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value
func (o *InlineObject37) GetOrganization() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetOrganizationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *InlineObject37) SetOrganization(v int32) {
	o.Organization = v
}

func (o InlineObject37) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject37 struct {
	value *InlineObject37
	isSet bool
}

func (v NullableInlineObject37) Get() *InlineObject37 {
	return v.value
}

func (v *NullableInlineObject37) Set(val *InlineObject37) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject37) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject37) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject37(val *InlineObject37) *NullableInlineObject37 {
	return &NullableInlineObject37{value: val, isSet: true}
}

func (v NullableInlineObject37) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject37) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


