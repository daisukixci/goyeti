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

// checks if the SuricataOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuricataOutput{}

// SuricataOutput Represents a Suricata rule.  Parsing and matching is yet TODO.
type SuricataOutput struct {
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
	Sid             *int32                           `json:"sid,omitempty"`
	Metadata        []string                         `json:"metadata,omitempty"`
	References      []string                         `json:"references,omitempty"`
	Id              string                           `json:"id"`
	Tags            map[string]TagRelationshipOutput `json:"tags"`
	RootType        string                           `json:"root_type"`
}

type _SuricataOutput SuricataOutput

// NewSuricataOutput instantiates a new SuricataOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuricataOutput(name string, pattern string, diamond DiamondModel, id string, tags map[string]TagRelationshipOutput, rootType string) *SuricataOutput {
	this := SuricataOutput{}
	this.Name = name
	var type_ string = "suricata"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	this.Pattern = pattern
	var location string = ""
	this.Location = &location
	this.Diamond = diamond
	var sid int32 = 0
	this.Sid = &sid
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	return &this
}

// NewSuricataOutputWithDefaults instantiates a new SuricataOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuricataOutputWithDefaults() *SuricataOutput {
	this := SuricataOutput{}
	var type_ string = "suricata"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	var location string = ""
	this.Location = &location
	var sid int32 = 0
	this.Sid = &sid
	return &this
}

// GetName returns the Name field value
func (o *SuricataOutput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SuricataOutput) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SuricataOutput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SuricataOutput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SuricataOutput) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SuricataOutput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SuricataOutput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SuricataOutput) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SuricataOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SuricataOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SuricataOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SuricataOutput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SuricataOutput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SuricataOutput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *SuricataOutput) GetValidFrom() time.Time {
	if o == nil || IsNil(o.ValidFrom) {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetValidFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *SuricataOutput) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *SuricataOutput) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *SuricataOutput) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *SuricataOutput) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *SuricataOutput) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetPattern returns the Pattern field value
func (o *SuricataOutput) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *SuricataOutput) SetPattern(v string) {
	o.Pattern = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SuricataOutput) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SuricataOutput) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SuricataOutput) SetLocation(v string) {
	o.Location = &v
}

// GetDiamond returns the Diamond field value
func (o *SuricataOutput) GetDiamond() DiamondModel {
	if o == nil {
		var ret DiamondModel
		return ret
	}

	return o.Diamond
}

// GetDiamondOk returns a tuple with the Diamond field value
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetDiamondOk() (*DiamondModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diamond, true
}

// SetDiamond sets field value
func (o *SuricataOutput) SetDiamond(v DiamondModel) {
	o.Diamond = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *SuricataOutput) GetKillChainPhases() []string {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []string
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetKillChainPhasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *SuricataOutput) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []string and assigns it to the KillChainPhases field.
func (o *SuricataOutput) SetKillChainPhases(v []string) {
	o.KillChainPhases = v
}

// GetRelevantTags returns the RelevantTags field value if set, zero value otherwise.
func (o *SuricataOutput) GetRelevantTags() []string {
	if o == nil || IsNil(o.RelevantTags) {
		var ret []string
		return ret
	}
	return o.RelevantTags
}

// GetRelevantTagsOk returns a tuple with the RelevantTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetRelevantTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelevantTags) {
		return nil, false
	}
	return o.RelevantTags, true
}

// HasRelevantTags returns a boolean if a field has been set.
func (o *SuricataOutput) HasRelevantTags() bool {
	if o != nil && !IsNil(o.RelevantTags) {
		return true
	}

	return false
}

// SetRelevantTags gets a reference to the given []string and assigns it to the RelevantTags field.
func (o *SuricataOutput) SetRelevantTags(v []string) {
	o.RelevantTags = v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *SuricataOutput) GetSid() int32 {
	if o == nil || IsNil(o.Sid) {
		var ret int32
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetSidOk() (*int32, bool) {
	if o == nil || IsNil(o.Sid) {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *SuricataOutput) HasSid() bool {
	if o != nil && !IsNil(o.Sid) {
		return true
	}

	return false
}

// SetSid gets a reference to the given int32 and assigns it to the Sid field.
func (o *SuricataOutput) SetSid(v int32) {
	o.Sid = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SuricataOutput) GetMetadata() []string {
	if o == nil || IsNil(o.Metadata) {
		var ret []string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetMetadataOk() ([]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SuricataOutput) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []string and assigns it to the Metadata field.
func (o *SuricataOutput) SetMetadata(v []string) {
	o.Metadata = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *SuricataOutput) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *SuricataOutput) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *SuricataOutput) SetReferences(v []string) {
	o.References = v
}

// GetId returns the Id field value
func (o *SuricataOutput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SuricataOutput) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *SuricataOutput) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *SuricataOutput) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *SuricataOutput) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *SuricataOutput) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *SuricataOutput) SetRootType(v string) {
	o.RootType = v
}

func (o SuricataOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuricataOutput) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Sid) {
		toSerialize["sid"] = o.Sid
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["root_type"] = o.RootType
	return toSerialize, nil
}

func (o *SuricataOutput) UnmarshalJSON(data []byte) (err error) {
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

	varSuricataOutput := _SuricataOutput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuricataOutput)

	if err != nil {
		return err
	}

	*o = SuricataOutput(varSuricataOutput)

	return err
}

type NullableSuricataOutput struct {
	value *SuricataOutput
	isSet bool
}

func (v NullableSuricataOutput) Get() *SuricataOutput {
	return v.value
}

func (v *NullableSuricataOutput) Set(val *SuricataOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableSuricataOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableSuricataOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuricataOutput(val *SuricataOutput) *NullableSuricataOutput {
	return &NullableSuricataOutput{value: val, isSet: true}
}

func (v NullableSuricataOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuricataOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}