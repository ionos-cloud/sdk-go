# ConnectableDatacenter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | Identifier of the virtual data center that can be connected to the Cross Connect. | [optional] |
|**Name** | Pointer to **string** | Name of the virtual data center that can be connected to the Cross Connect. | [optional] |
|**Location** | Pointer to **string** | Location of the virtual data center that can be connected to the Cross Connect. | [optional] |

## Methods

### NewConnectableDatacenter

`func NewConnectableDatacenter() *ConnectableDatacenter`

NewConnectableDatacenter instantiates a new ConnectableDatacenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectableDatacenterWithDefaults

`func NewConnectableDatacenterWithDefaults() *ConnectableDatacenter`

NewConnectableDatacenterWithDefaults instantiates a new ConnectableDatacenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectableDatacenter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectableDatacenter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectableDatacenter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectableDatacenter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConnectableDatacenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectableDatacenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectableDatacenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectableDatacenter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *ConnectableDatacenter) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConnectableDatacenter) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConnectableDatacenter) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ConnectableDatacenter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.



