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

// checks if the GetMemoryBlocksResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMemoryBlocksResult{}

// GetMemoryBlocksResult struct for GetMemoryBlocksResult
type GetMemoryBlocksResult struct {
	Result []MemoryBlockSummary `json:"result,omitempty"`
}

// NewGetMemoryBlocksResult instantiates a new GetMemoryBlocksResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMemoryBlocksResult() *GetMemoryBlocksResult {
	this := GetMemoryBlocksResult{}
	return &this
}

// NewGetMemoryBlocksResultWithDefaults instantiates a new GetMemoryBlocksResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMemoryBlocksResultWithDefaults() *GetMemoryBlocksResult {
	this := GetMemoryBlocksResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetMemoryBlocksResult) GetResult() []MemoryBlockSummary {
	if o == nil || isNil(o.Result) {
		var ret []MemoryBlockSummary
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMemoryBlocksResult) GetResultOk() ([]MemoryBlockSummary, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetMemoryBlocksResult) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []MemoryBlockSummary and assigns it to the Result field.
func (o *GetMemoryBlocksResult) SetResult(v []MemoryBlockSummary) {
	o.Result = v
}

func (o GetMemoryBlocksResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMemoryBlocksResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetMemoryBlocksResult struct {
	value *GetMemoryBlocksResult
	isSet bool
}

func (v NullableGetMemoryBlocksResult) Get() *GetMemoryBlocksResult {
	return v.value
}

func (v *NullableGetMemoryBlocksResult) Set(val *GetMemoryBlocksResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMemoryBlocksResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMemoryBlocksResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMemoryBlocksResult(val *GetMemoryBlocksResult) *NullableGetMemoryBlocksResult {
	return &NullableGetMemoryBlocksResult{value: val, isSet: true}
}

func (v NullableGetMemoryBlocksResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMemoryBlocksResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


