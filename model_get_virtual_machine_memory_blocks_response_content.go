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

// GetVirtualMachineMemoryBlocksResponseContent struct for GetVirtualMachineMemoryBlocksResponseContent
type GetVirtualMachineMemoryBlocksResponseContent struct {
	Data GetMemoryBlocksResult `json:"data"`
}

// NewGetVirtualMachineMemoryBlocksResponseContent instantiates a new GetVirtualMachineMemoryBlocksResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualMachineMemoryBlocksResponseContent(data GetMemoryBlocksResult) *GetVirtualMachineMemoryBlocksResponseContent {
	this := GetVirtualMachineMemoryBlocksResponseContent{}
	this.Data = data
	return &this
}

// NewGetVirtualMachineMemoryBlocksResponseContentWithDefaults instantiates a new GetVirtualMachineMemoryBlocksResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualMachineMemoryBlocksResponseContentWithDefaults() *GetVirtualMachineMemoryBlocksResponseContent {
	this := GetVirtualMachineMemoryBlocksResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetVirtualMachineMemoryBlocksResponseContent) GetData() GetMemoryBlocksResult {
	if o == nil {
		var ret GetMemoryBlocksResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetVirtualMachineMemoryBlocksResponseContent) GetDataOk() (*GetMemoryBlocksResult, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetVirtualMachineMemoryBlocksResponseContent) SetData(v GetMemoryBlocksResult) {
	o.Data = v
}

func (o GetVirtualMachineMemoryBlocksResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetVirtualMachineMemoryBlocksResponseContent struct {
	value *GetVirtualMachineMemoryBlocksResponseContent
	isSet bool
}

func (v NullableGetVirtualMachineMemoryBlocksResponseContent) Get() *GetVirtualMachineMemoryBlocksResponseContent {
	return v.value
}

func (v *NullableGetVirtualMachineMemoryBlocksResponseContent) Set(val *GetVirtualMachineMemoryBlocksResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualMachineMemoryBlocksResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualMachineMemoryBlocksResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualMachineMemoryBlocksResponseContent(val *GetVirtualMachineMemoryBlocksResponseContent) *NullableGetVirtualMachineMemoryBlocksResponseContent {
	return &NullableGetVirtualMachineMemoryBlocksResponseContent{value: val, isSet: true}
}

func (v NullableGetVirtualMachineMemoryBlocksResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualMachineMemoryBlocksResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


