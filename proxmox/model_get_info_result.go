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

// checks if the GetInfoResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInfoResult{}

// GetInfoResult struct for GetInfoResult
type GetInfoResult struct {
	Result *AgentInfoSummary `json:"result,omitempty"`
}

// NewGetInfoResult instantiates a new GetInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInfoResult() *GetInfoResult {
	this := GetInfoResult{}
	return &this
}

// NewGetInfoResultWithDefaults instantiates a new GetInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInfoResultWithDefaults() *GetInfoResult {
	this := GetInfoResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetInfoResult) GetResult() AgentInfoSummary {
	if o == nil || isNil(o.Result) {
		var ret AgentInfoSummary
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInfoResult) GetResultOk() (*AgentInfoSummary, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetInfoResult) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AgentInfoSummary and assigns it to the Result field.
func (o *GetInfoResult) SetResult(v AgentInfoSummary) {
	o.Result = &v
}

func (o GetInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInfoResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetInfoResult struct {
	value *GetInfoResult
	isSet bool
}

func (v NullableGetInfoResult) Get() *GetInfoResult {
	return v.value
}

func (v *NullableGetInfoResult) Set(val *GetInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInfoResult(val *GetInfoResult) *NullableGetInfoResult {
	return &NullableGetInfoResult{value: val, isSet: true}
}

func (v NullableGetInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


