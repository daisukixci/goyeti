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

// RegistryHive Registry Hive enum class.
type RegistryHive string

// List of RegistryHive
const (
	CURRENT_CONFIG         RegistryHive = "HKEY_CURRENT_CONFIG"
	CURRENT_USER           RegistryHive = "HKEY_CURRENT_USER"
	LOCAL_MACHINE_SAM      RegistryHive = "HKEY_LOCAL_MACHINE_SAM"
	LOCAL_MACHINE_SECURITY RegistryHive = "HKEY_LOCAL_MACHINE_Security"
	LOCAL_MACHINE_SOFTWARE RegistryHive = "HKEY_LOCAL_MACHINE_Software"
	LOCAL_MACHINE_SYSTEM   RegistryHive = "HKEY_LOCAL_MACHINE_System"
	USERS_DEFAULT          RegistryHive = "HKEY_USERS_DEFAULT"
)

// All allowed values of RegistryHive enum
var AllowedRegistryHiveEnumValues = []RegistryHive{
	"HKEY_CURRENT_CONFIG",
	"HKEY_CURRENT_USER",
	"HKEY_LOCAL_MACHINE_SAM",
	"HKEY_LOCAL_MACHINE_Security",
	"HKEY_LOCAL_MACHINE_Software",
	"HKEY_LOCAL_MACHINE_System",
	"HKEY_USERS_DEFAULT",
}

func (v *RegistryHive) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistryHive(value)
	for _, existing := range AllowedRegistryHiveEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistryHive", value)
}

// NewRegistryHiveFromValue returns a pointer to a valid RegistryHive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistryHiveFromValue(v string) (*RegistryHive, error) {
	ev := RegistryHive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistryHive: valid values are %v", v, AllowedRegistryHiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistryHive) IsValid() bool {
	for _, existing := range AllowedRegistryHiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegistryHive value
func (v RegistryHive) Ptr() *RegistryHive {
	return &v
}

type NullableRegistryHive struct {
	value *RegistryHive
	isSet bool
}

func (v NullableRegistryHive) Get() *RegistryHive {
	return v.value
}

func (v *NullableRegistryHive) Set(val *RegistryHive) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryHive) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryHive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryHive(val *RegistryHive) *NullableRegistryHive {
	return &NullableRegistryHive{value: val, isSet: true}
}

func (v NullableRegistryHive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryHive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
