/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
	"fmt"
)

// NatGatewayRuleType the model 'NatGatewayRuleType'
type NatGatewayRuleType string

// List of NatGatewayRuleType
const (
	SNAT NatGatewayRuleType = "SNAT"
)

func (v *NatGatewayRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NatGatewayRuleType(value)
	for _, existing := range []NatGatewayRuleType{ "SNAT",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NatGatewayRuleType", value)
}

// Ptr returns reference to NatGatewayRuleType value
func (v NatGatewayRuleType) Ptr() *NatGatewayRuleType {
	return &v
}

type NullableNatGatewayRuleType struct {
	value *NatGatewayRuleType
	isSet bool
}

func (v NullableNatGatewayRuleType) Get() *NatGatewayRuleType {
	return v.value
}

func (v *NullableNatGatewayRuleType) Set(val *NatGatewayRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGatewayRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGatewayRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGatewayRuleType(val *NatGatewayRuleType) *NullableNatGatewayRuleType {
	return &NullableNatGatewayRuleType{value: val, isSet: true}
}

func (v NullableNatGatewayRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGatewayRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
