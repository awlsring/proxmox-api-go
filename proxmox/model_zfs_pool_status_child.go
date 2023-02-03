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

// checks if the ZFSPoolStatusChild type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZFSPoolStatusChild{}

// ZFSPoolStatusChild struct for ZFSPoolStatusChild
type ZFSPoolStatusChild struct {
	Cksum float32 `json:"cksum"`
	Read float32 `json:"read"`
	Write float32 `json:"write"`
	Name string `json:"name"`
	State string `json:"state"`
	Msg *string `json:"msg,omitempty"`
	Children []ZFSPoolStatusChild `json:"children,omitempty"`
}

// NewZFSPoolStatusChild instantiates a new ZFSPoolStatusChild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZFSPoolStatusChild(cksum float32, read float32, write float32, name string, state string) *ZFSPoolStatusChild {
	this := ZFSPoolStatusChild{}
	this.Cksum = cksum
	this.Read = read
	this.Write = write
	this.Name = name
	this.State = state
	return &this
}

// NewZFSPoolStatusChildWithDefaults instantiates a new ZFSPoolStatusChild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZFSPoolStatusChildWithDefaults() *ZFSPoolStatusChild {
	this := ZFSPoolStatusChild{}
	return &this
}

// GetCksum returns the Cksum field value
func (o *ZFSPoolStatusChild) GetCksum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Cksum
}

// GetCksumOk returns a tuple with the Cksum field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetCksumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cksum, true
}

// SetCksum sets field value
func (o *ZFSPoolStatusChild) SetCksum(v float32) {
	o.Cksum = v
}

// GetRead returns the Read field value
func (o *ZFSPoolStatusChild) GetRead() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Read
}

// GetReadOk returns a tuple with the Read field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetReadOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Read, true
}

// SetRead sets field value
func (o *ZFSPoolStatusChild) SetRead(v float32) {
	o.Read = v
}

// GetWrite returns the Write field value
func (o *ZFSPoolStatusChild) GetWrite() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Write
}

// GetWriteOk returns a tuple with the Write field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetWriteOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Write, true
}

// SetWrite sets field value
func (o *ZFSPoolStatusChild) SetWrite(v float32) {
	o.Write = v
}

// GetName returns the Name field value
func (o *ZFSPoolStatusChild) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZFSPoolStatusChild) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *ZFSPoolStatusChild) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ZFSPoolStatusChild) SetState(v string) {
	o.State = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ZFSPoolStatusChild) GetMsg() string {
	if o == nil || isNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetMsgOk() (*string, bool) {
	if o == nil || isNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ZFSPoolStatusChild) HasMsg() bool {
	if o != nil && !isNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ZFSPoolStatusChild) SetMsg(v string) {
	o.Msg = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ZFSPoolStatusChild) GetChildren() []ZFSPoolStatusChild {
	if o == nil || isNil(o.Children) {
		var ret []ZFSPoolStatusChild
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusChild) GetChildrenOk() ([]ZFSPoolStatusChild, bool) {
	if o == nil || isNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ZFSPoolStatusChild) HasChildren() bool {
	if o != nil && !isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ZFSPoolStatusChild and assigns it to the Children field.
func (o *ZFSPoolStatusChild) SetChildren(v []ZFSPoolStatusChild) {
	o.Children = v
}

func (o ZFSPoolStatusChild) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZFSPoolStatusChild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cksum"] = o.Cksum
	toSerialize["read"] = o.Read
	toSerialize["write"] = o.Write
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	if !isNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !isNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableZFSPoolStatusChild struct {
	value *ZFSPoolStatusChild
	isSet bool
}

func (v NullableZFSPoolStatusChild) Get() *ZFSPoolStatusChild {
	return v.value
}

func (v *NullableZFSPoolStatusChild) Set(val *ZFSPoolStatusChild) {
	v.value = val
	v.isSet = true
}

func (v NullableZFSPoolStatusChild) IsSet() bool {
	return v.isSet
}

func (v *NullableZFSPoolStatusChild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZFSPoolStatusChild(val *ZFSPoolStatusChild) *NullableZFSPoolStatusChild {
	return &NullableZFSPoolStatusChild{value: val, isSet: true}
}

func (v NullableZFSPoolStatusChild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZFSPoolStatusChild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

