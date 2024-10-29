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

// checks if the YaraInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YaraInput{}

// YaraInput Represents a Yara rule.  Parsing and matching is yet TODO.
type YaraInput struct {
	Name            string       `json:"name"`
	Type            *string      `json:"type,omitempty"`
	Description     *string      `json:"description,omitempty"`
	Created         *time.Time   `json:"created,omitempty"`
	Modified        *time.Time   `json:"modified,omitempty"`
	ValidFrom       *time.Time   `json:"valid_from,omitempty"`
	ValidUntil      *time.Time   `json:"valid_until,omitempty"`
	Pattern         string       `json:"pattern"`
	Location        *string      `json:"location,omitempty"`
	Diamond         DiamondModel `json:"diamond"`
	KillChainPhases []string     `json:"kill_chain_phases,omitempty"`
	RelevantTags    []string     `json:"relevant_tags,omitempty"`
}

type _YaraInput YaraInput

// NewYaraInput instantiates a new YaraInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYaraInput(name string, pattern string, diamond DiamondModel) *YaraInput {
	this := YaraInput{}
	this.Name = name
	var type_ string = "yara"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	this.Pattern = pattern
	var location string = ""
	this.Location = &location
	this.Diamond = diamond
	return &this
}

// NewYaraInputWithDefaults instantiates a new YaraInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYaraInputWithDefaults() *YaraInput {
	this := YaraInput{}
	var type_ string = "yara"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	var location string = ""
	this.Location = &location
	return &this
}

// GetName returns the Name field value
func (o *YaraInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *YaraInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *YaraInput) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *YaraInput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *YaraInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *YaraInput) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *YaraInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *YaraInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *YaraInput) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *YaraInput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *YaraInput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *YaraInput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *YaraInput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *YaraInput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *YaraInput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *YaraInput) GetValidFrom() time.Time {
	if o == nil || IsNil(o.ValidFrom) {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetValidFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *YaraInput) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *YaraInput) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *YaraInput) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *YaraInput) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *YaraInput) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetPattern returns the Pattern field value
func (o *YaraInput) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *YaraInput) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *YaraInput) SetPattern(v string) {
	o.Pattern = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *YaraInput) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *YaraInput) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *YaraInput) SetLocation(v string) {
	o.Location = &v
}

// GetDiamond returns the Diamond field value
func (o *YaraInput) GetDiamond() DiamondModel {
	if o == nil {
		var ret DiamondModel
		return ret
	}

	return o.Diamond
}

// GetDiamondOk returns a tuple with the Diamond field value
// and a boolean to check if the value has been set.
func (o *YaraInput) GetDiamondOk() (*DiamondModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diamond, true
}

// SetDiamond sets field value
func (o *YaraInput) SetDiamond(v DiamondModel) {
	o.Diamond = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *YaraInput) GetKillChainPhases() []string {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []string
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetKillChainPhasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *YaraInput) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []string and assigns it to the KillChainPhases field.
func (o *YaraInput) SetKillChainPhases(v []string) {
	o.KillChainPhases = v
}

// GetRelevantTags returns the RelevantTags field value if set, zero value otherwise.
func (o *YaraInput) GetRelevantTags() []string {
	if o == nil || IsNil(o.RelevantTags) {
		var ret []string
		return ret
	}
	return o.RelevantTags
}

// GetRelevantTagsOk returns a tuple with the RelevantTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YaraInput) GetRelevantTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RelevantTags) {
		return nil, false
	}
	return o.RelevantTags, true
}

// HasRelevantTags returns a boolean if a field has been set.
func (o *YaraInput) HasRelevantTags() bool {
	if o != nil && !IsNil(o.RelevantTags) {
		return true
	}

	return false
}

// SetRelevantTags gets a reference to the given []string and assigns it to the RelevantTags field.
func (o *YaraInput) SetRelevantTags(v []string) {
	o.RelevantTags = v
}

func (o YaraInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YaraInput) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *YaraInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"pattern",
		"diamond",
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

	varYaraInput := _YaraInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varYaraInput)

	if err != nil {
		return err
	}

	*o = YaraInput(varYaraInput)

	return err
}

type NullableYaraInput struct {
	value *YaraInput
	isSet bool
}

func (v NullableYaraInput) Get() *YaraInput {
	return v.value
}

func (v *NullableYaraInput) Set(val *YaraInput) {
	v.value = val
	v.isSet = true
}

func (v NullableYaraInput) IsSet() bool {
	return v.isSet
}

func (v *NullableYaraInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYaraInput(val *YaraInput) *NullableYaraInput {
	return &NullableYaraInput{value: val, isSet: true}
}

func (v NullableYaraInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYaraInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}