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

// VirtualMachineTemplateDiskTarget the model 'VirtualMachineTemplateDiskTarget'
type VirtualMachineTemplateDiskTarget string

// List of VirtualMachineTemplateDiskTarget
const (
	VIRTUALMACHINETEMPLATEDISKTARGET_IDE0 VirtualMachineTemplateDiskTarget = "ide0"
	VIRTUALMACHINETEMPLATEDISKTARGET_IDE1 VirtualMachineTemplateDiskTarget = "ide1"
	VIRTUALMACHINETEMPLATEDISKTARGET_IDE2 VirtualMachineTemplateDiskTarget = "ide2"
	VIRTUALMACHINETEMPLATEDISKTARGET_IDE3 VirtualMachineTemplateDiskTarget = "ide3"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI0 VirtualMachineTemplateDiskTarget = "scsi0"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI1 VirtualMachineTemplateDiskTarget = "scsi1"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI2 VirtualMachineTemplateDiskTarget = "scsi2"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI3 VirtualMachineTemplateDiskTarget = "scsi3"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI4 VirtualMachineTemplateDiskTarget = "scsi4"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI5 VirtualMachineTemplateDiskTarget = "scsi5"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI6 VirtualMachineTemplateDiskTarget = "scsi6"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI7 VirtualMachineTemplateDiskTarget = "scsi7"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI8 VirtualMachineTemplateDiskTarget = "scsi8"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI9 VirtualMachineTemplateDiskTarget = "scsi9"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI10 VirtualMachineTemplateDiskTarget = "scsi10"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI11 VirtualMachineTemplateDiskTarget = "scsi11"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI12 VirtualMachineTemplateDiskTarget = "scsi12"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI13 VirtualMachineTemplateDiskTarget = "scsi13"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI14 VirtualMachineTemplateDiskTarget = "scsi14"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI15 VirtualMachineTemplateDiskTarget = "scsi15"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI16 VirtualMachineTemplateDiskTarget = "scsi16"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI17 VirtualMachineTemplateDiskTarget = "scsi17"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI18 VirtualMachineTemplateDiskTarget = "scsi18"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI19 VirtualMachineTemplateDiskTarget = "scsi19"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI20 VirtualMachineTemplateDiskTarget = "scsi20"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI21 VirtualMachineTemplateDiskTarget = "scsi21"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI22 VirtualMachineTemplateDiskTarget = "scsi22"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI23 VirtualMachineTemplateDiskTarget = "scsi23"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI24 VirtualMachineTemplateDiskTarget = "scsi24"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI25 VirtualMachineTemplateDiskTarget = "scsi25"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI26 VirtualMachineTemplateDiskTarget = "scsi26"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI27 VirtualMachineTemplateDiskTarget = "scsi27"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI28 VirtualMachineTemplateDiskTarget = "scsi28"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI29 VirtualMachineTemplateDiskTarget = "scsi29"
	VIRTUALMACHINETEMPLATEDISKTARGET_SCSI30 VirtualMachineTemplateDiskTarget = "scsi30"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO0 VirtualMachineTemplateDiskTarget = "virtio0"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO1 VirtualMachineTemplateDiskTarget = "virtio1"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO2 VirtualMachineTemplateDiskTarget = "virtio2"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO3 VirtualMachineTemplateDiskTarget = "virtio3"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO4 VirtualMachineTemplateDiskTarget = "virtio4"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO5 VirtualMachineTemplateDiskTarget = "virtio5"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO6 VirtualMachineTemplateDiskTarget = "virtio6"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO7 VirtualMachineTemplateDiskTarget = "virtio7"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO8 VirtualMachineTemplateDiskTarget = "virtio8"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO9 VirtualMachineTemplateDiskTarget = "virtio9"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO10 VirtualMachineTemplateDiskTarget = "virtio10"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO11 VirtualMachineTemplateDiskTarget = "virtio11"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO12 VirtualMachineTemplateDiskTarget = "virtio12"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO13 VirtualMachineTemplateDiskTarget = "virtio13"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO14 VirtualMachineTemplateDiskTarget = "virtio14"
	VIRTUALMACHINETEMPLATEDISKTARGET_VIRTIO15 VirtualMachineTemplateDiskTarget = "virtio15"
	VIRTUALMACHINETEMPLATEDISKTARGET_SATA0 VirtualMachineTemplateDiskTarget = "sata0"
	VIRTUALMACHINETEMPLATEDISKTARGET_SATA1 VirtualMachineTemplateDiskTarget = "sata1"
	VIRTUALMACHINETEMPLATEDISKTARGET_SATA2 VirtualMachineTemplateDiskTarget = "sata2"
	VIRTUALMACHINETEMPLATEDISKTARGET_SATA3 VirtualMachineTemplateDiskTarget = "sata3"
	VIRTUALMACHINETEMPLATEDISKTARGET_SATA4 VirtualMachineTemplateDiskTarget = "sata4"
	VIRTUALMACHINETEMPLATEDISKTARGET_SATA5 VirtualMachineTemplateDiskTarget = "sata5"
	VIRTUALMACHINETEMPLATEDISKTARGET_EFIDISK0 VirtualMachineTemplateDiskTarget = "efidisk0"
	VIRTUALMACHINETEMPLATEDISKTARGET_TPMSTATE0 VirtualMachineTemplateDiskTarget = "tpmstate0"
)

