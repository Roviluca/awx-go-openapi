# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Injectors** | Pointer to **map[string]interface{}** | Enter injectors using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**Inputs** | Pointer to **map[string]interface{}** | Enter inputs using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**Kind** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewInlineObject2

`func NewInlineObject2(kind string, name string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InlineObject2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInjectors

`func (o *InlineObject2) GetInjectors() map[string]interface{}`

GetInjectors returns the Injectors field if non-nil, zero value otherwise.

### GetInjectorsOk

`func (o *InlineObject2) GetInjectorsOk() (*map[string]interface{}, bool)`

GetInjectorsOk returns a tuple with the Injectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectors

`func (o *InlineObject2) SetInjectors(v map[string]interface{})`

SetInjectors sets Injectors field to given value.

### HasInjectors

`func (o *InlineObject2) HasInjectors() bool`

HasInjectors returns a boolean if a field has been set.

### GetInputs

`func (o *InlineObject2) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *InlineObject2) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *InlineObject2) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *InlineObject2) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetKind

`func (o *InlineObject2) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InlineObject2) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InlineObject2) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *InlineObject2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject2) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


