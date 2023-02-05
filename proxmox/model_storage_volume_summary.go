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

// checks if the StorageVolumeSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageVolumeSummary{}

// StorageVolumeSummary struct for StorageVolumeSummary
type StorageVolumeSummary struct {
	Format string `json:"format"`
	// The size of the storage content in bytes.
	Size float32 `json:"size"`
	Volid string `json:"volid"`
	// The creation time of the storage content in seconds since the epoch.
	Ctime *float32 `json:"ctime,omitempty"`
	Encrypted *string `json:"encrypted,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Parent *string `json:"parent,omitempty"`
	// An integer used to represent a boolean. 0 is false, 1 is true.
	Protected *float32 `json:"protected,omitempty"`
	// The amount of storage content used in bytes.
	Used *float32 `json:"used,omitempty"`
	// The id of the virtual machine as a string
	Vmid *string `json:"vmid,omitempty"`
}

// NewStorageVolumeSummary instantiates a new StorageVolumeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVolumeSummary(format string, size float32, volid string) *StorageVolumeSummary {
	this := StorageVolumeSummary{}
	this.Format = format
	this.Size = size
	this.Volid = volid
	return &this
}

// NewStorageVolumeSummaryWithDefaults instantiates a new StorageVolumeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVolumeSummaryWithDefaults() *StorageVolumeSummary {
	this := StorageVolumeSummary{}
	return &this
}

// GetFormat returns the Format field value
func (o *StorageVolumeSummary) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *StorageVolumeSummary) SetFormat(v string) {
	o.Format = v
}

// GetSize returns the Size field value
func (o *StorageVolumeSummary) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *StorageVolumeSummary) SetSize(v float32) {
	o.Size = v
}

// GetVolid returns the Volid field value
func (o *StorageVolumeSummary) GetVolid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Volid
}

// GetVolidOk returns a tuple with the Volid field value
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetVolidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volid, true
}

// SetVolid sets field value
func (o *StorageVolumeSummary) SetVolid(v string) {
	o.Volid = v
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetCtime() float32 {
	if o == nil || isNil(o.Ctime) {
		var ret float32
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetCtimeOk() (*float32, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given float32 and assigns it to the Ctime field.
func (o *StorageVolumeSummary) SetCtime(v float32) {
	o.Ctime = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetEncrypted() string {
	if o == nil || isNil(o.Encrypted) {
		var ret string
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetEncryptedOk() (*string, bool) {
	if o == nil || isNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasEncrypted() bool {
	if o != nil && !isNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given string and assigns it to the Encrypted field.
func (o *StorageVolumeSummary) SetEncrypted(v string) {
	o.Encrypted = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *StorageVolumeSummary) SetNotes(v string) {
	o.Notes = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetParent() string {
	if o == nil || isNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetParentOk() (*string, bool) {
	if o == nil || isNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasParent() bool {
	if o != nil && !isNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *StorageVolumeSummary) SetParent(v string) {
	o.Parent = &v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetProtected() float32 {
	if o == nil || isNil(o.Protected) {
		var ret float32
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetProtectedOk() (*float32, bool) {
	if o == nil || isNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasProtected() bool {
	if o != nil && !isNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given float32 and assigns it to the Protected field.
func (o *StorageVolumeSummary) SetProtected(v float32) {
	o.Protected = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetUsed() float32 {
	if o == nil || isNil(o.Used) {
		var ret float32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetUsedOk() (*float32, bool) {
	if o == nil || isNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given float32 and assigns it to the Used field.
func (o *StorageVolumeSummary) SetUsed(v float32) {
	o.Used = &v
}

// GetVmid returns the Vmid field value if set, zero value otherwise.
func (o *StorageVolumeSummary) GetVmid() string {
	if o == nil || isNil(o.Vmid) {
		var ret string
		return ret
	}
	return *o.Vmid
}

// GetVmidOk returns a tuple with the Vmid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVolumeSummary) GetVmidOk() (*string, bool) {
	if o == nil || isNil(o.Vmid) {
		return nil, false
	}
	return o.Vmid, true
}

// HasVmid returns a boolean if a field has been set.
func (o *StorageVolumeSummary) HasVmid() bool {
	if o != nil && !isNil(o.Vmid) {
		return true
	}

	return false
}

// SetVmid gets a reference to the given string and assigns it to the Vmid field.
func (o *StorageVolumeSummary) SetVmid(v string) {
	o.Vmid = &v
}

func (o StorageVolumeSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageVolumeSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	toSerialize["size"] = o.Size
	toSerialize["volid"] = o.Volid
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !isNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !isNil(o.Vmid) {
		toSerialize["vmid"] = o.Vmid
	}
	return toSerialize, nil
}

type NullableStorageVolumeSummary struct {
	value *StorageVolumeSummary
	isSet bool
}

func (v NullableStorageVolumeSummary) Get() *StorageVolumeSummary {
	return v.value
}

func (v *NullableStorageVolumeSummary) Set(val *StorageVolumeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVolumeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVolumeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVolumeSummary(val *StorageVolumeSummary) *NullableStorageVolumeSummary {
	return &NullableStorageVolumeSummary{value: val, isSet: true}
}

func (v NullableStorageVolumeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVolumeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


