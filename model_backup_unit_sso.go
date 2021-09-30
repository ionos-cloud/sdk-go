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

// BackupUnitSSO struct for BackupUnitSSO
type BackupUnitSSO struct {
	// The backup unit single sign on url
	SsoUrl *string `json:"ssoUrl,omitempty"`
}



// GetSsoUrl returns the SsoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupUnitSSO) GetSsoUrl() *string {
	if o == nil {
		return nil
	}


	return o.SsoUrl

}

// GetSsoUrlOk returns a tuple with the SsoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnitSSO) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.SsoUrl, true
}

// SetSsoUrl sets field value
func (o *BackupUnitSSO) SetSsoUrl(v string) {


	o.SsoUrl = &v

}

// HasSsoUrl returns a boolean if a field has been set.
func (o *BackupUnitSSO) HasSsoUrl() bool {
	if o != nil && o.SsoUrl != nil {
		return true
	}

	return false
}

func (o BackupUnitSSO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SsoUrl != nil {
		toSerialize["ssoUrl"] = o.SsoUrl
	}
	return json.Marshal(toSerialize)
}

type NullableBackupUnitSSO struct {
	value *BackupUnitSSO
	isSet bool
}

func (v NullableBackupUnitSSO) Get() *BackupUnitSSO {
	return v.value
}

func (v *NullableBackupUnitSSO) Set(val *BackupUnitSSO) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupUnitSSO) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupUnitSSO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupUnitSSO(val *BackupUnitSSO) *NullableBackupUnitSSO {
	return &NullableBackupUnitSSO{value: val, isSet: true}
}

func (v NullableBackupUnitSSO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupUnitSSO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


