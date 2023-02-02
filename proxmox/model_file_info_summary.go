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

// checks if the FileInfoSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileInfoSummary{}

// FileInfoSummary struct for FileInfoSummary
type FileInfoSummary struct {
	Path string `json:"path"`
	Index float32 `json:"index"`
	Message string `json:"message"`
	Kind string `json:"kind"`
}

// NewFileInfoSummary instantiates a new FileInfoSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInfoSummary(path string, index float32, message string, kind string) *FileInfoSummary {
	this := FileInfoSummary{}
	this.Path = path
	this.Index = index
	this.Message = message
	this.Kind = kind
	return &this
}

// NewFileInfoSummaryWithDefaults instantiates a new FileInfoSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInfoSummaryWithDefaults() *FileInfoSummary {
	this := FileInfoSummary{}
	return &this
}

// GetPath returns the Path field value
func (o *FileInfoSummary) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FileInfoSummary) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FileInfoSummary) SetPath(v string) {
	o.Path = v
}

// GetIndex returns the Index field value
func (o *FileInfoSummary) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *FileInfoSummary) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *FileInfoSummary) SetIndex(v float32) {
	o.Index = v
}

// GetMessage returns the Message field value
func (o *FileInfoSummary) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FileInfoSummary) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FileInfoSummary) SetMessage(v string) {
	o.Message = v
}

// GetKind returns the Kind field value
func (o *FileInfoSummary) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *FileInfoSummary) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *FileInfoSummary) SetKind(v string) {
	o.Kind = v
}

func (o FileInfoSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileInfoSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["index"] = o.Index
	toSerialize["message"] = o.Message
	toSerialize["kind"] = o.Kind
	return toSerialize, nil
}

type NullableFileInfoSummary struct {
	value *FileInfoSummary
	isSet bool
}

func (v NullableFileInfoSummary) Get() *FileInfoSummary {
	return v.value
}

func (v *NullableFileInfoSummary) Set(val *FileInfoSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInfoSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInfoSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInfoSummary(val *FileInfoSummary) *NullableFileInfoSummary {
	return &NullableFileInfoSummary{value: val, isSet: true}
}

func (v NullableFileInfoSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInfoSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


