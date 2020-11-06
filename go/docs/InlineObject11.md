# InlineObject11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Is this host online and available for running jobs? | [optional] 
**InstanceId** | Pointer to **string** | The value used by the remote inventory source to uniquely identify the host | [optional] 
**Inventory** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **string** | Host variables in JSON or YAML format. | [optional] 

## Methods

### NewInlineObject11

`func NewInlineObject11() *InlineObject11`

NewInlineObject11 instantiates a new InlineObject11 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject11WithDefaults

`func NewInlineObject11WithDefaults() *InlineObject11`

NewInlineObject11WithDefaults instantiates a new InlineObject11 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InlineObject11) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject11) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject11) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject11) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject11) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject11) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject11) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject11) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstanceId

`func (o *InlineObject11) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InlineObject11) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InlineObject11) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *InlineObject11) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInventory

`func (o *InlineObject11) GetInventory() int32`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InlineObject11) GetInventoryOk() (*int32, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InlineObject11) SetInventory(v int32)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InlineObject11) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetName

`func (o *InlineObject11) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject11) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject11) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject11) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariables

`func (o *InlineObject11) GetVariables() string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *InlineObject11) GetVariablesOk() (*string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *InlineObject11) SetVariables(v string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *InlineObject11) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


