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

// InlineObject13 struct for InlineObject13
type InlineObject13 struct {
	// Host variables in JSON or YAML format.
	Variables *string `json:"variables,omitempty"`
}

// NewInlineObject13 instantiates a new InlineObject13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject13() *InlineObject13 {
	this := InlineObject13{}
	return &this
}

// NewInlineObject13WithDefaults instantiates a new InlineObject13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject13WithDefaults() *InlineObject13 {
	this := InlineObject13{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *InlineObject13) GetVariables() string {
	if o == nil || o.Variables == nil {
		var ret string
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13) GetVariablesOk() (*string, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *InlineObject13) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given string and assigns it to the Variables field.
func (o *InlineObject13) SetVariables(v string) {
	o.Variables = &v
}

func (o InlineObject13) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject13 struct {
	value *InlineObject13
	isSet bool
}

func (v NullableInlineObject13) Get() *InlineObject13 {
	return v.value
}

func (v *NullableInlineObject13) Set(val *InlineObject13) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject13) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject13(val *InlineObject13) *NullableInlineObject13 {
	return &NullableInlineObject13{value: val, isSet: true}
}

func (v NullableInlineObject13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


