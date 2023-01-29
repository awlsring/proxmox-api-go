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

// NetworkInterfaceBondMode the model 'NetworkInterfaceBondMode'
type NetworkInterfaceBondMode string

// List of NetworkInterfaceBondMode
const (
	NETWORKINTERFACEBONDMODE_BALANCE_RR NetworkInterfaceBondMode = "balance-rr"
	NETWORKINTERFACEBONDMODE_ACTIVE_BACKUP NetworkInterfaceBondMode = "active-backup"
	NETWORKINTERFACEBONDMODE_BALANCE_XOR NetworkInterfaceBondMode = "balance-xor"
	NETWORKINTERFACEBONDMODE_BROADCAST NetworkInterfaceBondMode = "broadcast"
	NETWORKINTERFACEBONDMODE__802_3AD NetworkInterfaceBondMode = "802.3ad"
	NETWORKINTERFACEBONDMODE_BALANCE_TLB NetworkInterfaceBondMode = "balance-tlb"
	NETWORKINTERFACEBONDMODE_BALANCE_ALB NetworkInterfaceBondMode = "balance-alb"
	NETWORKINTERFACEBONDMODE_BALANCE_SLB NetworkInterfaceBondMode = "balance-slb"
	NETWORKINTERFACEBONDMODE_LACP_BALANCE_SLB NetworkInterfaceBondMode = "lacp-balance-slb"
	NETWORKINTERFACEBONDMODE_LACP_BALANCE_TCP NetworkInterfaceBondMode = "lacp-balance-tcp"
)

// All allowed values of NetworkInterfaceBondMode enum
var AllowedNetworkInterfaceBondModeEnumValues = []NetworkInterfaceBondMode{
	"balance-rr",
	"active-backup",
	"balance-xor",
	"broadcast",
	"802.3ad",
	"balance-tlb",
	"balance-alb",
	"balance-slb",
	"lacp-balance-slb",
	"lacp-balance-tcp",
}

func (v *NetworkInterfaceBondMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkInterfaceBondMode(value)
	for _, existing := range AllowedNetworkInterfaceBondModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkInterfaceBondMode", value)
}

// NewNetworkInterfaceBondModeFromValue returns a pointer to a valid NetworkInterfaceBondMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkInterfaceBondModeFromValue(v string) (*NetworkInterfaceBondMode, error) {
	ev := NetworkInterfaceBondMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkInterfaceBondMode: valid values are %v", v, AllowedNetworkInterfaceBondModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkInterfaceBondMode) IsValid() bool {
	for _, existing := range AllowedNetworkInterfaceBondModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkInterfaceBondMode value
func (v NetworkInterfaceBondMode) Ptr() *NetworkInterfaceBondMode {
	return &v
}

type NullableNetworkInterfaceBondMode struct {
	value *NetworkInterfaceBondMode
	isSet bool
}

func (v NullableNetworkInterfaceBondMode) Get() *NetworkInterfaceBondMode {
	return v.value
}

func (v *NullableNetworkInterfaceBondMode) Set(val *NetworkInterfaceBondMode) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaceBondMode) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaceBondMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaceBondMode(val *NetworkInterfaceBondMode) *NullableNetworkInterfaceBondMode {
	return &NullableNetworkInterfaceBondMode{value: val, isSet: true}
}

func (v NullableNetworkInterfaceBondMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaceBondMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

