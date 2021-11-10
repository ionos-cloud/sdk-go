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

// LabelProperties struct for LabelProperties
type LabelProperties struct {
	// A label key
	Key *string `json:"key,omitempty"`
	// A label value
	Value *string `json:"value,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"resourceId,omitempty"`
	// The type of the resource on which the label is applied.
	ResourceType *string `json:"resourceType,omitempty"`
	// URL to the Resource (absolute path) on which the label is applied.
	ResourceHref *string `json:"resourceHref,omitempty"`
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelProperties) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelProperties) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *LabelProperties) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *LabelProperties) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelProperties) GetValue() *string {
	if o == nil {
		return nil
	}

	return o.Value

}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelProperties) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Value, true
}

// SetValue sets field value
func (o *LabelProperties) SetValue(v string) {

	o.Value = &v

}

// HasValue returns a boolean if a field has been set.
func (o *LabelProperties) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// GetResourceId returns the ResourceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelProperties) GetResourceId() *string {
	if o == nil {
		return nil
	}

	return o.ResourceId

}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelProperties) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *LabelProperties) SetResourceId(v string) {

	o.ResourceId = &v

}

// HasResourceId returns a boolean if a field has been set.
func (o *LabelProperties) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// GetResourceType returns the ResourceType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelProperties) GetResourceType() *string {
	if o == nil {
		return nil
	}

	return o.ResourceType

}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelProperties) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *LabelProperties) SetResourceType(v string) {

	o.ResourceType = &v

}

// HasResourceType returns a boolean if a field has been set.
func (o *LabelProperties) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// GetResourceHref returns the ResourceHref field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabelProperties) GetResourceHref() *string {
	if o == nil {
		return nil
	}

	return o.ResourceHref

}

// GetResourceHrefOk returns a tuple with the ResourceHref field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabelProperties) GetResourceHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ResourceHref, true
}

// SetResourceHref sets field value
func (o *LabelProperties) SetResourceHref(v string) {

	o.ResourceHref = &v

}

// HasResourceHref returns a boolean if a field has been set.
func (o *LabelProperties) HasResourceHref() bool {
	if o != nil && o.ResourceHref != nil {
		return true
	}

	return false
}

func (o LabelProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}

	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}

	if o.ResourceHref != nil {
		toSerialize["resourceHref"] = o.ResourceHref
	}
	return json.Marshal(toSerialize)
}

type NullableLabelProperties struct {
	value *LabelProperties
	isSet bool
}

func (v NullableLabelProperties) Get() *LabelProperties {
	return v.value
}

func (v *NullableLabelProperties) Set(val *LabelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelProperties(val *LabelProperties) *NullableLabelProperties {
	return &NullableLabelProperties{value: val, isSet: true}
}

func (v NullableLabelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
