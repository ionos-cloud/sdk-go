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

// NicEntities struct for NicEntities
type NicEntities struct {
	Flowlogs       *FlowLogs       `json:"flowlogs,omitempty"`
	Firewallrules  *FirewallRules  `json:"firewallrules,omitempty"`
	Securitygroups *SecurityGroups `json:"securitygroups,omitempty"`
}

// NewNicEntities instantiates a new NicEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNicEntities() *NicEntities {
	this := NicEntities{}

	return &this
}

// NewNicEntitiesWithDefaults instantiates a new NicEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNicEntitiesWithDefaults() *NicEntities {
	this := NicEntities{}
	return &this
}

// GetFlowlogs returns the Flowlogs field value
// If the value is explicit nil, nil is returned
func (o *NicEntities) GetFlowlogs() *FlowLogs {
	if o == nil {
		return nil
	}

	return o.Flowlogs

}

// GetFlowlogsOk returns a tuple with the Flowlogs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicEntities) GetFlowlogsOk() (*FlowLogs, bool) {
	if o == nil {
		return nil, false
	}

	return o.Flowlogs, true
}

// SetFlowlogs sets field value
func (o *NicEntities) SetFlowlogs(v FlowLogs) {

	o.Flowlogs = &v

}

// HasFlowlogs returns a boolean if a field has been set.
func (o *NicEntities) HasFlowlogs() bool {
	if o != nil && o.Flowlogs != nil {
		return true
	}

	return false
}

// GetFirewallrules returns the Firewallrules field value
// If the value is explicit nil, nil is returned
func (o *NicEntities) GetFirewallrules() *FirewallRules {
	if o == nil {
		return nil
	}

	return o.Firewallrules

}

// GetFirewallrulesOk returns a tuple with the Firewallrules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicEntities) GetFirewallrulesOk() (*FirewallRules, bool) {
	if o == nil {
		return nil, false
	}

	return o.Firewallrules, true
}

// SetFirewallrules sets field value
func (o *NicEntities) SetFirewallrules(v FirewallRules) {

	o.Firewallrules = &v

}

// HasFirewallrules returns a boolean if a field has been set.
func (o *NicEntities) HasFirewallrules() bool {
	if o != nil && o.Firewallrules != nil {
		return true
	}

	return false
}

// GetSecuritygroups returns the Securitygroups field value
// If the value is explicit nil, nil is returned
func (o *NicEntities) GetSecuritygroups() *SecurityGroups {
	if o == nil {
		return nil
	}

	return o.Securitygroups

}

// GetSecuritygroupsOk returns a tuple with the Securitygroups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicEntities) GetSecuritygroupsOk() (*SecurityGroups, bool) {
	if o == nil {
		return nil, false
	}

	return o.Securitygroups, true
}

// SetSecuritygroups sets field value
func (o *NicEntities) SetSecuritygroups(v SecurityGroups) {

	o.Securitygroups = &v

}

// HasSecuritygroups returns a boolean if a field has been set.
func (o *NicEntities) HasSecuritygroups() bool {
	if o != nil && o.Securitygroups != nil {
		return true
	}

	return false
}

func (o NicEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flowlogs != nil {
		toSerialize["flowlogs"] = o.Flowlogs
	}

	if o.Firewallrules != nil {
		toSerialize["firewallrules"] = o.Firewallrules
	}

	if o.Securitygroups != nil {
		toSerialize["securitygroups"] = o.Securitygroups
	}

	return json.Marshal(toSerialize)
}

type NullableNicEntities struct {
	value *NicEntities
	isSet bool
}

func (v NullableNicEntities) Get() *NicEntities {
	return v.value
}

func (v *NullableNicEntities) Set(val *NicEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableNicEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableNicEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicEntities(val *NicEntities) *NullableNicEntities {
	return &NullableNicEntities{value: val, isSet: true}
}

func (v NullableNicEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
