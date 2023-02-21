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

// checks if the VirtualMachineNicBlockStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineNicBlockStatistics{}

// VirtualMachineNicBlockStatistics struct for VirtualMachineNicBlockStatistics
type VirtualMachineNicBlockStatistics struct {
	FlushTotalTimeNs *float32 `json:"flush_total_time_ns,omitempty"`
	RdBytes *float32 `json:"rd_bytes,omitempty"`
	TimedStats []string `json:"timed_stats,omitempty"`
	WrHighestOffset *float32 `json:"wr_highest_offset,omitempty"`
	RdTotalTimeNs *float32 `json:"rd_total_time_ns,omitempty"`
	FlushOperations *float32 `json:"flush_operations,omitempty"`
	WrOperations *float32 `json:"wr_operations,omitempty"`
	IdleTimeNs *float32 `json:"idle_time_ns,omitempty"`
	WrMerged *float32 `json:"wr_merged,omitempty"`
	InvalidRdOperations *float32 `json:"invalid_rd_operations,omitempty"`
	FailedFlushOperations *float32 `json:"failed_flush_operations,omitempty"`
	UnmapBytes *float32 `json:"unmap_bytes,omitempty"`
	FailedRdOperations *float32 `json:"failed_rd_operations,omitempty"`
	WrBytes *float32 `json:"wr_bytes,omitempty"`
	InvalidFlushOperations *float32 `json:"invalid_flush_operations,omitempty"`
	UnmapOperations *float32 `json:"unmap_operations,omitempty"`
	WrTotalTimeNs *float32 `json:"wr_total_time_ns,omitempty"`
	FailedWrOperations *float32 `json:"failed_wr_operations,omitempty"`
	UnmapMerged *float32 `json:"unmap_merged,omitempty"`
	InvalidWrOperations *float32 `json:"invalid_wr_operations,omitempty"`
	RdOperations *float32 `json:"rd_operations,omitempty"`
	UnmapTotalTimeNs *float32 `json:"unmap_total_time_ns,omitempty"`
	InvalidUnmapOperations *float32 `json:"invalid_unmap_operations,omitempty"`
	AccountFailed *bool `json:"account_failed,omitempty"`
	AccountInvalid *bool `json:"account_invalid,omitempty"`
	RdMerged *float32 `json:"rd_merged,omitempty"`
	FailedUnmapOperations *float32 `json:"failed_unmap_operations,omitempty"`
}

// NewVirtualMachineNicBlockStatistics instantiates a new VirtualMachineNicBlockStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineNicBlockStatistics() *VirtualMachineNicBlockStatistics {
	this := VirtualMachineNicBlockStatistics{}
	return &this
}

// NewVirtualMachineNicBlockStatisticsWithDefaults instantiates a new VirtualMachineNicBlockStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineNicBlockStatisticsWithDefaults() *VirtualMachineNicBlockStatistics {
	this := VirtualMachineNicBlockStatistics{}
	return &this
}

// GetFlushTotalTimeNs returns the FlushTotalTimeNs field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetFlushTotalTimeNs() float32 {
	if o == nil || IsNil(o.FlushTotalTimeNs) {
		var ret float32
		return ret
	}
	return *o.FlushTotalTimeNs
}

// GetFlushTotalTimeNsOk returns a tuple with the FlushTotalTimeNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetFlushTotalTimeNsOk() (*float32, bool) {
	if o == nil || IsNil(o.FlushTotalTimeNs) {
		return nil, false
	}
	return o.FlushTotalTimeNs, true
}

// HasFlushTotalTimeNs returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasFlushTotalTimeNs() bool {
	if o != nil && !IsNil(o.FlushTotalTimeNs) {
		return true
	}

	return false
}

// SetFlushTotalTimeNs gets a reference to the given float32 and assigns it to the FlushTotalTimeNs field.
func (o *VirtualMachineNicBlockStatistics) SetFlushTotalTimeNs(v float32) {
	o.FlushTotalTimeNs = &v
}

// GetRdBytes returns the RdBytes field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetRdBytes() float32 {
	if o == nil || IsNil(o.RdBytes) {
		var ret float32
		return ret
	}
	return *o.RdBytes
}

// GetRdBytesOk returns a tuple with the RdBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetRdBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.RdBytes) {
		return nil, false
	}
	return o.RdBytes, true
}

