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

// checks if the PatchObservableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchObservableRequest{}

// PatchObservableRequest struct for PatchObservableRequest
type PatchObservableRequest struct {
	Observable           Observable1 `json:"observable"`
	AdditionalProperties map[string]interface{}
}

type _PatchObservableRequest PatchObservableRequest

// NewPatchObservableRequest instantiates a new PatchObservableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchObservableRequest(observable Observable1) *PatchObservableRequest {
	this := PatchObservableRequest{}
	this.Observable = observable
	return &this
}

// NewPatchObservableRequestWithDefaults instantiates a new PatchObservableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchObservableRequestWithDefaults() *PatchObservableRequest {
	this := PatchObservableRequest{}
	return &this
}

// GetObservable returns the Observable field value
func (o *PatchObservableRequest) GetObservable() Observable1 {
	if o == nil {
		var ret Observable1
		return ret
	}

	return o.Observable
}

// GetObservableOk returns a tuple with the Observable field value
// and a boolean to check if the value has been set.
func (o *PatchObservableRequest) GetObservableOk() (*Observable1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observable, true
}

// SetObservable sets field value
func (o *PatchObservableRequest) SetObservable(v Observable1) {
	o.Observable = v
}

func (o PatchObservableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchObservableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["observable"] = o.Observable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchObservableRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"observable",
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

	varPatchObservableRequest := _PatchObservableRequest{}

	err = json.Unmarshal(data, &varPatchObservableRequest)

	if err != nil {
		return err
	}

	*o = PatchObservableRequest(varPatchObservableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "observable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchObservableRequest struct {
	value *PatchObservableRequest
	isSet bool
}

func (v NullablePatchObservableRequest) Get() *PatchObservableRequest {
	return v.value
}

func (v *NullablePatchObservableRequest) Set(val *PatchObservableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchObservableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchObservableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchObservableRequest(val *PatchObservableRequest) *NullablePatchObservableRequest {
	return &NullablePatchObservableRequest{value: val, isSet: true}
}

func (v NullablePatchObservableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchObservableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}