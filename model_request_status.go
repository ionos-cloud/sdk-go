/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RequestStatus struct for RequestStatus
type RequestStatus struct {
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	// The type of object that has been created
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	Metadata *RequestStatusMetadata `json:"metadata,omitempty"`
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequestStatus) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *RequestStatus) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *RequestStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *RequestStatus) GetType() *Type {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestStatus) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *RequestStatus) SetType(v Type) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *RequestStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequestStatus) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestStatus) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *RequestStatus) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *RequestStatus) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for RequestStatusMetadata will be returned
func (o *RequestStatus) GetMetadata() *RequestStatusMetadata {
	if o == nil {
		return nil
	}


	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestStatus) GetMetadataOk() (*RequestStatusMetadata, bool) {
	if o == nil {
		return nil, false
	}


	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RequestStatus) SetMetadata(v RequestStatusMetadata) {


	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *RequestStatus) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}


func (o RequestStatus) MarshalJSON() ([]byte, error) {
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
	
	return json.Marshal(toSerialize)
}

type NullableRequestStatus struct {
	value *RequestStatus
	isSet bool
}

func (v NullableRequestStatus) Get() *RequestStatus {
	return v.value
}

func (v *NullableRequestStatus) Set(val *RequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatus(val *RequestStatus) *NullableRequestStatus {
	return &NullableRequestStatus{value: val, isSet: true}
}

func (v NullableRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


