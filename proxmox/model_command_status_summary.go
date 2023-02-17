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

// checks if the CommandStatusSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandStatusSummary{}

// CommandStatusSummary struct for CommandStatusSummary
type CommandStatusSummary struct {
	Exited bool `json:"exited"`
	ErrData *string `json:"err-data,omitempty"`
	ErrTruncated *bool `json:"err-truncated,omitempty"`
	Exitcode *float32 `json:"exitcode,omitempty"`
	OutData *string `json:"out-data,omitempty"`
	OutTruncated *bool `json:"out-truncated,omitempty"`
	Signal *float32 `json:"signal,omitempty"`
}

// NewCommandStatusSummary instantiates a new CommandStatusSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandStatusSummary(exited bool) *CommandStatusSummary {
	this := CommandStatusSummary{}
	this.Exited = exited
	return &this
}

// NewCommandStatusSummaryWithDefaults instantiates a new CommandStatusSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandStatusSummaryWithDefaults() *CommandStatusSummary {
	this := CommandStatusSummary{}
	return &this
}

// GetExited returns the Exited field value
func (o *CommandStatusSummary) GetExited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exited
}

// GetExitedOk returns a tuple with the Exited field value
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetExitedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exited, true
}

// SetExited sets field value
func (o *CommandStatusSummary) SetExited(v bool) {
	o.Exited = v
}

// GetErrData returns the ErrData field value if set, zero value otherwise.
func (o *CommandStatusSummary) GetErrData() string {
	if o == nil || isNil(o.ErrData) {
		var ret string
		return ret
	}
	return *o.ErrData
}

// GetErrDataOk returns a tuple with the ErrData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetErrDataOk() (*string, bool) {
	if o == nil || isNil(o.ErrData) {
		return nil, false
	}
	return o.ErrData, true
}

// HasErrData returns a boolean if a field has been set.
func (o *CommandStatusSummary) HasErrData() bool {
	if o != nil && !isNil(o.ErrData) {
		return true
	}

	return false
}

// SetErrData gets a reference to the given string and assigns it to the ErrData field.
func (o *CommandStatusSummary) SetErrData(v string) {
	o.ErrData = &v
}

// GetErrTruncated returns the ErrTruncated field value if set, zero value otherwise.
func (o *CommandStatusSummary) GetErrTruncated() bool {
	if o == nil || isNil(o.ErrTruncated) {
		var ret bool
		return ret
	}
	return *o.ErrTruncated
}

// GetErrTruncatedOk returns a tuple with the ErrTruncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetErrTruncatedOk() (*bool, bool) {
	if o == nil || isNil(o.ErrTruncated) {
		return nil, false
	}
	return o.ErrTruncated, true
}

// HasErrTruncated returns a boolean if a field has been set.
func (o *CommandStatusSummary) HasErrTruncated() bool {
	if o != nil && !isNil(o.ErrTruncated) {
		return true
	}

	return false
}

// SetErrTruncated gets a reference to the given bool and assigns it to the ErrTruncated field.
func (o *CommandStatusSummary) SetErrTruncated(v bool) {
	o.ErrTruncated = &v
}

// GetExitcode returns the Exitcode field value if set, zero value otherwise.
func (o *CommandStatusSummary) GetExitcode() float32 {
	if o == nil || isNil(o.Exitcode) {
		var ret float32
		return ret
	}
	return *o.Exitcode
}

// GetExitcodeOk returns a tuple with the Exitcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetExitcodeOk() (*float32, bool) {
	if o == nil || isNil(o.Exitcode) {
		return nil, false
	}
	return o.Exitcode, true
}

// HasExitcode returns a boolean if a field has been set.
func (o *CommandStatusSummary) HasExitcode() bool {
	if o != nil && !isNil(o.Exitcode) {
		return true
	}

	return false
}

// SetExitcode gets a reference to the given float32 and assigns it to the Exitcode field.
func (o *CommandStatusSummary) SetExitcode(v float32) {
	o.Exitcode = &v
}

// GetOutData returns the OutData field value if set, zero value otherwise.
func (o *CommandStatusSummary) GetOutData() string {
	if o == nil || isNil(o.OutData) {
		var ret string
		return ret
	}
	return *o.OutData
}

// GetOutDataOk returns a tuple with the OutData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetOutDataOk() (*string, bool) {
	if o == nil || isNil(o.OutData) {
		return nil, false
	}
	return o.OutData, true
}

// HasOutData returns a boolean if a field has been set.
func (o *CommandStatusSummary) HasOutData() bool {
	if o != nil && !isNil(o.OutData) {
		return true
	}

	return false
}

// SetOutData gets a reference to the given string and assigns it to the OutData field.
func (o *CommandStatusSummary) SetOutData(v string) {
	o.OutData = &v
}

// GetOutTruncated returns the OutTruncated field value if set, zero value otherwise.
func (o *CommandStatusSummary) GetOutTruncated() bool {
	if o == nil || isNil(o.OutTruncated) {
		var ret bool
		return ret
	}
	return *o.OutTruncated
}

// GetOutTruncatedOk returns a tuple with the OutTruncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetOutTruncatedOk() (*bool, bool) {
	if o == nil || isNil(o.OutTruncated) {
		return nil, false
	}
	return o.OutTruncated, true
}

// HasOutTruncated returns a boolean if a field has been set.
func (o *CommandStatusSummary) HasOutTruncated() bool {
	if o != nil && !isNil(o.OutTruncated) {
		return true
	}

	return false
}

// SetOutTruncated gets a reference to the given bool and assigns it to the OutTruncated field.
func (o *CommandStatusSummary) SetOutTruncated(v bool) {
	o.OutTruncated = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *CommandStatusSummary) GetSignal() float32 {
	if o == nil || isNil(o.Signal) {
		var ret float32
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandStatusSummary) GetSignalOk() (*float32, bool) {
	if o == nil || isNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *CommandStatusSummary) HasSignal() bool {
	if o != nil && !isNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given float32 and assigns it to the Signal field.
func (o *CommandStatusSummary) SetSignal(v float32) {
	o.Signal = &v
}

func (o CommandStatusSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandStatusSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exited"] = o.Exited
	if !isNil(o.ErrData) {
		toSerialize["err-data"] = o.ErrData
	}
	if !isNil(o.ErrTruncated) {
		toSerialize["err-truncated"] = o.ErrTruncated
	}
	if !isNil(o.Exitcode) {
		toSerialize["exitcode"] = o.Exitcode
	}
	if !isNil(o.OutData) {
		toSerialize["out-data"] = o.OutData
	}
	if !isNil(o.OutTruncated) {
		toSerialize["out-truncated"] = o.OutTruncated
	}
	if !isNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	return toSerialize, nil
}

type NullableCommandStatusSummary struct {
	value *CommandStatusSummary
	isSet bool
}

func (v NullableCommandStatusSummary) Get() *CommandStatusSummary {
	return v.value
}

func (v *NullableCommandStatusSummary) Set(val *CommandStatusSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandStatusSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandStatusSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandStatusSummary(val *CommandStatusSummary) *NullableCommandStatusSummary {
	return &NullableCommandStatusSummary{value: val, isSet: true}
}

func (v NullableCommandStatusSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandStatusSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

