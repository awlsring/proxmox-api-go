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

// checks if the GetUserTokenResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserTokenResponseContent{}

// GetUserTokenResponseContent struct for GetUserTokenResponseContent
type GetUserTokenResponseContent struct {
	Data UserConfigurationTokenSummary `json:"data"`
}

// NewGetUserTokenResponseContent instantiates a new GetUserTokenResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserTokenResponseContent(data UserConfigurationTokenSummary) *GetUserTokenResponseContent {
	this := GetUserTokenResponseContent{}
	this.Data = data
	return &this
}

// NewGetUserTokenResponseContentWithDefaults instantiates a new GetUserTokenResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserTokenResponseContentWithDefaults() *GetUserTokenResponseContent {
	this := GetUserTokenResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetUserTokenResponseContent) GetData() UserConfigurationTokenSummary {
	if o == nil {
		var ret UserConfigurationTokenSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetUserTokenResponseContent) GetDataOk() (*UserConfigurationTokenSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetUserTokenResponseContent) SetData(v UserConfigurationTokenSummary) {
	o.Data = v
}

func (o GetUserTokenResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserTokenResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetUserTokenResponseContent struct {
	value *GetUserTokenResponseContent
	isSet bool
}

func (v NullableGetUserTokenResponseContent) Get() *GetUserTokenResponseContent {
	return v.value
}

func (v *NullableGetUserTokenResponseContent) Set(val *GetUserTokenResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserTokenResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserTokenResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserTokenResponseContent(val *GetUserTokenResponseContent) *NullableGetUserTokenResponseContent {
	return &NullableGetUserTokenResponseContent{value: val, isSet: true}
}

func (v NullableGetUserTokenResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserTokenResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


