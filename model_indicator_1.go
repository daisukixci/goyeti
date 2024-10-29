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
	"gopkg.in/validator.v2"
)

// Indicator1 - struct for Indicator1
type Indicator1 struct {
	ForensicArtifactInput *ForensicArtifactInput
	QueryInput            *QueryInput
	RegexInput            *RegexInput
	SigmaInput            *SigmaInput
	SuricataInput         *SuricataInput
	YaraInput             *YaraInput
}

// ForensicArtifactInputAsIndicator1 is a convenience function that returns ForensicArtifactInput wrapped in Indicator1
func ForensicArtifactInputAsIndicator1(v *ForensicArtifactInput) Indicator1 {
	return Indicator1{
		ForensicArtifactInput: v,
	}
}

// QueryInputAsIndicator1 is a convenience function that returns QueryInput wrapped in Indicator1
func QueryInputAsIndicator1(v *QueryInput) Indicator1 {
	return Indicator1{
		QueryInput: v,
	}
}

// RegexInputAsIndicator1 is a convenience function that returns RegexInput wrapped in Indicator1
func RegexInputAsIndicator1(v *RegexInput) Indicator1 {
	return Indicator1{
		RegexInput: v,
	}
}

// SigmaInputAsIndicator1 is a convenience function that returns SigmaInput wrapped in Indicator1
func SigmaInputAsIndicator1(v *SigmaInput) Indicator1 {
	return Indicator1{
		SigmaInput: v,
	}
}

// SuricataInputAsIndicator1 is a convenience function that returns SuricataInput wrapped in Indicator1
func SuricataInputAsIndicator1(v *SuricataInput) Indicator1 {
	return Indicator1{
		SuricataInput: v,
	}
}

// YaraInputAsIndicator1 is a convenience function that returns YaraInput wrapped in Indicator1
func YaraInputAsIndicator1(v *YaraInput) Indicator1 {
	return Indicator1{
		YaraInput: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Indicator1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ForensicArtifactInput
	err = newStrictDecoder(data).Decode(&dst.ForensicArtifactInput)
	if err == nil {
		jsonForensicArtifactInput, _ := json.Marshal(dst.ForensicArtifactInput)
		if string(jsonForensicArtifactInput) == "{}" { // empty struct
			dst.ForensicArtifactInput = nil
		} else {
			if err = validator.Validate(dst.ForensicArtifactInput); err != nil {
				dst.ForensicArtifactInput = nil
			} else {
				match++
			}
		}
	} else {
		dst.ForensicArtifactInput = nil
	}

	// try to unmarshal data into QueryInput
	err = newStrictDecoder(data).Decode(&dst.QueryInput)
	if err == nil {
		jsonQueryInput, _ := json.Marshal(dst.QueryInput)
		if string(jsonQueryInput) == "{}" { // empty struct
			dst.QueryInput = nil
		} else {
			if err = validator.Validate(dst.QueryInput); err != nil {
				dst.QueryInput = nil
			} else {
				match++
			}
		}
	} else {
		dst.QueryInput = nil
	}

	// try to unmarshal data into RegexInput
	err = newStrictDecoder(data).Decode(&dst.RegexInput)
	if err == nil {
		jsonRegexInput, _ := json.Marshal(dst.RegexInput)
		if string(jsonRegexInput) == "{}" { // empty struct
			dst.RegexInput = nil
		} else {
			if err = validator.Validate(dst.RegexInput); err != nil {
				dst.RegexInput = nil
			} else {
				match++
			}
		}
	} else {
		dst.RegexInput = nil
	}

	// try to unmarshal data into SigmaInput
	err = newStrictDecoder(data).Decode(&dst.SigmaInput)
	if err == nil {
		jsonSigmaInput, _ := json.Marshal(dst.SigmaInput)
		if string(jsonSigmaInput) == "{}" { // empty struct
			dst.SigmaInput = nil
		} else {
			if err = validator.Validate(dst.SigmaInput); err != nil {
				dst.SigmaInput = nil
			} else {
				match++
			}
		}
	} else {
		dst.SigmaInput = nil
	}

	// try to unmarshal data into SuricataInput
	err = newStrictDecoder(data).Decode(&dst.SuricataInput)
	if err == nil {
		jsonSuricataInput, _ := json.Marshal(dst.SuricataInput)
		if string(jsonSuricataInput) == "{}" { // empty struct
			dst.SuricataInput = nil
		} else {
			if err = validator.Validate(dst.SuricataInput); err != nil {
				dst.SuricataInput = nil
			} else {
				match++
			}
		}
	} else {
		dst.SuricataInput = nil
	}

	// try to unmarshal data into YaraInput
	err = newStrictDecoder(data).Decode(&dst.YaraInput)
	if err == nil {
		jsonYaraInput, _ := json.Marshal(dst.YaraInput)
		if string(jsonYaraInput) == "{}" { // empty struct
			dst.YaraInput = nil
		} else {
			if err = validator.Validate(dst.YaraInput); err != nil {
				dst.YaraInput = nil
			} else {
				match++
			}
		}
	} else {
		dst.YaraInput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ForensicArtifactInput = nil
		dst.QueryInput = nil
		dst.RegexInput = nil
		dst.SigmaInput = nil
		dst.SuricataInput = nil
		dst.YaraInput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Indicator1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Indicator1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Indicator1) MarshalJSON() ([]byte, error) {
	if src.ForensicArtifactInput != nil {
		return json.Marshal(&src.ForensicArtifactInput)
	}

	if src.QueryInput != nil {
		return json.Marshal(&src.QueryInput)
	}

	if src.RegexInput != nil {
		return json.Marshal(&src.RegexInput)
	}

	if src.SigmaInput != nil {
		return json.Marshal(&src.SigmaInput)
	}

	if src.SuricataInput != nil {
		return json.Marshal(&src.SuricataInput)
	}

	if src.YaraInput != nil {
		return json.Marshal(&src.YaraInput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Indicator1) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ForensicArtifactInput != nil {
		return obj.ForensicArtifactInput
	}

	if obj.QueryInput != nil {
		return obj.QueryInput
	}

	if obj.RegexInput != nil {
		return obj.RegexInput
	}

	if obj.SigmaInput != nil {
		return obj.SigmaInput
	}

	if obj.SuricataInput != nil {
		return obj.SuricataInput
	}

	if obj.YaraInput != nil {
		return obj.YaraInput
	}

	// all schemas are nil
	return nil
}

type NullableIndicator1 struct {
	value *Indicator1
	isSet bool
}

func (v NullableIndicator1) Get() *Indicator1 {
	return v.value
}

func (v *NullableIndicator1) Set(val *Indicator1) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicator1) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicator1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicator1(val *Indicator1) *NullableIndicator1 {
	return &NullableIndicator1{value: val, isSet: true}
}

func (v NullableIndicator1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicator1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
