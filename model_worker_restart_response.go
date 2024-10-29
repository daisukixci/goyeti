/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WorkerRestartResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkerRestartResponse{}

// WorkerRestartResponse Worker restart API response.
type WorkerRestartResponse struct {
	Successes []string `json:"successes"`
	Failures  []string `json:"failures"`
}

type _WorkerRestartResponse WorkerRestartResponse

// NewWorkerRestartResponse instantiates a new WorkerRestartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerRestartResponse(successes []string, failures []string) *WorkerRestartResponse {
	this := WorkerRestartResponse{}
	this.Successes = successes
	this.Failures = failures
	return &this
}

// NewWorkerRestartResponseWithDefaults instantiates a new WorkerRestartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerRestartResponseWithDefaults() *WorkerRestartResponse {
	this := WorkerRestartResponse{}
	return &this
}

// GetSuccesses returns the Successes field value
func (o *WorkerRestartResponse) GetSuccesses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Successes
}

// GetSuccessesOk returns a tuple with the Successes field value
// and a boolean to check if the value has been set.
func (o *WorkerRestartResponse) GetSuccessesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Successes, true
}

// SetSuccesses sets field value
func (o *WorkerRestartResponse) SetSuccesses(v []string) {
	o.Successes = v
}

// GetFailures returns the Failures field value
func (o *WorkerRestartResponse) GetFailures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value
// and a boolean to check if the value has been set.
func (o *WorkerRestartResponse) GetFailuresOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Failures, true
}

// SetFailures sets field value
func (o *WorkerRestartResponse) SetFailures(v []string) {
	o.Failures = v
}

func (o WorkerRestartResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerRestartResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["successes"] = o.Successes
	toSerialize["failures"] = o.Failures
	return toSerialize, nil
}

func (o *WorkerRestartResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"successes",
		"failures",
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

	varWorkerRestartResponse := _WorkerRestartResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkerRestartResponse)

	if err != nil {
		return err
	}

	*o = WorkerRestartResponse(varWorkerRestartResponse)

	return err
}

type NullableWorkerRestartResponse struct {
	value *WorkerRestartResponse
	isSet bool
}

func (v NullableWorkerRestartResponse) Get() *WorkerRestartResponse {
	return v.value
}

func (v *NullableWorkerRestartResponse) Set(val *WorkerRestartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerRestartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerRestartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerRestartResponse(val *WorkerRestartResponse) *NullableWorkerRestartResponse {
	return &NullableWorkerRestartResponse{value: val, isSet: true}
}

func (v NullableWorkerRestartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerRestartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}