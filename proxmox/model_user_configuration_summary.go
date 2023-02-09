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

// checks if the UserConfigurationSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserConfigurationSummary{}

// UserConfigurationSummary struct for UserConfigurationSummary
type UserConfigurationSummary struct {
	Email *string `json:"email,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	// An integer used to represent a boolean. 0 is false, 1 is true.
	Enable *float32 `json:"enable,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Keys *string `json:"keys,omitempty"`
	Expire *float32 `json:"expire,omitempty"`
	Tokens *map[string]UserConfigurationTokenSummary `json:"tokens,omitempty"`
}

// NewUserConfigurationSummary instantiates a new UserConfigurationSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConfigurationSummary() *UserConfigurationSummary {
	this := UserConfigurationSummary{}
	return &this
}

// NewUserConfigurationSummaryWithDefaults instantiates a new UserConfigurationSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConfigurationSummaryWithDefaults() *UserConfigurationSummary {
	this := UserConfigurationSummary{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserConfigurationSummary) SetEmail(v string) {
	o.Email = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetFirstname() string {
	if o == nil || isNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetFirstnameOk() (*string, bool) {
	if o == nil || isNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasFirstname() bool {
	if o != nil && !isNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *UserConfigurationSummary) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *UserConfigurationSummary) SetLastname(v string) {
	o.Lastname = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetEnable() float32 {
	if o == nil || isNil(o.Enable) {
		var ret float32
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetEnableOk() (*float32, bool) {
	if o == nil || isNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasEnable() bool {
	if o != nil && !isNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given float32 and assigns it to the Enable field.
func (o *UserConfigurationSummary) SetEnable(v float32) {
	o.Enable = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UserConfigurationSummary) SetComment(v string) {
	o.Comment = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetGroups() []string {
	if o == nil || isNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *UserConfigurationSummary) SetGroups(v []string) {
	o.Groups = v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetKeys() string {
	if o == nil || isNil(o.Keys) {
		var ret string
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetKeysOk() (*string, bool) {
	if o == nil || isNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasKeys() bool {
	if o != nil && !isNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given string and assigns it to the Keys field.
func (o *UserConfigurationSummary) SetKeys(v string) {
	o.Keys = &v
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetExpire() float32 {
	if o == nil || isNil(o.Expire) {
		var ret float32
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetExpireOk() (*float32, bool) {
	if o == nil || isNil(o.Expire) {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasExpire() bool {
	if o != nil && !isNil(o.Expire) {
		return true
	}

	return false
}

// SetExpire gets a reference to the given float32 and assigns it to the Expire field.
func (o *UserConfigurationSummary) SetExpire(v float32) {
	o.Expire = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *UserConfigurationSummary) GetTokens() map[string]UserConfigurationTokenSummary {
	if o == nil || isNil(o.Tokens) {
		var ret map[string]UserConfigurationTokenSummary
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfigurationSummary) GetTokensOk() (*map[string]UserConfigurationTokenSummary, bool) {
	if o == nil || isNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *UserConfigurationSummary) HasTokens() bool {
	if o != nil && !isNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given map[string]UserConfigurationTokenSummary and assigns it to the Tokens field.
func (o *UserConfigurationSummary) SetTokens(v map[string]UserConfigurationTokenSummary) {
	o.Tokens = &v
}

func (o UserConfigurationSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserConfigurationSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !isNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !isNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	if !isNil(o.Expire) {
		toSerialize["expire"] = o.Expire
	}
	if !isNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	return toSerialize, nil
}

type NullableUserConfigurationSummary struct {
	value *UserConfigurationSummary
	isSet bool
}

func (v NullableUserConfigurationSummary) Get() *UserConfigurationSummary {
	return v.value
}

func (v *NullableUserConfigurationSummary) Set(val *UserConfigurationSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConfigurationSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConfigurationSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConfigurationSummary(val *UserConfigurationSummary) *NullableUserConfigurationSummary {
	return &NullableUserConfigurationSummary{value: val, isSet: true}
}

func (v NullableUserConfigurationSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConfigurationSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

