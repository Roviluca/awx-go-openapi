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

// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	// 
	BecomeEnabled *bool `json:"become_enabled,omitempty"`
	// 
	Credential *int32 `json:"credential,omitempty"`
	// 
	DiffMode *bool `json:"diff_mode,omitempty"`
	// 
	ExtraVars *string `json:"extra_vars,omitempty"`
	// 
	Forks *int32 `json:"forks,omitempty"`
	// 
	Inventory *int32 `json:"inventory,omitempty"`
	// 
	JobType *string `json:"job_type,omitempty"`
	// 
	Limit *string `json:"limit,omitempty"`
	// 
	ModuleArgs *string `json:"module_args,omitempty"`
	// 
	ModuleName *string `json:"module_name,omitempty"`
	// 
	Verbosity *string `json:"verbosity,omitempty"`
}

// NewInlineObject7 instantiates a new InlineObject7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject7() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// NewInlineObject7WithDefaults instantiates a new InlineObject7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject7WithDefaults() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// GetBecomeEnabled returns the BecomeEnabled field value if set, zero value otherwise.
func (o *InlineObject7) GetBecomeEnabled() bool {
	if o == nil || o.BecomeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BecomeEnabled
}

// GetBecomeEnabledOk returns a tuple with the BecomeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetBecomeEnabledOk() (*bool, bool) {
	if o == nil || o.BecomeEnabled == nil {
		return nil, false
	}
	return o.BecomeEnabled, true
}

// HasBecomeEnabled returns a boolean if a field has been set.
func (o *InlineObject7) HasBecomeEnabled() bool {
	if o != nil && o.BecomeEnabled != nil {
		return true
	}

	return false
}

// SetBecomeEnabled gets a reference to the given bool and assigns it to the BecomeEnabled field.
func (o *InlineObject7) SetBecomeEnabled(v bool) {
	o.BecomeEnabled = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *InlineObject7) GetCredential() int32 {
	if o == nil || o.Credential == nil {
		var ret int32
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetCredentialOk() (*int32, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *InlineObject7) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given int32 and assigns it to the Credential field.
func (o *InlineObject7) SetCredential(v int32) {
	o.Credential = &v
}

// GetDiffMode returns the DiffMode field value if set, zero value otherwise.
func (o *InlineObject7) GetDiffMode() bool {
	if o == nil || o.DiffMode == nil {
		var ret bool
		return ret
	}
	return *o.DiffMode
}

// GetDiffModeOk returns a tuple with the DiffMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetDiffModeOk() (*bool, bool) {
	if o == nil || o.DiffMode == nil {
		return nil, false
	}
	return o.DiffMode, true
}

// HasDiffMode returns a boolean if a field has been set.
func (o *InlineObject7) HasDiffMode() bool {
	if o != nil && o.DiffMode != nil {
		return true
	}

	return false
}

// SetDiffMode gets a reference to the given bool and assigns it to the DiffMode field.
func (o *InlineObject7) SetDiffMode(v bool) {
	o.DiffMode = &v
}

// GetExtraVars returns the ExtraVars field value if set, zero value otherwise.
func (o *InlineObject7) GetExtraVars() string {
	if o == nil || o.ExtraVars == nil {
		var ret string
		return ret
	}
	return *o.ExtraVars
}

// GetExtraVarsOk returns a tuple with the ExtraVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetExtraVarsOk() (*string, bool) {
	if o == nil || o.ExtraVars == nil {
		return nil, false
	}
	return o.ExtraVars, true
}

// HasExtraVars returns a boolean if a field has been set.
func (o *InlineObject7) HasExtraVars() bool {
	if o != nil && o.ExtraVars != nil {
		return true
	}

	return false
}

// SetExtraVars gets a reference to the given string and assigns it to the ExtraVars field.
func (o *InlineObject7) SetExtraVars(v string) {
	o.ExtraVars = &v
}

// GetForks returns the Forks field value if set, zero value otherwise.
func (o *InlineObject7) GetForks() int32 {
	if o == nil || o.Forks == nil {
		var ret int32
		return ret
	}
	return *o.Forks
}

// GetForksOk returns a tuple with the Forks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetForksOk() (*int32, bool) {
	if o == nil || o.Forks == nil {
		return nil, false
	}
	return o.Forks, true
}

