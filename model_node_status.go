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

// NodeStatus the model 'NodeStatus'
type NodeStatus string

// List of NodeStatus
const (
	NODESTATUS_ONLINE NodeStatus = "online"
	NODESTATUS_OFFLINE NodeStatus = "offline"
)

// All allowed values of NodeStatus enum
var AllowedNodeStatusEnumValues = []NodeStatus{
	"online",
	"offline",
}

func (v *NodeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NodeStatus(value)
	for _, existing := range AllowedNodeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NodeStatus", value)
}

// NewNodeStatusFromValue returns a pointer to a valid NodeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNodeStatusFromValue(v string) (*NodeStatus, error) {
	ev := NodeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NodeStatus: valid values are %v", v, AllowedNodeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NodeStatus) IsValid() bool {
	for _, existing := range AllowedNodeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeStatus value
func (v NodeStatus) Ptr() *NodeStatus {
	return &v
}

type NullableNodeStatus struct {
	value *NodeStatus
	isSet bool
}

func (v NullableNodeStatus) Get() *NodeStatus {
	return v.value
}

func (v *NullableNodeStatus) Set(val *NodeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeStatus(val *NodeStatus) *NullableNodeStatus {
	return &NullableNodeStatus{value: val, isSet: true}
}

func (v NullableNodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

