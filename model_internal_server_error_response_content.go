/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InternalServerErrorResponseContent struct for InternalServerErrorResponseContent
type InternalServerErrorResponseContent struct {
	Message string `json:"message"`
}

// NewInternalServerErrorResponseContent instantiates a new InternalServerErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerErrorResponseContent(message string) *InternalServerErrorResponseContent {
	this := InternalServerErrorResponseContent{}
	this.Message = message
	return &this
}

// NewInternalServerErrorResponseContentWithDefaults instantiates a new InternalServerErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerErrorResponseContentWithDefaults() *InternalServerErrorResponseContent {
	this := InternalServerErrorResponseContent{}
	return &this
}

// GetMessage returns the Message field value
func (o *InternalServerErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InternalServerErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InternalServerErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o InternalServerErrorResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInternalServerErrorResponseContent struct {
	value *InternalServerErrorResponseContent
	isSet bool
}

func (v NullableInternalServerErrorResponseContent) Get() *InternalServerErrorResponseContent {
	return v.value
}

func (v *NullableInternalServerErrorResponseContent) Set(val *InternalServerErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerErrorResponseContent(val *InternalServerErrorResponseContent) *NullableInternalServerErrorResponseContent {
	return &NullableInternalServerErrorResponseContent{value: val, isSet: true}
}

func (v NullableInternalServerErrorResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalServerErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

