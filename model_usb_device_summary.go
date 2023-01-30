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

// UsbDeviceSummary struct for UsbDeviceSummary
type UsbDeviceSummary struct {
	Busnum float32 `json:"busnum"`
	Devnum float32 `json:"devnum"`
	Level float32 `json:"level"`
	Port float32 `json:"port"`
	Prodid string `json:"prodid"`
	Vendid string `json:"vendid"`
	Speed string `json:"speed"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Product *string `json:"product,omitempty"`
	Serial *string `json:"serial,omitempty"`
	Usbpath *string `json:"usbpath,omitempty"`
}

// NewUsbDeviceSummary instantiates a new UsbDeviceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsbDeviceSummary(busnum float32, devnum float32, level float32, port float32, prodid string, vendid string, speed string) *UsbDeviceSummary {
	this := UsbDeviceSummary{}
	this.Busnum = busnum
	this.Devnum = devnum
	this.Level = level
	this.Port = port
	this.Prodid = prodid
	this.Vendid = vendid
	this.Speed = speed
	return &this
}

// NewUsbDeviceSummaryWithDefaults instantiates a new UsbDeviceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsbDeviceSummaryWithDefaults() *UsbDeviceSummary {
	this := UsbDeviceSummary{}
	return &this
}

// GetBusnum returns the Busnum field value
func (o *UsbDeviceSummary) GetBusnum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Busnum
}

// GetBusnumOk returns a tuple with the Busnum field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetBusnumOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Busnum, true
}

// SetBusnum sets field value
func (o *UsbDeviceSummary) SetBusnum(v float32) {
	o.Busnum = v
}

// GetDevnum returns the Devnum field value
func (o *UsbDeviceSummary) GetDevnum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Devnum
}

// GetDevnumOk returns a tuple with the Devnum field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetDevnumOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Devnum, true
}

// SetDevnum sets field value
func (o *UsbDeviceSummary) SetDevnum(v float32) {
	o.Devnum = v
}

// GetLevel returns the Level field value
func (o *UsbDeviceSummary) GetLevel() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetLevelOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *UsbDeviceSummary) SetLevel(v float32) {
	o.Level = v
}

// GetPort returns the Port field value
func (o *UsbDeviceSummary) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetPortOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *UsbDeviceSummary) SetPort(v float32) {
	o.Port = v
}

// GetProdid returns the Prodid field value
func (o *UsbDeviceSummary) GetProdid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prodid
}

// GetProdidOk returns a tuple with the Prodid field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetProdidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Prodid, true
}

// SetProdid sets field value
func (o *UsbDeviceSummary) SetProdid(v string) {
	o.Prodid = v
}

// GetVendid returns the Vendid field value
func (o *UsbDeviceSummary) GetVendid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendid
}

// GetVendidOk returns a tuple with the Vendid field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetVendidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vendid, true
}

// SetVendid sets field value
func (o *UsbDeviceSummary) SetVendid(v string) {
	o.Vendid = v
}

// GetSpeed returns the Speed field value
func (o *UsbDeviceSummary) GetSpeed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetSpeedOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *UsbDeviceSummary) SetSpeed(v string) {
	o.Speed = v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *UsbDeviceSummary) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetManufacturerOk() (*string, bool) {
	if o == nil || isNil(o.Manufacturer) {
    return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *UsbDeviceSummary) HasManufacturer() bool {
	if o != nil && !isNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *UsbDeviceSummary) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UsbDeviceSummary) GetProduct() string {
	if o == nil || isNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetProductOk() (*string, bool) {
	if o == nil || isNil(o.Product) {
    return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UsbDeviceSummary) HasProduct() bool {
	if o != nil && !isNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *UsbDeviceSummary) SetProduct(v string) {
	o.Product = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *UsbDeviceSummary) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *UsbDeviceSummary) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *UsbDeviceSummary) SetSerial(v string) {
	o.Serial = &v
}

// GetUsbpath returns the Usbpath field value if set, zero value otherwise.
func (o *UsbDeviceSummary) GetUsbpath() string {
	if o == nil || isNil(o.Usbpath) {
		var ret string
		return ret
	}
	return *o.Usbpath
}

// GetUsbpathOk returns a tuple with the Usbpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsbDeviceSummary) GetUsbpathOk() (*string, bool) {
	if o == nil || isNil(o.Usbpath) {
    return nil, false
	}
	return o.Usbpath, true
}

// HasUsbpath returns a boolean if a field has been set.
func (o *UsbDeviceSummary) HasUsbpath() bool {
	if o != nil && !isNil(o.Usbpath) {
		return true
	}

	return false
}

// SetUsbpath gets a reference to the given string and assigns it to the Usbpath field.
func (o *UsbDeviceSummary) SetUsbpath(v string) {
	o.Usbpath = &v
}

func (o UsbDeviceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["busnum"] = o.Busnum
	}
	if true {
		toSerialize["devnum"] = o.Devnum
	}
	if true {
		toSerialize["level"] = o.Level
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["prodid"] = o.Prodid
	}
	if true {
		toSerialize["vendid"] = o.Vendid
	}
	if true {
		toSerialize["speed"] = o.Speed
	}
	if !isNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !isNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Usbpath) {
		toSerialize["usbpath"] = o.Usbpath
	}
	return json.Marshal(toSerialize)
}

type NullableUsbDeviceSummary struct {
	value *UsbDeviceSummary
	isSet bool
}

func (v NullableUsbDeviceSummary) Get() *UsbDeviceSummary {
	return v.value
}

func (v *NullableUsbDeviceSummary) Set(val *UsbDeviceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableUsbDeviceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableUsbDeviceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsbDeviceSummary(val *UsbDeviceSummary) *NullableUsbDeviceSummary {
	return &NullableUsbDeviceSummary{value: val, isSet: true}
}

func (v NullableUsbDeviceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsbDeviceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

