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
)

// checks if the GraphFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphFilter{}

// GraphFilter struct for GraphFilter
type GraphFilter struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}

type _GraphFilter GraphFilter

// NewGraphFilter instantiates a new GraphFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphFilter(key string, value string, operator string) *GraphFilter {
	this := GraphFilter{}
	this.Key = key
	this.Value = value
	this.Operator = operator
	return &this
}

// NewGraphFilterWithDefaults instantiates a new GraphFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphFilterWithDefaults() *GraphFilter {
	this := GraphFilter{}
	return &this
}

// GetKey returns the Key field value
func (o *GraphFilter) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GraphFilter) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GraphFilter) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *GraphFilter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GraphFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GraphFilter) SetValue(v string) {
	o.Value = v
}

// GetOperator returns the Operator field value
func (o *GraphFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *GraphFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *GraphFilter) SetOperator(v string) {
	o.Operator = v
}

func (o GraphFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *GraphFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
		"operator",
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

	varGraphFilter := _GraphFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphFilter)

	if err != nil {
		return err
	}

	*o = GraphFilter(varGraphFilter)

	return err
}

type NullableGraphFilter struct {
	value *GraphFilter
	isSet bool
}

func (v NullableGraphFilter) Get() *GraphFilter {
	return v.value
}

func (v *NullableGraphFilter) Set(val *GraphFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphFilter(val *GraphFilter) *NullableGraphFilter {
	return &NullableGraphFilter{value: val, isSet: true}
}

func (v NullableGraphFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
