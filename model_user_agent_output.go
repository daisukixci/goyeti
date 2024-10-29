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

// checks if the UserAgentOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAgentOutput{}

// UserAgentOutput struct for UserAgentOutput
type UserAgentOutput struct {
	Value        string                           `json:"value"`
	Type         *string                          `json:"type,omitempty"`
	Created      *time.Time                       `json:"created,omitempty"`
	Context      []map[string]interface{}         `json:"context,omitempty"`
	LastAnalysis map[string]time.Time             `json:"last_analysis,omitempty"`
	Id           string                           `json:"id"`
	Tags         map[string]TagRelationshipOutput `json:"tags"`
	RootType     string                           `json:"root_type"`
}

type _UserAgentOutput UserAgentOutput

// NewUserAgentOutput instantiates a new UserAgentOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAgentOutput(value string, id string, tags map[string]TagRelationshipOutput, rootType string) *UserAgentOutput {
	this := UserAgentOutput{}
	this.Value = value
	var type_ string = "user_agent"
	this.Type = &type_
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	return &this
}

// NewUserAgentOutputWithDefaults instantiates a new UserAgentOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAgentOutputWithDefaults() *UserAgentOutput {
	this := UserAgentOutput{}
	var type_ string = "user_agent"
	this.Type = &type_
	return &this
}

// GetValue returns the Value field value
func (o *UserAgentOutput) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserAgentOutput) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserAgentOutput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserAgentOutput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserAgentOutput) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserAgentOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserAgentOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserAgentOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *UserAgentOutput) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *UserAgentOutput) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *UserAgentOutput) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetLastAnalysis returns the LastAnalysis field value if set, zero value otherwise.
func (o *UserAgentOutput) GetLastAnalysis() map[string]time.Time {
	if o == nil || IsNil(o.LastAnalysis) {
		var ret map[string]time.Time
		return ret
	}
	return o.LastAnalysis
}

// GetLastAnalysisOk returns a tuple with the LastAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetLastAnalysisOk() (map[string]time.Time, bool) {
	if o == nil || IsNil(o.LastAnalysis) {
		return map[string]time.Time{}, false
	}
	return o.LastAnalysis, true
}

// HasLastAnalysis returns a boolean if a field has been set.
func (o *UserAgentOutput) HasLastAnalysis() bool {
	if o != nil && !IsNil(o.LastAnalysis) {
		return true
	}

	return false
}

// SetLastAnalysis gets a reference to the given map[string]time.Time and assigns it to the LastAnalysis field.
func (o *UserAgentOutput) SetLastAnalysis(v map[string]time.Time) {
	o.LastAnalysis = v
}

// GetId returns the Id field value
func (o *UserAgentOutput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserAgentOutput) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *UserAgentOutput) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *UserAgentOutput) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *UserAgentOutput) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *UserAgentOutput) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *UserAgentOutput) SetRootType(v string) {
	o.RootType = v
}

func (o UserAgentOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAgentOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.LastAnalysis) {
		toSerialize["last_analysis"] = o.LastAnalysis
	}
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["root_type"] = o.RootType
	return toSerialize, nil
}

func (o *UserAgentOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varUserAgentOutput := _UserAgentOutput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserAgentOutput)

	if err != nil {
		return err
	}

	*o = UserAgentOutput(varUserAgentOutput)

	return err
}

type NullableUserAgentOutput struct {
	value *UserAgentOutput
	isSet bool
}

func (v NullableUserAgentOutput) Get() *UserAgentOutput {
	return v.value
}

func (v *NullableUserAgentOutput) Set(val *UserAgentOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAgentOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAgentOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAgentOutput(val *UserAgentOutput) *NullableUserAgentOutput {
	return &NullableUserAgentOutput{value: val, isSet: true}
}

func (v NullableUserAgentOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAgentOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}