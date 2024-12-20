/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the QueryOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryOutput{}

// QueryOutput Represents a query that can be sent to another system.
type QueryOutput struct {
	Name            string                           `json:"name"`
	Type            *string                          `json:"type,omitempty"`
	Description     *string                          `json:"description,omitempty"`
	Created         *time.Time                       `json:"created,omitempty"`
	Modified        *time.Time                       `json:"modified,omitempty"`
	ValidFrom       *time.Time                       `json:"valid_from,omitempty"`
	ValidUntil      *time.Time                       `json:"valid_until,omitempty"`
	Pattern         string                           `json:"pattern"`
	Location        *string                          `json:"location,omitempty"`
	Diamond         DiamondModel                     `json:"diamond"`
	KillChainPhases []string                         `json:"kill_chain_phases,omitempty"`
	RelevantTags    []string                         `json:"relevant_tags,omitempty"`
	QueryType       string                           `json:"query_type"`
	TargetSystems   []string                         `json:"target_systems,omitempty"`
	Id              string                           `json:"id"`
	Tags            map[string]TagRelationshipOutput `json:"tags"`
	RootType        string                           `json:"root_type"`
}

type _QueryOutput QueryOutput

// NewQueryOutput instantiates a new QueryOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOutput(name string, pattern string, diamond DiamondModel, queryType string, id string, tags map[string]TagRelationshipOutput, rootType string) *QueryOutput {
	this := QueryOutput{}
	this.Name = name
	var type_ string = "query"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	this.Pattern = pattern
	var location string = ""
	this.Location = &location
	this.Diamond = diamond
	this.QueryType = queryType
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	return &this
}

// NewQueryOutputWithDefaults instantiates a new QueryOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOutputWithDefaults() *QueryOutput {
	this := QueryOutput{}
	var type_ string = "query"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	var location string = ""
	this.Location = &location
	return &this
}

// GetName returns the Name field value
func (o *QueryOutput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QueryOutput) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryOutput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryOutput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryOutput) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QueryOutput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QueryOutput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QueryOutput) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *QueryOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *QueryOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *QueryOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *QueryOutput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *QueryOutput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *QueryOutput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *QueryOutput) GetValidFrom() time.Time {
	if o == nil || IsNil(o.ValidFrom) {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetValidFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *QueryOutput) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *QueryOutput) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *QueryOutput) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *QueryOutput) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *QueryOutput) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetPattern returns the Pattern field value
func (o *QueryOutput) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *QueryOutput) SetPattern(v string) {
	o.Pattern = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *QueryOutput) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *QueryOutput) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *QueryOutput) SetLocation(v string) {
	o.Location = &v
}

// GetDiamond returns the Diamond field value
func (o *QueryOutput) GetDiamond() DiamondModel {
	if o == nil {
		var ret DiamondModel
		return ret
	}

	return o.Diamond
}

// GetDiamondOk returns a tuple with the Diamond field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetDiamondOk() (*DiamondModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diamond, true
}

// SetDiamond sets field value
func (o *QueryOutput) SetDiamond(v DiamondModel) {
	o.Diamond = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *QueryOutput) GetKillChainPhases() []string {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []string
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetKillChainPhasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *QueryOutput) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []string and assigns it to the KillChainPhases field.
func (o *QueryOutput) SetKillChainPhases(v []string) {
	o.KillChainPhases = v
}

// GetRelevantTags returns the RelevantTags field value if set, zero value otherwise.
func (o *QueryOutput) GetRelevantTags() []string {
	if o == nil || IsNil(o.RelevantTags) {
		var ret []string
		return ret
	}
	return o.RelevantTags
}

// GetRelevantTagsOk returns a tuple with the RelevantTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetRelevantTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelevantTags) {
		return nil, false
	}
	return o.RelevantTags, true
}

// HasRelevantTags returns a boolean if a field has been set.
func (o *QueryOutput) HasRelevantTags() bool {
	if o != nil && !IsNil(o.RelevantTags) {
		return true
	}

	return false
}

// SetRelevantTags gets a reference to the given []string and assigns it to the RelevantTags field.
func (o *QueryOutput) SetRelevantTags(v []string) {
	o.RelevantTags = v
}

// GetQueryType returns the QueryType field value
func (o *QueryOutput) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *QueryOutput) SetQueryType(v string) {
	o.QueryType = v
}

// GetTargetSystems returns the TargetSystems field value if set, zero value otherwise.
func (o *QueryOutput) GetTargetSystems() []string {
	if o == nil || IsNil(o.TargetSystems) {
		var ret []string
		return ret
	}
	return o.TargetSystems
}

// GetTargetSystemsOk returns a tuple with the TargetSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetTargetSystemsOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetSystems) {
		return nil, false
	}
	return o.TargetSystems, true
}

// HasTargetSystems returns a boolean if a field has been set.
func (o *QueryOutput) HasTargetSystems() bool {
	if o != nil && !IsNil(o.TargetSystems) {
		return true
	}

	return false
}

// SetTargetSystems gets a reference to the given []string and assigns it to the TargetSystems field.
func (o *QueryOutput) SetTargetSystems(v []string) {
	o.TargetSystems = v
}

// GetId returns the Id field value
func (o *QueryOutput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QueryOutput) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *QueryOutput) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *QueryOutput) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *QueryOutput) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *QueryOutput) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *QueryOutput) SetRootType(v string) {
	o.RootType = v
}

func (o QueryOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if !IsNil(o.ValidUntil) {
		toSerialize["valid_until"] = o.ValidUntil
	}
	toSerialize["pattern"] = o.Pattern
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["diamond"] = o.Diamond
	if !IsNil(o.KillChainPhases) {
		toSerialize["kill_chain_phases"] = o.KillChainPhases
	}
	if !IsNil(o.RelevantTags) {
		toSerialize["relevant_tags"] = o.RelevantTags
	}
	toSerialize["query_type"] = o.QueryType
	if !IsNil(o.TargetSystems) {
		toSerialize["target_systems"] = o.TargetSystems
	}
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["root_type"] = o.RootType
	return toSerialize, nil
}

func (o *QueryOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"pattern",
		"diamond",
		"query_type",
		"id",
		"tags",
		"root_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQueryOutput := _QueryOutput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryOutput)

	if err != nil {
		return err
	}

	*o = QueryOutput(varQueryOutput)

	return err
}

type NullableQueryOutput struct {
	value *QueryOutput
	isSet bool
}

func (v NullableQueryOutput) Get() *QueryOutput {
	return v.value
}

func (v *NullableQueryOutput) Set(val *QueryOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOutput(val *QueryOutput) *NullableQueryOutput {
	return &NullableQueryOutput{value: val, isSet: true}
}

func (v NullableQueryOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
