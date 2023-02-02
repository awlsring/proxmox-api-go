/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package proxmox

import (
	"encoding/json"
)

// checks if the OrderNodeCertificateRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderNodeCertificateRequestContent{}

// OrderNodeCertificateRequestContent struct for OrderNodeCertificateRequestContent
type OrderNodeCertificateRequestContent struct {
	// Overwrite existing custom certificate.
	Force *float32 `json:"force,omitempty"`
}

// NewOrderNodeCertificateRequestContent instantiates a new OrderNodeCertificateRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderNodeCertificateRequestContent() *OrderNodeCertificateRequestContent {
	this := OrderNodeCertificateRequestContent{}
	return &this
}

// NewOrderNodeCertificateRequestContentWithDefaults instantiates a new OrderNodeCertificateRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderNodeCertificateRequestContentWithDefaults() *OrderNodeCertificateRequestContent {
	this := OrderNodeCertificateRequestContent{}
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *OrderNodeCertificateRequestContent) GetForce() float32 {
	if o == nil || isNil(o.Force) {
		var ret float32
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderNodeCertificateRequestContent) GetForceOk() (*float32, bool) {
	if o == nil || isNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *OrderNodeCertificateRequestContent) HasForce() bool {
	if o != nil && !isNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given float32 and assigns it to the Force field.
func (o *OrderNodeCertificateRequestContent) SetForce(v float32) {
	o.Force = &v
}

func (o OrderNodeCertificateRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderNodeCertificateRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return toSerialize, nil
}

type NullableOrderNodeCertificateRequestContent struct {
	value *OrderNodeCertificateRequestContent
	isSet bool
}

func (v NullableOrderNodeCertificateRequestContent) Get() *OrderNodeCertificateRequestContent {
	return v.value
}

func (v *NullableOrderNodeCertificateRequestContent) Set(val *OrderNodeCertificateRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderNodeCertificateRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderNodeCertificateRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderNodeCertificateRequestContent(val *OrderNodeCertificateRequestContent) *NullableOrderNodeCertificateRequestContent {
	return &NullableOrderNodeCertificateRequestContent{value: val, isSet: true}
}

func (v NullableOrderNodeCertificateRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderNodeCertificateRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


