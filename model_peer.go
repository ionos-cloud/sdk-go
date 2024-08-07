/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Peer struct for Peer
type Peer struct {
	// Identifier of the LAN connected to the Cross Connect.
	Id *string `json:"id,omitempty"`
	// Name of the LAN connected to the Cross Connect.
	Name *string `json:"name,omitempty"`
	// Identifier of the virtual data center connected to the Cross Connect.
	DatacenterId *string `json:"datacenterId,omitempty"`
	// Name of the virtual data center connected to the Cross Connect.
	DatacenterName *string `json:"datacenterName,omitempty"`
	// Location of the virtual data center connected to the Cross Connect.
	Location *string `json:"location,omitempty"`
}

// NewPeer instantiates a new Peer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeer() *Peer {
	this := Peer{}

	return &this
}

// NewPeerWithDefaults instantiates a new Peer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeerWithDefaults() *Peer {
	this := Peer{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *Peer) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Peer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Peer) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Peer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *Peer) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Peer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *Peer) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *Peer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetDatacenterId returns the DatacenterId field value
// If the value is explicit nil, nil is returned
func (o *Peer) GetDatacenterId() *string {
	if o == nil {
		return nil
	}

	return o.DatacenterId

}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Peer) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DatacenterId, true
}

// SetDatacenterId sets field value
func (o *Peer) SetDatacenterId(v string) {

	o.DatacenterId = &v

}

// HasDatacenterId returns a boolean if a field has been set.
func (o *Peer) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// GetDatacenterName returns the DatacenterName field value
// If the value is explicit nil, nil is returned
func (o *Peer) GetDatacenterName() *string {
	if o == nil {
		return nil
	}

	return o.DatacenterName

}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Peer) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DatacenterName, true
}

// SetDatacenterName sets field value
func (o *Peer) SetDatacenterName(v string) {

	o.DatacenterName = &v

}

// HasDatacenterName returns a boolean if a field has been set.
func (o *Peer) HasDatacenterName() bool {
	if o != nil && o.DatacenterName != nil {
		return true
	}

	return false
}

// GetLocation returns the Location field value
// If the value is explicit nil, nil is returned
func (o *Peer) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location

}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Peer) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Location, true
}

// SetLocation sets field value
func (o *Peer) SetLocation(v string) {

	o.Location = &v

}

// HasLocation returns a boolean if a field has been set.
func (o *Peer) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

func (o Peer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.DatacenterId != nil {
		toSerialize["datacenterId"] = o.DatacenterId
	}

	if o.DatacenterName != nil {
		toSerialize["datacenterName"] = o.DatacenterName
	}

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	return json.Marshal(toSerialize)
}

type NullablePeer struct {
	value *Peer
	isSet bool
}

func (v NullablePeer) Get() *Peer {
	return v.value
}

func (v *NullablePeer) Set(val *Peer) {
	v.value = val
	v.isSet = true
}

func (v NullablePeer) IsSet() bool {
	return v.isSet
}

func (v *NullablePeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeer(val *Peer) *NullablePeer {
	return &NullablePeer{value: val, isSet: true}
}

func (v NullablePeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
