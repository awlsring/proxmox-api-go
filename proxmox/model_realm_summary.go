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

// checks if the RealmSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmSummary{}

// RealmSummary struct for RealmSummary
type RealmSummary struct {
	Realm string `json:"realm"`
	Type string `json:"type"`
	Comment *string `json:"comment,omitempty"`
	Tfa *TFAType `json:"tfa,omitempty"`
}

// NewRealmSummary instantiates a new RealmSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmSummary(realm string, type_ string) *RealmSummary {
	this := RealmSummary{}
	this.Realm = realm
	this.Type = type_
	return &this
}

// NewRealmSummaryWithDefaults instantiates a new RealmSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmSummaryWithDefaults() *RealmSummary {
	this := RealmSummary{}
	return &this
}

// GetRealm returns the Realm field value
func (o *RealmSummary) GetRealm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Realm
}

// GetRealmOk returns a tuple with the Realm field value
// and a boolean to check if the value has been set.
func (o *RealmSummary) GetRealmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Realm, true
}

// SetRealm sets field value
func (o *RealmSummary) SetRealm(v string) {
	o.Realm = v
}

// GetType returns the Type field value
func (o *RealmSummary) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealmSummary) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealmSummary) SetType(v string) {
	o.Type = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RealmSummary) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmSummary) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RealmSummary) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RealmSummary) SetComment(v string) {
	o.Comment = &v
}

// GetTfa returns the Tfa field value if set, zero value otherwise.
func (o *RealmSummary) GetTfa() TFAType {
	if o == nil || IsNil(o.Tfa) {
		var ret TFAType
		return ret
	}
	return *o.Tfa
}

// GetTfaOk returns a tuple with the Tfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmSummary) GetTfaOk() (*TFAType, bool) {
	if o == nil || IsNil(o.Tfa) {
		return nil, false
	}
	return o.Tfa, true
}

// HasTfa returns a boolean if a field has been set.
func (o *RealmSummary) HasTfa() bool {
	if o != nil && !IsNil(o.Tfa) {
		return true
	}

	return false
}

// SetTfa gets a reference to the given TFAType and assigns it to the Tfa field.
func (o *RealmSummary) SetTfa(v TFAType) {
	o.Tfa = &v
}

func (o RealmSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["realm"] = o.Realm
	toSerialize["type"] = o.Type
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Tfa) {
		toSerialize["tfa"] = o.Tfa
	}
	return toSerialize, nil
}

type NullableRealmSummary struct {
	value *RealmSummary
	isSet bool
}

func (v NullableRealmSummary) Get() *RealmSummary {
	return v.value
}

func (v *NullableRealmSummary) Set(val *RealmSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmSummary(val *RealmSummary) *NullableRealmSummary {
	return &NullableRealmSummary{value: val, isSet: true}
}

func (v NullableRealmSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


