/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// S3Bucket struct for S3Bucket
type S3Bucket struct {
	// Name of the S3 bucket
	Name *string `json:"name"`
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *S3Bucket) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Bucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *S3Bucket) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *S3Bucket) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

func (o S3Bucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableS3Bucket struct {
	value *S3Bucket
	isSet bool
}

func (v NullableS3Bucket) Get() *S3Bucket {
	return v.value
}

func (v *NullableS3Bucket) Set(val *S3Bucket) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Bucket) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Bucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Bucket(val *S3Bucket) *NullableS3Bucket {
	return &NullableS3Bucket{value: val, isSet: true}
}

func (v NullableS3Bucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Bucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
