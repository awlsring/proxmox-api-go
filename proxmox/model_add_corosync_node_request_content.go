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

// checks if the AddCorosyncNodeRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCorosyncNodeRequestContent{}

// AddCorosyncNodeRequestContent struct for AddCorosyncNodeRequestContent
type AddCorosyncNodeRequestContent struct {
	// The api version on the new node
	Apiversion *float32 `json:"apiversion,omitempty"`
	// Do not throw an error if the node is already in the cluster
	Force *bool `json:"force,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link0 *string `json:"link0,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link1 *string `json:"link1,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link2 *string `json:"link2,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link3 *string `json:"link3,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link4 *string `json:"link4,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link5 *string `json:"link5,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link6 *string `json:"link6,omitempty"`
	// Address and priority of the link. Input as string with format address=<ip>,priority=<int>
	Link7 *string `json:"link7,omitempty"`
	// The IP address of the node to add
	NewNodeIp *string `json:"new_node_ip,omitempty"`
	// NodeID of the node to add
	Nodeid *float32 `json:"nodeid,omitempty"`
	// Votes thew new node should have
	Votes *float32 `json:"votes,omitempty"`
}

// NewAddCorosyncNodeRequestContent instantiates a new AddCorosyncNodeRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCorosyncNodeRequestContent() *AddCorosyncNodeRequestContent {
	this := AddCorosyncNodeRequestContent{}
	return &this
}

// NewAddCorosyncNodeRequestContentWithDefaults instantiates a new AddCorosyncNodeRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCorosyncNodeRequestContentWithDefaults() *AddCorosyncNodeRequestContent {
	this := AddCorosyncNodeRequestContent{}
	return &this
}

// GetApiversion returns the Apiversion field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetApiversion() float32 {
	if o == nil || IsNil(o.Apiversion) {
		var ret float32
		return ret
	}
	return *o.Apiversion
}

// GetApiversionOk returns a tuple with the Apiversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetApiversionOk() (*float32, bool) {
	if o == nil || IsNil(o.Apiversion) {
		return nil, false
	}
	return o.Apiversion, true
}

// HasApiversion returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasApiversion() bool {
	if o != nil && !IsNil(o.Apiversion) {
		return true
	}

	return false
}

// SetApiversion gets a reference to the given float32 and assigns it to the Apiversion field.
func (o *AddCorosyncNodeRequestContent) SetApiversion(v float32) {
	o.Apiversion = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *AddCorosyncNodeRequestContent) SetForce(v bool) {
	o.Force = &v
}

// GetLink0 returns the Link0 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink0() string {
	if o == nil || IsNil(o.Link0) {
		var ret string
		return ret
	}
	return *o.Link0
}

// GetLink0Ok returns a tuple with the Link0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink0Ok() (*string, bool) {
	if o == nil || IsNil(o.Link0) {
		return nil, false
	}
	return o.Link0, true
}

// HasLink0 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink0() bool {
	if o != nil && !IsNil(o.Link0) {
		return true
	}

	return false
}

// SetLink0 gets a reference to the given string and assigns it to the Link0 field.
func (o *AddCorosyncNodeRequestContent) SetLink0(v string) {
	o.Link0 = &v
}

// GetLink1 returns the Link1 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink1() string {
	if o == nil || IsNil(o.Link1) {
		var ret string
		return ret
	}
	return *o.Link1
}

// GetLink1Ok returns a tuple with the Link1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink1Ok() (*string, bool) {
	if o == nil || IsNil(o.Link1) {
		return nil, false
	}
	return o.Link1, true
}

// HasLink1 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink1() bool {
	if o != nil && !IsNil(o.Link1) {
		return true
	}

	return false
}

// SetLink1 gets a reference to the given string and assigns it to the Link1 field.
func (o *AddCorosyncNodeRequestContent) SetLink1(v string) {
	o.Link1 = &v
}

// GetLink2 returns the Link2 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink2() string {
	if o == nil || IsNil(o.Link2) {
		var ret string
		return ret
	}
	return *o.Link2
}

// GetLink2Ok returns a tuple with the Link2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink2Ok() (*string, bool) {
	if o == nil || IsNil(o.Link2) {
		return nil, false
	}
	return o.Link2, true
}

// HasLink2 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink2() bool {
	if o != nil && !IsNil(o.Link2) {
		return true
	}

	return false
}

// SetLink2 gets a reference to the given string and assigns it to the Link2 field.
func (o *AddCorosyncNodeRequestContent) SetLink2(v string) {
	o.Link2 = &v
}

// GetLink3 returns the Link3 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink3() string {
	if o == nil || IsNil(o.Link3) {
		var ret string
		return ret
	}
	return *o.Link3
}

// GetLink3Ok returns a tuple with the Link3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink3Ok() (*string, bool) {
	if o == nil || IsNil(o.Link3) {
		return nil, false
	}
	return o.Link3, true
}

// HasLink3 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink3() bool {
	if o != nil && !IsNil(o.Link3) {
		return true
	}

	return false
}

// SetLink3 gets a reference to the given string and assigns it to the Link3 field.
func (o *AddCorosyncNodeRequestContent) SetLink3(v string) {
	o.Link3 = &v
}

// GetLink4 returns the Link4 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink4() string {
	if o == nil || IsNil(o.Link4) {
		var ret string
		return ret
	}
	return *o.Link4
}

// GetLink4Ok returns a tuple with the Link4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink4Ok() (*string, bool) {
	if o == nil || IsNil(o.Link4) {
		return nil, false
	}
	return o.Link4, true
}

// HasLink4 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink4() bool {
	if o != nil && !IsNil(o.Link4) {
		return true
	}

	return false
}

// SetLink4 gets a reference to the given string and assigns it to the Link4 field.
func (o *AddCorosyncNodeRequestContent) SetLink4(v string) {
	o.Link4 = &v
}

// GetLink5 returns the Link5 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink5() string {
	if o == nil || IsNil(o.Link5) {
		var ret string
		return ret
	}
	return *o.Link5
}

// GetLink5Ok returns a tuple with the Link5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink5Ok() (*string, bool) {
	if o == nil || IsNil(o.Link5) {
		return nil, false
	}
	return o.Link5, true
}

// HasLink5 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink5() bool {
	if o != nil && !IsNil(o.Link5) {
		return true
	}

	return false
}

// SetLink5 gets a reference to the given string and assigns it to the Link5 field.
func (o *AddCorosyncNodeRequestContent) SetLink5(v string) {
	o.Link5 = &v
}

// GetLink6 returns the Link6 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink6() string {
	if o == nil || IsNil(o.Link6) {
		var ret string
		return ret
	}
	return *o.Link6
}

// GetLink6Ok returns a tuple with the Link6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink6Ok() (*string, bool) {
	if o == nil || IsNil(o.Link6) {
		return nil, false
	}
	return o.Link6, true
}

// HasLink6 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink6() bool {
	if o != nil && !IsNil(o.Link6) {
		return true
	}

	return false
}

// SetLink6 gets a reference to the given string and assigns it to the Link6 field.
func (o *AddCorosyncNodeRequestContent) SetLink6(v string) {
	o.Link6 = &v
}

// GetLink7 returns the Link7 field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetLink7() string {
	if o == nil || IsNil(o.Link7) {
		var ret string
		return ret
	}
	return *o.Link7
}

// GetLink7Ok returns a tuple with the Link7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetLink7Ok() (*string, bool) {
	if o == nil || IsNil(o.Link7) {
		return nil, false
	}
	return o.Link7, true
}

// HasLink7 returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasLink7() bool {
	if o != nil && !IsNil(o.Link7) {
		return true
	}

	return false
}

// SetLink7 gets a reference to the given string and assigns it to the Link7 field.
func (o *AddCorosyncNodeRequestContent) SetLink7(v string) {
	o.Link7 = &v
}

// GetNewNodeIp returns the NewNodeIp field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetNewNodeIp() string {
	if o == nil || IsNil(o.NewNodeIp) {
		var ret string
		return ret
	}
	return *o.NewNodeIp
}

// GetNewNodeIpOk returns a tuple with the NewNodeIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetNewNodeIpOk() (*string, bool) {
	if o == nil || IsNil(o.NewNodeIp) {
		return nil, false
	}
	return o.NewNodeIp, true
}

// HasNewNodeIp returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasNewNodeIp() bool {
	if o != nil && !IsNil(o.NewNodeIp) {
		return true
	}

	return false
}

// SetNewNodeIp gets a reference to the given string and assigns it to the NewNodeIp field.
func (o *AddCorosyncNodeRequestContent) SetNewNodeIp(v string) {
	o.NewNodeIp = &v
}

// GetNodeid returns the Nodeid field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetNodeid() float32 {
	if o == nil || IsNil(o.Nodeid) {
		var ret float32
		return ret
	}
	return *o.Nodeid
}

// GetNodeidOk returns a tuple with the Nodeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetNodeidOk() (*float32, bool) {
	if o == nil || IsNil(o.Nodeid) {
		return nil, false
	}
	return o.Nodeid, true
}

// HasNodeid returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasNodeid() bool {
	if o != nil && !IsNil(o.Nodeid) {
		return true
	}

	return false
}

// SetNodeid gets a reference to the given float32 and assigns it to the Nodeid field.
func (o *AddCorosyncNodeRequestContent) SetNodeid(v float32) {
	o.Nodeid = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *AddCorosyncNodeRequestContent) GetVotes() float32 {
	if o == nil || IsNil(o.Votes) {
		var ret float32
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeRequestContent) GetVotesOk() (*float32, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *AddCorosyncNodeRequestContent) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given float32 and assigns it to the Votes field.
func (o *AddCorosyncNodeRequestContent) SetVotes(v float32) {
	o.Votes = &v
}

func (o AddCorosyncNodeRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCorosyncNodeRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apiversion) {
		toSerialize["apiversion"] = o.Apiversion
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.Link0) {
		toSerialize["link0"] = o.Link0
	}
	if !IsNil(o.Link1) {
		toSerialize["link1"] = o.Link1
	}
	if !IsNil(o.Link2) {
		toSerialize["link2"] = o.Link2
	}
	if !IsNil(o.Link3) {
		toSerialize["link3"] = o.Link3
	}
	if !IsNil(o.Link4) {
		toSerialize["link4"] = o.Link4
	}
	if !IsNil(o.Link5) {
		toSerialize["link5"] = o.Link5
	}
	if !IsNil(o.Link6) {
		toSerialize["link6"] = o.Link6
	}
	if !IsNil(o.Link7) {
		toSerialize["link7"] = o.Link7
	}
	if !IsNil(o.NewNodeIp) {
		toSerialize["new_node_ip"] = o.NewNodeIp
	}
	if !IsNil(o.Nodeid) {
		toSerialize["nodeid"] = o.Nodeid
	}
	if !IsNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	return toSerialize, nil
}

type NullableAddCorosyncNodeRequestContent struct {
	value *AddCorosyncNodeRequestContent
	isSet bool
}

func (v NullableAddCorosyncNodeRequestContent) Get() *AddCorosyncNodeRequestContent {
	return v.value
}

func (v *NullableAddCorosyncNodeRequestContent) Set(val *AddCorosyncNodeRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCorosyncNodeRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCorosyncNodeRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCorosyncNodeRequestContent(val *AddCorosyncNodeRequestContent) *NullableAddCorosyncNodeRequestContent {
	return &NullableAddCorosyncNodeRequestContent{value: val, isSet: true}
}

func (v NullableAddCorosyncNodeRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCorosyncNodeRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


