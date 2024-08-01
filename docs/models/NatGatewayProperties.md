# NatGatewayProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | Name of the NAT Gateway. | |
|**PublicIps** | **[]string** | Collection of public IP addresses of the NAT Gateway. Should be customer reserved IP addresses in that location. | |
|**Lans** | Pointer to [**[]NatGatewayLanProperties**](NatGatewayLanProperties.md) | Collection of LANs connected to the NAT Gateway. IPs must contain a valid subnet mask. If no IP is provided, the system will generate an IP with /24 subnet. | [optional] |

## Methods

### NewNatGatewayProperties

`func NewNatGatewayProperties(name string, publicIps []string, ) *NatGatewayProperties`

NewNatGatewayProperties instantiates a new NatGatewayProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayPropertiesWithDefaults

`func NewNatGatewayPropertiesWithDefaults() *NatGatewayProperties`

NewNatGatewayPropertiesWithDefaults instantiates a new NatGatewayProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NatGatewayProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatGatewayProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatGatewayProperties) SetName(v string)`

SetName sets Name field to given value.


### GetPublicIps

`func (o *NatGatewayProperties) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *NatGatewayProperties) GetPublicIpsOk() (*[]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *NatGatewayProperties) SetPublicIps(v []string)`

SetPublicIps sets PublicIps field to given value.


### GetLans

`func (o *NatGatewayProperties) GetLans() []NatGatewayLanProperties`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *NatGatewayProperties) GetLansOk() (*[]NatGatewayLanProperties, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *NatGatewayProperties) SetLans(v []NatGatewayLanProperties)`

SetLans sets Lans field to given value.

### HasLans

`func (o *NatGatewayProperties) HasLans() bool`

HasLans returns a boolean if a field has been set.



