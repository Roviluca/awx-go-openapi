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

// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
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

// NewInlineObject12 instantiates a new InlineObject12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject12() *InlineObject12 {
	this := InlineObject12{}
	return &this
}

// NewInlineObject12WithDefaults instantiates a new InlineObject12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject12WithDefaults() *InlineObject12 {
	this := InlineObject12{}
	return &this
}

// GetBecomeEnabled returns the BecomeEnabled field value if set, zero value otherwise.
func (o *InlineObject12) GetBecomeEnabled() bool {
	if o == nil || o.BecomeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BecomeEnabled
}

// GetBecomeEnabledOk returns a tuple with the BecomeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetBecomeEnabledOk() (*bool, bool) {
	if o == nil || o.BecomeEnabled == nil {
		return nil, false
	}
	return o.BecomeEnabled, true
}

// HasBecomeEnabled returns a boolean if a field has been set.
func (o *InlineObject12) HasBecomeEnabled() bool {
	if o != nil && o.BecomeEnabled != nil {
		return true
	}

	return false
}

// SetBecomeEnabled gets a reference to the given bool and assigns it to the BecomeEnabled field.
func (o *InlineObject12) SetBecomeEnabled(v bool) {
	o.BecomeEnabled = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *InlineObject12) GetCredential() int32 {
	if o == nil || o.Credential == nil {
		var ret int32
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetCredentialOk() (*int32, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *InlineObject12) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given int32 and assigns it to the Credential field.
func (o *InlineObject12) SetCredential(v int32) {
	o.Credential = &v
}

// GetDiffMode returns the DiffMode field value if set, zero value otherwise.
func (o *InlineObject12) GetDiffMode() bool {
	if o == nil || o.DiffMode == nil {
		var ret bool
		return ret
	}
	return *o.DiffMode
}

// GetDiffModeOk returns a tuple with the DiffMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetDiffModeOk() (*bool, bool) {
	if o == nil || o.DiffMode == nil {
		return nil, false
	}
	return o.DiffMode, true
}

// HasDiffMode returns a boolean if a field has been set.
func (o *InlineObject12) HasDiffMode() bool {
	if o != nil && o.DiffMode != nil {
		return true
	}

	return false
}

// SetDiffMode gets a reference to the given bool and assigns it to the DiffMode field.
func (o *InlineObject12) SetDiffMode(v bool) {
	o.DiffMode = &v
}

// GetExtraVars returns the ExtraVars field value if set, zero value otherwise.
func (o *InlineObject12) GetExtraVars() string {
	if o == nil || o.ExtraVars == nil {
		var ret string
		return ret
	}
	return *o.ExtraVars
}

// GetExtraVarsOk returns a tuple with the ExtraVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetExtraVarsOk() (*string, bool) {
	if o == nil || o.ExtraVars == nil {
		return nil, false
	}
	return o.ExtraVars, true
}

// HasExtraVars returns a boolean if a field has been set.
func (o *InlineObject12) HasExtraVars() bool {
	if o != nil && o.ExtraVars != nil {
		return true
	}

	return false
}

// SetExtraVars gets a reference to the given string and assigns it to the ExtraVars field.
func (o *InlineObject12) SetExtraVars(v string) {
	o.ExtraVars = &v
}

// GetForks returns the Forks field value if set, zero value otherwise.
func (o *InlineObject12) GetForks() int32 {
	if o == nil || o.Forks == nil {
		var ret int32
		return ret
	}
	return *o.Forks
}

// GetForksOk returns a tuple with the Forks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetForksOk() (*int32, bool) {
	if o == nil || o.Forks == nil {
		return nil, false
	}
	return o.Forks, true
}

// HasForks returns a boolean if a field has been set.
func (o *InlineObject12) HasForks() bool {
	if o != nil && o.Forks != nil {
		return true
	}

	return false
}

// SetForks gets a reference to the given int32 and assigns it to the Forks field.
func (o *InlineObject12) SetForks(v int32) {
	o.Forks = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *InlineObject12) GetInventory() int32 {
	if o == nil || o.Inventory == nil {
		var ret int32
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetInventoryOk() (*int32, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *InlineObject12) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given int32 and assigns it to the Inventory field.
func (o *InlineObject12) SetInventory(v int32) {
	o.Inventory = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *InlineObject12) GetJobType() string {
	if o == nil || o.JobType == nil {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetJobTypeOk() (*string, bool) {
	if o == nil || o.JobType == nil {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *InlineObject12) HasJobType() bool {
	if o != nil && o.JobType != nil {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *InlineObject12) SetJobType(v string) {
	o.JobType = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *InlineObject12) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *InlineObject12) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *InlineObject12) SetLimit(v string) {
	o.Limit = &v
}

// GetModuleArgs returns the ModuleArgs field value if set, zero value otherwise.
func (o *InlineObject12) GetModuleArgs() string {
	if o == nil || o.ModuleArgs == nil {
		var ret string
		return ret
	}
	return *o.ModuleArgs
}

// GetModuleArgsOk returns a tuple with the ModuleArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetModuleArgsOk() (*string, bool) {
	if o == nil || o.ModuleArgs == nil {
		return nil, false
	}
	return o.ModuleArgs, true
}

// HasModuleArgs returns a boolean if a field has been set.
func (o *InlineObject12) HasModuleArgs() bool {
	if o != nil && o.ModuleArgs != nil {
		return true
	}

	return false
}

// SetModuleArgs gets a reference to the given string and assigns it to the ModuleArgs field.
func (o *InlineObject12) SetModuleArgs(v string) {
	o.ModuleArgs = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *InlineObject12) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *InlineObject12) HasModuleName() bool {
	if o != nil && o.ModuleName != nil {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *InlineObject12) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetVerbosity returns the Verbosity field value if set, zero value otherwise.
func (o *InlineObject12) GetVerbosity() string {
	if o == nil || o.Verbosity == nil {
		var ret string
		return ret
	}
	return *o.Verbosity
}

// GetVerbosityOk returns a tuple with the Verbosity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetVerbosityOk() (*string, bool) {
	if o == nil || o.Verbosity == nil {
		return nil, false
	}
	return o.Verbosity, true
}

// HasVerbosity returns a boolean if a field has been set.
func (o *InlineObject12) HasVerbosity() bool {
	if o != nil && o.Verbosity != nil {
		return true
	}

	return false
}

// SetVerbosity gets a reference to the given string and assigns it to the Verbosity field.
func (o *InlineObject12) SetVerbosity(v string) {
	o.Verbosity = &v
}

func (o InlineObject12) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject12 struct {
	value *InlineObject12
	isSet bool
}

func (v NullableInlineObject12) Get() *InlineObject12 {
	return v.value
}

func (v *NullableInlineObject12) Set(val *InlineObject12) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject12) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject12(val *InlineObject12) *NullableInlineObject12 {
	return &NullableInlineObject12{value: val, isSet: true}
}

func (v NullableInlineObject12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


