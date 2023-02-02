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

// checks if the RepositorySummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositorySummary{}

// RepositorySummary struct for RepositorySummary
type RepositorySummary struct {
	Description string `json:"description"`
	Status float32 `json:"status"`
	Name string `json:"name"`
	Handle string `json:"handle"`
}

// NewRepositorySummary instantiates a new RepositorySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositorySummary(description string, status float32, name string, handle string) *RepositorySummary {
	this := RepositorySummary{}
	this.Description = description
	this.Status = status
	this.Name = name
	this.Handle = handle
	return &this
}

// NewRepositorySummaryWithDefaults instantiates a new RepositorySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositorySummaryWithDefaults() *RepositorySummary {
	this := RepositorySummary{}
	return &this
}

// GetDescription returns the Description field value
func (o *RepositorySummary) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RepositorySummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RepositorySummary) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *RepositorySummary) GetStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RepositorySummary) GetStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RepositorySummary) SetStatus(v float32) {
	o.Status = v
}

// GetName returns the Name field value
func (o *RepositorySummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RepositorySummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RepositorySummary) SetName(v string) {
	o.Name = v
}

// GetHandle returns the Handle field value
func (o *RepositorySummary) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *RepositorySummary) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *RepositorySummary) SetHandle(v string) {
	o.Handle = v
}

func (o RepositorySummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositorySummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["status"] = o.Status
	toSerialize["name"] = o.Name
	toSerialize["handle"] = o.Handle
	return toSerialize, nil
}

type NullableRepositorySummary struct {
	value *RepositorySummary
	isSet bool
}

func (v NullableRepositorySummary) Get() *RepositorySummary {
	return v.value
}

func (v *NullableRepositorySummary) Set(val *RepositorySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositorySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositorySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositorySummary(val *RepositorySummary) *NullableRepositorySummary {
	return &NullableRepositorySummary{value: val, isSet: true}
}

func (v NullableRepositorySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositorySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


