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

// GetVirtualMachineOperatingSystemInformationResponseContent struct for GetVirtualMachineOperatingSystemInformationResponseContent
type GetVirtualMachineOperatingSystemInformationResponseContent struct {
	Data GetOsInfoResult `json:"data"`
}

// NewGetVirtualMachineOperatingSystemInformationResponseContent instantiates a new GetVirtualMachineOperatingSystemInformationResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualMachineOperatingSystemInformationResponseContent(data GetOsInfoResult) *GetVirtualMachineOperatingSystemInformationResponseContent {
	this := GetVirtualMachineOperatingSystemInformationResponseContent{}
	this.Data = data
	return &this
}

// NewGetVirtualMachineOperatingSystemInformationResponseContentWithDefaults instantiates a new GetVirtualMachineOperatingSystemInformationResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualMachineOperatingSystemInformationResponseContentWithDefaults() *GetVirtualMachineOperatingSystemInformationResponseContent {
	this := GetVirtualMachineOperatingSystemInformationResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetVirtualMachineOperatingSystemInformationResponseContent) GetData() GetOsInfoResult {
	if o == nil {
		var ret GetOsInfoResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetVirtualMachineOperatingSystemInformationResponseContent) GetDataOk() (*GetOsInfoResult, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetVirtualMachineOperatingSystemInformationResponseContent) SetData(v GetOsInfoResult) {
	o.Data = v
}

func (o GetVirtualMachineOperatingSystemInformationResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetVirtualMachineOperatingSystemInformationResponseContent struct {
	value *GetVirtualMachineOperatingSystemInformationResponseContent
	isSet bool
}

func (v NullableGetVirtualMachineOperatingSystemInformationResponseContent) Get() *GetVirtualMachineOperatingSystemInformationResponseContent {
	return v.value
}

func (v *NullableGetVirtualMachineOperatingSystemInformationResponseContent) Set(val *GetVirtualMachineOperatingSystemInformationResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualMachineOperatingSystemInformationResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualMachineOperatingSystemInformationResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualMachineOperatingSystemInformationResponseContent(val *GetVirtualMachineOperatingSystemInformationResponseContent) *NullableGetVirtualMachineOperatingSystemInformationResponseContent {
	return &NullableGetVirtualMachineOperatingSystemInformationResponseContent{value: val, isSet: true}
}

func (v NullableGetVirtualMachineOperatingSystemInformationResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualMachineOperatingSystemInformationResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