// HasRdBytes returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasRdBytes() bool {
	if o != nil && !IsNil(o.RdBytes) {
		return true
	}

	return false
}

// SetRdBytes gets a reference to the given float32 and assigns it to the RdBytes field.
func (o *VirtualMachineNicBlockStatistics) SetRdBytes(v float32) {
	o.RdBytes = &v
}

// GetTimedStats returns the TimedStats field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetTimedStats() []string {
	if o == nil || IsNil(o.TimedStats) {
		var ret []string
		return ret
	}
	return o.TimedStats
}

// GetTimedStatsOk returns a tuple with the TimedStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetTimedStatsOk() ([]string, bool) {
	if o == nil || IsNil(o.TimedStats) {
		return nil, false
	}
	return o.TimedStats, true
}

// HasTimedStats returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasTimedStats() bool {
	if o != nil && !IsNil(o.TimedStats) {
		return true
	}

	return false
}

// SetTimedStats gets a reference to the given []string and assigns it to the TimedStats field.
func (o *VirtualMachineNicBlockStatistics) SetTimedStats(v []string) {
	o.TimedStats = v
}

// GetWrHighestOffset returns the WrHighestOffset field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetWrHighestOffset() float32 {
	if o == nil || IsNil(o.WrHighestOffset) {
		var ret float32
		return ret
	}
	return *o.WrHighestOffset
}

// GetWrHighestOffsetOk returns a tuple with the WrHighestOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetWrHighestOffsetOk() (*float32, bool) {
	if o == nil || IsNil(o.WrHighestOffset) {
		return nil, false
	}
	return o.WrHighestOffset, true
}

// HasWrHighestOffset returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasWrHighestOffset() bool {
	if o != nil && !IsNil(o.WrHighestOffset) {
		return true
	}

	return false
}

// SetWrHighestOffset gets a reference to the given float32 and assigns it to the WrHighestOffset field.
func (o *VirtualMachineNicBlockStatistics) SetWrHighestOffset(v float32) {
	o.WrHighestOffset = &v
}

// GetRdTotalTimeNs returns the RdTotalTimeNs field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetRdTotalTimeNs() float32 {
	if o == nil || IsNil(o.RdTotalTimeNs) {
		var ret float32
		return ret
	}
	return *o.RdTotalTimeNs
}

// GetRdTotalTimeNsOk returns a tuple with the RdTotalTimeNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetRdTotalTimeNsOk() (*float32, bool) {
	if o == nil || IsNil(o.RdTotalTimeNs) {
		return nil, false
	}
	return o.RdTotalTimeNs, true
}

// HasRdTotalTimeNs returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasRdTotalTimeNs() bool {
	if o != nil && !IsNil(o.RdTotalTimeNs) {
		return true
	}

	return false
}

// SetRdTotalTimeNs gets a reference to the given float32 and assigns it to the RdTotalTimeNs field.
func (o *VirtualMachineNicBlockStatistics) SetRdTotalTimeNs(v float32) {
	o.RdTotalTimeNs = &v
}

// GetFlushOperations returns the FlushOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetFlushOperations() float32 {
	if o == nil || IsNil(o.FlushOperations) {
		var ret float32
		return ret
	}
	return *o.FlushOperations
}

// GetFlushOperationsOk returns a tuple with the FlushOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetFlushOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.FlushOperations) {
		return nil, false
	}
	return o.FlushOperations, true
}

// HasFlushOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasFlushOperations() bool {
	if o != nil && !IsNil(o.FlushOperations) {
		return true
	}

	return false
}

// SetFlushOperations gets a reference to the given float32 and assigns it to the FlushOperations field.
func (o *VirtualMachineNicBlockStatistics) SetFlushOperations(v float32) {
	o.FlushOperations = &v
}

// GetWrOperations returns the WrOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetWrOperations() float32 {
	if o == nil || IsNil(o.WrOperations) {
		var ret float32
		return ret
	}
	return *o.WrOperations
}

// GetWrOperationsOk returns a tuple with the WrOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetWrOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.WrOperations) {
		return nil, false
	}
	return o.WrOperations, true
}

// HasWrOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasWrOperations() bool {
	if o != nil && !IsNil(o.WrOperations) {
		return true
	}

	return false
}

