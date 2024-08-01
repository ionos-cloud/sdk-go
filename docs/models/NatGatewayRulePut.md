# NatGatewayRulePut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Properties** | [**NatGatewayRuleProperties**](NatGatewayRuleProperties.md) |  | |

## Methods

### NewNatGatewayRulePut

`func NewNatGatewayRulePut(properties NatGatewayRuleProperties, ) *NatGatewayRulePut`

NewNatGatewayRulePut instantiates a new NatGatewayRulePut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayRulePutWithDefaults

`func NewNatGatewayRulePutWithDefaults() *NatGatewayRulePut`

NewNatGatewayRulePutWithDefaults instantiates a new NatGatewayRulePut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NatGatewayRulePut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NatGatewayRulePut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NatGatewayRulePut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NatGatewayRulePut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NatGatewayRulePut) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NatGatewayRulePut) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NatGatewayRulePut) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *NatGatewayRulePut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NatGatewayRulePut) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NatGatewayRulePut) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NatGatewayRulePut) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NatGatewayRulePut) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *NatGatewayRulePut) GetProperties() NatGatewayRuleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NatGatewayRulePut) GetPropertiesOk() (*NatGatewayRuleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NatGatewayRulePut) SetProperties(v NatGatewayRuleProperties)`

SetProperties sets Properties field to given value.




