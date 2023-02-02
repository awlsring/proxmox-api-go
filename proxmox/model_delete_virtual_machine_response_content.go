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

// checks if the DeleteVirtualMachineResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVirtualMachineResponseContent{}

// DeleteVirtualMachineResponseContent struct for DeleteVirtualMachineResponseContent
type DeleteVirtualMachineResponseContent struct {
	Data string `json:"data"`
}

// NewDeleteVirtualMachineResponseContent instantiates a new DeleteVirtualMachineResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVirtualMachineResponseContent(data string) *DeleteVirtualMachineResponseContent {
	this := DeleteVirtualMachineResponseContent{}
	this.Data = data
	return &this
}

// NewDeleteVirtualMachineResponseContentWithDefaults instantiates a new DeleteVirtualMachineResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVirtualMachineResponseContentWithDefaults() *DeleteVirtualMachineResponseContent {
	this := DeleteVirtualMachineResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteVirtualMachineResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteVirtualMachineResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteVirtualMachineResponseContent) SetData(v string) {
	o.Data = v
}

func (o DeleteVirtualMachineResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVirtualMachineResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableDeleteVirtualMachineResponseContent struct {
	value *DeleteVirtualMachineResponseContent
	isSet bool
}

func (v NullableDeleteVirtualMachineResponseContent) Get() *DeleteVirtualMachineResponseContent {
	return v.value
}

func (v *NullableDeleteVirtualMachineResponseContent) Set(val *DeleteVirtualMachineResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVirtualMachineResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVirtualMachineResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVirtualMachineResponseContent(val *DeleteVirtualMachineResponseContent) *NullableDeleteVirtualMachineResponseContent {
	return &NullableDeleteVirtualMachineResponseContent{value: val, isSet: true}
}

func (v NullableDeleteVirtualMachineResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVirtualMachineResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