// SetWrOperations gets a reference to the given float32 and assigns it to the WrOperations field.
func (o *VirtualMachineNicBlockStatistics) SetWrOperations(v float32) {
	o.WrOperations = &v
}

// GetIdleTimeNs returns the IdleTimeNs field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetIdleTimeNs() float32 {
	if o == nil || IsNil(o.IdleTimeNs) {
		var ret float32
		return ret
	}
	return *o.IdleTimeNs
}

// GetIdleTimeNsOk returns a tuple with the IdleTimeNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetIdleTimeNsOk() (*float32, bool) {
	if o == nil || IsNil(o.IdleTimeNs) {
		return nil, false
	}
	return o.IdleTimeNs, true
}

// HasIdleTimeNs returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasIdleTimeNs() bool {
	if o != nil && !IsNil(o.IdleTimeNs) {
		return true
	}

	return false
}

// SetIdleTimeNs gets a reference to the given float32 and assigns it to the IdleTimeNs field.
func (o *VirtualMachineNicBlockStatistics) SetIdleTimeNs(v float32) {
	o.IdleTimeNs = &v
}

// GetWrMerged returns the WrMerged field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetWrMerged() float32 {
	if o == nil || IsNil(o.WrMerged) {
		var ret float32
		return ret
	}
	return *o.WrMerged
}

// GetWrMergedOk returns a tuple with the WrMerged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetWrMergedOk() (*float32, bool) {
	if o == nil || IsNil(o.WrMerged) {
		return nil, false
	}
	return o.WrMerged, true
}

// HasWrMerged returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasWrMerged() bool {
	if o != nil && !IsNil(o.WrMerged) {
		return true
	}

	return false
}

// SetWrMerged gets a reference to the given float32 and assigns it to the WrMerged field.
func (o *VirtualMachineNicBlockStatistics) SetWrMerged(v float32) {
	o.WrMerged = &v
}

// GetInvalidRdOperations returns the InvalidRdOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetInvalidRdOperations() float32 {
	if o == nil || IsNil(o.InvalidRdOperations) {
		var ret float32
		return ret
	}
	return *o.InvalidRdOperations
}

// GetInvalidRdOperationsOk returns a tuple with the InvalidRdOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetInvalidRdOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.InvalidRdOperations) {
		return nil, false
	}
	return o.InvalidRdOperations, true
}

// HasInvalidRdOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasInvalidRdOperations() bool {
	if o != nil && !IsNil(o.InvalidRdOperations) {
		return true
	}

	return false
}

// SetInvalidRdOperations gets a reference to the given float32 and assigns it to the InvalidRdOperations field.
func (o *VirtualMachineNicBlockStatistics) SetInvalidRdOperations(v float32) {
	o.InvalidRdOperations = &v
}

// GetFailedFlushOperations returns the FailedFlushOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetFailedFlushOperations() float32 {
	if o == nil || IsNil(o.FailedFlushOperations) {
		var ret float32
		return ret
	}
	return *o.FailedFlushOperations
}

// GetFailedFlushOperationsOk returns a tuple with the FailedFlushOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetFailedFlushOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.FailedFlushOperations) {
		return nil, false
	}
	return o.FailedFlushOperations, true
}

// HasFailedFlushOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasFailedFlushOperations() bool {
	if o != nil && !IsNil(o.FailedFlushOperations) {
		return true
	}

	return false
}

// SetFailedFlushOperations gets a reference to the given float32 and assigns it to the FailedFlushOperations field.
func (o *VirtualMachineNicBlockStatistics) SetFailedFlushOperations(v float32) {
	o.FailedFlushOperations = &v
}

// GetUnmapBytes returns the UnmapBytes field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetUnmapBytes() float32 {
	if o == nil || IsNil(o.UnmapBytes) {
		var ret float32
		return ret
	}
	return *o.UnmapBytes
}

// GetUnmapBytesOk returns a tuple with the UnmapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetUnmapBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.UnmapBytes) {
		return nil, false
	}
	return o.UnmapBytes, true
}

// HasUnmapBytes returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasUnmapBytes() bool {
	if o != nil && !IsNil(o.UnmapBytes) {
		return true
	}

	return false
}

// SetUnmapBytes gets a reference to the given float32 and assigns it to the UnmapBytes field.
func (o *VirtualMachineNicBlockStatistics) SetUnmapBytes(v float32) {
	o.UnmapBytes = &v
}

