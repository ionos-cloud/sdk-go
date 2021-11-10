# NetworkLoadBalancerForwardingRule

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**NetworkLoadBalancerForwardingRuleProperties**](NetworkLoadBalancerForwardingRuleProperties.md) |  | |

## Methods

### NewNetworkLoadBalancerForwardingRule

`func NewNetworkLoadBalancerForwardingRule(properties NetworkLoadBalancerForwardingRuleProperties, ) *NetworkLoadBalancerForwardingRule`

NewNetworkLoadBalancerForwardingRule instantiates a new NetworkLoadBalancerForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerForwardingRuleWithDefaults

`func NewNetworkLoadBalancerForwardingRuleWithDefaults() *NetworkLoadBalancerForwardingRule`

NewNetworkLoadBalancerForwardingRuleWithDefaults instantiates a new NetworkLoadBalancerForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkLoadBalancerForwardingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLoadBalancerForwardingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLoadBalancerForwardingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkLoadBalancerForwardingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NetworkLoadBalancerForwardingRule) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkLoadBalancerForwardingRule) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkLoadBalancerForwardingRule) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkLoadBalancerForwardingRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NetworkLoadBalancerForwardingRule) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NetworkLoadBalancerForwardingRule) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NetworkLoadBalancerForwardingRule) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NetworkLoadBalancerForwardingRule) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkLoadBalancerForwardingRule) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkLoadBalancerForwardingRule) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkLoadBalancerForwardingRule) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkLoadBalancerForwardingRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *NetworkLoadBalancerForwardingRule) GetProperties() NetworkLoadBalancerForwardingRuleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NetworkLoadBalancerForwardingRule) GetPropertiesOk() (*NetworkLoadBalancerForwardingRuleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NetworkLoadBalancerForwardingRule) SetProperties(v NetworkLoadBalancerForwardingRuleProperties)`

SetProperties sets Properties field to given value.




