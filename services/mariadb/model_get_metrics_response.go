/*
STACKIT MariaDB API

The STACKIT MariaDB API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

import (
	"encoding/json"
)

// checks if the GetMetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMetricsResponse{}

// GetMetricsResponse struct for GetMetricsResponse
type GetMetricsResponse struct {
	CpuIdleTime *int64 `json:"cpuIdleTime,omitempty"`
	// REQUIRED
	CpuLoadPercent *int64 `json:"cpuLoadPercent"`
	CpuSystemTime  *int64 `json:"cpuSystemTime,omitempty"`
	CpuUserTime    *int64 `json:"cpuUserTime,omitempty"`
	// REQUIRED
	DiskEphemeralTotal *int64 `json:"diskEphemeralTotal"`
	// REQUIRED
	DiskEphemeralUsed *int64 `json:"diskEphemeralUsed"`
	// REQUIRED
	DiskPersistentTotal *int64 `json:"diskPersistentTotal"`
	// REQUIRED
	DiskPersistentUsed *int64 `json:"diskPersistentUsed"`
	// REQUIRED
	Load1 *int64 `json:"load1"`
	// REQUIRED
	Load15 *int64 `json:"load15"`
	// REQUIRED
	Load5 *int64 `json:"load5"`
	// REQUIRED
	MemoryTotal *int64 `json:"memoryTotal"`
	// REQUIRED
	MemoryUsed *int64 `json:"memoryUsed"`
	// REQUIRED
	ParachuteDiskEphemeralActivated *int64 `json:"parachuteDiskEphemeralActivated"`
	// REQUIRED
	ParachuteDiskEphemeralTotal *int64 `json:"parachuteDiskEphemeralTotal"`
	// REQUIRED
	ParachuteDiskEphemeralUsed *int64 `json:"parachuteDiskEphemeralUsed"`
	// REQUIRED
	ParachuteDiskEphemeralUsedPercent *int64 `json:"parachuteDiskEphemeralUsedPercent"`
	// REQUIRED
	ParachuteDiskEphemeralUsedThreshold *int64 `json:"parachuteDiskEphemeralUsedThreshold"`
	// REQUIRED
	ParachuteDiskPersistentActivated *int64 `json:"parachuteDiskPersistentActivated"`
	// REQUIRED
	ParachuteDiskPersistentTotal *int64 `json:"parachuteDiskPersistentTotal"`
	// REQUIRED
	ParachuteDiskPersistentUsed *int64 `json:"parachuteDiskPersistentUsed"`
	// REQUIRED
	ParachuteDiskPersistentUsedPercent *int64 `json:"parachuteDiskPersistentUsedPercent"`
	// REQUIRED
	ParachuteDiskPersistentUsedThreshold *int64 `json:"parachuteDiskPersistentUsedThreshold"`
}

type _GetMetricsResponse GetMetricsResponse

// NewGetMetricsResponse instantiates a new GetMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetricsResponse(cpuLoadPercent *int64, diskEphemeralTotal *int64, diskEphemeralUsed *int64, diskPersistentTotal *int64, diskPersistentUsed *int64, load1 *int64, load15 *int64, load5 *int64, memoryTotal *int64, memoryUsed *int64, parachuteDiskEphemeralActivated *int64, parachuteDiskEphemeralTotal *int64, parachuteDiskEphemeralUsed *int64, parachuteDiskEphemeralUsedPercent *int64, parachuteDiskEphemeralUsedThreshold *int64, parachuteDiskPersistentActivated *int64, parachuteDiskPersistentTotal *int64, parachuteDiskPersistentUsed *int64, parachuteDiskPersistentUsedPercent *int64, parachuteDiskPersistentUsedThreshold *int64) *GetMetricsResponse {
	this := GetMetricsResponse{}
	this.CpuLoadPercent = cpuLoadPercent
	this.DiskEphemeralTotal = diskEphemeralTotal
	this.DiskEphemeralUsed = diskEphemeralUsed
	this.DiskPersistentTotal = diskPersistentTotal
	this.DiskPersistentUsed = diskPersistentUsed
	this.Load1 = load1
	this.Load15 = load15
	this.Load5 = load5
	this.MemoryTotal = memoryTotal
	this.MemoryUsed = memoryUsed
	this.ParachuteDiskEphemeralActivated = parachuteDiskEphemeralActivated
	this.ParachuteDiskEphemeralTotal = parachuteDiskEphemeralTotal
	this.ParachuteDiskEphemeralUsed = parachuteDiskEphemeralUsed
	this.ParachuteDiskEphemeralUsedPercent = parachuteDiskEphemeralUsedPercent
	this.ParachuteDiskEphemeralUsedThreshold = parachuteDiskEphemeralUsedThreshold
	this.ParachuteDiskPersistentActivated = parachuteDiskPersistentActivated
	this.ParachuteDiskPersistentTotal = parachuteDiskPersistentTotal
	this.ParachuteDiskPersistentUsed = parachuteDiskPersistentUsed
	this.ParachuteDiskPersistentUsedPercent = parachuteDiskPersistentUsedPercent
	this.ParachuteDiskPersistentUsedThreshold = parachuteDiskPersistentUsedThreshold
	return &this
}

// NewGetMetricsResponseWithDefaults instantiates a new GetMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetricsResponseWithDefaults() *GetMetricsResponse {
	this := GetMetricsResponse{}
	return &this
}

// GetCpuIdleTime returns the CpuIdleTime field value if set, zero value otherwise.
func (o *GetMetricsResponse) GetCpuIdleTime() *int64 {
	if o == nil || IsNil(o.CpuIdleTime) {
		var ret *int64
		return ret
	}
	return o.CpuIdleTime
}

// GetCpuIdleTimeOk returns a tuple with the CpuIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetCpuIdleTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuIdleTime) {
		return nil, false
	}
	return o.CpuIdleTime, true
}

// HasCpuIdleTime returns a boolean if a field has been set.
func (o *GetMetricsResponse) HasCpuIdleTime() bool {
	if o != nil && !IsNil(o.CpuIdleTime) {
		return true
	}

	return false
}

// SetCpuIdleTime gets a reference to the given int64 and assigns it to the CpuIdleTime field.
func (o *GetMetricsResponse) SetCpuIdleTime(v *int64) {
	o.CpuIdleTime = v
}

// GetCpuLoadPercent returns the CpuLoadPercent field value
func (o *GetMetricsResponse) GetCpuLoadPercent() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.CpuLoadPercent
}

// GetCpuLoadPercentOk returns a tuple with the CpuLoadPercent field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetCpuLoadPercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuLoadPercent, true
}

// SetCpuLoadPercent sets field value
func (o *GetMetricsResponse) SetCpuLoadPercent(v *int64) {
	o.CpuLoadPercent = v
}

// GetCpuSystemTime returns the CpuSystemTime field value if set, zero value otherwise.
func (o *GetMetricsResponse) GetCpuSystemTime() *int64 {
	if o == nil || IsNil(o.CpuSystemTime) {
		var ret *int64
		return ret
	}
	return o.CpuSystemTime
}

// GetCpuSystemTimeOk returns a tuple with the CpuSystemTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetCpuSystemTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuSystemTime) {
		return nil, false
	}
	return o.CpuSystemTime, true
}

// HasCpuSystemTime returns a boolean if a field has been set.
func (o *GetMetricsResponse) HasCpuSystemTime() bool {
	if o != nil && !IsNil(o.CpuSystemTime) {
		return true
	}

	return false
}

// SetCpuSystemTime gets a reference to the given int64 and assigns it to the CpuSystemTime field.
func (o *GetMetricsResponse) SetCpuSystemTime(v *int64) {
	o.CpuSystemTime = v
}

// GetCpuUserTime returns the CpuUserTime field value if set, zero value otherwise.
func (o *GetMetricsResponse) GetCpuUserTime() *int64 {
	if o == nil || IsNil(o.CpuUserTime) {
		var ret *int64
		return ret
	}
	return o.CpuUserTime
}

// GetCpuUserTimeOk returns a tuple with the CpuUserTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetCpuUserTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuUserTime) {
		return nil, false
	}
	return o.CpuUserTime, true
}

// HasCpuUserTime returns a boolean if a field has been set.
func (o *GetMetricsResponse) HasCpuUserTime() bool {
	if o != nil && !IsNil(o.CpuUserTime) {
		return true
	}

	return false
}

// SetCpuUserTime gets a reference to the given int64 and assigns it to the CpuUserTime field.
func (o *GetMetricsResponse) SetCpuUserTime(v *int64) {
	o.CpuUserTime = v
}

// GetDiskEphemeralTotal returns the DiskEphemeralTotal field value
func (o *GetMetricsResponse) GetDiskEphemeralTotal() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.DiskEphemeralTotal
}

// GetDiskEphemeralTotalOk returns a tuple with the DiskEphemeralTotal field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetDiskEphemeralTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskEphemeralTotal, true
}

// SetDiskEphemeralTotal sets field value
func (o *GetMetricsResponse) SetDiskEphemeralTotal(v *int64) {
	o.DiskEphemeralTotal = v
}

// GetDiskEphemeralUsed returns the DiskEphemeralUsed field value
func (o *GetMetricsResponse) GetDiskEphemeralUsed() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.DiskEphemeralUsed
}

// GetDiskEphemeralUsedOk returns a tuple with the DiskEphemeralUsed field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetDiskEphemeralUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskEphemeralUsed, true
}

// SetDiskEphemeralUsed sets field value
func (o *GetMetricsResponse) SetDiskEphemeralUsed(v *int64) {
	o.DiskEphemeralUsed = v
}

// GetDiskPersistentTotal returns the DiskPersistentTotal field value
func (o *GetMetricsResponse) GetDiskPersistentTotal() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.DiskPersistentTotal
}

// GetDiskPersistentTotalOk returns a tuple with the DiskPersistentTotal field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetDiskPersistentTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskPersistentTotal, true
}

// SetDiskPersistentTotal sets field value
func (o *GetMetricsResponse) SetDiskPersistentTotal(v *int64) {
	o.DiskPersistentTotal = v
}

// GetDiskPersistentUsed returns the DiskPersistentUsed field value
func (o *GetMetricsResponse) GetDiskPersistentUsed() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.DiskPersistentUsed
}

// GetDiskPersistentUsedOk returns a tuple with the DiskPersistentUsed field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetDiskPersistentUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskPersistentUsed, true
}

// SetDiskPersistentUsed sets field value
func (o *GetMetricsResponse) SetDiskPersistentUsed(v *int64) {
	o.DiskPersistentUsed = v
}

// GetLoad1 returns the Load1 field value
func (o *GetMetricsResponse) GetLoad1() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.Load1
}

// GetLoad1Ok returns a tuple with the Load1 field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetLoad1Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Load1, true
}

// SetLoad1 sets field value
func (o *GetMetricsResponse) SetLoad1(v *int64) {
	o.Load1 = v
}

// GetLoad15 returns the Load15 field value
func (o *GetMetricsResponse) GetLoad15() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.Load15
}

// GetLoad15Ok returns a tuple with the Load15 field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetLoad15Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Load15, true
}

// SetLoad15 sets field value
func (o *GetMetricsResponse) SetLoad15(v *int64) {
	o.Load15 = v
}

// GetLoad5 returns the Load5 field value
func (o *GetMetricsResponse) GetLoad5() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.Load5
}

// GetLoad5Ok returns a tuple with the Load5 field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetLoad5Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Load5, true
}

// SetLoad5 sets field value
func (o *GetMetricsResponse) SetLoad5(v *int64) {
	o.Load5 = v
}

// GetMemoryTotal returns the MemoryTotal field value
func (o *GetMetricsResponse) GetMemoryTotal() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.MemoryTotal
}

// GetMemoryTotalOk returns a tuple with the MemoryTotal field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetMemoryTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryTotal, true
}

// SetMemoryTotal sets field value
func (o *GetMetricsResponse) SetMemoryTotal(v *int64) {
	o.MemoryTotal = v
}

// GetMemoryUsed returns the MemoryUsed field value
func (o *GetMetricsResponse) GetMemoryUsed() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.MemoryUsed
}

// GetMemoryUsedOk returns a tuple with the MemoryUsed field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetMemoryUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryUsed, true
}

// SetMemoryUsed sets field value
func (o *GetMetricsResponse) SetMemoryUsed(v *int64) {
	o.MemoryUsed = v
}

// GetParachuteDiskEphemeralActivated returns the ParachuteDiskEphemeralActivated field value
func (o *GetMetricsResponse) GetParachuteDiskEphemeralActivated() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskEphemeralActivated
}

// GetParachuteDiskEphemeralActivatedOk returns a tuple with the ParachuteDiskEphemeralActivated field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskEphemeralActivatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskEphemeralActivated, true
}

// SetParachuteDiskEphemeralActivated sets field value
func (o *GetMetricsResponse) SetParachuteDiskEphemeralActivated(v *int64) {
	o.ParachuteDiskEphemeralActivated = v
}

// GetParachuteDiskEphemeralTotal returns the ParachuteDiskEphemeralTotal field value
func (o *GetMetricsResponse) GetParachuteDiskEphemeralTotal() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskEphemeralTotal
}

// GetParachuteDiskEphemeralTotalOk returns a tuple with the ParachuteDiskEphemeralTotal field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskEphemeralTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskEphemeralTotal, true
}

// SetParachuteDiskEphemeralTotal sets field value
func (o *GetMetricsResponse) SetParachuteDiskEphemeralTotal(v *int64) {
	o.ParachuteDiskEphemeralTotal = v
}

// GetParachuteDiskEphemeralUsed returns the ParachuteDiskEphemeralUsed field value
func (o *GetMetricsResponse) GetParachuteDiskEphemeralUsed() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskEphemeralUsed
}

// GetParachuteDiskEphemeralUsedOk returns a tuple with the ParachuteDiskEphemeralUsed field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskEphemeralUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskEphemeralUsed, true
}

// SetParachuteDiskEphemeralUsed sets field value
func (o *GetMetricsResponse) SetParachuteDiskEphemeralUsed(v *int64) {
	o.ParachuteDiskEphemeralUsed = v
}

// GetParachuteDiskEphemeralUsedPercent returns the ParachuteDiskEphemeralUsedPercent field value
func (o *GetMetricsResponse) GetParachuteDiskEphemeralUsedPercent() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskEphemeralUsedPercent
}

// GetParachuteDiskEphemeralUsedPercentOk returns a tuple with the ParachuteDiskEphemeralUsedPercent field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskEphemeralUsedPercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskEphemeralUsedPercent, true
}

// SetParachuteDiskEphemeralUsedPercent sets field value
func (o *GetMetricsResponse) SetParachuteDiskEphemeralUsedPercent(v *int64) {
	o.ParachuteDiskEphemeralUsedPercent = v
}

// GetParachuteDiskEphemeralUsedThreshold returns the ParachuteDiskEphemeralUsedThreshold field value
func (o *GetMetricsResponse) GetParachuteDiskEphemeralUsedThreshold() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskEphemeralUsedThreshold
}

// GetParachuteDiskEphemeralUsedThresholdOk returns a tuple with the ParachuteDiskEphemeralUsedThreshold field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskEphemeralUsedThresholdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskEphemeralUsedThreshold, true
}

// SetParachuteDiskEphemeralUsedThreshold sets field value
func (o *GetMetricsResponse) SetParachuteDiskEphemeralUsedThreshold(v *int64) {
	o.ParachuteDiskEphemeralUsedThreshold = v
}

// GetParachuteDiskPersistentActivated returns the ParachuteDiskPersistentActivated field value
func (o *GetMetricsResponse) GetParachuteDiskPersistentActivated() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskPersistentActivated
}

// GetParachuteDiskPersistentActivatedOk returns a tuple with the ParachuteDiskPersistentActivated field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskPersistentActivatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskPersistentActivated, true
}

// SetParachuteDiskPersistentActivated sets field value
func (o *GetMetricsResponse) SetParachuteDiskPersistentActivated(v *int64) {
	o.ParachuteDiskPersistentActivated = v
}

// GetParachuteDiskPersistentTotal returns the ParachuteDiskPersistentTotal field value
func (o *GetMetricsResponse) GetParachuteDiskPersistentTotal() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskPersistentTotal
}

// GetParachuteDiskPersistentTotalOk returns a tuple with the ParachuteDiskPersistentTotal field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskPersistentTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskPersistentTotal, true
}

// SetParachuteDiskPersistentTotal sets field value
func (o *GetMetricsResponse) SetParachuteDiskPersistentTotal(v *int64) {
	o.ParachuteDiskPersistentTotal = v
}

// GetParachuteDiskPersistentUsed returns the ParachuteDiskPersistentUsed field value
func (o *GetMetricsResponse) GetParachuteDiskPersistentUsed() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskPersistentUsed
}

// GetParachuteDiskPersistentUsedOk returns a tuple with the ParachuteDiskPersistentUsed field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskPersistentUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskPersistentUsed, true
}

// SetParachuteDiskPersistentUsed sets field value
func (o *GetMetricsResponse) SetParachuteDiskPersistentUsed(v *int64) {
	o.ParachuteDiskPersistentUsed = v
}

// GetParachuteDiskPersistentUsedPercent returns the ParachuteDiskPersistentUsedPercent field value
func (o *GetMetricsResponse) GetParachuteDiskPersistentUsedPercent() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskPersistentUsedPercent
}

// GetParachuteDiskPersistentUsedPercentOk returns a tuple with the ParachuteDiskPersistentUsedPercent field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskPersistentUsedPercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskPersistentUsedPercent, true
}

// SetParachuteDiskPersistentUsedPercent sets field value
func (o *GetMetricsResponse) SetParachuteDiskPersistentUsedPercent(v *int64) {
	o.ParachuteDiskPersistentUsedPercent = v
}

// GetParachuteDiskPersistentUsedThreshold returns the ParachuteDiskPersistentUsedThreshold field value
func (o *GetMetricsResponse) GetParachuteDiskPersistentUsedThreshold() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ParachuteDiskPersistentUsedThreshold
}

// GetParachuteDiskPersistentUsedThresholdOk returns a tuple with the ParachuteDiskPersistentUsedThreshold field value
// and a boolean to check if the value has been set.
func (o *GetMetricsResponse) GetParachuteDiskPersistentUsedThresholdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParachuteDiskPersistentUsedThreshold, true
}

// SetParachuteDiskPersistentUsedThreshold sets field value
func (o *GetMetricsResponse) SetParachuteDiskPersistentUsedThreshold(v *int64) {
	o.ParachuteDiskPersistentUsedThreshold = v
}

func (o GetMetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuIdleTime) {
		toSerialize["cpuIdleTime"] = o.CpuIdleTime
	}
	toSerialize["cpuLoadPercent"] = o.CpuLoadPercent
	if !IsNil(o.CpuSystemTime) {
		toSerialize["cpuSystemTime"] = o.CpuSystemTime
	}
	if !IsNil(o.CpuUserTime) {
		toSerialize["cpuUserTime"] = o.CpuUserTime
	}
	toSerialize["diskEphemeralTotal"] = o.DiskEphemeralTotal
	toSerialize["diskEphemeralUsed"] = o.DiskEphemeralUsed
	toSerialize["diskPersistentTotal"] = o.DiskPersistentTotal
	toSerialize["diskPersistentUsed"] = o.DiskPersistentUsed
	toSerialize["load1"] = o.Load1
	toSerialize["load15"] = o.Load15
	toSerialize["load5"] = o.Load5
	toSerialize["memoryTotal"] = o.MemoryTotal
	toSerialize["memoryUsed"] = o.MemoryUsed
	toSerialize["parachuteDiskEphemeralActivated"] = o.ParachuteDiskEphemeralActivated
	toSerialize["parachuteDiskEphemeralTotal"] = o.ParachuteDiskEphemeralTotal
	toSerialize["parachuteDiskEphemeralUsed"] = o.ParachuteDiskEphemeralUsed
	toSerialize["parachuteDiskEphemeralUsedPercent"] = o.ParachuteDiskEphemeralUsedPercent
	toSerialize["parachuteDiskEphemeralUsedThreshold"] = o.ParachuteDiskEphemeralUsedThreshold
	toSerialize["parachuteDiskPersistentActivated"] = o.ParachuteDiskPersistentActivated
	toSerialize["parachuteDiskPersistentTotal"] = o.ParachuteDiskPersistentTotal
	toSerialize["parachuteDiskPersistentUsed"] = o.ParachuteDiskPersistentUsed
	toSerialize["parachuteDiskPersistentUsedPercent"] = o.ParachuteDiskPersistentUsedPercent
	toSerialize["parachuteDiskPersistentUsedThreshold"] = o.ParachuteDiskPersistentUsedThreshold
	return toSerialize, nil
}

type NullableGetMetricsResponse struct {
	value *GetMetricsResponse
	isSet bool
}

func (v NullableGetMetricsResponse) Get() *GetMetricsResponse {
	return v.value
}

func (v *NullableGetMetricsResponse) Set(val *GetMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetricsResponse(val *GetMetricsResponse) *NullableGetMetricsResponse {
	return &NullableGetMetricsResponse{value: val, isSet: true}
}

func (v NullableGetMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