// GetFailedRdOperations returns the FailedRdOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetFailedRdOperations() float32 {
	if o == nil || IsNil(o.FailedRdOperations) {
		var ret float32
		return ret
	}
	return *o.FailedRdOperations
}

// GetFailedRdOperationsOk returns a tuple with the FailedRdOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetFailedRdOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.FailedRdOperations) {
		return nil, false
	}
	return o.FailedRdOperations, true
}

// HasFailedRdOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasFailedRdOperations() bool {
	if o != nil && !IsNil(o.FailedRdOperations) {
		return true
	}

	return false
}

// SetFailedRdOperations gets a reference to the given float32 and assigns it to the FailedRdOperations field.
func (o *VirtualMachineNicBlockStatistics) SetFailedRdOperations(v float32) {
	o.FailedRdOperations = &v
}

// GetWrBytes returns the WrBytes field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetWrBytes() float32 {
	if o == nil || IsNil(o.WrBytes) {
		var ret float32
		return ret
	}
	return *o.WrBytes
}

// GetWrBytesOk returns a tuple with the WrBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetWrBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.WrBytes) {
		return nil, false
	}
	return o.WrBytes, true
}

// HasWrBytes returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasWrBytes() bool {
	if o != nil && !IsNil(o.WrBytes) {
		return true
	}

	return false
}

// SetWrBytes gets a reference to the given float32 and assigns it to the WrBytes field.
func (o *VirtualMachineNicBlockStatistics) SetWrBytes(v float32) {
	o.WrBytes = &v
}

// GetInvalidFlushOperations returns the InvalidFlushOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetInvalidFlushOperations() float32 {
	if o == nil || IsNil(o.InvalidFlushOperations) {
		var ret float32
		return ret
	}
	return *o.InvalidFlushOperations
}

// GetInvalidFlushOperationsOk returns a tuple with the InvalidFlushOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetInvalidFlushOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.InvalidFlushOperations) {
		return nil, false
	}
	return o.InvalidFlushOperations, true
}

// HasInvalidFlushOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasInvalidFlushOperations() bool {
	if o != nil && !IsNil(o.InvalidFlushOperations) {
		return true
	}

	return false
}

// SetInvalidFlushOperations gets a reference to the given float32 and assigns it to the InvalidFlushOperations field.
func (o *VirtualMachineNicBlockStatistics) SetInvalidFlushOperations(v float32) {
	o.InvalidFlushOperations = &v
}

// GetUnmapOperations returns the UnmapOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetUnmapOperations() float32 {
	if o == nil || IsNil(o.UnmapOperations) {
		var ret float32
		return ret
	}
	return *o.UnmapOperations
}

// GetUnmapOperationsOk returns a tuple with the UnmapOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetUnmapOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.UnmapOperations) {
		return nil, false
	}
	return o.UnmapOperations, true
}

// HasUnmapOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasUnmapOperations() bool {
	if o != nil && !IsNil(o.UnmapOperations) {
		return true
	}

	return false
}

// SetUnmapOperations gets a reference to the given float32 and assigns it to the UnmapOperations field.
func (o *VirtualMachineNicBlockStatistics) SetUnmapOperations(v float32) {
	o.UnmapOperations = &v
}

// GetWrTotalTimeNs returns the WrTotalTimeNs field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetWrTotalTimeNs() float32 {
	if o == nil || IsNil(o.WrTotalTimeNs) {
		var ret float32
		return ret
	}
	return *o.WrTotalTimeNs
}

// GetWrTotalTimeNsOk returns a tuple with the WrTotalTimeNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetWrTotalTimeNsOk() (*float32, bool) {
	if o == nil || IsNil(o.WrTotalTimeNs) {
		return nil, false
	}
	return o.WrTotalTimeNs, true
}

// HasWrTotalTimeNs returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasWrTotalTimeNs() bool {
	if o != nil && !IsNil(o.WrTotalTimeNs) {
		return true
	}

	return false
}

// SetWrTotalTimeNs gets a reference to the given float32 and assigns it to the WrTotalTimeNs field.
func (o *VirtualMachineNicBlockStatistics) SetWrTotalTimeNs(v float32) {
	o.WrTotalTimeNs = &v
}

