/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ConfigContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigContext{}

// ConfigContext Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ConfigContext struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Weight *int32 `json:"weight,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Regions []Region `json:"regions,omitempty"`
	SiteGroups []SiteGroup `json:"site_groups,omitempty"`
	Sites []Site `json:"sites,omitempty"`
	Locations []Location `json:"locations,omitempty"`
	DeviceTypes []DeviceType `json:"device_types,omitempty"`
	Roles []DeviceRole `json:"roles,omitempty"`
	Platforms []Platform `json:"platforms,omitempty"`
	ClusterTypes []ClusterType `json:"cluster_types,omitempty"`
	ClusterGroups []ClusterGroup `json:"cluster_groups,omitempty"`
	Clusters []Cluster `json:"clusters,omitempty"`
	TenantGroups []TenantGroup `json:"tenant_groups,omitempty"`
	Tenants []Tenant `json:"tenants,omitempty"`
	Tags []string `json:"tags,omitempty"`
	DataSource *BriefDataSource `json:"data_source,omitempty"`
	// Path to remote file (relative to data source root)
	DataPath string `json:"data_path"`
	DataFile BriefDataFile `json:"data_file"`
	DataSynced NullableTime `json:"data_synced"`
	Data interface{} `json:"data"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
}

type _ConfigContext ConfigContext

// NewConfigContext instantiates a new ConfigContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigContext(id int32, url string, display string, name string, dataPath string, dataFile BriefDataFile, dataSynced NullableTime, data interface{}, created NullableTime, lastUpdated NullableTime) *ConfigContext {
	this := ConfigContext{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.DataPath = dataPath
	this.DataFile = dataFile
	this.DataSynced = dataSynced
	this.Data = data
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewConfigContextWithDefaults instantiates a new ConfigContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigContextWithDefaults() *ConfigContext {
	this := ConfigContext{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigContext) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigContext) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConfigContext) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConfigContext) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ConfigContext) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConfigContext) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *ConfigContext) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigContext) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ConfigContext) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ConfigContext) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ConfigContext) SetWeight(v int32) {
	o.Weight = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigContext) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigContext) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigContext) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ConfigContext) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ConfigContext) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ConfigContext) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *ConfigContext) GetRegions() []Region {
	if o == nil || IsNil(o.Regions) {
		var ret []Region
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetRegionsOk() ([]Region, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *ConfigContext) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []Region and assigns it to the Regions field.
func (o *ConfigContext) SetRegions(v []Region) {
	o.Regions = v
}

// GetSiteGroups returns the SiteGroups field value if set, zero value otherwise.
func (o *ConfigContext) GetSiteGroups() []SiteGroup {
	if o == nil || IsNil(o.SiteGroups) {
		var ret []SiteGroup
		return ret
	}
	return o.SiteGroups
}

// GetSiteGroupsOk returns a tuple with the SiteGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetSiteGroupsOk() ([]SiteGroup, bool) {
	if o == nil || IsNil(o.SiteGroups) {
		return nil, false
	}
	return o.SiteGroups, true
}

// HasSiteGroups returns a boolean if a field has been set.
func (o *ConfigContext) HasSiteGroups() bool {
	if o != nil && !IsNil(o.SiteGroups) {
		return true
	}

	return false
}

// SetSiteGroups gets a reference to the given []SiteGroup and assigns it to the SiteGroups field.
func (o *ConfigContext) SetSiteGroups(v []SiteGroup) {
	o.SiteGroups = v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ConfigContext) GetSites() []Site {
	if o == nil || IsNil(o.Sites) {
		var ret []Site
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetSitesOk() ([]Site, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *ConfigContext) HasSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []Site and assigns it to the Sites field.
func (o *ConfigContext) SetSites(v []Site) {
	o.Sites = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ConfigContext) GetLocations() []Location {
	if o == nil || IsNil(o.Locations) {
		var ret []Location
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetLocationsOk() ([]Location, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ConfigContext) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []Location and assigns it to the Locations field.
func (o *ConfigContext) SetLocations(v []Location) {
	o.Locations = v
}

// GetDeviceTypes returns the DeviceTypes field value if set, zero value otherwise.
func (o *ConfigContext) GetDeviceTypes() []DeviceType {
	if o == nil || IsNil(o.DeviceTypes) {
		var ret []DeviceType
		return ret
	}
	return o.DeviceTypes
}

// GetDeviceTypesOk returns a tuple with the DeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetDeviceTypesOk() ([]DeviceType, bool) {
	if o == nil || IsNil(o.DeviceTypes) {
		return nil, false
	}
	return o.DeviceTypes, true
}

// HasDeviceTypes returns a boolean if a field has been set.
func (o *ConfigContext) HasDeviceTypes() bool {
	if o != nil && !IsNil(o.DeviceTypes) {
		return true
	}

	return false
}

// SetDeviceTypes gets a reference to the given []DeviceType and assigns it to the DeviceTypes field.
func (o *ConfigContext) SetDeviceTypes(v []DeviceType) {
	o.DeviceTypes = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ConfigContext) GetRoles() []DeviceRole {
	if o == nil || IsNil(o.Roles) {
		var ret []DeviceRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetRolesOk() ([]DeviceRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ConfigContext) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []DeviceRole and assigns it to the Roles field.
func (o *ConfigContext) SetRoles(v []DeviceRole) {
	o.Roles = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *ConfigContext) GetPlatforms() []Platform {
	if o == nil || IsNil(o.Platforms) {
		var ret []Platform
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetPlatformsOk() ([]Platform, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *ConfigContext) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []Platform and assigns it to the Platforms field.
func (o *ConfigContext) SetPlatforms(v []Platform) {
	o.Platforms = v
}

// GetClusterTypes returns the ClusterTypes field value if set, zero value otherwise.
func (o *ConfigContext) GetClusterTypes() []ClusterType {
	if o == nil || IsNil(o.ClusterTypes) {
		var ret []ClusterType
		return ret
	}
	return o.ClusterTypes
}

// GetClusterTypesOk returns a tuple with the ClusterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetClusterTypesOk() ([]ClusterType, bool) {
	if o == nil || IsNil(o.ClusterTypes) {
		return nil, false
	}
	return o.ClusterTypes, true
}

// HasClusterTypes returns a boolean if a field has been set.
func (o *ConfigContext) HasClusterTypes() bool {
	if o != nil && !IsNil(o.ClusterTypes) {
		return true
	}

	return false
}

// SetClusterTypes gets a reference to the given []ClusterType and assigns it to the ClusterTypes field.
func (o *ConfigContext) SetClusterTypes(v []ClusterType) {
	o.ClusterTypes = v
}

// GetClusterGroups returns the ClusterGroups field value if set, zero value otherwise.
func (o *ConfigContext) GetClusterGroups() []ClusterGroup {
	if o == nil || IsNil(o.ClusterGroups) {
		var ret []ClusterGroup
		return ret
	}
	return o.ClusterGroups
}

// GetClusterGroupsOk returns a tuple with the ClusterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetClusterGroupsOk() ([]ClusterGroup, bool) {
	if o == nil || IsNil(o.ClusterGroups) {
		return nil, false
	}
	return o.ClusterGroups, true
}

// HasClusterGroups returns a boolean if a field has been set.
func (o *ConfigContext) HasClusterGroups() bool {
	if o != nil && !IsNil(o.ClusterGroups) {
		return true
	}

	return false
}

// SetClusterGroups gets a reference to the given []ClusterGroup and assigns it to the ClusterGroups field.
func (o *ConfigContext) SetClusterGroups(v []ClusterGroup) {
	o.ClusterGroups = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *ConfigContext) GetClusters() []Cluster {
	if o == nil || IsNil(o.Clusters) {
		var ret []Cluster
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetClustersOk() ([]Cluster, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *ConfigContext) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []Cluster and assigns it to the Clusters field.
func (o *ConfigContext) SetClusters(v []Cluster) {
	o.Clusters = v
}

// GetTenantGroups returns the TenantGroups field value if set, zero value otherwise.
func (o *ConfigContext) GetTenantGroups() []TenantGroup {
	if o == nil || IsNil(o.TenantGroups) {
		var ret []TenantGroup
		return ret
	}
	return o.TenantGroups
}

// GetTenantGroupsOk returns a tuple with the TenantGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetTenantGroupsOk() ([]TenantGroup, bool) {
	if o == nil || IsNil(o.TenantGroups) {
		return nil, false
	}
	return o.TenantGroups, true
}

// HasTenantGroups returns a boolean if a field has been set.
func (o *ConfigContext) HasTenantGroups() bool {
	if o != nil && !IsNil(o.TenantGroups) {
		return true
	}

	return false
}

// SetTenantGroups gets a reference to the given []TenantGroup and assigns it to the TenantGroups field.
func (o *ConfigContext) SetTenantGroups(v []TenantGroup) {
	o.TenantGroups = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ConfigContext) GetTenants() []Tenant {
	if o == nil || IsNil(o.Tenants) {
		var ret []Tenant
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetTenantsOk() ([]Tenant, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ConfigContext) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []Tenant and assigns it to the Tenants field.
func (o *ConfigContext) SetTenants(v []Tenant) {
	o.Tenants = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigContext) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigContext) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConfigContext) SetTags(v []string) {
	o.Tags = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ConfigContext) GetDataSource() BriefDataSource {
	if o == nil || IsNil(o.DataSource) {
		var ret BriefDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetDataSourceOk() (*BriefDataSource, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ConfigContext) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given BriefDataSource and assigns it to the DataSource field.
func (o *ConfigContext) SetDataSource(v BriefDataSource) {
	o.DataSource = &v
}

// GetDataPath returns the DataPath field value
func (o *ConfigContext) GetDataPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataPath
}

// GetDataPathOk returns a tuple with the DataPath field value
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetDataPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataPath, true
}

// SetDataPath sets field value
func (o *ConfigContext) SetDataPath(v string) {
	o.DataPath = v
}

// GetDataFile returns the DataFile field value
func (o *ConfigContext) GetDataFile() BriefDataFile {
	if o == nil {
		var ret BriefDataFile
		return ret
	}

	return o.DataFile
}

// GetDataFileOk returns a tuple with the DataFile field value
// and a boolean to check if the value has been set.
func (o *ConfigContext) GetDataFileOk() (*BriefDataFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataFile, true
}

// SetDataFile sets field value
func (o *ConfigContext) SetDataFile(v BriefDataFile) {
	o.DataFile = v
}

// GetDataSynced returns the DataSynced field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConfigContext) GetDataSynced() time.Time {
	if o == nil || o.DataSynced.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DataSynced.Get()
}

// GetDataSyncedOk returns a tuple with the DataSynced field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContext) GetDataSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSynced.Get(), o.DataSynced.IsSet()
}

// SetDataSynced sets field value
func (o *ConfigContext) SetDataSynced(v time.Time) {
	o.DataSynced.Set(&v)
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConfigContext) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContext) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConfigContext) SetData(v interface{}) {
	o.Data = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConfigContext) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContext) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ConfigContext) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConfigContext) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContext) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ConfigContext) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ConfigContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
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
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}
	toSerialize["data_path"] = o.DataPath
	toSerialize["data_file"] = o.DataFile
	toSerialize["data_synced"] = o.DataSynced.Get()
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	return toSerialize, nil
}

func (o *ConfigContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"data_path",
		"data_file",
		"data_synced",
		"data",
		"created",
		"last_updated",
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

	varConfigContext := _ConfigContext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigContext)

	if err != nil {
		return err
	}

	*o = ConfigContext(varConfigContext)

	return err
}

type NullableConfigContext struct {
	value *ConfigContext
	isSet bool
}

func (v NullableConfigContext) Get() *ConfigContext {
	return v.value
}

func (v *NullableConfigContext) Set(val *ConfigContext) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigContext) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigContext(val *ConfigContext) *NullableConfigContext {
	return &NullableConfigContext{value: val, isSet: true}
}

func (v NullableConfigContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

