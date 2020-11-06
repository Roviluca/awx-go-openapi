# InlineObject20

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**PodSpecOverride** | Pointer to **string** |  | [optional] 
**PolicyInstanceList** | Pointer to **[]string** | List of exact-match Instances that will be assigned to this group | [optional] 
**PolicyInstanceMinimum** | Pointer to **int32** | Static minimum number of Instances that will be automatically assign to this group when new instances come online. | [optional] 
**PolicyInstancePercentage** | Pointer to **int32** | Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. | [optional] 

## Methods

### NewInlineObject20

`func NewInlineObject20(name string, ) *InlineObject20`

NewInlineObject20 instantiates a new InlineObject20 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject20WithDefaults

`func NewInlineObject20WithDefaults() *InlineObject20`

NewInlineObject20WithDefaults instantiates a new InlineObject20 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *InlineObject20) GetCredential() int32`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InlineObject20) GetCredentialOk() (*int32, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InlineObject20) SetCredential(v int32)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InlineObject20) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetName

`func (o *InlineObject20) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject20) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject20) SetName(v string)`

SetName sets Name field to given value.


### GetPodSpecOverride

`func (o *InlineObject20) GetPodSpecOverride() string`

GetPodSpecOverride returns the PodSpecOverride field if non-nil, zero value otherwise.

### GetPodSpecOverrideOk

`func (o *InlineObject20) GetPodSpecOverrideOk() (*string, bool)`

GetPodSpecOverrideOk returns a tuple with the PodSpecOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSpecOverride

`func (o *InlineObject20) SetPodSpecOverride(v string)`

SetPodSpecOverride sets PodSpecOverride field to given value.

### HasPodSpecOverride

`func (o *InlineObject20) HasPodSpecOverride() bool`

HasPodSpecOverride returns a boolean if a field has been set.

### GetPolicyInstanceList

`func (o *InlineObject20) GetPolicyInstanceList() []string`

GetPolicyInstanceList returns the PolicyInstanceList field if non-nil, zero value otherwise.

### GetPolicyInstanceListOk

`func (o *InlineObject20) GetPolicyInstanceListOk() (*[]string, bool)`

GetPolicyInstanceListOk returns a tuple with the PolicyInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyInstanceList

`func (o *InlineObject20) SetPolicyInstanceList(v []string)`

SetPolicyInstanceList sets PolicyInstanceList field to given value.

### HasPolicyInstanceList

`func (o *InlineObject20) HasPolicyInstanceList() bool`

HasPolicyInstanceList returns a boolean if a field has been set.

### GetPolicyInstanceMinimum

`func (o *InlineObject20) GetPolicyInstanceMinimum() int32`

GetPolicyInstanceMinimum returns the PolicyInstanceMinimum field if non-nil, zero value otherwise.

### GetPolicyInstanceMinimumOk

`func (o *InlineObject20) GetPolicyInstanceMinimumOk() (*int32, bool)`

GetPolicyInstanceMinimumOk returns a tuple with the PolicyInstanceMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyInstanceMinimum

`func (o *InlineObject20) SetPolicyInstanceMinimum(v int32)`

SetPolicyInstanceMinimum sets PolicyInstanceMinimum field to given value.

### HasPolicyInstanceMinimum

`func (o *InlineObject20) HasPolicyInstanceMinimum() bool`

HasPolicyInstanceMinimum returns a boolean if a field has been set.

### GetPolicyInstancePercentage

`func (o *InlineObject20) GetPolicyInstancePercentage() int32`

GetPolicyInstancePercentage returns the PolicyInstancePercentage field if non-nil, zero value otherwise.

### GetPolicyInstancePercentageOk

`func (o *InlineObject20) GetPolicyInstancePercentageOk() (*int32, bool)`

GetPolicyInstancePercentageOk returns a tuple with the PolicyInstancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyInstancePercentage

`func (o *InlineObject20) SetPolicyInstancePercentage(v int32)`

SetPolicyInstancePercentage sets PolicyInstancePercentage field to given value.

### HasPolicyInstancePercentage

`func (o *InlineObject20) HasPolicyInstancePercentage() bool`

HasPolicyInstancePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


