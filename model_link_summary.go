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

// LinkSummary struct for LinkSummary
type LinkSummary struct {
	Linknumber *string `json:"linknumber,omitempty"`
}

// NewLinkSummary instantiates a new LinkSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSummary() *LinkSummary {
	this := LinkSummary{}
	return &this
}

// NewLinkSummaryWithDefaults instantiates a new LinkSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSummaryWithDefaults() *LinkSummary {
	this := LinkSummary{}
	return &this
}

// GetLinknumber returns the Linknumber field value if set, zero value otherwise.
func (o *LinkSummary) GetLinknumber() string {
	if o == nil || isNil(o.Linknumber) {
		var ret string
		return ret
	}
	return *o.Linknumber
}

// GetLinknumberOk returns a tuple with the Linknumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkSummary) GetLinknumberOk() (*string, bool) {
	if o == nil || isNil(o.Linknumber) {
    return nil, false
	}
	return o.Linknumber, true
}

// HasLinknumber returns a boolean if a field has been set.
func (o *LinkSummary) HasLinknumber() bool {
	if o != nil && !isNil(o.Linknumber) {
		return true
	}

	return false
}

// SetLinknumber gets a reference to the given string and assigns it to the Linknumber field.
func (o *LinkSummary) SetLinknumber(v string) {
	o.Linknumber = &v
}

func (o LinkSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Linknumber) {
		toSerialize["linknumber"] = o.Linknumber
	}
	return json.Marshal(toSerialize)
}

type NullableLinkSummary struct {
	value *LinkSummary
	isSet bool
}

func (v NullableLinkSummary) Get() *LinkSummary {
	return v.value
}

func (v *NullableLinkSummary) Set(val *LinkSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSummary(val *LinkSummary) *NullableLinkSummary {
	return &NullableLinkSummary{value: val, isSet: true}
}

func (v NullableLinkSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


