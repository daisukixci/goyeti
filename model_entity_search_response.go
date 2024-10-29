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

// checks if the EntitySearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySearchResponse{}

// EntitySearchResponse struct for EntitySearchResponse
type EntitySearchResponse struct {
	Entities             []EntitiesInner `json:"entities"`
	Total                int32           `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _EntitySearchResponse EntitySearchResponse

// NewEntitySearchResponse instantiates a new EntitySearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySearchResponse(entities []EntitiesInner, total int32) *EntitySearchResponse {
	this := EntitySearchResponse{}
	this.Entities = entities
	this.Total = total
	return &this
}

// NewEntitySearchResponseWithDefaults instantiates a new EntitySearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySearchResponseWithDefaults() *EntitySearchResponse {
	this := EntitySearchResponse{}
	return &this
}

// GetEntities returns the Entities field value
func (o *EntitySearchResponse) GetEntities() []EntitiesInner {
	if o == nil {
		var ret []EntitiesInner
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *EntitySearchResponse) GetEntitiesOk() ([]EntitiesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entities, true
}

// SetEntities sets field value
func (o *EntitySearchResponse) SetEntities(v []EntitiesInner) {
	o.Entities = v
}

// GetTotal returns the Total field value
func (o *EntitySearchResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *EntitySearchResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *EntitySearchResponse) SetTotal(v int32) {
	o.Total = v
}

func (o EntitySearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entities"] = o.Entities
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitySearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entities",
		"total",
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

	varEntitySearchResponse := _EntitySearchResponse{}

	err = json.Unmarshal(data, &varEntitySearchResponse)

	if err != nil {
		return err
	}

	*o = EntitySearchResponse(varEntitySearchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entities")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitySearchResponse struct {
	value *EntitySearchResponse
	isSet bool
}

func (v NullableEntitySearchResponse) Get() *EntitySearchResponse {
	return v.value
}

func (v *NullableEntitySearchResponse) Set(val *EntitySearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySearchResponse(val *EntitySearchResponse) *NullableEntitySearchResponse {
	return &NullableEntitySearchResponse{value: val, isSet: true}
}

func (v NullableEntitySearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
