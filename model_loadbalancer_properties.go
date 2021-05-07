/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// LoadbalancerProperties struct for LoadbalancerProperties
type LoadbalancerProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// IPv4 address of the loadbalancer. All attached NICs will inherit this IP. Leaving value null will assign IP automatically
	Ip *string `json:"ip,omitempty"`
	// Indicates if the loadbalancer will reserve an IP using DHCP
	Dhcp *bool `json:"dhcp,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LoadbalancerProperties) GetName() *string {
	if o == nil {
		return nil
	}


	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadbalancerProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Name, true
}

// SetName sets field value
func (o *LoadbalancerProperties) SetName(v string) {


	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *LoadbalancerProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LoadbalancerProperties) GetIp() *string {
	if o == nil {
		return nil
	}


	return o.Ip

}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadbalancerProperties) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Ip, true
}

// SetIp sets field value
func (o *LoadbalancerProperties) SetIp(v string) {


	o.Ip = &v

}

// HasIp returns a boolean if a field has been set.
func (o *LoadbalancerProperties) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}



// GetDhcp returns the Dhcp field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *LoadbalancerProperties) GetDhcp() *bool {
	if o == nil {
		return nil
	}


	return o.Dhcp

}

// GetDhcpOk returns a tuple with the Dhcp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadbalancerProperties) GetDhcpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.Dhcp, true
}

// SetDhcp sets field value
func (o *LoadbalancerProperties) SetDhcp(v bool) {


	o.Dhcp = &v

}

// HasDhcp returns a boolean if a field has been set.
func (o *LoadbalancerProperties) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}


func (o LoadbalancerProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	

	if o.Dhcp != nil {
		toSerialize["dhcp"] = o.Dhcp
	}
	
	return json.Marshal(toSerialize)
}

type NullableLoadbalancerProperties struct {
	value *LoadbalancerProperties
	isSet bool
}

func (v NullableLoadbalancerProperties) Get() *LoadbalancerProperties {
	return v.value
}

func (v *NullableLoadbalancerProperties) Set(val *LoadbalancerProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerProperties(val *LoadbalancerProperties) *NullableLoadbalancerProperties {
	return &NullableLoadbalancerProperties{value: val, isSet: true}
}

func (v NullableLoadbalancerProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