// HasForks returns a boolean if a field has been set.
func (o *InlineObject7) HasForks() bool {
	if o != nil && o.Forks != nil {
		return true
	}

	return false
}

// SetForks gets a reference to the given int32 and assigns it to the Forks field.
func (o *InlineObject7) SetForks(v int32) {
	o.Forks = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *InlineObject7) GetInventory() int32 {
	if o == nil || o.Inventory == nil {
		var ret int32
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetInventoryOk() (*int32, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *InlineObject7) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given int32 and assigns it to the Inventory field.
func (o *InlineObject7) SetInventory(v int32) {
	o.Inventory = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *InlineObject7) GetJobType() string {
	if o == nil || o.JobType == nil {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetJobTypeOk() (*string, bool) {
	if o == nil || o.JobType == nil {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *InlineObject7) HasJobType() bool {
	if o != nil && o.JobType != nil {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *InlineObject7) SetJobType(v string) {
	o.JobType = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *InlineObject7) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *InlineObject7) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *InlineObject7) SetLimit(v string) {
	o.Limit = &v
}

// GetModuleArgs returns the ModuleArgs field value if set, zero value otherwise.
func (o *InlineObject7) GetModuleArgs() string {
	if o == nil || o.ModuleArgs == nil {
		var ret string
		return ret
	}
	return *o.ModuleArgs
}

// GetModuleArgsOk returns a tuple with the ModuleArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetModuleArgsOk() (*string, bool) {
	if o == nil || o.ModuleArgs == nil {
		return nil, false
	}
	return o.ModuleArgs, true
}

// HasModuleArgs returns a boolean if a field has been set.
func (o *InlineObject7) HasModuleArgs() bool {
	if o != nil && o.ModuleArgs != nil {
		return true
	}

	return false
}

// SetModuleArgs gets a reference to the given string and assigns it to the ModuleArgs field.
func (o *InlineObject7) SetModuleArgs(v string) {
	o.ModuleArgs = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *InlineObject7) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *InlineObject7) HasModuleName() bool {
	if o != nil && o.ModuleName != nil {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *InlineObject7) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetVerbosity returns the Verbosity field value if set, zero value otherwise.
func (o *InlineObject7) GetVerbosity() string {
	if o == nil || o.Verbosity == nil {
		var ret string
		return ret
	}
	return *o.Verbosity
}

// GetVerbosityOk returns a tuple with the Verbosity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetVerbosityOk() (*string, bool) {
	if o == nil || o.Verbosity == nil {
		return nil, false
	}
	return o.Verbosity, true
}

// HasVerbosity returns a boolean if a field has been set.
func (o *InlineObject7) HasVerbosity() bool {
	if o != nil && o.Verbosity != nil {
		return true
	}

	return false
}

// SetVerbosity gets a reference to the given string and assigns it to the Verbosity field.
func (o *InlineObject7) SetVerbosity(v string) {
	o.Verbosity = &v
}

func (o InlineObject7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BecomeEnabled != nil {
		toSerialize["become_enabled"] = o.BecomeEnabled
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	if o.DiffMode != nil {
		toSerialize["diff_mode"] = o.DiffMode
	}
	if o.ExtraVars != nil {
		toSerialize["extra_vars"] = o.ExtraVars
	}
	if o.Forks != nil {
		toSerialize["forks"] = o.Forks
	}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	if o.JobType != nil {
		toSerialize["job_type"] = o.JobType
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.ModuleArgs != nil {
		toSerialize["module_args"] = o.ModuleArgs
	}
	if o.ModuleName != nil {
		toSerialize["module_name"] = o.ModuleName
	}
	if o.Verbosity != nil {
		toSerialize["verbosity"] = o.Verbosity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject7 struct {
	value *InlineObject7
	isSet bool
}

func (v NullableInlineObject7) Get() *InlineObject7 {
	return v.value
}

func (v *NullableInlineObject7) Set(val *InlineObject7) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject7) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject7(val *InlineObject7) *NullableInlineObject7 {
	return &NullableInlineObject7{value: val, isSet: true}
}

func (v NullableInlineObject7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


