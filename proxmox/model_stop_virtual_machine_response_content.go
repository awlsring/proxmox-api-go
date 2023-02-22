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

// checks if the StopVirtualMachineResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopVirtualMachineResponseContent{}

// StopVirtualMachineResponseContent struct for StopVirtualMachineResponseContent
type StopVirtualMachineResponseContent struct {
	Data string `json:"data"`
}

// NewStopVirtualMachineResponseContent instantiates a new StopVirtualMachineResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopVirtualMachineResponseContent(data string) *StopVirtualMachineResponseContent {
	this := StopVirtualMachineResponseContent{}
	this.Data = data
	return &this
}

// NewStopVirtualMachineResponseContentWithDefaults instantiates a new StopVirtualMachineResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopVirtualMachineResponseContentWithDefaults() *StopVirtualMachineResponseContent {
	this := StopVirtualMachineResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *StopVirtualMachineResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StopVirtualMachineResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StopVirtualMachineResponseContent) SetData(v string) {
	o.Data = v
}

func (o StopVirtualMachineResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopVirtualMachineResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableStopVirtualMachineResponseContent struct {
	value *StopVirtualMachineResponseContent
	isSet bool
}

func (v NullableStopVirtualMachineResponseContent) Get() *StopVirtualMachineResponseContent {
	return v.value
}

func (v *NullableStopVirtualMachineResponseContent) Set(val *StopVirtualMachineResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableStopVirtualMachineResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableStopVirtualMachineResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopVirtualMachineResponseContent(val *StopVirtualMachineResponseContent) *NullableStopVirtualMachineResponseContent {
	return &NullableStopVirtualMachineResponseContent{value: val, isSet: true}
}

func (v NullableStopVirtualMachineResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopVirtualMachineResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

