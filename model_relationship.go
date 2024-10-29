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

// checks if the Relationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Relationship{}

// Relationship struct for Relationship
type Relationship struct {
	Source      string    `json:"source"`
	Target      string    `json:"target"`
	Type        string    `json:"type"`
	Count       *int32    `json:"count,omitempty"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Id          string    `json:"id"`
}

type _Relationship Relationship

// NewRelationship instantiates a new Relationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship(source string, target string, type_ string, description string, created time.Time, modified time.Time, id string) *Relationship {
	this := Relationship{}
	this.Source = source
	this.Target = target
	this.Type = type_
	var count int32 = 1
	this.Count = &count
	this.Description = description
	this.Created = created
	this.Modified = modified
	this.Id = id
	return &this
}

// NewRelationshipWithDefaults instantiates a new Relationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWithDefaults() *Relationship {
	this := Relationship{}
	var count int32 = 1
	this.Count = &count
	return &this
}

// GetSource returns the Source field value
func (o *Relationship) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Relationship) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *Relationship) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Relationship) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value
func (o *Relationship) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Relationship) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Relationship) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Relationship) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Relationship) SetCount(v int32) {
	o.Count = &v
}

// GetDescription returns the Description field value
func (o *Relationship) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Relationship) SetDescription(v string) {
	o.Description = v
}

// GetCreated returns the Created field value
func (o *Relationship) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Relationship) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *Relationship) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *Relationship) SetModified(v time.Time) {
	o.Modified = v
}

// GetId returns the Id field value
func (o *Relationship) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Relationship) SetId(v string) {
	o.Id = v
}

func (o Relationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Relationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["description"] = o.Description
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *Relationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"target",
		"type",
		"description",
		"created",
		"modified",
		"id",
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

	varRelationship := _Relationship{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelationship)

	if err != nil {
		return err
	}

	*o = Relationship(varRelationship)

	return err
}

type NullableRelationship struct {
	value *Relationship
	isSet bool
}

func (v NullableRelationship) Get() *Relationship {
	return v.value
}

func (v *NullableRelationship) Set(val *Relationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship(val *Relationship) *NullableRelationship {
	return &NullableRelationship{value: val, isSet: true}
}

func (v NullableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}