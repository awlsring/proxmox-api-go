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

// DiskInformationSummary struct for DiskInformationSummary
type DiskInformationSummary struct {
	Target *float32 `json:"target,omitempty"`
	Bus *float32 `json:"bus,omitempty"`
	PciController *PciControllerSummary `json:"pci-controller,omitempty"`
	Unit *float32 `json:"unit,omitempty"`
	Dev *string `json:"dev,omitempty"`
	BusType *string `json:"bus-type,omitempty"`
	Serial *string `json:"serial,omitempty"`
}

// NewDiskInformationSummary instantiates a new DiskInformationSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskInformationSummary() *DiskInformationSummary {
	this := DiskInformationSummary{}
	return &this
}

// NewDiskInformationSummaryWithDefaults instantiates a new DiskInformationSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskInformationSummaryWithDefaults() *DiskInformationSummary {
	this := DiskInformationSummary{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetTarget() float32 {
	if o == nil || isNil(o.Target) {
		var ret float32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetTargetOk() (*float32, bool) {
	if o == nil || isNil(o.Target) {
    return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given float32 and assigns it to the Target field.
func (o *DiskInformationSummary) SetTarget(v float32) {
	o.Target = &v
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetBus() float32 {
	if o == nil || isNil(o.Bus) {
		var ret float32
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetBusOk() (*float32, bool) {
	if o == nil || isNil(o.Bus) {
    return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasBus() bool {
	if o != nil && !isNil(o.Bus) {
		return true
	}

	return false
}

// SetBus gets a reference to the given float32 and assigns it to the Bus field.
func (o *DiskInformationSummary) SetBus(v float32) {
	o.Bus = &v
}

// GetPciController returns the PciController field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetPciController() PciControllerSummary {
	if o == nil || isNil(o.PciController) {
		var ret PciControllerSummary
		return ret
	}
	return *o.PciController
}

// GetPciControllerOk returns a tuple with the PciController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetPciControllerOk() (*PciControllerSummary, bool) {
	if o == nil || isNil(o.PciController) {
    return nil, false
	}
	return o.PciController, true
}

// HasPciController returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasPciController() bool {
	if o != nil && !isNil(o.PciController) {
		return true
	}

	return false
}

// SetPciController gets a reference to the given PciControllerSummary and assigns it to the PciController field.
func (o *DiskInformationSummary) SetPciController(v PciControllerSummary) {
	o.PciController = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetUnit() float32 {
	if o == nil || isNil(o.Unit) {
		var ret float32
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetUnitOk() (*float32, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given float32 and assigns it to the Unit field.
func (o *DiskInformationSummary) SetUnit(v float32) {
	o.Unit = &v
}

// GetDev returns the Dev field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetDev() string {
	if o == nil || isNil(o.Dev) {
		var ret string
		return ret
	}
	return *o.Dev
}

// GetDevOk returns a tuple with the Dev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetDevOk() (*string, bool) {
	if o == nil || isNil(o.Dev) {
    return nil, false
	}
	return o.Dev, true
}

// HasDev returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasDev() bool {
	if o != nil && !isNil(o.Dev) {
		return true
	}

	return false
}

// SetDev gets a reference to the given string and assigns it to the Dev field.
func (o *DiskInformationSummary) SetDev(v string) {
	o.Dev = &v
}

// GetBusType returns the BusType field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetBusType() string {
	if o == nil || isNil(o.BusType) {
		var ret string
		return ret
	}
	return *o.BusType
}

// GetBusTypeOk returns a tuple with the BusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetBusTypeOk() (*string, bool) {
	if o == nil || isNil(o.BusType) {
    return nil, false
	}
	return o.BusType, true
}

// HasBusType returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasBusType() bool {
	if o != nil && !isNil(o.BusType) {
		return true
	}

	return false
}

// SetBusType gets a reference to the given string and assigns it to the BusType field.
func (o *DiskInformationSummary) SetBusType(v string) {
	o.BusType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *DiskInformationSummary) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskInformationSummary) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *DiskInformationSummary) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *DiskInformationSummary) SetSerial(v string) {
	o.Serial = &v
}

func (o DiskInformationSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.Bus) {
		toSerialize["bus"] = o.Bus
	}
	if !isNil(o.PciController) {
		toSerialize["pci-controller"] = o.PciController
	}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !isNil(o.Dev) {
		toSerialize["dev"] = o.Dev
	}
	if !isNil(o.BusType) {
		toSerialize["bus-type"] = o.BusType
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableDiskInformationSummary struct {
	value *DiskInformationSummary
	isSet bool
}

func (v NullableDiskInformationSummary) Get() *DiskInformationSummary {
	return v.value
}

func (v *NullableDiskInformationSummary) Set(val *DiskInformationSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskInformationSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskInformationSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskInformationSummary(val *DiskInformationSummary) *NullableDiskInformationSummary {
	return &NullableDiskInformationSummary{value: val, isSet: true}
}

func (v NullableDiskInformationSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskInformationSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


