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

// InlineObject33 struct for InlineObject33
type InlineObject33 struct {
	// 
	Name string `json:"name"`
	// Organization this label belongs to.
	Organization int32 `json:"organization"`
}

// NewInlineObject33 instantiates a new InlineObject33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject33(name string, organization int32, ) *InlineObject33 {
	this := InlineObject33{}
	this.Name = name
	this.Organization = organization
	return &this
}

// NewInlineObject33WithDefaults instantiates a new InlineObject33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject33WithDefaults() *InlineObject33 {
	this := InlineObject33{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject33) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject33) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject33) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value
func (o *InlineObject33) GetOrganization() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *InlineObject33) GetOrganizationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *InlineObject33) SetOrganization(v int32) {
	o.Organization = v
}

func (o InlineObject33) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject33 struct {
	value *InlineObject33
	isSet bool
}

func (v NullableInlineObject33) Get() *InlineObject33 {
	return v.value
}

func (v *NullableInlineObject33) Set(val *InlineObject33) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject33) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject33(val *InlineObject33) *NullableInlineObject33 {
	return &NullableInlineObject33{value: val, isSet: true}
}

func (v NullableInlineObject33) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


