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

// checks if the GetVirtualMachineFeatureSupportResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVirtualMachineFeatureSupportResponseContent{}

// GetVirtualMachineFeatureSupportResponseContent struct for GetVirtualMachineFeatureSupportResponseContent
type GetVirtualMachineFeatureSupportResponseContent struct {
	Data VirtualMachineFeatureSupportSummary `json:"data"`
}

// NewGetVirtualMachineFeatureSupportResponseContent instantiates a new GetVirtualMachineFeatureSupportResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualMachineFeatureSupportResponseContent(data VirtualMachineFeatureSupportSummary) *GetVirtualMachineFeatureSupportResponseContent {
	this := GetVirtualMachineFeatureSupportResponseContent{}
	this.Data = data
	return &this
}

// NewGetVirtualMachineFeatureSupportResponseContentWithDefaults instantiates a new GetVirtualMachineFeatureSupportResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualMachineFeatureSupportResponseContentWithDefaults() *GetVirtualMachineFeatureSupportResponseContent {
	this := GetVirtualMachineFeatureSupportResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetVirtualMachineFeatureSupportResponseContent) GetData() VirtualMachineFeatureSupportSummary {
	if o == nil {
		var ret VirtualMachineFeatureSupportSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetVirtualMachineFeatureSupportResponseContent) GetDataOk() (*VirtualMachineFeatureSupportSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetVirtualMachineFeatureSupportResponseContent) SetData(v VirtualMachineFeatureSupportSummary) {
	o.Data = v
}

func (o GetVirtualMachineFeatureSupportResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVirtualMachineFeatureSupportResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetVirtualMachineFeatureSupportResponseContent struct {
	value *GetVirtualMachineFeatureSupportResponseContent
	isSet bool
}

func (v NullableGetVirtualMachineFeatureSupportResponseContent) Get() *GetVirtualMachineFeatureSupportResponseContent {
	return v.value
}

func (v *NullableGetVirtualMachineFeatureSupportResponseContent) Set(val *GetVirtualMachineFeatureSupportResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualMachineFeatureSupportResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualMachineFeatureSupportResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualMachineFeatureSupportResponseContent(val *GetVirtualMachineFeatureSupportResponseContent) *NullableGetVirtualMachineFeatureSupportResponseContent {
	return &NullableGetVirtualMachineFeatureSupportResponseContent{value: val, isSet: true}
}

func (v NullableGetVirtualMachineFeatureSupportResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualMachineFeatureSupportResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

