/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CircuitCircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitCircuitTerminationRequest{}

// CircuitCircuitTerminationRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type CircuitCircuitTerminationRequest struct {
	Site NullableBriefSiteRequest `json:"site"`
	ProviderNetwork NullableBriefProviderNetworkRequest `json:"provider_network"`
	// Physical circuit speed
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId *string `json:"xconnect_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

type _CircuitCircuitTerminationRequest CircuitCircuitTerminationRequest

// NewCircuitCircuitTerminationRequest instantiates a new CircuitCircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitCircuitTerminationRequest(site NullableBriefSiteRequest, providerNetwork NullableBriefProviderNetworkRequest) *CircuitCircuitTerminationRequest {
	this := CircuitCircuitTerminationRequest{}
	this.Site = site
	this.ProviderNetwork = providerNetwork
	return &this
}

// NewCircuitCircuitTerminationRequestWithDefaults instantiates a new CircuitCircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitCircuitTerminationRequestWithDefaults() *CircuitCircuitTerminationRequest {
	this := CircuitCircuitTerminationRequest{}
	return &this
}

// GetSite returns the Site field value
// If the value is explicit nil, the zero value for BriefSiteRequest will be returned
func (o *CircuitCircuitTerminationRequest) GetSite() BriefSiteRequest {
	if o == nil || o.Site.Get() == nil {
		var ret BriefSiteRequest
		return ret
	}

	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// SetSite sets field value
func (o *CircuitCircuitTerminationRequest) SetSite(v BriefSiteRequest) {
	o.Site.Set(&v)
}

// GetProviderNetwork returns the ProviderNetwork field value
// If the value is explicit nil, the zero value for BriefProviderNetworkRequest will be returned
func (o *CircuitCircuitTerminationRequest) GetProviderNetwork() BriefProviderNetworkRequest {
	if o == nil || o.ProviderNetwork.Get() == nil {
		var ret BriefProviderNetworkRequest
		return ret
	}

	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// SetProviderNetwork sets field value
func (o *CircuitCircuitTerminationRequest) SetProviderNetwork(v BriefProviderNetworkRequest) {
	o.ProviderNetwork.Set(&v)
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTerminationRequest) GetPortSpeed() int32 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *CircuitCircuitTerminationRequest) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}
// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *CircuitCircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *CircuitCircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTerminationRequest) GetUpstreamSpeed() int32 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *CircuitCircuitTerminationRequest) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}
// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *CircuitCircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *CircuitCircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *CircuitCircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CircuitCircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CircuitCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitCircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["site"] = o.Site.Get()
	toSerialize["provider_network"] = o.ProviderNetwork.Get()
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *CircuitCircuitTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"site",
		"provider_network",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCircuitCircuitTerminationRequest := _CircuitCircuitTerminationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCircuitCircuitTerminationRequest)

	if err != nil {
		return err
	}

	*o = CircuitCircuitTerminationRequest(varCircuitCircuitTerminationRequest)

	return err
}

type NullableCircuitCircuitTerminationRequest struct {
	value *CircuitCircuitTerminationRequest
	isSet bool
}

func (v NullableCircuitCircuitTerminationRequest) Get() *CircuitCircuitTerminationRequest {
	return v.value
}

func (v *NullableCircuitCircuitTerminationRequest) Set(val *CircuitCircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitCircuitTerminationRequest(val *CircuitCircuitTerminationRequest) *NullableCircuitCircuitTerminationRequest {
	return &NullableCircuitCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullableCircuitCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

