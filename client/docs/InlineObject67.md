# InlineObject67

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**IsSuperuser** | Pointer to **bool** | Designates that this user has all permissions without explicitly assigning them. | [optional] 
**IsSystemAuditor** | Pointer to **bool** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | Write-only field used to change the password. | [optional] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 

## Methods

### NewInlineObject67

`func NewInlineObject67(username string, ) *InlineObject67`

NewInlineObject67 instantiates a new InlineObject67 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject67WithDefaults

`func NewInlineObject67WithDefaults() *InlineObject67`

NewInlineObject67WithDefaults instantiates a new InlineObject67 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InlineObject67) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject67) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject67) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineObject67) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *InlineObject67) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineObject67) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineObject67) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InlineObject67) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetIsSuperuser

`func (o *InlineObject67) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *InlineObject67) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *InlineObject67) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *InlineObject67) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetIsSystemAuditor

`func (o *InlineObject67) GetIsSystemAuditor() bool`

GetIsSystemAuditor returns the IsSystemAuditor field if non-nil, zero value otherwise.

### GetIsSystemAuditorOk

`func (o *InlineObject67) GetIsSystemAuditorOk() (*bool, bool)`

GetIsSystemAuditorOk returns a tuple with the IsSystemAuditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAuditor

`func (o *InlineObject67) SetIsSystemAuditor(v bool)`

SetIsSystemAuditor sets IsSystemAuditor field to given value.

### HasIsSystemAuditor

`func (o *InlineObject67) HasIsSystemAuditor() bool`

HasIsSystemAuditor returns a boolean if a field has been set.

### GetLastName

`func (o *InlineObject67) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineObject67) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineObject67) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineObject67) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject67) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject67) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject67) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject67) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *InlineObject67) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineObject67) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineObject67) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


