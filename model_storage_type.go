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

// StorageType the model 'StorageType'
type StorageType string

// List of StorageType
const (
	STORAGETYPE_ZFSPOOL StorageType = "zfspool"
	STORAGETYPE_DIR StorageType = "dir"
	STORAGETYPE_BTRFS StorageType = "btrfs"
	STORAGETYPE_NFS StorageType = "nfs"
	STORAGETYPE_CIFS StorageType = "cifs"
	STORAGETYPE_PBS StorageType = "pbs"
	STORAGETYPE_GLUSTERFS StorageType = "glusterfs"
	STORAGETYPE_CEPHFS StorageType = "cephfs"
	STORAGETYPE_LVMTHIN StorageType = "lvmthin"
	STORAGETYPE_LVM StorageType = "lvm"
	STORAGETYPE_ISCSI StorageType = "iscsi"
	STORAGETYPE_ISCSIDIRECT StorageType = "iscsidirect"
	STORAGETYPE_RBD StorageType = "rbd"
	STORAGETYPE_ZFS StorageType = "zfs"
)

// All allowed values of StorageType enum
var AllowedStorageTypeEnumValues = []StorageType{
	"zfspool",
	"dir",
	"btrfs",
	"nfs",
	"cifs",
	"pbs",
	"glusterfs",
	"cephfs",
	"lvmthin",
	"lvm",
	"iscsi",
	"iscsidirect",
	"rbd",
	"zfs",
}

func (v *StorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageType(value)
	for _, existing := range AllowedStorageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageType", value)
}

// NewStorageTypeFromValue returns a pointer to a valid StorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageTypeFromValue(v string) (*StorageType, error) {
	ev := StorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageType: valid values are %v", v, AllowedStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageType) IsValid() bool {
	for _, existing := range AllowedStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageType value
func (v StorageType) Ptr() *StorageType {
	return &v
}

type NullableStorageType struct {
	value *StorageType
	isSet bool
}

func (v NullableStorageType) Get() *StorageType {
	return v.value
}

func (v *NullableStorageType) Set(val *StorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageType(val *StorageType) *NullableStorageType {
	return &NullableStorageType{value: val, isSet: true}
}

func (v NullableStorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

