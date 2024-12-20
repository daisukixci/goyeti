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

// checks if the IPv4Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv4Output{}

// IPv4Output struct for IPv4Output
type IPv4Output struct {
	Value        string                           `json:"value"`
	Type         *string                          `json:"type,omitempty"`
	Created      *time.Time                       `json:"created,omitempty"`
	Context      []map[string]interface{}         `json:"context,omitempty"`
	LastAnalysis map[string]time.Time             `json:"last_analysis,omitempty"`
	Id           string                           `json:"id"`
	Tags         map[string]TagRelationshipOutput `json:"tags"`
	RootType     string                           `json:"root_type"`
}

type _IPv4Output IPv4Output

// NewIPv4Output instantiates a new IPv4Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv4Output(value string, id string, tags map[string]TagRelationshipOutput, rootType string) *IPv4Output {
	this := IPv4Output{}
	this.Value = value
	var type_ string = "ipv4"
	this.Type = &type_
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	return &this
}

// NewIPv4OutputWithDefaults instantiates a new IPv4Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv4OutputWithDefaults() *IPv4Output {
	this := IPv4Output{}
	var type_ string = "ipv4"
	this.Type = &type_
	return &this
}

// GetValue returns the Value field value
func (o *IPv4Output) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IPv4Output) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv4Output) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv4Output) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv4Output) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IPv4Output) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IPv4Output) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IPv4Output) SetCreated(v time.Time) {
	o.Created = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *IPv4Output) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *IPv4Output) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *IPv4Output) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetLastAnalysis returns the LastAnalysis field value if set, zero value otherwise.
func (o *IPv4Output) GetLastAnalysis() map[string]time.Time {
	if o == nil || IsNil(o.LastAnalysis) {
		var ret map[string]time.Time
		return ret
	}
	return o.LastAnalysis
}

// GetLastAnalysisOk returns a tuple with the LastAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetLastAnalysisOk() (map[string]time.Time, bool) {
	if o == nil || IsNil(o.LastAnalysis) {
		return map[string]time.Time{}, false
	}
	return o.LastAnalysis, true
}

// HasLastAnalysis returns a boolean if a field has been set.
func (o *IPv4Output) HasLastAnalysis() bool {
	if o != nil && !IsNil(o.LastAnalysis) {
		return true
	}

	return false
}

// SetLastAnalysis gets a reference to the given map[string]time.Time and assigns it to the LastAnalysis field.
func (o *IPv4Output) SetLastAnalysis(v map[string]time.Time) {
	o.LastAnalysis = v
}

// GetId returns the Id field value
func (o *IPv4Output) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPv4Output) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *IPv4Output) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *IPv4Output) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *IPv4Output) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *IPv4Output) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *IPv4Output) SetRootType(v string) {
	o.RootType = v
}

func (o IPv4Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv4Output) ToMap() (map[string]interface{}, error) {
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

func (o *IPv4Output) UnmarshalJSON(data []byte) (err error) {
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

	varIPv4Output := _IPv4Output{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIPv4Output)

	if err != nil {
		return err
	}

	*o = IPv4Output(varIPv4Output)

	return err
}

type NullableIPv4Output struct {
	value *IPv4Output
	isSet bool
}

func (v NullableIPv4Output) Get() *IPv4Output {
	return v.value
}

func (v *NullableIPv4Output) Set(val *IPv4Output) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv4Output) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv4Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv4Output(val *IPv4Output) *NullableIPv4Output {
	return &NullableIPv4Output{value: val, isSet: true}
}

func (v NullableIPv4Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv4Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
