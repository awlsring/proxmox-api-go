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

// checks if the CreateVirtualMachineTemplateRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVirtualMachineTemplateRequestContent{}

// CreateVirtualMachineTemplateRequestContent struct for CreateVirtualMachineTemplateRequestContent
type CreateVirtualMachineTemplateRequestContent struct {
	Disk *VirtualMachineDiskTarget `json:"disk,omitempty"`
}

// NewCreateVirtualMachineTemplateRequestContent instantiates a new CreateVirtualMachineTemplateRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVirtualMachineTemplateRequestContent() *CreateVirtualMachineTemplateRequestContent {
	this := CreateVirtualMachineTemplateRequestContent{}
	return &this
}

// NewCreateVirtualMachineTemplateRequestContentWithDefaults instantiates a new CreateVirtualMachineTemplateRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVirtualMachineTemplateRequestContentWithDefaults() *CreateVirtualMachineTemplateRequestContent {
	this := CreateVirtualMachineTemplateRequestContent{}
	return &this
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *CreateVirtualMachineTemplateRequestContent) GetDisk() VirtualMachineDiskTarget {
	if o == nil || isNil(o.Disk) {
		var ret VirtualMachineDiskTarget
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualMachineTemplateRequestContent) GetDiskOk() (*VirtualMachineDiskTarget, bool) {
	if o == nil || isNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *CreateVirtualMachineTemplateRequestContent) HasDisk() bool {
	if o != nil && !isNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given VirtualMachineDiskTarget and assigns it to the Disk field.
func (o *CreateVirtualMachineTemplateRequestContent) SetDisk(v VirtualMachineDiskTarget) {
	o.Disk = &v
}

func (o CreateVirtualMachineTemplateRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVirtualMachineTemplateRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	return toSerialize, nil
}

type NullableCreateVirtualMachineTemplateRequestContent struct {
	value *CreateVirtualMachineTemplateRequestContent
	isSet bool
}

func (v NullableCreateVirtualMachineTemplateRequestContent) Get() *CreateVirtualMachineTemplateRequestContent {
	return v.value
}

func (v *NullableCreateVirtualMachineTemplateRequestContent) Set(val *CreateVirtualMachineTemplateRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVirtualMachineTemplateRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVirtualMachineTemplateRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVirtualMachineTemplateRequestContent(val *CreateVirtualMachineTemplateRequestContent) *NullableCreateVirtualMachineTemplateRequestContent {
	return &NullableCreateVirtualMachineTemplateRequestContent{value: val, isSet: true}
}

func (v NullableCreateVirtualMachineTemplateRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVirtualMachineTemplateRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

