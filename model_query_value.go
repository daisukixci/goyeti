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

// QueryValue struct for QueryValue
type QueryValue struct {
	ArrayOfAny *[]interface{}
	Int32      *int32
	String     *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *QueryValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfAny
	err = json.Unmarshal(data, &dst.ArrayOfAny)
	if err == nil {
		jsonArrayOfAny, _ := json.Marshal(dst.ArrayOfAny)
		if string(jsonArrayOfAny) == "{}" { // empty struct
			dst.ArrayOfAny = nil
		} else {
			return nil // data stored in dst.ArrayOfAny, return on the first match
		}
	} else {
		dst.ArrayOfAny = nil
	}

	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(QueryValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QueryValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfAny != nil {
		return json.Marshal(&src.ArrayOfAny)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableQueryValue struct {
	value *QueryValue
	isSet bool
}

func (v NullableQueryValue) Get() *QueryValue {
	return v.value
}

func (v *NullableQueryValue) Set(val *QueryValue) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryValue) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryValue(val *QueryValue) *NullableQueryValue {
	return &NullableQueryValue{value: val, isSet: true}
}

func (v NullableQueryValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
