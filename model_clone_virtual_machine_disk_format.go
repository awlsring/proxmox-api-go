/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CloneVirtualMachineDiskFormat the model 'CloneVirtualMachineDiskFormat'
type CloneVirtualMachineDiskFormat string

// List of CloneVirtualMachineDiskFormat
const (
	CLONEVIRTUALMACHINEDISKFORMAT_RAW CloneVirtualMachineDiskFormat = "raw"
	CLONEVIRTUALMACHINEDISKFORMAT_QCOW2 CloneVirtualMachineDiskFormat = "qcow2"
	CLONEVIRTUALMACHINEDISKFORMAT_VMDK CloneVirtualMachineDiskFormat = "vmdk"
)

// All allowed values of CloneVirtualMachineDiskFormat enum
var AllowedCloneVirtualMachineDiskFormatEnumValues = []CloneVirtualMachineDiskFormat{
	"raw",
	"qcow2",
	"vmdk",
}

func (v *CloneVirtualMachineDiskFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloneVirtualMachineDiskFormat(value)
	for _, existing := range AllowedCloneVirtualMachineDiskFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloneVirtualMachineDiskFormat", value)
}

// NewCloneVirtualMachineDiskFormatFromValue returns a pointer to a valid CloneVirtualMachineDiskFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloneVirtualMachineDiskFormatFromValue(v string) (*CloneVirtualMachineDiskFormat, error) {
	ev := CloneVirtualMachineDiskFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloneVirtualMachineDiskFormat: valid values are %v", v, AllowedCloneVirtualMachineDiskFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloneVirtualMachineDiskFormat) IsValid() bool {
	for _, existing := range AllowedCloneVirtualMachineDiskFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloneVirtualMachineDiskFormat value
func (v CloneVirtualMachineDiskFormat) Ptr() *CloneVirtualMachineDiskFormat {
	return &v
}

type NullableCloneVirtualMachineDiskFormat struct {
	value *CloneVirtualMachineDiskFormat
	isSet bool
}

func (v NullableCloneVirtualMachineDiskFormat) Get() *CloneVirtualMachineDiskFormat {
	return v.value
}

func (v *NullableCloneVirtualMachineDiskFormat) Set(val *CloneVirtualMachineDiskFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneVirtualMachineDiskFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneVirtualMachineDiskFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneVirtualMachineDiskFormat(val *CloneVirtualMachineDiskFormat) *NullableCloneVirtualMachineDiskFormat {
	return &NullableCloneVirtualMachineDiskFormat{value: val, isSet: true}
}

func (v NullableCloneVirtualMachineDiskFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneVirtualMachineDiskFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

