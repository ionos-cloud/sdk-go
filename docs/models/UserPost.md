# UserPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**UserPropertiesPost**](UserPropertiesPost.md) |  | |

## Methods

### NewUserPost

`func NewUserPost(properties UserPropertiesPost, ) *UserPost`

NewUserPost instantiates a new UserPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostWithDefaults

`func NewUserPostWithDefaults() *UserPost`

NewUserPostWithDefaults instantiates a new UserPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *UserPost) GetProperties() UserPropertiesPost`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserPost) GetPropertiesOk() (*UserPropertiesPost, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserPost) SetProperties(v UserPropertiesPost)`

SetProperties sets Properties field to given value.




