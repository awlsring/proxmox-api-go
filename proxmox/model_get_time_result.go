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

// checks if the GetTimeResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeResult{}

// GetTimeResult struct for GetTimeResult
type GetTimeResult struct {
	Result *float32 `json:"result,omitempty"`
}

// NewGetTimeResult instantiates a new GetTimeResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeResult() *GetTimeResult {
	this := GetTimeResult{}
	return &this
}

// NewGetTimeResultWithDefaults instantiates a new GetTimeResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeResultWithDefaults() *GetTimeResult {
	this := GetTimeResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetTimeResult) GetResult() float32 {
	if o == nil || IsNil(o.Result) {
		var ret float32
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeResult) GetResultOk() (*float32, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetTimeResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given float32 and assigns it to the Result field.
func (o *GetTimeResult) SetResult(v float32) {
	o.Result = &v
}

func (o GetTimeResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetTimeResult struct {
	value *GetTimeResult
	isSet bool
}

func (v NullableGetTimeResult) Get() *GetTimeResult {
	return v.value
}

func (v *NullableGetTimeResult) Set(val *GetTimeResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeResult(val *GetTimeResult) *NullableGetTimeResult {
	return &NullableGetTimeResult{value: val, isSet: true}
}

func (v NullableGetTimeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


