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

// JoinClusterResponseContent struct for JoinClusterResponseContent
type JoinClusterResponseContent struct {
	Data string `json:"data"`
}

// NewJoinClusterResponseContent instantiates a new JoinClusterResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinClusterResponseContent(data string) *JoinClusterResponseContent {
	this := JoinClusterResponseContent{}
	this.Data = data
	return &this
}

// NewJoinClusterResponseContentWithDefaults instantiates a new JoinClusterResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinClusterResponseContentWithDefaults() *JoinClusterResponseContent {
	this := JoinClusterResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *JoinClusterResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *JoinClusterResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *JoinClusterResponseContent) SetData(v string) {
	o.Data = v
}

func (o JoinClusterResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableJoinClusterResponseContent struct {
	value *JoinClusterResponseContent
	isSet bool
}

func (v NullableJoinClusterResponseContent) Get() *JoinClusterResponseContent {
	return v.value
}

func (v *NullableJoinClusterResponseContent) Set(val *JoinClusterResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinClusterResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinClusterResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinClusterResponseContent(val *JoinClusterResponseContent) *NullableJoinClusterResponseContent {
	return &NullableJoinClusterResponseContent{value: val, isSet: true}
}

func (v NullableJoinClusterResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinClusterResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


