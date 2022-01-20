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

// Datacenter struct for Datacenter
type Datacenter struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                    `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata `json:"metadata,omitempty"`
	Properties *DatacenterProperties      `json:"properties"`
	Entities   *DataCenterEntities        `json:"entities,omitempty"`
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Datacenter) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Datacenter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Datacenter) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Datacenter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *Datacenter) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Datacenter) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *Datacenter) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Datacenter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Datacenter) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Datacenter) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Datacenter) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Datacenter) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for DatacenterElementMetadata will be returned
func (o *Datacenter) GetMetadata() *DatacenterElementMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Datacenter) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Datacenter) SetMetadata(v DatacenterElementMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *Datacenter) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for DatacenterProperties will be returned
func (o *Datacenter) GetProperties() *DatacenterProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Datacenter) GetPropertiesOk() (*DatacenterProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *Datacenter) SetProperties(v DatacenterProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *Datacenter) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for DataCenterEntities will be returned
func (o *Datacenter) GetEntities() *DataCenterEntities {
	if o == nil {
		return nil
	}

	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Datacenter) GetEntitiesOk() (*DataCenterEntities, bool) {
	if o == nil {
		return nil, false
	}

	return o.Entities, true
}

// SetEntities sets field value
func (o *Datacenter) SetEntities(v DataCenterEntities) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *Datacenter) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

func (o Datacenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	return json.Marshal(toSerialize)
}

type NullableDatacenter struct {
	value *Datacenter
	isSet bool
}

func (v NullableDatacenter) Get() *Datacenter {
	return v.value
}

func (v *NullableDatacenter) Set(val *Datacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenter(val *Datacenter) *NullableDatacenter {
	return &NullableDatacenter{value: val, isSet: true}
}

func (v NullableDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
