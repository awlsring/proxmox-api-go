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

// checks if the ZFSPoolStatusSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZFSPoolStatusSummary{}

// ZFSPoolStatusSummary struct for ZFSPoolStatusSummary
type ZFSPoolStatusSummary struct {
	// The pool configuration including vdevs for each section. May be nested.
	Children []ZFSPoolStatusChild `json:"children"`
	// Errors on the pool
	Errors string `json:"errors"`
	// The pool name
	Name string `json:"name"`
	// The state of the pool
	State string `json:"state"`
	// Recommended action to address the pool state
	Action *string `json:"action,omitempty"`
	// Information on the last or current scrub.
	Scan *string `json:"scan,omitempty"`
	// Information on the state of the pool
	Status *string `json:"status,omitempty"`
}

// NewZFSPoolStatusSummary instantiates a new ZFSPoolStatusSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZFSPoolStatusSummary(children []ZFSPoolStatusChild, errors string, name string, state string) *ZFSPoolStatusSummary {
	this := ZFSPoolStatusSummary{}
	this.Children = children
	this.Errors = errors
	this.Name = name
	this.State = state
	return &this
}

// NewZFSPoolStatusSummaryWithDefaults instantiates a new ZFSPoolStatusSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZFSPoolStatusSummaryWithDefaults() *ZFSPoolStatusSummary {
	this := ZFSPoolStatusSummary{}
	return &this
}

// GetChildren returns the Children field value
func (o *ZFSPoolStatusSummary) GetChildren() []ZFSPoolStatusChild {
	if o == nil {
		var ret []ZFSPoolStatusChild
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetChildrenOk() ([]ZFSPoolStatusChild, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *ZFSPoolStatusSummary) SetChildren(v []ZFSPoolStatusChild) {
	o.Children = v
}

// GetErrors returns the Errors field value
func (o *ZFSPoolStatusSummary) GetErrors() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetErrorsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *ZFSPoolStatusSummary) SetErrors(v string) {
	o.Errors = v
}

// GetName returns the Name field value
func (o *ZFSPoolStatusSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZFSPoolStatusSummary) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *ZFSPoolStatusSummary) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ZFSPoolStatusSummary) SetState(v string) {
	o.State = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ZFSPoolStatusSummary) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ZFSPoolStatusSummary) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ZFSPoolStatusSummary) SetAction(v string) {
	o.Action = &v
}

// GetScan returns the Scan field value if set, zero value otherwise.
func (o *ZFSPoolStatusSummary) GetScan() string {
	if o == nil || IsNil(o.Scan) {
		var ret string
		return ret
	}
	return *o.Scan
}

// GetScanOk returns a tuple with the Scan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetScanOk() (*string, bool) {
	if o == nil || IsNil(o.Scan) {
		return nil, false
	}
	return o.Scan, true
}

// HasScan returns a boolean if a field has been set.
func (o *ZFSPoolStatusSummary) HasScan() bool {
	if o != nil && !IsNil(o.Scan) {
		return true
	}

	return false
}

// SetScan gets a reference to the given string and assigns it to the Scan field.
func (o *ZFSPoolStatusSummary) SetScan(v string) {
	o.Scan = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ZFSPoolStatusSummary) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZFSPoolStatusSummary) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ZFSPoolStatusSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ZFSPoolStatusSummary) SetStatus(v string) {
	o.Status = &v
}

func (o ZFSPoolStatusSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZFSPoolStatusSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["children"] = o.Children
	toSerialize["errors"] = o.Errors
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Scan) {
		toSerialize["scan"] = o.Scan
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableZFSPoolStatusSummary struct {
	value *ZFSPoolStatusSummary
	isSet bool
}

func (v NullableZFSPoolStatusSummary) Get() *ZFSPoolStatusSummary {
	return v.value
}

func (v *NullableZFSPoolStatusSummary) Set(val *ZFSPoolStatusSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableZFSPoolStatusSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableZFSPoolStatusSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZFSPoolStatusSummary(val *ZFSPoolStatusSummary) *NullableZFSPoolStatusSummary {
	return &NullableZFSPoolStatusSummary{value: val, isSet: true}
}

func (v NullableZFSPoolStatusSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZFSPoolStatusSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


