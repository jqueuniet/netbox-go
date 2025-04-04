/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_6

import (
	"encoding/json"
	"fmt"
)

// checks if the WritableConfigContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableConfigContextRequest{}

// WritableConfigContextRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableConfigContextRequest struct {
	Name          string   `json:"name"`
	Weight        *int32   `json:"weight,omitempty"`
	Description   *string  `json:"description,omitempty"`
	IsActive      *bool    `json:"is_active,omitempty"`
	Regions       []int32  `json:"regions,omitempty"`
	SiteGroups    []int32  `json:"site_groups,omitempty"`
	Sites         []int32  `json:"sites,omitempty"`
	Locations     []int32  `json:"locations,omitempty"`
	DeviceTypes   []int32  `json:"device_types,omitempty"`
	Roles         []int32  `json:"roles,omitempty"`
	Platforms     []int32  `json:"platforms,omitempty"`
	ClusterTypes  []int32  `json:"cluster_types,omitempty"`
	ClusterGroups []int32  `json:"cluster_groups,omitempty"`
	Clusters      []int32  `json:"clusters,omitempty"`
	TenantGroups  []int32  `json:"tenant_groups,omitempty"`
	Tenants       []int32  `json:"tenants,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	// Remote data source
	DataSource           NullableInt32 `json:"data_source,omitempty"`
	Data                 interface{}   `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _WritableConfigContextRequest WritableConfigContextRequest

// NewWritableConfigContextRequest instantiates a new WritableConfigContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableConfigContextRequest(name string, data interface{}) *WritableConfigContextRequest {
	this := WritableConfigContextRequest{}
	this.Name = name
	this.Data = data
	return &this
}

// NewWritableConfigContextRequestWithDefaults instantiates a new WritableConfigContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableConfigContextRequestWithDefaults() *WritableConfigContextRequest {
	this := WritableConfigContextRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableConfigContextRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableConfigContextRequest) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *WritableConfigContextRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableConfigContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *WritableConfigContextRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetRegions() []int32 {
	if o == nil || IsNil(o.Regions) {
		var ret []int32
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetRegionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []int32 and assigns it to the Regions field.
func (o *WritableConfigContextRequest) SetRegions(v []int32) {
	o.Regions = v
}

// GetSiteGroups returns the SiteGroups field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetSiteGroups() []int32 {
	if o == nil || IsNil(o.SiteGroups) {
		var ret []int32
		return ret
	}
	return o.SiteGroups
}

// GetSiteGroupsOk returns a tuple with the SiteGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetSiteGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SiteGroups) {
		return nil, false
	}
	return o.SiteGroups, true
}

// HasSiteGroups returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasSiteGroups() bool {
	if o != nil && !IsNil(o.SiteGroups) {
		return true
	}

	return false
}

// SetSiteGroups gets a reference to the given []int32 and assigns it to the SiteGroups field.
func (o *WritableConfigContextRequest) SetSiteGroups(v []int32) {
	o.SiteGroups = v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetSites() []int32 {
	if o == nil || IsNil(o.Sites) {
		var ret []int32
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetSitesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []int32 and assigns it to the Sites field.
func (o *WritableConfigContextRequest) SetSites(v []int32) {
	o.Sites = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetLocations() []int32 {
	if o == nil || IsNil(o.Locations) {
		var ret []int32
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetLocationsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []int32 and assigns it to the Locations field.
func (o *WritableConfigContextRequest) SetLocations(v []int32) {
	o.Locations = v
}

// GetDeviceTypes returns the DeviceTypes field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetDeviceTypes() []int32 {
	if o == nil || IsNil(o.DeviceTypes) {
		var ret []int32
		return ret
	}
	return o.DeviceTypes
}

// GetDeviceTypesOk returns a tuple with the DeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetDeviceTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.DeviceTypes) {
		return nil, false
	}
	return o.DeviceTypes, true
}

// HasDeviceTypes returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasDeviceTypes() bool {
	if o != nil && !IsNil(o.DeviceTypes) {
		return true
	}

	return false
}

// SetDeviceTypes gets a reference to the given []int32 and assigns it to the DeviceTypes field.
func (o *WritableConfigContextRequest) SetDeviceTypes(v []int32) {
	o.DeviceTypes = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetRoles() []int32 {
	if o == nil || IsNil(o.Roles) {
		var ret []int32
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetRolesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int32 and assigns it to the Roles field.
func (o *WritableConfigContextRequest) SetRoles(v []int32) {
	o.Roles = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetPlatforms() []int32 {
	if o == nil || IsNil(o.Platforms) {
		var ret []int32
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetPlatformsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []int32 and assigns it to the Platforms field.
func (o *WritableConfigContextRequest) SetPlatforms(v []int32) {
	o.Platforms = v
}

// GetClusterTypes returns the ClusterTypes field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetClusterTypes() []int32 {
	if o == nil || IsNil(o.ClusterTypes) {
		var ret []int32
		return ret
	}
	return o.ClusterTypes
}

// GetClusterTypesOk returns a tuple with the ClusterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetClusterTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.ClusterTypes) {
		return nil, false
	}
	return o.ClusterTypes, true
}

// HasClusterTypes returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasClusterTypes() bool {
	if o != nil && !IsNil(o.ClusterTypes) {
		return true
	}

	return false
}

// SetClusterTypes gets a reference to the given []int32 and assigns it to the ClusterTypes field.
func (o *WritableConfigContextRequest) SetClusterTypes(v []int32) {
	o.ClusterTypes = v
}

// GetClusterGroups returns the ClusterGroups field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetClusterGroups() []int32 {
	if o == nil || IsNil(o.ClusterGroups) {
		var ret []int32
		return ret
	}
	return o.ClusterGroups
}

// GetClusterGroupsOk returns a tuple with the ClusterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetClusterGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ClusterGroups) {
		return nil, false
	}
	return o.ClusterGroups, true
}

// HasClusterGroups returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasClusterGroups() bool {
	if o != nil && !IsNil(o.ClusterGroups) {
		return true
	}

	return false
}

