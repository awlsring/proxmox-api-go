/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package proxmox

import (
	"encoding/json"
	"fmt"
)

// VirtualMachineStatus the model 'VirtualMachineStatus'
type VirtualMachineStatus string

// List of VirtualMachineStatus
const (
	VIRTUALMACHINESTATUS_RUNNING VirtualMachineStatus = "running"
	VIRTUALMACHINESTATUS_STOPPED VirtualMachineStatus = "stopped"
)

// All allowed values of VirtualMachineStatus enum
var AllowedVirtualMachineStatusEnumValues = []VirtualMachineStatus{
	"running",
	"stopped",
}

func (v *VirtualMachineStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachineStatus(value)
	for _, existing := range AllowedVirtualMachineStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachineStatus", value)
}

// NewVirtualMachineStatusFromValue returns a pointer to a valid VirtualMachineStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualMachineStatusFromValue(v string) (*VirtualMachineStatus, error) {
	ev := VirtualMachineStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualMachineStatus: valid values are %v", v, AllowedVirtualMachineStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualMachineStatus) IsValid() bool {
	for _, existing := range AllowedVirtualMachineStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualMachineStatus value
func (v VirtualMachineStatus) Ptr() *VirtualMachineStatus {
	return &v
}

type NullableVirtualMachineStatus struct {
	value *VirtualMachineStatus
	isSet bool
}

func (v NullableVirtualMachineStatus) Get() *VirtualMachineStatus {
	return v.value
}

func (v *NullableVirtualMachineStatus) Set(val *VirtualMachineStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineStatus(val *VirtualMachineStatus) *NullableVirtualMachineStatus {
	return &NullableVirtualMachineStatus{value: val, isSet: true}
}

func (v NullableVirtualMachineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