// GetFailedWrOperations returns the FailedWrOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetFailedWrOperations() float32 {
	if o == nil || IsNil(o.FailedWrOperations) {
		var ret float32
		return ret
	}
	return *o.FailedWrOperations
}

// GetFailedWrOperationsOk returns a tuple with the FailedWrOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetFailedWrOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.FailedWrOperations) {
		return nil, false
	}
	return o.FailedWrOperations, true
}

// HasFailedWrOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasFailedWrOperations() bool {
	if o != nil && !IsNil(o.FailedWrOperations) {
		return true
	}

	return false
}

// SetFailedWrOperations gets a reference to the given float32 and assigns it to the FailedWrOperations field.
func (o *VirtualMachineNicBlockStatistics) SetFailedWrOperations(v float32) {
	o.FailedWrOperations = &v
}

// GetUnmapMerged returns the UnmapMerged field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetUnmapMerged() float32 {
	if o == nil || IsNil(o.UnmapMerged) {
		var ret float32
		return ret
	}
	return *o.UnmapMerged
}

// GetUnmapMergedOk returns a tuple with the UnmapMerged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetUnmapMergedOk() (*float32, bool) {
	if o == nil || IsNil(o.UnmapMerged) {
		return nil, false
	}
	return o.UnmapMerged, true
}

// HasUnmapMerged returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasUnmapMerged() bool {
	if o != nil && !IsNil(o.UnmapMerged) {
		return true
	}

	return false
}

// SetUnmapMerged gets a reference to the given float32 and assigns it to the UnmapMerged field.
func (o *VirtualMachineNicBlockStatistics) SetUnmapMerged(v float32) {
	o.UnmapMerged = &v
}

// GetInvalidWrOperations returns the InvalidWrOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetInvalidWrOperations() float32 {
	if o == nil || IsNil(o.InvalidWrOperations) {
		var ret float32
		return ret
	}
	return *o.InvalidWrOperations
}

// GetInvalidWrOperationsOk returns a tuple with the InvalidWrOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetInvalidWrOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.InvalidWrOperations) {
		return nil, false
	}
	return o.InvalidWrOperations, true
}

// HasInvalidWrOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasInvalidWrOperations() bool {
	if o != nil && !IsNil(o.InvalidWrOperations) {
		return true
	}

	return false
}

// SetInvalidWrOperations gets a reference to the given float32 and assigns it to the InvalidWrOperations field.
func (o *VirtualMachineNicBlockStatistics) SetInvalidWrOperations(v float32) {
	o.InvalidWrOperations = &v
}

// GetRdOperations returns the RdOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetRdOperations() float32 {
	if o == nil || IsNil(o.RdOperations) {
		var ret float32
		return ret
	}
	return *o.RdOperations
}

// GetRdOperationsOk returns a tuple with the RdOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetRdOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.RdOperations) {
		return nil, false
	}
	return o.RdOperations, true
}

// HasRdOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasRdOperations() bool {
	if o != nil && !IsNil(o.RdOperations) {
		return true
	}

	return false
}

// SetRdOperations gets a reference to the given float32 and assigns it to the RdOperations field.
func (o *VirtualMachineNicBlockStatistics) SetRdOperations(v float32) {
	o.RdOperations = &v
}

// GetUnmapTotalTimeNs returns the UnmapTotalTimeNs field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetUnmapTotalTimeNs() float32 {
	if o == nil || IsNil(o.UnmapTotalTimeNs) {
		var ret float32
		return ret
	}
	return *o.UnmapTotalTimeNs
}

// GetUnmapTotalTimeNsOk returns a tuple with the UnmapTotalTimeNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetUnmapTotalTimeNsOk() (*float32, bool) {
	if o == nil || IsNil(o.UnmapTotalTimeNs) {
		return nil, false
	}
	return o.UnmapTotalTimeNs, true
}

// HasUnmapTotalTimeNs returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasUnmapTotalTimeNs() bool {
	if o != nil && !IsNil(o.UnmapTotalTimeNs) {
		return true
	}

	return false
}

// SetUnmapTotalTimeNs gets a reference to the given float32 and assigns it to the UnmapTotalTimeNs field.
func (o *VirtualMachineNicBlockStatistics) SetUnmapTotalTimeNs(v float32) {
	o.UnmapTotalTimeNs = &v
}

