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

// checks if the ListPackagesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPackagesResponseContent{}

// ListPackagesResponseContent struct for ListPackagesResponseContent
type ListPackagesResponseContent struct {
	Data []PackageSummary `json:"data"`
}

// NewListPackagesResponseContent instantiates a new ListPackagesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackagesResponseContent(data []PackageSummary) *ListPackagesResponseContent {
	this := ListPackagesResponseContent{}
	this.Data = data
	return &this
}

// NewListPackagesResponseContentWithDefaults instantiates a new ListPackagesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackagesResponseContentWithDefaults() *ListPackagesResponseContent {
	this := ListPackagesResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListPackagesResponseContent) GetData() []PackageSummary {
	if o == nil {
		var ret []PackageSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPackagesResponseContent) GetDataOk() ([]PackageSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListPackagesResponseContent) SetData(v []PackageSummary) {
	o.Data = v
}

func (o ListPackagesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPackagesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListPackagesResponseContent struct {
	value *ListPackagesResponseContent
	isSet bool
}

func (v NullableListPackagesResponseContent) Get() *ListPackagesResponseContent {
	return v.value
}

func (v *NullableListPackagesResponseContent) Set(val *ListPackagesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackagesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackagesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackagesResponseContent(val *ListPackagesResponseContent) *NullableListPackagesResponseContent {
	return &NullableListPackagesResponseContent{value: val, isSet: true}
}

func (v NullableListPackagesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPackagesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


