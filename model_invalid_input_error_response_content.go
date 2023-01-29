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

// InvalidInputErrorResponseContent struct for InvalidInputErrorResponseContent
type InvalidInputErrorResponseContent struct {
	Message string `json:"message"`
}

// NewInvalidInputErrorResponseContent instantiates a new InvalidInputErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidInputErrorResponseContent(message string) *InvalidInputErrorResponseContent {
	this := InvalidInputErrorResponseContent{}
	this.Message = message
	return &this
}

// NewInvalidInputErrorResponseContentWithDefaults instantiates a new InvalidInputErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidInputErrorResponseContentWithDefaults() *InvalidInputErrorResponseContent {
	this := InvalidInputErrorResponseContent{}
	return &this
}

// GetMessage returns the Message field value
func (o *InvalidInputErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidInputErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidInputErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o InvalidInputErrorResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidInputErrorResponseContent struct {
	value *InvalidInputErrorResponseContent
	isSet bool
}

func (v NullableInvalidInputErrorResponseContent) Get() *InvalidInputErrorResponseContent {
	return v.value
}

func (v *NullableInvalidInputErrorResponseContent) Set(val *InvalidInputErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidInputErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidInputErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidInputErrorResponseContent(val *InvalidInputErrorResponseContent) *NullableInvalidInputErrorResponseContent {
	return &NullableInvalidInputErrorResponseContent{value: val, isSet: true}
}

func (v NullableInvalidInputErrorResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidInputErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


