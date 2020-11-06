# InlineObject70

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationGrantType** | **string** | The Grant type the user must use for acquire tokens for this application. | 
**ClientType** | **string** | Set to Public or Confidential depending on how secure the client device is. | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Organization** | **int32** | Organization containing this application. | 
**RedirectUris** | Pointer to **string** | Allowed URIs list, space separated | [optional] 
**SkipAuthorization** | Pointer to **bool** | Set True to skip authorization step for completely trusted applications. | [optional] 

## Methods

### NewInlineObject70

`func NewInlineObject70(authorizationGrantType string, clientType string, name string, organization int32, ) *InlineObject70`

NewInlineObject70 instantiates a new InlineObject70 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject70WithDefaults

`func NewInlineObject70WithDefaults() *InlineObject70`

NewInlineObject70WithDefaults instantiates a new InlineObject70 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationGrantType

`func (o *InlineObject70) GetAuthorizationGrantType() string`

GetAuthorizationGrantType returns the AuthorizationGrantType field if non-nil, zero value otherwise.

### GetAuthorizationGrantTypeOk

`func (o *InlineObject70) GetAuthorizationGrantTypeOk() (*string, bool)`

GetAuthorizationGrantTypeOk returns a tuple with the AuthorizationGrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGrantType

`func (o *InlineObject70) SetAuthorizationGrantType(v string)`

SetAuthorizationGrantType sets AuthorizationGrantType field to given value.


### GetClientType

`func (o *InlineObject70) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *InlineObject70) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *InlineObject70) SetClientType(v string)`

SetClientType sets ClientType field to given value.


### GetDescription

`func (o *InlineObject70) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject70) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject70) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject70) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *InlineObject70) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject70) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject70) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *InlineObject70) GetOrganization() int32`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InlineObject70) GetOrganizationOk() (*int32, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InlineObject70) SetOrganization(v int32)`

SetOrganization sets Organization field to given value.


### GetRedirectUris

`func (o *InlineObject70) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *InlineObject70) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *InlineObject70) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *InlineObject70) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetSkipAuthorization

`func (o *InlineObject70) GetSkipAuthorization() bool`

GetSkipAuthorization returns the SkipAuthorization field if non-nil, zero value otherwise.

### GetSkipAuthorizationOk

`func (o *InlineObject70) GetSkipAuthorizationOk() (*bool, bool)`

GetSkipAuthorizationOk returns a tuple with the SkipAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAuthorization

`func (o *InlineObject70) SetSkipAuthorization(v bool)`

SetSkipAuthorization sets SkipAuthorization field to given value.

### HasSkipAuthorization

`func (o *InlineObject70) HasSkipAuthorization() bool`

HasSkipAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


