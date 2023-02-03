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

// checks if the GetPendingVirtualMachineCloudInitChangesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPendingVirtualMachineCloudInitChangesResponseContent{}

// GetPendingVirtualMachineCloudInitChangesResponseContent struct for GetPendingVirtualMachineCloudInitChangesResponseContent
type GetPendingVirtualMachineCloudInitChangesResponseContent struct {
	Data []PendingVirtualMachineCloudInitField `json:"data"`
}

// NewGetPendingVirtualMachineCloudInitChangesResponseContent instantiates a new GetPendingVirtualMachineCloudInitChangesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPendingVirtualMachineCloudInitChangesResponseContent(data []PendingVirtualMachineCloudInitField) *GetPendingVirtualMachineCloudInitChangesResponseContent {
	this := GetPendingVirtualMachineCloudInitChangesResponseContent{}
	this.Data = data
	return &this
}

// NewGetPendingVirtualMachineCloudInitChangesResponseContentWithDefaults instantiates a new GetPendingVirtualMachineCloudInitChangesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPendingVirtualMachineCloudInitChangesResponseContentWithDefaults() *GetPendingVirtualMachineCloudInitChangesResponseContent {
	this := GetPendingVirtualMachineCloudInitChangesResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetPendingVirtualMachineCloudInitChangesResponseContent) GetData() []PendingVirtualMachineCloudInitField {
	if o == nil {
		var ret []PendingVirtualMachineCloudInitField
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetPendingVirtualMachineCloudInitChangesResponseContent) GetDataOk() ([]PendingVirtualMachineCloudInitField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetPendingVirtualMachineCloudInitChangesResponseContent) SetData(v []PendingVirtualMachineCloudInitField) {
	o.Data = v
}

func (o GetPendingVirtualMachineCloudInitChangesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPendingVirtualMachineCloudInitChangesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetPendingVirtualMachineCloudInitChangesResponseContent struct {
	value *GetPendingVirtualMachineCloudInitChangesResponseContent
	isSet bool
}

func (v NullableGetPendingVirtualMachineCloudInitChangesResponseContent) Get() *GetPendingVirtualMachineCloudInitChangesResponseContent {
	return v.value
}

func (v *NullableGetPendingVirtualMachineCloudInitChangesResponseContent) Set(val *GetPendingVirtualMachineCloudInitChangesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPendingVirtualMachineCloudInitChangesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPendingVirtualMachineCloudInitChangesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPendingVirtualMachineCloudInitChangesResponseContent(val *GetPendingVirtualMachineCloudInitChangesResponseContent) *NullableGetPendingVirtualMachineCloudInitChangesResponseContent {
	return &NullableGetPendingVirtualMachineCloudInitChangesResponseContent{value: val, isSet: true}
}

func (v NullableGetPendingVirtualMachineCloudInitChangesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPendingVirtualMachineCloudInitChangesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

