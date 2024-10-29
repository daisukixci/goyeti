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

// PathsInnerInner struct for PathsInnerInner
type PathsInnerInner struct {
	Relationship          *Relationship
	TagRelationshipOutput *TagRelationshipOutput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PathsInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Relationship
	err = json.Unmarshal(data, &dst.Relationship)
	if err == nil {
		jsonRelationship, _ := json.Marshal(dst.Relationship)
		if string(jsonRelationship) == "{}" { // empty struct
			dst.Relationship = nil
		} else {
			return nil // data stored in dst.Relationship, return on the first match
		}
	} else {
		dst.Relationship = nil
	}

	// try to unmarshal JSON data into TagRelationshipOutput
	err = json.Unmarshal(data, &dst.TagRelationshipOutput)
	if err == nil {
		jsonTagRelationshipOutput, _ := json.Marshal(dst.TagRelationshipOutput)
		if string(jsonTagRelationshipOutput) == "{}" { // empty struct
			dst.TagRelationshipOutput = nil
		} else {
			return nil // data stored in dst.TagRelationshipOutput, return on the first match
		}
	} else {
		dst.TagRelationshipOutput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PathsInnerInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PathsInnerInner) MarshalJSON() ([]byte, error) {
	if src.Relationship != nil {
		return json.Marshal(&src.Relationship)
	}

	if src.TagRelationshipOutput != nil {
		return json.Marshal(&src.TagRelationshipOutput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePathsInnerInner struct {
	value *PathsInnerInner
	isSet bool
}

func (v NullablePathsInnerInner) Get() *PathsInnerInner {
	return v.value
}

func (v *NullablePathsInnerInner) Set(val *PathsInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePathsInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePathsInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathsInnerInner(val *PathsInnerInner) *NullablePathsInnerInner {
	return &NullablePathsInnerInner{value: val, isSet: true}
}

func (v NullablePathsInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathsInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}