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

// ObservableSearchRequestType struct for ObservableSearchRequestType
type ObservableSearchRequestType struct {
	ObservableTypeInput *ObservableTypeInput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ObservableSearchRequestType) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ObservableTypeInput
	err = json.Unmarshal(data, &dst.ObservableTypeInput)
	if err == nil {
		jsonObservableTypeInput, _ := json.Marshal(dst.ObservableTypeInput)
		if string(jsonObservableTypeInput) == "{}" { // empty struct
			dst.ObservableTypeInput = nil
		} else {
			return nil // data stored in dst.ObservableTypeInput, return on the first match
		}
	} else {
		dst.ObservableTypeInput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ObservableSearchRequestType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ObservableSearchRequestType) MarshalJSON() ([]byte, error) {
	if src.ObservableTypeInput != nil {
		return json.Marshal(&src.ObservableTypeInput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableObservableSearchRequestType struct {
	value *ObservableSearchRequestType
	isSet bool
}

func (v NullableObservableSearchRequestType) Get() *ObservableSearchRequestType {
	return v.value
}

func (v *NullableObservableSearchRequestType) Set(val *ObservableSearchRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableObservableSearchRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableObservableSearchRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservableSearchRequestType(val *ObservableSearchRequestType) *NullableObservableSearchRequestType {
	return &NullableObservableSearchRequestType{value: val, isSet: true}
}

func (v NullableObservableSearchRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservableSearchRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