// SetClusterGroups gets a reference to the given []int32 and assigns it to the ClusterGroups field.
func (o *WritableConfigContextRequest) SetClusterGroups(v []int32) {
	o.ClusterGroups = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetClusters() []int32 {
	if o == nil || IsNil(o.Clusters) {
		var ret []int32
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetClustersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []int32 and assigns it to the Clusters field.
func (o *WritableConfigContextRequest) SetClusters(v []int32) {
	o.Clusters = v
}

// GetTenantGroups returns the TenantGroups field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetTenantGroups() []int32 {
	if o == nil || IsNil(o.TenantGroups) {
		var ret []int32
		return ret
	}
	return o.TenantGroups
}

// GetTenantGroupsOk returns a tuple with the TenantGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetTenantGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TenantGroups) {
		return nil, false
	}
	return o.TenantGroups, true
}

// HasTenantGroups returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasTenantGroups() bool {
	if o != nil && !IsNil(o.TenantGroups) {
		return true
	}

	return false
}

// SetTenantGroups gets a reference to the given []int32 and assigns it to the TenantGroups field.
func (o *WritableConfigContextRequest) SetTenantGroups(v []int32) {
	o.TenantGroups = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetTenants() []int32 {
	if o == nil || IsNil(o.Tenants) {
		var ret []int32
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetTenantsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []int32 and assigns it to the Tenants field.
func (o *WritableConfigContextRequest) SetTenants(v []int32) {
	o.Tenants = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableConfigContextRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigContextRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WritableConfigContextRequest) SetTags(v []string) {
	o.Tags = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConfigContextRequest) GetDataSource() int32 {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret int32
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConfigContextRequest) GetDataSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *WritableConfigContextRequest) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableInt32 and assigns it to the DataSource field.
func (o *WritableConfigContextRequest) SetDataSource(v int32) {
	o.DataSource.Set(&v)
}

// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *WritableConfigContextRequest) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *WritableConfigContextRequest) UnsetDataSource() {
	o.DataSource.Unset()
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WritableConfigContextRequest) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConfigContextRequest) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WritableConfigContextRequest) SetData(v interface{}) {
	o.Data = v
}

func (o WritableConfigContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableConfigContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.SiteGroups) {
		toSerialize["site_groups"] = o.SiteGroups
	}
	if !IsNil(o.Sites) {
		toSerialize["sites"] = o.Sites
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.DeviceTypes) {
		toSerialize["device_types"] = o.DeviceTypes
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	if !IsNil(o.ClusterTypes) {
		toSerialize["cluster_types"] = o.ClusterTypes
	}
	if !IsNil(o.ClusterGroups) {
		toSerialize["cluster_groups"] = o.ClusterGroups
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.TenantGroups) {
		toSerialize["tenant_groups"] = o.TenantGroups
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.DataSource.IsSet() {
		toSerialize["data_source"] = o.DataSource.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableConfigContextRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"data",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varWritableConfigContextRequest := _WritableConfigContextRequest{}

	err = json.Unmarshal(data, &varWritableConfigContextRequest)

	if err != nil {
		return err
	}

	*o = WritableConfigContextRequest(varWritableConfigContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "description")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "regions")
		delete(additionalProperties, "site_groups")
		delete(additionalProperties, "sites")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "device_types")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "platforms")
		delete(additionalProperties, "cluster_types")
		delete(additionalProperties, "cluster_groups")
		delete(additionalProperties, "clusters")
		delete(additionalProperties, "tenant_groups")
		delete(additionalProperties, "tenants")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableConfigContextRequest struct {
	value *WritableConfigContextRequest
	isSet bool
}

func (v NullableWritableConfigContextRequest) Get() *WritableConfigContextRequest {
	return v.value
}

func (v *NullableWritableConfigContextRequest) Set(val *WritableConfigContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableConfigContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableConfigContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableConfigContextRequest(val *WritableConfigContextRequest) *NullableWritableConfigContextRequest {
	return &NullableWritableConfigContextRequest{value: val, isSet: true}
}

func (v NullableWritableConfigContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableConfigContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
