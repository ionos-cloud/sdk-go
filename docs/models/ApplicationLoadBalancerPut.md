# ApplicationLoadBalancerPut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Properties** | [**ApplicationLoadBalancerProperties**](ApplicationLoadBalancerProperties.md) |  | |

## Methods

### NewApplicationLoadBalancerPut

`func NewApplicationLoadBalancerPut(properties ApplicationLoadBalancerProperties, ) *ApplicationLoadBalancerPut`

NewApplicationLoadBalancerPut instantiates a new ApplicationLoadBalancerPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerPutWithDefaults

`func NewApplicationLoadBalancerPutWithDefaults() *ApplicationLoadBalancerPut`

NewApplicationLoadBalancerPutWithDefaults instantiates a new ApplicationLoadBalancerPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationLoadBalancerPut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationLoadBalancerPut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationLoadBalancerPut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationLoadBalancerPut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationLoadBalancerPut) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationLoadBalancerPut) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationLoadBalancerPut) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationLoadBalancerPut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ApplicationLoadBalancerPut) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplicationLoadBalancerPut) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplicationLoadBalancerPut) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplicationLoadBalancerPut) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *ApplicationLoadBalancerPut) GetProperties() ApplicationLoadBalancerProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ApplicationLoadBalancerPut) GetPropertiesOk() (*ApplicationLoadBalancerProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ApplicationLoadBalancerPut) SetProperties(v ApplicationLoadBalancerProperties)`

SetProperties sets Properties field to given value.




