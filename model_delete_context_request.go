/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"encoding/json"
	"fmt"
)

// checks if the DeleteContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteContextRequest{}

// DeleteContextRequest struct for DeleteContextRequest
type DeleteContextRequest struct {
	Source               string                 `json:"source"`
	Context              map[string]interface{} `json:"context"`
	SkipCompare          []interface{}          `json:"skip_compare,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteContextRequest DeleteContextRequest

// NewDeleteContextRequest instantiates a new DeleteContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteContextRequest(source string, context map[string]interface{}) *DeleteContextRequest {
	this := DeleteContextRequest{}
	this.Source = source
	this.Context = context
	return &this
}

// NewDeleteContextRequestWithDefaults instantiates a new DeleteContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteContextRequestWithDefaults() *DeleteContextRequest {
	this := DeleteContextRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *DeleteContextRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DeleteContextRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DeleteContextRequest) SetSource(v string) {
	o.Source = v
}

// GetContext returns the Context field value
func (o *DeleteContextRequest) GetContext() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *DeleteContextRequest) GetContextOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// SetContext sets field value
func (o *DeleteContextRequest) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetSkipCompare returns the SkipCompare field value if set, zero value otherwise.
func (o *DeleteContextRequest) GetSkipCompare() []interface{} {
	if o == nil || IsNil(o.SkipCompare) {
		var ret []interface{}
		return ret
	}
	return o.SkipCompare
}

// GetSkipCompareOk returns a tuple with the SkipCompare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteContextRequest) GetSkipCompareOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.SkipCompare) {
		return nil, false
	}
	return o.SkipCompare, true
}

// HasSkipCompare returns a boolean if a field has been set.
func (o *DeleteContextRequest) HasSkipCompare() bool {
	if o != nil && !IsNil(o.SkipCompare) {
		return true
	}

	return false
}

// SetSkipCompare gets a reference to the given []interface{} and assigns it to the SkipCompare field.
func (o *DeleteContextRequest) SetSkipCompare(v []interface{}) {
	o.SkipCompare = v
}

func (o DeleteContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["context"] = o.Context
	if !IsNil(o.SkipCompare) {
		toSerialize["skip_compare"] = o.SkipCompare
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteContextRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"context",
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

	varDeleteContextRequest := _DeleteContextRequest{}

	err = json.Unmarshal(data, &varDeleteContextRequest)

	if err != nil {
		return err
	}

	*o = DeleteContextRequest(varDeleteContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "context")
		delete(additionalProperties, "skip_compare")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteContextRequest struct {
	value *DeleteContextRequest
	isSet bool
}

func (v NullableDeleteContextRequest) Get() *DeleteContextRequest {
	return v.value
}

func (v *NullableDeleteContextRequest) Set(val *DeleteContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteContextRequest(val *DeleteContextRequest) *NullableDeleteContextRequest {
	return &NullableDeleteContextRequest{value: val, isSet: true}
}

func (v NullableDeleteContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}