// GetInvalidUnmapOperations returns the InvalidUnmapOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetInvalidUnmapOperations() float32 {
	if o == nil || IsNil(o.InvalidUnmapOperations) {
		var ret float32
		return ret
	}
	return *o.InvalidUnmapOperations
}

// GetInvalidUnmapOperationsOk returns a tuple with the InvalidUnmapOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetInvalidUnmapOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.InvalidUnmapOperations) {
		return nil, false
	}
	return o.InvalidUnmapOperations, true
}

// HasInvalidUnmapOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasInvalidUnmapOperations() bool {
	if o != nil && !IsNil(o.InvalidUnmapOperations) {
		return true
	}

	return false
}

// SetInvalidUnmapOperations gets a reference to the given float32 and assigns it to the InvalidUnmapOperations field.
func (o *VirtualMachineNicBlockStatistics) SetInvalidUnmapOperations(v float32) {
	o.InvalidUnmapOperations = &v
}

// GetAccountFailed returns the AccountFailed field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetAccountFailed() bool {
	if o == nil || IsNil(o.AccountFailed) {
		var ret bool
		return ret
	}
	return *o.AccountFailed
}

// GetAccountFailedOk returns a tuple with the AccountFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetAccountFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountFailed) {
		return nil, false
	}
	return o.AccountFailed, true
}

// HasAccountFailed returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasAccountFailed() bool {
	if o != nil && !IsNil(o.AccountFailed) {
		return true
	}

	return false
}

// SetAccountFailed gets a reference to the given bool and assigns it to the AccountFailed field.
func (o *VirtualMachineNicBlockStatistics) SetAccountFailed(v bool) {
	o.AccountFailed = &v
}

// GetAccountInvalid returns the AccountInvalid field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetAccountInvalid() bool {
	if o == nil || IsNil(o.AccountInvalid) {
		var ret bool
		return ret
	}
	return *o.AccountInvalid
}

// GetAccountInvalidOk returns a tuple with the AccountInvalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetAccountInvalidOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountInvalid) {
		return nil, false
	}
	return o.AccountInvalid, true
}

// HasAccountInvalid returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasAccountInvalid() bool {
	if o != nil && !IsNil(o.AccountInvalid) {
		return true
	}

	return false
}

// SetAccountInvalid gets a reference to the given bool and assigns it to the AccountInvalid field.
func (o *VirtualMachineNicBlockStatistics) SetAccountInvalid(v bool) {
	o.AccountInvalid = &v
}

// GetRdMerged returns the RdMerged field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetRdMerged() float32 {
	if o == nil || IsNil(o.RdMerged) {
		var ret float32
		return ret
	}
	return *o.RdMerged
}

// GetRdMergedOk returns a tuple with the RdMerged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetRdMergedOk() (*float32, bool) {
	if o == nil || IsNil(o.RdMerged) {
		return nil, false
	}
	return o.RdMerged, true
}

// HasRdMerged returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasRdMerged() bool {
	if o != nil && !IsNil(o.RdMerged) {
		return true
	}

	return false
}

// SetRdMerged gets a reference to the given float32 and assigns it to the RdMerged field.
func (o *VirtualMachineNicBlockStatistics) SetRdMerged(v float32) {
	o.RdMerged = &v
}

// GetFailedUnmapOperations returns the FailedUnmapOperations field value if set, zero value otherwise.
func (o *VirtualMachineNicBlockStatistics) GetFailedUnmapOperations() float32 {
	if o == nil || IsNil(o.FailedUnmapOperations) {
		var ret float32
		return ret
	}
	return *o.FailedUnmapOperations
}

// GetFailedUnmapOperationsOk returns a tuple with the FailedUnmapOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineNicBlockStatistics) GetFailedUnmapOperationsOk() (*float32, bool) {
	if o == nil || IsNil(o.FailedUnmapOperations) {
		return nil, false
	}
	return o.FailedUnmapOperations, true
}

// HasFailedUnmapOperations returns a boolean if a field has been set.
func (o *VirtualMachineNicBlockStatistics) HasFailedUnmapOperations() bool {
	if o != nil && !IsNil(o.FailedUnmapOperations) {
		return true
	}

	return false
}

// SetFailedUnmapOperations gets a reference to the given float32 and assigns it to the FailedUnmapOperations field.
func (o *VirtualMachineNicBlockStatistics) SetFailedUnmapOperations(v float32) {
	o.FailedUnmapOperations = &v
}

func (o VirtualMachineNicBlockStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineNicBlockStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlushTotalTimeNs) {
		toSerialize["flush_total_time_ns"] = o.FlushTotalTimeNs
	}
	if !IsNil(o.RdBytes) {
		toSerialize["rd_bytes"] = o.RdBytes
	}
	if !IsNil(o.TimedStats) {
		toSerialize["timed_stats"] = o.TimedStats
	}
	if !IsNil(o.WrHighestOffset) {
		toSerialize["wr_highest_offset"] = o.WrHighestOffset
	}
	if !IsNil(o.RdTotalTimeNs) {
		toSerialize["rd_total_time_ns"] = o.RdTotalTimeNs
	}
	if !IsNil(o.FlushOperations) {
		toSerialize["flush_operations"] = o.FlushOperations
	}
	if !IsNil(o.WrOperations) {
		toSerialize["wr_operations"] = o.WrOperations
	}
	if !IsNil(o.IdleTimeNs) {
		toSerialize["idle_time_ns"] = o.IdleTimeNs
	}
	if !IsNil(o.WrMerged) {
		toSerialize["wr_merged"] = o.WrMerged
	}
	if !IsNil(o.InvalidRdOperations) {
		toSerialize["invalid_rd_operations"] = o.InvalidRdOperations
	}
	if !IsNil(o.FailedFlushOperations) {
		toSerialize["failed_flush_operations"] = o.FailedFlushOperations
	}
	if !IsNil(o.UnmapBytes) {
		toSerialize["unmap_bytes"] = o.UnmapBytes
	}
	if !IsNil(o.FailedRdOperations) {
		toSerialize["failed_rd_operations"] = o.FailedRdOperations
	}
	if !IsNil(o.WrBytes) {
		toSerialize["wr_bytes"] = o.WrBytes
	}
	if !IsNil(o.InvalidFlushOperations) {
		toSerialize["invalid_flush_operations"] = o.InvalidFlushOperations
	}
	if !IsNil(o.UnmapOperations) {
		toSerialize["unmap_operations"] = o.UnmapOperations
	}
	if !IsNil(o.WrTotalTimeNs) {
		toSerialize["wr_total_time_ns"] = o.WrTotalTimeNs
	}
	if !IsNil(o.FailedWrOperations) {
		toSerialize["failed_wr_operations"] = o.FailedWrOperations
	}
	if !IsNil(o.UnmapMerged) {
		toSerialize["unmap_merged"] = o.UnmapMerged
	}
	if !IsNil(o.InvalidWrOperations) {
		toSerialize["invalid_wr_operations"] = o.InvalidWrOperations
	}
	if !IsNil(o.RdOperations) {
		toSerialize["rd_operations"] = o.RdOperations
	}
	if !IsNil(o.UnmapTotalTimeNs) {
		toSerialize["unmap_total_time_ns"] = o.UnmapTotalTimeNs
	}
	if !IsNil(o.InvalidUnmapOperations) {
		toSerialize["invalid_unmap_operations"] = o.InvalidUnmapOperations
	}
	if !IsNil(o.AccountFailed) {
		toSerialize["account_failed"] = o.AccountFailed
	}
	if !IsNil(o.AccountInvalid) {
		toSerialize["account_invalid"] = o.AccountInvalid
	}
	if !IsNil(o.RdMerged) {
		toSerialize["rd_merged"] = o.RdMerged
	}
	if !IsNil(o.FailedUnmapOperations) {
		toSerialize["failed_unmap_operations"] = o.FailedUnmapOperations
	}
	return toSerialize, nil
}

type NullableVirtualMachineNicBlockStatistics struct {
	value *VirtualMachineNicBlockStatistics
	isSet bool
}

func (v NullableVirtualMachineNicBlockStatistics) Get() *VirtualMachineNicBlockStatistics {
	return v.value
}

func (v *NullableVirtualMachineNicBlockStatistics) Set(val *VirtualMachineNicBlockStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineNicBlockStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineNicBlockStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineNicBlockStatistics(val *VirtualMachineNicBlockStatistics) *NullableVirtualMachineNicBlockStatistics {
	return &NullableVirtualMachineNicBlockStatistics{value: val, isSet: true}
}

func (v NullableVirtualMachineNicBlockStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineNicBlockStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


