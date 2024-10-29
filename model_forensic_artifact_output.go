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

// checks if the ForensicArtifactOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForensicArtifactOutput{}

// ForensicArtifactOutput Represents a Forensic Artifact  As defined in https://github.com/ForensicArtifacts/artifacts
type ForensicArtifactOutput struct {
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
	Sources         []map[string]interface{}         `json:"sources,omitempty"`
	Aliases         []string                         `json:"aliases,omitempty"`
	SupportedOs     []string                         `json:"supported_os,omitempty"`
	Id              string                           `json:"id"`
	Tags            map[string]TagRelationshipOutput `json:"tags"`
	RootType        string                           `json:"root_type"`
}

type _ForensicArtifactOutput ForensicArtifactOutput

// NewForensicArtifactOutput instantiates a new ForensicArtifactOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForensicArtifactOutput(name string, pattern string, diamond DiamondModel, id string, tags map[string]TagRelationshipOutput, rootType string) *ForensicArtifactOutput {
	this := ForensicArtifactOutput{}
	this.Name = name
	var type_ string = "forensicartifact"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	this.Pattern = pattern
	var location string = ""
	this.Location = &location
	this.Diamond = diamond
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	return &this
}

// NewForensicArtifactOutputWithDefaults instantiates a new ForensicArtifactOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForensicArtifactOutputWithDefaults() *ForensicArtifactOutput {
	this := ForensicArtifactOutput{}
	var type_ string = "forensicartifact"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	var location string = ""
	this.Location = &location
	return &this
}

// GetName returns the Name field value
func (o *ForensicArtifactOutput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ForensicArtifactOutput) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ForensicArtifactOutput) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ForensicArtifactOutput) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ForensicArtifactOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ForensicArtifactOutput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetValidFrom() time.Time {
	if o == nil || IsNil(o.ValidFrom) {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetValidFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *ForensicArtifactOutput) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *ForensicArtifactOutput) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetPattern returns the Pattern field value
func (o *ForensicArtifactOutput) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *ForensicArtifactOutput) SetPattern(v string) {
	o.Pattern = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ForensicArtifactOutput) SetLocation(v string) {
	o.Location = &v
}

// GetDiamond returns the Diamond field value
func (o *ForensicArtifactOutput) GetDiamond() DiamondModel {
	if o == nil {
		var ret DiamondModel
		return ret
	}

	return o.Diamond
}

// GetDiamondOk returns a tuple with the Diamond field value
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetDiamondOk() (*DiamondModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diamond, true
}

// SetDiamond sets field value
func (o *ForensicArtifactOutput) SetDiamond(v DiamondModel) {
	o.Diamond = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetKillChainPhases() []string {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []string
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetKillChainPhasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []string and assigns it to the KillChainPhases field.
func (o *ForensicArtifactOutput) SetKillChainPhases(v []string) {
	o.KillChainPhases = v
}

// GetRelevantTags returns the RelevantTags field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetRelevantTags() []string {
	if o == nil || IsNil(o.RelevantTags) {
		var ret []string
		return ret
	}
	return o.RelevantTags
}

// GetRelevantTagsOk returns a tuple with the RelevantTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetRelevantTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelevantTags) {
		return nil, false
	}
	return o.RelevantTags, true
}

// HasRelevantTags returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasRelevantTags() bool {
	if o != nil && !IsNil(o.RelevantTags) {
		return true
	}

	return false
}

// SetRelevantTags gets a reference to the given []string and assigns it to the RelevantTags field.
func (o *ForensicArtifactOutput) SetRelevantTags(v []string) {
	o.RelevantTags = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetSources() []map[string]interface{} {
	if o == nil || IsNil(o.Sources) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetSourcesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []map[string]interface{} and assigns it to the Sources field.
func (o *ForensicArtifactOutput) SetSources(v []map[string]interface{}) {
	o.Sources = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *ForensicArtifactOutput) SetAliases(v []string) {
	o.Aliases = v
}

// GetSupportedOs returns the SupportedOs field value if set, zero value otherwise.
func (o *ForensicArtifactOutput) GetSupportedOs() []string {
	if o == nil || IsNil(o.SupportedOs) {
		var ret []string
		return ret
	}
	return o.SupportedOs
}

// GetSupportedOsOk returns a tuple with the SupportedOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetSupportedOsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedOs) {
		return nil, false
	}
	return o.SupportedOs, true
}

// HasSupportedOs returns a boolean if a field has been set.
func (o *ForensicArtifactOutput) HasSupportedOs() bool {
	if o != nil && !IsNil(o.SupportedOs) {
		return true
	}

	return false
}

// SetSupportedOs gets a reference to the given []string and assigns it to the SupportedOs field.
func (o *ForensicArtifactOutput) SetSupportedOs(v []string) {
	o.SupportedOs = v
}

// GetId returns the Id field value
func (o *ForensicArtifactOutput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ForensicArtifactOutput) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *ForensicArtifactOutput) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ForensicArtifactOutput) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *ForensicArtifactOutput) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *ForensicArtifactOutput) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *ForensicArtifactOutput) SetRootType(v string) {
	o.RootType = v
}

func (o ForensicArtifactOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForensicArtifactOutput) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.SupportedOs) {
		toSerialize["supported_os"] = o.SupportedOs
	}
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["root_type"] = o.RootType
	return toSerialize, nil
}

func (o *ForensicArtifactOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"pattern",
		"diamond",
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

	varForensicArtifactOutput := _ForensicArtifactOutput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForensicArtifactOutput)

	if err != nil {
		return err
	}

	*o = ForensicArtifactOutput(varForensicArtifactOutput)

	return err
}

type NullableForensicArtifactOutput struct {
	value *ForensicArtifactOutput
	isSet bool
}

func (v NullableForensicArtifactOutput) Get() *ForensicArtifactOutput {
	return v.value
}

func (v *NullableForensicArtifactOutput) Set(val *ForensicArtifactOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableForensicArtifactOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableForensicArtifactOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForensicArtifactOutput(val *ForensicArtifactOutput) *NullableForensicArtifactOutput {
	return &NullableForensicArtifactOutput{value: val, isSet: true}
}

func (v NullableForensicArtifactOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForensicArtifactOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}