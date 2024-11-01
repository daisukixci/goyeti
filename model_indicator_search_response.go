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

// checks if the IndicatorSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorSearchResponse{}

// IndicatorSearchResponse struct for IndicatorSearchResponse
type IndicatorSearchResponse struct {
	Indicators           []IndicatorsInner `json:"indicators"`
	Total                int32             `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _IndicatorSearchResponse IndicatorSearchResponse

// NewIndicatorSearchResponse instantiates a new IndicatorSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorSearchResponse(indicators []IndicatorsInner, total int32) *IndicatorSearchResponse {
	this := IndicatorSearchResponse{}
	this.Indicators = indicators
	this.Total = total
	return &this
}

// NewIndicatorSearchResponseWithDefaults instantiates a new IndicatorSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorSearchResponseWithDefaults() *IndicatorSearchResponse {
	this := IndicatorSearchResponse{}
	return &this
}

// GetIndicators returns the Indicators field value
func (o *IndicatorSearchResponse) GetIndicators() []IndicatorsInner {
	if o == nil {
		var ret []IndicatorsInner
		return ret
	}

	return o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value
// and a boolean to check if the value has been set.
func (o *IndicatorSearchResponse) GetIndicatorsOk() ([]IndicatorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indicators, true
}

// SetIndicators sets field value
func (o *IndicatorSearchResponse) SetIndicators(v []IndicatorsInner) {
	o.Indicators = v
}

// GetTotal returns the Total field value
func (o *IndicatorSearchResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *IndicatorSearchResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *IndicatorSearchResponse) SetTotal(v int32) {
	o.Total = v
}

func (o IndicatorSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["indicators"] = o.Indicators
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndicatorSearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"indicators",
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

	varIndicatorSearchResponse := _IndicatorSearchResponse{}

	err = json.Unmarshal(data, &varIndicatorSearchResponse)

	if err != nil {
		return err
	}

	*o = IndicatorSearchResponse(varIndicatorSearchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "indicators")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndicatorSearchResponse struct {
	value *IndicatorSearchResponse
	isSet bool
}

func (v NullableIndicatorSearchResponse) Get() *IndicatorSearchResponse {
	return v.value
}

func (v *NullableIndicatorSearchResponse) Set(val *IndicatorSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorSearchResponse(val *IndicatorSearchResponse) *NullableIndicatorSearchResponse {
	return &NullableIndicatorSearchResponse{value: val, isSet: true}
}

func (v NullableIndicatorSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