// All allowed values of VirtualMachineTemplateDiskTarget enum
var AllowedVirtualMachineTemplateDiskTargetEnumValues = []VirtualMachineTemplateDiskTarget{
	"ide0",
	"ide1",
	"ide2",
	"ide3",
	"scsi0",
	"scsi1",
	"scsi2",
	"scsi3",
	"scsi4",
	"scsi5",
	"scsi6",
	"scsi7",
	"scsi8",
	"scsi9",
	"scsi10",
	"scsi11",
	"scsi12",
	"scsi13",
	"scsi14",
	"scsi15",
	"scsi16",
	"scsi17",
	"scsi18",
	"scsi19",
	"scsi20",
	"scsi21",
	"scsi22",
	"scsi23",
	"scsi24",
	"scsi25",
	"scsi26",
	"scsi27",
	"scsi28",
	"scsi29",
	"scsi30",
	"virtio0",
	"virtio1",
	"virtio2",
	"virtio3",
	"virtio4",
	"virtio5",
	"virtio6",
	"virtio7",
	"virtio8",
	"virtio9",
	"virtio10",
	"virtio11",
	"virtio12",
	"virtio13",
	"virtio14",
	"virtio15",
	"sata0",
	"sata1",
	"sata2",
	"sata3",
	"sata4",
	"sata5",
	"efidisk0",
	"tpmstate0",
}

func (v *VirtualMachineTemplateDiskTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachineTemplateDiskTarget(value)
	for _, existing := range AllowedVirtualMachineTemplateDiskTargetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachineTemplateDiskTarget", value)
}

// NewVirtualMachineTemplateDiskTargetFromValue returns a pointer to a valid VirtualMachineTemplateDiskTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualMachineTemplateDiskTargetFromValue(v string) (*VirtualMachineTemplateDiskTarget, error) {
	ev := VirtualMachineTemplateDiskTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualMachineTemplateDiskTarget: valid values are %v", v, AllowedVirtualMachineTemplateDiskTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualMachineTemplateDiskTarget) IsValid() bool {
	for _, existing := range AllowedVirtualMachineTemplateDiskTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualMachineTemplateDiskTarget value
func (v VirtualMachineTemplateDiskTarget) Ptr() *VirtualMachineTemplateDiskTarget {
	return &v
}

type NullableVirtualMachineTemplateDiskTarget struct {
	value *VirtualMachineTemplateDiskTarget
	isSet bool
}

func (v NullableVirtualMachineTemplateDiskTarget) Get() *VirtualMachineTemplateDiskTarget {
	return v.value
}

func (v *NullableVirtualMachineTemplateDiskTarget) Set(val *VirtualMachineTemplateDiskTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineTemplateDiskTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineTemplateDiskTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineTemplateDiskTarget(val *VirtualMachineTemplateDiskTarget) *NullableVirtualMachineTemplateDiskTarget {
	return &NullableVirtualMachineTemplateDiskTarget{value: val, isSet: true}
}

func (v NullableVirtualMachineTemplateDiskTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineTemplateDiskTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
