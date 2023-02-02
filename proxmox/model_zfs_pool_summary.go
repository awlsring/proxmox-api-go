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

// checks if the ZFSPoolSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZFSPoolSummary{}

// ZFSPoolSummary struct for ZFSPoolSummary
type ZFSPoolSummary struct {
	// The allocated space on the pool in bytes
	Alloc float32 `json:"alloc"`
	Dedup float32 `json:"dedup"`
	Frag float32 `json:"frag"`
	// The health of the pool
	Health string `json:"health"`
	// The free space on the pool in bytes
	Free float32 `json:"free"`
	// The total size of the pool in bytes
	Size float32 `json:"size"`
	// Name of the pool
	Name string `json:"name"`
}

// NewZFSPoolSummary instantiates a new ZFSPoolSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZFSPoolSummary(alloc float32, dedup float32, frag float32, health string, free float32, size float32, name string) *ZFSPoolSummary {
	this := ZFSPoolSummary{}
	this.Alloc = alloc
	this.Dedup = dedup
	this.Frag = frag
	this.Health = health
	this.Free = free
	this.Size = size
	this.Name = name
	return &this
}

// NewZFSPoolSummaryWithDefaults instantiates a new ZFSPoolSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZFSPoolSummaryWithDefaults() *ZFSPoolSummary {
	this := ZFSPoolSummary{}
	return &this
}

// GetAlloc returns the Alloc field value
func (o *ZFSPoolSummary) GetAlloc() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Alloc
}

// GetAllocOk returns a tuple with the Alloc field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetAllocOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alloc, true
}

// SetAlloc sets field value
func (o *ZFSPoolSummary) SetAlloc(v float32) {
	o.Alloc = v
}

// GetDedup returns the Dedup field value
func (o *ZFSPoolSummary) GetDedup() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Dedup
}

// GetDedupOk returns a tuple with the Dedup field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetDedupOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dedup, true
}

// SetDedup sets field value
func (o *ZFSPoolSummary) SetDedup(v float32) {
	o.Dedup = v
}

// GetFrag returns the Frag field value
func (o *ZFSPoolSummary) GetFrag() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Frag
}

// GetFragOk returns a tuple with the Frag field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetFragOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frag, true
}

// SetFrag sets field value
func (o *ZFSPoolSummary) SetFrag(v float32) {
	o.Frag = v
}

// GetHealth returns the Health field value
func (o *ZFSPoolSummary) GetHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *ZFSPoolSummary) SetHealth(v string) {
	o.Health = v
}

// GetFree returns the Free field value
func (o *ZFSPoolSummary) GetFree() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetFreeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *ZFSPoolSummary) SetFree(v float32) {
	o.Free = v
}

// GetSize returns the Size field value
func (o *ZFSPoolSummary) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ZFSPoolSummary) SetSize(v float32) {
	o.Size = v
}

// GetName returns the Name field value
func (o *ZFSPoolSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZFSPoolSummary) SetName(v string) {
	o.Name = v
}

func (o ZFSPoolSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZFSPoolSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alloc"] = o.Alloc
	toSerialize["dedup"] = o.Dedup
	toSerialize["frag"] = o.Frag
	toSerialize["health"] = o.Health
	toSerialize["free"] = o.Free
	toSerialize["size"] = o.Size
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableZFSPoolSummary struct {
	value *ZFSPoolSummary
	isSet bool
}

func (v NullableZFSPoolSummary) Get() *ZFSPoolSummary {
	return v.value
}

func (v *NullableZFSPoolSummary) Set(val *ZFSPoolSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableZFSPoolSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableZFSPoolSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZFSPoolSummary(val *ZFSPoolSummary) *NullableZFSPoolSummary {
	return &NullableZFSPoolSummary{value: val, isSet: true}
}

func (v NullableZFSPoolSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZFSPoolSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


