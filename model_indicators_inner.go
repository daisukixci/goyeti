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

// IndicatorsInner struct for IndicatorsInner
type IndicatorsInner struct {
	ForensicArtifactOutput *ForensicArtifactOutput
	QueryOutput            *QueryOutput
	RegexOutput            *RegexOutput
	SigmaOutput            *SigmaOutput
	SuricataOutput         *SuricataOutput
	YaraOutput             *YaraOutput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IndicatorsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ForensicArtifactOutput
	err = json.Unmarshal(data, &dst.ForensicArtifactOutput)
	if err == nil {
		jsonForensicArtifactOutput, _ := json.Marshal(dst.ForensicArtifactOutput)
		if string(jsonForensicArtifactOutput) == "{}" { // empty struct
			dst.ForensicArtifactOutput = nil
		} else {
			return nil // data stored in dst.ForensicArtifactOutput, return on the first match
		}
	} else {
		dst.ForensicArtifactOutput = nil
	}

	// try to unmarshal JSON data into QueryOutput
	err = json.Unmarshal(data, &dst.QueryOutput)
	if err == nil {
		jsonQueryOutput, _ := json.Marshal(dst.QueryOutput)
		if string(jsonQueryOutput) == "{}" { // empty struct
			dst.QueryOutput = nil
		} else {
			return nil // data stored in dst.QueryOutput, return on the first match
		}
	} else {
		dst.QueryOutput = nil
	}

	// try to unmarshal JSON data into RegexOutput
	err = json.Unmarshal(data, &dst.RegexOutput)
	if err == nil {
		jsonRegexOutput, _ := json.Marshal(dst.RegexOutput)
		if string(jsonRegexOutput) == "{}" { // empty struct
			dst.RegexOutput = nil
		} else {
			return nil // data stored in dst.RegexOutput, return on the first match
		}
	} else {
		dst.RegexOutput = nil
	}

	// try to unmarshal JSON data into SigmaOutput
	err = json.Unmarshal(data, &dst.SigmaOutput)
	if err == nil {
		jsonSigmaOutput, _ := json.Marshal(dst.SigmaOutput)
		if string(jsonSigmaOutput) == "{}" { // empty struct
			dst.SigmaOutput = nil
		} else {
			return nil // data stored in dst.SigmaOutput, return on the first match
		}
	} else {
		dst.SigmaOutput = nil
	}

	// try to unmarshal JSON data into SuricataOutput
	err = json.Unmarshal(data, &dst.SuricataOutput)
	if err == nil {
		jsonSuricataOutput, _ := json.Marshal(dst.SuricataOutput)
		if string(jsonSuricataOutput) == "{}" { // empty struct
			dst.SuricataOutput = nil
		} else {
			return nil // data stored in dst.SuricataOutput, return on the first match
		}
	} else {
		dst.SuricataOutput = nil
	}

	// try to unmarshal JSON data into YaraOutput
	err = json.Unmarshal(data, &dst.YaraOutput)
	if err == nil {
		jsonYaraOutput, _ := json.Marshal(dst.YaraOutput)
		if string(jsonYaraOutput) == "{}" { // empty struct
			dst.YaraOutput = nil
		} else {
			return nil // data stored in dst.YaraOutput, return on the first match
		}
	} else {
		dst.YaraOutput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IndicatorsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IndicatorsInner) MarshalJSON() ([]byte, error) {
	if src.ForensicArtifactOutput != nil {
		return json.Marshal(&src.ForensicArtifactOutput)
	}

	if src.QueryOutput != nil {
		return json.Marshal(&src.QueryOutput)
	}

	if src.RegexOutput != nil {
		return json.Marshal(&src.RegexOutput)
	}

	if src.SigmaOutput != nil {
		return json.Marshal(&src.SigmaOutput)
	}

	if src.SuricataOutput != nil {
		return json.Marshal(&src.SuricataOutput)
	}

	if src.YaraOutput != nil {
		return json.Marshal(&src.YaraOutput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIndicatorsInner struct {
	value *IndicatorsInner
	isSet bool
}

func (v NullableIndicatorsInner) Get() *IndicatorsInner {
	return v.value
}

func (v *NullableIndicatorsInner) Set(val *IndicatorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorsInner(val *IndicatorsInner) *NullableIndicatorsInner {
	return &NullableIndicatorsInner{value: val, isSet: true}
}

func (v NullableIndicatorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
