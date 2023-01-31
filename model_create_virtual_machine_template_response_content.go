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

// CreateVirtualMachineTemplateResponseContent struct for CreateVirtualMachineTemplateResponseContent
type CreateVirtualMachineTemplateResponseContent struct {
	Data *string `json:"data,omitempty"`
}

// NewCreateVirtualMachineTemplateResponseContent instantiates a new CreateVirtualMachineTemplateResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVirtualMachineTemplateResponseContent() *CreateVirtualMachineTemplateResponseContent {
	this := CreateVirtualMachineTemplateResponseContent{}
	return &this
}

// NewCreateVirtualMachineTemplateResponseContentWithDefaults instantiates a new CreateVirtualMachineTemplateResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVirtualMachineTemplateResponseContentWithDefaults() *CreateVirtualMachineTemplateResponseContent {
	this := CreateVirtualMachineTemplateResponseContent{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateVirtualMachineTemplateResponseContent) GetData() string {
	if o == nil || isNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualMachineTemplateResponseContent) GetDataOk() (*string, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateVirtualMachineTemplateResponseContent) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *CreateVirtualMachineTemplateResponseContent) SetData(v string) {
	o.Data = &v
}

func (o CreateVirtualMachineTemplateResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVirtualMachineTemplateResponseContent struct {
	value *CreateVirtualMachineTemplateResponseContent
	isSet bool
}

func (v NullableCreateVirtualMachineTemplateResponseContent) Get() *CreateVirtualMachineTemplateResponseContent {
	return v.value
}

func (v *NullableCreateVirtualMachineTemplateResponseContent) Set(val *CreateVirtualMachineTemplateResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVirtualMachineTemplateResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVirtualMachineTemplateResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVirtualMachineTemplateResponseContent(val *CreateVirtualMachineTemplateResponseContent) *NullableCreateVirtualMachineTemplateResponseContent {
	return &NullableCreateVirtualMachineTemplateResponseContent{value: val, isSet: true}
}

func (v NullableCreateVirtualMachineTemplateResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVirtualMachineTemplateResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

