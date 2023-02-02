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

// checks if the UpdateNetworkInterfaceRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkInterfaceRequestContent{}

// UpdateNetworkInterfaceRequestContent struct for UpdateNetworkInterfaceRequestContent
type UpdateNetworkInterfaceRequestContent struct {
	Type NetworkInterfaceType `json:"type"`
	Address *string `json:"address,omitempty"`
	Address6 *string `json:"address6,omitempty"`
	Autostart *bool `json:"autostart,omitempty"`
	BondPrimary *string `json:"bond_primary,omitempty"`
	BondMode *NetworkInterfaceBondMode `json:"bond_mode,omitempty"`
	BondXmitHashPolicy *NetworkInterfaceBondHashPolicy `json:"bond_xmit_hash_policy,omitempty"`
	BridgePorts *string `json:"bridge_ports,omitempty"`
	BridgeVlanAware *bool `json:"bridge_vlan_aware,omitempty"`
	Cidr *string `json:"cidr,omitempty"`
	Cidr6 *string `json:"cidr6,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Comments6 *string `json:"comments6,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Gateway6 *string `json:"gateway6,omitempty"`
	Mtu *float32 `json:"mtu,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Netmask6 *float32 `json:"netmask6,omitempty"`
	OvsBonds *string `json:"ovs_bonds,omitempty"`
	OvsOptions *string `json:"ovs_options,omitempty"`
	OvsPorts *string `json:"ovs_ports,omitempty"`
	OvsTag *float32 `json:"ovs_tag,omitempty"`
	OvsBridge *string `json:"ovs_bridge,omitempty"`
	Slaves []string `json:"slaves,omitempty"`
	VlanId *float32 `json:"vlan-id,omitempty"`
	VlanRawDevice *string `json:"vlan-raw-device,omitempty"`
}

// NewUpdateNetworkInterfaceRequestContent instantiates a new UpdateNetworkInterfaceRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkInterfaceRequestContent(type_ NetworkInterfaceType) *UpdateNetworkInterfaceRequestContent {
	this := UpdateNetworkInterfaceRequestContent{}
	this.Type = type_
	return &this
}

// NewUpdateNetworkInterfaceRequestContentWithDefaults instantiates a new UpdateNetworkInterfaceRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkInterfaceRequestContentWithDefaults() *UpdateNetworkInterfaceRequestContent {
	this := UpdateNetworkInterfaceRequestContent{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateNetworkInterfaceRequestContent) GetType() NetworkInterfaceType {
	if o == nil {
		var ret NetworkInterfaceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetTypeOk() (*NetworkInterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateNetworkInterfaceRequestContent) SetType(v NetworkInterfaceType) {
	o.Type = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateNetworkInterfaceRequestContent) SetAddress(v string) {
	o.Address = &v
}

// GetAddress6 returns the Address6 field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetAddress6() string {
	if o == nil || isNil(o.Address6) {
		var ret string
		return ret
	}
	return *o.Address6
}

// GetAddress6Ok returns a tuple with the Address6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetAddress6Ok() (*string, bool) {
	if o == nil || isNil(o.Address6) {
		return nil, false
	}
	return o.Address6, true
}

// HasAddress6 returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasAddress6() bool {
	if o != nil && !isNil(o.Address6) {
		return true
	}

	return false
}

// SetAddress6 gets a reference to the given string and assigns it to the Address6 field.
func (o *UpdateNetworkInterfaceRequestContent) SetAddress6(v string) {
	o.Address6 = &v
}

// GetAutostart returns the Autostart field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetAutostart() bool {
	if o == nil || isNil(o.Autostart) {
		var ret bool
		return ret
	}
	return *o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetAutostartOk() (*bool, bool) {
	if o == nil || isNil(o.Autostart) {
		return nil, false
	}
	return o.Autostart, true
}

// HasAutostart returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasAutostart() bool {
	if o != nil && !isNil(o.Autostart) {
		return true
	}

	return false
}

// SetAutostart gets a reference to the given bool and assigns it to the Autostart field.
func (o *UpdateNetworkInterfaceRequestContent) SetAutostart(v bool) {
	o.Autostart = &v
}

// GetBondPrimary returns the BondPrimary field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetBondPrimary() string {
	if o == nil || isNil(o.BondPrimary) {
		var ret string
		return ret
	}
	return *o.BondPrimary
}

// GetBondPrimaryOk returns a tuple with the BondPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetBondPrimaryOk() (*string, bool) {
	if o == nil || isNil(o.BondPrimary) {
		return nil, false
	}
	return o.BondPrimary, true
}

// HasBondPrimary returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasBondPrimary() bool {
	if o != nil && !isNil(o.BondPrimary) {
		return true
	}

	return false
}

// SetBondPrimary gets a reference to the given string and assigns it to the BondPrimary field.
func (o *UpdateNetworkInterfaceRequestContent) SetBondPrimary(v string) {
	o.BondPrimary = &v
}

// GetBondMode returns the BondMode field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetBondMode() NetworkInterfaceBondMode {
	if o == nil || isNil(o.BondMode) {
		var ret NetworkInterfaceBondMode
		return ret
	}
	return *o.BondMode
}

// GetBondModeOk returns a tuple with the BondMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetBondModeOk() (*NetworkInterfaceBondMode, bool) {
	if o == nil || isNil(o.BondMode) {
		return nil, false
	}
	return o.BondMode, true
}

// HasBondMode returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasBondMode() bool {
	if o != nil && !isNil(o.BondMode) {
		return true
	}

	return false
}

// SetBondMode gets a reference to the given NetworkInterfaceBondMode and assigns it to the BondMode field.
func (o *UpdateNetworkInterfaceRequestContent) SetBondMode(v NetworkInterfaceBondMode) {
	o.BondMode = &v
}

// GetBondXmitHashPolicy returns the BondXmitHashPolicy field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetBondXmitHashPolicy() NetworkInterfaceBondHashPolicy {
	if o == nil || isNil(o.BondXmitHashPolicy) {
		var ret NetworkInterfaceBondHashPolicy
		return ret
	}
	return *o.BondXmitHashPolicy
}

// GetBondXmitHashPolicyOk returns a tuple with the BondXmitHashPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetBondXmitHashPolicyOk() (*NetworkInterfaceBondHashPolicy, bool) {
	if o == nil || isNil(o.BondXmitHashPolicy) {
		return nil, false
	}
	return o.BondXmitHashPolicy, true
}

// HasBondXmitHashPolicy returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasBondXmitHashPolicy() bool {
	if o != nil && !isNil(o.BondXmitHashPolicy) {
		return true
	}

	return false
}

// SetBondXmitHashPolicy gets a reference to the given NetworkInterfaceBondHashPolicy and assigns it to the BondXmitHashPolicy field.
func (o *UpdateNetworkInterfaceRequestContent) SetBondXmitHashPolicy(v NetworkInterfaceBondHashPolicy) {
	o.BondXmitHashPolicy = &v
}

// GetBridgePorts returns the BridgePorts field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetBridgePorts() string {
	if o == nil || isNil(o.BridgePorts) {
		var ret string
		return ret
	}
	return *o.BridgePorts
}

// GetBridgePortsOk returns a tuple with the BridgePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetBridgePortsOk() (*string, bool) {
	if o == nil || isNil(o.BridgePorts) {
		return nil, false
	}
	return o.BridgePorts, true
}

// HasBridgePorts returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasBridgePorts() bool {
	if o != nil && !isNil(o.BridgePorts) {
		return true
	}

	return false
}

// SetBridgePorts gets a reference to the given string and assigns it to the BridgePorts field.
func (o *UpdateNetworkInterfaceRequestContent) SetBridgePorts(v string) {
	o.BridgePorts = &v
}

// GetBridgeVlanAware returns the BridgeVlanAware field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetBridgeVlanAware() bool {
	if o == nil || isNil(o.BridgeVlanAware) {
		var ret bool
		return ret
	}
	return *o.BridgeVlanAware
}

// GetBridgeVlanAwareOk returns a tuple with the BridgeVlanAware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetBridgeVlanAwareOk() (*bool, bool) {
	if o == nil || isNil(o.BridgeVlanAware) {
		return nil, false
	}
	return o.BridgeVlanAware, true
}

// HasBridgeVlanAware returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasBridgeVlanAware() bool {
	if o != nil && !isNil(o.BridgeVlanAware) {
		return true
	}

	return false
}

// SetBridgeVlanAware gets a reference to the given bool and assigns it to the BridgeVlanAware field.
func (o *UpdateNetworkInterfaceRequestContent) SetBridgeVlanAware(v bool) {
	o.BridgeVlanAware = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *UpdateNetworkInterfaceRequestContent) SetCidr(v string) {
	o.Cidr = &v
}

// GetCidr6 returns the Cidr6 field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetCidr6() string {
	if o == nil || isNil(o.Cidr6) {
		var ret string
		return ret
	}
	return *o.Cidr6
}

// GetCidr6Ok returns a tuple with the Cidr6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetCidr6Ok() (*string, bool) {
	if o == nil || isNil(o.Cidr6) {
		return nil, false
	}
	return o.Cidr6, true
}

// HasCidr6 returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasCidr6() bool {
	if o != nil && !isNil(o.Cidr6) {
		return true
	}

	return false
}

// SetCidr6 gets a reference to the given string and assigns it to the Cidr6 field.
func (o *UpdateNetworkInterfaceRequestContent) SetCidr6(v string) {
	o.Cidr6 = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetComments() string {
	if o == nil || isNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetCommentsOk() (*string, bool) {
	if o == nil || isNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasComments() bool {
	if o != nil && !isNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *UpdateNetworkInterfaceRequestContent) SetComments(v string) {
	o.Comments = &v
}

// GetComments6 returns the Comments6 field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetComments6() string {
	if o == nil || isNil(o.Comments6) {
		var ret string
		return ret
	}
	return *o.Comments6
}

// GetComments6Ok returns a tuple with the Comments6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetComments6Ok() (*string, bool) {
	if o == nil || isNil(o.Comments6) {
		return nil, false
	}
	return o.Comments6, true
}

// HasComments6 returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasComments6() bool {
	if o != nil && !isNil(o.Comments6) {
		return true
	}

	return false
}

// SetComments6 gets a reference to the given string and assigns it to the Comments6 field.
func (o *UpdateNetworkInterfaceRequestContent) SetComments6(v string) {
	o.Comments6 = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *UpdateNetworkInterfaceRequestContent) SetGateway(v string) {
	o.Gateway = &v
}

// GetGateway6 returns the Gateway6 field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetGateway6() string {
	if o == nil || isNil(o.Gateway6) {
		var ret string
		return ret
	}
	return *o.Gateway6
}

// GetGateway6Ok returns a tuple with the Gateway6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetGateway6Ok() (*string, bool) {
	if o == nil || isNil(o.Gateway6) {
		return nil, false
	}
	return o.Gateway6, true
}

// HasGateway6 returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasGateway6() bool {
	if o != nil && !isNil(o.Gateway6) {
		return true
	}

	return false
}

// SetGateway6 gets a reference to the given string and assigns it to the Gateway6 field.
func (o *UpdateNetworkInterfaceRequestContent) SetGateway6(v string) {
	o.Gateway6 = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetMtu() float32 {
	if o == nil || isNil(o.Mtu) {
		var ret float32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetMtuOk() (*float32, bool) {
	if o == nil || isNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasMtu() bool {
	if o != nil && !isNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given float32 and assigns it to the Mtu field.
func (o *UpdateNetworkInterfaceRequestContent) SetMtu(v float32) {
	o.Mtu = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetNetmask() string {
	if o == nil || isNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetNetmaskOk() (*string, bool) {
	if o == nil || isNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasNetmask() bool {
	if o != nil && !isNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *UpdateNetworkInterfaceRequestContent) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNetmask6 returns the Netmask6 field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetNetmask6() float32 {
	if o == nil || isNil(o.Netmask6) {
		var ret float32
		return ret
	}
	return *o.Netmask6
}

// GetNetmask6Ok returns a tuple with the Netmask6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetNetmask6Ok() (*float32, bool) {
	if o == nil || isNil(o.Netmask6) {
		return nil, false
	}
	return o.Netmask6, true
}

// HasNetmask6 returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasNetmask6() bool {
	if o != nil && !isNil(o.Netmask6) {
		return true
	}

	return false
}

// SetNetmask6 gets a reference to the given float32 and assigns it to the Netmask6 field.
func (o *UpdateNetworkInterfaceRequestContent) SetNetmask6(v float32) {
	o.Netmask6 = &v
}

// GetOvsBonds returns the OvsBonds field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsBonds() string {
	if o == nil || isNil(o.OvsBonds) {
		var ret string
		return ret
	}
	return *o.OvsBonds
}

// GetOvsBondsOk returns a tuple with the OvsBonds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsBondsOk() (*string, bool) {
	if o == nil || isNil(o.OvsBonds) {
		return nil, false
	}
	return o.OvsBonds, true
}

// HasOvsBonds returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasOvsBonds() bool {
	if o != nil && !isNil(o.OvsBonds) {
		return true
	}

	return false
}

// SetOvsBonds gets a reference to the given string and assigns it to the OvsBonds field.
func (o *UpdateNetworkInterfaceRequestContent) SetOvsBonds(v string) {
	o.OvsBonds = &v
}

// GetOvsOptions returns the OvsOptions field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsOptions() string {
	if o == nil || isNil(o.OvsOptions) {
		var ret string
		return ret
	}
	return *o.OvsOptions
}

// GetOvsOptionsOk returns a tuple with the OvsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsOptionsOk() (*string, bool) {
	if o == nil || isNil(o.OvsOptions) {
		return nil, false
	}
	return o.OvsOptions, true
}

// HasOvsOptions returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasOvsOptions() bool {
	if o != nil && !isNil(o.OvsOptions) {
		return true
	}

	return false
}

// SetOvsOptions gets a reference to the given string and assigns it to the OvsOptions field.
func (o *UpdateNetworkInterfaceRequestContent) SetOvsOptions(v string) {
	o.OvsOptions = &v
}

// GetOvsPorts returns the OvsPorts field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsPorts() string {
	if o == nil || isNil(o.OvsPorts) {
		var ret string
		return ret
	}
	return *o.OvsPorts
}

// GetOvsPortsOk returns a tuple with the OvsPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsPortsOk() (*string, bool) {
	if o == nil || isNil(o.OvsPorts) {
		return nil, false
	}
	return o.OvsPorts, true
}

// HasOvsPorts returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasOvsPorts() bool {
	if o != nil && !isNil(o.OvsPorts) {
		return true
	}

	return false
}

// SetOvsPorts gets a reference to the given string and assigns it to the OvsPorts field.
func (o *UpdateNetworkInterfaceRequestContent) SetOvsPorts(v string) {
	o.OvsPorts = &v
}

// GetOvsTag returns the OvsTag field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsTag() float32 {
	if o == nil || isNil(o.OvsTag) {
		var ret float32
		return ret
	}
	return *o.OvsTag
}

// GetOvsTagOk returns a tuple with the OvsTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsTagOk() (*float32, bool) {
	if o == nil || isNil(o.OvsTag) {
		return nil, false
	}
	return o.OvsTag, true
}

// HasOvsTag returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasOvsTag() bool {
	if o != nil && !isNil(o.OvsTag) {
		return true
	}

	return false
}

// SetOvsTag gets a reference to the given float32 and assigns it to the OvsTag field.
func (o *UpdateNetworkInterfaceRequestContent) SetOvsTag(v float32) {
	o.OvsTag = &v
}

// GetOvsBridge returns the OvsBridge field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsBridge() string {
	if o == nil || isNil(o.OvsBridge) {
		var ret string
		return ret
	}
	return *o.OvsBridge
}

// GetOvsBridgeOk returns a tuple with the OvsBridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetOvsBridgeOk() (*string, bool) {
	if o == nil || isNil(o.OvsBridge) {
		return nil, false
	}
	return o.OvsBridge, true
}

// HasOvsBridge returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasOvsBridge() bool {
	if o != nil && !isNil(o.OvsBridge) {
		return true
	}

	return false
}

// SetOvsBridge gets a reference to the given string and assigns it to the OvsBridge field.
func (o *UpdateNetworkInterfaceRequestContent) SetOvsBridge(v string) {
	o.OvsBridge = &v
}

// GetSlaves returns the Slaves field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetSlaves() []string {
	if o == nil || isNil(o.Slaves) {
		var ret []string
		return ret
	}
	return o.Slaves
}

// GetSlavesOk returns a tuple with the Slaves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetSlavesOk() ([]string, bool) {
	if o == nil || isNil(o.Slaves) {
		return nil, false
	}
	return o.Slaves, true
}

// HasSlaves returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasSlaves() bool {
	if o != nil && !isNil(o.Slaves) {
		return true
	}

	return false
}

// SetSlaves gets a reference to the given []string and assigns it to the Slaves field.
func (o *UpdateNetworkInterfaceRequestContent) SetSlaves(v []string) {
	o.Slaves = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetVlanId() float32 {
	if o == nil || isNil(o.VlanId) {
		var ret float32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetVlanIdOk() (*float32, bool) {
	if o == nil || isNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given float32 and assigns it to the VlanId field.
func (o *UpdateNetworkInterfaceRequestContent) SetVlanId(v float32) {
	o.VlanId = &v
}

// GetVlanRawDevice returns the VlanRawDevice field value if set, zero value otherwise.
func (o *UpdateNetworkInterfaceRequestContent) GetVlanRawDevice() string {
	if o == nil || isNil(o.VlanRawDevice) {
		var ret string
		return ret
	}
	return *o.VlanRawDevice
}

// GetVlanRawDeviceOk returns a tuple with the VlanRawDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkInterfaceRequestContent) GetVlanRawDeviceOk() (*string, bool) {
	if o == nil || isNil(o.VlanRawDevice) {
		return nil, false
	}
	return o.VlanRawDevice, true
}

// HasVlanRawDevice returns a boolean if a field has been set.
func (o *UpdateNetworkInterfaceRequestContent) HasVlanRawDevice() bool {
	if o != nil && !isNil(o.VlanRawDevice) {
		return true
	}

	return false
}

// SetVlanRawDevice gets a reference to the given string and assigns it to the VlanRawDevice field.
func (o *UpdateNetworkInterfaceRequestContent) SetVlanRawDevice(v string) {
	o.VlanRawDevice = &v
}

func (o UpdateNetworkInterfaceRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkInterfaceRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Address6) {
		toSerialize["address6"] = o.Address6
	}
	if !isNil(o.Autostart) {
		toSerialize["autostart"] = o.Autostart
	}
	if !isNil(o.BondPrimary) {
		toSerialize["bond_primary"] = o.BondPrimary
	}
	if !isNil(o.BondMode) {
		toSerialize["bond_mode"] = o.BondMode
	}
	if !isNil(o.BondXmitHashPolicy) {
		toSerialize["bond_xmit_hash_policy"] = o.BondXmitHashPolicy
	}
	if !isNil(o.BridgePorts) {
		toSerialize["bridge_ports"] = o.BridgePorts
	}
	if !isNil(o.BridgeVlanAware) {
		toSerialize["bridge_vlan_aware"] = o.BridgeVlanAware
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Cidr6) {
		toSerialize["cidr6"] = o.Cidr6
	}
	if !isNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !isNil(o.Comments6) {
		toSerialize["comments6"] = o.Comments6
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.Gateway6) {
		toSerialize["gateway6"] = o.Gateway6
	}
	if !isNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !isNil(o.Netmask) {
		toSerialize["netmask"] = o.Netmask
	}
	if !isNil(o.Netmask6) {
		toSerialize["netmask6"] = o.Netmask6
	}
	if !isNil(o.OvsBonds) {
		toSerialize["ovs_bonds"] = o.OvsBonds
	}
	if !isNil(o.OvsOptions) {
		toSerialize["ovs_options"] = o.OvsOptions
	}
	if !isNil(o.OvsPorts) {
		toSerialize["ovs_ports"] = o.OvsPorts
	}
	if !isNil(o.OvsTag) {
		toSerialize["ovs_tag"] = o.OvsTag
	}
	if !isNil(o.OvsBridge) {
		toSerialize["ovs_bridge"] = o.OvsBridge
	}
	if !isNil(o.Slaves) {
		toSerialize["slaves"] = o.Slaves
	}
	if !isNil(o.VlanId) {
		toSerialize["vlan-id"] = o.VlanId
	}
	if !isNil(o.VlanRawDevice) {
		toSerialize["vlan-raw-device"] = o.VlanRawDevice
	}
	return toSerialize, nil
}

type NullableUpdateNetworkInterfaceRequestContent struct {
	value *UpdateNetworkInterfaceRequestContent
	isSet bool
}

func (v NullableUpdateNetworkInterfaceRequestContent) Get() *UpdateNetworkInterfaceRequestContent {
	return v.value
}

func (v *NullableUpdateNetworkInterfaceRequestContent) Set(val *UpdateNetworkInterfaceRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkInterfaceRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkInterfaceRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkInterfaceRequestContent(val *UpdateNetworkInterfaceRequestContent) *NullableUpdateNetworkInterfaceRequestContent {
	return &NullableUpdateNetworkInterfaceRequestContent{value: val, isSet: true}
}

func (v NullableUpdateNetworkInterfaceRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkInterfaceRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


