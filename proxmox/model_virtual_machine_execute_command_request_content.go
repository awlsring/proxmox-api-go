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

// checks if the VirtualMachineExecuteCommandRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineExecuteCommandRequestContent{}

// VirtualMachineExecuteCommandRequestContent struct for VirtualMachineExecuteCommandRequestContent
type VirtualMachineExecuteCommandRequestContent struct {
	Command *string `json:"command,omitempty"`
	InputData *string `json:"input-data,omitempty"`
}

// NewVirtualMachineExecuteCommandRequestContent instantiates a new VirtualMachineExecuteCommandRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineExecuteCommandRequestContent() *VirtualMachineExecuteCommandRequestContent {
	this := VirtualMachineExecuteCommandRequestContent{}
	return &this
}

// NewVirtualMachineExecuteCommandRequestContentWithDefaults instantiates a new VirtualMachineExecuteCommandRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineExecuteCommandRequestContentWithDefaults() *VirtualMachineExecuteCommandRequestContent {
	this := VirtualMachineExecuteCommandRequestContent{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *VirtualMachineExecuteCommandRequestContent) GetCommand() string {
	if o == nil || isNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineExecuteCommandRequestContent) GetCommandOk() (*string, bool) {
	if o == nil || isNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *VirtualMachineExecuteCommandRequestContent) HasCommand() bool {
	if o != nil && !isNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *VirtualMachineExecuteCommandRequestContent) SetCommand(v string) {
	o.Command = &v
}

// GetInputData returns the InputData field value if set, zero value otherwise.
func (o *VirtualMachineExecuteCommandRequestContent) GetInputData() string {
	if o == nil || isNil(o.InputData) {
		var ret string
		return ret
	}
	return *o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineExecuteCommandRequestContent) GetInputDataOk() (*string, bool) {
	if o == nil || isNil(o.InputData) {
		return nil, false
	}
	return o.InputData, true
}

// HasInputData returns a boolean if a field has been set.
func (o *VirtualMachineExecuteCommandRequestContent) HasInputData() bool {
	if o != nil && !isNil(o.InputData) {
		return true
	}

	return false
}

// SetInputData gets a reference to the given string and assigns it to the InputData field.
func (o *VirtualMachineExecuteCommandRequestContent) SetInputData(v string) {
	o.InputData = &v
}

func (o VirtualMachineExecuteCommandRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineExecuteCommandRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !isNil(o.InputData) {
		toSerialize["input-data"] = o.InputData
	}
	return toSerialize, nil
}

type NullableVirtualMachineExecuteCommandRequestContent struct {
	value *VirtualMachineExecuteCommandRequestContent
	isSet bool
}

func (v NullableVirtualMachineExecuteCommandRequestContent) Get() *VirtualMachineExecuteCommandRequestContent {
	return v.value
}

func (v *NullableVirtualMachineExecuteCommandRequestContent) Set(val *VirtualMachineExecuteCommandRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineExecuteCommandRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineExecuteCommandRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineExecuteCommandRequestContent(val *VirtualMachineExecuteCommandRequestContent) *NullableVirtualMachineExecuteCommandRequestContent {
	return &NullableVirtualMachineExecuteCommandRequestContent{value: val, isSet: true}
}

func (v NullableVirtualMachineExecuteCommandRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineExecuteCommandRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


