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

// VirtualMachineHighAvailabilityStatus struct for VirtualMachineHighAvailabilityStatus
type VirtualMachineHighAvailabilityStatus struct {
	// An integer used to represent a boolean. 0 is false, 1 is true.
	Managed float32 `json:"managed"`
}

// NewVirtualMachineHighAvailabilityStatus instantiates a new VirtualMachineHighAvailabilityStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineHighAvailabilityStatus(managed float32) *VirtualMachineHighAvailabilityStatus {
	this := VirtualMachineHighAvailabilityStatus{}
	this.Managed = managed
	return &this
}

// NewVirtualMachineHighAvailabilityStatusWithDefaults instantiates a new VirtualMachineHighAvailabilityStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineHighAvailabilityStatusWithDefaults() *VirtualMachineHighAvailabilityStatus {
	this := VirtualMachineHighAvailabilityStatus{}
	return &this
}

// GetManaged returns the Managed field value
func (o *VirtualMachineHighAvailabilityStatus) GetManaged() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineHighAvailabilityStatus) GetManagedOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *VirtualMachineHighAvailabilityStatus) SetManaged(v float32) {
	o.Managed = v
}

func (o VirtualMachineHighAvailabilityStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["managed"] = o.Managed
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualMachineHighAvailabilityStatus struct {
	value *VirtualMachineHighAvailabilityStatus
	isSet bool
}

func (v NullableVirtualMachineHighAvailabilityStatus) Get() *VirtualMachineHighAvailabilityStatus {
	return v.value
}

func (v *NullableVirtualMachineHighAvailabilityStatus) Set(val *VirtualMachineHighAvailabilityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineHighAvailabilityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineHighAvailabilityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineHighAvailabilityStatus(val *VirtualMachineHighAvailabilityStatus) *NullableVirtualMachineHighAvailabilityStatus {
	return &NullableVirtualMachineHighAvailabilityStatus{value: val, isSet: true}
}

func (v NullableVirtualMachineHighAvailabilityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineHighAvailabilityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


