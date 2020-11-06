# InlineObject28

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | **int32** | Specify the type of credential you want to create. Refer to the Ansible Tower documentation for details on each type. | 
**Description** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to **map[string]interface{}** | Enter inputs using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**Name** | **string** |  | 
**Organization** | Pointer to **int32** |  | [optional] 

## Methods

### NewInlineObject28

`func NewInlineObject28(credentialType int32, name string, ) *InlineObject28`

NewInlineObject28 instantiates a new InlineObject28 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject28WithDefaults

`func NewInlineObject28WithDefaults() *InlineObject28`

NewInlineObject28WithDefaults instantiates a new InlineObject28 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *InlineObject28) GetCredentialType() int32`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *InlineObject28) GetCredentialTypeOk() (*int32, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *InlineObject28) SetCredentialType(v int32)`

SetCredentialType sets CredentialType field to given value.


### GetDescription

`func (o *InlineObject28) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject28) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject28) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject28) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputs

`func (o *InlineObject28) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *InlineObject28) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *InlineObject28) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *InlineObject28) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetName

`func (o *InlineObject28) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject28) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject28) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *InlineObject28) GetOrganization() int32`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InlineObject28) GetOrganizationOk() (*int32, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InlineObject28) SetOrganization(v int32)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InlineObject28) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


