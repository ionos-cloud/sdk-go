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

// NicProperties struct for NicProperties
type NicProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// The MAC address of the NIC
	Mac *string `json:"mac,omitempty"`
	// Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips *[]string `json:"ips,omitempty"`
	// Indicates if the nic will reserve an IP using DHCP
	Dhcp *bool `json:"dhcp,omitempty"`
	// The LAN ID the NIC will sit on. If the LAN ID does not exist it will be implicitly created
	Lan *int32 `json:"lan"`
	// Activate or deactivate the firewall. By default an active firewall without any defined rules will block all incoming network traffic except for the firewall rules that explicitly allows certain protocols, ip addresses and ports.
	FirewallActive *bool `json:"firewallActive,omitempty"`
	// Indicates if NAT is enabled on this NIC. This is now deprecated.
	Nat *bool `json:"nat,omitempty"`
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NicProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *NicProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *NicProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetMac returns the Mac field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NicProperties) GetMac() *string {
	if o == nil {
		return nil
	}

	return o.Mac

}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Mac, true
}

// SetMac sets field value
func (o *NicProperties) SetMac(v string) {

	o.Mac = &v

}

// HasMac returns a boolean if a field has been set.
func (o *NicProperties) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// GetIps returns the Ips field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *NicProperties) GetIps() *[]string {
	if o == nil {
		return nil
	}

	return o.Ips

}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ips, true
}

// SetIps sets field value
func (o *NicProperties) SetIps(v []string) {

	o.Ips = &v

}

// HasIps returns a boolean if a field has been set.
func (o *NicProperties) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// GetDhcp returns the Dhcp field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *NicProperties) GetDhcp() *bool {
	if o == nil {
		return nil
	}

	return o.Dhcp

}

// GetDhcpOk returns a tuple with the Dhcp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetDhcpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Dhcp, true
}

// SetDhcp sets field value
func (o *NicProperties) SetDhcp(v bool) {

	o.Dhcp = &v

}

// HasDhcp returns a boolean if a field has been set.
func (o *NicProperties) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}

// GetLan returns the Lan field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NicProperties) GetLan() *int32 {
	if o == nil {
		return nil
	}

	return o.Lan

}

// GetLanOk returns a tuple with the Lan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetLanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Lan, true
}

// SetLan sets field value
func (o *NicProperties) SetLan(v int32) {

	o.Lan = &v

}

// HasLan returns a boolean if a field has been set.
func (o *NicProperties) HasLan() bool {
	if o != nil && o.Lan != nil {
		return true
	}

	return false
}

// GetFirewallActive returns the FirewallActive field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *NicProperties) GetFirewallActive() *bool {
	if o == nil {
		return nil
	}

	return o.FirewallActive

}

// GetFirewallActiveOk returns a tuple with the FirewallActive field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetFirewallActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.FirewallActive, true
}

// SetFirewallActive sets field value
func (o *NicProperties) SetFirewallActive(v bool) {

	o.FirewallActive = &v

}

// HasFirewallActive returns a boolean if a field has been set.
func (o *NicProperties) HasFirewallActive() bool {
	if o != nil && o.FirewallActive != nil {
		return true
	}

	return false
}

// GetNat returns the Nat field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *NicProperties) GetNat() *bool {
	if o == nil {
		return nil
	}

	return o.Nat

}

// GetNatOk returns a tuple with the Nat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicProperties) GetNatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nat, true
}

// SetNat sets field value
func (o *NicProperties) SetNat(v bool) {

	o.Nat = &v

}

// HasNat returns a boolean if a field has been set.
func (o *NicProperties) HasNat() bool {
	if o != nil && o.Nat != nil {
		return true
	}

	return false
}

func (o NicProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.Dhcp != nil {
		toSerialize["dhcp"] = o.Dhcp
	}
	if o.Lan != nil {
		toSerialize["lan"] = o.Lan
	}
	if o.FirewallActive != nil {
		toSerialize["firewallActive"] = o.FirewallActive
	}
	if o.Nat != nil {
		toSerialize["nat"] = o.Nat
	}
	return json.Marshal(toSerialize)
}

type NullableNicProperties struct {
	value *NicProperties
	isSet bool
}

func (v NullableNicProperties) Get() *NicProperties {
	return v.value
}

func (v *NullableNicProperties) Set(val *NicProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNicProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNicProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicProperties(val *NicProperties) *NullableNicProperties {
	return &NullableNicProperties{value: val, isSet: true}
}

func (v NullableNicProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
