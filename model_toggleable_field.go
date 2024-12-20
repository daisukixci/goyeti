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

// ToggleableField the model 'ToggleableField'
type ToggleableField string

// List of ToggleableField
const (
	ENABLED ToggleableField = "enabled"
	ADMIN   ToggleableField = "admin"
)

// All allowed values of ToggleableField enum
var AllowedToggleableFieldEnumValues = []ToggleableField{
	"enabled",
	"admin",
}

func (v *ToggleableField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ToggleableField(value)
	for _, existing := range AllowedToggleableFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ToggleableField", value)
}

// NewToggleableFieldFromValue returns a pointer to a valid ToggleableField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToggleableFieldFromValue(v string) (*ToggleableField, error) {
	ev := ToggleableField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToggleableField: valid values are %v", v, AllowedToggleableFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToggleableField) IsValid() bool {
	for _, existing := range AllowedToggleableFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToggleableField value
func (v ToggleableField) Ptr() *ToggleableField {
	return &v
}

type NullableToggleableField struct {
	value *ToggleableField
	isSet bool
}

func (v NullableToggleableField) Get() *ToggleableField {
	return v.value
}

func (v *NullableToggleableField) Set(val *ToggleableField) {
	v.value = val
	v.isSet = true
}

func (v NullableToggleableField) IsSet() bool {
	return v.isSet
}

func (v *NullableToggleableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToggleableField(val *ToggleableField) *NullableToggleableField {
	return &NullableToggleableField{value: val, isSet: true}
}

func (v NullableToggleableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToggleableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
