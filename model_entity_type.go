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

// EntityType the model 'EntityType'
type EntityType string

// List of EntityType
const (
	INVESTIGATION    EntityType = "investigation"
	COURSE_OF_ACTION EntityType = "course-of-action"
	IDENTITY         EntityType = "identity"
	INTRUSION_SET    EntityType = "intrusion-set"
	TOOL             EntityType = "tool"
	THREAT_ACTOR     EntityType = "threat-actor"
	CAMPAIGN         EntityType = "campaign"
	COMPANY          EntityType = "company"
	MALWARE          EntityType = "malware"
	PHONE            EntityType = "phone"
	NOTE             EntityType = "note"
	ATTACK_PATTERN   EntityType = "attack-pattern"
	VULNERABILITY    EntityType = "vulnerability"
	HONEYPOT         EntityType = "honeypot"
	THREAT_FINDING   EntityType = "threat-finding"
)

// All allowed values of EntityType enum
var AllowedEntityTypeEnumValues = []EntityType{
	"investigation",
	"course-of-action",
	"identity",
	"intrusion-set",
	"tool",
	"threat-actor",
	"campaign",
	"company",
	"malware",
	"phone",
	"note",
	"attack-pattern",
	"vulnerability",
	"honeypot",
	"threat-finding",
}

func (v *EntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntityType(value)
	for _, existing := range AllowedEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntityType", value)
}

// NewEntityTypeFromValue returns a pointer to a valid EntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntityTypeFromValue(v string) (*EntityType, error) {
	ev := EntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntityType: valid values are %v", v, AllowedEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntityType) IsValid() bool {
	for _, existing := range AllowedEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityType value
func (v EntityType) Ptr() *EntityType {
	return &v
}

type NullableEntityType struct {
	value *EntityType
	isSet bool
}

func (v NullableEntityType) Get() *EntityType {
	return v.value
}

func (v *NullableEntityType) Set(val *EntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityType(val *EntityType) *NullableEntityType {
	return &NullableEntityType{value: val, isSet: true}
}

func (v NullableEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
