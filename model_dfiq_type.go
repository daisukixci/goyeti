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

// DFIQType the model 'DFIQType'
type DFIQType string

// List of DFIQType
const (
	SCENARIO DFIQType = "scenario"
	FACET    DFIQType = "facet"
	QUESTION DFIQType = "question"
)

// All allowed values of DFIQType enum
var AllowedDFIQTypeEnumValues = []DFIQType{
	"scenario",
	"facet",
	"question",
}

func (v *DFIQType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DFIQType(value)
	for _, existing := range AllowedDFIQTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DFIQType", value)
}

// NewDFIQTypeFromValue returns a pointer to a valid DFIQType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDFIQTypeFromValue(v string) (*DFIQType, error) {
	ev := DFIQType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DFIQType: valid values are %v", v, AllowedDFIQTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DFIQType) IsValid() bool {
	for _, existing := range AllowedDFIQTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DFIQType value
func (v DFIQType) Ptr() *DFIQType {
	return &v
}

type NullableDFIQType struct {
	value *DFIQType
	isSet bool
}

func (v NullableDFIQType) Get() *DFIQType {
	return v.value
}

func (v *NullableDFIQType) Set(val *DFIQType) {
	v.value = val
	v.isSet = true
}

func (v NullableDFIQType) IsSet() bool {
	return v.isSet
}

func (v *NullableDFIQType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDFIQType(val *DFIQType) *NullableDFIQType {
	return &NullableDFIQType{value: val, isSet: true}
}

func (v NullableDFIQType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDFIQType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
