# ApplicationLoadBalancerForwardingRule

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**ApplicationLoadBalancerForwardingRuleProperties**](ApplicationLoadBalancerForwardingRuleProperties.md) |  | |

## Methods

### NewApplicationLoadBalancerForwardingRule

`func NewApplicationLoadBalancerForwardingRule(properties ApplicationLoadBalancerForwardingRuleProperties, ) *ApplicationLoadBalancerForwardingRule`

NewApplicationLoadBalancerForwardingRule instantiates a new ApplicationLoadBalancerForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerForwardingRuleWithDefaults

`func NewApplicationLoadBalancerForwardingRuleWithDefaults() *ApplicationLoadBalancerForwardingRule`

NewApplicationLoadBalancerForwardingRuleWithDefaults instantiates a new ApplicationLoadBalancerForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationLoadBalancerForwardingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationLoadBalancerForwardingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationLoadBalancerForwardingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationLoadBalancerForwardingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationLoadBalancerForwardingRule) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationLoadBalancerForwardingRule) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationLoadBalancerForwardingRule) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationLoadBalancerForwardingRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ApplicationLoadBalancerForwardingRule) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplicationLoadBalancerForwardingRule) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplicationLoadBalancerForwardingRule) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplicationLoadBalancerForwardingRule) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationLoadBalancerForwardingRule) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationLoadBalancerForwardingRule) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationLoadBalancerForwardingRule) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationLoadBalancerForwardingRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ApplicationLoadBalancerForwardingRule) GetProperties() ApplicationLoadBalancerForwardingRuleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ApplicationLoadBalancerForwardingRule) GetPropertiesOk() (*ApplicationLoadBalancerForwardingRuleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ApplicationLoadBalancerForwardingRule) SetProperties(v ApplicationLoadBalancerForwardingRuleProperties)`

SetProperties sets Properties field to given value.




