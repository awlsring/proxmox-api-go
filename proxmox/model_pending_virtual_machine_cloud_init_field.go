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

// checks if the PendingVirtualMachineCloudInitField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingVirtualMachineCloudInitField{}

// PendingVirtualMachineCloudInitField struct for PendingVirtualMachineCloudInitField
type PendingVirtualMachineCloudInitField struct {
	Key string `json:"key"`
	New *string `json:"new,omitempty"`
	Old *string `json:"old,omitempty"`
}

// NewPendingVirtualMachineCloudInitField instantiates a new PendingVirtualMachineCloudInitField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingVirtualMachineCloudInitField(key string) *PendingVirtualMachineCloudInitField {
	this := PendingVirtualMachineCloudInitField{}
	this.Key = key
	return &this
}

// NewPendingVirtualMachineCloudInitFieldWithDefaults instantiates a new PendingVirtualMachineCloudInitField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingVirtualMachineCloudInitFieldWithDefaults() *PendingVirtualMachineCloudInitField {
	this := PendingVirtualMachineCloudInitField{}
	return &this
}

// GetKey returns the Key field value
func (o *PendingVirtualMachineCloudInitField) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PendingVirtualMachineCloudInitField) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PendingVirtualMachineCloudInitField) SetKey(v string) {
	o.Key = v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *PendingVirtualMachineCloudInitField) GetNew() string {
	if o == nil || IsNil(o.New) {
		var ret string
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingVirtualMachineCloudInitField) GetNewOk() (*string, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *PendingVirtualMachineCloudInitField) HasNew() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given string and assigns it to the New field.
func (o *PendingVirtualMachineCloudInitField) SetNew(v string) {
	o.New = &v
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *PendingVirtualMachineCloudInitField) GetOld() string {
	if o == nil || IsNil(o.Old) {
		var ret string
		return ret
	}
	return *o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingVirtualMachineCloudInitField) GetOldOk() (*string, bool) {
	if o == nil || IsNil(o.Old) {
		return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *PendingVirtualMachineCloudInitField) HasOld() bool {
	if o != nil && !IsNil(o.Old) {
		return true
	}

	return false
}

// SetOld gets a reference to the given string and assigns it to the Old field.
func (o *PendingVirtualMachineCloudInitField) SetOld(v string) {
	o.Old = &v
}

func (o PendingVirtualMachineCloudInitField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingVirtualMachineCloudInitField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.New) {
		toSerialize["new"] = o.New
	}
	if !IsNil(o.Old) {
		toSerialize["old"] = o.Old
	}
	return toSerialize, nil
}

type NullablePendingVirtualMachineCloudInitField struct {
	value *PendingVirtualMachineCloudInitField
	isSet bool
}

func (v NullablePendingVirtualMachineCloudInitField) Get() *PendingVirtualMachineCloudInitField {
	return v.value
}

func (v *NullablePendingVirtualMachineCloudInitField) Set(val *PendingVirtualMachineCloudInitField) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingVirtualMachineCloudInitField) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingVirtualMachineCloudInitField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingVirtualMachineCloudInitField(val *PendingVirtualMachineCloudInitField) *NullablePendingVirtualMachineCloudInitField {
	return &NullablePendingVirtualMachineCloudInitField{value: val, isSet: true}
}

func (v NullablePendingVirtualMachineCloudInitField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingVirtualMachineCloudInitField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


